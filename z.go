package main
import "fmt"

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    fmt.Println("Original Rectangle:", rect)
    fmt.Println("Area:", rect.Area())

    rect.Scale(2)
    fmt.Println("Scaled Rectangle:", rect)
    fmt.Println("New Area:", rect.Area())
}