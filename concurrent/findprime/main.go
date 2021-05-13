package main
import (
   "fmt"
   "github.com/jieliuhk/cs345-lang-pesent/oo/mytype"
   "sync"
)

func printPrime(num int, wg *sync.WaitGroup){
   defer wg.Done()

   myint := mytype.MyInt{
     Value: num,
   }

   isPrime := myint.IsPrime()

   if isPrime {
      fmt.Printf("%d ", num)
   }
}

func printPrimeNumbers(num1, num2 int){
   if num1<2 || num2<2{
      fmt.Println("Numbers must be greater than 2.")
      return
   }

   var wg sync.WaitGroup

   for num1 <= num2 {
      wg.Add(1)
      printPrime(num1, &wg)
      num1++
   }

   wg.Wait()

   fmt.Println()
}

func main(){
   printPrimeNumbers(5, 19)
   printPrimeNumbers(0, 2)
   printPrimeNumbers(13, 100)
}
