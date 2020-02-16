package task

import (
	"log"
	"yuanzheng/adb"
)

// 物品奖励确认框
func JiangLiConfirm(deviceName string) {

	log.Printf("[%s] 物品奖励确认框\n", deviceName)
	templateImg := "images/template/back.png"
	img := adb.ShellScreenCapPullRm(deviceName)
	_, err := adb.GetPoint(img, templateImg)
	if err != nil {
		log.Println("JiangLiConfirm not fount point")
		return
	} else {
		//点击确认
		SureClick(deviceName)
	}

}

func SureClick(deviceName string) {
	templateImg := "images/template/sureBtn.png"
	img := adb.ShellScreenCapPullRm(deviceName)
	point, err := adb.GetPoint(img, templateImg)
	if err != nil {
		templateImg = "images/template/sureBtn1.png"
		img = adb.ShellScreenCapPullRm(deviceName)
		point, err = adb.GetPoint(img, templateImg)
		if err != nil {
			templateImg = "images/template/sureBtn2.png"
			img = adb.ShellScreenCapPullRm(deviceName)
			point, err = adb.GetPoint(img, templateImg)
			if err != nil {
				log.Printf("[%s]  未找到完成确认按钮", deviceName)
			}
		}
	}
	adb.Click(point, deviceName)
}

func Back(deviceName string) {
	log.Printf("[%s] 返回\n", deviceName)
	templateImg := "images/template/back.png"
	img := adb.ShellScreenCapPullRm(deviceName)
	point, err := adb.GetPoint(img, templateImg)
	if err != nil {
		log.Println("Back not fount point")
		return
	}
	adb.Click(point, deviceName)
}

func Cancel(deviceName string)  {
	templateImg := "images/template/cancel.png"
	img := adb.ShellScreenCapPullRm(deviceName)
	point, err := adb.GetPoint(img, templateImg)
	if err != nil {
		log.Println("cancel not fount point")
		return
	}
	adb.Click(point, deviceName)
}

//远征滑动解锁
func Unlock(deviceName string) {
	adb.Swipe(adb.LockStart, adb.LockEnd, deviceName)
}
