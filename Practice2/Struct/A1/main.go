package main
import "fmt"
type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel//nested struct
    backWheel wheel
}
type wheel struct {
  radius int
  material string
}
func main() {
	myCar := car{}
    myCar.frontWheel.radius = 5
	fmt.Println(myCar.frontWheel.radius)
}