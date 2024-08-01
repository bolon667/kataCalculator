package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arabToRimNum(val1 string) int {
	result := -1
	switch strings.TrimSpace(val1) {
	case "I":
		result = 1
	case "II":
		result = 2
	case "III":
		result = 3
	case "IV":
		result = 4
	case "V":
		result = 5
	case "VI":
		result = 6
	case "VII":
		result = 7
	case "VIII":
		result = 8
	case "IX":
		result = 9
	case "X":
		result = 10
	}
	return result
}

func ToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			result += sym[i]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	vals := make([]string, 3)
	result := 0
	isRimNum1 := false //false - arab, true - rim
	isRimNum2 := false
	sign := ""

	//Ищем операторы в строке
	if strings.Contains(input, "+") {
		sign = "+"
	} else if strings.Contains(input, "-") {
		sign = "-"
	} else if strings.Contains(input, "*") {
		sign = "*"
	} else if strings.Contains(input, "/") {
		sign = "/"
	} else {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	//Делим строку на аргументы
	vals = strings.Split(input, sign)
	//Проверяем на число аргументов
	if len(vals) > 2 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	//Конвертируем аргументы(текст) в число
	val1, err1 := strconv.Atoi(strings.TrimSpace(vals[0]))
	val2, err2 := strconv.Atoi(strings.TrimSpace(vals[1]))

	//Если не получилось сконвертировать в число, то, скорей всего, это римская цифра
	if err1 != nil {
		val1 = arabToRimNum(vals[0])
		if val1 != -1 {
			isRimNum1 = true
		} else {
			panic("Выдача паники, так как числа не должны быть меньше 10 включительно")
		}
	}
	//Второе число так же проверяем
	if err2 != nil {
		val2 = arabToRimNum(vals[1])
		if val2 != -1 {
			isRimNum2 = true
		} else {
			panic("Выдача паники, так как числа не должны быть меньше 10 включительно")
		}
	}
	if val1 > 10 || val1 < 1 || val2 > 10 || val2 < 1 {
		panic("Выдача паники, так как числа должны быть >= 1 и <= 10")
	}
	//По условию задачи, типы чисел должны быть одинаковыми
	if isRimNum1 == isRimNum2 {
		//Вычисляем результат
		switch sign {
		case "+":
			result = val1 + val2
		case "-":
			result = val1 - val2
		case "*":
			result = val1 * val2
		case "/":
			result = val1 / val2
		}
		if (isRimNum1 && isRimNum2) && (result <= 0) {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	} else {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	//Выводим результат соотв. цифрами
	if isRimNum1 {
		fmt.Println(ToRoman(result))
	} else {
		fmt.Println(result)
	}

}
