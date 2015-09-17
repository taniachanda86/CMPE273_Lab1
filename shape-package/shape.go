package newShape
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

// type Triangle struct{
// 	l1, l2, l3 float64
// }

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

func GetArea(shape Shape) float64{
	return shape.area()	
}

func GetPerimeter(shape Shape) float64{
	return shape.perimeter()	
}