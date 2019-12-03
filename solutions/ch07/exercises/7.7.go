/*
Explain why the help message contains °C when the default value of 20.0 does
not.
 */
package main

import "fmt"

/*
func (f *FlagSet) Var(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}
*/

func main() {
	explanation := `
Explanation:
	func (f *FlagSet) Var(value Value, name string, usage string) {
		// Remember the default value as a string; it won't change.
		flag := &Flag{name, usage, value, value.String()}
	...
	}
	From the code above, we can see that value.String() called
`
	fmt.Println(explanation)
}