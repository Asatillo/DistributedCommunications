package main
import(
	"fmt"
	"time"
)

func main(){
	c := make(chan int)
	go func(){
		// TIME OUT
		// time.Sleep(4*time.Second)
		c <- 1
		fmt.Println("g1 sent 1")
	}()

	select{
	case msg := <-c:
		fmt.Println("Main received ", msg)
	case <- time.After(3*time.Second):
		fmt.Println("Timeout, Quit")
	}
}