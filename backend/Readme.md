To start server: go run cmd/main.go
Index page url: http://localhost:8080

This app is written in layers format keeping in mind testability of code, 
There are three layers, 
    Routes
    Controller
    Models
    
**Routes** takes care of routing the request to appropriate controller

**Controller** gets the required data from models and responds to the queries

**Models** specifies the objects and methods to get them


  