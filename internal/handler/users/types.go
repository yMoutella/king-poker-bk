package users

type GetUserParam struct {
	ID string `uri:"id" binding:"required"`
}

type CreateUserParam struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type GetAllUsersParam struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}
