package user

type UserInfoRequest struct {
	Name string `json:"name" form:"name"`
}
