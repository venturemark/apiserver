package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

// Search provides a filter primitive to lookup metrics associated with a
// timeline. An example of a timeline ID is shown below.
//
//     tml-al9qy
//
// A timeline refers to many updates. A timeline does also refer to many
// metrics. Between updates and metrics there is a one to one relationship.
//
//     tml:tml-al9qy:met    met-id23g, met-ofn2j, ...
//     tml:tml-al9qy:upd    upd-oh3f4, upd-lk34n, ...
//
// Referring to the example above, the metric object met-id23g and the update
// object upd-oh3f4 belong together. Each references the other via ID.
//
//     met:met-id23g    update_id: upd-oh3f4
//     upd:upd-oh3f4    metric_id: met-id23g
//
// With redis we use ZREVRANGE where scores are unix timestamps. That allows us
// to search for objects while having support for chunking.
//
// With redis we use ZRANGEBYSCORE where scores are unix timestamps. That allows
// us to search for objects while having support for the "bet" operator. One
// example is to show updates and metrics within a certain timerange.
//
func (t *Timeline) Search(obj *metric.SearchI) (*metric.SearchO, error) {
	return &metric.SearchO{}, nil
}
