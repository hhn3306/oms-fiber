package request

type Login struct {
	Username  string `json:"username" validate:"required" ` // 用户名
	Password  string `json:"password" validate:"required"`  // 密码
	Captcha   string `json:"captcha"`                       // 验证码
	CaptchaId string `json:"captchaId"`                     // 验证码ID
}
