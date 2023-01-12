# rest-go-1
First Restful API with Go, Gin And Gorm

This is my first tryout project with Go, the rest using Gin as HTTP web web framework, also using GORM as Obeject Relational Maping (ORM).

#How To Use This

- First install go lang enviroment : https://go.dev/doc/install
- Clone these repo
- Install all dependencies : go mod download
- Install nodemon as golbally, i use nodemon to live reload file changes while develop : npm install -g nodemon
- Create your .env file by copy from .env.example and customize base on your preferences.
- Run this project : make run.

#There's 5 endpoints I've write, its about create, read, update, and delete book.
- GET /books        => Get All Books
- GET /books/:id    => Get Single Book
- POST /books       => Create New Book
- UPDATE /books/:id => Update Stored Book
- DELETE /books/:id => Remove Book from databse

This project i write by tutorial from https://blog.logrocket.com/rest-api-golang-gin-gorm/ , Thanks you a lot for nice tutorial.
