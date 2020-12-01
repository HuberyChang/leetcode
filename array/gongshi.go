package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
)

func sum(q, r float64) float64 {
	R1 := 14 * q / (math.Pow(r, 3))
	//fmt.Println(R1)
	//R2 := 4.3 + math.Pow(q, 0.6666666666666)/math.Pow(r, 2)
	R2 := 4.3 * math.Cbrt(math.Pow(q, 2)) / math.Pow(r, 2)
	//fmt.Println(math.Cbrt(math.Pow(q, 2)), math.Pow(r, 2), math.Cbrt(math.Pow(q, 2))/math.Pow(r, 2))
	//fmt.Println(R2)
	//R3 := 1.1 + math.Pow(q, 0.3333333333333)/r
	R3 := 1.1 * math.Cbrt(q) / r
	//fmt.Println(math.Cbrt(q), math.Cbrt(q)/r)
	//fmt.Println(R3)
	sum := (R1 + R2 + R3) / 10
	return sum
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	var R, Q float64
	fmt.Println("请输入Q：")
	fmt.Scanln(&Q)
	fmt.Println("请输入R：")
	fmt.Scanln(&R)
	fmt.Printf("结果是：%.2f", sum(Q, R)) 
	<-done
}
