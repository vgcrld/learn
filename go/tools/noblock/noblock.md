# Non Blocking Routines

A channel gets data and must have a receiver to not crash:

```go
ok  := make(chan string, 2)
bad := make(chan string)

// ok <- "info"
// bad <- "bad info"
```

OK is ok because it can take 2 "messages"
BAD is bad because it has nothing to read it.

```bash
rdavis@richs-mbp-2 noblock % go run main.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/rdavis/code/go/src/github.com/vgcrld/tools/noblock/main.go:9 +0x92
exit status 2
```

You cant send on a channel if nothing is open to get the data. If you buffer the channel then you are good.

