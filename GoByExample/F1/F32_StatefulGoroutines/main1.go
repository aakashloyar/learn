//do not communicate by sharing memory; share memory by communicating.
//do not let multiple goroutine touch same data
//do let 1 goroutine to own the data and let other send message to it 


// stateful goroutine:

// Owns some data (state) internally

// Only that goroutine is allowed to read/write that state

// Other goroutines interact with it only via channels

// This avoids:

// Race conditions

// Mutex complexity

// Shared-memory bugs

//A goroutine that behaves like a mini server or actor with private memory.



// Why Use Stateful Goroutines?
// Problems with Shared State (Bad Pattern)
// var counter int

// go func() {
//     counter++ // ❌ data race
// }()

// go func() {
//     counter++ // ❌ data race
// }()


// You now need:

// sync.Mutex

// Careful lock discipline

// Risk of deadlocks

//everyone editing same excel
//locking row conflicting everywhere

//one accoutant own the sheet 
//other just send request
//accountant update sequentailly
//accountant=stateful goroutine