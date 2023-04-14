package main

import "fmt"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("canot divide by zero")
	}
	return x / y, nil
}

func main() {
	var a, b float64

	fmt.Print("digite dois números")
	fmt.Scan(&a)
	fmt.Scan(&b)

	x := a
	y := b

	resultado, err := divide(x, y)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao dividr x e y: %s\n", err)
		return
	}

	fmt.Println(resultado)

}
