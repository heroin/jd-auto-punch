GOPATH=$(PWD)

OBJ=$(PWD)/_obj
TARG=jd-auto-punch

main:
	$(GOROOT)/bin/go build -o $(OBJ)/$(TARG) ./src
	@echo "Build Done."

clean:
	rm -f $(OBJ)/$(TARG)
	rm -rf $(OBJ)

