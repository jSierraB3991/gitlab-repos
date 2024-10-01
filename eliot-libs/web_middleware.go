package eliotlibs

import "strings"

func PublicMiddleWare(route string, method string) bool {

	var NO_AUTH_NEED = []string{"public"}

	if method == "OPTIONS" {
		return true
	}

	for _, value := range NO_AUTH_NEED {
		if strings.Contains(route, value) {
			return true
		}
	}
	return false
}
