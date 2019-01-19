package gen

import "unicode"

func firstRuneToLower(origin string) string {
	return modifyFirstRune(origin, unicode.ToLower)
}

func firstRuneToUpper(origin string) string {
	return modifyFirstRune(origin, unicode.ToUpper)
}

func modifyFirstRune(origin string, f func(rune) rune) string {
	runes := []rune(origin)
	runes[0] = f(runes[0])
	return string(runes)
}
