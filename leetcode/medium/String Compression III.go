package medium

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	var simbArr []rune
	var counts []int

	if word == "" {
		return ""
	}

	simbArr = append(simbArr, rune(word[0]))
	counts = append(counts, 1)

	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			if counts[len(counts)-1] == 9 {
				simbArr = append(simbArr, rune(word[i]))
				counts = append(counts, 1)
				continue
			}
			counts[len(counts)-1]++
		} else {
			simbArr = append(simbArr, rune(word[i]))
			counts = append(counts, 1)
		}
	}

	var str []string
	for i := 0; i < len(simbArr); i++ {
		str = append(str, fmt.Sprintf("%d", counts[i]))
		str = append(str, string(simbArr[i]))
	}

	result := strings.Join(str, "")

	return result
}
