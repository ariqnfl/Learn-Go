package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	
  rand.Seed(time.Now().UnixNano())

  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
  if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000

  switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
	  fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
	  fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
