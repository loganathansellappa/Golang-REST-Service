echo "Tests starting......"
cd app/tests && go test -v
echo "App starting......"
cd ../../ && go run main.go

