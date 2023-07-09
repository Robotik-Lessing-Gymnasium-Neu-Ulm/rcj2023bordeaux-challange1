package vision

import (
	"fmt"

	"gocv.io/x/gocv"
)

func CordCalculator(corners [][]gocv.Point2f, ids []int) {
	if len(ids) > 0 {
		for i := 0; i < len(ids); i++ {
			for j := 0; j < len(ids); j++ {
				fmt.Printf("ID %d:X-Cord: %f; Y-Cord: %f\n",ids[i], corners[i][j].X, corners[i][i].Y)
			}
		}
	}
}
