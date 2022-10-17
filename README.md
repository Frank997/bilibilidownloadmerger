# bilibilidownloadmerger
b站手机离线视频合并


简述：本应用可将b站离线下载的audio.m4s和video.m4s合并为mp4文件


使用方法(命令格式)：
mergeBiliDown.exe <b站离线下载文件夹所在路径> <输出目录>

示例：mergeBiliDown.exe D:/bilibili/download D:/bilibiliOut

输出文件名格式：视频名称_bvid_avid_下载时间戳.mp4

依赖：
ffmpeg.exe。


注：mergeBiliDown会尝试从当前目录和环境变量PATH中查找ffmpeg.exe，若两者同时存在，优先当前目录


操作步骤：
- 拷贝b站离线下载视频目录到电脑。(b站离线下载文件夹位置：存储空间/Android/tv.danmaku.bili/download)
- 执行程序


构建：
go build mergeBiliDown.go
