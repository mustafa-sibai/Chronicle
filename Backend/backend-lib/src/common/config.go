package common

type ApplicationName string

const (
	ApplicationNameUnknown              ApplicationName = "unknown"
	ApplicationNameRegistrationServer   ApplicationName = "registration-server"
	ApplicationNameAuthenticationServer ApplicationName = "authentication-server"
)

type EnvironmentType string

const (
	EnvironmentTypeDevelopment EnvironmentType = "development"
	EnvironmentTypeStaging     EnvironmentType = "staging"
	EnvironmentTypeProduction  EnvironmentType = "production"
)

type Config struct {
	ApplicationName ApplicationName
	EnvironmentType EnvironmentType
}

var config = Config{
	ApplicationName: ApplicationNameUnknown,
	EnvironmentType: EnvironmentTypeDevelopment,
}

func GetConfig() Config {
	return config
}

func SetConfig(c Config) {
	config = c
}
