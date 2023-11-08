setup:
	go install golang.org/x/perf/cmd/benchstat@latest
bench-check:
	go test -bench=. -count 10 -benchmem
bench-gen-new:
	go test -bench=. -count 10 -benchmem | tee new.txt
bench-gen-old:
	go test -bench=. -count 10 -benchmem | tee old.txt
bench-compare:
	benchstat old.txt new.txt
run:
	go run main.go