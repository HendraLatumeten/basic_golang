package interfaces

type Hitung2d interface {
	luas() float64
	keliling() float64
}
type Hitung3d interface {
	volume() float64
}
type Hitung interface {
	hitung2d()
	hitung3d()
}
