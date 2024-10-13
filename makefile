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

run11:
	bash run11_latency.sh
	bash run11_throughput.sh
run12:
	echo "we are running experiments multiple times for accuracy"
	echo "it may takes up to 3 hours to finish"
	bash run12_latency.sh
	python3 run12_throughput.py 10

plot11: run11
	bash plot11_latency.sh
	python3 plot11_throughput.py
plot12: run12
	bash plot12_latency.sh
	python3 plot12_throughput.py 10

part11: run11 plot11
	pdflatex part11.tex
part12: run12 plot12
	pdflatex part12.tex