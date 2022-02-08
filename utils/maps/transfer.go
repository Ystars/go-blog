package maps

import "encoding/json"

type Transfer struct{}

// MapToJson map转换json
func (t *Transfer) MapToJson(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

// JsonToMap json转map
func (t *Transfer) JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	_ = json.Unmarshal([]byte(str), &tempMap)
	return tempMap
}

func (t *Transfer) ErrorJsonToMap(err error) interface{} {
	msg := err.Error()
	if json.Valid([]byte(msg)) {
		return t.JsonToMap(msg)
	}
	return msg
}
