package main

import (
	"fmt"
)

func main() {

	// strings.Replace()

	var number int
	fmt.Scan(&number)

	if number < 0 {
		number = -1 * number

		fmt.Print("- ")
	}

	numberToWord(number)

}

func numberToWord(number int) { // 56_987

	var dig = 0             // 5
	var tenPower = 1        // 10_000
	var workNumber = number //56_987

	for workNumber != 0 {
		workNumber = workNumber / 10
		dig++
	}

	for i := 1; i < dig; i++ {
		tenPower *= 10
	}

	for dig >= 1 { // 2

		workNumber = number / tenPower // 87 / 10 = 8

		if dig%3 == 2 { // 2%3 == 2 true

			switch workNumber { // 8

			case 1:
				fmt.Print("o'n ")

			case 2:
				fmt.Print("yigirma ")

			case 3:
				fmt.Print("o'ttiz ")

			case 4:
				fmt.Print("qirq ")

			case 5:
				fmt.Print("ellik ") // ellik

			case 6:
				fmt.Print("oltmish ")

			case 7:
				fmt.Print("yetmish ")

			case 8:
				fmt.Print("sakson ") // sakson

			case 9:
				fmt.Print("to'qson ")

			}
		} else {

			switch workNumber {
			case 1:
				fmt.Print("bir ")
			case 2:
				fmt.Print("ikki ")
			case 3:
				fmt.Print("uch ")
			case 4:
				fmt.Print("to`rt ")
			case 5:
				fmt.Print("besh ")
			case 6:
				fmt.Print("olti ") // olti
			case 7:
				fmt.Print("yetti ")
			case 8:
				fmt.Print("sakkiz ")
			case 9:
				fmt.Print("to'qqiz ") // to`qqiz
			}

		}

		if dig%3 == 0 && workNumber != 0 { // 3
			fmt.Print("yuz ") // yuz
		}

		if workNumber != 0 { // 9
			switch dig { // 3
			case 4:
				fmt.Print("ming ") // ming
			case 7:
				fmt.Print("million ")
			case 10:
				fmt.Print("milliard ")
			case 13:
				fmt.Print("trillion ")
			}
		}

		number %= tenPower // 87 % 10 = 7
		tenPower /= 10     // 1
		dig--              // 1

		// ellik olti ming to`qqiz yuz sakson yetti

	}

	fmt.Println()

}

// 	   var counterA, counterB int32

//     for i := 0; i < len(a); i++ {

//        if a[i] < b[i] {
//            counterB++
//        } else if a[i] > b[i] {
//            counterA++
//        }

//     }

//     var result = [2] int32 {counterA, counterB}

//     return result[:]
