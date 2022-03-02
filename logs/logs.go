package logs

import "unicode/utf8"

func Application(log string) string {
	reference := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}

	for _, val := range log {
		app, valExists := reference[val]
		if valExists {
			return app
		}
	}
	return "default"
}
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, val := range log {
		if val == oldRune {
			val = newRune
		}
		newLog += string(val)
	}
	return newLog
}
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
