# ticket_booking_project

**Set up your MySql database**
1. Create a database called tickets_db:
   create database tickets_db;
   show databases;
3. Create three tables in your database for movies, show timings and tickets booked:\n
   a. CREATE TABLE movies (
      movie_id int NOT NULL PRIMARY KEY AUTO_INCREMENT, 
      movie_name VARCHAR(100) NOT NULL);
   b. Create table show_timings (
      show_id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
      show_time TIME,
      show_date date,
      screen_no Varchar(10));
   c. Create table tickets_booked(
      booking_id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
      movie_id int,
      FOREIGN KEY (movie_id) REFERENCES movies(movie_id),
      show_id int,
      FOREIGN KEY (show_id) REFERENCES show_timings(show_id),
      firstname VARCHAR(50) NOT NULL,
      lastname VARCHAR(50),
      confirmation_no VARCHAR(50),
      status bool);
4. Insert dummy values to start off:
   a. INSERT INTO Movies (movie_id, movie_name) VALUES (1, 'Iron man');INSERT INTO Movies 
      (movie_id, movie_name) VALUES (2, 'avengers');INSERT INTO Movies (movie_id, movie_name) VALUES 
      (3, 'Pulp Fiction');
   b. INSERT INTO show_timings (show_id, show_time, show_date, screen_no)
      VALUES (1, '12:00:00', '2023-10-28', 1); INSERT INTO show_timings (show_id, show_time, 
      show_date, screen_no) VALUES (2, '15:30:00', '2023-10-28', 2);
      INSERT INTO show_timings (show_id, show_time, show_date, screen_no) 
      VALUES (3, '18:45:00', '2023-10-28', 3);
   c. INSERT INTO tickets_booked (booking_id, movie_id, show_id, firstname, lastname, 
      confirmation_no, status) VALUES (1, 1, 1, 'John', 'Doe', 'ABC123', 1);
      INSERT INTO tickets_booked (booking_id, movie_id, show_id, firstname, lastname, confirmation_no,
      status) VALUES (2, 2, 2, 'Alice', 'Smith', 'DEF456', 1);
      INSERT INTO tickets_booked (booking_id, movie_id, show_id, firstname, lastname, 
      confirmation_no, status) VALUES (3, 3, 3, 'Eva', 'Johnson', 'GHI789', 1);

**Usage**

go mod init ticket_booking

**Dependencies**
1. go get gorm.io/gorm
2. gorm.io/driver/mysql

1. To book tickets
   go run booking_module.go
2. To search information
   go run search_module.go
3. To perform admin tasks
   go run admin_module.go
   

   
