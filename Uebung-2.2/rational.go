package main

import "fmt"

type Rational struct {
	numerator int
	denominator int
}

func NewRational (numerator int, denominator int) Rational {
	if denominator == 0 {
		panic("division by zero")
		}
		r := Rational{}
		divisor := gcd(numerator, denominator)
		r.numerator = numerator / divisor
		r.denominator = denominator / divisor
		return r
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func (x Rational) String() string {
	return fmt.Sprintf("(%v/%v)", x.numerator, x.denominator)
	}

func main() {
	r1 := NewRational(20, 16)
	fmt.Println(r1)
}