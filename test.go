package main

import (
    "github.com/jimlawless/cfg"
        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot-spark"
        "github.com/hybridgroup/gobot-gpio"
//        "time"
        "fmt"
        "net/http"
        "log"
)

// Used for to pass a chanel to the handler

func handler(controlChan chan int, w http.ResponseWriter, r *http.Request) {
    controlChan <- 1
    fmt.Fprintf(w, "<html><head><body><form name=SetLights, action=\"/lights\", method=\"post\"><input type=\"radio\" name=\"light\">Red</input><input type=\"submit\" name=\"Set\"></form></body></html>")
}


func main() {

    // device_id and access_token
    config := make(map[string]string)
    err := cfg.Load("test.cfg", config)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%v\n", config)

        var controlChan = make(chan int)

        spark := new(gobotSpark.SparkAdaptor)
        spark.Name = "spark"
        spark.Params = make(map[string]interface{})
        spark.Params["device_id"] = config["device_id"]
        spark.Params["access_token"] = config["access_token"]

        led := gobotGPIO.NewLed(spark)
        led.Name = "led"
        led.Pin = "D7"


        work := func() {
            //gobot.Every("1s", func() {
            //            led.Toggle()
            //})
            for {
                _ = <- controlChan
                led.Toggle()
            }

        }

        robot := gobot.Robot{
                Connections: []interface{}{ spark }, 
                Devices:     []interface{}{ led },
                Work:        work,
        }

        go robot.Start()
        //for {
            //control <- 1
            //time.Sleep(1000 * time.Millisecond)
        //}


        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            handler(controlChan, w, r)
        })

        http.ListenAndServe(":8080", nil)

}
