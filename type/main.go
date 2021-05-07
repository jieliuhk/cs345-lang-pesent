package main
import (
   "fmt"
   "math"
)

func printPrimeNumbers(num1, num2 int){
   if num1<2 || num2<2{
      fmt.Println("Numbers must be greater than 2.")
      return
   }

   for num1 <= num2 {
      //implicit type
      //var isPrime bool
      //isPrime = true
      isPrime := true
      //mainifest type, the following code won't compile
      //isPrime = "true"

      //static type
      for i:=2; i<=int(math.Sqrt(float64(num1))); i++{
         if num1 % i == 0{
            isPrime = false
            break
         }
      }
      if isPrime {
         fmt.Printf("%d ", num1)
      }
      num1++
   }
   fmt.Println()
   //strong type, the following code won't compile
   //fmt.Println("complie error" + 23)
   //this will work
   fmt.Println("compiled: " + string(66))
}

func main(){
   printPrimeNumbers(5, 19)
   printPrimeNumbers(0, 2)
   printPrimeNumbers(13, 100)
}
