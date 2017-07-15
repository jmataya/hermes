install:
	glide install

migrate-ci:
	go run main.go migrate --database ci_test --user postgres

test:
	go test `glide nv`

.PHONY: install test