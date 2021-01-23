// Package metadata provides label constants used for resource identification
// and information sharing. We need to consider the following resources within
// the system.
//
//     audience
//     message
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
