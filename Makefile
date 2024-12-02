DAYS = "1" "2"

run-day:
	go run $(day)/main.go

run-all: $(DAYS)

$(DAYS):
	go run $@/main.go

.PHONY: run-day run-all