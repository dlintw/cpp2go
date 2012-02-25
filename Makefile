all : cpp2go

% : %.go
	go tool fix $<
	go tool vet $<
	gofmt -s -w $<
	go build -o $@ $<
clean:
	rm -f cpp2go

