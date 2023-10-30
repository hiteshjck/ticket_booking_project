package main
 
import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"
  "strconv"
)
 
func main() {
	
  db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/tickets_db"))
  if err != nil {
    panic("failed to connect database")
  }
  
  // querying the last entry
  result := map[string]interface{}{}
  db.Table("tickets_booked").Select("*").Order("booking_id desc").Limit(1).Scan(&result)
  fmt.Printf("The last booking ID is: %T", result["booking_id"])
  
  // Adding new show
  bk_id := result["booking_id"].(int32) 
  
	var number int
    fmt.Println("enter your number 1 for IRON MAN")
    fmt.Println("enter your number 2 for AVENGERS")
    fmt.Println("enter your number 3 for PULP FICTION")
    fmt.Scan(&number)
    fmt.Println("number entered is", number)
    var num int
    fmt.Println("enter your number 1 for 2023-10-28")
    fmt.Println("enter your number 2 for 2023-10-28")
    fmt.Println("enter your number 3 for 2023-10-28")
    fmt.Scan(&num)
    fmt.Println("number entered is", num)
    var first_name string
    var last_name string
    fmt.Println("enter your first_name")
 
    fmt.Scan(&first_name)
    fmt.Println("enter your last_name")
    fmt.Scan(&last_name)
    
    fmt.Printf("first_name : %s\n", first_name)
    fmt.Printf("last_name : %s", last_name)
    
  db.Exec("INSERT INTO tickets_booked VALUES (" + strconv.Itoa(int(bk_id)+1) + 
  	", " + strconv.Itoa(number) + ", " + strconv.Itoa(num) + ", '" + first_name + "', '" + 
  	last_name +"', 'a123', 1);")
  
}