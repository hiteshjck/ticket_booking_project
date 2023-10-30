package main
 
import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"
  "strconv"
)
 
func main() {
	
  db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/tickets_db?parseTime=true"))
  if err != nil {
    panic("failed to connect database")
  }
  
  // querying the last entry gor getting latest booking id
  result := map[string]interface{}{}
  db.Table("tickets_booked").Select("*").Order("booking_id desc").Limit(1).Scan(&result)
  //fmt.Printf("The last booking ID is: %T", result["booking_id"])
  
  // query for movie names and taking movie choice
  result1 := []map[string]interface{}{}
  db.Table("movies").Select("movie_name").Scan(&result1)
  var movie_choice int
  for i:=1; i<=len(result1); i++ {
  	fmt.Println("Enter " + strconv.Itoa(i) + " for " + result1[i-1]["movie_name"].(string))
  }
  fmt.Printf("\nEnter your choice: ")
  fmt.Scanf("%d\n", &movie_choice)
  
  // query for show timings and taking show choice
  result2 := []map[string]interface{}{}
  db.Table("show_timings").Select("*").Scan(&result2)
  var show_choice int
  for i:=1; i<=len(result2); i++ {
  	fmt.Printf("Enter " + strconv.Itoa(i) + " for:")
  	fmt.Println(result2[i-1]["show_date"])
	fmt.Println(result2[i-1]["show_time"])
	  }
  fmt.Printf("\nEnter your choice: ")
  fmt.Scanf("%d\n", &show_choice)
  
  var first_name string
  var last_name string
  fmt.Printf("\nEnter your first_name: ") 
  fmt.Scan(&first_name)
  fmt.Printf("\nEnter your last_name: ")
  fmt.Scan(&last_name)
  
  bk_id := result["booking_id"].(int32) 
  db.Exec("INSERT INTO tickets_booked VALUES (" + strconv.Itoa(int(bk_id)+1) + 
  	", " + strconv.Itoa(movie_choice) + ", " + strconv.Itoa(show_choice) + ", '" + first_name + "', '" + 
  	last_name +"', 'a123', 1);")
  
  fmt.Println("Ticket booked succesfully!")
}
