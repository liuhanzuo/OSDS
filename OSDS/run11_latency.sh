echo "started"
go run sync.go LATENCYTEST 300 300
go run sync.go LATENCYTEST 300 30
go run sync.go LATENCYTEST 30 30