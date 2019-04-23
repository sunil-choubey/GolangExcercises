package main

import "fmt"
import "math"

//objective of the function is to find the squareroot of a given number
// by starting with a guess and then using newton method to reach to the //required value.
//Newton method say
// Let suppose we need  to find the squareroot of x
// we start by supposing that a number Z which starts with 1
// Z -= ( Z*Z - x ) / 2*Z


func main() {

		var x float64 = 10.0
		// Print mathematical squareroot of the X
		fmt.Printf ("Mathematical square root value of %f is %f\n",x , math.Sqrt(x) )
		// Finding the squareroot using newton method
		var Y float64 = Sqrt(x)
		fmt.Printf ("Using newton method square root of %f is %f",x, Y )
		}

	func Sqrt ( a float64 ) float64 {

		var hold float64 = a
		var z float64 = 1
		var cnt int64


		if a <= 0 {
		    fmt.Printf (" Squareroot of less than equal to 0 is not supported %f\n " , a )
			return a
			}

		for i := 0 ;i < 10;i++ {
			z -= ( ( math.Pow(z, 2) - a ) / ( z * 2 ) )
			fmt.Printf("Value of z is %f\n", z )

			if diff := hold - z ; diff < 0.00001 {
				fmt.Printf ("The differenc %f is very minimal , so this is the near guess calculated using newton method \n", diff )
			fmt.Printf ( "The number of times taken to solve the problem is %d\n",cnt )
				return z
				}	else 	{
				   hold = z
				}
				cnt = cnt + 1

		  }
		 fmt.Printf ( "The number of times taken to solve the problem is %d\n",cnt )

		 return z
	}
