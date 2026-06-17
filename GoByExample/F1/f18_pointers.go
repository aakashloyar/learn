package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    // i := 1
    // fmt.Println("initial:", i)

    // zeroval(i)
    // fmt.Println("zeroval:", i)

    // zeroptr(&i)
    // fmt.Println("zeroptr:", i)

    // fmt.Println("pointer:", &i)

	x:=42
	y:=42;
	z:=&x;
	fmt.Println(&x);
	fmt.Println(&y);
	fmt.Println(x==y)
	//&x=21;//error
	fmt.Println(z);
	fmt.Println(*z)
	*z=21;
}