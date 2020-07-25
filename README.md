# net/http backend template for [Create Go App CLI](https://github.com/create-go-app/cli)

<img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://gocover.io/github.com/create-go-app/net_http-go-template/pkg/apiserver" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-49%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/net_http-go-template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" />

Package `net` provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. Although the package provides access to low-level networking primitives.

Package [`net/http`](https://golang.org/pkg/net/http/) provides HTTP client and server implementations.

## ⚡️ Quick start

1. Create a new project:

```bash
cgapp create
```

2. Run project by this command:

```bash
task -s
```

> ☝️ We're using `Taskfile` as task manager for running project on a local machine by default. If you've never heard of `Taskfile`, we recommend to read the [Docs](https://taskfile.dev/#/usage?id=getting-started) and use it, instead of `Makefile`.

## ✅ Used packages

- [gorilla/mux](https://github.com/gorilla/mux) `v1.7.4`
- [go-yaml/yaml](https://github.com/go-yaml/yaml) `v2.3.0`

## 🗄 Template structure

```bash
.
├── .dockerignore
├── .editorconfig
├── .gitignore
├── Dockerfile
├── Taskfile.yml
├── go.mod
├── go.sum
├── main.go
├── configs
│   └── apiserver.yml
├── static
│   └── index.html
└── pkg
    └── apiserver
        ├── config.go
        ├── config_test.go
        ├── error_checker.go
        ├── error_checker_test.go
        ├── new_server.go
        ├── new_server_test.go
        ├── routes.go
        ├── utils.go
        └── utils_test.go

4 directories, 17 files
```

## ⚙️ Configuration

```yaml
# ./configs/apiserver.yml

# Server config
server:
  host: 0.0.0.0
  port: 5000

# Database config
database:
  host: 127.0.0.1
  port: 5432
  username: postgres
  password: 1234

# Static files config
static:
  prefix: /
  path: ./static
```

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
