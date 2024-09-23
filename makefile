.PHONY: clean

clean:
	python3 clean.py
hello:
	go build -o mlr.exe hello.go
	./mlr.exe
helloall: hello clean
queue:
	go build -o queue.exe queue.go
	./queue.exe
queueall: queue clean