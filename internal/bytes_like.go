package internal

import (
	"errors"
)

type BencodexBytesLike struct {
	isBytes bool
	raw     string
}

func (b *BencodexBytesLike) MustAsString() string {
	if b.isBytes {
		panic("It is not string.")
	}

	return b.raw
}

func (b *BencodexBytesLike) MustAsBytes() []byte {
	if !b.isBytes {
		panic("It is not bytes.")
	}

	return S2B(b.raw)
}

func (b *BencodexBytesLike) AsString() (string, error) {
	if b.isBytes {
		return "", errors.New("It is bytes.")
	}

	return b.raw, nil
}

func (b *BencodexBytesLike) IsString() bool {
	return !b.isBytes
}

func (b *BencodexBytesLike) IsBytes() bool {
	return b.isBytes
}

func (b *BencodexBytesLike) AsBytes() ([]byte, error) {
	if !b.isBytes {
		return nil, errors.New("It is string.")
	}

	return S2B(b.raw), nil
}

func (b *BencodexBytesLike) Len() int {
	return len(b.raw)
}

func NewBytes(b []byte) BencodexBytesLike {
	return BencodexBytesLike{
		isBytes: true,
		raw:     B2S(b),
	}
}

func NewString(s string) BencodexBytesLike {
	return BencodexBytesLike{
		isBytes: false,
		raw:     s,
	}
}
