package main

import (
	"fmt" // fmt is a standard library package no need of full path
	"os"
	"strings"

	"github.com/lauramoyano/semana1/bmi"
	"github.com/lauramoyano/semana1/evenorodd" // full path including the mod name
	"github.com/lauramoyano/semana1/foobar"
	"github.com/lauramoyano/semana1/mario"
	"github.com/lauramoyano/semana1/ohm"
	// full path including the mod name
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(os.Args[1])
		exercise := os.Args[1]
		exercise = strings.ToLower(exercise)
		switch exercise {
		case "evenodd":
			fmt.Print("Enter a number: ")
			var num int
			_, err := fmt.Scanln(&num)
			if err != nil {
				fmt.Println("Invalid input:", err)
				return
			}

			result := evenorodd.EvenOrOdd(num)
			fmt.Printf("Result: %s\n", result)
		case "foobar":
			fmt.Print("Enter a number: ")
			var num int
			_, err := fmt.Scanln(&num)
			if err != nil {
				fmt.Println("Invalid input:", err)
				return
			}
			// %s string
			result := foobar.Foobar(num)
			fmt.Printf("Result: %s\n", result)

		case "ohm":
			var voltagePtr *int
			var resistancePtr *int
			var currentPtr *int

			fmt.Print("Enter voltage (V): ")
			var voltage int
			voltageSet := false
			_, err := fmt.Scanln(&voltage)
			if err == nil {
				voltagePtr = &voltage
				voltageSet = true
			}

			fmt.Print("Enter resistance (R): ")
			var resistance int
			resistanceSet := false
			_, err = fmt.Scanln(&resistance)
			if err == nil {
				resistancePtr = &resistance
				resistanceSet = true
			}

			fmt.Print("Enter current (I): ")
			var current int
			currentSet := false
			_, err = fmt.Scanln(&current)
			if err == nil {
				currentPtr = &current
				currentSet = true
			}

			if voltageSet || resistanceSet || currentSet {
				result := ohm.Ohm(voltagePtr, resistancePtr, currentPtr)
				if result != nil {
					fmt.Printf("Result: %d\n", *result)
				} else {
					fmt.Println("nil")
				}
			} else {
				fmt.Println("Invalid input. At least one value (V, R, or I) must be provided.")
			}

		case "bmi": // Add a new case for bmi
			fmt.Print("Enter your weight (kg): ")
			var weight float64
			_, err := fmt.Scanln(&weight)
			if err != nil {
				fmt.Println("Invalid input:", err)
				return
			}

			fmt.Print("Enter your height (m): ")
			var height float64
			_, err = fmt.Scanln(&height)
			if err != nil {
				fmt.Println("Invalid input:", err)
				return
			}
			//.2 float number
			result := bmi.CalculateBMI(weight, height)
			fmt.Printf("Your BMI is %.2f\n", result)

		case "mario": // Add a new case for bmi
			fmt.Print("Pyramid height: ")
			var height int
			_, err := fmt.Scanln(&height)
			if err != nil {
				fmt.Println("Invalid input:", err)
				return
			}

			mario.Mario(height)

		default:
			fmt.Println("Unknown exercise:", exercise)
		}
	}

}
