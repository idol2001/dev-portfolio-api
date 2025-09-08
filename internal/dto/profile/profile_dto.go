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

type HomeResp struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type AboutResp struct {
	About       string `json:"about"`
	ImageSource string `json:"imageSource"`
}

type Logo struct {
	Source string `json:"source"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type NavItem struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type NavBarResp struct {
	Logo     Logo      `json:"logo"`
	Sections []NavItem `json:"sections"`
}

type SocialNetwork struct {
	Network string `json:"network"`
	Href    string `json:"href"`
}

type SocialResp struct {
	Social []SocialNetwork `json:"social"`
}
