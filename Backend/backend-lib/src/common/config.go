package common

type ApplicationName string

const (
	ApplicationName_Unknown               ApplicationName = "unknown"
	ApplicationName_RegistrationService   ApplicationName = "registration-service"
	ApplicationName_AuthenticationService ApplicationName = "authentication-service"
	ApplicationName_SessionService        ApplicationName = "session-service"
	ApplicationName_AccountService        ApplicationName = "account-service"
	ApplicationName_CharacterService      ApplicationName = "character-service"
	ApplicationName_ContentService        ApplicationName = "content-service"
)

type EnvironmentType string

const (
	EnvironmentType_Development EnvironmentType = "development"
	EnvironmentType_Staging     EnvironmentType = "staging"
	EnvironmentType_Production  EnvironmentType = "production"
)

type Config struct {
	ApplicationName ApplicationName
	EnvironmentType EnvironmentType
}

var config = Config{
	ApplicationName: ApplicationName_Unknown,
	EnvironmentType: EnvironmentType_Development,
}

func GetConfig() Config {
	return config
}

func SetConfig(c Config) {
	config = c
}
