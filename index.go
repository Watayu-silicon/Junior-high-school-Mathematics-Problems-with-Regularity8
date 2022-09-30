package main
import "fmt"
func main(){
    var num1 int
    var num2 int
    result := 0
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    if num1<num2 {
        a := num2
        b := num1
        dif := a-b
        mid := (a-1) * (a-1) + a;
        result := mid - dif
        fmt.Println(result)
    } else if num1>num2 {
        a := num1
        b := num2
        dif := a-b
        mid := (a-1) * (a-1) + a;
        result := mid + dif
        fmt.Println(result)
    } else if num1 == num2 {
        result = (num1-1) * (num1-1) + num2;
        fmt.Println(result)
    }
}