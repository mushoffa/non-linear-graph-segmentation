DATASET=signal1

run:
ifdef SIGNAL
	go run main.go -s $(SIGNAL) < ./dataset/$(SIGNAL)
else
	go run main.go -s $(DATASET) < ./dataset/$(DATASET)
endif