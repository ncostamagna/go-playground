package main

/*
#include <stdio.h>

void helloFromC() {
    printf("Hello from C!\n");
}
*/
import "C"

func main() {
    C.helloFromC()
}