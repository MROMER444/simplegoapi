Simple Go API

This project is a simple RESTful API written in Go using the Gin web framework. The API provides endpoints for managing members, performing basic mathematical operations, and checking user inputs.
Table of Contents

    Members API
    Mathematical Operations API
    User Input API
    Data Model
    Sample Data
    How to Run

Members API
1. Get all Members

    Endpoint: GET /members
    Description: Retrieve a list of all members.
    Response Format:


    200 OK
    [
      {
        "ID": "1",
        "Name": "alex",
        "Age": 23,
        "Major": "backend"
      },
      // ... other members
    ]

2. Add a Member

    Endpoint: POST /addmember
    Description: Add a new member to the list.
    Request Format:


{
  "ID": "5",
  "Name": "new member",
  "Age": 30,
  "Major": "data science"
}

Response Format:


    200 OK
    {
      "ID": "5",
      "Name": "new member",
      "Age": 30,
      "Major": "data science"
    }

3. Get Member by ID

    Endpoint: GET /getmemberbyid/:id
    Description: Retrieve a specific member by ID.
    Response Format:


    200 OK
    {
      "ID": "2",
      "Name": "mikle",
      "Age": 21,
      "Major": "frontend"
    }

4. Update Member by ID

    Endpoint: PUT /updatememberbyid/:id
    Description: Update a specific member's details by ID.
    Request Format:


{
  "ID": "2",
  "Name": "updated mikle",
  "Age": 25,
  "Major": "updated frontend"
}

Response Format:

    200 OK
    {
      "ID": "2",
      "Name": "updated mikle",
      "Age": 25,
      "Major": "updated frontend"
    }

Mathematical Operations API
5. Add

    Endpoint: GET /add/:num1/:num2
    Description: Perform addition of two numbers.
    Response Format:


    200 OK
    {
      "result": 10
    }

6. Divide

    Endpoint: GET /div/:number1/:number2
    Description: Perform division of two numbers.
    Response Format:


    200 OK
    {
      "result": 2
    }

User Input API
7. Check

    Endpoint: GET /check/:userinput
    Description: Check user input and provide a response.
    Response Format:


    200 OK
    {
      "msg": "u r right",
      "success": true
    }

8. Password Validation

    Endpoint: GET /pass/:inp
    Description: Validate a password based on a custom pattern.
    Response Format:


    200 OK
    {
      "password": "example",
      "isvalid": true
    }

Data Model

    The member struct represents a member with properties such as ID, Name, Age, and Major.

go

type member struct {
  ID    string `validate:"required"`
  Name  string `validate:"required,min=3,max=20"`
  Age   int    `validate:"required,gte=18,lte=60"`
  Major string `validate:"required"`
}

Sample Data

    Sample member data is provided in the Members slice within the data.go file.

go

var Members = []member{
  {ID: "1", Name: "alex", Age: 23, Major: "backend"},
  {ID: "2", Name: "mikle", Age: 21, Major: "frontend"},
  {ID: "3", Name: "test", Age: 27, Major: "fullstack"},
  {ID: "4", Name: "adam", Age: 43, Major: "mobileapplication"},
}

How to Run

    Install Go and Gin if not already installed.
    Clone the repository.
    Navigate to the project directory.
    Run the following command to start the server:

    bash

    go run main.go

The server will run on localhost:8080. Use the provided API endpoints to interact with the API.
