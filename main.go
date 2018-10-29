package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
)

type All interface{}

func main() {
	var t All

	input, _ := ioutil.ReadAll(os.Stdin)
	json.Unmarshal(input, &t)

	jsonIndent, _ := json.MarshalIndent(&t, "", "　　")
	fmt.Println("-------------")
	fmt.Println(string(jsonIndent))
}

