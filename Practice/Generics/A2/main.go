package main
import "fmt"

type sendingChannel string

const (
    Email sendingChannel = "email"
    SMS   sendingChannel = "sms"
    Phone sendingChannel = "phone"
)

func main() {
	fmt.Println(Email)
}