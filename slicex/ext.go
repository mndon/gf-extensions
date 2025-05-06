package slicex

import "reflect"

// InSlice 一个元素是否存在于切片中
// Deprecated: 使用Contains替代
func InSlice[T comparable](val T, slices []T) bool {
	return Contains(slices, val)
}

// InSliceAny 一个元素是否存在于切片中，参数为interface，可以比较复杂结构
// 注意下面这个坑点！
//   - InSliceAny(float64(0), []float64{1, 0.0}) => false
//   - InSliceAny(0, []float64{1, 0.0}) => true
func InSliceAny(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	//case reflect.Map:
	//	for _, k := range val.MapKeys() {
	//		if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
	//			return true
	//		}
	//	}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

// GroupBy 按指定字段对切片 s 进行分组。
// fn 是一个函数，用于从切片元素中提取分组字段。
// 返回值是一个映射，键是字段值，值是属于该字段值的元素切片。
// 示例：
// people := []Person{{Name: "Alice", ClassID: 25}, {Name: "Bob", ClassID: 25}, {Name: "Charlie", ClassID: 30}}
// grouped := GroupBy(people, func(p Person) int { return p.ClassID })
// fmt.Println(grouped) // map[25:[{Alice 25} {Bob 25}] 30:[{Charlie 30}]]
func GroupBy[T any, T2 comparable](s []T, fn func(t T) T2) map[T2][]T {
	m := make(map[T2][]T)
	for i := range s {
		m[fn(s[i])] = append(m[fn(s[i])], s[i])
	}
	return m
}

// GroupByAndFlatten 按指定字段对切片 s 进行分组，并展平，只保留最后一个匹配的元素。
// fn 是一个函数，用于从切片元素中提取分组字段。
// 返回值是一个映射，键是字段值，值是属于该字段值的元素。
// 示例：
// people := []Person{{Name: "Bob", ID: 25}, {Name: "Charlie", ID: 30}}
// grouped := GroupByAndFlatten(people, func(p Person) int { return p.ID })
// fmt.Println(grouped) // map[25:{Bob 25} 30:{Charlie 30}]
func GroupByAndFlatten[T any, T2 comparable](s []T, fn func(t T) T2) map[T2]T {
	m := make(map[T2]T)
	for i := range s {
		m[fn(s[i])] = s[i]
	}
	return m
}

// Map 对切片 s 的每个元素应用函数 fn 并返回新切片。
// 示例：
// nums := []int{1, 2, 3}
// doubled := Map(nums, func(n int) int { return n * 2 })
// fmt.Println(doubled) // [2, 4, 6]
func Map[T any, K any](s []T, fn func(item T) K) []K {
	k := make([]K, len(s))
	for i := range s {
		k[i] = fn(s[i])
	}
	return k
}

// Every 检查切片 s 的所有元素是否都满足 fn 函数。
// 如果所有元素都满足，则返回 true，否则返回 false。
// 示例：
// nums := []int{2, 4, 6}
// allEven := Every(nums, func(n int) bool { return n%2 == 0 })
// fmt.Println(allEven) // true
func Every[T any](s []T, fn func(item T) bool) bool {
	for i := range s {
		if !fn(s[i]) {
			return false
		}
	}
	return true
}

// Any 检查切片 s 中是否至少有一个元素满足 fn 函数。
// 如果存在满足条件的元素，则返回 true，否则返回 false。
// 示例：
// nums := []int{1, 3, 5, 6}
// hasEven := Any(nums, func(n int) bool { return n%2 == 0 })
// fmt.Println(hasEven) // true
func Any[T any](s []T, fn func(item T) bool) bool {
	for i := range s {
		if fn(s[i]) {
			return true
		}
	}
	return false
}

// Union 合并多个切片并去重，返回一个新的切片。
// 示例：
// s1 := []int{1, 2, 3}
// s2 := []int{3, 4, 5}
// unionList := Union(s1, s2)
// fmt.Println(unionList) // [1, 2, 3, 4, 5] （顺序可能不同）
func Union[T comparable](s ...[]T) []T {
	if len(s) == 0 {
		return []T{}
	}
	hash := make(map[T]struct{})
	for i := range s {
		for j := range s[i] {
			hash[s[i][j]] = struct{}{}
		}
	}
	output := make([]T, len(hash))
	i := 0
	for k := range hash {
		output[i] = k
		i++
	}
	return output
}
