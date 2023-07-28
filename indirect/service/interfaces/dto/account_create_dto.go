package dto

type AccountCreateInDto struct {
	Account  string
	Password string
}
type AccountCreateOutDto struct {
	Account  string
	Password string
	Count    int64
}
