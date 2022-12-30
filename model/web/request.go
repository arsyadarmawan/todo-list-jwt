package web

type TaskCreateRequest struct {
	User_Id    int `validate:"omitempty" json:"user_id"`
	Product_Id int `validate:"omitempty" json:"product_id"`
	Total      int `validate:"required" json:"total"`
}

type TaskUpdateRequest struct {
	Id    int `validate:"required"`
	Total int `validate:"required" json:"total"`
}
