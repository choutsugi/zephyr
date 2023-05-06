package rotatelog

import (
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type RotateLog struct {
	logPath  string
	file     *os.File
	fileSize int64
	mutex    *sync.Mutex
	rotate   <-chan time.Time // notify rotate event
	close    chan struct{}    // close file and write goroutine
	opts     Opts
}

func NewRotateLog(logPath string, opts ...OptFunc) (*RotateLog, error) {

	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	rl := &RotateLog{
		logPath: logPath,
		mutex:   &sync.Mutex{},
		close:   make(chan struct{}, 1),
		opts:    o,
	}

	if err := os.Mkdir(filepath.Dir(rl.logPath), 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	if err := rl.rotateFileByTime(time.Now()); err != nil {
		return nil, err
	}

	if rl.opts.rotateTime != 0 {
		go rl.handleEvent()
	}

	return rl, nil
}

func (r *RotateLog) Write(bytes []byte) (int, error) {

	writeLen := int64(len(bytes))
	if writeLen > r.opts.maxFileSize {
		return 0, errors.Errorf("write length %d exceeds max file size %d", writeLen, r.opts.maxFileSize)
	}
	if r.fileSize+writeLen > r.opts.maxFileSize {
		if err := r.rotateFileBySize(); err != nil {
			return 0, err
		}
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	n, err := r.file.Write(bytes)
	r.fileSize += int64(n)
	return n, err
}

func (r *RotateLog) Close() error {
	r.close <- struct{}{}
	return r.file.Close()
}

func (r *RotateLog) handleEvent() {
	for {
		select {
		case <-r.close:
			return
		case now := <-r.rotate:
			_ = r.rotateFileByTime(now)
		}
	}
}

func (r *RotateLog) rotateFileByTime(now time.Time) error {
	if r.opts.rotateTime != 0 {
		nextRotateTime := r.calRotateTimeDuration(now, r.opts.rotateTime)
		r.rotate = time.After(nextRotateTime)
	}

	latestLogPath := r.getLatestLogPath(now)
	r.mutex.Lock()
	defer r.mutex.Unlock()
	file, err := os.OpenFile(latestLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if r.file != nil {
		_ = r.file.Close()
	}
	r.file = file
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	r.fileSize = stat.Size()

	if len(r.opts.curLogLinkPath) > 0 {
		_ = os.Remove(r.opts.curLogLinkPath)
		_ = os.Link(latestLogPath, r.opts.curLogLinkPath)
	}

	if r.opts.maxAge > 0 && len(r.opts.delFileWildcard) > 0 {
		go r.deleteExpiredFile(now)
	}

	return nil
}

func (r *RotateLog) rotateFileBySize() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.file.Close()

	path, err := filepath.Abs(filepath.Dir(r.file.Name()))
	if err != nil {
		return err
	}
	filename := filepath.Base(r.file.Name())
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	files := make([]fs.FileInfo, 0)
	if err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), prefix) {
			files = append(files, info)
		}
		return nil
	}); err != nil {
		return err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds() <
			files[j].Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds()
	})

	for _, file := range files {
		oldName := filepath.Join(path, file.Name())
		ext := filepath.Ext(file.Name())
		if ext == "" || len(ext) < 2 {
			continue
		}
		fileNo, err := strconv.ParseInt(ext[1:], 10, 64)
		if err != nil {
			return err
		}
		newExt := ext[:1] + strconv.FormatInt(fileNo+1, 10)
		newName := filepath.Join(path, prefix+newExt)
		if err := os.Rename(oldName, newName); err != nil {
			return err
		}
	}

	latestLogPath := r.getLatestLogPath(time.Now())
	file, err := os.OpenFile(latestLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if r.file != nil {
		_ = r.file.Close()
	}
	r.file = file
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	r.fileSize = stat.Size()

	if len(r.opts.curLogLinkPath) > 0 {
		_ = os.Remove(r.opts.curLogLinkPath)
		_ = os.Link(latestLogPath, r.opts.curLogLinkPath)
	}

	if r.opts.maxAge > 0 && len(r.opts.delFileWildcard) > 0 {
		go r.deleteExpiredFile(time.Now())
	}

	return nil
}

// Judge expired by last modify time
func (r *RotateLog) deleteExpiredFile(now time.Time) {
	cutoffTime := now.Add(-r.opts.maxAge)
	matches, err := filepath.Glob(r.opts.delFileWildcard)
	if err != nil {
		return
	}

	toUnlink := make([]string, 0, len(matches))
	for _, path := range matches {
		fileInfo, err := os.Stat(path)
		if err != nil {
			continue
		}

		if r.opts.maxAge > 0 && fileInfo.ModTime().After(cutoffTime) {
			continue
		}

		if len(r.opts.curLogLinkPath) > 0 && fileInfo.Name() == filepath.Base(r.opts.curLogLinkPath) {
			continue
		}
		toUnlink = append(toUnlink, path)
	}

	for _, path := range toUnlink {
		_ = os.Remove(path)
	}
}

func (r *RotateLog) getLatestLogPath(t time.Time) string {
	return fmt.Sprintf("%s.0", t.Format(r.logPath))
}

func (r *RotateLog) calRotateTimeDuration(now time.Time, duration time.Duration) time.Duration {
	nowUnixNao := now.UnixNano()
	NanoSecond := duration.Nanoseconds()
	nextRotateTime := NanoSecond - (nowUnixNao % NanoSecond)
	return time.Duration(nextRotateTime)
}
