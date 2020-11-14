# Playground



You can build and run it like this if you have `go` installed. Note that the
server does not log anything at this point.

```
go build && ./apiserver daemon
```



This is how you can run redis given you have `docker` installed. Redis is used
as datastore.

```
docker run --rm -p 127.0.0.1:6379:6379 redis
```



Searching for metrics the first time against the apiserver does not result in
any data, since non got created yet. Note you need `grpcurl` installed.

```
$ grpcurl -plaintext -d '{"filter":{"property":[{"timeline":"tml-al9qy"}]}}' 127.0.0.1:7777 metric.API.Search
{

}
```



Creating the first metric update against the apiserver can be done like this.
Right now only the timestamp of creation is returned.

```
$ grpcurl -plaintext -d '{"yaxis":[5,40],"text":"Lorem ipsum ...","timeline":"tml-al9qy"}' 127.0.0.1:7777 metupd.API.Create
{
  "timestamp": "1605311478"
}
```



Searching for metrics again shows the metric object we just created.

```
$ grpcurl -plaintext -d '{"filter":{"property":[{"timeline":"tml-al9qy"}]}}' 127.0.0.1:7777 metric.API.Search
{
  "result": [
    {
      "yaxis": [
        "5",
        "40"
      ],
      "timeline": "tml-al9qy",
      "timestamp": "1605311478"
    }
  ]
}
```
