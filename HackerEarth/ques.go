package main

import (
	"fmt"
)

func main() {
	var n, x, y, temp, flag int
	var a[] int
	fmt.Scan(&n, &x, &y)

	for i:=0; i<n; i++ {
		fmt.Scan(&temp)

		a = append(a, temp)
	}

	for i:=0; i<n; i++ {
		if a[i]<x || a[i]>y {
			flag = 0
			break
		}else {
			flag = 1
		}
	}

	if flag == 0 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}