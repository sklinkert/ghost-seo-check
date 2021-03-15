package seo

import (
	"errors"
	"fmt"
)

type CheckFunc func(page *Page) error

func checkTitle(page *Page) error {
	const minLength = 10
	const maxLength = 170
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
	if page.MetaDescription == "" {
		return errors.New("meta description is empty")
	}
	return nil
}

func checkFeatureImage(page *Page) error {
	if page.FeatureImage == "" {
		return errors.New("feature image is missing")
	}
	return nil
}
