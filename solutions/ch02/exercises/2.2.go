// Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no arguments,
// and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pound s and kilograms, and the like.

package main

import (
	"fmt"
	"os"
	"strconv"
)

type (
	// Temperature
	Celsius    float64
	Fahrenheit float64

	// Length
	Feet  float64
	Meter float64

	// Weight
	Pound    float64
	Kilogram float64
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Conversion between Feet & Meter
func FToM(f Feet) Meter { return Meter(f / 3.281) }
func MToF(m Meter) Feet { return Feet(m * 3.281) }

// Conversion between Pound & Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / 2.205) }
func KToP(k Kilogram) Pound { return Pound(k * 2.205) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		tf := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			tf, FToC(tf), c, CToF(c))

		lf := Feet(t)
		m := Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			lf, FToM(lf), m, MToF(m))

		p := Pound(t)
		k := Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, PToK(p), k, KToP(k))
	}
}
