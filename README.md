![Build Status](https://github.com/zemags/go_workshop_1/actions/workflows/pipeline.yml/badge.svg)<br>

**App structure**
* cmd
  * app name folder (in case **workshop**)
    * main.go
* internal (interтal libs, like a helper)
  * handler folder
    * handler.go (structures, all dependencies)
* Dockerfile
* docker-compose.yml
* workflows

##### Run linters local
```bash
golangci-lint run cmd/workshop/main.go
```
##### Generate mocks
```bash
mockery --name=Client --dir=internal/api --output=internal/api/mocks
```
##### For mocks generating
- mockery, ginkgo
