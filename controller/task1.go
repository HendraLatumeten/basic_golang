package controller

import (
	"fmt"
	"strconv"
)

func GetRound(x float64) string {
	s := fmt.Sprintf("%f", x)
	// a := strconv.FormatFloat(s, 'f', -1, 64)
	// z := strconv.Itoa(s)
	// join := strings.Split(s, ",")
	// z :=
	return s
}

func RoundFloat(x float64) string {
	z := x
	a := strconv.FormatFloat(z, 'f', -1, 64)
	sx := int(x)
	cv := a[3:]
	//rows, err := strconv.Atoi(cv)
	fmt.Println(cv)
	//	check := strings

	for i := 0; i < len(cv); i++ {

		fmt.Println(cv[i])
	}
	//g := int(len(a))
	//res := strconv.Itoa(cv)
	dj, err := strconv.Atoi(a[3:])
	if err != nil {
		panic(err)
	}
	ak := strconv.Itoa(sx) + "." + strconv.Itoa(dj)

	//fmt.Println(ak)
	return ak
}
