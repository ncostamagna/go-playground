package main

import (
	//"fmt"
	"testing"

	"unique"
)

type (
	StructA struct {
		Content string
	}
	StructB struct {
		Content string
	}	
)

var (
	s1 = "it's a fucking long string, its sooooo loooong, i'm sooooo tired"
	s2 = "it's a fucking long string, its sooooo loooong, i'm sooooo tired"
	a1 = StructA{Content: "it's a fucking long string, its sooooo loooong, i'm sooooo tired"}
	a2 = StructA{Content: "it's a fucking long string, its sooooo loooong, i'm sooooo tired"}
	sameStructA = StructA{Content: "it's a fucking long string, its sooooo loooong, i'm sooooo tired"}
	sameStructB = StructA{Content: "it's a fucking long string, its sooooo loooong, i'm sooooo tired"}
)

/*
func main() {
	fmt.Println(s1 == s2)
}*/

func BenchmarkStringComparisong(b *testing.B) {
	for b.Loop() {
		_ = s1 == s2
	}
}

func BenchmarkUniqueHandleComparisong(b *testing.B) {
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)

	for b.Loop() {
		_ = h1 == h2
	}
}

func BenchmarkDiffStructStringComparisong(b *testing.B) {
	for b.Loop() {
		_ = a1.Content == a2.Content
	}
}

func BenchmarkDiffStructUniqueHandleComparisong(b *testing.B) {
	h1 := unique.Make(a1.Content)
	h2 := unique.Make(a2.Content)

	for b.Loop() {
		_ = h1 == h2
	}
}


func BenchmarkSameStructStringComparisong(b *testing.B) {
	for b.Loop() {
		_ = sameStructA == sameStructB
	}
}

func BenchmarkSameStructUniqueHandleComparisong(b *testing.B) {
	h1 := unique.Make(sameStructA)
	h2 := unique.Make(sameStructB)

	for b.Loop() {
		_ = h1 == h2
	}
}