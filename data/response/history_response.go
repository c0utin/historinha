package response

type HistoryResponse struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
	History string	`json:"history"`
	UserID	int	`json:"user_id"`
}
