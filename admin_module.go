package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"
  "strconv"
  "encoding/json"
)

// Adding new movie 
func addMovie(db *gorm.DB) {

  // querying the last entry
  result := map[string]interface{}{}
  db.Table("movies").Select("*").Order("movie_id desc").Limit(1).Scan(&result)
  fmt.Println("The last movie ID is:", result["movie_id"])

  // Adding new movie
  mov_id := result["movie_id"].(int32)
  var movie_name string
  fmt.Println("\nEnter the name of movie: ")
  fmt.Scanf("%s\n", &movie_name)
  db.Exec("INSERT INTO movies VALUES (" + strconv.Itoa(int(mov_id)+1) + ", '" + movie_name +  "');")

  // confirmation of insertion
  fmt.Println("Insertion of movie - " + movie_name + " is success!")
}

// Deleting a movie 
func deleteMovie(db *gorm.DB) {

  //deleting a movie
  var movie_id_delete int
  fmt.Printf("\nEnter the movie id for the movie that you want to delete: ")
  fmt.Scanf("%d\n",&movie_id_delete)  
  db.Exec("DELETE FROM movies WHERE movie_id =" + strconv.Itoa(movie_id_delete))

  // confirmation of deletion
  fmt.Println("Deletion of movie ID - " + strconv.Itoa(movie_id_delete) + " is success!")
}

// searching
func searchTicket(db *gorm.DB) {
	
	// getting user choice
	var search int
	var search_val string
	my_map := map[int]string{1: "booking_id", 2: "movie_id", 3: "show_id", 
									4: "firstname", 5: "lastname",}
	fmt.Printf("\n")
	for key, val := range my_map{
		fmt.Println("Enter " + strconv.Itoa(key) + " to search with " + val)}
	fmt.Printf("\nEnter your choice: ")
	fmt.Scanf("%d\n", &search)	
 	fmt.Printf("\nEnter the " + my_map[search] + " value: ")
 	fmt.Scanf("%s\n", &search_val)
 	
	// Query to get matching results
	result := map[string]interface{}{}
	db.Table("tickets_booked").Select("*").Where(my_map[search] + "=?", search_val).Limit(1).Scan(&result)
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	if string(b) == "{}"{
		fmt.Printf("\nData does not exist!")
		return
	}
	fmt.Printf("\nData found!\n")
	fmt.Println(string(b))
}
 
func main(){
	
	// connecting to database
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/tickets_db"))
	if err != nil {
		panic("failed to connect database")
	}
	
	// choice is taken from the user
	var choice int
	start:
	
	switch choice {
		case 1:
			addMovie(db)
			choice = 0 
			goto start

		case 2:
			deleteMovie(db)
			choice = 0 
			goto start

		case 3:
			searchTicket(db)
			choice = 0 
			goto start
			
		case 4:
			break
		
		default:
			fmt.Printf("\nWelcome Admin!\n" + 
			"Enter 1 for adding a movie\n" +
			"Enter 2 for deleting a movie\n" +
			"Enter 3 for searching ticket booking info\n" +
			"Enter 4 to exit\n\n" + 
			"Enter your choice: ")
			fmt.Scanf("%d\n",&choice)
			goto start
	}
}