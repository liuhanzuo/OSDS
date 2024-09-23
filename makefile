.PHONY: clean

clean:
	python3 clean.py
queue:
	go build -o queue.exe queue.go
	./queue.exe
queueall: queue clean
sync:
	go build -o sync.exe sync.go
	./sync.exe
syncall: sync clean