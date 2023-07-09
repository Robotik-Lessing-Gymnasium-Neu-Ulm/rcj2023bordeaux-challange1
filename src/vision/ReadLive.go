package vision

import (
	"log"

	"gocv.io/x/gocv"
)

func Init() {
	// webcam, err := gocv.VideoCaptureDevice(0)
	webcam, err := gocv.VideoCaptureFile("ch.mp4")
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error reading webcam: %s\n", err)
	}

	window := gocv.NewWindow("GPS Tracker")

	dictionary := gocv.GetPredefinedDictionary(gocv.ArucoDict4x4_100)

	parameters := gocv.NewArucoDetectorParameters()
	parameters.SetCornerRefinementMethod(1)

	detector := gocv.NewArucoDetectorWithParams(dictionary, parameters)

	for {
		frame := gocv.NewMat()
		if ok := webcam.Read(&frame); !ok {
			break
		}

		if frame.Empty() {
			continue
		}

		corners, ids, _ := detector.DetectMarkers(frame)
		if len(ids) > 0 {
			gocv.ArucoDrawDetectedMarkers(frame, corners, ids, gocv.NewScalar(20, 100, 200, 50))
		}

		CordCalculator(corners, ids)

		window.IMShow(frame)
		if window.WaitKey(1) >= 0 {
		    break
		}
	}
}
