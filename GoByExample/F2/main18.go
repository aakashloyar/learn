package main 
import "fmt"

type Cache interface {
	Set(key string, value int)
	Get(key string) (int, bool)
}

type InMemoryCache struct {
	m map[string]int
}

type ReadOnlyCache struct {
	m map[string]int
}
func (c *InMemoryCache) Set(key string, value int) {
	c.m[key]=value;
}
func (c ReadOnlyCache) Set(key string, value int) {
	panic("This is read only cache")
	//c.m[key]=value;
}
func (c ReadOnlyCache) Get(key string) (int, bool) {
	return c.m[key],true
}
func (c InMemoryCache) Get(key string) (int, bool) {
	return c.m[key],true
}
func main() {
	var x Cache= &InMemoryCache{m : make(map[string]int)}
	var y Cache= &ReadOnlyCache{m : make(map[string]int)}

	x.Set("Aakash",1);
	fmt.Println(y)
	fmt.Println(x);

}

//so here learn 1 thing 
//let us suppose one of the method of interface receive pointer then when declaring just use pointer
//otherwise it will not work