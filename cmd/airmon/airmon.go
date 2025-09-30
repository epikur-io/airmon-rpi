package main

// import (
// 	"github.com/epikur-io/airmon-rpi/internal/model/config"
// 	"github.com/epikur-io/airmon-rpi/internal/metrics"
// 	"github.com/labstack/echo/v4"
// )

// See: https://echo.labstack.com/docs/middleware/prometheus

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

var GIT_COMMIT = ""

func main() {
	fmt.Println("GIT_COMMIT =", GIT_COMMIT)
	e := echo.New()
	e.HideBanner = true

	customRegistry := prometheus.NewRegistry() // create custom registry for your custom metrics

	// Create a gauge metric without any label dimensions.
	tempGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temp_a0_c",
		Help: "Just for testing.",
	})

	// Register the metric with the registry.
	if err := customRegistry.Register(tempGauge); err != nil { // register your new counter metric with metrics registry
		log.Fatal(err)
	}

	// Set the gauge's value to 99.78.
	tempGauge.Set(99.78)

	customCounter := prometheus.NewCounter( // create new counter metric. This is replacement for `prometheus.Metric` struct
		prometheus.CounterOpts{
			Name: "custom_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
	)
	if err := customRegistry.Register(customCounter); err != nil { // register your new counter metric with metrics registry
		log.Fatal(err)
	}

	// ---- Dump metrics from the custom registry ----
	mfs, err := customRegistry.Gather()
	if err != nil {
		log.Fatal(err)
	}
	sr := &bytes.Buffer{} // os.Stdout
	enc := expfmt.NewEncoder(sr, expfmt.NewFormat(expfmt.TypeTextPlain))
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(sr.String())

	// ---- Dump metrics from the *default* registry too ----
	fmt.Println("\n--- Default Registry ---")
	mfs, err = prometheus.DefaultGatherer.Gather()
	if err != nil {
		log.Fatal(err)
	}
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			log.Fatal(err)
		}
	}

	// os.Exit(0) // !DEBUG

	e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		AfterNext: func(c echo.Context, err error) {
			customCounter.Inc() // use our custom metric in middleware. after every request increment the counter
		},
		Registerer: customRegistry, // use our custom registry instead of default Prometheus registry
	}))
	e.GET("/metrics", echoprometheus.NewHandlerWithConfig(echoprometheus.HandlerConfig{Gatherer: customRegistry})) // register route for getting gathered metrics data from our custom Registry

	if err := e.Start(":8042"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
