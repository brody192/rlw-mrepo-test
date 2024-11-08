CONTAINER_NAME=service_test
IMAGE_NAME=tinygodsdev/service:test

docker:
	docker build -f Dockerfile -t $(IMAGE_NAME) --progress=plain ..
	docker run --rm --name $(CONTAINER_NAME) $(IMAGE_NAME)

run:
	go run cmd/main.go
