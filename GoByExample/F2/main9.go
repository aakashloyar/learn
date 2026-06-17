package main 
import "fmt"

type Equaler interface {
	Equal(other any) bool
}

type MyInt int

func (this MyInt) Equal(other any) bool {
	if o, ok := other.(MyInt); ok {
		return this==o
	}
	return false;
}

func IsEqual(a, b Equaler) bool {
	return a.Equal(b);
}

func main() {
	fmt.Println(IsEqual(MyInt(2),MyInt(2)));
}