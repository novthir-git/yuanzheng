/**
Adb 基本操作集合
*/
package adb

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	appPkgName = ""
	className  = ""
	adbCmd     = "D:/tools/adb/adb.exe"
	sleepTime  = 2 * time.Second
)

// adb链接模拟器
func Connect(ip string, port int) {
	host := fmt.Sprintf("%s:%d", ip, port)

	log.Printf("connect  host [%s]", host)

	output, err := exec.Command(adbCmd, "connect", host).Output()
	if err != nil {
		log.Printf("connect error %s", err)

	}

	log.Printf("connect result %s", output)

}
func Disconnect(ip string, port int) {
	host := fmt.Sprintf("%s:%d", ip, port)

	log.Printf("disconnect  host [%s]", host)

	output, err := exec.Command("adb", "disconnect", host).Output()
	if err != nil {
		log.Printf("disconnect error %s", err)

	}
	log.Printf("disconnect result %s", output)
}

//截屏并保存到当前目录下。
//由于需在手机和电脑上复制文件，必要时可增加延时或用下面的PathExists()判断文件是否存在，如：
//time.Sleep(time.Duration(2) * time.Second)
func ShellScreenCapPullRm(deviceName string) image.Image {
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%s-%d.png", deviceName, timestamp)

	cmd := exec.Command(adbCmd, "-s", deviceName, "shell", "screencap", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Printf("screencap error %s", err)
	}
	x := bytes.Replace(out.Bytes(), []byte("\r\r\n"), []byte("\n"), -1)

	img, _, _ := image.Decode(bytes.NewReader(x))
	f, _ := os.Create(fmt.Sprintf("./images/screen/%s", fileName))

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return img
}

//模拟屏幕点击
//adb shell input tap  900 800
func Click(point Point, deviceName string) {
	x2 := strconv.Itoa(point.X)
	y2 := strconv.Itoa(point.Y)
	output, err := exec.Command(adbCmd, "-s", deviceName, "shell", "input", "tap", x2, y2).Output()
	if err != nil {
		log.Printf("client error: %s", err)
	}
	log.Printf("[%s] click result %s", deviceName, string(output))
	time.Sleep(sleepTime)
}

//模拟滑动
//adb shell input swipe  0 0  600 600
func Swipe(start Point, end Point, deviceName string) {
	xx1 := strconv.Itoa(start.X)
	yy1 := strconv.Itoa(start.Y)
	xx2 := strconv.Itoa(end.X)
	yy2 := strconv.Itoa(end.Y)
	_, err := exec.Command(adbCmd, "-s", deviceName, "shell", "input", "swipe", xx1, yy1, xx2, yy2).Output()
	if err != nil {
		fmt.Printf("swip error : %s", err)
	}
	time.Sleep(sleepTime)
}

//func DeviceNames() {
//	_output, _ := exec.Command("adb", "devices").Output()
//	log.Printf("%s", _output)
//
//	var listStr = string(_output)
//	var array = strings.Split(listStr, "\n")
//	for i := 0; i < len(array); i++ {
//		var content = array[i]
//		if content != "" {
//			if i != 0 {
//
//			}
//		}
//	}
//}

func SendCommand(deviceName string, cmd string) {
	output, err := exec.Command("adb", "-s", deviceName, "shell", cmd).Output()
	if err != nil {
		log.Printf("SendCommand error %s", err)
	}
	log.Printf("SendCommand result %s", output)

}

//func AdbShellInputKeyEvent2(s string) {
//	exec.Command("adb", "shell", "input", "keyevent", s).Run()
//}
