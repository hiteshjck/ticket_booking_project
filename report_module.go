package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"
  "encoding/json"
  "os"
)

func main(){
	
	// connecting to database
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/tickets_db"))
	if err != nil {
		panic("failed to connect database")
	}
	
	result := []map[string]interface{}{}
	db.Table("tickets_booked").Select("*").Scan(&result)
	f, _ := os.Create("output.txt")
	
	x, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Data read\n")
	//fmt.Println(string(x))
	
	for i:=0; i<len(result); i++ {
		y, err := json.Marshal(result[i])
		if err != nil {
			panic(err)
		}
		f.Write(y)
		f.WriteString("\n")
	}
	
	fmt.Printf("\nReport generated successfully!")
}
