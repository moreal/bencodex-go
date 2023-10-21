package bencode

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/moreal/bencodex-go/internal"
)

func convertBencodexKeyToJson(x interface{}) (string, error) {
	switch v := x.(type) {
	case string:
		return fmt.Sprintf("\uFEFF%s", v), nil
	case internal.BencodexBytesLike:
		if v.IsString() {
			return fmt.Sprintf("\uFEFF%s", v.MustAsString()), nil
		}

		return fmt.Sprintf("0x%s", hex.EncodeToString(v.MustAsBytes())), nil
	}

	return "", errors.New("convertBencodexKeyToJson: not supported type.")
}

func ConvertToBencodexJson(x interface{}) (interface{}, error) {
	switch v := x.(type) {
	case string:
		return fmt.Sprintf("\uFEFF%s", v), nil
	case internal.BencodexBytesLike:
		if v.IsString() {
			return fmt.Sprintf("\uFEFF%s", v.MustAsString()), nil
		}

		return fmt.Sprintf("0x%s", hex.EncodeToString(v.MustAsBytes())), nil
	case int64:
		return v, nil
	case bool:
		return v, nil
	case nil:
		return v, nil
	case []interface{}:
		result := make([]interface{}, len(v))
		for i := 0; i < len(v); i++ {
			newVal, err := ConvertToBencodexJson(v[i])
			if err != nil {
				return nil, err
			}

			result[i] = newVal
		}

		return result, nil
	case map[internal.BencodexBytesLike]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			newKey, err := convertBencodexKeyToJson(key)
			if err != nil {
				return nil, err
			}

			newVal, err := ConvertToBencodexJson(value)
			if err != nil {
				return nil, err
			}

			result[newKey] = newVal
		}

		return result, nil
	}

	return nil, errors.New("ConvertToBencodexJson: not supported type.")
}
