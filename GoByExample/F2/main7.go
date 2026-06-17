package main 
import "fmt"

type Direction string
const (
	east = "123"
	west 
	north 
	south
)


type Status int

const (
    Pending Status = iota << 1//now this will make these 2 times
    Approved
    Rejected
	Last 
)

func main() {
	x:=Direction(east)
	fmt.Println(x)
	y:=Direction(west)
	fmt.Println(y)

	z1:=Status(Pending)
	z2:=Status(Approved)
	z3:=Status(Rejected)
	z4:=Status(Last)
	fmt.Println(z1,z2,z3,z4)

}


// why do we use enum 
// let us think of work pending, approved, rejected
// as suppose we are switching between cases
// then we call a function updatestatus then we can do some wrong casing
// then it will get rejected
// instead use enum
// func updateOrder(status Status) {}
// so switching between cases enum helps here

// now another use case is also switching between cases
// type Role string

// const (
//     Admin Role = "ADMIN"
//     User  Role = "USER"
// )
