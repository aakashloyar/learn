// 2️⃣ Why do we need signals in Go?

// Without signal handling:

// Ctrl+C → program stops immediately ❌


// With signal handling:

// Ctrl+C → close DB → stop server → exit cleanly ✅

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Create channel to receive signals
	ch := make(chan os.Signal, 1)

	// Notify on Ctrl+C (SIGINT)
	signal.Notify(ch, os.Interrupt)

	fmt.Println("Program running... Press Ctrl+C")

	// Block until signal received
	sig := <-ch

	fmt.Println("Received signal:", sig)
	fmt.Println("Exiting gracefully")
}
