package main
import "fmt"
import "math"

type Shape interface{
	area() float64
	perimeter() float64
}

type Rectangle struct{
	height, width float64
}

type Circle struct{
	rad float64
}

func (r Rectangle) area() float64 {
    return r.height * r.width
}
func (r Rectangle) perimeter() float64 {
    return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
    return math.Pi * c.rad * c.rad
}
func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.rad
}

func getArea(shape Shape) float64{
	return shape.area()	
}

func getPerimeter(shape Shape) float64{
	return shape.perimeter()	
}

func main() {
    var rad float64
    var height float64
    var width float64

    fmt.Print("Please enter the radius for circle : ")
    fmt.Scanf("%f", &rad)
    fmt.Print("Please enter the height for rectangle : ")
    fmt.Scanf("%f", &height)
    fmt.Print("Please enter the width for rectangle : ")
    fmt.Scanf("%f", &width)

    c := Circle {rad}
	r := Rectangle {height, width}
    
    fmt.Println("Circle area : ", getArea(c))
    fmt.Println("Circle perimeter : ", getPerimeter(c))
    fmt.Println("Rectangle area : ", getArea(r))
    fmt.Println("Rectangle perimeter : ", getPerimeter(r))
}