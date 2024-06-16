BINARY := toodoo 

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/n] ' && read ans && [ $${ans:-N} = y]

# ensure there are no local unstaged changes
.PHONY: no-dirty
no-dirty:
	git diff --exit-code

# ==================================================================================== #
# QUALITY CHECKS 
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	@echo 'Tidying and formatting code...'
	go fmt ./...
	go mod tidy -v

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: test
test:
	@echo 'Executing tests...'
	go test -v -race -buildvcs ./internal/...