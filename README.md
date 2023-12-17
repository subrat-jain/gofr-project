# gofr-project <br />
Mini GoLang framework for writing http api using gofr <br />
API server and Cmd creation <br />
open the terminal in the project location and run the following command first <br />
•	go mod init github.com/example <br />
•	go get gofr.dev <br />
code and create a main file main.go <br />
run the command in terminal <br />
•	go mod tidy <br />
connecting mysql run the command <br />
•	docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.30 <br />
•	docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE garage (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME);" <br />
•	docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE complete (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME, out_time DATETIME);" <br />

Postman collection for trying out the APIs <br />
for inserting into garage <br />
•	POST localhost:9000/customer/name/car_name/car_brand/car_no. <br />
To See the list of cars currently in Garage <br />
•	GET localhost:9000/customer <br />
To See the list of cars leaves in Garage <br />
•	GET localhost:9000/old <br />
To Update the entry in DB <br />
•	PUT localhost:9000/id/name/car_name/car_brand/car_no <br />
To Delete the entry from DB when the car leaves the Garage <br />
•	DELETE localhost:9000/customer/id <br />
