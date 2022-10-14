###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building bdtool binary..."
	@go build -mod=readonly -o build/bdtool.exe .
else
	@echo "building bdtool binary..."
	@go build -mod=readonly -o build/bdtool .
endif
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing bdtool binary..."
	@go install -mod=readonly .
.PHONY: install