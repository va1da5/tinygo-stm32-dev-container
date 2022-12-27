.PHONY: all
all: build flash

.PHONY: build
build:
	tinygo build -target=bluepill -o main.bin main.go

.PHONY: flash
flash:
	tinygo flash -target=bluepill


.PHONY: image
image:
	docker build -t tinygo-stm32 .


.PHONY: do
do:
	docker run --rm -it \
		--privileged \
		-u root \
		-v /dev:/dev \
		-v "$$PWD:/stm32" \
		-w "/stm32" tinygo-stm32 make all
