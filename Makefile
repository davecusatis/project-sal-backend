.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

install:
	go get ./cmd/project-sal

test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" ./cmd/project-sal

serve: build
	./news

clean:
	rm ./news

pack:
	GOOS=linux make build
	sudo docker build -t sal:$(TAG) .
	sudo docker tag sal:$(TAG) 319131104487.dkr.ecr.us-west-2.amazonaws.com/sal:$(TAG)

upload:
	AWS_PROFILE=sal-prod sudo docker push 319131104487.dkr.ecr.us-west-2.amazonaws.com/sal:$(TAG)

deploy:
	envsubst < k8s/deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean