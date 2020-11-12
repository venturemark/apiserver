# apiserver

Daemon for serving the venturemark grpc api. After creating a kubernetes cluster
using https://github.com/xh3b4sd/kia this api server app can be deployed and
used.



### EKS

```
hlm -n infra install apiserver ./hlm/apiserver --set cluster.name=kia02 --set cluster.zone=aws.venturemark.co
```

```
grpcurl apiserver.kia02.aws.venturemark.co:443 post.API/Search
```



### OSX

```
hlm -n infra install apiserver ./hlm/apiserver
```

```
grpcurl -plaintext 127.0.0.1:7777 post.API/Search
```



### Usage

```
$ ./apiserver
Daemon for serving the venturemark grpc api.

Usage:
  apiserver [flags]
  apiserver [command]

Available Commands:
  daemon      Run the apiserver process and server traffic.
  help        Help about any command
  version     Print version information of this command line tool.

Flags:
  -h, --help   help for apiserver

Use "apiserver [command] --help" for more information about a command.
```

```
$ ./apiserver daemon -h
Run the apiserver process and server traffic.

Usage:
  apiserver daemon [flags]

Flags:
  -h, --help          help for daemon
      --host string   The host for binding the grpc server to. (default "127.0.0.1")
      --port string   The port for binding the grpc server to. (default "7777")
```

```
$ ./apiserver version
Git Commit    n/a
Go Version    go1.15.2
Go Arch       amd64
Go OS         darwin
Source        https://github.com/venturemark/apiserver
Version       n/a
```
