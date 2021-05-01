package main
import "fmt"
func max_two(a int, b int) int {
        if  a > b {
                return a
        return b
        }
func max_three(a int, b int, c int){
        return max_two(a, max_two(b,c))
        }
func main() {
        fmt.Println("max number will be",max_three(3,6,4))
}

