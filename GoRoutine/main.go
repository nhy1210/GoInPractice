package main
import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int){
	for i:=0; i< 10; i++{
		fmt.Println(n, ":", i)
		amt:= time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond*amt)
	}
}
func main(){
	for i:=0; i <5; i++ {	
		switch i{
			case 0: 
				go f(0)
			case 1:
				go f(1)
			default: 
				go f(5) 
		}
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("First program to check for Go routine")
}
