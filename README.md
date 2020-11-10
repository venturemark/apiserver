# apiserver

Daemon for serving the venturemark grpc api. After creating a kubernetes cluster
using https://github.com/xh3b4sd/kia this api server app can be deployed and
used.



### EKS

```
helm -n infra install apiserver ./helm/apiserver --set cluster.name=kia02 --set cluster.zone=aws.venturemark.co
```

```
grpcurl apiserver.kia02.aws.venturemark.co:443 post.API/Search
```



### OSX

```
helm -n infra install apiserver ./helm/apiserver
```

```
grpcurl -plaintext 127.0.0.1:7777 post.API/Search
```
