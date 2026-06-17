//now this http client 
package main
import (
	"bufio"
	"fmt"
	"net/http"
	"log"
)
func main(){
	resp,err:=http.Get("https://gobyexample.com")
	if err!=nil {
		log.Fatal("Error:",err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:",resp.Status)
	scanner:=bufio.NewScanner(resp.Body)
	for i:=0;i<5 && scanner.Scan();i++ {
		fmt.Println(scanner.Text())
	}//scanner does not return an error and donot exit if error occurs //also it reads line by line
	//it skips if error happens
	if err := scanner.Err(); err != nil {
        panic(err)
    }

}

//http.get makes a network connection 
//resp.Body is like a stream (file)
//stream must be closed after use
//if not closed it will lead to resource leak 
// and network connection will remain open
//so we will have to close it using resp.Body.Close()




// Scan() hides the error

// Error is stored internally

// Must be checked after scanning finishes

//network connection dropped, invalid respone, read failure



//why the connection is still open even after we have got resp
//because resp.Body is a stream and it is not completely come in 1 request
//so we have to read from the stream until the end
// so the connection is still open



// for {
//     resp, _ := http.Get(url)
//     // forgot resp.Body.Close()
// }//in tis case //error too many open files



// func main() {
//     http.Get("https://example.com")
// }
// When the program terminates:

// OS will reclaim memory

// OS will close file descriptors

// TCP connections are closed

// So eventually, resources are released.

// 👉 BUT this is not safe or acceptable design.

//this will be ok


// in the above case use function inside for loop to use defer properly
// for {
//     func() {
//         resp, err := http.Get(url)
//         if err != nil {
//             return
//         }
//         defer resp.Body.Close()

//         body, err := io.ReadAll(resp.Body)
//         if err != nil {
//             return
//         }

//         fmt.Println(len(body))
//     }()
// }



//scanner vs reader
//scanner is for line by line reading
//small inputs
//reader is for large inputs
//reader reads a chunk at a time
//reader is more complex to use
//scanner is easier to use
//scanner has a buffer limit of 64KB by default
//reader has no buffer limit
//scanner is built on top of reader
//scanner is more convenient for line by line reading
//reader is more flexible for custom reading	