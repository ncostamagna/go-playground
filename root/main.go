package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

func ReadFile(path string){

	// all operation are done in the root directory
	// you dont have access to outside of the root directory
	root, err := os.OpenRoot("public")
	if err != nil {
		log.Fatal("Failed to open root directory:", err)
	}
	defer root.Close()
	
	file, err := root.Open(path)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}
	
	fmt.Println("File content:", string(data))
}

func main() {

	os.WriteFile("public/safe.txt", []byte("This is a safe file"), 0644)

	os.WriteFile("secret.txt", []byte("Sensitive data"), 0644)
	ReadFile("safe.txt")
	ReadFile("../secret.txt") // be careful, it try to read a sensitive file
}
