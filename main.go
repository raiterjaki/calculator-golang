/*This has been coded during the learning process of golang.
  This project has been made with less effort with help of math lib.
  Do not attempt to judge me for publishing something simple like this on github. I know that it is simple. (I will just block you)
  Console input is missing because there are some of the functions that do not require it. (If you want it you can do it by yourself it isn't something hard)
  I highly estimate that I will update this calculator and add more futures to this anytime in the future.
  github.com/raiterjaki
  raiterjaki#0796
  */

package main

import (
    "fmt"
    "math"
)

//math funcs (start)

func addNum(a1 int, a2 int) int { return a1 + a2 }
func subtractNum(a1 int, a2 int) int { return a1 - a2 }
func multiplyNum(a1 int, a2 int) int { return a1 * a2 }
func divideNum(a1 int, a2 int) int { return a1 / a2 }
func sinNum(a1 float64) float64 { return math.Sin(a1) }
func cosNum(a1 float64) float64 { return math.Cos(a1) }
func sqrtNum(a1 int) int { return a1 * a1 }
func cubeNum(a1 int) int { return a1 * a1 * a1 }
func powerNum(a1 int, power int) float64 { return math.Pow(3, 4) }
func tanNum(a1 float64) float64 { return math.Tan(a1) }

//math funcs (end)

func consoleOutput(a1 int) { fmt.Println("Your answer is: ", a1) } 

func main() {
    
    consoleOutput(addNum(5, 15)) //adding numbers
    consoleOutput(int(sinNum(50))) // you can do this to find sine but you will get rounded version. to get the exact number, use the way that I have added below.
    fmt.Println("Your answer is ", sinNum(50)) // this will provide you your sine result with exact number.
}