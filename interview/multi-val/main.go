package main
import (
	"fmt"
)

func reverseValues(a,b string)(string, string){
    //notice how multiple values are returned
    return b,a
}

func main(){
    // notice how multiple values are assigned
    val1,val2:= reverseValues("interview","bit")
    fmt.Println(val1, val2)
}
