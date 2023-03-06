package domain

import "regexp"

func GetMentions(text string) []string {
	/*
		The regular expression @\w{1,30} matches a string that
		starts with "@" followed by 1 to 30 word characters,
		where a word character is any letter, digit, or underscore.
		This ensures that usernames do not include symbols or punctuation marks.
	*/
	pattern := `@(\w{1,30})`

	// Compile regular expression pattern into a regular expression object
	regex := regexp.MustCompile(pattern)

	var mentions []string
	for _, match := range regex.FindAllStringSubmatch(text, -1) {
		mentions = append(mentions, match[1])
	}

	return mentions
}
