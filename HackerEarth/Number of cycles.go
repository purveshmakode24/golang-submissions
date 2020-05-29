//passed all test cases

package main

import (
	"fmt"
)

func main(){
	var n, p int
	var cnt[] int
	fmt.Scanln(&n)

	for i:=0; i<n; i++ {
		fmt.Scanln(&p)
		cnt = append(cnt, p*(p-1)+1)
	}

	for i:=0; i<len(cnt); i++{
		fmt.Print(cnt[i], "\n")
	}
}