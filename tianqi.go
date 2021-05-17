package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
  "encoding/json"
  "time"
  "strconv"
  "math"
)

var cityCodes = make(map[string]string)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("天气查询命令，版本2.0")
    fmt.Println("请提供城市名称或拼音，例如\"北京\"或\"beijing\"")
    fmt.Println("后面再加一个数字可以查询未来数小时的具体天气，例如：`tianqi guangzhou 6`可查询未来6小时的天气。")
    fmt.Println("作者：黄冠能(hgneng at gmail.com)")
    os.Exit(0)
  }

  city := os.Args[1]
  code := getCityCode(city)

  hours := 0;
  if len(os.Args) > 2 {
    hours, _ = strconv.Atoi(os.Args[2])
  }

  if len(code) > 0 {
    if hours > 0 {
      queryTianqiApi(code, hours);
    } else {
      queryXiaomiApi(code);
    }
  } else {
    fmt.Println("没有找到城市" + city)
  }
}

func queryTianqiApi(code string, hours int) {
  // 可精确到小时的接口（每日限300次，付费2000元后终身每日10万）
  // 接口文档：https://www.tianqiapi.com/index/doc?version=v1
  var api2 = "https://www.tianqiapi.com/api?version=v1&appid=62864148&appsecret=XQj5TooL&cityid="

  out, err := exec.Command("sh", "-c", "wget -qO - '" + api2 + code + "'").Output();
  if err != nil {
    fmt.Println("查询失败，请检查网络连接")
    return
  } 

  if hours > 24 {
    fmt.Println("暂时只支持查询未来24小时的具体天气")
    return
  }

  if len(out) > 0 {
    var ret interface{}
    err := json.Unmarshal(out, &ret)
    if err != nil {
      fmt.Println("json error:", err)
      return
    }

    retMap := ret.(map[string]interface{})
    _, ok := retMap["errcode"]
    if ok {
      fmt.Println("天气接口查询出错，请联系作者hgneng at gmail.com")
      return
    }

    dataArray := retMap["data"].([]interface{})

    currentHour := time.Now().Hour()

    count := 0
    passed := false
    dayIndex := 0;

    for ; count < hours; dayIndex++ {
      dayMap := dataArray[dayIndex].(map[string]interface{})
      hoursArray := dayMap["hours"].([]interface{})

      for i := 0; i < len(hoursArray) && count < hours; i++ {
        hourMap := hoursArray[i].(map[string]interface{})
        h, _ := strconv.Atoi(hourMap["hours"].(string)[0:2])
        if !passed && h < currentHour {
          continue
        }

        passed = true

        fmt.Println(hourMap["hours"].(string) + "：" +
            hourMap["wea"].(string) + "，" +
            hourMap["tem"].(string) + "度")
        count++
      }
    }
  }
}

func queryXiaomiApi(code string) {
  // 小米接口（无限制）
  var api = "http://weatherapi.market.xiaomi.com/wtr-v2/weather?imei=e32c8a29d0e8633283737f5d9f381d47&device=HM2013023&miuiVersion=JHBCNBD16.0&modDevice=&source=miuiWeatherApp&cityId="

  out, err := exec.Command("sh", "-c", "wget -qO - '" + api + code + "'").Output();
  if err != nil {
    fmt.Println("查询失败，请检查网络连接")
  } 

  if len(out) > 0 {
    var data interface{}
    err := json.Unmarshal(out, &data)
    if err != nil {
      fmt.Println("json error:", err)
    }
    dataMap := data.(map[string]interface{})
    forecastMap := dataMap["forecast"].(map[string]interface{})
    realtimeMap := dataMap["realtime"].(map[string]interface{})
    day4 := strconv.Itoa(int(math.Mod(float64(time.Now().Weekday() + 3), 7))) 
    if day4 == "0" {
      day4 = "日"
    }
    day5 := strconv.Itoa(int(math.Mod(float64(time.Now().Weekday() + 4), 7))) 
    if day5 == "0" {
      day5 = "日"
    }

    fmt.Println("当前温度" + realtimeMap["temp"].(string) +
         "摄氏度，湿度百分之" + strings.Replace(realtimeMap["SD"].(string), "%", "", 1) +
         "，" + realtimeMap["weather"].(string) +
         "。天气预报：今天，" + forecastMap["weather1"].(string) + "，" +
         strings.Replace(strings.Replace(forecastMap["temp1"].(string), "℃~", "到", 1), "℃", "摄氏度", 1) +
         "；明天，" + forecastMap["weather2"].(string) + "，" +
         strings.Replace(strings.Replace(forecastMap["temp2"].(string), "℃~", "到", 1), "℃", "摄氏度", 1) +
         "；后天，" + forecastMap["weather3"].(string) + "，" +
         strings.Replace(strings.Replace(forecastMap["temp3"].(string), "℃~", "到", 1), "℃", "摄氏度", 1) +
         "；周" + day4 + "，" + forecastMap["weather4"].(string) + "，" +
         strings.Replace(strings.Replace(forecastMap["temp4"].(string), "℃~", "到", 1), "℃", "摄氏度", 1) +
         "；周" + day5 + "，" + forecastMap["weather5"].(string) + "，" +
         strings.Replace(strings.Replace(forecastMap["temp5"].(string), "℃~", "到", 1), "℃", "摄氏度", 1))
  }
}

func getCityCode(city string) string {
  citylist := "/usr/local/share/tianqi/citylist"
  out, err := exec.Command("sh", "-c", "grep '\"" + city + "\"' " + citylist + "| head -n 1").Output();
  if err != nil {
    fmt.Println(err)
  }
  
  if len(out) > 0 {
    items := strings.Split(string(out), ":")
    if len(items) == 2 {
      return strings.TrimRight(strings.TrimLeft(items[1], "\""), "\"")
    }
  }

  return "";
}
