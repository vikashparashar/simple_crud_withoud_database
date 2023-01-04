package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// var Users = make(map[string]User)

// Users ["1"] = User{ID: 1, Name: "Vikash", Age: 27}

var Users = []User{
	{ID: 1, Name: "Vikash", Age: 27},
	{ID: 2, Name: "Pawan", Age: 19},
	{ID: 3, Name: "Romeo", Age: 32},
}
