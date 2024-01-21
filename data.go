package main


type member struct {
	ID    string `validate:"required"`
	Name  string `validate:"required,min=3,max=20"`
	Age   int    `validate:"required,gte=18,lte=60"`
	Major string `validate:"required"`
}

var Members = []member{
	{ID: "1", Name: "alex", Age: 23, Major: "backend"},
	{ID: "2", Name: "mikle", Age: 21, Major: "frontend"},
	{ID: "3", Name: "test", Age: 27, Major: "fullstack"},
	{ID: "4", Name: "adam", Age: 43, Major: "mobileapplication"},
}