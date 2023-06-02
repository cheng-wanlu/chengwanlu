# Airtest

[官方文档](https://airtest.doc.io.netease.com/)

为什么要使用命令行来执行脚本
想同时用多个命令行运行多台手机、多个脚本等情况, 以及对于一些Python开发者来说, 可能需要在脚本中使用其他功能强大的Python第三方库. 因此推荐在本地python环境中安装airtest和pocoui, 然后用命令行运行脚本.

准备环境
首先必须要有python环境

使用 pip 安装Airtest模块
```bash
$ pip install airtest
```

注意: 在Mac/Linux系统下, 需要手动赋予adb可执行权限, 否则可能在执行脚本时遇到Permission denied的报错
```bash
# mac系统
$ cd {your_python_path}/site-packages/airtest/core/android/static/adb/mac
# linux系统
# cd {your_python_path}/site-packages/airtest/core/android/static/adb/linux
$ chmod +x adb
```

使用 pip 安装poco框架
```bash
$ pip install pocoui
```
环境部署完成后, 我们就能够脱离AirtestIDE, 在不同的宿主机器和被测平台上运行脚本了.

执行命令
```bash
# 这两个命令行的效果是相同的，我们用airtest运行了一个叫做untitled.air的脚本
>airtest run untitled.air --device Android:///手机设备号 --log log/
>python -m airtest run untitled.air --device Android:///手机设备号 --log log/
```

# 可以使用python sys模块查看site-packages文件夹所在的路径
```python
import sys
print(sys.path)
```

参数说明:
airtest run 后面接的的脚本文件的路径.
--device, 是我们的手机设备.
--log, log输出目录.

关于--device的说明
在刚才的命令行中使用的--device参数, 传入的是一个设备字符串, 以安卓设备为例, 字串完整定义如下:
```text
Android://<adbhost[localhost]>:<adbport[5037]>/<serialno>
```

其中, adbhost是adb server所在主机的ip(本地host为127.0.0.1), adb port默认是5037, serialno是android手机的序列号. 更多adb的方面的内容请参考文档developer.google.

在平时的脚本运行中, 我们一般可以这样写:
```text
# 什么都不填写, 会默认取当前连接中的第一台手机
Android:///
# 连接本机默认端口连的一台设备号为79d03fa的手机
Android://127.0.0.1:5037/79d03fa
# 连接一个Windows窗口, 窗口句柄为123456
Windows:///123456
# 连接一个Windows窗口, 窗口名称匹配某个正则表达式
Windows:///?title_re=Unity.*
# 连接iOS手机
iOS:///127.0.0.1:8100
```

生成测试报告
```text
参数说明
--log_root: 指定脚本执行完毕后生成log文件夹的路径
--outfile: 指定输出的html文件路径
--lang: 指定输出的语言, zh为中文
```

```bash
$ airtest report untitled.air --log_root log/ --outfile log/log.html --lang zh
```
