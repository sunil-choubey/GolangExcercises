package main

import (
"golang.org/x/tour/pic"
"fmt"
"math"
)

func Pic(dx, dy int) [][]uint8 {
	x := 0
	y := 0
 	var newarray = make ([][]uint8 ,dy)
	fmt.Printf ( "Length of Newarray %d " , len(newarray))

	for  v  := range newarray {

		newarray[v] = make ( []uint8 , dx )
		for i := 0 ; i < dx ; i++ {
			newarray[v][i] = uint8( x * y )
			}
		fmt.Printf ( "Value of array is %d\n", newarray[v])
	    x++
	    y++
	}


	return newarray
}

func main() {
	pic.Show(Pic)

 	TestArray :=  ShowArray(3,25)
	 for test := range TestArray {
	 	fmt.Printf ( "Array Values are %d ",TestArray[test] )
			for test1 := 0; test1 < 25 ; test1 ++ {
				fmt.Printf ("Values of Inner slices are %d \n", TestArray[test][test1] )
			}
		}

}

func ShowArray ( x , y int ) [][]int {
	fmt.Printf("Input length provided are %d and %d " , x, y )
	newarray := make ( [][]int , x )

	for v := range newarray {

		//newarray[v] = []int { v ,v }
		newarray[v] = make([]int , y )
		for i := 0; i < y ; i++ {
			newarray[v][i] = int(math.Sqrt(float64(i)))
			}
		}
	return newarray
	}
	
