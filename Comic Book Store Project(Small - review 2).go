package main
import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber, costint int
  var grade, cost float32
  
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist, publisher, year, pageNumber,grade = "Jewel Tampson", "DizzyBooks Publishing Inc.", 1997, 14, 6.5

  costint = (((2020-year)+pageNumber)*2)
  cost = float32(costint)
  cost = cost * (grade/10)

  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in ", year,"yr.\nIt has", pageNumber, "pages.\nIt's condition grade is", grade,"points.\nIt Costs ", cost, "₹\n")

  title, writer, artist, publisher, year, pageNumber, grade ="Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Inc.", 2013, 160, 9.0

  costint = (((2020-year)+pageNumber)*2)
  cost = float32(costint)
  cost = cost * (grade/10)

   fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in ", year,"yr.\nIt has", pageNumber, "pages.\nIt's condition grade is", grade,"points.\nIt Costs ", cost, "₹\n")
}