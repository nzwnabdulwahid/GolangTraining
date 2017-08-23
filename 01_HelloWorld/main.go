package main
import "fmt"


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

}