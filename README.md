# GO-Gin-Boilerplate

This boilerplate is build with traditional web applications and microservices in Go using the [Gin](https://github.com/gin-gonic/gin "Gin") framework with [MongoDB](https://www.mongodb.com/ "MongoDB") as a database.

Gin is a framework that reduces boilerplate code that would normally go into building these applications. It also lends itself very well to creating reusable and extensible pieces of code.

This boilerplate will help you to set up your projects with following prerequisites while developing the rest API.

  - Environments config
  - Logger
  - Auto-reload while changing the code with [Air](https://github.com/cosmtrek/air "air")
  - Database connection
  - [Swagger](https://swagger.io/ "Swagger")


------------


### Installation
##### Prerequisites:

You will need the following prerequisites installed on your machine.
1. Go ([Setup Go Enviourment](https://golang.org/doc/install "Setup Go Enviourment"))
1. MongoDB ([Setup MongoDB Enviourment](https://docs.mongodb.com/manual/administration/install-community/ "Setup MongoDB Enviourment"))

##### Install Dependencies
Install all Go.mod  dependecies 

    go list -m all

Install Swagger 

    go get github.com/go-swagger/go-swagger/cmd/swagger

Install Air

    go get -u github.com/cosmtrek/air

------------

### RUN 

Run the application just write follwoing command

    air

##### For Production Mode

Create the follwing enviourment varibale



    APP_MODE = prod


------------


