package metrics

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Gauges for current values
	temperatureGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "Current temperature in Celsius",
	})

	humidityGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "humidity_percent",
		Help: "Current humidity percentage",
	})

	gpioStateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gpio_state",
		Help: "Current GPIO pin states (0=low, 1=high)",
	}, []string{"pin"})

	// Counters for events
	i2cReadsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "i2c_reads_total",
		Help: "Total number of I2C read operations",
	})

	gpioChangesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "gpio_changes_total",
		Help: "Total number of GPIO state changes",
	}, []string{"pin"})

	// Histograms for timing
	i2cOperationDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "i2c_operation_duration_seconds",
		Help:    "Duration of I2C operations",
		Buckets: prometheus.DefBuckets,
	})
)

func Example() {
	// Start metrics server
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting metrics server on :9090")
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Application started")

	// Main application loop
	for {
		start := time.Now()

		// Simulate reading from I2C devices
		readSensors()

		// Simulate GPIO operations
		readGPIOs()

		// Record operation duration
		i2cOperationDuration.Observe(time.Since(start).Seconds())

		time.Sleep(5 * time.Second)
	}
}

func readSensors() {
	// Simulate sensor readings
	temperatureGauge.Set(23.5 + (float64(time.Now().Second()%10) - 5))
	humidityGauge.Set(45.0 + (float64(time.Now().Second()%15) - 7.5))
	i2cReadsTotal.Inc()
}

func readGPIOs() {
	// Simulate GPIO readings
	pins := []string{"17", "18", "23", "24"}
	for _, pin := range pins {
		state := float64((time.Now().Unix() / 10) % 2) // Alternating state
		gpioStateGauge.WithLabelValues(pin).Set(state)

		// Simulate occasional changes
		if time.Now().Second()%30 == 0 {
			gpioChangesTotal.WithLabelValues(pin).Inc()
		}
	}
}
