# gofr-project

REQUIREMENTS :-
MySQL , Docker Desktop , Postman

Open the terminal in project location and run the following command 

    go mod init github.com/example            (here we have to paste a link of a empty repo of users account)
    go get gofr.dev

  run the command in terminal

      go mod tidy

  ** command for connection of mySQL run the following commands
  
    docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.30

    docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE garage (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME);"

    docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE complete (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, car_name VARCHAR(255) NOT NULL, car_brand VARCHAR(255) NOT NULL, car_no VARCHAR(255) NOT NULL, in_time DATETIME, out_time DATETIME);"

    go run main.go 

  ## Now To check , use postman ##

  -to insert data of cars ,type the following url using POST method

      localhost:9000/customer/name/car_name/car_brand/car_no
      
      The above url resembles with the one given below
      localhost:9000/customer/{name}/{car_name}/{car_brand}/{car_no}

  -to read data of cars ,type the following url using GET method

      localhost:9000/customer

  -to read data of old cars which left the garrage ,type the following url using GET method

      localhost:9000/old
    
  -to update data of a car ,type the following url using PUT method

      localhost:9000/customer/1/name/car_name/car_brand/car_no

      The above url resembles with the one given below
      localhost:9000/customer/{id}/{name}/{car_name}/{car_brand}/{car_no}

-to delete data of a car ,type the following url using DELETE method

      localhost:9000/customer/1

      The above url resembles with the one given below
      localhost:9000/customer/{id}

 
