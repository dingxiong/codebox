<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Contents**

- [run locally](#run-locally)
- [build docker](#build-docker)
- [k8s](#k8s)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# run locally

```
cd app
uvicorn app:app --reload --port 8123
```

# build docker

```
docker buildx build --platform linux/amd64,linux/arm64 --push -t dingxiong/fastapi-strawberry-example:latest .
```

# k8s

```
k apply -f fastapi-strawberry-example.yaml -n sample
```
