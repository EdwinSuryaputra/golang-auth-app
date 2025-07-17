package config

type config struct {
	Application     applicationConfig     `mapstructure:"application"`
	Datasource      datasourceConfig      `mapstructure:"datasource"`
	Route           routeConfig           `mapstructure:"route"`
	InternalService internalServiceConfig `mapstructure:"internal_service"`
	ExternalService externalServiceConfig `mapstructure:"external_service"`
	Module          moduleConfig          `mapstructure:"module"`
	Monitoring      monitoringConfig      `mapstructure:"monitoring"`
}

type applicationConfig struct {
	Name string      `mapstructure:"name"`
	Port int         `mapstructure:"port"`
	Env  Environment `mapstructure:"env"`
	Tls  tlsConfig   `mapstructure:"tls"`
	Cors corsConfig  `mapstructure:"cors"`
}

type tlsConfig struct {
	IsEnabled bool   `mapstructure:"is_enabled"`
	CertFile  string `mapstructure:"cert_file"`
	KeyFile   string `mapstructure:"key_file"`
}

type corsConfig struct {
	Origins          string `mapstructure:"origins"`
	Methods          string `mapstructure:"methods"`
	Headers          string `mapstructure:"headers"`
	AllowCredentials bool   `mapstructure:"allow_credentials"`
}

type datasourceConfig struct {
	Postgres postgresConfig `mapstructure:"postgres"`
	Redis    redisConfig    `mapstructure:"redis"`
}

type postgresConfig struct {
	Host          string `mapstructure:"host"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Port          int    `mapstructure:"port"`
	DBName        string `mapstructure:"database"`
	Retry         int    `mapstructure:"retry"`
	SSLMode       string `mapstructure:"sslmode"`
	PoolMaxConns  int    `mapstructure:"pool_max_conns"`
	RetryInterval string `mapstructure:"retry_interval"`
}

type redisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type routeConfig struct {
}

type internalServiceConfig struct {
	Medio internalServiceMedioConfig `mapstructure:"medio"`
}

type internalServiceMedioConfig struct {
	Host   string                           `mapstructure:"host"`
	Routes internalServiceMedioRoutesConfig `mapstructure:"routes"`
}

type internalServiceMedioRoutesConfig struct {
	GetList        string `mapstructure:"get_list"`
	GetActivityLog string `mapstructure:"get_activity_log"`
}

type externalServiceConfig struct {
	Smtp smtpConfig `mapstructure:"smtp"`
}

type smtpConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type moduleConfig struct {
	Auth authModuleConfig `mapstructure:"auth"`
	User userModuleConfig `mapstructure:"user"`
	Role roleModuleConfig `mapstructure:"role"`
}

type authModuleConfig struct {
	Jwt        authModuleJwtConfig        `mapstructure:"jwt"`
	WebPageUrl authModuleWebPageUrlConfig `mapstructure:"web_page_url"`
}

type authModuleJwtConfig struct {
	SecretKey                     string `mapstructure:"secret_key"`
	TokenExpiryInSeconds          int    `mapstructure:"token_expiry_in_seconds"`
	IsForeverLoginPersistForever  bool   `mapstructure:"is_forever_login_persist_forever"`
	ForeverLoginDurationInSeconds int    `mapstructure:"forever_login_duration_in_seconds"`
}

type authModuleWebPageUrlConfig struct {
	LoginUrl         string `mapstructure:"login_url"`
	ResetPasswordUrl string `mapstructure:"reset_password_url"`
}

type userModuleConfig struct {
	PayloadValidation userModulePayloadValidationConfig `mapstructure:"payload_validation"`
}

type userModulePayloadValidationConfig struct {
	UsernameMinDigit    int `mapstructure:"username_min_digit"`
	PasswordMinDigit    int `mapstructure:"password_min_digit"`
	FullNameMinDigit    int `mapstructure:"fullname_min_digit"`
	EmailMinDigit       int `mapstructure:"email_min_digit"`
	PhoneNumberMinDigit int `mapstructure:"phone_number_min_digit"`
	DescriptionMinDigit int `mapstructure:"description_min_digit"`
}

type roleModuleConfig struct {
	PayloadValidation roleModulePayloadValidationConfig `mapstructure:"payload_validation"`
}

type roleModulePayloadValidationConfig struct {
	NameMinDigit        int `mapstructure:"name_min_digit"`
	DescriptionMinDigit int `mapstructure:"description_min_digit"`
}

type monitoringConfig struct {
}
