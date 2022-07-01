# Introduction 
This is a service to handle user profile related process on users authentication

# Getting Started
1. 	Depedencies

run this to download all go dependencies
```bash
go mod vendor
```

2.	Run the service
```bash
go run main.go
```

3. Create DB MySQL
CREATE DATABASE `lemonilo`

4. Create Table
CREATE TABLE users (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	username varchar(15),
	email varchar(30),
	address varchar(50),
	password varchar(20)
)