package response

type UserLoginInfo struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}

// UserDTO 用户信息
type UserVO struct {
	UserId   int64  `json:"userId"`   //用户id
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
	Role     string `json:"role"`     //角色
	Name     string `json:"name" `    //姓名
	Phone    string `json:"phone"`    //手机号
	Email    string `json:"email"`    //邮箱
}

type UserList struct {
	List []UserVO `json:"list"`
}
