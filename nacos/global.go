package nacos

var (
	ConfigAll   Config
	ConfigYmAll ConfigYm
	NacosConfig *ConfigNa
	NacosT      *T
)

type T struct {
	Mysql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Library  string `json:"library"`
	} `json:"mysql"`
	Grpc struct {
		Agreement string `json:"agreement"`
		Port      string `json:"port"`
	} `json:"grpc"`
}
type ConfigNa struct {
	NamespaceId string `mapstructure:"namespaceId"`
	IpAddr      string `mapstructure:"ipAddr"`
	Port        int    `mapstructure:"port"`
}
type Config struct {
	Mysql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Library  string `json:"library"`
	} `json:"mysql"`
	Grpc struct {
		Agreement string `json:"agreement"`
		Port      string `json:"port"`
	} `json:"register"`
}
type ConfigYm struct {
	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"mysql"`
}
