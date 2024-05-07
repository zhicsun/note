package pkg

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

func TestCopier(t *testing.T) {
	var (
		user     = User{Name: "tom", Age: 18, Role: "Admin", Salary: 200000, EmployeeCode: 123456}
		employee = Employee{Salary: 150000}

		users     = []User{{Name: "tom", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jerry", Age: 30, Role: "Dev", Salary: 60000}}
		employees = make([]Employee, 0)
	)

	copier.Copy(&employee, &user)
	fmt.Printf("%+v\n", employee)

	copier.Copy(&employees, &user)
	fmt.Printf("%+v\n", employees)

	employees = []Employee{}
	copier.Copy(&employees, &users)
	fmt.Printf("%+v\n", employees)

	map1 := map[int]int{3: 6, 4: 8}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)
	fmt.Printf("%+v\n", map2)
}

type User struct {
	Name         string
	Age          int32
	Role         string
	Salary       int
	EmployeeCode int64
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name       string `copier:"must"`
	Age        int32  `copier:"must,nopanic"`
	Salary     int    `copier:"-"`
	DoubleAge  int32
	EmployeeId int64 `copier:"EmployeeCode"`
	SuperRole  string
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}
