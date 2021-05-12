package mytype

import (
    "math"
)

type MyInt struct {
    Value int
}


func (myint *MyInt) IsPrime() bool {

    for i:=2; i<=int(math.Sqrt(float64(myint.Value))); i++ {
         if myint.Value % i == 0{
            return false
         }
    }

    return true
}



