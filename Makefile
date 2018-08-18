.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

install:
	go get ./cmd/project-sal

test:
	go test ./...

build: install
	GOOS=linux CGO_ENABLED=0 go build -ldflags "-X main.version=$(TAG)" ./cmd/project-sal

serve: build
	./project-sal

clean:
	rm ./project-sal

pack:
	GOOS=linux make build
	CGO_ENABLED=0
	sudo docker build -t sal:$(TAG) .
	sudo docker tag sal:$(TAG) 319131104487.dkr.ecr.us-west-2.amazonaws.com/sal:$(TAG)

upload:
	AWS_PROFILE=sal-prod sudo docker push 319131104487.dkr.ecr.us-west-2.amazonaws.com/sal:$(TAG)

deploy:
	envsubst < kubernetes/deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean
