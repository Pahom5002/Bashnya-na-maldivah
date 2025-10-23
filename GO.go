package main

import "fmt"

func main() {
	var num int
	var count int
	var des []string
	var ed []string
	var hundreds []string

	count = 0
	ed = []string{"один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять", "одинадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	des = []string{"двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семдесят", "восемдесят", "девяноста"}
	hundreds = []string{"сто", "двести", "триста", "четыриста", "пятьсот", "шестьсот", "семсот", "восемсот", "девятьсот"}
	fmt.Println("введите число меньше 12307")

begin1:

	fmt.Scan(&num)

	if num > 12307 {
		fmt.Println("вы ввели чило болше 12307, введите число меньше этого числа, пожалуйста;)")
		goto begin1
	}
begin:
	for num < 12307 {

		count++
		fmt.Println("итерация №", count)
		fmt.Println(num)

		if num%13 == 0 {
			if num%9 == 0 {
				fmt.Println("services error")
				goto last
			}
		}

		if num < 0 {

			num = num * -1
			goto begin

		}
		if num%7 == 0 {

			num = num * 39
			goto begin

		}
		if num%13 != 0 {
			if num%9 != 0 {
				num = num + 1
				goto begin
			}
		}
		if num%9 == 0 {

			num = num*13 + 1
			goto begin

		} else {
			num = (num + 2) * 3
			goto begin

		}

	}

	fmt.Println("итерация №", count+1)
	fmt.Println(num)
	fmt.Println("________________________")
	fmt.Println("Итоговое значение:", num)

	if num < 20 {
		fmt.Println(ed[num-1])
	}
	if num < 100 {
		var num1 int
		num1 = num - num%10

		fmt.Println(des[num1/10-2], ed[num%10-1])

	}
	if num < 1000 {
		var num2 int
		num2 = num - num%100
		fmt.Println(hundreds[num2/100-1], des[(num%100-num%10)/10-2], ed[num%10-1])
	}
	if num > 1000 {
		if num < 2000 {
			var num3 int
			num3 = num - num%1000
			fmt.Println("тысяча", hundreds[num3/1000-1], des[(num%100-num&10)/10-2], ed[num%10-1])
		}
	}
	if num > 2000 {
		if num < 3000 {
			var num4 int
			num4 = num - num%1000
			fmt.Println("две тысячи", hundreds[num4/1000-1], des[(num%100-num&10)/10-2], ed[num%10-1])

		}

	}
	if num > 3000 {
		if num < 20000 {
			if (num - (num - num%10)) == 0 {
				fmt.Println(ed[(num-num%1000)/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2])
			} else {
				fmt.Println(ed[(num-num%1000)/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
			}

		}
	}
	if num > 20000 {
		if num < 100000 {
			var num7 int
			var num6 int
			var num8 int
			var num9 int
			num6 = num - num%10000
			num7 = num%1000 - num%100
			num8 = num%100 - num%10
			num9 = num % 10
			if ((num - num&1000 - (num - num%10000)) / 1000) == 1 {

				fmt.Println(des[num6/10000-2], "одна тысяча", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
				goto last
			}
			if ((num - num&1000 - (num - num%10000)) / 1000) == 2 {

				fmt.Println(des[num6/10000-2], "две тысячи", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
				goto last
			}
			if ((num - num&1000 - (num - num%10000)) / 1000) > 2 {
				if ((num - num&1000 - (num - num%10000)) / 1000) < 5 {

					fmt.Println(des[num6/10000-2], ed[(num-num&1000-(num-num%10000))/1000-1], "тысячи", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
					goto last
				}
			}
			if num7 == 0 {
				fmt.Println(des[num6/10000-2], ed[(num-num&1000-(num-num%10000))/1000-1], "тысяч", des[(num%100-num%10)/10-2], ed[num%10-1])
				goto last
			}
			if num8 == 0 {
				fmt.Println(des[num6/10000-2], ed[(num-num&1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], ed[num-(num-num%10)-1])

			}
			if num9 == 0 {
				fmt.Println(des[num6/10000-2], ed[(num-num&1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2])
			}
			if (num - num%1000 - (num - num%10000)) == 0 {
				fmt.Println(des[num6/10000-2], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])

			} else {

				fmt.Println(des[num6/10000-2], ed[(num-num&1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
			}

		}

	}
	if num > 100000 {

		if ((num - num%10000 - (num - num%100000)) / 10000) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], ed[(num-num%1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
			goto last
		}

		if ((num - num%1000 - (num - num%10000)) / 1000) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
			goto last
		}

		if ((num - num%100 - (num - num%1000)) / 100) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], ed[(num-num%1000-(num-num%10000))/1000-1], "тысяч", des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
			goto last
		}

		if ((num - num%10 - (num - num%100)) / 10) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], ed[(num-num%1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], ed[num-(num-num%10)-1])
			goto last
		}

		if (num - (num - num%10)) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], ed[(num-num%1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2])
			goto last
		}

		if ((num-num%1000-(num-num%10000))/1000) == 0 && (num-(num-num%10)) == 0 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2])
			goto last
		}

		if ((num - num%1000 - (num - num%10000)) / 1000) == 1 {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], "одна тысяча", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num%10-1])
		}
		if ((num - num%1000 - (num - num%10000)) / 1000) < 5 {
			if ((num - num%1000 - (num - num%10000)) / 1000) > 1 {
				fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], ed[(num-num%1000-(num-num%10000))/1000-1], "тысячи", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num%10-1])
			}
		} else {
			fmt.Println(hundreds[(num-num%100000)/100000-1], des[(num-num%10000-(num-num%100000))/10000-2], ed[(num-num%1000-(num-num%10000))/1000-1], "тысяч", hundreds[(num-num%100-(num-num%1000))/100-1], des[(num-num%10-(num-num%100))/10-2], ed[num-(num-num%10)-1])
		}

	}
last:
	fmt.Println(" ")
	fmt.Println("Благодарю за число, всего Вам хорошего)")
}
