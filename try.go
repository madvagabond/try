package try 

type Try struct {
	bool
	error
}


 


func OK() Try {
	return Try{true, nil}
}



func Fail(e error) Try {
	return Try{false, e}
}

// A wrapper around errors that checks the error 
// if error != nil returns failure if it's nil returns ok
func Check(e error) Try {

	if e != nil {
		return Try{false, e}
	}

	return OK() 
}




func Run(f func() error) Try {
	e := f()
	return Check(e) 

}

func (t *Try) Success() bool {
	return t.bool 
} 


func (t *Try) Error() error {
	return t.error 
}

func (t *Try) Map(f func() error) Try {
	if !t.Success() {
		return Fail(t.Error())
	} 

	return Run(f) 

}


func (t *Try) Then(f func() Try) Try {
	if !t.Success() {
		return Fail(t.Error())
	} 

	return f() 
} 


func (t *Try) OnFail(f func(e error)) {
	if !t.Success() {
		f(t.Error()) 
	}
}

// takes the error and trys to repair and returns true if it does
type RescueFunc = func(error) bool




func (t *Try) Rescue(rf RescueFunc) Try {
	if !t.Success() {
		if rf(t.Error()) {
			return OK()
		}
		return Fail(t.Error()) 
	}

	return OK() 
	
}

func (t *Try) OnSuccess(f func() ) {
	if t.Success() {
		f() 
	}
}
