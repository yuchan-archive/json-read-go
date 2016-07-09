package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonobject struct {
	Domains []string
}

func main() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))

	var jsontype jsonobject
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
	fmt.Println(jsontype.Domains[0])
}
