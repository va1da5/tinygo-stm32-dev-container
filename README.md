# TinyGo on STM32 Blue Pill

The repo contains a basic I2C device scanner. It focuses on providing an initial development environment for Tinygo with STM32 "Blue Pill". The development should be done in a container environment simplifying setup on different systems (*had issues on Fedora 37*).

## Development Container

Follow [Dev Containers tutorial](https://code.visualstudio.com/docs/devcontainers/tutorial) to get started with the development in containers. The repository contains [the development container configuration](./.devcontainer/devcontainer.json) for Tinygo development.

### Manual Approach

It is possible to use container only for building and flashing binaries to STM32 micro-controller. The commands below builds required container image and performs required actions.

```bash
# build dedicated container
make image

# access container shell
docker run --rm -it --entrypoint=/bin/bash \
    --privileged \
    -u root \
    -v /dev:/dev \
    -v "${PWD}:/${PWD##*/}" \
    -w "/${PWD##*/}" tinygo-stm32 make all

```

## Device Drivers

The below command pulls [device drivers](https://tinygo.org/docs/reference/devices/) for Tinygo development environment.

```bash
go get tinygo.org/x/drivers
```

## Serial

```bash
#/dev/ttyUSB0 @ 115200bit/s
# monitor serial output
tinygo monitor -baudrate 115200 

```

## Connection Diagram

![connection diagram](./diagram.png)


## STM32 "Blue Pill" Pin Diagram

![pinout](./STM32-Pin-Details.png)

## References

- [ST Micro STM32F103XX "Bluepill"](https://tinygo.org/docs/reference/microcontrollers/bluepill/)
- [Blue Pill](https://stm32-base.org/boards/STM32F103C8T6-Blue-Pill.html)
- [Blinking LED](https://tinygo.org/docs/tutorials/blinky/)
- [Container tinygo/tinygo](https://hub.docker.com/r/tinygo/tinygo)
- [Unable to start debug session with OpenOCD/Clion](https://stackoverflow.com/questions/71608471/unable-to-start-debug-session-with-openocd-clion)
- [STMicroelectronics/OpenOCD](https://github.com/STMicroelectronics/OpenOCD)
- [Developing inside a Container](https://code.visualstudio.com/docs/devcontainers/containers)
- [Using the I2C Bus](https://www.robot-electronics.co.uk/i2c-tutorial)
- [Go by Example](https://gobyexample.com/)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
