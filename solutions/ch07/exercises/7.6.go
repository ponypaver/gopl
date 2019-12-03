package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type celsiusFlag struct{ Celsius }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		c.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	c := celsiusFlag{value}
	flag.CommandLine.Var(&c, name, usage)
	return &c.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
