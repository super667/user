IS_RPC := $(shell echo $(AppName) | grep -ic rpc)
IS_WORKER := $(shell echo $(AppName) | grep -ic worker)

ifeq "$(IS_RPC)" "1"
	WORKDIR := rpc
else ifeq "$(IS_WORKER)" "1"
	WORKDIR := worker
else
	WORKDIR := api
endif

BUILD_IMAGE_NAME := $(AppName):$(BUILD_ID)
CONTAINER_NAME := $(AppName)_$(BUILD_ID)


build:
	@echo 应用类型: $(WORKDIR)
	@docker build . -t $(BUILD_IMAGE_NAME) --build-arg MAINDIR=$(WORKDIR)
	@docker run -v `pwd`:/local --name $(CONTAINER_NAME) $(BUILD_IMAGE_NAME) cp /app /local
	@ls -al
	@docker rm -f $(CONTAINER_NAME)
	@docker rmi $(BUILD_IMAGE_NAME)
