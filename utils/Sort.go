package utils

import (
	"reflect"
	"sort"
)

//通用排序
//结构体排序，必须重写数组Len() Swap() Less()函数
type arr_wrapper struct {
	Arr []interface{}
	by  func(p, q *interface{}) bool //内部Less()函数会用到
}

//数组长度Len()
func (e arr_wrapper) Len() int {
	return len(e.Arr)
}

//元素交换
func (e arr_wrapper) Swap(i, j int) {
	e.Arr[i], e.Arr[j] = e.Arr[j], e.Arr[i]
}

//比较函数，使用外部传入的by比较函数
func (e arr_wrapper) Less(i, j int) bool {
	return e.by(&e.Arr[i], &e.Arr[j])
}

//自定义排序字段，参考SortBodyByCreateTime中的传入函数
func SortInterface(arr []interface{}, by func(p, q *interface{}) bool) {
	sort.Sort(arr_wrapper{arr, by})
}

func SortArrMap(maps []map[string]interface{}, by func(a, b map[string]interface{}) bool) {
	arr := []interface{}{}
	for _, tmp := range maps {
		arr = append(arr, tmp)
	}
	sort.Sort(arr_wrapper{arr, func(p, q *interface{}) bool {
		a := (*p).(map[string]interface{})
		b := (*q).(map[string]interface{})
		return by(a, b)
	}})
	for i := 0; i < len(arr); i++ {
		maps[i] = arr[i].(map[string]interface{})
	}
}

//按照createtime排序，需要注意是否有createtime
func SortBodyByCreateTime(arr []interface{}) {
	sort.Sort(arr_wrapper{arr, func(p, q *interface{}) bool {
		v := reflect.ValueOf(*p)
		i := v.FieldByName("Create_time")
		v = reflect.ValueOf(*q)
		j := v.FieldByName("Create_time")
		return i.String() > j.String()
	}})
}
