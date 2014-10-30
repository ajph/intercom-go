package intercom

func NewIntercom(appId, apiKey string) *Intercom_t {
	return &Intercom_t{
		appId,
		apiKey,
	} //Intercom_t
} //NewIntercom
