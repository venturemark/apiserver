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



Searching for timelines in the first place should not lead to any result if we
never created any timeline before.

```
$ grpcurl -d '{"obj":[{"metadata":{"user.venturemark.co/id":"usr-al9qy"}}]}' -plaintext 127.0.0.1:7777 timeline.API.Search
{

}
```



Searching for metrics the first time against the apiserver does not result in
any data, since none got created yet. Note you need `grpcurl` installed.

```
$ grpcurl -d '{"obj":[{"metadata":{"timeline.venturemark.co/id":"1606488079","user.venturemark.co/id":"usr-al9qy"}}]}' -plaintext 127.0.0.1:7777 metric.API.Search
{

}
```



Before we can create a metric update we need to create a timeline like shown
below. We get a timeline ID which we further use working with the apiserver.

```
$ grpcurl -d '{"obj":{"metadata":{"user.venturemark.co/id":"usr-al9qy"},"property":{"name":"user retention"}}}' -plaintext 127.0.0.1:7777 timeline.API.Create
{
  "obj": {
    "metadata": {
      "timeline.venturemark.co/id": "1606488079"
    }
  }
}
```



Searching for timelines again shows the timeline we just created.

```
$ grpcurl -d '{"obj":[{"metadata":{"user.venturemark.co/id":"usr-al9qy"}}]}' -plaintext 127.0.0.1:7777 timeline.API.Search
{
  "obj": [
    {
      "metadata": {
        "timeline.venturemark.co/id": "1606488079",
        "user.venturemark.co/id": "usr-al9qy"
      },
      "property": {
        "name": "user retention"
      }
    }
  ]
}
```



Creating the first metric update against the apiserver can be done like this.
Right now only the timestamp of creation is returned.

```
$ grpcurl -d '{"obj":{"metadata":{"timeline.venturemark.co/id":"1606488079","user.venturemark.co/id":"usr-al9qy"},"property":{"data":[{"space":"y","value":[23]}],"text":"foo bar baz"}}}' -plaintext 127.0.0.1:7777 metupd.API.Create
{
  "obj": {
    "metadata": {
      "metric.venturemark.co/id": "1606488140",
      "update.venturemark.co/id": "1606488140"
    }
  }
}
```



Searching for metrics again shows the metric object we just created. Note the
automatically injected dimension t tracking the unix timestamp of each datapoint
emitted.

```
$ grpcurl -d '{"obj":[{"metadata":{"timeline.venturemark.co/id":"1606488079","user.venturemark.co/id":"usr-al9qy"}}]}' -plaintext 127.0.0.1:7777 metric.API.Search
{
  "obj": [
    {
      "metadata": {
        "metric.venturemark.co/id": "1606488140",
        "timeline.venturemark.co/id": "1606488079",
        "user.venturemark.co/id": "usr-al9qy"
      },
      "property": {
        "data": [
          {
            "space": "t",
            "value": [
              1.60648814e+09
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
$ grpcurl -d '{"obj":{"metadata":{"timeline.venturemark.co/id":"1606488079","update.venturemark.co/id":"1606488140","user.venturemark.co/id":"usr-al9qy"},"property":{"text":"changed"}}}' -plaintext 127.0.0.1:7777 metupd.API.Update
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
$ grpcurl -d '{"obj":[{"metadata":{"timeline.venturemark.co/id":"1606488079","user.venturemark.co/id":"usr-al9qy"}}]}' -plaintext 127.0.0.1:7777 update.API.Search
{
  "obj": [
    {
      "metadata": {
        "timeline.venturemark.co/id": "1606488079",
        "update.venturemark.co/id": "1606488140",
        "user.venturemark.co/id": "usr-al9qy"
      },
      "property": {
        "text": "changed"
      }
    }
  ]
}
```
