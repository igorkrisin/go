
package main
import "fmt"
import ("strconv")


type list struct {

data interface{}
nextdata *list // * указатель (тип)

}

func printList(xs interface{}) {
	//fmt.Print("(")
	switch tempEl := xs.(type){
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
			return v ==  str
			
		default:
			return false					
	}
	 
}

//(1 (2 4)) -> "(" "1" "(" "2" "4" ")" ")"  ->  
//stack: {1 {2, 4}}

func listReverse(xs *list) *list {
	
	
	var tempList *list = nil
	for xs != nil{
		tempList = &list {data: xs.data, nextdata: tempList}
		xs = xs.nextdata
	}
	return tempList
}

func lenList(xs *list) int {
	count := 0
	for xs != nil {
		xs = xs.nextdata
		count ++
	}
	return count
}

func parse(arr []string) interface{} {
	var stack *list  = nil 
	for i:=0; i < len(arr); i++ {
		if !equalEl(arr[i], ")") {
			j, err := strconv.Atoi(arr[i])
			fmt.Println("Atoi: ", j, err, arr[i])
			if err != nil {
				stack = &list {data: arr[i], nextdata: stack}
				fmt.Println("\n", "stack: ")
				printList(stack)
			} else {
			stack = &list {data: j, nextdata: stack}
			}
		} else if equalEl(arr[i], ")") {
			var tempList *list = nil
			for  lenList(stack) != 0 && !equalEl(stack.data, "(") {
				tempList = &list {data: stack.data, nextdata: tempList}
				fmt.Println("\n", "tempList: ")
				printList(tempList)
				stack = stack.nextdata
				//fmt.Println("stack: ",stack[:len(stack)-1])
			}
			stack = stack.nextdata
			stack = &list {data: tempList, nextdata: stack}
		}
	}
	printList( stack)
	fmt.Println("\n", "return: ")
	return stack.data//некий костыль
}

func tokenize(data string) []string {
	storeStr := ""
	arr := []string {}
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
				storeStr =""
			}
		} else {
			storeStr += data[i:i+1]
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

func evalList(tempEl *list, dict map[string]interface{}) *list {

    var tempList *list = nil
    tempEl = tempEl.nextdata 
    for tempEl != nil {   
		tempList = &list{data: eval(tempEl.data, dict), nextdata: tempList}
		tempEl = tempEl.nextdata
    }
    return listReverse(tempList)
} 

func evalListRecur(tempEl *list, dict map[string]interface{}) *list {
	if tempEl != nil {
		return  &list{data: eval(tempEl.data, dict), nextdata: evalListRecur(tempEl.nextdata, dict)}
		
	}
	return tempEl
}




//ДЗ сделать функцию evalList рекурсивной, отрезание аргумента добавить при вызове функции
// почитать про словари в го (map)
// null -- сделать функция в eval


//(+ x 42)

func eval(xs interface{}, dict  map[string]interface{}) interface{} {
	switch tempEl := xs.(type){
		case *list:
			if  equalEl(tempEl.data, "+") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type) {
					case int:
						switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {	
							case int:
								return el1 + el2
						}
				}
			} else if equalEl(tempEl.data, "if") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type){
					case string:
						if el1 == "true" {
							return eval(tempEl.nextdata.nextdata.data, dict)
						} else if el1 == "false" {
							return eval(tempEl.nextdata.nextdata.nextdata.data, dict)
						}
				}
			} else if  equalEl(tempEl.data, "-") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type) {
					case int:
						switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {	
							case int:
								return el1 - el2
						}
				}
			} else if  equalEl(tempEl.data, "*") {
				switch el1 := eval(tempEl.nextdata.data,dict).(type) {
					case int:
						switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {	
							case int:
								return el1 * el2
						}
				}
			} else if  equalEl(tempEl.data, "/") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type) {
					case int:
						switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {	
							case int:
								return el1 / el2
						}
				}
			} else if  equalEl(tempEl.data, "=") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type) {
					case int:
						switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {	
							case int:
								if el1 == el2{
									return "true"
								} else {
									return "false"
								}
						}
				}
			} else if  equalEl(tempEl.data, "quote") {
					return tempEl.nextdata.data

			} else if equalEl(tempEl.data, "car") {
				//fmt.Println("car")
				switch el1 := eval(tempEl.nextdata.data, dict).(type){
					case *list:
					return el1.data	}
			}  else if equalEl(tempEl.data, "cdr") {
				switch el1 := eval(tempEl.nextdata.data, dict).(type){
					case *list:
					return el1.nextdata	}
			} else if equalEl(tempEl.data, "cons") {
				var el1 interface{} = eval(tempEl.nextdata.data, dict)
				switch el2 := eval(tempEl.nextdata.nextdata.data, dict).(type) {
					case *list:
					    var tempList *list = &list{data: el1, nextdata: el2}
					return tempList
					}
					
			} else if equalEl(tempEl.data, "list") {
				tempEl = tempEl.nextdata
				fmt.Print("list print")
				return evalListRecur(tempEl, dict)
		
				//return evalList(tempEl)
				/* var tempList *list = nil
				tempEl = tempEl.nextdata //вызов функции без первого элемента
				for tempEl != nil {   
				    tempList = &list{data: eval(tempEl.data), nextdata: tempList}
				    tempEl = tempEl.nextdata
				}
				return listReverse(tempList)  */
			} else if equalEl(tempEl.data, "null") {
				switch el1 := tempEl.nextdata.data.(type) {
					case *list:
						if lenList(el1) != 0 {
							return "false"
						} else {
							return "true"
						}
				}
				
			} 
			
			switch el1 := tempEl.data.(type) {
			    case *list: 
					var clue interface{} = el1.nextdata.data
					var val *list = tempEl.nextdata
					if equalEl(el1.data, "lambda") {
						
						printList(clue)
						fmt.Println("clue fitst")
						printList(val)
						fmt.Println("val first")
					}
				
					switch el3 := clue.(type) {
					case *list:
						
						for el3 != nil && val!= nil {
							switch el4 := el3.data.(type) {
							case string:
								switch el5 := val.data.(type){
								case interface{}:
								
									dict[el4] = el5
									for key,val := range dict {
										fmt.Println("key: ",key,"val: ", val)
								}
								}
							}
							val = val.nextdata
							el3 = el3.nextdata
						
						}
						
/* 						switch el6 := el1.nextdata.nextdata.data.(type){
						case *list:
							eval(el6.data,dict)
							fmt.Println(el6.nextdata.data)
								

							
						}
						printList(eval(el1.nextdata.nextdata.data, dict))
 */					}
					fmt.Println(dict)
			    
			}
				//cons 23 (1 2) -> (23 1 2)
			
				
		case int:
			return xs
	}
	return xs
}

// ((lambda (x y) (+ x y)) 3 4)


func main() {

dict := make(map[string]interface{})

structVar1 := list {data: "1", nextdata: nil} //
structVar2 := list {data: "2", nextdata: &structVar1} // * - разименование
structVar3 := list {data: "3", nextdata: &structVar2}
s4 := list {data: &structVar3, nextdata: nil}
s5 := list {data:"5", nextdata: &s4}
s6 := list {data:"6", nextdata: &s5}

printList(&s6)   // (6 5 (3 2 1))

printList(listReverse(&s6))
//fmt.Print("EQUALMAIN",")" == ")")
//fmt.Println(tokenize())
printList(eval(parse(tokenize("((lambda (x y) (+ x y)) 3 4)")), dict))

//fmt.Println(structVar1, structVar2, structVar3)

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
