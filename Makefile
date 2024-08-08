.SILENT:

default: test audit auditstrict

.PHONY: test
test: 
	echo "test race"
	go test -race -buildvcs -count=1 ./...

.PHONY: audit
audit:
	echo "modcheck"
	go mod verify
	echo "vetcheck"
	go vet ./...
	echo "staticcheck"
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	echo "lintcheck"
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --max-same-issues 0 --max-issues-per-linter 0

.PHONY: auditstrict
auditstrict:
	echo "vulncheck"
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	echo "nilaway"
	go run go.uber.org/nilaway/cmd/nilaway@latest ./...
