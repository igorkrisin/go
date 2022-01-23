/*  
4 |     |   677 |
    | 90  |       |
    |     |       |
    |  123|       |
func transpose(matrix [][]int) [][]int {
	arrNew := []int{}
	doubleMatr := [][]int{}
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			arrNew = append(arrNew, matrix[y][x])
		}
		doubleMatr = append(doubleMatr, arrNew)
		arrNew = []int{}
	}
	return doubleMatr
}

func getRight(matrix[][] int) [][]int {
	return getDown(transpose(transpose(transpose(getDown(transpose(matrix))))))
}

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os/exec"
	"os"
	
)

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < 1 {
				fmt.Print("      |")
			} else {
				fmt.Printf("% 5d |", matrix[i][j])
			}

			/*f matrix[i][j] < 10 {
				fmt.Print("   ", matrix[i][j], "|")
			} else if matrix[i][j] < 100 {
				fmt.Print("  ", matrix[i][j], "|")
			} else if matrix[i][j] < 1000 {
				fmt.Print(" ", matrix[i][j], "|")
			} else {
				fmt.Print(matrix[i][j], "|")
			}*/

		}
		fmt.Println()
	}
	//fmt.Println("\033[H")
}

func appInt(matrix [][]int, num int) [][]int {

	y := rand.Intn(len(matrix))
	x := rand.Intn(len(matrix[y]))
	fmt.Println("m[y][x]", matrix[x][y])
	if matrix[y][x] == 0 {
		num = rand.Intn(2048)
		matrix[y][x] = num
		return matrix
	}

	return appInt(matrix, 0)
}

func appUpgrateInt(matrix [][]int, num int) [][]int {
	arrTempX := []int{}
	arrTempY := []int{}
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				arrTempY = append(arrTempY, y)
				arrTempX = append(arrTempX, x)
			}
		}
	}
	numbEl := rand.Intn(len(arrTempX))
	matrix[arrTempY[numbEl]][arrTempX[numbEl]] = num
	return matrix

}

func getDown(matrix [][]int) [][]int {
	var count int
	for true {
		count = 0
		for y := 0; y < len(matrix)-1; y++ {
			for x := 0; x < len(matrix[y]); x++ {
				if matrix[y][x] != 0 && matrix[y+1][x] == 0 {
					matrix[y][x], matrix[y+1][x] = matrix[y+1][x], matrix[y][x]
					count++
				}
			}
		}
		if count == 0 {
			break
		}
	}
	return matrix
}

func transpose2(matrix [][]int) [][]int {
	arrNew := []int{}
	doubleMatr := [][]int{}
	for x := len(matrix) - 1; x >= 0; x-- {
		for y := len(matrix[x]) - 1; y >= 0; y-- {
			arrNew = append(arrNew, matrix[y][x])
			//fmt.Println("arrNew" ,arrNew)
		}
		for i, j := 0, len(arrNew)-1; i < j; i, j = i+1, j-1 {
			arrNew[i], arrNew[j] = arrNew[j], arrNew[i]
		}
		
		doubleMatr = append(doubleMatr, arrNew)
		arrNew = []int{}

	}
	
	return doubleMatr
}

func getRight(matrix[][] int) [][]int {
	return transpose2(getDown(summInt(getDown(transpose2(transpose2(transpose2(matrix)))))))
}
func getLeft(matrix[][] int) [][]int {
	return (transpose2(transpose2(transpose2(getDown(summInt(getDown(transpose2(matrix))))))))
}

func getUp(matrix[][] int) [][]int {
	return transpose2(transpose2(getDown(summInt(getDown(transpose2(transpose2(matrix)))))))
}

func gameEnd(matrix [][]int) bool {
	var count int
	for true {
		count = 0
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[y]); x++ {
				if matrix[y][x] == 0 {
					count++
				}
			}
		}
		if count == 0 {
			return true
		} else {
			break
		}
	}
	return false	
}

func gameWin(matrix [][]int) bool {
	
		
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[y]); x++ {
				if matrix[y][x] == 2048 {
					fmt.Println("You WIN")
					return true
				} 
			}
		}
	
	return false	
}

func summInt(matrix [][]int) [][]int {
	for y := 0; y < len(matrix) - 1; y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] != 0 && matrix[y + 1][x] == matrix[y][x] {
				matrix[y + 1][x] = matrix[y][x] + matrix[y + 1][x]
				matrix[y][x] = 0
			}
		}
	}
	return matrix
}





// функция, проверяет что игра закончилась. Если в матрице нет 0, то игра закончена.

func main() {

	/*submit := exec.Command("ls")
	result, codError := submit.Output()
	fmt.Println(string(result)) 
	fmt.Println(codError)
	//fmt.Printf("You pressed: %q\r\n", char)*/

	submit := exec.Command("stty -F /dev/tty cbreak min 1")
	submit.Run()
//	os.stdin.read()

	var b []byte = make([]byte, 1)
    LOOP:  // метка
    for {
        os.Stdin.Read(b)
        var n uint8 = b[0]
        switch n {
        case 27, 113, 81 : fmt.Println("ESC, or 'q' or 'Q' was hitted!")
            break LOOP
        default:
            fmt.Printf("You typed : %d\n", n)
        }
    }

	

	rand.Seed(time.Now().UnixNano())
	const num = 2
	arr := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	arrWork := appUpgrateInt(appUpgrateInt(arr, num), num)
	
	
	//printMatrix(arr)
	//fmt.Println(appInt(arr, 0)
	//printMatrix(arr)
	//print("\n", "/////////////////", "\n")
	//printMatrix(getDown(arr))
	//fmt.Println(appUpgrateInt(arr, 2))
	//fmt.Println(transpose(arr))
	//printMatrix(transpose(getDown(transpose(arr))))
	//fmt.Println("not for")
	for !gameEnd(arrWork) && !gameWin(arr) {
		//fmt.Println(gameWin(arrWork))
		var arrows string
		printMatrix(arrWork)
		fmt.Scanf("%s", &arrows)
		if arrows == "a" {
			arrWork = appUpgrateInt(getLeft(arrWork), num)
		} else if arrows == "d" {
			arrWork = appUpgrateInt(getRight(arrWork), num)
		} else if arrows == "s" {
			arrWork = appUpgrateInt(getDown(summInt(getDown(arrWork))), num)
		} else if arrows == "w" {
			arrWork = appUpgrateInt(getUp(arrWork), num)
		}
		fmt.Println("\033[K", "\033[H", "\033[2J")
	} 
	/* fmt.Println(arrWork)
	fmt.Println(gameWin(arrWork))
	fmt.Println(gameEnd(arrWork)) */
}
