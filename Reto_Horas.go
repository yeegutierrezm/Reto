package main

import "fmt"

func horadecir(hora int) string {
	h := make(map[int]string)
	h[1] = "one"
	h[2] = "two"
	h[3] = "three"
	h[4] = "four"
	h[5] = "five"
	h[6] = "six"
	h[7] = "seven"
	h[8] = "eigth"
	h[9] = "nine"
	h[10] = "ten"
	h[11] = "eleven"
	h[12] = "twelve"
	h[13] = "thirteen"
	h[14] = "fourteen"
	h[15] = "fiveteen"
	h[16] = "sixteen"
	h[17] = "seventeen"
	h[18] = "eigthteen"
	h[19] = "nineteen"
	h[20] = "twenty"
	h[21] = "twenty one"
	h[22] = "twenty two"
	h[23] = "twenty three"
	h[24] = "twenty four"
	h[25] = "twenty five"
	h[26] = "twenty six"
	h[27] = "twenty seven"
	h[28] = "twenty eight"
	h[29] = "twenty nine"
	h[30] = "thirty"
	value, _ := h[hora]
	return value
}

func respuesta(hora, min int) {
	minnuevo := 60 - min
	horanuevo := hora + 1
	if horanuevo > 12 {
		horanuevo = 1
	}

	switch {
	case min > 30 && min != 45:
		fmt.Printf("%s minutes to %s", horadecir(minnuevo), horadecir(horanuevo))
	case min < 30 && min != 00 && min != 15:
		fmt.Printf("%s minutes past %s", horadecir(min), horadecir(hora))
	case min == 15:
		fmt.Printf("quarter past %s", horadecir(hora))
	case min == 45:
		fmt.Printf("quarter to %s", horadecir(horanuevo))
	case min == 30:
		fmt.Printf("half past %s", horadecir(hora))
	case min == 00:
		fmt.Printf("%s o' clock", horadecir(hora))
	default:
		fmt.Println("No condicion")

	}
}

func main() {

	var hora int
	var min int

	fmt.Println("coloca la hora:")
	fmt.Scanf("%d \n", &hora)

	fmt.Println("coloca los minutos:")
	fmt.Scanf("%d \n", &min)

	respuesta(hora, min)

}
