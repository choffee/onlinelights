package main

import (
    "github.com/hybridgroup/gobot"
    "github.com/hybridgroup/gobot-gpio"
    "github.com/hybridgroup/gobot-spark"
    "github.com/jimlawless/cfg"
    //        "time"
    "fmt"
    "log"
    "net/http"
)

type Rgb struct {
    red   uint8
    green uint8
    blue  uint8
}

// Used for to pass a chanel to the handler

func handler(controlChan chan Rgb, w http.ResponseWriter, r *http.Request) {
    controlChan <- Rgb{105, 105, 105}
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

    var controlChan = make(chan Rgb)

    spark := new(gobotSpark.SparkAdaptor)
    spark.Name = "spark"
    spark.Params = make(map[string]interface{})
    spark.Params["device_id"] = config["device_id"]
    spark.Params["access_token"] = config["access_token"]

    ledred := gobotGPIO.NewLed(spark)
    ledred.Name = "led"
    ledred.Pin = "A4"
    ledblue := gobotGPIO.NewLed(spark)
    ledblue.Name = "led"
    ledblue.Pin = "A5"
    ledgreen := gobotGPIO.NewLed(spark)
    ledgreen.Name = "led"
    ledgreen.Pin = "A6"

    work := func() {
        //gobot.Every("1s", func() {
        //            led.Toggle()
        //})
        var led Rgb
        for {
            led = <-controlChan
            ledred.Brightness(led.red)
            ledgreen.Brightness(led.green)
            ledblue.Brightness(led.blue)
        }

    }

    robot := gobot.Robot{
        Connections: []interface{}{spark},
        Devices:     []interface{}{ledred, ledblue, ledgreen},
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
