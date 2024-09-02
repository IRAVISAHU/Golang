package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Login struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

var db *sql.DB

func initDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/employee_management"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected to the database")
}

// it is used like global variable but it is not changeable
const PI = 3.14

// struct
type Address struct {
	Name    string
	City    string
	Pincode int
}

// Creating structure
type Student struct {
	Name   string
	Branch string
	Year   int
}

// Creating nested structure
type Teacher struct {
	Name    string
	Subject string
	Exp     int
	Details Student
}

// Creating a structure
// with anonymous fields
type Students struct {
	int
	string
	float64
}

func pointer(a *int) {
	*a = 748
}

type Employee struct {
	Name  string
	Empid int
}

func display(str string) {
	for w := 0; w < 3; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func sending(s chan<- string) {
	s <- "Hello"
}

func portal1(channel1 chan string) {

	// time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// function 2
func portal2(channel2 chan string) {

	// time.Sleep(9 * time.Second) 
	channel2 <- "Welcome to channel 2"
}
func main() {
	// var is used like global variable and := it is used inside function
	// vaiable declaration with initial value
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred
	fmt.Println(x)
	fmt.Println(student1)
	fmt.Println(student2)

	// vaiable declaration without initial value
	var a string
	fmt.Println(a)
	// initialize value
	a = "John"
	fmt.Println(a)

	// The :=  sign operator can only be used within functions and not for declaring package-level (global) variables. if we declare outside function this give  syntax error: non-declaration statement outside function body

	// initialize and declare both same time it can be changed
	b := "new value"
	fmt.Println(b)

	// Multiple variable
	var t, u = 6, "Hello"
	c, d := 7, "World!"
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(c)
	fmt.Println(d)
	// variable Declaration in a Block
	var (
		m int
		v int    = 1
		w string = "hello"
	)
	fmt.Println(m)
	fmt.Println(v)
	fmt.Println(w)

	// const global
	fmt.Println(PI)

	// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
	// arrays have a fixed length.
	// var array_name = [length]datatype{values} // here length is defined

	// or

	// var array_name = [...]datatype{values} // here length is inferred

	// 	sign :=
	// 	array_name := [length]datatype{values} // here length is defined

	// or
	// array_name := [...]datatype{values} // here length is inferred
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	fmt.Println(cars[0])

	// Initialize Only Specific Elements
	arr3 := [5]int{1: 10, 2: 40}
	fmt.Println(arr3)
	fmt.Println(len(cars))

	// Slice are arrays are same but slice length can grow and shrink,it powerful and flexible then array.,it also store data in same type and single variable
	// len() function - returns the length of the slice
	// cap() function - returns the capacity of the slice
	// 1 Create a Slice With []datatype{values}
	// slice_name := []datatype{values}
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	arr9 := myslice2[1:3]
	fmt.Println(arr9)

	//2 Create a Slice From an Array
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array
	arr12 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr12[2:4]

	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	// Note: If the capacity parameter is not defined, it will be equal to length.

	//3 Create a Slice With The make() Function
	// slice_name := make([]type, length, capacity)
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// access slice
	prices := []int{10, 20, 30}
	fmt.Println(prices[1])
	// change slice
	prices[2] = 50
	fmt.Println(prices[2])

	// append()function: to add element in last in slice
	// slice_name = append(slice_name, element1, element2, ...)
	// it store inside another slice
	pricess := append(prices, 40, 21)
	fmt.Println(pricess)
	fmt.Println(len(pricess))
	fmt.Println(cap(pricess))
	// Memory Efficiency
	// The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
	// copy(dest, src)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	fmt.Println(numbersCopy)

	// iterating for loop
	for e := 0; e < len(numbersCopy); e++ {
		fmt.Println(numbersCopy[e])
	}
	// iterating using range
	for index, element := range numbersCopy {
		fmt.Println(index, element)
	}
	// without index with us _ blank space
	for _, element := range numbersCopy {
		fmt.Println(element)
	}

	// if you made any changes in the slice, then it will also reflect in the array
	slc := numbersCopy[0:3]
	slc[0] = 100
	slc[1] = 200
	fmt.Println(slc)
	fmt.Println(numbersCopy)
	fmt.Println(slc == nil)

	// Multi-dimensional slice are just like the multidimensional array, except that slice does not contain the size.

	Marry := [][]int{
		{1, 2, 3},
		{99, 3, 44},
		{33, 44, 55},
	}
	fmt.Println(Marry)

	// Composite literals are used to construct the values for arrays, structs, slices, and maps.Each time they are evaluated, it will create new value. They consist of the type of the literal followed by a brace-bound list of elements.
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// sort bt Ints function
	intSlice := []int{-23, 567, -34, 67, 0, 12, -5}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// IntsAreSorted it returns true if they are sorted
	sort.IntsAreSorted(intSlice)
	fmt.Println(intSlice)

	// Trim() function  - what i want to remove from the string
	// func Trim(ori_slice[]byte, cut_string string) []byte
	// This function returns a subslice of the original slice

	Name := []byte{'!', '3', 'r', 'a', 'v', 'i', '4', 'l'}
	res := bytes.Trim(Name, "!34l")
	fmt.Println(string(res))

	// Split function -This function splits a slice of bytes into all subslices separated by the given separator and returns a slice which contains all these subslices.
	// split function remove what i want in string ,thi remove from last string
	// func Split(o_slice, sep []byte) [][]byte

	newres := string(Name)
	fmt.Println(strings.Split(newres, "vi4l"))

	// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.
	// 1 . In structure fields name starts with capital letter
	// Declaring a structure:
	// type Address struct {
	// 	name, street, city, state string
	// 	Pincode int
	// }

	// var a Address
	var addres Address
	fmt.Println(addres)
	addres = Address{"Akshay", "Dehradun", 3623572} // Just assigning a new value to an existing variable
	fmt.Println(addres)
	fmt.Println(addres.City)

	// nested structure field
	// a structure within another structure is known as a Nested Structure
	result := Teacher{
		Name:    "Suman",
		Subject: "Java",
		Exp:     5,
		Details: Student{"Bongo", "CSE", 2},
	}

	fmt.Println(result)

	// An anonymous structure is a structure which does not contain a name. It useful when you want to create a one-time usable structure.
	// variable_name := struct{
	// 	 fields
	// 	}{Field_values}
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}
	fmt.Println(Element)

	// Assigning values to the anonymous
	// fields of the student structure
	value := Students{123, "Bud", 8900.23}
	fmt.Println(value)
	// Pointer - it store the memory address of another variable
	// it starting with 0x like 0xFFAAF etc
	// Declaration and Initialization of Pointers
	// var pointer_name *Data_Typ

	// eg ....
	var ab = 45
	var s *int = &ab
	fmt.Println("Value of ab : ", ab)
	fmt.Println("Address of ab : ", &ab)
	fmt.Println("Address of s : ", s)
	fmt.Println("Value of s : ", *s)

	var st *int
	fmt.Println("s = ", st)
	fmt.Println("st = ", st)

	var y = 458
	var p = &y
	fmt.Println("Value stored in pointer variable p = ", p)
	fmt.Println("Address of y = ", &y)

	// passing pointer in function
	// Passing an address of the variable to the function

	var rs = 100
	fmt.Println("Rupee amount ", rs)

	var pa *int = &rs
	pointer(pa)
	fmt.Println("Rupee amount after function call ", rs)

	//  pointer comparision
	var data = "ravi"
	var data1 = &data
	var data2 = &data
	fmt.Println(data1 == data2)

	// Double Pointer -this looks like a chain of pointers. When we define a pointer to pointer then the first pointer is used to store the address of the second pointer.
	var V int = 100
	var pt1 *int = &V
	var pt2 **int = &pt1
	fmt.Println(V)
	fmt.Println(&V)
	fmt.Println("address", pt1)
	fmt.Println(*pt1)
	fmt.Println(pt2)
	fmt.Println(*pt2)
	fmt.Println(**pt2)
	**pt2 = 233
	fmt.Println(**pt2)
	emp := Employee{"ABC", 19078}
	pts := &emp
	var ptst = &pts
	fmt.Println(**ptst)

	// Concurrency- A Goroutine is a function or method that executes independently and simultaneously in connection with any other Goroutines in your program.

	// Every program contains at least a single Goroutine and that Goroutine is known as the main Goroutine. All the Goroutines are working under the main Goroutines if the main Goroutine is terminated, then all the Goroutines in the program are also terminated.

	// 1 goroutine function execute afer main goroutine
	// Sleep() method in our program which makes the main Goroutine sleeps for 1 second in between 1-second the new Goroutine executes,
	// after 1-second main Goroutine re-schedule and perform its operation
	// 	Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
	// Goroutines can communicate using the channel
	// we can create multiple goroutines
	go display("Ravi")
	display("sahu")

	// Anonymous Goroutine -it has no name a.  its called also anonymous function . it is written in main goroutine

	// declarataion

	go func() {
		time.Sleep(8 * time.Second)
		fmt.Println("Welcome!")
	}()

	//  a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel

	// Creating a Channel - var Channel_name chan Type  , channel_name:= make(chan Type)

	// Send operation -Mychannel <- element
	// Receive operation: - element := <-Mychannel
	// Closing a Channel- (1) close()  ,(2) ele, ok:= <- Mychannel

	// Unidirectional Channel  -  By default a channel is bidirectional but you can create a unidirectional channel also. A channel that can only receive data or a channel that can only send data is the unidirectional channel. The unidirectional channel can also create with the help of make() function

	// mychanl1 := make(<-chan string)
	// mychanl2 := make(chan<- string)
	// fmt.Printf("%T", mychanl1)

	// fmt.Printf("\n%T", mychanl2)
	mychanl := make(chan string)
	go sending(mychanl)
	fmt.Println(<-mychanl)

	R1 := make(chan string)
	R2 := make(chan string)
	go portal1(R1)
	go portal2(R2)
	select {
	case opt1 := <-R1:
		fmt.Println("Portal 1:", opt1)
	case opt2 := <-R2:
		fmt.Println("Portal 2:", opt2)
	}

	initDB() 
	http.HandleFunc("/postlogin", createLogin)
	http.HandleFunc("/dellogin", deleteLogin)
	http.HandleFunc("/updateLogin", updateLogin)
	http.HandleFunc("/getdata", getdata)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createLogin(w http.ResponseWriter, r *http.Request) {
	var login Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	_, err := db.Exec("INSERT INTO login (id, password) VALUES (?, ?)", login.Id, login.Password)

	if err != nil {
		http.Error(w, "Database insert error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Login created successfully")
}

func deleteLogin(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM login WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Database delete error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id %s deleted successfully", id)
}

func updateLogin(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var login Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	_, err := db.Exec("UPDATE login SET password = ? WHERE id = ?", login.Password, id)
	if err != nil {
		http.Error(w, "Database update error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id %s updated successfully", id)
}

func getdata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data []Login
	rows, err := db.Query("SELECT id, password FROM login")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var login Login
		if err := rows.Scan(&login.Id, &login.Password); err != nil {
			http.Error(w, "Error processing database row", http.StatusInternalServerError)
			return
		}
		data = append(data, login)
	}
	json.NewEncoder(w).Encode(data)
}
