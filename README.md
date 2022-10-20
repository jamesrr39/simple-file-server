# simple-file-server

simple-file-server implements a very simple HTTP file server, serving from disk.

It's ideal as a quick tool for exposing a directory across the network when e.g. sharing files between one computer or another.

It has functionality similar to running `python3 -m http.server --directory $DIR $PORT`

It's very easy to install and also very easy to audit, given that the codebase is very small.

## Installation

```
go install github.com/jamesrr39/simple-file-server@latest
```

## Usage

```
simple-file-server $DIR
```

## Limitations

This project does not support things such as allow/deny rules, setitng custom HTTP headers, etc. For these, there are plenty of other projects, e.g. Nginx, Apache web server, Caddy...
