#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

data=$(cat << EOF
https://beta.breadcrumb.so/1643331603055234833/feed
https://beta.breadcrumb.so/1643331603055234833/1643337330998008180/feed
https://beta.breadcrumb.so/1643331603055234833/1643336388127630533/feed
https://beta.breadcrumb.so/1643331603055234833/1643336623044103662/feed
https://beta.breadcrumb.so/1643331603055234833/1643337297735528779/feed
https://beta.breadcrumb.so/1643329691890366113/feed
https://beta.breadcrumb.so/1643329691890366113/1643329890005312500/feed
https://beta.breadcrumb.so/1643329691890366113/1643329722864159260/feed
https://beta.breadcrumb.so/1643329691890366113/1643329795470833993/feed
https://beta.breadcrumb.so/1643336212549921413/feed
https://beta.breadcrumb.so/1643336212549921413/1643336295940014089/feed
https://beta.breadcrumb.so/1643336212549921413/1643337968623704422/feed
https://beta.breadcrumb.so/1643336212549921413/1643337944086818125/feed
https://beta.breadcrumb.so/1643336212549921413/1643336255853257793/feed
EOF)

for i in $data; do
  venture=$(echo $i | cut -d '/' -f 4)
  timeline=$(echo $i | cut -d '/' -f 5)
  if [ $timeline = "feed" ]; then
    echo extracting venture $venture
    kubectl -n infra exec -it -c redis rfr-redis-failover-0 -- redis-cli --raw get ven:$venture | jq -r | jq > venture-$venture.json
    timelines=$(kubectl -n infra exec -it -c redis rfr-redis-failover-0 -- redis-cli --raw zrange ven:$venture:tim 0 100)
    for j in $timelines; do
      raw=$(echo $j | jq -r)
      timeline=$(echo $raw | jq -r '.obj.metadata."timeline.venturemark.co/id"')
      echo extracting timeline $venture/$timeline
      echo $raw | jq > timeline-$venture-$timeline.json
      updates=$(kubectl -n infra exec -it -c redis rfr-redis-failover-0 -- redis-cli --raw zrange ven:$venture:tim:$timeline:upd 0 100)
      echo updates $updates
      for k in $updates; do
        raw=$(echo $k | jq -r)
        update=$(echo $raw | jq -r '.obj.metadata."update.venturemark.co/id"')
        echo extracting update $venture/$timeline/$update
        echo $raw | jq > update-$venture-$timeline-$update.json
      done
    done
  fi
done
