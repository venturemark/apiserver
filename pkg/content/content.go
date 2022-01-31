package content

import (
	"encoding/json"
	"github.com/xh3b4sd/tracer"
	"os"
)

type SerializedUpdate struct {
	Head string `json:"head"`
	Text string `json:"text"`
}

type SerializedTimeline struct {
	Desc    string             `json:"desc"`
	Name    string             `json:"name"`
	Updates []SerializedUpdate `json:"updates"`
}

type SerializedVenture struct {
	Prepopulate string               `json:"prepopulate"`
	Desc        string               `json:"desc"`
	Name        string               `json:"name"`
	Timelines   []SerializedTimeline `json:"timelines"`
}

func GetTemplateVenture(prepopulate string) (*SerializedVenture, error) {
	content, err := os.ReadFile("ventures.json")
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var ventures []SerializedVenture
	err = json.Unmarshal(content, &ventures)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var search string
	switch prepopulate {
	case "choiceNetwork":
		search = "ku"
	case "choiceProgress":
		search = "tmp"
	case "choiceTeam":
		search = "sis"
	}

	for _, v := range ventures {
		if v.Prepopulate == search {
			return &v, nil
		}
	}

	return nil, nil
}
