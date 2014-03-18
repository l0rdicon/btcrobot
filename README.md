BTC Robot
===========
BTC操盘手自动化交易引擎

  btcrobot is a Bitcoin, Litecoin and Altcoin trading bot written in golang,
  it features multiple trading methods using technical analysis.

  Disclaimer:

  USE AT YOUR OWN RISK!

  The author of this project is NOT responsible for any damage or loss caused
  by this software. There can be bugs and the bot may not perform as expected
  or specified. Please consider testing it first with paper trading /
  backtesting on historical data. Also look at the code to see what how
  it's working.

  Weibo:http://weibo.com/bocaicfa

# 使用指南， 用前必读！！！！
https://github.com/philsong/btcrobot/wiki/%E6%9C%BA%E5%99%A8%E4%BA%BA%E4%BD%BF%E7%94%A8%E6%8C%87%E5%8D%97 # 

BTC捐赠地址：1NDnnWCUu926z4wxA3sNBGYWNQD3mKyes8


# 本地搭建 #

0、

0.0安装golang开发运行环境，选择适合自己电脑操作系统的安装包
  
  http://code.google.com/p/go/downloads/list

0.1安装git环境

  http://code.google.com/p/msysgit/downloads/list

1、下载 btcrobot 代码
	
	git clone https://github.com/philsong/btcrobot

2、下载安装依赖库（如果依赖库下载不下来可以联系我）

	cd btcrobot/thirdparty
	// windows 下执行
	getpkg.bat
	// linux/mac 下执行
	sh getpkg

3、编译 btcrobot

先编译

	// 接着上一步
	cd ../
	// windows 下执行
	install.bat
	// linux/mac 下执行
	sh install
	
这样便编译好了 btcrobot

4、配置秘钥API文件

修改btcrobot/websites/code/btcbot/conf目录下的secret.sample文件名为secret.json

5、运行 btcrobot。

	// windows 下执行
	start.bat
	// linux/mac 下执行
	sh start

一切顺利的话，btcrobot应该就启动了。

6、浏览器中查看

在浏览器中输入：http://127.0.0.1:9090

应该就能看到了。

此时，你可以通过WEB界面配置各种参数，参数实时生效。


注：在第5步运行前可以根据需要修改 conf目录里的 配置，亦可在第6步配置。
