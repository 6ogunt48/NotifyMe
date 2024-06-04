IMAGE_NAME = checkemailbot
IMAGE_TAG = latest

build:
	docker buildx build -t $(IMAGE_NAME):$(IMAGE_TAG) .

run:
	docker run --rm $(IMAGE_NAME):$(IMAGE_TAG)

clean:
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG)