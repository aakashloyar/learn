package main 
import "fmt"


//this patter is not good as robot implements eat method and resolve it during run time 
//while it must be resolved during compile time
type Worker interface {
	Work()
	Eat()
}

type Human struct {

}
type Robot struct {

}

func (w Human)String() string {
	//return fmt.Sprintf("%v",w);//this will do stackoverflow error as it will recursively call .string method
	return "Human"
}
func (w Robot)String() string {
	//return fmt.Sprintf("%v",w);
	return "Robot"
}
func (w Human)Work() {
	fmt.Println(w.String()+" worker is working")
}
func (w Human)Eat() {
	fmt.Println(w.String()+" worker is eating")
}

func (w Robot)Work() {
	fmt.Println(w.String()+" worker is working")
}
func (w Robot)Eat() {
	panic("Robot cannot eat")
}
func main() {
	var w Worker
	w = Human{}
	w.Work()
	w.Eat()

	w = Robot{}
	w.Work()
	w.Eat() // ❌ should be impossible at compile time
}