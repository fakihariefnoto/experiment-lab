usage: FORCE
	# See
	exit 1

FORCE:

sqlite : FORCE
	@echo "building app.."
	@CGO_ENABLED=1 go build
	@echo "run app"
	@./sqlite

.PHONY: default
default: sqlite

.DEFAULT_GOAL := sqlite