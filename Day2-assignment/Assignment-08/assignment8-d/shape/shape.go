package shape

import (
	"assignment8-d/model"
	"math"
)

func AreaofACircle(c model.Circle) float64 { //defining AreaofACircle func accepts circle struct
	return math.Pi * c.Radius * c.Radius
}
