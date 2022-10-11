install: go.sum
	@echo "installing bdtool binary..."
	@go install -mod=readonly
.PHONY: install