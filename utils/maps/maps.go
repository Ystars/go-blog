package maps

type maps struct{}

func NewMaps() *maps {
	return &maps{}
}

// MergeIntString 合并整型字符串
func (t *maps) MergeIntString(mObj ...map[int]string) map[int]string {
	newObj := map[int]string{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

// MergeStringInterface 合并万能类型
func (t *maps) MergeStringInterface(mObj ...map[string]interface{}) map[string]interface{} {
	newObj := map[string]interface{}{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}
