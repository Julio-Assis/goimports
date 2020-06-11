package main

import (
    "demoproject/src/model"
    "fmt"
)

func main() {
    fmt.Println("Main function")

    var employee = model.Employee{
        Id: 1,
        FirstName: "First Name",
        LastName: "Last Name",
        BadgeNumber: 1000,
    }

    fmt.Printf(employee.FirstName)
}
