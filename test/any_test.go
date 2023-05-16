package test

import (
	"encoding/json"
	"github.com/luvcurt/zephyr/bytezip"
	"math/rand"
	"sync/atomic"
	"testing"
)

func TestAny(t *testing.T) {}

func BenchmarkByteZip(b *testing.B) {

	type Player struct {
		Pid  int64 `json:"pid,omitempty"`
		PosX int64 `json:"posX,omitempty"`
		PosY int64 `json:"posY,omitempty"`
		PosZ int64 `json:"posZ,omitempty"`
		PosV int64 `json:"posV,omitempty"`
	}

	baseUid := int64(10000)

	for i := 0; i < b.N; i++ {
		ps := make([]Player, 0, 10000)
		for i := 0; i < 10000; i++ {
			ps = append(ps, Player{
				Pid:  atomic.AddInt64(&baseUid, 1),
				PosX: int64(rand.Intn(50000)),
				PosY: int64(rand.Intn(50000)),
				PosZ: int64(rand.Intn(50000)),
				PosV: int64(rand.Intn(36000)),
			})
		}
		bytes, _ := json.Marshal(&ps)
		_, _ = bytezip.GZipBytes(bytes, bytezip.DefaultCompression)
	}
}
