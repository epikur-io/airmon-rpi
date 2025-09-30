# Wiring of the sensors and the raspberry pi

The SPS30 uses 5V and all the other sensors use 3.3V.
All the sensors can be accessed via I2C by connecting all the sensor outputs to the corresponding SDA1 (GPIO2) and SCL1 (GPIO3) inputs 

## Raspberry Pi 4 Pinout overview

![Raspberry Pi 4 Pinout](./../assets/images/rpi4-pinout.png)

Connect all the sensors in-/outputs to the I2C1 SDA (GPIO2) & I2C1 SCL (GPIO3), they can all share the same I2C/GPIO pins.