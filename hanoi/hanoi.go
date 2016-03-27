package main

import "fmt"

func move(n, start, end int){

	if n < 1 {
		return
	}

	mid := 6-start-end
	move(n-1, start, mid)
	fmt.Println(start, "->", end)
	move(n-1, mid, end)
}

func hanoi(n int){
	fmt.Println("Number of disks :", n)
	move(n, 1, 2)
}

func main(){
	hanoi(3)
}