package main

import (
	"beginnerGo/internal/app"
	"beginnerGo/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		// ! panic is keyword to stop the program immediately
		panic(err)
	}

	// ! close the database connection when main function ends
	// ! defer will ensure the function is called at the end of the surrounding function
	defer app.DB.Close()

	app.Logger.Println("Logging from now..")
	fmt.Println(("Hello frontend master!"))

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running on the port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
	// sample function for reference purpose
	// Tutorial("Golang", "Golang is an open-source programming language.")
}

type Location struct {
	latitude  int
	longitude float64
}

// func tutorial(){}
// * define CAPITAL for function to be public
// * define small letter for function to be private
// * function have to define params + return type
func Tutorial(title, description string) string {
	// constant declaration
	const pi = 3.14
	// default variable declaration
	var fakeName string = "Golang"
	// shorthand variable declaration
	fakeName2 := "Golang"
	// multiple variable declaration
	var (
		myName string = "Jason"
		age    int    = 30
		isCool bool   = true
	)

	// print line for print in new line
	fmt.Println(fakeName, fakeName2)

	// printf for formatted string, need manual newline
	fmt.Printf("%s is %d years old. Cool? %t\n", myName, age, isCool)

	// switch case declaration
	switch myName {
	case "Jason":
		fmt.Println("Hello Jason")
	case "Bob", "Alice":
		fmt.Println("Hello Bob")
	default:
		fmt.Println("Hello Stranger")
	}

	// if else declaration
	if isCool {
		fmt.Println("You are cool")
	} else {
		fmt.Println("You are not cool")
	}

	// for loop declaration
	for i := 0; i < 5; i++ {
		fmt.Println("This is current index:", i)
	}

	// while loop declaration
	// will auto loop until condition is false
	// * in loop, break keyword can break the loop immediately
	// * in loop, continue keyword can skip the current iteration and continue to next iteration
	counter := 0
	for counter < 5 {
		fmt.Println("This is current index:", counter)
		counter++
	}

	// Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("This is current index:", i, "with value:", numbers[i])
	}
	// slice is dynamic array
	firstSlice := []string{"apple", "banana", "cherry"}
	secondSlice := []string{"durian", "elderberry", "fig"}

	// * append function to add new element to slice
	// * spread operator ... to append one slice to another slice
	secondSlice = append(secondSlice, firstSlice...)
	for i, v := range secondSlice {
		fmt.Println("This is current index:", i, "with value:", v)
	}

	// * ":" starting from index
	// * "x:" starting from index x to end
	// * ":x" starting from index 0 to x-1
	// sample below to remove index 2 from slice
	// thirdSlice = append(secondSlice[:2], secondSlice[3:]...)

	// map declaration
	capitals := map[string]string{
		"name":    "Jason",
		"country": "USA",
		"city":    "New York",
	}
	for key, value := range capitals {
		fmt.Println("This is key:", key, "with value:", value)
	}

	// first return is value, second return is boolean to check if key exists
	capital, exists := capitals["name"]
	if exists {
		fmt.Println("The capital is:", capital)
	}

	// * can delete key from map using delete function
	delete(capitals, "city")

	// struct declaration
	type Person struct {
		name    string
		age     int
		country string
	}
	person := Person{name: "Jason", age: 30, country: "USA"}
	fmt.Println("This is person:", person.name, person.age, person.country)

	// anonymous struct declaration
	anonymousPerson := struct {
		name    string
		age     int
		country string
	}{
		name:    "Anonymous",
		age:     0,
		country: "Unknown",
	}
	fmt.Println("This is anonymous person:", anonymousPerson.name, anonymousPerson.age, anonymousPerson.country)

	// sample for updating struct by passing pointer
	location := Location{latitude: 10, longitude: 20.5}
	fmt.Println("Before function call:", location.latitude, location.longitude)
	// * pass by reference using &
	updateSimpleStruct(&location)
	fmt.Println("After function call:", location.latitude, location.longitude)

	// method receiver showcase
	location.updateStructWithReceiver(50, 60.5)
	fmt.Println("After method receiver call:", location.latitude, location.longitude)

	return "success"
}

// update struct by passing pointer
// * receive pointer using "*"
func updateSimpleStruct(l *Location) {
	l.latitude = 100
	l.longitude = 200.5
}

// method receiver defined after 'func' and before function name
// * receiver can be pointer or value
// * receiver is like 'this' or 'self' in other languages
// * receiver can access struct fields and methods
// * receiver ensure its tightly coupled with the struct, only struct can call this method
func (l *Location) updateStructWithReceiver(latitude int, longitude float64) {
	l.latitude = latitude
	l.longitude = longitude
}
