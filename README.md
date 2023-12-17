gofr-project
Mini GoLang framework for writing http api using gofr 
API server and Cmd creation
open the terminal in the project location and run the following command first
•	go mod init github.com/example
•	go get gofr.dev
code and create a main file main.go
run the command in terminal
•	go mod tidy
connecting mysql run the command
•	docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.30
•	docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE garage (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME);"
•	docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE complete (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME, out_time DATETIME);"

Postman collection for trying out the APIs
for inserting into garage
•	POST localhost:9000/customer/name/car_name/car_brand/car_no.
To See the list of cars currently in Garage
•	GET localhost:9000/customer
To See the list of cars leaves in Garage
•	GET localhost:9000/old
To Update the entry in DB
•	PUT localhost:9000/id/name/car_name/car_brand/car_no
To Delete the entry from DB when the car leaves the Garage
•	DELETE localhost:9000/customer/id
