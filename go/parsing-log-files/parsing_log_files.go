package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\*|~|-|=)*>+`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\"(?i).*(password)+\"+`)
	c := 0
	for _, line := range lines {
		if re.MatchString(line) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d+)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+)(\w+)`)
	logs := []string{}
	for _, line := range lines {
		lg := line
		if re.MatchString(line) {
			lns := re.FindStringSubmatch(line)
			lg = fmt.Sprintf("[USR] %s %s", lns[2], line)
		}
		logs = append(logs, lg)
	}
	return logs
}
