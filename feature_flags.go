package featureflags

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type FeatureFlags struct {
	Prefix string
}

var instance *FeatureFlags
var once sync.Once

func GetInstance() *FeatureFlags {
	once.Do(func() {
		instance = &FeatureFlags{
			Prefix: "FeatureFlag",
		}
	})
	return instance
}

func (featureFlags *FeatureFlags) IsEnabled(flag string) bool {
	featureFlagKey := fmt.Sprintf("%s_%s", featureFlags.Prefix, flag)
	envValue := os.Getenv(featureFlagKey)
	if envValue == "" {
		return false
	}

	parsed, err := strconv.ParseBool(os.Getenv(featureFlagKey))

	if err != nil {
		return false
	}
	return parsed
}
