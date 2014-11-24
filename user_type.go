package intercom

type User_t struct {
	IntercomType           string                 `json:"type"`
	Id                     string                 `json:"id"`
	CreatedAt              int64                  `json:"created_at"`
	RemoteCreatedAt        int64                  `json:"remote_created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	UserId                 string                 `json:"user_id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	LastRequestAt          int64                  `json:"last_request_at,omitempty"`
	SessionCount           uint                   `json:"session_count,omitempty"`
	Avatar                 Avatar_t               `json:"avatar,omitempty"`
	UnsubscribedFromEmails bool                   `json:"unsubscribed_from_emails,omitempty"`
	LocationData           Location_t             `json:"location_data,omitempty"`
	UserAgentData          string                 `json:"user_agent_data,omitempty"`
	LastSeenIP             string                 `json:"last_seen_ip,omitempty"`
	Companies              []string               `json"companies,omitempty"`
} //User_t

type Avatar_t struct {
	IntercomType string `json:"type"`
	ImageUrl     string `json:"image_url"`
} //Avatar_t

type Location_t struct {
	IntercomType  string `json:"type"`
	CityName      string `json:"city_name,omitempty"`
	ContinentCode string `json:"continent_code,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	CountryName   string `json:"country_name,omitempty"`
	Latitude      uint   `json:"latitude,omitempty"`
	Longitude     uint   `json:"longitude,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	RegionName    string `json:"region_name,omitempty"`
	Timezone      string `json:"timezone,omitempty"`
} //Location_t

type SocialProfile_t struct {
	IntercomType string `json:"type"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Url          string `json:"url"`
	Id           string `json:"id"`
} //SocialProfile_t

type UserList_t struct {
	IntercomType string   `json:"type"`
	TotalCount   uint     `json:"total_count"`
	Pages        Pages_t  `json:"pages"`
	Users        []User_t `json"users"`
} //UserList_t

type Pages_t struct {
	Next       string `json:"next,omitempty"`
	Page       uint   `json:"page,omitempty"`
	PerPage    uint   `json:"per_page,omitempty"`
	TotalPages uint   `json:"total_pages,omitempty"`
} //Pages_t
