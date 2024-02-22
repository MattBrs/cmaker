BINARY_NAME="cmaker"

build:
	go build -o ${BINARY_NAME} main.go

deps:
	echo "no deps"

clean:
	rm -rf ${BINARY_NAME}


