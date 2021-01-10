package key

const (
	Audience = "usr:%s:aud"
)

// TODO we need to map all of the resources below to audiences instead of users.
const (
	Metric   = "usr:%s:tml:%s:met"
	Timeline = "usr:%s:tml"
	Update   = "usr:%s:tml:%s:upd"
)
