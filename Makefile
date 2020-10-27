

ENV_FLAGS=CGO_ENABLED=0 \
	GO111MODULE=on \
	GOOS=linux \
	GOARCH=amd64 \
	
golang-api:	
	cd golang && ${ENV_FLAGS} go build -mod=vendor -o bin/golang-api .