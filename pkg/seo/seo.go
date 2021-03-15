package seo

import (
	"strings"
)

type Page struct {
	Title           string
	Text            string
	Excerpt         string
	MetaDescription string
	MetaTitle       string
	FeatureImage    string
}

func wordCount(s string) int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return len(m)
}

func CheckPost(page Page) (seoErrors []error) {
	var checks = []CheckFunc{
		checkTitle,
		checkExcerpt,
		checkMetaDescription,
		checkFeatureImage,
	}
	for _, check := range checks {
		var seoError = check(&page)
		if seoError != nil {
			seoErrors = append(seoErrors, seoError)
		}
	}
	return seoErrors
}
