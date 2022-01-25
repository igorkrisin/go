package main

import (
	"fmt"
	"strconv"
)

var global map[string]interface{}

type list struct {
	data     interface{}
	nextdata *list // * указатель (тип)

}

func printList(xs interface{}) {
	//fmt.Print("(")
	switch tempEl := xs.(type) {
	case *list:
		fmt.Print("(")
		for tempEl != nil {
			printList(tempEl.data)
			tempEl = tempEl.nextdata
			if tempEl != nil {
				fmt.Print(" ")
			} else {
				continue
			}
		}
		fmt.Print(")")
	default:
		fmt.Print(tempEl)

	}
	//fmt.Print(")")

}

func equalEl(elem interface{}, str string) bool {

	switch v := elem.(type) {
	case string:

		//fmt.Println("v == str",v == str, "v: ", v, "str: ", str)
		return v == str

	default:
		return false
	}

}

//(1 (2 4)) -> "(" "1" "(" "2" "4" ")" ")"  ->
//stack: {1 {2, 4}}

func listReverse(xs *list) *list {

	var tempList *list = nil
	for xs != nil {
		tempList = &list{data: xs.data, nextdata: tempList}
		xs = xs.nextdata
	}
	return tempList
}

func lenList(xs *list) int {
	count := 0
	for xs != nil {
		xs = xs.nextdata
		count++
	}
	return count
}

func parse(arr []string) (interface{}, bool) {
	var stack *list = nil
	count := 0
	count2 := 0
	for i := 0; i < len(arr); i++ {
		if equalEl(arr[i], ")") {
			fmt.Println()
			count++
		} else if equalEl(arr[i], "(") {
			count2++
		}
		if !equalEl(arr[i], ")") {
			j, err := strconv.Atoi(arr[i])
			fmt.Println("Atoi: ", j, err, arr[i])
			if err != nil {
				stack = &list{data: arr[i], nextdata: stack}
				fmt.Println("\n", "stack: ")
				printList(stack)
			} else {
				stack = &list{data: j, nextdata: stack}
			}
		} else if equalEl(arr[i], ")") {
			var tempList *list = nil
			for lenList(stack) != 0 && !equalEl(stack.data, "(") {
				tempList = &list{data: stack.data, nextdata: tempList}
				fmt.Println("\n", "tempList: ")
				printList(tempList)
				stack = stack.nextdata
				//fmt.Println("stack: ",stack[:len(stack)-1])
			}
			stack = stack.nextdata
			stack = &list{data: tempList, nextdata: stack}
		}
	}
	if count != count2 {
		return "error parenthesses", false
	}
	printList(stack)
	fmt.Println("\n", "return parse: ")
	return stack.data, true //некий костыль
}

func tokenize(data string) []string {
	storeStr := ""
	arr := []string{}
	for i := range data {
		if data[i] == 40 || data[i] == 41 {
			if len(storeStr) != 0 {
				arr = append(arr, storeStr)
			}
			arr = append(arr, data[i:i+1])
			storeStr = ""
		} else if data[i] == 32 || data[i] == 12 {
			if len(storeStr) != 0 {
				arr = append(arr, storeStr)
				storeStr = ""
			}
		} else {
			storeStr += data[i : i+1]
		}
	}
	if len(storeStr) != 0 {
		arr = append(arr, storeStr)
	}
	fmt.Println("storeStr:", storeStr)
	fmt.Print("return: ")

	return arr
}

func arrReverse(arr []interface{}) []interface{} {
	var tempArr []interface{} = []interface{}{}
	for i := len(arr) - 1; i >= 0; i-- {
		tempArr = append(tempArr, arr[i])
	}
	return tempArr
}

func evalList(tempEl *list, dict map[string]interface{}) (interface{}, bool) {

	var tempList *list = nil
	//tempEl = tempEl.nextdata
	for tempEl != nil {
		elem, mess := eval(tempEl.data, dict)
		if !mess {
			return elem, mess
		} else {
			tempList = &list{data: elem, nextdata: tempList}
			//elem,mess := tempEl.data
			tempEl = tempEl.nextdata
		}
	}
	return listReverse(tempList), true
}

