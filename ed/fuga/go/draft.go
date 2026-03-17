package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)
    var c string
    if d == 1 {
        for{
            if f == h {
                c = "S"
                break
            } 
            if f == p {
                c = "N"
                break
            }  
            f++
            if f == 16 {
                f = 0
            } 
        }
    } else {
        for {
            if f == h {
                c = "S"
                break
            } 
            if f == p {
                c = "N"
                break
            } 
            f--
            if f == -1 {
                f = 15
            } 
        }
    }
    fmt.Println(c)
}
