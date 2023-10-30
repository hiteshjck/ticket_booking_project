package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Adding new show
func addShow(showTime, showDate, ScreenNo string, db *gorm.DB) {

	// querying the last entry
	result := map[string]interface{}{}
	db.Table("show_timings").Select("show_id").Order("show_id desc").Limit(1).Scan(&result)
	fmt.Println("The last show ID is:", result["show_id"])

	// Adding new show
	show_id := result["show_id"].(int32)

	db.Exec("INSERT INTO show_timings VALUES (" + strconv.Itoa(int(show_id)+1) +
		", '" + showTime + "', '" + showDate + "' , " + ScreenNo + ")")

	// confirmation of insertion
	fmt.Println("Insertion of show - " + showTime + showDate + " is success!")
}

func editShow(db *gorm.DB) {
	my_map := map[int]string{1: "show_timing", 2: "show_Date", 3: "screen_no"}
	fmt.Printf("\n")
	var showid string
	fmt.Println("Enter the show_id that you want to edit")
	fmt.Scanf("%s\n", &showid)
	for key, val := range my_map {
		fmt.Println("Enter " + strconv.Itoa(key) + " to replace  " + val)
	}
	var choice int
	fmt.Scanf("%d\n", &choice)
	var update_val string
	fmt.Printf("\nEnter your value to be updated: ")
	fmt.Scanf("%s\n", &update_val)

	db.Exec("update show_timings set " + my_map[choice] + "=" + update_val + " where show_id = " + showid)

	// Confirmation of update
	fmt.Println("Show ID - " + showid + " has been updated successfully!")
}

func deleteShow(show_id_delete int, db *gorm.DB) {

	db.Exec("DELETE FROM show_timings WHERE show_id =" + strconv.Itoa(show_id_delete))

	// confirmation of deletion
	fmt.Println("Deletion of show ID - " + strconv.Itoa(show_id_delete) + " is success!")
}

func getUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err == nil && (choice >= 1 && choice <= 4) {
			return choice
		}
		fmt.Println("Invalid choice. Please try again.")
	}
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	var choice, id int
	var showTime, showDate, screenNo string
	// connecting to database
	db, err := gorm.Open(mysql.Open("root:kAshish@2301@tcp(127.0.0.1:3306)/tickets_db"))
	if err != nil {
		panic("failed to connect database")
	}

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Show")
		fmt.Println("2. Edit Show")
		fmt.Println("3. Delete Show")
		fmt.Println("4. Exit")

		choice = getUserChoice()

		switch choice {
		case 1:
			showTime = getUserInput("Enter show time:(hh:mm:ss) ")
			showDate = getUserInput("Enter show date:(yyyy-mm-dd) ")
			screenNo = getUserInput("Enter screen number: ")
			addShow(showTime, showDate, screenNo, db)
		case 2:
			editShow(db)
		case 3:
			id, _ = strconv.Atoi(getUserInput("Enter show ID to delete: "))
			deleteShow(id, db)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println("Shows:")
		for _, show := range shows {
			fmt.Printf("ID: %d, Name: %s, ShowTime: %s, ShowDate: %s, ScreenNo: %s\n",
				show.ID, show.Name, show.ShowTime, show.ShowDate, show.ScreenNo)
		}
	}
}
