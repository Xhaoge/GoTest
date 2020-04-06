package config

import (
	"strconv"
	"strings"

	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var profileEnvKey string = "activeProfile"
var profile string
var additionalConfigPathKey string = "additionalConfigPath"
var additionalConfigPath string

func init() {

	viper.SetDefault(profileEnvKey, "")

	viper.BindEnv(profileEnvKey, "ACTIVE_PROFILE")
	viper.BindEnv(additionalConfigPathKey, "ADDITIONAL_CONFIG_PATH")
	profile = viper.GetString(profileEnvKey)
	additionalConfigPath := viper.GetString(additionalConfigPathKey)
	configFile := fmt.Sprintf("application-%s", profile)
	configFile = strings.TrimRight(configFile, "-")

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/applications/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("./internal/configs")
	viper.AddConfigPath("../internal/configs/")
	viper.AddConfigPath("../../internal/configs/")
	viper.AddConfigPath("../../../internal/configs/")
	viper.AddConfigPath("../../../../internal/configs/")
	viper.AddConfigPath("../../../../../internal/configs/")
	if len(additionalConfigPath) > 0 {
		viper.AddConfigPath(additionalConfigPath)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to search config files: %s\n", err))
	}
}

func splitByComma(ce string) []string {
	var result []string
	for _, value := range strings.Split(ce, ",") {
		s := strings.TrimSpace(value)
		if len(s) > 0 {
			result = append(result, s)
		}
	}

	return result
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	val := GetString(key)
	return strings.ToUpper(val) == "TRUE"
}

func GetValueSplitByComma(key string) []string {
	v := GetString(key)
	if len(v) == 0 {
		return []string{}
	}
	return splitByComma(v)
}

func GetInt(key string) (int, error) {
	return strconv.Atoi(GetString(key))
}

func GetLoggingLevel() zapcore.Level {
	level := viper.GetString("logging.level")
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func IsLocal() bool {
	return profile == "" || profile == "local"
}

func IsDevel() bool {
	return profile == "devel"
}

func IsPreProd() bool {
	return profile == "pre-prod"
}

func IsProd() bool {
	return profile == "prod"
}

func GetProfile() string {
	return profile
}

func GetApplicationName() string {
	return GetString("application.name")
}
