package main
import (
   "fmt"
   "github.com/jieliuhk/cs345-lang-pesent/oo/mytype"
)

func printPrimeNumbers(num1, num2 int){
   if num1<2 || num2<2{
      fmt.Println("Numbers must be greater than 2.")
      return
   }

   for num1 <= num2 {
       myint := mytype.MyInt{
        Value: num1,
       }
      isPrime := myint.IsPrime()

      if isPrime {
         fmt.Printf("%d ", num1)
      }

      num1++
   }
   fmt.Println()
}

func main(){
   printPrimeNumbers(5, 19)
   printPrimeNumbers(0, 2)
   printPrimeNumbers(13, 100)
}
