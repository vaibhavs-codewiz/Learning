1. hello world :what happens when run the go file : go run <file path> , what are versions of golang , what are different pacakges like fmt time
2. simple values : 
3. vairables : go infers the variable type , when use shorthand syntax set the value on the go 
var name string = "some"
var name = "some" // type infers
name := "some" // shorthand , cannot be use outside function
var name string //only declare
name = "some" // initialize

4. constants : vairble with const keyword and cannot be reinitialize , 
grouping const : const ( port = , host = )

5. for loop : only construct for looping in go 
 while can be implemented using for style ,for i<= 3 {}
 infinite no condition , for {}
 classic for loop , for i:=3; i < 3; i++ {}
 break and continue keywords are available 
range for loop , for i := range 3 {}
if else
we can declare variable inside if construct : if age := 12; age = 15{} 
go does not have ternaryy operator in version 1.22

6. simple switch in go 
multiple condition switch , (use time library package) 
switch time.NOW().weekday() {
case time.Saturday, time.Sunday:
	...
}

type switch : 
whoAmI := func(i interface{}){
	switch i.(type) {
	case int:	
	 ...
	case string:
		...
	}
}

whoAmI("golang")

7. arrays :
var nums [4]int //default 0 value
len(nums)
fmt.Println(nums) 

nums := [3]int{1,2,3}
2d array , nums := [2][2]int{{},{}}

8. slices , most used , 
dynamic array, 
var nums = make([]int, 2, 5); type size capacity 2 zero initially so keep size 0 
when cap got full it gets increase by 2x . append(nums,2)
cap(nums) len(nums)
creating array without size is also a slice 
can use Slice package and use method of it like Equal()

9. map 
create m := make(map[key type]value type) , m1 := map[string]int{"": 3}
setting value , m[key type] = value
access m[key type]  
delete, delete(m,key)
khali krne k liye , clear(m)
for comparing , maps.Equal(m1,m2)

10. range
use for iterating data structure
nums can be slice or array
for idx,val := range nums {
	...
}
can also be used for map to access key, value

can be use for string 
range prints starting byte of rune (character) and unicode in case of string 
for i, c := range "golang" {
	fmt.Println(i,string(c)) , just c print unicode
}

11. functions 
func add(a int,b int) int {
	return a+b
}
returning multiple values
func sub() (int, string) {
	return 1,"sub"
}

can return the function from one function : 
func printfn() func fn(int) int
{
	return fn(a int) int{
		return a
	}
}
func main(){
	fn := printfn()
	fn(6)
}

12. variadic function
func sum(nums ...int) int { here nums of type int for any type use interface{}
  total := 0
  for val := range nums {
	total = total + val
  }
  return total
}
func main() {
	// sum(3,4,5,6)
	num := []int{3,4,5,6,7}
	result := sum(num...) // spread the slice
}

13. closure in go  
function can use variable which is not in scope 

14. pointers in go
instead of pass by value pass by reference
pass by ref : change(&num) //1
func change(num *int) {
  *num = 3 //3 dereferencing
}

15. structs
can be use to implement oop in golang 
struct handles the dereferencing in case of pointer recevier method of any struct
convention : instead of creating directly object we should create by using the constructor and
this constructor will initialize the value in this struct 
ex. func initialize(id int, name string) *order {
	myOrder := order{
		id : id,
		name : name
	}
	return &myOrder
}
now from main method for creating new order : 
 newOrder := initialize(1,go)

 inline structs 
 lang := struct{
  id int,
  name string
 }{1,"go"}

 16. struct embedding
 struct k andar struct

 17. interface
  

 18. enums

 19. generics

 20. Goroutines
 running tasks parallely 
 how go execute this internally : understand this 
 understand the uscase of goroutines

 21. wait groups in goroutines
 mechanism to synchronize goroutine
 use methods Add done wait of it 

 22. Channels
    create channel make(chan string)
    sending data in channel, 