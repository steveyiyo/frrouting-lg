# FRRouting Looking Glass

This is a looking glass for the Internet Routing Daemon "FRRouting"

### How it work?

Software is split in two parts:

- server.go:

Responsible for processing all requests, and forward it to the router(client.go)

- client.go:

Receive server.go request

### How to install?

Download ["FRRouting-LG.zip"](https://github.com/steveyiyo/frrouting-lg/releases/download/v1.1/FRRouting-LG.zip) to your server and run it.

```
wget https://github.com/steveyiyo/frrouting-lg/releases/download/v1.1/FRRouting-LG.zip && unzip FRRouting-LG.zip
```