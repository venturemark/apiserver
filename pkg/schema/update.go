package schema

type Update struct {
	Obj UpdateObj `json:"obj"`
}

type UpdateObj struct {
	Metadata map[string]string `json:"metadata"`
	Property UpdateObjProperty `json:"property"`
}

type UpdateObjProperty struct {
	Text string `json:"text"`
}
