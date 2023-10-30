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
  db.Table("movies").Select("*").Order("movie_id desc").Limit(1).Scan(&result)
  fmt.Println("The last movie ID is:", result["movie_id"])

  // Adding new movie
  mov_id := result["movie_id"].(int32) 
  var movie_name string
  fmt.Println("enter the name of movie")
  // the scanf function only reads single word movie name
  // got to find an alternative
  // check here - https://www.codekru.com/go/how-to-take-inputs-from-users-in-go-language
  fmt.Scanf("%s\n", &movie_name)
  db.Exec("INSERT INTO movies VALUES (" + strconv.Itoa(int(mov_id)+1) + ", '" + movie_name +  "');")
  
  // confirmation of insertion
  fmt.Println("Insertion of movie - " + movie_name + " is success!")
  
  //deleting a movie
  var movie_id_delete int
  fmt.Println("enter the movie id for the movie that you want to delete")
  fmt.Scanf("%d",&movie_id_delete)  
  db.Exec("DELETE FROM movies WHERE movie_id =" + strconv.Itoa(movie_id_delete))
  
  // confirmation of deletion
  fmt.Println("Deletion of movie ID - " + strconv.Itoa(movie_id_delete) + " is success!")
}