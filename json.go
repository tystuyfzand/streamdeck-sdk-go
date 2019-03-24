package sdk

import "github.com/valyala/fastjson"

func JsonStringValue(v *fastjson.Value, keys ...string) string {
	value := v.Get(keys...)

	if value == nil {
		return ""
	}

	strB, err := value.StringBytes()

	if err != nil {
		return ""
	}

	return string(strB)
}

