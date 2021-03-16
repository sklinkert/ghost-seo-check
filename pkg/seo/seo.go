package seo

type Page struct {
	Title           string
	Text            string
	Excerpt         string
	MetaDescription string
	MetaTitle       string
	FeatureImage    string
	Tags            []string
}

func CheckPost(page Page) (seoErrors []error) {
	var checks = []CheckFunc{
		checkTitle,
		checkExcerpt,
		checkMetaDescription,
		checkFeatureImage,
		checkText,
		checkTags,
	}
	for _, check := range checks {
		var seoError = check(&page)
		if seoError != nil {
			seoErrors = append(seoErrors, seoError)
		}
	}
	return seoErrors
}
