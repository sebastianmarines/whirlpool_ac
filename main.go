package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

var metrics = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "appliance_metrics",
		Help: "Appliance metrics",
	},
	[]string{"appliance", "metric"},
)

func main() {
	router := mux.NewRouter()

	prometheus.MustRegister(metrics)

	go func() {
		for {
			err, data := getData("WPR4AJV7DN97D")

			if err != nil {
				fmt.Println(err)
				return
			}

			temperature := data.Attributes.SysOpstatusdisplaytemp.Value
			temperature = temperature[0:len(temperature)-1] + "." + temperature[len(temperature)-1:]
			fmt.Println(temperature)
			temperatureFloat, err := strconv.ParseFloat(temperature, 64)
			if err != nil {
				fmt.Println(err)
				return
			}

			metrics.WithLabelValues(data.ApplianceId, "temperature").Add(temperatureFloat)
			time.Sleep(60 * time.Second)
		}
	}()

	router.Path("/").HandlerFunc(HomeHandler)
	router.Path("/metrics").Handler(promhttp.Handler())

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}
