package rotatelog

import (
	"time"
)

type OptFunc func(opts *Opts)

type Opts struct {
	curLogLinkPath  string
	delFileWildcard string
	rotateTime      time.Duration
	maxAge          time.Duration
	maxFileSize     int64
}

func defaultOpts() Opts {
	return Opts{
		rotateTime:  time.Hour * 24,
		maxAge:      time.Hour * 24 * 7,
		maxFileSize: 1024 * 1024 * 50,
	}
}

func WithRotateTime(duration time.Duration) OptFunc {
	return func(o *Opts) {
		o.rotateTime = duration
	}
}

func WithMaxFileSize(size int64) OptFunc {
	return func(o *Opts) {
		o.maxFileSize = size
	}
}

func WithCurLogLinkPath(linkPath string) OptFunc {
	return func(o *Opts) {
		o.curLogLinkPath = linkPath
	}
}

// WithDeleteExpiredFile Judge expired by last modify time
// cutoffTime = now - maxAge
// Only delete satisfying file wildcard filename
func WithDeleteExpiredFile(maxAge time.Duration, fileWildcard string) OptFunc {
	return func(o *Opts) {
		o.maxAge = maxAge
		o.delFileWildcard = fileWildcard
	}
}
