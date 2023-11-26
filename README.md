# b站app离线缓存合并器 —— bilibilidownloadmerger


简述：本应用可将b站app离线缓存功能下载的audio.m4s和video.m4s合并为mp4文件


使用方法(命令格式)：
```
mergeBiliDown.exe <b站离线下载文件夹所在路径> <输出目录>
```

示例：
```
mergeBiliDown.exe D:/bilibili/download D:/bilibiliOut
```
输出格式：视频名称_bvid_avid/视频分p名_下载时间戳.mp4

依赖：
ffmpeg.exe。


注：mergeBiliDown会尝试从当前目录和环境变量PATH中查找ffmpeg.exe，若两者同时存在，优先当前目录


### 操作步骤：
```
- 拷贝b站离线下载视频目录到电脑。(b站安卓app离线缓存文件夹位置：存储空间/Android/tv.danmaku.bili/download)
- 下载ffmpeg：https://ffmpeg.org/download.html#build-windows
- 下载本程序可执行文件：https://github.com/Frank997/bilibilidownloadmerger/releases
- 打开powershell，定位到程序所在目录，执行：mergeBiliDown.exe <b站离线下载文件夹所在路径> <输出目录>
```
注：程序执行时间取决于视频数量，请耐心等待，执行完毕后会打印处理成功数和失败数，若有错误，请拷贝程序输出，自行检查或发issue


构建：
```
go build mergeBiliDown.go
```
