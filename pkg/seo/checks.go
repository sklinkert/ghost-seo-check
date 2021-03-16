package seo

import (
	"errors"
	"fmt"
	"strings"
)

type CheckFunc func(page *Page) error

func checkTitle(page *Page) error {
	// https://moz.com/learn/seo/title-tag
	const minLength = 20
	const maxLength = 60

	if page.Title == "" {
		return errors.New("title is missing")
	}
	if len(page.Title) > maxLength {
		return fmt.Errorf("title is too long (%d > %d)", len(page.Title), maxLength)
	} else if len(page.Title) < minLength {
		return fmt.Errorf("title is too short (%d < %d)", len(page.Title), minLength)
	} else {
		return nil
	}
}

func checkExcerpt(page *Page) error {
	if page.Excerpt == "" {
		return errors.New("excerpt is empty")
	}
	return nil
}

func checkMetaDescription(page *Page) error {
	// https://moz.com/learn/seo/meta-description
	const minLength = 50
	const maxLength = 160

	if page.MetaDescription == "" {
		return errors.New("meta description is empty")
	}
	if len(page.MetaDescription) > maxLength {
		return fmt.Errorf("meta description is too long (%d > %d)", len(page.MetaDescription), maxLength)
	} else if len(page.MetaDescription) < minLength {
		return fmt.Errorf("meta description is too short (%d < %d)", len(page.MetaDescription), minLength)
	} else {
		return nil
	}
}

func checkFeatureImage(page *Page) error {
	if page.FeatureImage == "" {
		return errors.New("feature image is missing")
	}
	return nil
}

func checkText(page *Page) error {
	if page.Text == "" {
		return errors.New("text is missing")
	}

	const minWordCount = 300
	words := wordCount(page.Text)
	if words < minWordCount {
		return fmt.Errorf("text is too short (%d < %d)", words, minWordCount)
	}
	return nil
}

func checkTags(page *Page) error {
	if len(page.Tags) == 0 {
		return errors.New("no tags set")
	}
	return nil
}

func wordCount(s string) int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return len(m)
}
