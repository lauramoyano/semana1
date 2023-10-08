package foobar

import "strconv"

func Foobar(num int) string {
	var chainstr string
	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0 {
			chainstr += "FooBar"
		} else if i%3 == 0 {
			chainstr += "Foo"
		} else if i%5 == 0 {
			chainstr += "Bar"
		} else {
			chainstr += strconv.Itoa(i)
		}

		if i < num {
			chainstr += "->"
		}
	}

	return chainstr
	// Nueva línea al final para un formato limpio.
}
