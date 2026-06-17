package main 
import "fmt"
type Cache interface {
	Set(a string, b int)
	Get(a string) int
}

type InMemoryCache struct {
	m map[string]int 
}
type ReadOnlyCache struct {
	inner Cache
}

func (c InMemoryCache) Set(key string, val int) {
	c.m[key]=val
}

func (c InMemoryCache) Get(key string) int{
	return c.m[key]
}

func (c ReadOnlyCache) Get(key string) int{
	return c.inner.Get(key)
}

func (c ReadOnlyCache) Set(key string, val int) {
	panic("Not possible")
}
func main() {
	var x Cache = &InMemoryCache{m: make(map[string]int)}
	x.Set("aakash",1)
	var y Cache = &ReadOnlyCache{inner: x}
	fmt.Println(y.Get("aakash"))
	//fmt.Println(y.inner.Set("krish",34))
	//fmt.Println(y.Get("krish"))
	//fmt.Println(x.m);//this one you cannot print
	fmt.Println(x);//this one you can print
	var z InMemoryCache = InMemoryCache{m: make(map[string]int)}
	z.Set("krish",2)
	fmt.Println(z.Get("krish"))
	fmt.Println(z.m);//here it is accessible 

}
//interface in go expose only methods not expose any fields
//