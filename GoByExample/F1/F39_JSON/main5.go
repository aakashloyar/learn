//JSON → map[string]interface{}
//when json structure is not known


package main
import (
	"fmt"
	"encoding/json"
)
type Data struct {
	ID     int      `json:"id"`
	Fruits []string `json:"fruits"`
}



func main() {
	res := &Data{ID: 1}
    fmt.Println(res)


	b, _ := json.MarshalIndent(u, "", "  ")
    fmt.Println(string(b))

}

//⚠️ JSON numbers → float64 by default.