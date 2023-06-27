package utils

type VerifyPhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
//func VerifyCodePhone(data interface{}, e *echo.Context) map[string][]string {
//	// 1. 制定认证规则
//
//	rules := govalidator.MapData{
//		"phone":          []string{"required", "digits:11"},
//		"captcha_id":     []string{"required"},
//		"captcha_answer": []string{"required", "digits:6"},
//	}
//
//	// 定制错误信息
//	messages := govalidator.MapData{
//		"phone": []string{
//			"required:手机号为必填项，参数名称phone",
//			"digits:手机号长度必须为11位数字",
//		},
//		"captcha_id": []string{
//			"required:图片验证码的ID为必填",
//		},
//		"captcha_answer": []string{
//			"required:图片验证码答案为必填",
//			"digits:图片验证码长度必须为6位的数字",
//		},
//	}
//
//	errs := validate
//
//}
