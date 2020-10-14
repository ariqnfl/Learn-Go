package main

import "fmt"

func main() {
	template := "I wish I had a %v."
  pet := "puppy"
  
  var wish string
  
  // Add your code below:
  wish = fmt.Sprintf(template,pet)
  
  fmt.Println(wish)

  fmt.Println("What would you like for lunch?")
  
  // Add your code below:
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.", food)
}