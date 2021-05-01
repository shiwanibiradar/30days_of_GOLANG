package main
import "fmt"
func add(a int, b int) int {
	var c = a+b
	return c
}
func main() {
	var a,b,c int 
	fmt.Println("ENTER FIRST NUMBERS")
	fmt.Scanf("%d",&a)
	fmt.Println("ENTER SECOND NUMBERS")
	fmt.Scanf("%d",&b)
	c = add(a,b)
	fmt.Printf("SUM IS %d\n",c)
}
