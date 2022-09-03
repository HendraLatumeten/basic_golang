package controller

import (
	"fmt"
	"strconv"
)

type Deret struct {
	prima     int
	ganjil    int
	genap     int
	fibonacci int
}

func DeretBilangan(x int) {
	BilPrima := Deret{prima: x}
	fmt.Println("Bilangan Prima :", BilPrima.GetPrima())

	BilGanjil := Deret{ganjil: x}
	fmt.Println("Bilangan Ganjil :", BilGanjil.GetGanjil())

	BilGenap := Deret{genap: x}
	fmt.Println("Bilangan Genap :", BilGenap.GetGenap())

	BilFibonacci := Deret{fibonacci: x}
	fmt.Println("Bilangan Fibonacci :", BilFibonacci.GetFibonacci())

}

func (value *Deret) GetPrima() string {
	x := value.prima
	var result string

	for bil := 1; bil < x; bil++ {
		i := 0
		for bag := 1; bag < x; bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			result += strconv.Itoa(bil) + " "
		}
	}
	return result

}

func (value *Deret) GetGanjil() string {
	x := value.ganjil
	var result string

	for i := 0; i < x; i++ {
		if i%2 == 1 {
			result += strconv.Itoa(i) + " "

		}
	}
	return result

}

func (value *Deret) GetGenap() string {
	x := value.genap
	var result string
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			result += strconv.Itoa(i) + " "
		}
	}
	return result
}

func (value *Deret) GetFibonacci() string {
	num := value.fibonacci
	var result string
	a := 0
	b := 1
	c := b
	for true {
		c = b
		b = a + b
		if b >= num {
			break
		}
		a = c
		result += strconv.Itoa(b) + " "
	}
	return result
}
