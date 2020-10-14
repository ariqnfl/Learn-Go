package main

import "fmt"

func main() {
	var congrats string
  congrats = "Congratulations"
  congrats += "!!!"
  fmt.Println(congrats)
  
  var challenge string = "What else can you do?"
  fmt.Println(challenge)
  
  reminder := "Pratice is important!"
  fmt.Println(reminder)

  var publisher,writer,artist,title string
  var year,pageNumber int
  var grade float32

  title = "Mr.GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages", "and have rating", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn."
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 16
  grade = 9.0
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages", "and have rating", grade)

}