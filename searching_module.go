package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tickets_db")
    if err != nil {
        fmt.Println("Something is wrong ", err)
        return
    }
    fmt.Println("Connected")
    
    obj, err := db.Query("select * from booking_details;") // SQL is not case sensitive language
    if err != nil {
        fmt.Println("Something Wrong ", err)
    }
 
    for obj.Next() {
		
		var id int
		var name string
		var show_date_and_time, firstname, lastname, confirmation_number, status string
		
		err = obj.Scan(&id, &name, &show_date_and_time, &firstname, &lastname, &confirmation_number, &status)
		
		// handle error
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("booking ID: %d Movie name: %s\n", id, name)
		fmt.Println(show_date_and_time, firstname, lastname, confirmation_number, status)
		
	}
    defer db.Close()
}