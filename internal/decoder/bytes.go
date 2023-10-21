package decoder

import (
	"bytes"
	"errors"
	"fmt"
	bencodex "github.com/moreal/bencodex-go/internal"
)

func (d *Decoder) decodeBytes(isUnicode bool) (bencodex.BencodexBytesLike, error) {
	if d.data[d.cursor] < '0' || d.data[d.cursor] > '9' {
		return bencodex.BencodexBytesLike{}, errors.New(fmt.Sprintf("bencode: invalid string field '%c', expected '0'~'9'", d.data[d.cursor]))
	}
	index := bytes.IndexByte(d.data[d.cursor:], ':')
	if index == -1 {
		return bencodex.BencodexBytesLike{}, errors.New(fmt.Sprintf("bencode: invalid string field '%c', expected ':'", d.data[d.cursor]))
	}
	index += d.cursor
	stringLength, err := d.parseInt(d.data[d.cursor:index])
	if err != nil {
		return bencodex.BencodexBytesLike{}, err
	}
	index += 1
	endIndex := index + int(stringLength)
	if endIndex > d.length {
		return bencodex.BencodexBytesLike{}, errors.New("bencode: not a valid bencoded string")
	}
	value := d.data[index:endIndex]
	d.cursor = endIndex
	if isUnicode {
		return bencodex.NewString(bencodex.B2S(value)), nil
	}

	return bencodex.NewBytes(value), nil
}
