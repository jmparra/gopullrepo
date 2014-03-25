gopullrepo
==========

Pulls specified git repo every time the webserver is accessed (GitHub webhooks).

# Install

Install go, and then:

```
        go build .
        sudo mv gopullrepo /usr/bin # Or whatever you like
        gopullrepo --repo /path/to/repo
```

From now on, every time someone requests (with any method) `localhost:8080/`, the repo will be pulled.

You can change the listening address and port, as well as path:

```
        gopullrepo --repo /path/to/repo --address ":8080" --path "/webhook"
```
