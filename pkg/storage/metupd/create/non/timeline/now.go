package timeline

import "time"

func now() func() time.Time {
	return time.Now
}
