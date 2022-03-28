# tianqi(天气)
A command line tool to query weather report in China. It can be run on Linux and Mac.

It's written in Go lang.

Run following command to build the package:
$ go build tianqi.go

Run following command to install the package:
$ ./deploy.sh

Usage:
$ tianqi guangzhou
当前温度31摄氏度，湿度百分之71，多云。天气预报：今天，中雨转暴雨，33到25摄氏度；明天，暴雨转多云转暴雨，30到25摄氏度；后天，多云，32到25摄氏度；周4，多云，33到27摄氏度；周5，多云转多云转暴雨，34到27摄氏度

$ tianqi 广州
当前温度31摄氏度，湿度百分之71，多云。天气预报：今天，中雨转暴雨，33到25摄氏度；明天，暴雨转多云转暴雨，30到25摄氏度；后天，多云，32到25摄氏度；周4，多云，33到27摄氏度；周5，多云转多云转暴雨，34到27摄氏度

$ tianqi guangzhou 6
09时：小雨，16度
10时：小雨，16度
11时：小雨，16度
12时：小雨，18度
13时：小雨，18度
14时：小雨，19度