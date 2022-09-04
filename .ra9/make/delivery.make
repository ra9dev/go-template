DOCKER_TAG ?= go-template
docker-build:
	$(call describe_job,"Building docker image '$(DOCKER_TAG)'")
	docker build -f .ra9/Dockerfile -t $(DOCKER_TAG) .