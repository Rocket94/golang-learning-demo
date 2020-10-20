package main

import (
	"fmt"
	"flag"
	"strconv"
)
type Celsius float64

type Fahrenheit float64 // 华氏温度
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f	, name, usage)
	return &f.Celsius
}
func (f *celsiusFlag)String()string{
	 return fmt.Sprintf("%g°C", f)
	return strconv.FormatFloat(float64(f.Celsius),'e',-1,64)
}
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
//var period = flag.Duration("period", 1*time.Second, "sleep period")
//var i=flag.Int("integer",0,"the number")
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println()
}
