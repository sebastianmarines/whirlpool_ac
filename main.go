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

var metrics = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "appliance_metrics",
		Help: "Appliance metrics",
	},
	[]string{"appliance", "metric"},
)

func main() {
	router := mux.NewRouter()

	prometheus.MustRegister(metrics)

	go func() {
		err, token := getToken()

		if err != nil {
			fmt.Println(err)
			return
		}

		appliances := token.SAID

		for {
			for _, appliance := range appliances {
				err, data := getData(appliance)

				if err != nil {
					fmt.Println(err)
					return
				}

				setMetricsForAppliance(data)

			}
			time.Sleep(60 * time.Second)
		}
	}()

	router.Path("/").HandlerFunc(HomeHandler)
	router.Path("/metrics").Handler(promhttp.Handler())

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

func HomeHandler(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func setMetricsForAppliance(appliance ApplianceData) {
	if appliance.Attributes.Online.Value != "1" {
		metrics.WithLabelValues(appliance.ApplianceId, "online").Set(0)
		metrics.WithLabelValues(appliance.ApplianceId, "humidity").Set(0)
	}
	// Temperature
	temperature := appliance.Attributes.SysOpstatusdisplaytemp.Value
	temperature = temperature[0:len(temperature)-1] + "." + temperature[len(temperature)-1:]
	temperatureFloat, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	metrics.WithLabelValues(appliance.ApplianceId, "temperature").Set(temperatureFloat)

	// Humidity
	humidity := appliance.Attributes.SysOpstatusdisplayhumidity.Value
	humidityFloat, err := strconv.ParseFloat(humidity, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	metrics.WithLabelValues(appliance.ApplianceId, "humidity").Set(humidityFloat)
}
