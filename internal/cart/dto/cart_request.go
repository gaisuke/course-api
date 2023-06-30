package cart

type CartRequestBody struct {
	ProductID int64 `json:"product_id" binding:"required,number"`
	UserID    int64 `json:"user_id"`
	CreatedBy int64
	UpdatedBy int64
}

type CartRequestUpdateBody struct {
	IsChecked bool   `json:"is_checked"`
	UserID    *int64 `json:"user_id"`
}
