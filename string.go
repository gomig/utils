package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ExtractNumbers extract numbers from string
func ExtractNumbers(str string) string {
	rx := regexp.MustCompile("[0-9]+")
	return strings.Join(rx.FindAllString(str, -1), "")
}

// ExtractAlphaNum extract alpha and numbers from string [a-zA-Z0-9]
func ExtractAlphaNum(str string, includes ...string) string {
	rx := regexp.MustCompile(fmt.Sprintf("[^a-zA-Z0-9%s]", strings.Join(includes, "")))
	return rx.ReplaceAllString(str, "")
}

// ExtractAlphaNumPersian extract persian alpha, alpha and numbers from string [ا-یa-zA-Z0-9]
func ExtractAlphaNumPersian(str string, includes ...string) string {
	rx := regexp.MustCompile(fmt.Sprintf("[^\u0600-\u06FF\uFB8A\u067E\u0686\u06AFa-zA-Z0-9%s]", strings.Join(includes, "")))
	return rx.ReplaceAllString(str, "")
}

// RandomStringFromCharset generate random string from character list
func RandomStringFromCharset(n uint, letters string) (res string, err error) {
	randomer := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n)
	_, err = randomer.Read(bytes)
	if err != nil {
		return
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	res = string(bytes)
	return
}

// RandomString generate random string from Alpha-Num Chars
func RandomString(n uint) (string, error) {
	return RandomStringFromCharset(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
}

// Slugify make slugify string
func Slugify(str ...string) string {
	r := regexp.MustCompile(`\s+`)
	r2 := regexp.MustCompile(`\-+`)
	return string(r2.ReplaceAllString(string(r.ReplaceAllString(strings.Join(str, "-"), "-")), "-"))
}

// SlugifyPersian make slugify string for persian string
func SlugifyPersian(str ...string) string {
	rx := regexp.MustCompile("[^\u0600-\u06FF\uFB8A\u067E\u0686\u06AF a-zA-Z0-9\\-]") // standard chars
	return Slugify(rx.ReplaceAllString(strings.Join(str, "-"), ""))
}

// ConcatStr join strings with separator
func ConcatStr(sep string, str ...string) string {
	res := make([]string, 0)
	for _, v := range str {
		if strings.TrimSpace(v) != "" {
			res = append(res, v)
		}
	}
	return strings.Join(res, sep)
}

// FormatNumber format number with comma separator
func FormatNumber(format string, v ...any) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf(format, v...)
}

// FormatRx format string using regex pattern
// use () for group and $1, $2 for output placeholder
// example FormatRx("123456", `^(\d{3})(\d{2})(\d{1})$`, "($1) $2-$3")
func FormatRx(data, pattern, repl string) string {
	rx := regexp.MustCompile(pattern)
	return rx.ReplaceAllString(data, repl)
}
