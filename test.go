package main

import (
	"fmt"
	"os"
)

func main() {
	list:=os.Args
	fmt.Println(list[1])
	/*for a,b:= range list {
		fmt.Println(a,b)
	}*/
}
