package main
import "fmt"


type Person struct {
	Name string
	Age int
}
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func (p Person) NHaveBirthday() {
    p.Age++
}

func (p *Person) HaveBirthday() {
    p.Age++
}

func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
} 

func main() {
	p:=Person {
		Name: "Alice",
		Age: 30,
	}
	fmt.Println(p)
	// fmt.Println(p.Name)
	// fmt.Println(p.Age)
	
	// p1:=Person{"Bob",25}
	// fmt.Println(p1)

	p2:=Person{}
	p2.Name="Charlie"
	p2.Age=28
	//fmt.Println(p2)
	fmt.Printf("%s\n", p2.String()) //prints field names too
	p2.NHaveBirthday();
	fmt.Printf("After NHaveBirthday: %s\n", p2.String()) //age will not change
	p2.HaveBirthday();//go automatically takes address
	(&p2).HaveBirthday();//cannot do directly
	p5:=&p2;
	p5.HaveBirthday();
	fmt.Printf("After HaveBirthday: %s\n", p2.String()) //age will change

	// dog := struct {
    //     name   string
    //     isGood bool
    // }{
    //     "Rex",
    //     true,
    // }
    // fmt.Println(dog)

	// type Men struct {
	// 	name   string
	// 	age    int
	// 	isGood bool
	// }
	// m:=Men{"John",40,true}

	// fmt.Println(m)

	// p4:=newPerson("Diana",35)
	// fmt.Printf("New Person: %s\n", p4.String())


}




// &p2.HaveBirthday()
// Go parser reads it like:

// less
// Copy code
// &(p2.HaveBirthday())
// Which means:

// call p2.HaveBirthday()

// then take address of the result

// But your method returns nothing, so it's invalid.
