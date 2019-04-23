package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	type KeyValues struct
	{
		key string
		value int
	}

	var m map[string]int

	m = make(map[string]int)
	var keyvalue  []KeyValues
	//var s string = "Hello Sunil!"
	var newarray = strings.Fields(s)

	for v := range newarray {
		fmt.Print(newarray[v])
		counter := strings.Count( s, newarray[v] )
		fmt.Printf("\nCounter using strings.count function %d \n", counter )
		keyvalue = append ( keyvalue , KeyValues{ newarray[v],counter} )
		counter = countword ( s, newarray[v])
		fmt.Printf ("Counter using function count word %d\n", counter )
		m[newarray[v]] = counter
		}

	return m
}

func countword ( a , b string ) int {
	   var counter int = 0
	   var newarray []string = strings.Fields ( a)
	   for i := range newarray {
	   			if strings.Compare(newarray[i], b ) == 0 {
					counter++
				}
		}
		return counter
	}


func main() {
	wc.Test(WordCount)
}
