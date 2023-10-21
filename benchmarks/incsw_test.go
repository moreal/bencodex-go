package benchmarks

import (
	"testing"

	bencode "github.com/moreal/bencodex-go"
	"github.com/stretchr/testify/assert"
)

func Benchmark_Incsw_Marshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buffer, err = bencodex.Marshal(bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Incsw_MarshalTo(b *testing.B) {
	b.ReportAllocs()
	buffer = make([]byte, 512)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer, err = bencodex.MarshalTo(buffer, bytesInt64TestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, string(unmarshalTestData), string(buffer))
}

func Benchmark_Incsw_Unmarshal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		torrent, err = bencodex.Unmarshal(unmarshalTestData)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	assert.Equal(b, bytesInt64TestData, torrent)
}

func Benchmark_Incsw_RealWorld(b *testing.B) {
	b.ReportAllocs()
	b.Run("Unmarshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			torrent, err = bencodex.Unmarshal(realWorldData)
			if err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
	b.Run("Marshal", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencodex.Marshal(torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("MarshalTo", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			buffer, err = bencodex.MarshalTo(buffer, torrent)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	assert.Equal(b, string(realWorldData), string(buffer))
}
