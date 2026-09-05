package profile

type SkillItem struct {
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type SkillGroup struct {
	Title string      `json:"title"`
	Items []SkillItem `json:"items"`
}

type SkillsResp struct {
	Intro  string       `json:"intro"`
	Skills []SkillGroup `json:"skills"`
}
