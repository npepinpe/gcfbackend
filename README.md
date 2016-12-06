# Geo Campaign Framework Backend
Backend server for the Geo Campaigns Framework

# Development

Install [godep](https://github.com/tools/godep), clone the project, then install dependencies.

```
go get github.com/tools/godep
git clone git@github.com:npepinpe/gcfbackend.git
cd gcfbackend
godep get
```

When testing with dependencies, remember to preface all `go` commands with `godep`, e.g.:

```
godep go build
godep go test
```

# Secrets

You will need a secrets file for development, which you can obtain by pinging any of the devs.
