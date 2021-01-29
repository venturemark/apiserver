package schema

type Message struct {
	Obj MessageObj `json:"obj"`
}

type MessageObj struct {
	Metadata map[string]string  `json:"metadata"`
	Property MessageObjProperty `json:"property"`
}

type MessageObjProperty struct {
	Text string `json:"text"`
	Reid string `json:"reid"`
}
