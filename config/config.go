package config

type Server struct {
	Zap     Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System          `mapstructure:"system" json:"system" yaml:"system"`
	JWT     JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha Captcha         `mapstructure:"captcha" json:"captcha" yaml:"captcha"` //验证码
	Redis   Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql   Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList  []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
