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
		t.Fatalf("Doesn't satisfy properties bich")
	} 

}
