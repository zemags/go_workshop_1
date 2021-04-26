![Build Status](https://github.com/zemags/go_workshop_1/actions/workflows/pipeline.yml/badge.svg)<br>

##### Run linters local
```bash
golangci-lint run cmd/workshop/main.go
```
##### Generate mocks
```bash
mockery --name=Client --dir=internal/api --output=internal/api/mocks
```