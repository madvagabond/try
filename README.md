# Try

## Idiomatic monadic-ish error handling for go. 



### Examples 

```go
import ( '
	"github.com/xnukernpoll/try"
    "io"
    "net"
)

	func simpleExample() error {

		conn, e := net.Dial("tcp", "0.0.0.0:21") 
    	// wraps over error passes if nil, fails if non-empty  
    	t := try.Check(e)
    	
    
		t1 := t.Map(func() error {
			_, e := conn.Write(b"")
    		return e 
		})
    
    
    
    return t1.Error() 

}





func onExample() {
	e := simpleExample()
    a := try.Check(e)
    
    a.OnSuccess(func(){
    	fmt.Println("hello I've grinded to succeed but now I am all alone and empty inside")
    } )
    
    a.OnFail(func() {fmt.Println("Sorry I Screwed up")} )
} 

```





