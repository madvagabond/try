package try 
import (
	"testing"
	"errors"
)



func TestMonad(t *testing.T) {
	res := OK()
	failed := Fail(errors.New("shit"))

	err_res := res.Then(func() Try {return failed })
	if err_res.Success() == true {
		t.Fatalf("Bind test failed")
	} 


	res1 := res.Map(func() error {return nil } ) 
	
	if res1.Success() != true {
		t.Fatalf("Map test failed ") 
	}
}
