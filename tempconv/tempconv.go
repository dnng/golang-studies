// gopl.io/ch2/tempconv0
// Package tempconv0 performs Celcius and Farenheit temperature computations.
package tempconv

import "fmt"

type Celcius float64
type Farenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

// The declaration below, in which the Celsius parameter c appears before the
// function name, associates with the Celcius type a method named String that
// returns c's numeric value followed by °C
func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

func CToF(c Celcius) Farenheit { return Farenheit(c*9/5 +32) }
func CToK(c Celcius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

func FToC(f Farenheit) Celcius { return Celcius((f-32) * 5 / 9) }
func FToK(f Farenheit) Kelvin { return CToK(FToC(f)) }

func KToC(k Kelvin) Celcius { return Celcius(k - Kelvin(AbsoluteZeroC)) }
func KToF(k Kelvin) Farenheit { return CToF(KToC(k)) }

