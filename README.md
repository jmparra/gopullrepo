Go Pull Repo
============

Pulls specified git repo every time the webserver is accessed (Git webhooks).

# Usage

```
        docker run -d --name puller -v /path/to/repo:/repo -v /path/.ssh/id_rsa:/root/.ssh/id_rsa -p 8080:8080 antarkt/pullrepo
```

# Install

Install Go and Docker, and then compile:

```
        env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o puller *.go
        docker build -t puller .
```

From now on, every time someone requests (with any method) `localhost:8080/webhook`, the repo will be pulled.

You can change the listening address and port, as well as path:

```
        docker run -d --name puller -v /path/to/repo:/newrepo -v /path/.ssh/id_rsa:/root/.ssh/id_rsa -p 8080:8081 antarkt/pullrepo --repo /newrepo --address ":8081" --path "/webhook"
```
