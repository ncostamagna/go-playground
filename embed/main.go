package main

import _ "embed"

//go:embed my-shit-script.sql
var s string

func main(){
	
	print(s)
}