# Making member functions to Go Structure
### It is an antipattern, don't do it.

C++ and Java users tend use Go Structure like a class. A C++ way to add a method is to make a member function. 

This can be closely mimiced as follows.

## Antipatern !!

```
type twice_func_type func(input int) int
```

Define a structure and make a member variable to a go function as functions are first class objects.

```
type myNum struct {
	data  int
	twice twice_func_type
}
```

Make a concrete function which implements ```doTwice``` functionality.

```
func doTwice(input int) int {
	return input * 2
}
```
## Pattern. :) 
However this approach is convoluted and less than ideal. </span> 
Go doesn't have member functions to structure but a method can be bound to any type.

```
func (input myNum) idealTwice() int {
	return input.data * 2
}
```


A call to member function as a call to method gives a same result. However defining a method is a correct way to bind the method.

```
func main() {
	x := myNum{twice: doTwice, data: 100}
	fmt.Println(x.twice(x.data))
	fmt.Println(x.idealTwice())
}
```
Complete code can be found at https://github.com/forvaidya/go-learn-stuff/blob/main/antipatterns/test.go and can be tested using ```go run test.go```
