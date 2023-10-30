# ticket_booking_project

**Set up your MySql database**
<div>
1. Create a database called tickets_db:
   create database tickets_db;
   show databases;
   
</div>

2. Create three tables in your database for movies, show timings and tickets booked:<br>
   a. CREATE TABLE movies (<br>
      movie_id int NOT NULL PRIMARY KEY AUTO_INCREMENT, <br>
      movie_name VARCHAR(100) NOT NULL);<br>
   b. Create table show_timings (<br>
      show_id int NOT NULL PRIMARY KEY AUTO_INCREMENT,<br>
      show_time TIME,<br>
      show_date date,<br>
      screen_no Varchar(10));<br>
   c. Create table tickets_booked(<br>
      booking_id int NOT NULL PRIMARY KEY AUTO_INCREMENT,<br>
      movie_id int,<br>
      FOREIGN KEY (movie_id) REFERENCES movies(movie_id),<br>
      show_id int,<br>
      FOREIGN KEY (show_id) REFERENCES show_timings(show_id),<br>
      firstname VARCHAR(50) NOT NULL,<br>
      lastname VARCHAR(50),<br>
      confirmation_no VARCHAR(50),<br>
      status bool);<br>
4. Insert dummy values to start off:<br>
   a. INSERT INTO Movies (movie_id, movie_name) VALUES (1, 'Iron man');INSERT INTO Movies <br>
      (movie_id, movie_name) VALUES (2, 'avengers');INSERT INTO Movies (movie_id, movie_name) VALUES <br>
      (3, 'Pulp Fiction');<br>
   b. INSERT INTO show_timings (show_id, show_time, show_date, screen_no)<br>
      VALUES (1, '12:00:00', '2023-10-28', 1); INSERT INTO show_timings (show_id, show_time, <br>
      show_date, screen_no) VALUES (2, '15:30:00', '2023-10-28', 2);<br>
      INSERT INTO show_timings (show_id, show_time, show_date, screen_no) <br>
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
   

   
