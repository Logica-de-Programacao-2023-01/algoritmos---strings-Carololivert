package main

import (
	"database/sql"
	"fmt"
)

func media(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("canot divide by zero")
	}
	soma := 0
	for _, x := range list {

	}

	return soma / len(list), nil

}

func main() {
	result, err := media ([] int{})
	if err != nil{
		fmt.Printf("teve um erro : %s", err){

		}

	}

}
