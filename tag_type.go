package intercom

type TagList_t struct {
	IntercomType string      `json:"type,omitempty"`
	Tags         []Tag_t     `json:"tags,omitempty"`
	Pages        interface{} `json:"pages,omitempty"`
} //TagList_t

type Tag_t struct {
	IntercomType string         `json:"type,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Users        []TagUser_t    `json:"users,omitempty"`
	Companies    []TagCompany_t `json:"companies,omitempty"`
} //Tag_t

type TagUser_t struct {
	Id     string `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
	Untag  bool   `json:"untag,omitempty"`
} //TagUser_t

type TagCompany_t struct {
	Id        string `json:"id,omitempty"`
	CompanyId string `json:"company_id,omitempty"`
	Untag     bool   `json:"untag,omitempty"`
} //TagCompany_t
