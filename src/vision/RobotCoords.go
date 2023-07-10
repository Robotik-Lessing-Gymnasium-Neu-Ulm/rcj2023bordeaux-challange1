package vision

import (
	"fmt"

	"gocv.io/x/gocv"
)

func CordCalculator(corners [][]gocv.Point2f, ids []int) {
	if len(ids) == 5 {
		//referenz zum Mittelpunkt
		dy := corners[2][0].X - 434
		dx := corners[2][0].Y - 332

		var dx_incm float32 = (0.3 * dx)
		var dy_incm float32 = (0.3 * dy)

		fmt.Printf("X: %f -- Y: %f\n", dx_incm, dy_incm)

		/*neutralSpotTopRight = gocv.NewPoint2f(425, 634)
		neutralSpotTopLeft = gocv.NewPoint2f(178, 628)
		neutralSpotDownRight = gocv.NewPoint2f(433, 225)
		neutralSpotDownLeft = gocv.NewPoint2f(206, 223)
		*/

		//anglToNeutralPoint := math.Atan2(float64(-dy), float64(-dx)) * 180 / math.Pi

		//fmt.Printf("Angle: %f\n", anglToNeutralPoint) vc 

	} else {
		fmt.Println("Missin Points")
	}
}
