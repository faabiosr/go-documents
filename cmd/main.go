package main

import (
	"fmt"
	"github.com/brazanation/go-documents"
)

func main() {
	doc := brdocs.NewCNH("53197826392")
	f(doc)
	fmt.Println(v(doc))
	//cnh := brdocs.CNH{"53197826392"}
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

func f(ft brdocs.Formatter) {
	fmt.Println(ft.Format())
}

func v(vt brdocs.Validator) error {
	if !vt.IsValid() {
		return fmt.Errorf("Validation failed")
	}

	return nil
}
