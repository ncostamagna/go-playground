package main

import "fmt"

// zero memory overhead for the values themselves, compare with 
// boolean this is the best optimal way to implement a set
type Set[T comparable] = map[T]struct{}


func Add[T comparable](s Set[T], value T) {
	s[value] = struct{}{}
}

func Contains[T comparable](s Set[T], value T) bool {
	_, ok := s[value]
	return ok
}


func main() {
	fruits := Set[string]{}
	Add(fruits, "apple")
	Add(fruits, "banana")
	Add(fruits, "cherry")

	fmt.Println(Contains(fruits, "apple"))
	fmt.Println(Contains(fruits, "banana"))
	fmt.Println(Contains(fruits, "cherry"))


	numbers := Set[int]{}
	Add(numbers, 1)
	Add(numbers, 2)
	Add(numbers, 3)

	fmt.Println(Contains(numbers, 1))
	fmt.Println(Contains(numbers, 2))
}
