package encoder

import bencodex "github.com/moreal/bencodex-go/internal"

//go:nosplit
func (e *Encoder) encodeBytesLike(data bencodex.BencodexBytesLike) {
	dataLength := data.Len()
	e.grow(dataLength + 23)
	if data.IsString() {
		e.writeByte('u')
	}
	e.writeInt(int64(dataLength))
	e.writeByte(':')
	if data.IsBytes() {
		e.write(data.MustAsBytes())
	} else {
		e.write(bencodex.S2B(data.MustAsString()))
	}
}
