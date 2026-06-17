//✅ 7. Retry with Delay (Backoff Logic)
package main
import (
	"fmt"
	"time"
)
func tryRequest() bool {
	// Simulate a request that fails the first two times and succeeds the third time
	return false;
}
func main() {
	for i := 1; i <= 3; i++ {
		success := tryRequest()

		if success {
			break
		}

		fmt.Println("Retrying...")
		<-time.After(2 * time.Second)
	}

}
//✅ Used in:

//Network retries

//Payment gateways

//Job polling