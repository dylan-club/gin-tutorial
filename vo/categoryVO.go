package vo

type CategoryVO struct {
	Name string `json:"name" binding:"required"`
}
