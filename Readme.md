# Majoo BE Test
## Simple API for Majoo BE Test
Backend API USING GO , GORM, Postgre , with seeder and schema auto migrate


## Instalation
**I already hosting postgre db in this app the credential can be viewed at .env files so you dont have to migrate it in your local mechine**
-Clone this project 
-go mod tidy
-go run main.go


## Note
- This app using postgre and gorm but i make seeder and automigrate using gorm so you just need to connect using your favorite database


## Features

- Login Function
- Login API with  jwt
- Report function
- Report API with header auth


## Routes 
http://127.0.0.1:8000/login [POST]

http://127.0.0.1:8000/report?page=1&page_size=5 [GET]

Pada file ini akan disertakan postman collection dan dokumentasi API

