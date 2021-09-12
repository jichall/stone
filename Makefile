CC = `which go`

SRC = $(shell (find ./src -maxdepth 1 \( -name "*.go" ! -name "*_test.go" \)))
EXE = stone

all:
	$(CC) run $(SRC)

docker-build:
	docker build -t stone .

docker-run: docker-build
	docker run -d --network=host stone

docker-run-t:
	sudo docker run -it --entrypoint /bin/bash stone

vars:
	@echo COMPILER........$(CC)
	@echo SOURCE FILES...$(SRC)