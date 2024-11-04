package request

// UserLogin 用户登录
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserEditPassword 用户修改密码
type UserEditPassword struct {
	UserId      int64  `json:"userId"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

// UserDTO 用户信息
type UserDTO struct {
	UserId   int64  `json:"userId"`    //用户id
	Username string `json:"username" ` //用户名
	Password string `json:"password"`  //密码
	Role     string `json:"role"`      //角色
	Name     string `json:"name" `     //姓名
	Phone    string `json:"phone"`     //手机号
	Email    string `json:"email"`     //邮箱
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserID struct {
	UserId int64 `json:"userId" binding:"required"`
}

type Username struct {
	Username string `json:"username" binding:"required"`
}
