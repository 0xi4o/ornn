# Set the default shell
set shell := ["bash", "-c"]

# Default recipe (runs when you just type 'just')
default:
    @just --list

# Build the project
build:
    @echo "Building ornn..."
    go build .
    @echo "Done."


# Format the project
fmt:
    @echo "Formatting..."
    go fmt .
    @echo "Done."

# Run tidy
tidy:
    @echo "Tidying up..."
    go mod tidy
    @echo "Done."
