name := "pack"
outDir := "out"
target := outDir + "/" + name

# show recipes
default:
    just --list

# build main
build:
    go build -o {{target}}

# test
test:
    go test -v ./...

# build and run
run: build
    ./{{target}}

# clean project
clean:
    go clean
    rm -fr {{target}}