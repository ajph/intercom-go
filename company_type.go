package intercom

type ExistingCompany_t struct {
	Type             string                 `json:"type"`
	Id               uint                   `json:"id"`
	Name             string                 `json:"name"`
	Plan             Plan_t                 `json:"plan"`
	CompanyId        uint                   `json:"company_id"`
	RemoteCreatedAt  int64                  `json:"remote_created_at"`
	CreatedAt        int64                  `json:"created_at"`
	UpdatedAt        int64                  `json:"updated_at"`
	LastRequestAt    int64                  `json:"last_request_at"`
	MonthlySpend     uint                   `json:"monthly_spend"`
	SessionCount     uint                   `json:"session_count"`
	UserCount        uint                   `json:"user_count"`
	CustomAttributes map[string]interface{} `json:"custom_attributes"`
} //ExistingCompany_t

type NewCompany_t struct {
	Name             string                 `json:"name"`
	Plan             string                 `json:"plan"`
	CompanyId        uint                   `json:"company_id"`
	RemoteCreatedAt  int64                  `json:"remote_created_at"`
	MonthlySpend     uint                   `json:"monthly_spend"`
	CustomAttributes map[string]interface{} `json:"custom_attributes"`
} //NewCompany_t

type Plan_t struct {
	Type string `json:"type"`
	Id   uint   `json:"id"`
	Name string `json:"name"`
} //Plan_t
