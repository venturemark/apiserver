package schema

type Timeline struct {
	Obj TimelineObj `json:"obj"`
}

type TimelineObj struct {
	Metadata map[string]string   `json:"metadata"`
	Property TimelineObjProperty `json:"property"`
}

type TimelineObjProperty struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
	Stat string `json:"stat"`
}
