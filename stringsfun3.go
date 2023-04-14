package main

import "fmt"

func main() {

}

func count(s string) (int, error) {
	chars := len(s)
	if chars == 0 {
		return 0, fmt.Errorf("ta vazio")
	}
	return chars, nil
}