package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var count = 0

// TODO 如果所有entry.json键排序固定的话，可以改成一个正则，一次匹配出这几个元素
var regexDownTitle = regexp.MustCompile(`^.*"download_subtitle":"([^",]*)[,"].*$`) //优先选这个名字，因为这个名字更长
var regexTitle = regexp.MustCompile(`^.*"title":"([^",]*)[,"].*$`)                 //[^",] 作用是匹配不包含引号和逗号的所有字符
var regexBvid = regexp.MustCompile(`^.*"bvid":"([^",]*)[,"].*$`)
var regexAvid = regexp.MustCompile(`^.*"avid":([^",]*)[,"].*$`)                  //avid 是数字，不需要引号
var regexDownTime = regexp.MustCompile(`^.*"time_create_stamp":([^",]*)[,"].*$`) //视频下载时间戳，单位为毫秒
var regexReplaceWindowsIllegalChars = regexp.MustCompile(`[/\\:\*\?"<>\|]`)      //非法windows文件名

func main() {
	//如果参数数量不符预期，输出命令格式，若传参多，忽略多余参数
	if len(os.Args) < 3 {
		fmt.Println("命令格式：mergeBiliDown bilibiliDownLoadDir outputDir")
		os.Exit(0) //正常退出
	}

	//获取输入输出路径
	rootDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(err)
	}
	outDir, err := filepath.Abs(os.Args[2])
	if err != nil {
		panic(err)
	}
	//为输入输出路径末尾添加路径分隔符
	rootDir += string(os.PathSeparator)
	outDir += string(os.PathSeparator)

	var files []string
	//取出所有video.m4s路径
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "video.m4s") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(files) < 1 {
		fmt.Println("共发现0个视频，程序退出")
		os.Exit(0)
	}

	vidCounts := 0
	vidM4sSuffixLength := len("video.m4s")

	//查找ffmpeg，首先查找环境变量path，如果没有，查找当前目录
	ffmpegExeStr := "ffmpeg.exe"
	//先查找 .\ 是否有 ffmpeg可执行文件
	ffmpegBin, err := exec.LookPath("." + string(os.PathSeparator) + ffmpegExeStr)
	if err != nil { //如果没有在当前目录找到ffmpeg.exe，则去Path变量(环境变量)中查找
		ffmpegBin, err = exec.LookPath(ffmpegExeStr)
		if err != nil { //找不到ffmpeg直接退出
			fmt.Println("错误：未找到", ffmpegExeStr, "，请将", ffmpegExeStr, "设置于环境变量中或放到当前目录后再执行本程序")
			fmt.Printf("Error: Unable to find binary at: %v\n", ffmpegExeStr)
			os.Exit(-1)
		}
	}

	//创建输出目录
	os.MkdirAll(outDir, os.ModePerm)

	for _, file := range files {
		audio := file[0:len(file)-vidM4sSuffixLength] + "audio.m4s"                    //audio.m4s和video.m4s在同一目录
		vidName := getVidName(file[0:len(file)-vidM4sSuffixLength] + "..\\entry.json") //entry.json在video.m4s的上级目录
		isSucceed := merge(ffmpegBin, audio, file, outDir+vidName+".mp4")
		// fmt.Println(file)
		if isSucceed {
			vidCounts++
		}
	}

	//打印结果
	fmt.Println("程序执行完毕：")
	fmt.Println("视频总数：", len(files))
	fmt.Println("成功处理：", vidCounts)
	if len(files) > vidCounts {
		fmt.Println("有", len(files)-vidCounts, "个视频合并失败，请查看控制台输出。")
	}
}

func getVidName(entryPath string) string {
	content, err := ioutil.ReadFile(entryPath)
	if err != nil {
		panic(err)
	}
	contentStr := string(content)
	// get video title
	match := regexDownTitle.FindStringSubmatch(contentStr)
	if len(match) < 2 {
		match = regexTitle.FindStringSubmatch(contentStr)
	}
	// fmt.Println(match[0])
	// match[0] is raw string, regex group match from index 1
	title := match[1]
	//get bvid
	match = regexBvid.FindStringSubmatch(contentStr)
	bvid := match[1]
	// println(title, bvid)
	//get avid
	match = regexAvid.FindStringSubmatch(contentStr)
	avid := match[1]
	//get download timestamp
	match = regexDownTime.FindStringSubmatch(contentStr)
	downTime := match[1]
	//gen filename for output video
	filename := title + "_" + bvid + "_" + avid + "_" + downTime
	// fmt.Println(filename)
	//remove windows illegal chars for filename
	filename = regexReplaceWindowsIllegalChars.ReplaceAllString(filename, "")
	// fmt.Println(filename)

	return filename
}

func merge(ffmpegBin string, audio string, video string, outFile string) bool {
	isSucceed := true
	//ffmpeg -i video.m4s -i audio.m4s -codec copy -n Output.mp4
	cmd := exec.Command(ffmpegBin, "-i", audio, "-i", video, "-codec", "copy", "-n", outFile)
	out, err := cmd.CombinedOutput() //out 为命令输出，发生错误时输出，平时不用输
	if err != nil {                  //发生错误，打印命令输出并返回false给调用者
		fmt.Printf("===================Error-Split-Line-Start=================\n")
		fmt.Printf("ffmpeg got error:\n%s\n", string(out))
		fmt.Printf("cmd.CombinedOutput() failed with: %s\n", err)
		fmt.Printf("===================Error-Split-Line-End=================\n")
		return !isSucceed
	}
	return isSucceed
}
