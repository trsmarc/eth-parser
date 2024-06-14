mock:
	mockgen -source=txstore/memory.go \
		-package mocks \
		-destination=mocks/txstore/memory.go
