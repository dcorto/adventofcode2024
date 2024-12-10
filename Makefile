DAYS = "1" "2" "3" "4" "5" "6" "7" "8" "9" "10"

run-day:
	go run $(day)/main.go

run-all: $(DAYS)

$(DAYS):
	go run $@/main.go

.PHONY: run-day run-all