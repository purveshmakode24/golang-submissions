// All test cases passed

package main

import (
    "fmt"
    "math"
)

func prime(num int) bool{
    if num < 2 {
        return false
    }else if num == 2 {
        return true
    }else {	
        for i:=2; i<=int(math.Sqrt(float64(num))); i++ {
            if num%i == 0 {
                return false
                break
            }
        }
        return true
    }
}

func main(){

    var n int
    fmt.Scanln(&n)

    for num:=0; num<n; num++ {
        if prime(num) == true {
            fmt.Print(num, " ")
        }
    }

}
