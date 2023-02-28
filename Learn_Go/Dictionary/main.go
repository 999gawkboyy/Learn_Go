package main

import (
	"fmt"

	"github.com/mingeun3669/Dictionary/mydict"
)

func errorChack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dict := mydict.Dictionary{}
	var baseword string
	fmt.Print("KEY ? ")
	fmt.Scan(&baseword)
	dict.Add(baseword, "World")
	err := dict.Update(baseword, "Guys")
	errorChack(err)
	err2 := dict.Delete(baseword)
	errorChack(err2)
	word, err3 := dict.Search(baseword)
	errorChack(err3)
	fmt.Println(word)
}
