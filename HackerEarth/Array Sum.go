// All test cases passed

package main

import (
	"fmt"
)

func main() {

	var n, nums int
	var a[] int
	su:=0

	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		fmt.Scan(&nums)
		a = append(a, nums)
	}

	for i:=0; i<n; i++ {
		su = su + a[i]
	}

	fmt.Println(su)

}