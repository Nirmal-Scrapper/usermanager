package json

import "encoding/json"

type JSON struct {
	Jmap map[string]interface{}
}

//String to byte to json
func ParseString(data string) JSON {
	return Parse([]byte(data))
}

//byte to kind of json type(map)
func Parse(data []byte) (parsed JSON) {
	var f interface{}
	if err := json.Unmarshal(data, &f); err != nil {
		return
	}
	switch f.(type) {
	case []interface{}:
		//      log.Println("Found Array, Not Supported")
		return
	default:
		m := f.(map[string]interface{})
		parsed.Jmap = m
		return parsed
	}
}

//get value by key
func (jobj *JSON) GetString(k string) string {
	data := jobj.Jmap[k]
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}
