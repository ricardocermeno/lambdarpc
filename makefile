.PHONY: all build push

IMAGE	?= ricardocermeno/lambdarpc
LATEST ?= latest

ifeq ($(shell test -e VERSION && echo -n yes),yes)
    VERSION = $(shell cat VERSION)
endif

all: build push-latest tag-version push-version

build:
	docker build -f ./build/package/Dockerfile -t  $(IMAGE):$(LATEST) .

push-latest:
	docker push $(IMAGE):$(LATEST)

tag-version:
	docker tag $(IMAGE):$(LATEST) $(IMAGE):$(shell cat VERSION)

push-version:
	docker push $(IMAGE):$(shell cat VERSION)


# Development tasks

run:
	@docker run --rm $(IMAGE)

cli:
	@docker run --rm -it $(IMAGE) ash

aws: command ?= --version
aws:
	make run command="aws $(command)"

sam: command ?= --version
sam:
	make run command="sam $(command)"

test:
	@docker run --rm -it $(IMAGE) help