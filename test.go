package main

import (
	"fmt"
	//"bytes"
	"strconv"
)

func main() {
//	var r rune = 'h'
	var hi string = "gg"
	var ty  []byte = []byte{}
//	fmt.Println(rune(hi))
	var hida string = "helli" + "guss"
	fmt.Println(strconv.AppendQuote(ty,hi),hida)
}