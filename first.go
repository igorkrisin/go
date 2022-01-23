package main
import "fmt"
import "C"


func fact (numb int) int {
	if numb == 1 {
		return 1
	} else {
		return fact(numb - 1) * numb 
	}
}


func fib (numb int) int {
	if numb == 1 || numb == 2 {
		return 1
	} else {
		return fib(numb - 1) + fib(numb - 2)
	}
}

func fibMem(numb int, dict map[int]int) int {
	val, ok:= dict[numb]
	if ok {
		return val
	} else if numb == 1 || numb == 2 {
		dict[numb] = 1
		return 1
	} else {
		temp := fibMem(numb - 1, dict) + fibMem(numb - 2, dict)
		print(temp)
		dict[numb] = temp
		return temp
	}
}

func ack(m int, n int) int {
	fmt.Println(m, n)
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return ack(m - 1, 1)
	}
	var n2 = ack(m, n - 1)
	return ack(m - 1, n2)

}


func perm(str string) []string {
	if str == "" {
		return [] string {""}
	} else {
		
		arr:= []string{}
		for i:= 0; i<len(str); i++ {
			temp:= perm(str[:i]+str[i+1:])
			for j:= 0; j<len(temp); j++ {
				arr = append(arr, temp[j]+str[i:i+1])
			}
		}
	return arr
	}
}


func main() {
	//fmt.Println("Hello World")
	//fmt.Println(fact(4))
	//fmt.Println(fib(5))

	//for i:= 1; i < 51; i++ {
	//	fmt.Println(fib(i))
//	}
     /*m := make(map[string]int)
     m["Python"]=1990
     m["Lisp"]=1960
     fmt.Println(m["Lisp"])
     fmt.Println(m["Go"])
     val,ok:=m["Go"]
     fmt.Println(ok, val)*/

     //fmt.Println(fibMem(200, make(map[int]int)))
   	// fmt.Println(ack(3,4))
   	
	//fmt.Println(perm("abc123"))
	//fmt.Print("\033[H", "adb")

	/* arr := [] int{1, 2, 3}
	int1 := 6
	arr = append(arr , int1)
	fmt.Print(arr) */
	carta := map[string][]int {
		"1": {1, 2},
		"2": {2,10},
	}

	for key,val := range carta {
		fmt.Println(key,val)
		
		
	}
	



/*

1*2*3=6
perm("abc") => {"abc","cba","acb", "bac","cab", "bca"}
perm("bc") => {"bc", "cb"} => {"abc","acb"}
perm("ac") => {"ac", "ca"} => {"bac", "bca"}
per,(ab) => {"ab", "ba"} => {"cab", "cba"}

1!=1
2!=2
3!=1*2*3=6
4!=1*2*3*4=24=3!*4*/
}