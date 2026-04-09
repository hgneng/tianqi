# tianqi(天气)
A command line tool to query weather report in China. It can be run on Linux and Mac.

It's written in Go lang.

Run following command to build the package:

```$ go build tianqi.go```

Run following command to install the package:

```$ ./deploy.sh```

Usage:
```shell
$ tianqi guangzhou
当前天气阴温度26摄氏度，湿度百分之77，风力2级。

天气预报：
今天：多云，23到29摄氏度
明天：晴，23到31摄氏度
后天：多云，24到30摄氏度

$ tianqi 广州
当前天气阴温度26摄氏度，湿度百分之77，风力2级。

天气预报：
今天：多云，23到29摄氏度
明天：晴，23到31摄氏度
后天：多云，24到30摄氏度

$ tianqi guangzhou 6
09时：小雨，16度
10时：小雨，16度
11时：小雨，16度
12时：小雨，18度
13时：小雨，18度
14时：小雨，19度
```
