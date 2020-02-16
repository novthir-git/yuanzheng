package task

import (
	"log"
	"time"
	"yuanzheng/adb"
)

func GetPointAmend(templateImg string, deviceName string, count int) (adb.Point, error) {
	img := adb.ShellScreenCapPullRm(deviceName)
	point, err := adb.GetPoint(img, templateImg)
	if err == nil {
		return adb.Point{
			X: point.X + adb.RiChangE,
			Y: point.Y,
		}, err
	} else if count < 3 {
		//翻下页查找
		adb.Swipe(adb.RiChangPageStart, adb.RiChangPageEnd, deviceName)
		count = count + 1
		return GetPointAmend(templateImg, deviceName, count)
	}
	Back(deviceName)
	return point, err
}

// 循环等待获取Point
func EachGetPoint(deviceName string, templateImg string, count int, sleep int) (adb.Point, error) {

	//step2
	var i = 1
	for {
		log.Printf("[%s] 第[%d]查找", deviceName, i)
		time.Sleep(time.Duration(sleep) * time.Second)
		img := adb.ShellScreenCapPullRm(deviceName)
		point, err := adb.GetPoint(img, templateImg)
		if err == nil || i >= count {
			return point, err
		}
		i++
	}
}

// 寻龙诀
func XunLongJue(deviceName string) {
	log.Printf("[%s] 开始寻龙诀任务\n", deviceName)
	templateImg := "images/template/richang/xunlongjue.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Println("XunLongJue not fount point")
		Back(deviceName) //返回主页
		return
	}
	adb.Click(point, deviceName)
	//step2 一键扫荡
	log.Printf("[%s] 寻龙诀一键扫荡", deviceName)
	adb.Click(adb.XunLongJueExecBtn, deviceName)

	time.Sleep(5 * time.Second)
	//step3 确认完成
	log.Printf("[%s] 寻龙诀确认完成", deviceName)
	SureClick(deviceName)

	//step4 Back
	log.Printf("[%s] 寻龙诀返回", deviceName)
	adb.Click(adb.XunLongJueBackBtn, deviceName)
}

//国王参拜
func King(deviceName string) {
	log.Printf("[%s] 开始国王参拜任务", deviceName)
	//step1 定位国王参拜入口
	templateImg := "images/template/richang/guowang.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Println("King not fount point")
		return
	}
	adb.Click(point, deviceName)

	//step2
	log.Printf("[%s] 执行参拜 %d", deviceName)
	templateImg = "images/template/richang/canbaiBtn.png"
	EachGetPoint(deviceName, templateImg, 10, 5)
	if err == nil {
		adb.Click(point, deviceName)
	}
}

//丝绸之路
func SilkRoad(deviceName string) {
	log.Printf("[%s] Silk Road start", deviceName)
	//step1 定位丝绸之路入口
	templateImg := "images/template/richang/silkRoad.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Println("Silk Road fount point")
		return
	}
	adb.Click(point, deviceName)
}

//官府任务
func GuanFu(deviceName string) {
	log.Printf("[%s] 官府任务 start", deviceName)
	//step1 官府任务
	templateImg := "images/template/richang/guanFu.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Println("官府任务fount point")
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(60 * 5 * time.Second) //等待5分钟做任务
	Unlock(deviceName)
	//判断第一轮是否完成
	var i = 1
	for {
		log.Printf("[%s] 官方任务判断第一轮是否完成 %d", deviceName, i)
		time.Sleep(10 * time.Second)
		img := adb.ShellScreenCapPullRm(deviceName)
		templateImg := "images/template/richang/guanFuOver.png"
		_, err := adb.GetPoint(img, templateImg)
		if err == nil {
			SureClick(deviceName)
			break
		} else if i > 10 || err != nil {
			break
		}
		i++
	}
	//第二轮官方任务
	RiChangHome(deviceName)
	log.Println("官府第二轮任务start")

	templateImg = "images/template/richang/guanFu1.png"
	point, err = GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Println("官府第二轮任务 not fount point")
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(60 * 6 * time.Second) //等待5分钟做任务
	Unlock(deviceName)
	log.Printf("[%s] ")
}

