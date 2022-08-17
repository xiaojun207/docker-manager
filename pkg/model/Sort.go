package model

type SortInt struct {
	Arr      []string
	ValueMap SyncMap
}

func (e SortInt) Len() int      { return len(e.Arr) }
func (e SortInt) Swap(i, j int) { e.Arr[i], e.Arr[j] = e.Arr[j], e.Arr[i] }
func (e SortInt) Less(i, j int) bool {
	s := e.Arr
	return e.ValueMap.GetInt(s[i]) < e.ValueMap.GetInt(s[j])
}
