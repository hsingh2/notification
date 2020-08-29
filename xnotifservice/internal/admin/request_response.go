package admin

//HealthAdminTemplateResponse ...
type HealthAdminTemplateResponse struct {
	Status string `json:"status"`
	Err    error  `json:"error,omitempty"`
}

//TokenAdminTemplateResponse ...
type TokenAdminTemplateResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}
