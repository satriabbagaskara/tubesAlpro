package main

import "fmt"
type arr struct{
	x int
}
var jmhArr [100] arr
func cari(n int){
	var i, a, left, right, mid int
	var found bool = false
	for i = 0; i < n; i++{
		fmt.Scan(&jmhArr[i].x)
	}
	fmt.Scan(&a)
	left = 0
	right = n-1
	for left <= right && !found{
		mid = (left + right) / 2
		if jmhArr[mid].x == a{
			fmt.Println("Ketemu di index", mid+1)
			found = true
		} else if a < jmhArr[mid].x{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found{
		fmt.Println("Tidak ketemu")
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	cari(n)
}
