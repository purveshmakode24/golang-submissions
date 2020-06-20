//Passed all test cases

package main

import(
	"fmt"
)

func main() {
	var n, a, temp int
	var arr[] int
	
	fmt.Scan(&n)

	for i:=0; i<n; i++ {
	   	fmt.Scan(&a)
		arr = append(arr, a)
	}

	cnt := 0
	swapped := true

	for swapped != false {
		swapped = false
		cnt = cnt + 1
		
		for i:=0; i<n-1; i++ {
			if arr[i]>arr[i+1] {
				temp = arr[i+1]
            			arr[i+1] = arr[i]
            			arr[i] = temp
            			swapped = true
			}
		}
	}
	fmt.Print(cnt)
}
