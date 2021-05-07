package main
import (
   "fmt"
   "math"
)

func isPrime() func(num int) bool{
    f := func(num int) bool {
        for i:=2; i<=int(math.Sqrt(float64(num))); i++{
            if num % i == 0{
                return false
            }
        }

        return true
    }

    return f
}

func createPrimes(low, high int, f func(n int) bool, res []int) []int{
    if(low > high) {
        return res
    }

    if(f(low)) {
        return createPrimes(low + 1, high, f, append(res, low))
    } else {
        return createPrimes(low + 1, high, f, res)
    }
}

func findPrimeNumbers(low, high int) []int{
    if low<2 || high<2{
       fmt.Println("Numbers must be greater than 2.")
       return []int{}
    }

    return createPrimes(low, high, isPrime(), []int{})
}

func printRes(primes []int) {
    fmt.Println(primes)
}

func main(){
    printRes(findPrimeNumbers(5, 19))
    printRes(findPrimeNumbers(0, 2))
    printRes(findPrimeNumbers(13, 100))
}

