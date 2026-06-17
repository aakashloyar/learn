package main
import "fmt"

func main() {
	i:=1;
	for i<=3 {
		fmt.Println(i);
		i++;
	}
	for i:=2;i<=5;i++ {
		fmt.Println(i);
	}

	for i:=range 5 {
		fmt.Println(i);
	}//0->4
	for {
		fmt.Println("loop");
		break;
	}
}
