package email

type SendContactRequest struct {
	FirstName   string `json:"firstName" binding:"required,min=2"`
	LastName    string `json:"lastName" binding:"required,min=2"`
	Email       string `json:"email" binding:"required,email"`
	CompanyName string `json:"companyName" binding:"omitempty,max=100"`
	ProjectType string `json:"projectType" binding:"required,oneof=new-website e-commerce redesign web-app optimization other"`
	Budget      string `json:"budget" binding:"required,oneof=under-50k 50k-100k 100k-200k 200k-plus"`
	Message     string `json:"message" binding:"required,min=10"`
}

type SendResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
