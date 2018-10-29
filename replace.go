package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println("-------------")
	fmt.Println(strings.Replace(string(input), " ", "　", -1))
}