/*func evalListRecur(tempEl *list, dict map[string]interface{}) *list {
	if tempEl != nil {
		return &list{data: eval(tempEl.data, dict), nextdata: evalListRecur(tempEl.nextdata, dict)}

	}
	return tempEl
}*/

//ДЗ сделать функцию evalList рекурсивной, отрезание аргумента добавить при вызове функции
// почитать про словари в го (map)
// null -- сделать функция в eval

//(+ x 42)

func eval(xs interface{}, dict map[string]interface{}) (interface{}, bool) {
	/* fmt.Print("xs: ")
	printList(xs)
	fmt.Println("") */
	switch tempEl := xs.(type) {
	case *list:
		//fmt.Println("tempEl.data: ")
		//printList(tempEl.data)
		//fmt.Println(" ")
		// true - нет ошибок
		// false - произошла ошибка
		if equalEl(tempEl.data, "+") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func +", false //todo: написать код недостатка аргументов для остальных функций
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}
			switch el1 := elem.(type) {
			case int:
				switch el2 := elem2.(type) {
				case int:
					return el1 + el2, true
				}
			default:
				return "arguments type not int +", false
			}
		} else if equalEl(tempEl.data, "if") {
			if lenList(tempEl) < 4 {
				return "not enough arguments in func if", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}

			switch el1 := elem.(type) {
			case string:
				if el1 == "true" {
					elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
					if !mess2 {
						return elem2, mess2
					}
					return elem2, true
				} else if el1 == "false" {
					elem3, mess3 := eval(tempEl.nextdata.nextdata.nextdata.data, dict) //todo: исправить путем присвоения в нужном месте элементов (не сразу все 3 вподряд)
					if !mess3 {
						return elem3, mess3
					}
					return elem3, true
				}
			default:
				return "arguments type not string if", false

			}
		} else if equalEl(tempEl.data, "cond") {
			//fmt.Println(lenList(tempEl))
			if lenList(tempEl) < 2 {
				return "not enough arguments in func if", false //question: why cond fatal error if count argument < 2
			}

			for tempEl.nextdata != nil {
				switch el1 := tempEl.nextdata.data.(type) {
				case *list:
					elem, mess := eval(el1.data, dict)
					if !mess {
						return elem, mess
					}
					if elem == "false" {
						tempEl = tempEl.nextdata
						eval(el1.data, dict)
					} else if elem == "true" {
						return el1.nextdata.data, true
					}
				default:
					return "arguments type not list cond", false

				}
			}
			/* for tempEl.nextdata != nil {
				switch el1 := tempEl.nextdata.data.(type) {
				case *list:
					if eval(el1.data, dict) == "true" {
						return el1.nextdata.data
					}
					tempEl = tempEl.nextdata
				}

			} */
		} else if equalEl(tempEl.data, "-") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func -", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			switch el1 := elem.(type) {
			case int:
				switch el2 := elem2.(type) {
				case int:
					return el1 - el2, true
				}
			default:
				return "arguments type not int -", false
			}
		} else if equalEl(tempEl.data, "*") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func *", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			switch el1 := elem.(type) {
			case int:
				switch el2 := elem2.(type) {
				case int:
					return el1 * el2, true
				}
			default:
				return "arguments type is not int *", false
			}
		} else if equalEl(tempEl.data, "/") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func /", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			switch el1 := elem.(type) {
			case int:
				switch el2 := elem2.(type) {
				case int:
					return el1 / el2, true
				}
			default:
				return "arguments type is not int /", false
			}

		} else if equalEl(tempEl.data, "=") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func =", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			switch el1 := elem.(type) {
			case int:
				switch el2 := elem2.(type) {
				case int:
					if el1 == el2 {
						return "true", true
					} else {
						return "false", true
					}
				}
			default:
				return "arguments type is not int =", false
			}
		} else if equalEl(tempEl.data, "quote") {
			if lenList(tempEl) < 1 {
				return "not enough arguments in func quote", false
			}
			return tempEl.nextdata.data, true

		} else if equalEl(tempEl.data, "car") {
			if lenList(tempEl) < 2 {
				return "not enough arguments in func car", false
			}
			//fmt.Println("car")
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}

			switch el1 := elem.(type) {
			case *list:
				return el1.data, true
			default:
				return "arguments type is not list car", false

			}
		} else if equalEl(tempEl.data, "cdr") {
			if lenList(tempEl) < 2 {
				return "not enough arguments in func cdr", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			switch el1 := elem.(type) {
			case *list:
				return el1.nextdata, true
			default:
				return "arguments type is not list cdr", false
			}
		} else if equalEl(tempEl.data, "cons") {
			//fmt.Println(lenList(tempEl))
			if lenList(tempEl) < 3 {
				return "not enough arguments in func cons", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			var el1 interface{} = elem //какая-то ошибка возникает при замене присвоения el2 на elem2
			switch el2 := elem2.(type) {
			case *list:
				var tempList *list = &list{data: el1, nextdata: el2}
				return tempList, true
			default:
				return "arguments type is not list cons", false
			}

		} else if equalEl(tempEl.data, "list") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func list", false
			}
			//tempEl = tempEl.nextdata
			/* fmt.Print("tempEldata: ")
			printList(tempEl.data)
			fmt.Println("")
			fmt.Print("tempEldata: ")
			printList(tempEl.nextdata.data)
			fmt.Println("") */
			elem, mess := evalList(tempEl.nextdata, dict)
			/* fmt.Print("elem: ")
			printList(elem)
			fmt.Println("") */
			if !mess {
				return elem, mess
			}
			return elem, true
			//return evalList(tempEl)
			/* var tempList *list = nil
			tempEl = tempEl.nextdata //вызов функции без первого элемента
			for tempEl != nil {
			    tempList = &list{data: eval(tempEl.data), nextdata: tempList}
			    tempEl = tempEl.nextdata
			}
			return listReverse(tempList)  */
		} else if equalEl(tempEl.data, "null") {
			if lenList(tempEl) < 2 {
				return "not enough arguments in func null", false
			}
			elem, mess := eval(tempEl.nextdata.data, dict)
			if !mess {
				return elem, mess
			}

			switch el1 := elem.(type) {
			case *list:
				//printList(el1)
				// (null (foo 42))
				if lenList(el1) != 0 {
					return "false", true
				} else {
					return "true", true
				}
			default:
				return "arguments type is not list null", false
			}

		} else if equalEl(tempEl.data, "define") {
			if lenList(tempEl) < 2 {
				return "not enough arguments in func define", false
			}

			//fmt.Println("define enter")
			switch el1 := tempEl.nextdata.data.(type) {
			case string:
				//fmt.Println("define exit")
				/* fmt.Print("tempEl.nextdata.nextdata.data define: ")
				printList(tempEl.nextdata.nextdata.data)
				fmt.Println("") */
				global[el1] = tempEl.nextdata.nextdata.data
				/*  fmt.Print("el1 define: ")
				fmt.Print(el1)
				fmt.Println("") */
				return el1, true
			}
		} else if equalEl(tempEl.data, "let") {
			if lenList(tempEl) < 2 {
				return "not enough arguments in func let", false
			}

			elem, mess := eval(tempEl.nextdata.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			eval(tempEl.nextdata.nextdata.data, dict)
			switch el1 := tempEl.nextdata.data.(type) {
			case *list:

				for el1 != nil {
					switch el2 := el1.data.(type) {
					case *list:
						switch el3 := el2.data.(type) {
						case string:

							elem2, mess2 := eval(el2.nextdata.data, dict)
							if !mess2 {
								return elem2, mess2
							}

							dict[el3] = elem2
						}
					}
					el1 = el1.nextdata
				}
			default:
				return "arguments type is not list let", false

			}
			return elem, true

		} else if equalEl(tempEl.data, "progn") {
			if lenList(tempEl) < 3 {
				return "not enough arguments in func progn", false
			}
			for tempEl != nil {

				eval(tempEl.data, dict)
				tempEl = tempEl.nextdata
				if tempEl.nextdata == nil {
					elem, mess := eval(tempEl.data, dict)
					if !mess {
						return elem, mess
					}
					return elem, true
				}

			}
		}
		// (define foo 42)
		switch el1 := tempEl.data.(type) {
		case *list:
			if lenList(tempEl) < 2 {
				return "not enough arguments in func lambda", false
			}

			var clue interface{} = el1.nextdata.data
			var val *list = tempEl.nextdata
			if equalEl(el1.data, "lambda") {

				switch el3 := clue.(type) {
				case *list:
					for el3 != nil && val != nil {
						switch el4 := el3.data.(type) {
						case string:
							switch el5 := val.data.(type) {
							case interface{}:
								dict[el4] = el5
							}
						}
						val = val.nextdata
						el3 = el3.nextdata

					}
					elem, mess := eval(el1.nextdata.nextdata.data, dict)
					if !mess {
						return elem, mess
					}
					return elem, true
				default:
					return "arguments type is not string lambda", false
				}

			}

		case string:

			j, err := global[el1]
			elem2, mess2 := evalList(tempEl.nextdata, dict)
			if !mess2 {
				return elem2, mess2
			}

			elem, mess := eval(&list{data: j, nextdata: elem2.(*list)}, dict)
			if !mess {
				return elem, mess
			}
			if err == false {
				fmt.Print("сase string global false: ")

				return j, false
			} else {

				return elem, true

			}

		}

		//cons 23 (1 2) -> (23 1 2)
	case string:
		if tempEl == "false" {
			return "false", true
		} else if tempEl == "true" {
			return "true", true
		}
		j, err := dict[tempEl]

		if err != false {
			return j, true
		}

		h, err2 := global[tempEl]
		if err2 == true {
			return h, true
		} else {
			//printList(h)
			//fmt.Println("err2", err2)
			return "variable " + tempEl + " is not defindet", err2
			//fmt.Println("")
		}

	case int:
		return xs, true
	}
	return xs, true
}

