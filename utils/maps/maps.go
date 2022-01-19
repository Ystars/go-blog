package maps

type maps struct{}

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

var MapsStruct = new(maps)
