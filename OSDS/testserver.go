package main

func main() {
	sum := 0
	for i := 1; i < 1000000000; i++ {
		sum += i * i * (i + 1)
	}
}
