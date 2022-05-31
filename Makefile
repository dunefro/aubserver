create:
	mkdir -p bin
	GOOS=linux go build -o bin/aub

run:
	./bin/aub

build:
	DOCKER_BUILDKIT=1 docker image build -t quay.io/vedant99/aubserver:v0.1.0  .
	docker image push quay.io/vedant99/aubserver:v0.1.0

deploy:
	kubectl apply -f kubernetes/deploy.yaml

delete:
	kubectl delete -f kubernetes/deploy.yaml