// ((lambda (x y) (+ x y)) 3 4)

func main() {
	global = make(map[string]interface{})
	dict := make(map[string]interface{})

	structVar1 := list{data: "1", nextdata: nil}         //
	structVar2 := list{data: "2", nextdata: &structVar1} // * - разименование
	structVar3 := list{data: "3", nextdata: &structVar2}
	s4 := list{data: &structVar3, nextdata: nil}
	s5 := list{data: "5", nextdata: &s4}
	s6 := list{data: "6", nextdata: &s5}

	printList(&s6) // (6 5 (3 2 1))

	// 5 = (len '(a b c d e)) =(+ (len '(b c d e)) 1) = 4

	printList(listReverse(&s6))
	//fmt.Print("EQUALMAIN",")" == ")")
	//fmt.Println(tokenize())
	elem2, mess2 := parse(tokenize("(progn (define len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(a b c))))"))
	fmt.Println(mess2, "\n")
	if !mess2 {
		fmt.Println(elem2)
		fmt.Println("Error return parse")
	} else {
		elem, mess := eval(elem2, dict)
		if !mess {
			fmt.Println(elem)
			fmt.Println("Error return eval")
		} else {
			printList(elem)
		}
	}

	// todo: поменять типы го на типы лист, модифиц парсер, что бы он возвращал 2 значения и реагировал на незакрытые скобки

	//(list(define len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(a b c))))
	//(list(define len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(1 2 3 4 5))))
	//(list(define x (lambda (y) 25)) (x 13)
	//(list(define cudr (lambda (y)(cdr y))) (cudr (1 2 3 4 5))) cdr
	//(list(define sq (lambda (y) (* (* y y) y))) (sq 3))
	//fmt.Println(structVar1, structVar2, structVar3)
	//(cons (quote(a b c)))
	//(cond((= 3 3)42))
	/* var bar interface{}
	   bar=42
	   bar = bar.(int) + 1
	   fmt.Println(bar)
	   bar="bar"
	   fmt.Println(bar)

	   if = 2 2 (/ 4 2) 42
	*/

}

//(1 2 4 6)
//(2 4 (5 6))
