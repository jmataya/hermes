install:
	glide install

test:
	go test `glide nv`

.PHONY: install test