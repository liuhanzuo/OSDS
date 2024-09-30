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
cleantxt:
	python3 clean_txt.py
multi_resize:
	go build -o multi_resize.exe multi_resize.go
	./multi_resize.exe
multi_resizeall: multi_resize clean