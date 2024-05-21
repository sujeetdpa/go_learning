package main

import "fmt"

type Car struct {
	Name       string
	Model      string
	SNO        int
	FrontWheel Wheel //Nested structs
	RearWheel  Wheel
}
type Wheel struct {
	Radius           float32
	IsChromeFinished bool
}

// Embeded Structs
type shape struct {
	Dimension1 float32
	GetArea    func() float32
}
type circle struct {
	shape //Embeded
}

//Methods in structs

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) getArea() float32 {
	return r.Breadth * r.Length
}

func main() {

	var caar Car
	caar.Name = "Ford"
	car := Car{
		Name:  "Porche",
		Model: "Canmera",
		SNO:   45874984,
	}
	car.FrontWheel.IsChromeFinished = true
	car.RearWheel.IsChromeFinished = false
	fmt.Println(car)

	//Anonymous Structs
	acar := struct {
		Name  string
		Model string
		SNO   int
	}{
		Name:  "Hello",
		Model: "Hello Model",
		SNO:   484376454,
	}
	fmt.Println("Anonymous Structs", acar)

	//Embeded Structs
	circ := circle{}
	circ.Dimension1 = 6.56
	circ.GetArea = func() float32 {
		return 3.14 * circ.Dimension1 * circ.Dimension1
	}
	fmt.Println("Circle", circ.GetArea())

	//Methods in structs

	rec := Rectangle{
		Breadth: 2,
		Length:  2,
	}
	fmt.Printf("Rectangle Dimensions Are: %v and its area is: %v\n", rec, rec.getArea())

}
