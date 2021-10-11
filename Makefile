
.PHONY: build
build:
	$(info #Building...)
	go build -o ./bin/md5-generator ./cmd/md5-generator

.PHONY: run
run:
	$(info #Running...)
#	@$(MAKE) generate
	@$(MAKE) start-server


.PHONY: start-server
start-server:
	$(info #Start Application...)
	./bin/md5-generator

#PHONY: generate
#generate:
#	$(info #Generate JSON file...)
#	go run generate.go