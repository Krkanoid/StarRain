package main

import (
    "fmt"
	"time"
    "math/rand"
)

func main() {
	fmt.Print("Press enter to continue.")
	fmt.Scanln()
	
	for i := 0; i < 10000; i++ {
		space := [...]int{
			4,3,6,8,10,13,18,21,22,25,
			29,31,33,38,40,43,46,48,51,
		}
		symbol := [...]string{"✵","✶","✸","☆","."}
		rSpace := rand.Intn(len(space))
		rSymbol := rand.Intn(len(symbol))
		item := space[rSpace]
		sItem := symbol[rSymbol]

		for j := 0; j < item; j++ {
			fmt.Print(" ")
		}

		if(i / 10 == 10 || i / 10 == 1) {
			fmt.Println(`░`)
		fmt.Print(sItem)
		} else {
			fmt.Println(sItem)
		}
		time.Sleep(time.Duration(1000))
	}
}
