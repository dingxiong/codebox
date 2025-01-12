<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Contents**

- [Build](#build)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Build

```
GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags '-extldflags "-static"'
k cp kafka-sarama-example <pod_name>:/
```
