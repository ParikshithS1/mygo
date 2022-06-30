package entities

type Users struct {
	UserId        string `json:"userId"`
	UserName      string `json:"userName"`
	Password      string `json:"passWord"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	Email         string `json:"email"`
	// RoleId string `json:"roleId"`
	// CreatedBy string `json:"createdBy"`
	// CreatedDate string `json:"createdDate"`
	// ModifiedBy string `json:"modifiedBy"`
	// ModifiedDate string `json:"modifiedDate"`
	// ActiveStatus string `json:"activeStatus"`

}
