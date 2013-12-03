package main
import (
  "fmt"
  "math"
  )


type Point struct{
  x int
  y int
}

func sum(x,y int) int{
  return x+y
}

func swap(x,y int) (int, int){
  return y,x
}

func main(){
  summe := sum(2,3)
  fmt.Printf("Sum: %v\n",summe)
  var x,y int
  x = 4
  y = 8

  x,y = swap(x,y)
  fmt.Printf("X: %v, Y: %v\n", x, y)
  fmt.Printf("Abs(-10): %v\n",math.Abs(-10))
  p2 := Point{4,5}
  var p Point = Point{2,3}
  fmt.Printf("Point: %v",p2)
  fmt.Printf("Point: %v",p)
}


