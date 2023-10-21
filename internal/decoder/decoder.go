package decoder

import (
	"errors"
	bencodex "github.com/moreal/bencodex-go/internal"
)

type Decoder struct {
	data   []byte
	length int
	cursor int
}

func (d *Decoder) Decode(data []byte) (interface{}, error) {
	d.data = data
	d.length = len(data)
	return d.decode()
}

func (d *Decoder) decode() (interface{}, error) {
	switch d.data[d.cursor] {
	case 'i':
		return d.decodeInt()
	case 'n':
		d.cursor += 1
		return nil, nil
	case 't':
		d.cursor += 1
		return true, nil
	case 'f':
		d.cursor += 1
		return false, nil
	case 'l':
		d.cursor += 1
		list := []interface{}{}
		for {
			if d.cursor == d.length {
				return nil, errors.New("bencode: invalid list field")
			}
			if d.data[d.cursor] == 'e' {
				d.cursor += 1
				return list, nil
			}
			value, err := d.decode()
			if err != nil {
				return nil, err
			}
			list = append(list, value)
		}
	case 'd':
		d.cursor += 1
		dictionary := map[bencodex.BencodexBytesLike]interface{}{}
		for {
			if d.cursor == d.length {
				return nil, errors.New("bencode: invalid dictionary field")
			}
			if d.data[d.cursor] == 'e' {
				d.cursor += 1
				return dictionary, nil
			}
			var (
				key bencodex.BencodexBytesLike
				err error
			)
			if '0' <= d.data[d.cursor] && d.data[d.cursor] <= '9' {
				key, err = d.decodeBytes(false)
			}

			if d.data[d.cursor] == 'u' {
				d.cursor += 1
				key, err = d.decodeBytes(true)
			}

			if err != nil {
				return nil, errors.New("bencode: non-string dictionary key")
			}
			value, err := d.decode()
			if err != nil {
				return nil, err
			}
			dictionary[key] = value
		}
	case 'u':
		d.cursor += 1
		unicode, err := d.decodeBytes(true)
		if err != nil {
			return nil, err
		}

		return unicode, nil
	default:
		return d.decodeBytes(false)
	}
}
