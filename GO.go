package main

import "fmt"

func main() {
	var num float64
	fmt.Println("введите число меньше 12307")

begin:

	fmt.Scan(&num)
	num = num

	if num > 12307 {
		fmt.Println("вы ввели чило болше 12307, введите число меньше этого числа, пожалуйста;)")
		goto begin
	}
	fmt.Println("итерации в цикле:")
	for num < 12307 {
		if int(num)%13 == 0 {
			if int(num)%9 == 0 {
				fmt.Println("servise error")
			}
		}
		if int(num)%13 != 0 {
			if int(num)%9 != 0 {
				num = num + 1
			}
		}

		if num < 0 {

			num = num * -1

		}
		if int(num)%7 == 0 {

			num = num * 39

		}
		if int(num)%9 == 0 {

			num = num*13 + 1

		} else {
			num = (num + 2) * 3

		}
		fmt.Println(num)

	}
	fmt.Println("________________________")
	fmt.Println("Итоговое значение:", num)

}
