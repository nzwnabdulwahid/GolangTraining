package main
import "fmt"
//import "./intro"


const name string = "Niezwan"

const (
	fullname = "Niezwan Abdul Wahid"
	realAge = 24
)

func main() {
	fmt.Println("Hello World")

	//Variables
	var a int = 10
	var b string = "Ayam"
	var c float32 = 10.1
	d := 20.3
	var e, f string = "this is into e", "this is into f"

	//Also equal to :
	//a:= 10
	//b:= "Ayam"
	//c:= 10.1

	fmt.Printf("%v, %v and %v and also %v \n",a,b,c, d)
	fmt.Printf("Not to forget %v and %v \n", e, f)

	// %v is default value
	// %T is type

	//Constant
	const age = 23
	fmt.Printf("%v is a constant and also %v \n", name, age)

	//Print from multiple declaration
	fmt.Printf("%v is a constant and also %v \n", fullname, realAge)
	//Explain_IOTA()



}