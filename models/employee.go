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

type EmpUserDetail struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Logincheck struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
