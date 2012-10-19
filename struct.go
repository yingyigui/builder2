package main;
import(
	"fmt"
)
func init() {
	
}
type Person struct{
	name string;
	age int;
}
type Employee struct{
	Person;
	name string;
	phone string;
}
func main() {
	var person Person;
	person.name="yingyigui";
	person.age=30;
	fmt.Println(person);
	bob:=Employee{Person{name:"ander",age:30},"ander.liu","15010116177"};
	fmt.Println(bob);
	fmt.Println(bob.name);
	fmt.Println(bob.Person.name);
}