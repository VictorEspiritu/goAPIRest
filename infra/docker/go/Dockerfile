FROM golang:1.16.3

RUN go get github.com/cespare/reflex

ENTRYPOINT reflex -r "(\.go$|go\.mod)" -s go run cmd/http/http.go 
