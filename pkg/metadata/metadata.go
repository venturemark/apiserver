// Package metadata provides label constants used for resource identification
// and information sharing. We need to consider the following resources within
// the system.
//
//     audience
//     metric
//     organization
//     timeline
//     update
//     user
//
// The above resources have different properties. It might be necessary to
// expose different pieces of information throughout the lifecycle of the
// resources when interacting with them.
//
//     id        unix timestamp normalized to UTC timezone
//     status    state change information e.g. "updated"
//
package metadata

const (
	AudienceID     = "audience.venturemark.co/id"
	AudienceStatus = "audience.venturemark.co/status"
)

const (
	MetricID     = "metric.venturemark.co/id"
	MetricStatus = "metric.venturemark.co/status"
)

const (
	MessageID = "message.venturemark.co/id"
)

const (
	OrganizationID = "organization.venturemark.co/id"
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
