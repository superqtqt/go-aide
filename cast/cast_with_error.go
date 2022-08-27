package cast

import (
	"encoding/json"
	"fmt"
	"github.com/superqtqt/go-aide/reflect_aid"
	"strconv"
	"strings"
)

func ToIntWithError(v interface{}, opt ...func(options *Options)) (int, error) {
	t := reflect_aid.Indirect(v)

	switch s := t.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		options := &Options{}
		for i := range opt {
			opt[i](options)
		}
		castV, err := strconv.ParseInt(strings.ReplaceAll(s, " \n\t", ""), options.base, options.bitSize)
		if err == nil {
			return int(castV), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", v, t)
	case json.Number:
		return ToIntWithError(string(s))
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", v, t)
	}
}

func ToFloatWithError(v interface{}, opt ...func(options *Options)) (float64, error) {
	t := reflect_aid.Indirect(v)

	switch s := t.(type) {
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case string:
		options := &Options{}
		for i := range opt {
			opt[i](options)
		}
		castV, err := strconv.ParseFloat(strings.ReplaceAll(s, " \n\t", ""), options.bitSize)
		if err == nil {
			return castV, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float", v, t)
	case json.Number:
		return ToFloatWithError(string(s))
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float", v, t)
	}
}
