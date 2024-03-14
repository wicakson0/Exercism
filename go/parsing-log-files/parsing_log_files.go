package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}
func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[-=~*]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)
	count := 0

	for _, v := range lines {
		if re.MatchString(v) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+(\w+)`)

	var taggedUser []string

	for _, v := range lines {
		containsUser := re.FindStringSubmatch(v)

		if containsUser != nil {
			v = fmt.Sprintf("[USR] %s %s", containsUser[1], v)
		}
		taggedUser = append(taggedUser, v)
	}

	return taggedUser
}
