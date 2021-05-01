package main
import "fmt" 
func main() {
	var a,b,temp int
	fmt.Printf("Enter two numbers\n")
	fmt.Scanf("%d %d",&a,&b)
	fmt.Printf("Two Numbers are %d and %d\n",a,b)
	temp=0
	temp=a
	a=b
	b=temp
	fmt.Printf("After swapping, numbers are %d and %d",a,b)
}
