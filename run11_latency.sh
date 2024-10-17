echo "started"
go run sync.go LATENCYTEST 3000 3000
go run sync.go LATENCYTEST 3000 300
go run sync.go LATENCYTEST 3000 30
go run sync.go LATENCYTEST 30000 30000