# Vectorworks-Utility
Utility app for the Vectorworks application suite


To get started:
```shell
go get -u github.com/asticode/go-astilectron-bundler/...
go install github.com/asticode/go-astilectron-bundler/astilectron-bundler

astilectron-bundler
```

Build arguments:
-ldflags="-H windowsgui"

used to reduce file size of build
go build -ldflags="-s -w"