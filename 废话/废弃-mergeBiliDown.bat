bat 真他妈恶心，换用别的语言写

@chcp 65001
@echo off
rem 改变读取文件时的编码为utf8
rem 检查参数，若参数不全，打印使用格式
rem TODO

rem 创建输出目录
rem md %2

rem 遍历目录，找出所有video.m4s文件
for /r E:\tmp-winscp-trans-20220811\bilidownload %%I in (*video.m4s) do (
    rem echo %%I
    echo "中文"
    rem set video=%%I
    rem set audio=%%~dpIaudio.m4s
    set entry=%%~dpI..\entry.json
    rem echo %str:~0,10%
    rem echo %video%
    rem echo %audio%
    echo %entry:~0,10%
    rem findstr "[0-9]" %entry%
    rem 拼接audio.m4s文件路径
    rem 拼接 entry.json 文件路径
    rem 从entry.json路径取出视频名称
    rem 移除视频名称中的windows禁止作为文件名的字符
    rem 拼接输出文件路径
    rem 用ffmpeg合并视频
    rem 计数
 )

rem 打印合并的视频数目