package request

type CreateHistoryRequest struct{
	Name	string	`validate:"required,min=1,max=200" json:"name"`
	UserID	int		`validate:"required" json:"user_id"`
	History	string	`valdiate:"required,min=1,max=1000" json:"history"`
}

