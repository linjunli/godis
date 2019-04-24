
servergo=./src/main/main.go
serverbin=./bin/server

all : $(servergo)
$(servergo) : $(shell find $(servergo) ./bin/)
	go vet $(servergo)
	go build -o $(serverbin) -i $(servergo)
clean:
	rm bin/*
rebuild: $(clean) $(all)
