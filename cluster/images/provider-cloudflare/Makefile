# ====================================================================================
# Setup Project

PLATFORMS := linux_amd64 linux_arm64
include ../../../build/makelib/common.mk

# ====================================================================================
# Options
IMAGE ?= $(BUILD_REGISTRY)/provider-cloudflare:$(VERSION)
TARGETOS := $(word 1, $(subst _, ,$(PLATFORM)))
TARGETARCH := $(word 2, $(subst _, ,$(PLATFORM)))
include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Image Build
img.build:
	@$(INFO) docker build $(IMAGE)
	@$(MAKE) -C ../../.. go.build PLATFORM=$(PLATFORM)
	@$(MAKE) img.build.shared
	@$(OK) docker build $(IMAGE)

img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/
	@docker build $(BUILD_ARGS) --build-arg TARGETOS=$(TARGETOS) --build-arg TARGETARCH=$(TARGETARCH) -t $(IMAGE) $(IMAGE_TEMP_DIR)

# ====================================================================================
# Image Publish
img.publish: img.build
	@$(INFO) docker push $(IMAGE)
	@docker push $(IMAGE)
	@$(OK) docker push $(IMAGE)

# ====================================================================================
# Image Promote
img.promote:
	@$(INFO) Promoting image from $(FROM_IMAGE) to $(TO_IMAGE)
	@docker pull $(FROM_IMAGE)
	@docker tag $(FROM_IMAGE) $(TO_IMAGE)
	@docker push $(TO_IMAGE)
	@$(OK) Promoted image from $(FROM_IMAGE) to $(TO_IMAGE)