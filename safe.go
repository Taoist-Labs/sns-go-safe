package sns_go_safe

import api "github.com/Taoist-Labs/sns-go-api"

func IsSafe(name, apiUrl string) bool {
	return api.IsSafe(name, apiUrl)
}
func Safe(name []string, apiUrl string) []string {
	return api.Safe(name, apiUrl)
}

func IsAvailable(name, apiUrl string) bool {
	return api.IsAvailable(name, apiUrl)
}

func Available(name []string, apiUrl string) []string {
	return api.Available(name, apiUrl)
}
