//passed all test cases

package main

import (
	"fmt"
)

func main(){
	var t, a, b, noc int
	var c[] int

	fmt.Scanln(&t)

	for k:=0; k<t; k++ {
		fmt.Scan(&a)
		fmt.Scan(&b)

		if a > b {
			noc = a/b
		} else {
			noc = b/a
		}

		c = append(c, noc)
	}

	for i:=0; i<len(c); i++ {
		fmt.Print(c[i], "\n")
	}

}