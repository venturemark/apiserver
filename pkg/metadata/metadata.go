// Package metadata provides label constants used for resource identification
// and information sharing. We need to consider the following resources within
// the system.
//
//     metric
//     timeline
//     update
//     user
//
// The above resources have different properties. It might be necessary to
// expose different pieces of information throughout the lifecycle of the
// resources when interacting with them.
//
//     id        unix timestamp normalized to UTC timezone (does not apply to users)
//     status    state change information like "updated"
//
package metadata

const (
	MetricID     = "metric.venturemark.co/id"
	MetricStatus = "metric.venturemark.co/status"
)

const (
	TimelineID     = "timeline.venturemark.co/id"
	TimelineStatus = "timeline.venturemark.co/status"
)

const (
	UpdateID     = "update.venturemark.co/id"
	UpdateStatus = "update.venturemark.co/status"
)

const (
	UserID = "user.venturemark.co/id"
)
