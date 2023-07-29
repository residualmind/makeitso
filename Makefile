INSTALL_DIR = /usr/local/bin
BIN_NAME = makeitso

all: $(BIN_NAME)

$(BIN_NAME): main.go prompts.go megabrain.go
	go build -o $(BIN_NAME)  .

install:
	cp $(BIN_NAME) $(INSTALL_DIR)/$(BIN_NAME)

clean:
	rm -f $(BIN_NAME)
