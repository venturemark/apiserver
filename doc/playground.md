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
any data, since none got created yet. Note you need `grpcurl` installed.

```
$ grpcurl -d '{"obj":[{"metadata":{"venturemark.co/timeline":"tml-kn433"}}]}' -plaintext 127.0.0.1:7777 metric.API.Search
{

}
```



Creating the first metric update against the apiserver can be done like this.
Right now only the timestamp of creation is returned.

```
$ grpcurl -d '{"obj":{"metadata":{"venturemark.co/timeline":"tml-kn433"},"property":{"data":[{"space":"y","value":[23]}],"text":"foo bar baz"}}}' -plaintext 127.0.0.1:7777 metupd.API.Create
{
  "obj": {
    "metadata": {
      "venturemark.co/unixtime": "1605741130"
    }
  }
}
```



Searching for metrics again shows the metric object we just created. Note the
automatically injected dimension t tracking the unix timestamp of each datapoint
emitted.

```
$ grpcurl -d '{"obj":[{"metadata":{"venturemark.co/timeline":"tml-kn433"}}]}' -plaintext 127.0.0.1:7777 metric.API.Search
{
  "obj": [
    {
      "metadata": {
        "venturemark.co/timeline": "tml-kn433",
        "venturemark.co/unixtime": "1605741130"
      },
      "property": {
        "data": [
          {
            "space": "t",
            "value": [
              1605741130
            ]
          },
          {
            "space": "y",
            "value": [
              23
            ]
          }
        ]
      }
    }
  ]
}
```



Updating the text of a metric update is shown below. Note the response metadata
indicating which property got updated.

```
$ grpcurl -d '{"obj":{"metadata":{"venturemark.co/timeline":"tml-kn433","venturemark.co/unixtime": "1605741130"},"property":{"text":"changed"}}}' -plaintext 127.0.0.1:7777 metupd.API.Update
{
  "obj": {
    "metadata": {
      "update.venturemark.co/status": "updated"
    }
  }
}
```



Searching for the text of a metric update shows the updated content after the
update call above.

```
$ grpcurl -d '{"obj":[{"metadata":{"venturemark.co/timeline":"tml-kn433"}}]}' -plaintext 127.0.0.1:7777 update.API.Search
{
  "obj": [
    {
      "metadata": {
        "venturemark.co/timeline": "tml-kn433",
        "venturemark.co/unixtime": "1605741130"
      },
      "property": {
        "text": "changed"
      }
    }
  ]
}
```
