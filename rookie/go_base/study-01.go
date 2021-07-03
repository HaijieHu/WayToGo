package main

import(
	"fmt"
	"os"
)

func init() {
	fmt.Println("Program Start!")
}

func main()  {
	fmt.Println("Hello World1")
	var s, sep string
	for i:=0; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}