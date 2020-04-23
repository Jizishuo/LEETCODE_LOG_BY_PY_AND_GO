package Spelling_words_1160

func countCharacters(words []string, chars string) int {
	res := 0

	out:for _, word := range words {
		for _, w := range word {
			if strings.Count(word, string(w)) > strings.Count(chars, string(w)) {

				continue out
			}

		}
		res += len(word)
	}

	return res
}