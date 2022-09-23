publish-dev:
	docker build --no-cache \
		-f .ci/Dockerfile \
		-t 074531296166.dkr.ecr.ap-southeast-1.amazonaws.com/service-campaign-slip:dev \
		--build-arg ENV_FILE=.env.dev \
		.
	docker push 074531296166.dkr.ecr.ap-southeast-1.amazonaws.com/service-campaign-slip:dev


publish-prod:
	docker build --no-cache \
		-f .ci/Dockerfile \
		-t 074531296166.dkr.ecr.ap-southeast-1.amazonaws.com/service-campaign-slip:latest \
		--build-arg ENV_FILE=.env.prod \
		.
	docker push 074531296166.dkr.ecr.ap-southeast-1.amazonaws.com/service-campaign-slip:latest
