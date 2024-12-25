DAYS = "1" "2" "3" "4" "5" "6" "7" "8" "9" "10" "11" "12" "13" "14" "15" "16" "17" "18" "19" "20" "21" "22" "23" "24" "25"

run-day:
	go run $(day)/main.go

run-all: $(DAYS)

$(DAYS):
	go run $@/main.go

.PHONY: run-day run-all