package content

import (
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
)

var DefaultVenture = venture.CreateI_Obj_Property{
	Desc: "CHANGEME",
	Name: "My Network",
}

var GettingStartedTimeline = timeline.CreateI_Obj_Property{
	Desc: "CHANGEME",
	Name: "Getting Started",
}

var GettingStartedPosts = []texupd.CreateI_Obj_Property{
	{
		Head: "Post 1: What is a network?",
		Text: "A Breadcrumb Network is",
	},
}

// Keep in sync with https://github.com/venturemark/webclient/blob/767d411768e3e4fd45b42f16c8ded208233d7698/src/component/OnboardingGroup.tsx#L104-L106
var DefaultTimelinesMap = map[string][]*timeline.CreateI_Obj_Property{
	"choiceNetwork": {
		{
			Desc: "a",
			Name: "a",
		},
	},
	"choiceProgress": {
		{
			Desc: "b",
			Name: "b",
		},
	},
	"choiceTeam": {
		{
			Desc: "b",
			Name: "b",
		},
	},
}
