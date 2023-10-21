package encoder

import (
	"sync"

	"github.com/moreal/bencodex-go/internal"
)

const bytesLikeArrayLen = 20

var stringsArrayPool = sync.Pool{
	New: func() interface{} {
		return &[bytesLikeArrayLen]internal.BencodexBytesLike{}
	},
}

// -1 = left go first, 0 = equal, 1 = right go first
func compareBytes(x, y []byte) int {
	var minLen int
	if len(x) > len(y) {
		minLen = len(y)
	} else {
		minLen = len(x)
	}

	for i := 0; i < minLen; i++ {
		if x[i] > y[i] {
			return 1
		} else if x[i] < y[i] {
			return -1
		}
	}

	if len(x) == len(y) {
		return 0
	}

	if len(x) > len(y) {
		return 1
	}

	return -1
}

func sortKeys(keys []internal.BencodexBytesLike) {
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0; j-- {
			if keys[j].IsString() && keys[j-1].IsBytes() {
				break
			} else if keys[j].IsBytes() && keys[j-1].IsString() {
				keys[j], keys[j-1] = keys[j-1], keys[j]
			} else if keys[j].IsBytes() && keys[j-1].IsBytes() {
				res := compareBytes(keys[j-1].MustAsBytes(), keys[j].MustAsBytes())
				if res >= 0 {
					break
				}
			} else if keys[j].MustAsString() >= keys[j-1].MustAsString() {
				break
			}

			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
}

func (e *Encoder) encodeDictionary(data map[internal.BencodexBytesLike]interface{}) error {
	e.grow(1)
	e.writeByte('d')
	var keys []internal.BencodexBytesLike
	if len(data) <= bytesLikeArrayLen {
		stringsArray := stringsArrayPool.Get().(*[bytesLikeArrayLen]internal.BencodexBytesLike)
		defer stringsArrayPool.Put(stringsArray)
		keys = stringsArray[:0:len(data)]
	} else {
		keys = make([]internal.BencodexBytesLike, 0, len(data))
	}
	for key, _ := range data {
		keys = append(keys, key)
	}
	sortKeys(keys)
	for _, key := range keys {
		e.encodeBytesLike(key)
		err := e.encode(data[key])
		if err != nil {
			return err
		}
	}
	e.grow(1)
	e.writeByte('e')
	return nil
}
