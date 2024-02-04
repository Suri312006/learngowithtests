package structsmethodsinterfaces

import "math"


// the interface is applied to anything with the Area method that returns 
// a float64
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

func (r Rectangle) Perimeter() float64 {
    return 2*(r.Height+r.Width)
}

// this is the reviever  w the receiver type
func (r Rectangle) Area() float64 {
    return r.Height*r.Width
}


type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}


type Triangle struct {
    Base float64
    Height float64
}

func (t Triangle) Area() float64 {
    return 0.5*(t.Height*t.Base)
}

