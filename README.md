#  AirMon - Air quality monitoring for the Raspberry Pi

Monitor environmental air quality and air pollution

Supported sensors:

- [SPS30](https://sensirion.com/products/catalog/SPS30) (Particular sensor PM1.0, PM2.5, PM4, and PM10; ~ 0.3 µm to 10 µm in diameter)
- [SCD40](https://sensirion.com/de/produkte/katalog/SCD40) (CO2, Temperature)
- [BME280](https://www.az-delivery.de/products/gy-bme280) (Temperature, Humidity, Pressure)
- [MQ-2](https://www.az-delivery.de/en/products/gas-sensor-modul) Gas sensor (Methan/Propan, Hydrogen, Smoke, Alcohol vapors). Note that you might require a [ADS1115 A/D converter](https://www.az-delivery.de/en/products/analog-digitalwandler-ads1115-mit-i2c-interface)

Airmon-rpi exposes a prometheus metrics endpoint ([http://localhost:8042/metrics](http://localhost:8042/metrics)) containing all the sensor metrics, those will be displayed by Grafana ([http://localhost:3000/](http://localhost:3000/)). 

## Setup

You can find the wiring description of the hardware (sensors and raspberry pi) [here](./docs/raspberry_pi_wiring.md).

### Manually build the binary (no docker)

```bash
task build:linux:arm64 DEBUG=1 # include debug symbols, see `build` task inside `./Taskfile.yaml` for all possible build parameters.
# or for an older Raspberry Pi:
#   $ task build:linux:armhf
# build for current OS/architecture:
#   $ task build
```

### Docker compose setup (recommended)

**1. Prepare config & environment variables**

Change the environment variables (credentials etc.) for Grafana and the other containers if needed. You find the env-files in:

- `./docker/airmon-rpi/docker.env`
- `./docker/grafana/docker.env`
- `./docker/prometheus/docker.env`

**2. Copy the files to the Raspberry Pi**

```bash
# using scp:
scp ./ pi@raspberrypi:/home/pi/airmon

# using rsync;
rsync -avz ./ pi@raspberrypi:/home/pi/airmon
# Also copy default config...

# using rclone:
# First configure your Raspberry Pi as a remote:
rclone config

# Then copy:
rclone copy ./ pi-remote:/home/pi/airmon
# ...
```

**3. Run airmon on the Pi**

Login to the raspberry pi and build/start the docker containers

```bash
cd /home/pi/airmon; docker compose build; docker compose up -d
```

You can now access the Grafana dashboard via [http://raspberrypi:3000](http://raspberrypi:3000) and see the sensor outputs. The user credentials (`GF_SECURITY_ADMIN_USER` and `GF_SECURITY_ADMIN_PASSWORD`) are defined inside the `./docker/grafana/docker.env` file.
