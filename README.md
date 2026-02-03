go tool cover -html=coverage.out

go test ./service/... -coverprofile=coverage.out
