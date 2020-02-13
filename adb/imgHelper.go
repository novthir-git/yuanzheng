package adb

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

//# 获取灰度值
//def get_gray_level(self, img_uri, point):
//img = cv2.imread(img_uri)
//gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
//log.d('点（%s,%s）灰度值为：%s' % (point.x, point.y, gray[point.y][point.x]))
//return gray[point.y][point.x]

func GetPoint(matImageFile string, matTemplateFile string) *image.Point {

	matImage := gocv.IMRead(matImageFile, gocv.IMReadGrayScale)
	// gocv.Canny(matImage, &matImage, 200, 400)
	matTemplate := gocv.IMRead(matTemplateFile, gocv.IMReadGrayScale)

	//获取模版图片 宽/高
	tplCols := matTemplate.Cols()
	tplRow := matTemplate.Rows()

	// gocv.Canny(matTemplate, &matTemplate, 20, 40)
	matResult := gocv.NewMat()
	mask := gocv.NewMat()
	gocv.MatchTemplate(matImage, matTemplate, &matResult, gocv.TmCcoeffNormed, mask)
	mask.Close()
	minConfidence, maxConfidence, minLoc, maxLoc := gocv.MinMaxLoc(matResult)

	fmt.Println(minConfidence, maxConfidence, minLoc, maxLoc)

	gocv.Normalize(matImage, &matTemplate, 0, 0, 32)

	//gocv.Rectangle(&matImage, rect, color.RGBA{0, 0, 255, 0}, 1)
	point := image.Point{
		X: maxLoc.X + (tplCols / 2),
		Y: maxLoc.Y + (tplRow / 2),
	}
	gocv.Circle(&matImage, point, 5, color.RGBA{0, 0, 255,0}, 10)
	gocv.IMWrite("out.png", matImage)

	if maxConfidence > 0.5 {
		return &point
	}

	return nil
}
