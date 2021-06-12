package config

var ServerConfig = ServerConfiguration{}

type ServerConfiguration struct {
	Host          string                   `mapstructure:"host"`
	Port          int                      `mapstructure:"port"`
	ServiceConfig ServiceNameConfiguration `mapstructure:"svc_name"`
	LogConfig     LogConfiguration         `mapstructure:"log"`
	ConsulConfig  ConsulConfiguration      `mapstructure:"consul"`
	JWTConfig     JWTConfiguration         `mapstructure:"jwt"`
}
type ServiceNameConfiguration struct {
	UserServiceName      string `mapstructure:"user_svc_name"`
	EmailServiceName     string `mapstructure:"email_svc_name"`
	GoodsServiceName     string `mapstructure:"goods_svc_name"`
	InventoryServiceName string `mapstructure:"inventory_svc_name"`
	OrderServiceName     string `mapstructure:"order_svc_name"`
	PaymentServiceName   string `mapstructure:"payment_svc_name"`
}
type LogConfiguration struct {
	LogPath string `mapstructure:"log_path"`
}
type ConsulConfiguration struct {
	Name  string                   `mapstructure:"name"`
	Host  string                   `mapstructure:"host"`
	Port  int                      `mapstructure:"port"`
	Id    string                   `mapstructure:"id"`
	Tags  []string                 `mapstructure:"tags"`
	Url   string                   `mapstructure:"url"`
	Check ConsulCheckConfiguration `mapstructure:"check"`
}
type ConsulCheckConfiguration struct {
	CheckMethod string `mapstructure:"check_method"`
	Method      string `mapstructure:"method"`
	Interval    string `mapstructure:"interval"`
	Uri         string `mapstructure:"uri"`
}
type JWTConfiguration struct {
	Secret   string `mapstructure:"secret"`
	ExpireAt int64  `mapstructure:"expire_at"`
	Issuer   string `mapstructure:"issuer"`
}
