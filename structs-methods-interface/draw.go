package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"
)

/*
Implement a painting program
It should support :
- Circle with location (x,y) . Color and radius
- Rectangle with location (x,y) , width ,height and color
Each type should implement a 'Draw(d Device)' method

*/
type (
	Shape struct {
		X , Y int
		Color color.Color
	}
	Circle struct {
		Shape
		Radius int
	}
	Rectangle struct {
		Shape
		Height , Width int
	}
	ImageCanavas struct {
		w , h int
		shapes []Drawer
	}
	Drawer interface {
		Draw(d Device)
	}
	Device interface {
		Set(int,int,color.Color)
	}
)
var (
	RED = color.RGBA{
		R: 0xFF,
		G: 0,
		B: 0,
		A: 0xFF,
	}
	GREEN = color.RGBA{
		R: 0,
		G: 0xFF,
		B: 0,
		A: 0xFF,
	}
	BLUE = color.RGBA{
		R: 0,
		G: 0,
		B: 0xFF,
		A: 0xFF,
	}
)

func NewCircle(x,y,r int , c color.Color) *Circle  {
	cr := Circle{
		Shape: Shape{
			X:     x,
			Y:     y,
			Color: c,
		},
		Radius: r,
	}
	return &cr
}
func(c *Circle) Draw (d Device)  {
	minX , minY := c.X-c.Radius , c.Y-c.Radius
	maxX , maxY := c.X+c.Radius , c.Y+c.Radius
	for x := minX ; x <= maxX ; x++ {
		for y := minY ; y <= maxY ; y++ {
			dx ,dy := x-c.X , y-c.Y
			if int(math.Sqrt(float64(dx*dx+dy*dy))) <= c.Radius {
				d.Set(x,y,c.Color)
			}
		}
	}
}
func NewRectangle(x,y,h,w int , c color.Color) *Rectangle  {
	r := Rectangle{
		Shape:  Shape{x,y,c},
		Height: h,
		Width:  w,
	}
	return &r
}
func(r *Rectangle) Draw (d Device)  {
	minX , minY := r.X-r.Width/2 , r.Y-r.Height/2
	maxX , maxY := r.X+r.Width/2 , r.Y+r.Height/2
	for x := minX ; x <= maxX ; x++ {
		for y := minY ; y <= maxY ; y++ {
			d.Set(x,y,r.Color)
		}
	}
}
func NewImageCanvas (width , height int) (*ImageCanavas , error)  {
	if width <= 0 || height <= 0 {
		return nil , fmt.Errorf("negative size : width=%d , height=%d",width,height)
	}
	c := ImageCanavas{
		w:      width,
		h:      height,
	}
	return &c , nil
}
func(ic *ImageCanavas) Draw (w io.Writer) error  {
	img := image.NewRGBA(image.Rect(0,0,ic.w,ic.h))
	for _ , s := range ic.shapes {
		s.Draw(img)
	}
	return png.Encode(w,img)
}
func(ic *ImageCanavas) Add (d Drawer)   {
	ic.shapes = append(ic.shapes , d)
}
func main()  {
	ic , err := NewImageCanvas(200,200)
	if err != nil {
		log.Fatal(err)
	}
	ic.Add(NewCircle(100,100,80,GREEN))
	ic.Add(NewCircle(60,60,10,BLUE))
	ic.Add(NewCircle(140,60,10,BLUE))
	ic.Add(NewRectangle(100,130,10,80,RED))
	f,err := os.Create("face.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := ic.Draw(f);err != nil{
		log.Fatal(err)
	}
}