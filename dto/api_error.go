package dto

type ApiError struct {
	Status uint   `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (v ApiError) Error() string {
	return v.Detail
}
