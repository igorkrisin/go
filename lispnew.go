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
	switch exp := xs.(type) {
	case *list:
		fmt.Print("(")
		for exp != nil {
			printList(exp.data)
			exp = exp.nextdata
			if exp != nil {
				fmt.Print(" ")
			} else {
				continue
			}
		}
		fmt.Print(")")
		//fmt.Print("\n")
	default:
		fmt.Print(exp)
		//fmt.Print("\n")

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
			//fmt.Println("i: ", i)
			if len(storeStr) != 0 {
				arr = append(arr, storeStr)
			}
			arr = append(arr, data[i:i+1])
			storeStr = ""
		} else if data[i] == 32/*sp*/ || data[i] == 10/*nl*/ || data[i] == 9 /*ht*/ { //todo convertion integer to literal
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

func evalList(exp *list, dict map[string]interface{}) (interface{}, bool) {

	var tempList *list = nil
	//exp = exp.nextdata
	for exp != nil {
		elem, mess := eval(exp.data, dict)
		if !mess {
			return elem, mess
		} else {
			tempList = &list{data: elem, nextdata: tempList}
			//elem,mess := exp.data
			exp = exp.nextdata
		}
	}
	return listReverse(tempList), true
}

func mapCopy(dict map[string]interface{}) map[string]interface{} {
	var tempDict = make(map[string]interface{})
	for key, val := range dict {
		tempDict[key] = val
	}
	return tempDict
}

func printMap(dict map[string]interface{}) {
    for key, val := range dict{
	printList(key)
	
	fmt.Print(" : ")
	printList(val)
	fmt.Print(" ")
    }
}

/*func evalListRecur(exp *list, dict map[string]interface{}) *list {
	if exp != nil {
		return &list{data: eval(exp.data, dict), nextdata: evalListRecur(exp.nextdata, dict)}

	}
	return exp
}*/

//ДЗ сделать функцию evalList рекурсивной, отрезание аргумента добавить при вызове функции
// почитать про словари в го (map)
// null -- сделать функция в eval

//(+ x 42)

func eval(xs interface{}, dict map[string]interface{}) (interface{}, bool) {//to do renmame xs to expr
	fmt.Print("expr: ")
	printList(xs)
	fmt.Println("")
	printMap(dict)
	fmt.Println("")
	switch exp := xs.(type) {
	case *list:
		// true - нет ошибок
		// false - произошла ошибка
		if equalEl(exp.data, "+") {
			if lenList(exp) < 3 {
				return "not enough arguments in func +", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func +", false
			}
			elem, mess := evalList(exp.nextdata, dict)
			if !mess {
				return elem, mess
			}

			switch el1 := elem.(*list).data.(type) {
			case int:
				switch el2 := elem.(*list).nextdata.data.(type) {
				case int:
					return el1 + el2, true
				}
			default:
				return "arguments type not int in func +", false
			}
		} else if equalEl(exp.data, "if") {
			if lenList(exp) < 4 {
				fmt.Println("lenList: ",lenList(exp))
				return "not enough arguments in func if", false // todo разобраться почему не работет количественное определение аргументов  if
			} else if lenList(exp) > 4 {
				return "too many arguments in func if", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			switch el1 := elem.(type) {
			case string:
			    //fmt.Println("el1: ", el1)
				if el1 == "true" {
					elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
					if !mess2 {
						return elem2, mess2
					}
					//fmt.Println("elem2: ")
					printList(elem2)
					
					return elem2, true
				} else if el1 == "false" {
					elem3, mess3 := eval(exp.nextdata.nextdata.nextdata.data, dict)
					if !mess3 {
						return elem3, mess3
					}
					//fmt.Println("elem3: ")
					printList(elem3)
					return elem3, true
				}
			default:
				return "arguments type not string in func if", false

			}
		} else if equalEl(exp.data, "cond") {

			if lenList(exp) < 2 {
				return "not enough arguments in func cond", false //question: why cond fatal error if count argument < 2
			} else if lenList(exp) > 2 {
				return "too many arguments in func cond", false
			}
			for exp.nextdata != nil {
				switch el1 := exp.nextdata.data.(type) {
				case *list:
					elem, mess := eval(el1.data, dict)
					if !mess {
						return elem, mess
					}
					if elem == "false" {
						exp = exp.nextdata
						eval(el1.data, dict)
					} else if elem == "true" {
						return el1.nextdata.data, true
					}
				default:
					return "arguments type not list in func cond", false

				}
			}
			/* for exp.nextdata != nil {
				switch el1 := exp.nextdata.data.(type) {
				case *list:
					if eval(el1.data, dict) == "true" {
						return el1.nextdata.data
					}
					exp = exp.nextdata
				}

			} */
		} else if equalEl(exp.data, "-") {
			if lenList(exp) < 3 {
				return "not enough arguments in func -", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func -", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
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
				return "arguments type not int in func -", false
			}
		} else if equalEl(exp.data, "*") {
			if lenList(exp) < 3 {
				return "not enough arguments in func *", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func *", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
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
				return "arguments type is not int in func *", false
			}
		} else if equalEl(exp.data, "/") {
			if lenList(exp) < 3 {
				return "not enough arguments in func /", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func /", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
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
				return "arguments type is not int in func /", false
			}

		} else if equalEl(exp.data, "=") {
			if lenList(exp) < 3 {
				return "not enough arguments in func =", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func =", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
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
				return "arguments type is not intin func  =", false
			}
		} else if equalEl(exp.data, "quote") {
			if lenList(exp) < 1 {
				return "not enough arguments in func quote", false
			} else if lenList(exp) > 2 {
				return "too many arguments in func quote", false
			}
			return exp.nextdata.data, true

		} else if equalEl(exp.data, "car") {
			
			if lenList(exp) < 2 {
				fmt.Println("len exp: ",lenList(exp) )
				return "not enough arguments in func car", false
			} else if lenList(exp) > 2 {
				return "too many arguments in func car", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			
			switch el1 := elem.(type) {
			case *list:
				return el1.data, true
			default:
				return "arguments type is not list in func car", false

			}
		} else if equalEl(exp.data, "cdr") {
			if lenList(exp) < 2 {
				return "not enough arguments in func cdr", false
			} else if lenList(exp) > 2 {
				return "too many arguments in func cdr", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			switch el1 := elem.(type) {
			case *list:
				return el1.nextdata, true
			default:
				return "arguments type is not list in func cdr", false
			}
		} else if equalEl(exp.data, "cons") {
			if lenList(exp) < 3 {
				return "not enough arguments in func cons", false
			} else if lenList(exp) > 3 {
				return "too many arguments in func cons", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			elem2, mess2 := eval(exp.nextdata.nextdata.data, dict)
			if !mess2 {
				return elem2, mess2
			}

			//var el1 interface{} = elem
			switch el2 := elem2.(type) {
			case *list:
				var tempList *list = &list{data: elem, nextdata: el2}
				return tempList, true
			default:
				return "second argument type is not list in func cons", false
			}

		} else if equalEl(exp.data, "list") {
			if lenList(exp) < 3 {
				return "not enough arguments in func list", false
			}
			elem, mess := evalList(exp.nextdata, dict)
			if !mess {
				return elem, mess
			}
			return elem, true
		} else if equalEl(exp.data, "null") {
			if lenList(exp) < 2 {
				return "not enough arguments in func null", false
			} else if lenList(exp) > 2 {
				return "too many arguments in func null", false
			}
			elem, mess := eval(exp.nextdata.data, dict)
			if !mess {
				return elem, mess
			}

			switch el1 := elem.(type) {
			case *list:
				if lenList(el1) != 0 {
					return "false", true
				} else {
					return "true", true
				}
			default:
				return "arguments type is not list in func null", false 
			}

		} else if equalEl(exp.data, "define") {
			if lenList(exp) < 2 {
				return "not enough arguments in func define", false
			}
			switch el1 := exp.nextdata.data.(type) {
			case string:
				global[el1] = exp.nextdata.nextdata.data
				return el1, true
			}
		} else if equalEl(exp.data, "let") {
			if lenList(exp) < 2 {
				return "not enough arguments in func let", false
			} else if lenList(exp) > 2 {
				return "too many arguments in func let", false
			}
			switch el1 := exp.nextdata.data.(type) {
			case *list:
				dict = mapCopy(dict)
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
				return "arguments type is not list in func let", false

			}
			elem, mess := eval(exp.nextdata.nextdata.data, dict)
			if !mess {
				return elem, mess
			}
			return elem, true

		//} else if equalEl
		
		} else if equalEl(exp.data, "progn") {
			if lenList(exp) < 3 {
				return "not enough arguments in func progn", false
			}
			for exp != nil {

				eval(exp.data, dict)
				exp = exp.nextdata
				if exp.nextdata == nil {
					elem, mess := eval(exp.data, dict)
					if !mess {
						return elem, mess
					}
					return elem, true
				}
			}
		} 
		// (define foo 42)
		switch el1 := exp.data.(type) {
		case *list:
			if lenList(el1) != 3 {
				return "wrong amount of arguments in func lambda", false
			}
				var actualArgs *list = exp.nextdata
				if equalEl(el1.data, "lambda") {
				switch formalArgs := el1.nextdata.data.(type) {
				case *list:
				    if lenList(actualArgs) != lenList(formalArgs) {
					return "lenght actualArgs not equal lenght formalArgs in lambda", false
				    }
					 fmt.Print("formalArgs: ")
					newDict := make(map[string]interface{})
					newDict = mapCopy(dict)
					for formalArgs != nil {
						switch varName := formalArgs.data.(type) {
						case string:
							switch val := actualArgs.data.(type) {
							case interface{}:
								elem2, mess2 := eval(val, dict)
								if !mess2 {
									return elem2, mess2
								}

								newDict[varName] = elem2
							}
							default:
								return "type formal arguments lambda not  string", false
						}
						actualArgs = actualArgs.nextdata
						formalArgs = formalArgs.nextdata
					}
					elem, mess := eval(el1.nextdata.nextdata.data, newDict)
					if !mess {
						return elem, mess
					}
					return elem, true
				default:
					return "arguments type is not list in function lambda", false
				}
			}
			case string:

			j, err := global[el1]
		
			
			if err == false {
				//fmt.Print("сase string global false: ")
				return "finction " +el1+  " not defined", false
			}

			elem, mess := eval(&list{data: j, nextdata: exp.nextdata}, dict)
			if !mess {
				return elem, mess
			
				
			} else {

				return elem, true

			}
		    }
		

			//cons 23 (1 2) -> (23 1 2)
			case string:
				if exp == "false" {
					return "false", true
				} else if exp == "true" {
					return "true", true
				}
				j, err := dict[exp]

				if err != false {
					return j, true
				}

			h, err2 := global[exp]
			if err2 == true {
				return h, true
			} else {
				//printList(h)
				//fmt.Println("err2", err2)
				return "variable " + exp + " is not defindet", err2
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
	// (append '(a b c) '(d e)) -> '(a b c d e)
	printList(listReverse(&s6))
	//fmt.Print("EQUALMAIN",")" == ")")
	//fmt.Println(tokenize())
	// (butLast '(a)) -> '()
	elem2, mess2 := parse(tokenize(`
	
	
	(progn
	    (define append(lambda (bs ys)
	         (if (null bs)
	             ys
	             (cons (car bs)(append (cdr bs)ys)))))
	    (define len (lambda (y)
		 (if (null y)
		  0 
		  (+ (len(cdr y)) 1))))
	    (define revList(lambda (ys)
		(if (null ys)
		    ys
		    (append(revList(cdr ys))(cons(car ys)(quote()))))))
	    (define butLast(lambda (ys)
		(if (=(len ys)  1)
		    (quote())
		    
		    (cons(car ys)(butLast(cdr ys))))))
	    (define member(lst x)
		(if(null lst)(if(=((car lst) x)
		true
		
		(member(cdr lst))))
		false))
		
	    (butLast(quote(a b c)) b))`))//to do изучить трассировку, разобраться в работе  lambda
	    
	    
	    
	    
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
	//
     // (reverse(quote(a b c))) = ...(append(reverse(quote( b c)))(quote(a)))...
	//(progn(defined revappend(lambda (bs ys)(if (null bs) ys (revappend (cdr bs)(cons (car bs)ys)))))(revappend(quote(a b c))(quote(d e))))
	//(progn (defined len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(a b c)))) подсчет элементов с progn
	//(list(defined len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(a b c))))
	//(list(defined len (lambda (y) (if (null y) 0 (+ (len(cdr y)) 1)))) (len (quote(1 2 3 4 5))))
	//(list(defined x (lambda (y) 25)) (x 13)
	//(list(defined cudr (lambda (y)(cdr y))) (cudr (1 2 3 4 5))) cdr
	//(list(defined sq (lambda (y) (* (* y y) y))) (sq 3))
	//fmt.Println(structVar1, structVar2, structVar3)
	//(cons (quote(a b c)))
	//(cond((= 3 3)42))
	//((lambda (x y) (+ x y)) 3 4)  (3 4)
	//((lambda (x) x) (+ 1 2) 8)

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
