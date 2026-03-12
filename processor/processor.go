package processor

import (
	"fmt"
	"strconv"
	"strings"
)

func DecimalConversion(words []string) []string {

	for i := 0; i < len(words); i++ {
		match := words[i]
		// for hexadecimal conversion
		if strings.Contains(match, "(hex)") && i > 0 {
			// convert the previous word to decimal
			if val, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(val, 10)
			}

			//Remove the marker "(hex)" from current words
			clean := strings.ReplaceAll(match, "(hex)", "")
			if clean != "" {
				words[i] = clean
			} else {
				words = RemoveWords(words, i)
				i--
			}
		} else if strings.Contains(match, "(bin)") && i > 0 {
			if val, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(val, 10)
				// Remove the "(bin)" from the slice and adjust the index (before the index of "(bin)" and after the index of "(bin)")
			}
			//Remove the marker "(bin)" from current words
			clean := strings.ReplaceAll(match, "(bin)", "")
			if clean != "" {
				words[i] = clean
			} else {
				words = RemoveWords(words, i)
				i--
			}
		}
	}
	return words

}

func CaseConversion(words []string) []string {

	for i := 0; i < len(words); i++ {
		match := words[i]
		// strings.Tolower, Toupper
		if match == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = RemoveWords(words, i)
			i--
		}

		// change to lowercase
		if match == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = RemoveWords(words, i)
			i--
		}

		// // for capitalize
		if match == "(cap)" && i > 0 {
			// Capitalize the first letter and make the rest lowercase
			if len(words[i-1]) > 0 {
				words[i-1] = strings.ToUpper(words[i-1][0:1]) + strings.ToLower(words[i-1][1:])
			}
			words = RemoveWords(words, i)
			i--
		}

		// for (up, N), (low, N) and (cap, N)
		if match == "(up," || match == "(low," || match == "(cap," {

			//ensure that the words exist
			if i+1 >= len(words) {
				continue
			}
			ndx := strings.TrimSuffix(words[i+1], ")") // Remove the closing parenthesis from the index
			ndxInt, err := strconv.Atoi(ndx)           // Convert the index from string to integer
			if err != nil {
				fmt.Printf("Error converting %s to integer: %v\n", ndx, err)
				continue
			}
			if i-ndxInt < 0 {
				fmt.Printf("Error: Not enough words to convert for %s\n", words[i])
				continue
			}

			//loop for multi-words
			for j := i - ndxInt; j < i; j++ {
				// Use a switch statement to determine the case conversion based on the match
				switch match {
				case "(up,":
					words[j] = strings.ToUpper(words[j])
				case "(low,":
					words[j] = strings.ToLower(words[j])
				case "(cap,":
					words[j] = strings.ToUpper(words[j][0:1]) + strings.ToLower(words[j][1:])
				}
			}
			words = RemoveWords(words, i) // Remove the marker "(up,", "(low," or "(cap," from the slice
			words = RemoveWords(words, i) // Remove the  "N)" from the slice (after removing the marker, the index will be at the same position)
			i--
		}

	}
	return words
}

func RemoveWords(words []string, index int) []string {
	result := append(words[:index], words[index+1:]...)
	return result

}

func isPunct(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';':
		return true
	}
	return false
}

func PunctuationHandler(text string) string {
	runes := []rune(text)
	var b strings.Builder
	var last rune // Declare last to match space ' ' in rune to handle leading punctuation

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if isPunct(r) {
			if last == ' ' {
				out := b.String()
				b.Reset()
				b.WriteString(out[:len(out)-1]) // Remove the trailing space before punctuation
			}

			for i < len(runes) && isPunct(runes[i]) { // Handle consecutive punctuation
				b.WriteRune(runes[i])
				last = runes[i]
				i++
			}

			if i < len(runes) && runes[i] != ' ' { // Add a space after punctuation if the next character is not a space
				b.WriteRune(' ')
				last = ' '
			}

			i--
			continue // Skip the rest of the loop to avoid adding punctuation again
		}

		if r == ' ' && last == ' ' { // Handle multiple spaces
			continue
		}

		b.WriteRune(r)
		last = r
	}

	return b.String()
}

func QuoteHandler(s string) string {

	//  Fixes Quotes Through This Process: "' word '" -> "'word'"
	parts := strings.Split(s, "'")
	for i := 1; i < len(parts); i += 2 {
		parts[i] = strings.TrimSpace(parts[i]) // Strip spaces inside quotes
	}
	s = strings.Join(parts, "'")

	return strings.Join(strings.Fields(s), " ")
}

func VowelHandler(words []string) []string {

	for i, word := range words {
		// Check if the current word is "a" or "A"
		if (word == "a" || word == "A") && i+1 < len(words) {
			// Look at the first character of the next word
			nextWord := words[i+1]
			firstRune := rune(nextWord[0])

			if IsVowelOrHLetter(firstRune) {
				if word == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return words
}

func IsVowelOrHLetter(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	}
	return false
}
