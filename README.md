go-chalk
========

A port of @sindresorhus' chalk library to golang. In node.js on windows libuv
(nodes compatibility layer) automatically parses ANSI codes and calls the
relevant Windows API to change the color. We don't get such niceties in golang
so currently the functions simply noop on windows. I have plans to create an
`os.Stdout` wrapper that will parse the ANSI codes but this is a very long term
goal.

# Install

```
go get github.com/danielchatfield/go-chalk
```

# Usage

```go
string msg = chalk.Cyan("some cyan text")
```
