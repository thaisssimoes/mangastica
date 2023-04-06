package files

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func GetFileList(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func regexGetBeforeExpression(expression, delimiter string) (string, error) {
	cleanExpression := regexp.MustCompile(fmt.Sprintf((`^[^%s]+`), delimiter)).FindString(expression)
	if cleanExpression == "" {
		return "", fmt.Errorf("Uncapable of finding expression")
	}
	cleanExpression = strings.TrimRight(cleanExpression, "1234567890 ")

	return cleanExpression, nil
}

func DigitPrefix(s string) string {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return s[:i]
		}
	}
	return s
}
func DigitBetween(s string) string {
	for i, r := range s {
		if unicode.IsDigit(r) {
			splitStr := strings.Split(s[i:], " ")
			return splitStr[0]
		}
	}
	return s
}

func regexGetBetweenExpression(leftDelimiter, rightDelimiter, str string) (string, error) {
	r := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(leftDelimiter) + `(.*?)` + regexp.QuoteMeta(rightDelimiter))
	matches := r.FindAllStringSubmatch(str, -1)
	for _, v := range matches {
		return v[1], nil
	}
	return "", fmt.Errorf("Uncapable of finding expression")
}

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
