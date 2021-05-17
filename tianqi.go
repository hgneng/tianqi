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

var api = "http://weatherapi.market.xiaomi.com/wtr-v2/weather?imei=e32c8a29d0e8633283737f5d9f381d47&device=HM2013023&miuiVersion=JHBCNBD16.0&modDevice=&source=miuiWeatherApp&cityId="
var cityCodes = make(map[string]string)

func main() {
  if (len(os.Args) < 2) {
    fmt.Println("请提供城市名称或拼音，例如\"北京\"或\"beijing\"")
    os.Exit(0)
  }

  city := os.Args[1];
  code := getCityCode(city);

  if len(code) > 0 {
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
  } else {
    fmt.Println("没有找到城市" + city)
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
