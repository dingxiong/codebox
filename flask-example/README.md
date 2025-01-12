<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Contents**

- [build](#build)
  - [run in docker](#run-in-docker)
  - [run in k8s](#run-in-k8s)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# build

docker buildx build --platform linux/amd64,linux/arm64 --push -t
dingxiong/flask-example:latest .

## run in docker

```
docker pull dingxiong/flask-example:latest
docker run -i -p 8000:8000 -v app:/app --name flask-example dingxiong/flask-example
```

## run in k8s

k apply -f flask-example.yaml -n sample
