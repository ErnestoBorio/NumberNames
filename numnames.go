package main

import "fmt"

var digit = [...]string{
	"", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve",
}

var tens = [...]string{
	"", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa",
}

var hundreds = [...]string{
	"", "ciento", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos",
	"ochocientos", "novecientos",
}

func main() {
	for i := 999999; ; i++ {
		alpha := int2string(i)
		if alpha == "" {
			break
		}
		fmt.Println(alpha)
	}
}

func int2string(num int) string {
	switch {
	case num == 0:
		return ""
	case num <= 9:
		return digit[num]
	case num == 10:
		return "diez"
	case num == 11:
		return "once"
	case num == 12:
		return "doce"
	case num == 13:
		return "trece"
	case num == 14:
		return "catorce"
	case num == 15:
		return "quince"
	case num == 16:
		return "dieciséis"
	case num <= 19:
		return "dieci" + digit[num-10]
	case num == 20:
		return "veinte"
	case num == 22:
		return "veintidós"
	case num == 23:
		return "veintitrés"
	case num == 26:
		return "veintiséis"
	case num <= 29:
		return "veinti" + digit[num-20]

	case num <= 99: // 30 .. 99
		decade := num / 10
		if num%10 == 0 {
			return tens[decade]
		}
		return tens[decade] + " y " + int2string(num-decade*10)

	case num == 100:
		return "cien"
	case num <= 999: // 101 .. 999
		century := num / 100
		return hundreds[century] + " " + int2string(num-century*100)
		// multiples of 100 will have a trailing space

	case num <= 999999: // 1000 .. 999.999
		millennium := num / 1000
		tail := num - millennium*1000 // last 3 digits
		right := int2string(tail)
		left := ""
		if millennium > 1 {
			left = int2string(millennium)
			if millennium >= 21 && millennium%10 == 1 { // millennia ending in 1: doscientos un (mil...)
				left = left[:len(left)-1] // cut the trailing "o" of "uno" to make it "un"
			}
			left = left + " "
		}
		return left + "mil " + right

	default:
		return ""
	}
}
