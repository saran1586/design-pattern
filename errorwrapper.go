package main

import (
"fmt"
"errors"
)

type ErrorWrapper struct{

msg string
inner error
}

func (e *ErrorWrapper) Error() string {
return fmt.Sprintf("%s: %v",e.msg,e.inner)
}

func WrapError(msg string , err error) *ErrorWrapper{
return &ErrorWrapper{
msg : msg,
inner:err,
}
}
func main(){
err:=errors.New("database connection failed")
wrappedErr := WrapError("unable to connect",err)
fmt.Println(wrappedErr)
}
