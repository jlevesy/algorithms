package strings

import (
	"strconv"
)

// HasUniqueChars implements an algorithm to determine if a string has all unique characters
// An alternative would have been to use an array of bools with a length > 128
// Time Complexity O(n) (n is length of string)
// Space Complexity O(k) (k unique chars in string)
// Question asked:
// - Supported charset (Unicode in that case)
// - Maximum legnth of string (unlimited)
// - Authorized to use external datastruct (Yes)
// - Authorized to use golang map (Yes in that case)
func HasUniqueChars(s string) bool {
	collector := map[rune]struct{}{}

	for _, r := range s {
		if _, ok := collector[r]; ok {
			return false
		}

		collector[r] = struct{}{}
	}

	return true
}

// HasUniqueChars2 solves the  same problem, without additional data structure
// Compare each char with others, return false each time we meet the same character,
// Time complexity O(n²) (n is length of string)
// Space complexity O(1)
// Another option would have been to sort the string (inplace, can be costly)
func HasUniqueChars2(s string) bool {
	for i1, r1 := range s {
		for i2, r2 := range s {
			if i1 == i2 {
				continue
			}

			if r1 == r2 {
				return false
			}
		}
	}

	return true
}

// Returns a count per unique character
// Time complexity: O(n) (n length of string)
func getCharacterCount(s string) map[rune]uint {
	res := map[rune]uint{}

	for _, r := range s {
		res[r]++
	}

	return res
}

// IsPermutationOf shows if one string s1 is a permutation of the other.
// Works with Unicode thanks to Go support
// Time complexity:  O(n) (n length of string)
// Space complexity: O(k) (k unique chars)
// Questions asked:
// - Supported character set (unicode here)
// - Use of external data structures (yes)
func IsPermutationOf(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	c1 := getCharacterCount(s1)
	c2 := getCharacterCount(s2)

	for k, v := range c1 {
		if c2[k] != v {
			return false
		}
	}

	return true
}

// ReplaceSpaces replaces all spaces with %20 chars
// It takes a slice of rune and the size of the original string
// You can assume given string has the length of final string
// And is right blank padded.
// Time complexity O(n) (n lenth of string)
// Space complexity O(1) (in place)
// Questions asked:
// - Should it be done inplace (yes)
// - Character set (a-zA-Z)
func ReplaceSpaces(s []rune, length int) []rune {
	totalLength := len(s)
	writeCursor := 1

	for readCursor := length - 1; readCursor >= 0; readCursor-- {
		if s[readCursor] == ' ' {
			s[totalLength-writeCursor], s[totalLength-writeCursor-1], s[totalLength-writeCursor-2] = '0', '2', '%'
			writeCursor += 3
			continue
		}

		s[totalLength-writeCursor] = s[readCursor]
		writeCursor++
	}

	return s
}

// IsEquilibrated is a which tells if a given string have the same number of opening
// and closing character. It deals with 2 character point smileys like :( or :)
// It can deal with an arbitraty number of opening and closing character.
// Time complexity: O(kn) (n length of string / k total characters to check) k << n => O(n)
// Space complexity: O(1)
// Question(s) asked:
// - Supported chracter set (Unicode)
func IsEquilibrated(s string, chars map[[2]rune]uint) bool {
	for i, r := range s {
		if i > 0 && s[i-1] == ':' {
			continue
		}

		for c := range chars {
			if r == c[0] {
				chars[c]++
			}

			if r == c[1] {
				chars[c]--
			}
		}
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

// Return the expected compression length
// Time complexity O(n)
// Space complexity O(1)
func compressedLength(s []rune) int {
	last := s[0]
	total, count := 0, 1

	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
			continue
		}

		total += 1 + len(strconv.Itoa(count))
		last = s[i]
		count = 1
	}

	return total + 1 + len(strconv.Itoa(count))
}

func writeChunk(s []rune, c rune, writeIndex, count int) int {
	s[writeIndex] = c
	writeIndex++

	countAsString := strconv.Itoa(count)

	for j, r := range countAsString {
		s[writeIndex+j] = r
	}

	return writeIndex + len(countAsString)
}

// CompressString performs basic string compression
// FI: aaaabbddddEEE => a4b2d4E3
// If the compressed string would not become smaller that the original
// the origin should be returned
// string has only upper and lowercase letters (a-zA-Z)
func CompressString(s []rune) []rune {
	if len(s) == 0 {
		return s
	}

	size := compressedLength(s)
	if size >= len(s) {
		return s
	}

	res := make([]rune, size)
	writeIndex, count := 0, 1
	last := s[0]

	for i := 1; i < len(s); i++ {
		if last == s[i] {
			count++
			continue
		}

		writeIndex = writeChunk(res, last, writeIndex, count)

		last = s[i]
		count = 1
	}

	writeChunk(res, last, writeIndex, count)

	return res
}
