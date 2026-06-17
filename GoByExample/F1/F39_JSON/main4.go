//JSON → map[string]interface{}
//when json structure is not known


package main
import (
	"fmt"
	"encoding/json"
)


func main() {
	data := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	fmt.Println(m)

	x := m["num"]
	fmt.Println(x)
	num := m["num"].(float64)
	//fmt.Println(m["strs"][0])//cannot do directly
	strs := m["strs"].([]interface{})

	fmt.Println(num)
	fmt.Println(strs[0].(string))


}

//⚠️ JSON numbers → float64 by default.