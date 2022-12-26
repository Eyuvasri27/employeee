package models

//structures are given

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}
type Employees struct {
	Employees []Employee `json:"employees"`
}
