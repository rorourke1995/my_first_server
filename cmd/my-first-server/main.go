package main

import (
	"fmt"

	"github.com/rorourke1995/my_first_server/bank"
)

func main() {
	b := bank.NewBank(100)
	fmt.Println(b)
}
