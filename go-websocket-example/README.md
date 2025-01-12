<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Contents**

- [golang websocket exmaple](#golang-websocket-exmaple)
  - [Test steps](#test-steps)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# golang websocket exmaple

Follow https://blog.logrocket.com/using-websockets-go/

Want to verify that golang's goroutine can execute concurrently.

## Test steps

1. `go run server.go`
2. open browser localhost:8080/job, so server will print every second
3. Meanwhile open localhost:8000, and type some message to check websocket is
   working

This veryfies that go can serve two requests concurrently.
