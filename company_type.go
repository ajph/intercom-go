package intercom

type Company_t struct {
	IntercomType     string                 `json:"type,omitempty"`
	Id               string                 `json:"id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Plan             interface{}            `json:"plan,omitempty"`
	CompanyId        string                 `json:"company_id,omitempty"`
	RemoteCreatedAt  int64                  `json:"remote_created_at,omitempty"`
	CreatedAt        int64                  `json:"created_at,omitempty"`
	UpdatedAt        int64                  `json:"updated_at,omitempty"`
	LastRequestAt    int64                  `json:"last_request_at,omitempty"`
	MonthlySpend     uint                   `json:"monthly_spend,omitempty"`
	SessionCount     uint                   `json:"session_count,omitempty"`
	UserCount        uint                   `json:"user_count,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
} //Company_t

type CompanyPlan_t struct {
	IntercomType string `json:"type"`
	Id           uint   `json:"id"`
	Name         string `json:"name"`
} //Plan_t

type CompanyUsers_t struct {
	IntercomType string          `json:"type"`
	TotalCount   uint            `json:"total_count"`
	Users        []CompanyUser_t `json:"users"`
} //CompanyUsers_t

type CompanyUser_t struct {
	IntercomType string `json:"type"`
	Id           string `json:"id"`
} //CompanyUser_t
