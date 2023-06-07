//*****************ADIM 1**************

// package main

// import (
// 	"fmt"

// 	"github.com/ahmetb/go-linq"
// )

// func main() {
// 	// Bir slice oluşturalım
// 	numbers := []int{1, 2, 3, 4, 5}

// 	// LINQ benzeri işlemler yapalım
// 	linq.From(numbers).
// 		Where(func(i interface{}) bool { return i.(int) > 2 }).
// 		Select(func(i interface{}) interface{} { return i.(int) * 2 }).
// 		ToSlice(&numbers)

// 	// Sonuçları yazdıralım
// 	fmt.Println(numbers)

// }

//----------->> {Where, Map, Sort, GroupBy, Join} -> temel LINQ sorguları

//**********************ADIM 2**************

// package main

// import (
// 	"fmt"
// 	//  "github.com/ahmetb/go-linq"
// )

// func filter(numbers []int, f func(int) bool) []int {
// 	result := []int{}
// 	for _, n := range numbers {
// 		if f(n) {
// 			result = append(result, n)
// 		}
// 	}
// 	return result
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6}
// 	filterNumbers := filter(numbers, func(n int) bool {
// 		return n > 2
// 	})

// 	fmt.Print(filterNumbers)
// }

//****************ADIM 3**********************

// package main

// import "fmt"

// func mapSlice(numbers []int, f func(int) int) []int {
// 	result := []int{}
// 	for _, n := range numbers {
// 		mathresult := f(n)
// 		result = append(result, mathresult)

// 	}
// 	return result
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6}
// 	newSlice := mapSlice(numbers, func(n int) int {
// 		return n * 2
// 	})

// 	fmt.Println(newSlice)
// }

//****************ADIM 4*****************

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	numbers := []int{5, 4, 6, 2, 1, 3}
// 	sort.Slice(numbers, func(i, j int) bool {
// 		return numbers[i] < numbers[j]
// 	})
// 	fmt.Print(numbers)
// }

//*********************************************************************

// package main

// import "fmt"

// func main() {
// 	words := []string{"apple", "banana", "kiwi", "cherry"}
// 	groupedWords := groupBy(words, func(s string) rune {
// 		return rune(s[0])
// 	})

// 	fmt.Println(groupedWords)
// }

// func groupBy(words []string, key func(string) rune) map[rune]string {
// 	groups := make(map[rune]string)

// 	for _, v := range words {
// 		tempKey := key(v)
// 		groups[tempKey] = append(groups[tempKey], v)
// 	}
// 	return groups
// }

// func groupBy(words []string, key func(string) rune) map[rune]string {
// 	groups := make(map[rune]string)

// 	// for _, v := range words {
// 	// 	tempKey := key(v)
// 	// 	groups[tempKey] = append(groups[tempKey], v)
// 	// }

// 	for i:=0;i<len(words);i++{
// tempKey:=key(words[i])
// groups[tempKey]=append(groups[tempKey],words[i])
// 	}
// 	return groups
// }

//***********************ADIM 5******************************

// package main

// import "fmt"

// func main() {
// 	words := []string{"apple", "banana", "kiwi", "cherry"}
// 	groupedWords := groupBy(words, func(s string) rune {
// 		return rune(s[0])
// 	})

// 	fmt.Println(groupedWords)
// }

// func groupBy(words []string, key func(string) rune) map[rune][]string {
// 	groups := make(map[rune][]string)

// 	for _, v := range words {
// 		tempKey := key(v)
// 		groups[tempKey] = append(groups[tempKey], v)
// 	}
// 	return groups
// }

//********************************************************************

// package main

// type student struct {
// 	Name string
// 	ID   int
// }

// type grade struct {
// 	ID    int
// 	Grade string
// }

// func main() {
// 	students := []student{
// 		{Name: "John", ID: 1},
// 		{Name: "Mary", ID: 2},
// 		{Name: "Steve", ID: 3},
// 	}

// 	grades := []grade{
// 		{ID: 1, Grade: "A"},
// 		{ID: 2, Grade: "B"},
// 		{ID: 3, Grade: "C"},
// 	}

// 	joinedData := join(students, grades, func(s student) interface{} {
// 		return s.ID
// 	}, func(g grade) interface{} {
// 		return g.ID
// 	}, func(s student, g grade) interface{} {
// 		return struct {
// 			Name  string
// 			Grade string
// 		}{
// 			Name:  s.Name,
// 			Grade: g.Grade,
// 		}
// 	},
// 	)
// 	// joinedData = [{John A} {Mary B} {Steve C}]
// }

//*********************************ADIM 6***********************************************

package main

import "fmt"

type student struct {
	Name string
	ID   int
}

type grade struct {
	ID    int
	Grade string
}

func main() {
	students := []student{
		{Name: "John", ID: 1},
		{Name: "Mary", ID: 2},
		{Name: "Steve", ID: 3},
	}

	grades := []grade{
		{ID: 1, Grade: "A"},
		{ID: 2, Grade: "B"},
		{ID: 3, Grade: "C"},
	}

	joinedData := join(students, grades, func(s interface{}) interface{} {
		return s.(student).ID
	}, func(g interface{}) interface{} {
		return g.(grade).ID
	}, func(s interface{}, g interface{}) interface{} {
		return struct {
			Name  string
			Grade string
		}{
			Name:  s.(student).Name,
			Grade: g.(grade).Grade,
		}
	},
	)
	// joinedData = [{John A} {Mary B} {Steve C}]

	fmt.Print(joinedData...)
}

func join(table1 interface{}, table2 interface{}, keyFunc1 func(interface{}) interface{}, keyFunc2 func(interface{}) interface{}, resultFunc func(interface{}, interface{}) interface{}) []interface{} {
	m1 := make(map[interface{}]interface{})
	m2 := make(map[interface{}]interface{})

	// Table 1
	for _, v1 := range table1.([]student) {
		m1[keyFunc1(v1)] = v1
	}

	// Table 2
	for _, v2 := range table2.([]grade) {
		m2[keyFunc2(v2)] = v2
	}

	// Join
	result := []interface{}{}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			result = append(result, resultFunc(v1, v2))
		}
	}
	return result
}
