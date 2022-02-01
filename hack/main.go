package main

import (
	"encoding/json"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apiserver/pkg/content"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var ventures []schema.Venture
	var timelines []schema.Timeline
	var updates []schema.Update

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		if file.Name() == "data.json" {
			continue
		}

		content, err := os.ReadFile(file.Name())
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(file.Name(), "venture-") {
			var object schema.Venture
			err = json.Unmarshal(content, &object)
			if err != nil {
				panic(err)
			}

			ventures = append(ventures, object)
		} else if strings.HasPrefix(file.Name(), "timeline-") {
			var object schema.Timeline
			err = json.Unmarshal(content, &object)
			if err != nil {
				panic(err)
			}

			timelines = append(timelines, object)
		} else if strings.HasPrefix(file.Name(), "update-") {
			var object schema.Update
			err = json.Unmarshal(content, &object)
			if err != nil {
				panic(err)
			}

			updates = append(updates, object)
		}
	}

	var output []content.SerializedVenture

	for _, v := range ventures {
		ventureID := v.Obj.Metadata[metadata.VentureID]
		var ventureTimelines []content.SerializedTimeline
		for _, t := range timelines {
			if t.Obj.Metadata[metadata.VentureID] != ventureID {
				continue
			}
			timelineID := t.Obj.Metadata[metadata.TimelineID]
			var timelineUpdates []content.SerializedUpdate
			for _, u := range updates {
				if u.Obj.Metadata[metadata.VentureID] != ventureID || u.Obj.Metadata[metadata.TimelineID] != timelineID {
					continue
				}
				timelineUpdates = append(timelineUpdates, content.SerializedUpdate{
					Head: u.Obj.Property.Head,
					Text: u.Obj.Property.Text,
				})
			}
			ventureTimelines = append(ventureTimelines, content.SerializedTimeline{
				Desc:    t.Obj.Property.Desc,
				Name:    t.Obj.Property.Name,
				Updates: timelineUpdates,
			})
		}
		splitName := strings.Split(v.Obj.Property.Name, " ")
		ventureName := strings.Join(splitName[0:len(splitName)-1], " ")
		prepopulate := strings.ToLower(strings.Trim(splitName[len(splitName)-1], "()"))
		output = append(output, content.SerializedVenture{
			Desc:        v.Obj.Property.Desc,
			Name:        ventureName,
			Prepopulate: prepopulate,
			Timelines:   ventureTimelines,
		})
	}

	outputContent, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("../ventures.json", outputContent, 0644)
	if err != nil {
		panic(err)
	}
}
