package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	arr, err := getColl()
	if err != nil {
		log.Fatal(err)
	}
	for _, element := range arr {
		fmt.Println(element)
	}

}

func getColl() ([]string, error) {
	coll := []string{"GoLang", "JavaScript", "T-SQL", "Angular", "MongoDB"}
	err := errors.New("Error happens")
	return coll, err
}
