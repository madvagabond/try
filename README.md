# Try

## Clean, composable, and idiomatic error handling in go. 


### Examples 

```go
import ( '

    "github.com/xnukernpoll/try"
    "io"
    "net"
    "fmt"
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
    
    a.OnFail(func(e error) {fmt.Printf("Sorry I Screwed up and got %v", e)} )
} 

```




### By comparison this is how it is in plain go
``` 



```
