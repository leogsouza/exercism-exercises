// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith receives one string parameter. if string is empty who should be replaced by "you"
func ShareWith(who string) string {

	if who == "" {
		who = "you"
	}
	return "One for " + who + ", one for me."
}
