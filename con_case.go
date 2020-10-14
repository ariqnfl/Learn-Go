package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Still not a good work, keep up")
  } else {
    isHeistOn = false
    fmt.Println("Better next time?")
  }

  fmt.Println("Heist is currently :",isHeistOn)
  openedVault := rand.Intn(100)

  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn && openedVault < 70 {
    isHeistOn = false
    fmt.Println("Better luck next time!")
  }
  leftSafely := rand.Intn(5)
  amtStolen := 10000 + rand.Intn(1000)
  if isHeistOn {
    switch leftSafely {
      case 0:
      isHeistOn = false
      fmt.Println("Failed")
      case 1:
      isHeistOn = false
      fmt.Println("Waduh bro...")
      case 2:
      isHeistOn = true
      fmt.Println("Nice! u stole $",amtStolen)
      case 3:
      isHeistOn = true
      fmt.Println("Nice! u stole $",amtStolen)
      default:
      isHeistOn = true
	  fmt.Println("Nice! u stole $",amtStolen)
    }
  }
}
