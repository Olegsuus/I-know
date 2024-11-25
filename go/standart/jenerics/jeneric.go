package main

import "fmt"

//func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}
//
//func main() {
//	mi := map[string]int64{"a": 5, "b": 3, "c": 2}
//	ri := SumIntsOrFloats(mi)
//	println(ri)
//
//	mf := map[string]float64{"a": 5.8, "b": 3.1, "c": 2.11}
//	rf := SumIntsOrFloats(mf)
//	fmt.Printf("%v\n", rf)
//	//Проблема точности с плавающей
//	mf2 := map[string]float64{"a": 5.1, "b": 3.1, "c": 2.1}
//	rf = SumIntsOrFloats(mf2)
//	fmt.Printf("%v\n", rf)
//}
//
//

func main() {
	var i int = 3

	i = i << 1
	fmt.Println(i)
}
