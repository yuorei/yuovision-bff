fmt:
	./shell/fmt.sh

gen:
	./shell/gen.sh
	./shell/fmt.sh

lint:
	./shell/lint.sh

test:
	go test -v ./...

prod:
	set -a && source .env.prod && set +a&&\
	go run main.go

dev:
	set -a && source .env.dev && set +a&&\
	go run main.go
