package vision

import (
	"fmt"

	"gocv.io/x/gocv"
)

func CordCalculator(corners [][]gocv.Point2f, ids []int) {
	if len(ids) < 5 {
		fmt.Println("Robot not founds")
	} else {
		markercorners := corners[4]
		fmt.Println(markercorners[0].X, " :: ", markercorners[0].Y)
	}

}
