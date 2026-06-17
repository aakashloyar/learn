//http server
package main
import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello, World!")//io.writer output goes to client not console
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers: range r.Header {
		for _, h:= range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/hello",hello)//handleFunc -> register handler function for given pattern// with signature func(w,r)
	//http.Handle("/hello", http.HandlerFunc(hello))//same
	http.HandleFunc("/headers",headers)

	http.ListenAndServe(":8080",nil)
}


// w http.ResponseWriter 
// r *http.Request

// type ResponseWriter interface {
//     Header() Header
//     Write([]byte) (int, error)
//     WriteHeader(statusCode int)
// }
// this is an interface 

// type Request struct {
//     Method string
//     URL    *url.URL
//     Header Header
//     Body   io.ReadCloser
//     ...
// }
// // this is a large struct



//when to use struct vs interface
//use struct when you have a concrete type with fields and methods
//use interface when you want to define a contract that other types can implement
//means you donot care how but you care about what methods are implemented


//type Speaker interface { Speak() string }
//  type Person struct { Name string } 
// func (p Person) Speak() string { return "Hello, my name is " + p.Name } 
// var s Speaker s = Person{Name: "Aakash"} 
// fmt.Println(s.Speak())
// s is not a object it is an interface variable



// r *http.Request — pointer

// Request is a large struct (headers, body, URL, context…)

// Passing a pointer:

// avoids copying all that data

// ensures all handlers/middleware share the same request

// allows reading the body safely

// ✅ Rule: Pass big structs or things you want to share as pointers.


// w http.ResponseWriter — interface

// ResponseWriter is an interface, not a struct

// Interfaces are already references internally

// Writing through it directly updates the underlying connection

// No pointer needed — adding * would be wrong

// ✅ Rule: Interfaces are already like pointers. Pass them as values.




//type Header map[string][]string




// Alternative using http.Handler
// type MyHandler struct{}

// func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello from MyHandler")
// }

// func main() {
// 	http.Handle("/myhandler", MyHandler{})
// 	http.ListenAndServe(":8080", nil)
// }


// type CounterHandler struct {
// 	count int
// }

// func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	c.count++
// 	fmt.Fprintf(w, "Visitor number: %d", c.count)
// }

// func main() {
// 	handler := &CounterHandler{}
// 	http.Handle("/counter", handler)
// 	http.ListenAndServe(":8080", nil)
// }
//now see this can handle state using struct fields
// http.Handler → More structured, better for complex or stateful handlers.

// http.HandleFunc → Simpler, quicker for stateless endpoints.


// What companies actually use in production

// At scale, large applications typically use frameworks like:

// • Gin
// • Echo
// • Chi
// • Fiber (inspired by Express)

// These libraries give you:
// ✔ Routing (like /users/:id)
// ✔ Middleware (auth, logging, tracing)
// ✔ JSON handling, validation
// ✔ Better performance and cleaner code
