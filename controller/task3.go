package controller

import (
	"fmt"
	"math"
	"strconv"
)

type hitung struct {
	r float64 //jari jari
	//luas float64 // luas
	phi float64 // phi
	k   float64 //kelilingc
	p   int
	l   int
	t   int
}

func Hitung2d() {
	Luas := hitung{phi: 3.14}
	fmt.Println("Luas : ", Luas.luas())

	Keliling := hitung{phi: 3.14}
	fmt.Println("Keliling : ", Keliling.keliling())

}

func Hitung3d() {
	Volume := hitung{phi: 3.14}
	fmt.Println("Keliling : ", Volume.volume())

}

func (value *hitung) luas() float64 {
	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&value.p)
	fmt.Print("Masukkan Lebar: ")
	fmt.Scanln(&value.l)
	result := float64(value.p) * float64(value.l)
	return result
}
func (value *hitung) keliling() float64 {
	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&value.p)
	fmt.Print("Masukkan Lebar: ")
	fmt.Scanln(&value.l)
	value.k = 2 * (float64(value.p) + float64(value.l))
	s := fmt.Sprintf("%.2f", value.k)
	result, err := strconv.ParseFloat(s, 8)
	if err != nil {
		panic(err)
	}
	return result
}

func (value *hitung) volume() float64 {
	fmt.Print("Masukkan Jari Jari ")
	fmt.Scanln(&value.r)
	fmt.Print("Masukkan Tinggi: ")
	fmt.Scanln(&value.t)
	result := value.phi * value.r * value.r * float64(value.t)
	return math.Ceil(result)
}
