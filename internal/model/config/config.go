package config

import (
	"time"
)

type Config struct {
	Sensors struct {
		SPS30  Sensor `json:"sps30" yaml:"sps30"`
		BME280 Sensor `json:"bme280" yaml:"bme280"`
		MQ2    Sensor `json:"mq2" yaml:"mq2"` // This sensor requires an ADC in order to be able to read precise measurements.
		SCD40  Sensor `json:"scd40" yaml:"scd40"`
	}
}

type Sensor struct {
	Enabled      bool          `json:"enabled" yaml:"enabled"`
	I2cAddress   uint16        `json:"i2c_address" yaml:"i2c_address"`
	ReadInterval time.Duration `json:"read_interval" yaml:"read_interval"`
}
