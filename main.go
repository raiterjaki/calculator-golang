/*This has been coded during the learning process of golang.
  This project has been made with less effort with help of math lib.
  Do not attempt to judge me for publishing something simple like this on github. I know that it is simple. (I will just block you)
  Console input is missing because there are some of the functions that do not require it. (If you want it you can do it by yourself it isn't something hard)
  I highly estimate that I will update this calculator and add more futures to this anytime in the future.

  CHANGELOG [5th of March, 2021]:
    - Added rectangle struct
    - Added functions for rectangle such as area finder and perimeter finder

  [5th of March 2021 (2nd addition today)]
    - Added isosceles triangle struct
    - Added functions for isosceles triangle such as area finder and perimeter finder
    
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
type rectangle struct{ height, length int }
func findRecPerimeter(r rectangle) int { return r.height + r.height + r.length + r.length }
func findRecArea(r rectangle) int { return r.height * r.length }
type isosceles_triangle struct { A_side, B_side, C_side, height int }
func findIsTrPerimeter(t isosceles_triangle) int {
  if t.A_side != t.B_side {
    fmt.Println("[error] your variables are wrong. there are 2 equal sides in isosceles triangle. if you have 2 equal numbers but mentioned them as c or b, you can modify this code or just rearrange them. Even thought it is wrong, we still solved it for you.")
  } else if t.C_side < t.A_side | t.B_side {
    fmt.Println("[error] base of an isosceles triangle can't be smaller than the a and b. But we will still solve it for you.")
  }
  return t.A_side + t.B_side + t.C_side
}
func findIsTrArea(t isosceles_triangle) int {
  return (t.C_side * t.height) / 2 // C_side is the base.
}

//math funcs (end)a

func consoleOutput(a1 int) { fmt.Println("Your answer is: ", a1) } 

func main() {
    consoleOutput(addNum(5, 15)) //adding numbers
    consoleOutput(int(sinNum(50))) // you can do this to find sine but you will get rounded version. to get the exact number, use the way that I have added below.
    fmt.Println("Your answer is ", sinNum(50)) // this will provide you your sine result with exact number.
    fmt.Println("Perimeter of rectangle with height: 5 length: 7 is ", findRecPerimeter(rectangle{5, 7}))
    fmt.Println("Area of rectangle with height: 5 length: 7 is ", findRecArea(rectangle{5, 7}))
    fmt.Println("Perimeter of triange with a: 3, b: 3, c: 5 is ", findIsTrPerimeter(isosceles_triangle{3, 3, 5, 0})) // 0 = height
  fmt.Println("Area of a triangle with a: 5, b: 5, c (base): 7, height: 4 is ", findIsTrArea(isosceles_triangle{0, 0, 7, 4})) //base 7, height 4
}
