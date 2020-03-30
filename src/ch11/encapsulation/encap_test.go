package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployee(t *testing.T) {
	employeeA := Employee{"0", "Bob", 18}
	employeeB := Employee{Name: "Mike", Age: 30}
	employeeC := new(Employee)
	employeeC.Id = "2"
	employeeC.Age = 15
	employeeC.Name = "Employee C"
	t.Log(employeeA)
	t.Log(employeeB)
	t.Log(employeeC)
	t.Logf("Employee A %T", employeeA)
	t.Logf("Employee A %T", &employeeA)
	t.Logf("Employee C %T", employeeC)

}

/**
结构体 会进行复制
*/
func (e Employee) String() string {
	fmt.Printf("Address of Employee e is %x ", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s - Name: %s - Age : %d", e.Id, e.Name, e.Age)
}

/**
指针操作没有进行拷贝
*/
//func (e *Employee) String() string {
//	fmt.Printf("Address of Employee e is %x", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s - Name: %s - Age : %d", e.Id, e.Name, e.Age)
//}

/**
使用指针和使用类型
*/
func TestStructOperations(t *testing.T) {
	//e := &Employee{"0", "Bob", 20}
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address of Employee e is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
