package main

import (
	"fmt"
	"github.com/brazanation/go-documents"
)

func main() {
	cnh := brdocs.CNH("53197826392")

	fmt.Println(cnh)
	fmt.Println(cnh.Format())
	fmt.Println(cnh.IsValid())
	//	cnh := brdocs.NewCNH("53197826392")
	//
	//	fmt.Println(cnh.Format())
	//
	//	fmt.Println(cnh.IsValid())

	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(cnh.Format())
	//	fmt.Println(cnh)
	//
	//	// Something different
	//	var cnh cnh = "3227492687A"
	//
	//	if ok, err := cnh.IsValid(); !ok {
	//		panic(err)
	//	}
}
