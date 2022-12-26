package storages

import (
	"fmt"
	"goecho/models"
)

func (c *cursor) Createemp(u *models.Employee) error {
	sqlStatement := "INSERT INTO employee (name, salary,age)VALUES ($1, $2, $3)"
	_, err := c.Db.Query(sqlStatement, u.Name, u.Salary, u.Age)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *cursor) Listemployee() (*models.Employees, error) {
	sqlStatement := "SELECT id, name, salary, age FROM employee order by id"
	rows, err := c.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := models.Employees{}

	for rows.Next() {
		//employee := Employees{}
		var employee models.Employee
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			return nil, err2
		}
		result.Employees = append(result.Employees, employee)
	}
	return &result, nil
}

func (c *cursor) Geteempbyid(id int) (models.Employee, error) {
	sqlStatement := "SELECT * from employee where id = $1"
	fmt.Println(sqlStatement)
	var employee models.Employee
	res := c.Db.QueryRow(sqlStatement, id)
	err := res.Scan(&employee.Name, &employee.Salary, &employee.Age, &employee.Id)
	if err != nil {
		fmt.Println(err)
	}

	return employee, nil
}

func (c *cursor) Updateeempbyid(id int, u *models.Employee) error {
	sqlStatement := "UPDATE employee SET name=$1,salary=$2,age=$3 WHERE id=$4"
	_, err := c.Db.Query(sqlStatement, u.Name, u.Salary, u.Age, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *cursor) Deleteemp(id int) error {
	sqlStatement := "DELETE FROM employee WHERE id = $1"
	_, err := c.Db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	return nil
}
