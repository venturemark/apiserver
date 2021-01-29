package schema

type Audience struct {
	Obj AudienceObj `json:"obj"`
}

type AudienceObj struct {
	Metadata map[string]string   `json:"metadata"`
	Property AudienceObjProperty `json:"property"`
}

type AudienceObjProperty struct {
	Name string   `json:"name"`
	Tmln []string `json:"tmln"`
	User []string `json:"user"`
}
