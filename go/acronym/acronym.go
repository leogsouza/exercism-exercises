// Acronym for a string
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	re := regexp.MustCompile(`\b(\w)`)
	str := strings.Join(re.FindAllString(s, -1), "")
	return strings.ToUpper(str)
}
