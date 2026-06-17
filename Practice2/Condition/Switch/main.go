package main
import "fmt"

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
        return 20.0
	case "enterprise":
        return 50.0	
	default:
		return 0.0
	}
}
func main() {
	fmt.Println("hi")
	plan := "basic"
	fmt.Println(billingCost(plan))
	//fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}


