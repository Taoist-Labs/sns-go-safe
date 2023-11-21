package sns_go_safe

import api "github.com/Taoist-Labs/sns-go-api"

func IsSafe(name, safeHost string) bool {
	return api.IsSafe(name, safeHost)
}
func Safe(name []string, safeHost string) []string {
	return api.Safe(name, safeHost)
}

func IsAvailable(name, safeHost string) bool {
	return api.IsAvailable(name, safeHost)
}

func Available(name []string, safeHost string) []string {
	return api.Available(name, safeHost)
}