//一条龙
func YiTiaoLong(deviceName string) {
	log.Printf("[%s] 一条龙 start", deviceName)
	//step1 一条龙
	templateImg := "images/template/richang/yiTiaoLong.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Printf("[%s] 一条龙fail", deviceName)

		return
	}
	adb.Click(point, deviceName)
	time.Sleep(10 * time.Second)

	//step2 接受任务
	templateImg = "images/template/richang/yiTiaoLongAccept.png"
	point, err = EachGetPoint(deviceName, templateImg, 10, 6)
	if err != nil {
		log.Printf("[%s] 一条龙无法接受任务", deviceName)
		return
	}
	adb.Click(point, deviceName)

	//step3 确认接受
	time.Sleep(3 * time.Second)
	SureClick(deviceName)

	time.Sleep(16 * 60 * time.Second) //做任务中

}

//帮会环任务
func BangHuiHuan(deviceName string) {
	log.Printf("[%s] 帮会环 start", deviceName)
	//step1 一条龙
	templateImg := "images/template/richang/banghuiHuan.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Printf("[%s] 帮会环fail", deviceName)

		return
	}
	adb.Click(point, deviceName)
	time.Sleep(60 * 3 * time.Second)
}

// 种树
func Tree(deviceName string) {
	log.Printf("[%s] 种树 start", deviceName)
	//step1 种树
	templateImg := "images/template/richang/zhongShu.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Printf("[%s] 种树fail", deviceName)
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(30 * time.Second)
	//step2 接受任务
	templateImg = "images/template/richang/zhiShuBtn.png"
	point, err = EachGetPoint(deviceName, templateImg, 10, 6)
	if err != nil {
		log.Printf("[%s] 种树无法接受任务", deviceName)
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(30 * time.Second)
	adb.ShellScreenCapPullRm(deviceName)

}

//帮贡任务
func BangGong(deviceName string) {
	log.Printf("[%s] 帮贡任务 start", deviceName)
	//step1 帮贡任务
	templateImg := "images/template/richang/bangGong.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Printf("[%s] 帮贡任务fail", deviceName)
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(5 * time.Second)

	//step2 接受任务
	for i := 0; i < 6; i++ {
		templateImg = "images/template/richang/bangGong1000.png"
		point, err = EachGetPoint(deviceName, templateImg, 1, 2)
		if err != nil {
			log.Printf("[%s] 帮贡10000 失败", deviceName)
			return
		}
		log.Printf("[%s] 第%d次帮贡10000 ", deviceName, i)
		adb.Click(point, deviceName)
		time.Sleep(5 * time.Second)
	}
	Back(deviceName)

}

//封神榜
func FengShenBang(deviceName string) {
	log.Printf("[%s] 封神榜 start", deviceName)
	//step1 帮贡任务
	templateImg := "images/template/richang/fengShenBang.png"
	point, err := GetPointAmend(templateImg, deviceName, 0)
	if err != nil {
		log.Printf("[%s] 封神榜fail", deviceName)
		return
	}
	adb.Click(point, deviceName)
	time.Sleep(2 * time.Second)
	// 开始挑战
	if err == nil {
		for i := 1; i < 14; i++ {
			templateImg = "images/template/richang/tiaoZhan.png"
			img := adb.ShellScreenCapPullRm(deviceName)
			point, err = adb.GetPoint(img, templateImg)
			if err == nil {
				adb.Click(point, deviceName)
				time.Sleep(1 * time.Second)
			}
		}
	}
	Cancel(deviceName)
	Back(deviceName)
}

func RiChangHome(deviceName string) {
	adb.Click(adb.RiChangHome, deviceName)
}

//日常任務
func RiChangMain(deviceName string) {
	Unlock(deviceName)
	//step1 打开日常
	//adb.ShellScreenCapPullRm(deviceName)
	//point := adb.GetPoint(fmt.Sprintf("images/screen/%s", "emulator-5556-1581598332.png"), "images/template/richang_home.png")
	//point, err := adb.GetPoint(img, "images/template/richang/richang_home.png")
	//if err == nil {
	//	adb.Click(*point, deviceName)
	//} else {
	//	fmt.Println(err)
	//}

	RiChangHome(deviceName)
	//寻龙诀
	XunLongJue(deviceName)
	//

	//国王参拜
	RiChangHome(deviceName)
	King(deviceName)

	//丝绸之路
	RiChangHome(deviceName)
	SilkRoad(deviceName)

	//官府任务
	RiChangHome(deviceName)
	GuanFu(deviceName)

	//一条龙任务
	RiChangHome(deviceName)
	YiTiaoLong(deviceName)

	//帮会环任务
	RiChangHome(deviceName)
	BangHuiHuan(deviceName)

	//种树
	RiChangHome(deviceName)
	Tree(deviceName)

	//帮贡任务
	RiChangHome(deviceName)
	BangGong(deviceName)

	//封神榜
	RiChangHome(deviceName)
	FengShenBang(deviceName)
}
