install:
	glide install

migrate:
	go run main.go migrate --database hermes --user hermes

migrate-ci:
	go run main.go migrate --database hermes_ci_test --user postgres

migrate-test:
	go run main.go migrate --database hermes_test --user hermes

reset:
	dropdb --if-exists hermes
	dropuser --if-exists hermes
	createuser hermes
	createdb hermes

reset-test:
	dropdb --if-exists hermes_test
	createdb hermes_test

test:
	GOENV=test go test `glide nv`

test-ci:
 GOENV=ci go test `glide nv`

.PHONY: install test