package main
import "fmt"

type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct
  wheel struct {
    radius int
    material string
  }
}
type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car //embedded struct
  bedSize int
}
func main() {
	//anonymous struct
	var myCar = car{
	brand:   "Rezvani",
	model:   "Vengeance",
	doors:   4,
	mileage: 35000,
	wheel: struct {
		radius   int
		material string
	}{
		radius:   35,
		material: "alloy",
	},
	}
	fmt.Println(myCar)
}