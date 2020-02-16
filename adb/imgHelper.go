package adb

import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"log"
	"time"
)

//# 获取灰度值
//def get_gray_level(self, img_uri, point):
//img = cv2.imread(img_uri)
//gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
//log.d('点（%s,%s）灰度值为：%s' % (point.x, point.y, gray[point.y][point.x]))
//return gray[point.y][point.x]

func GetPoint(img image.Image, matTemplateFile string) (Point, error) {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	matImage, _ := gocv.ImageGrayToMatGray(grayImg)
	defer matImage.Close()

	matTemplate := gocv.IMRead(matTemplateFile, gocv.IMReadGrayScale)
	if matTemplate.Empty() {
		fmt.Printf("Invalid read of %s ", matTemplateFile)
	}
	defer matTemplate.Close()
	//获取模版图片 宽/高
	tplCols := matTemplate.Cols()
	tplRow := matTemplate.Rows()

	matResult := gocv.NewMat()
	defer matResult.Close()
	mask := gocv.NewMat()
	gocv.MatchTemplate(matImage, matTemplate, &matResult, gocv.TmCcoeffNormed, mask)
	mask.Close()
	minConfidence, maxConfidence, minLoc, maxLoc := gocv.MinMaxLoc(matResult)
	log.Println(minConfidence, maxConfidence, minLoc, maxLoc)
	p := image.Point{
		X: maxLoc.X + (tplCols / 2) ,
		Y: maxLoc.Y + (tplRow / 2),
	}

	gocv.Circle(&matImage, p, 5, color.RGBA{0, 0, 255, 0}, 20)
	gocv.IMWrite(fmt.Sprintf("out-%d.png", time.Now().Unix()), matImage)

	if maxConfidence > 0.99 {
		return Point{
			X: p.X ,
			Y: p.Y,
		}, nil
	}
	return Point{}, errors.New("not found point")
}

func TestMatchTemplate(matImageFile string, matTemplateFile string) {
	imgScene := gocv.IMRead(matImageFile, gocv.IMReadGrayScale)
	if imgScene.Empty() {
		fmt.Println("Invalid read of face.jpg in MatchTemplate test")
	}
	defer imgScene.Close()

	imgTemplate := gocv.IMRead(matTemplateFile, gocv.IMReadGrayScale)
	if imgTemplate.Empty() {
		fmt.Println("Invalid read of toy.jpg in MatchTemplate test")
	}
	defer imgTemplate.Close()

	result := gocv.NewMat()
	defer result.Close()
	m := gocv.NewMat()
	gocv.MatchTemplate(imgScene, imgTemplate, &result, gocv.TmCcoeffNormed, m)
	m.Close()
	_, maxConfidence, _, _ := gocv.MinMaxLoc(result)
	if maxConfidence < 0.95 {
		fmt.Println("Max confidence of %f is too low. MatchTemplate could not find template in scene.", maxConfidence)
	}
}
