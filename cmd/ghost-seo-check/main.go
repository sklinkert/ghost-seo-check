package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sklinkert/ghost"
	"github.com/sklinkert/ghost-seo-check/pkg/seo"
	"os"
)

func main() {
	const contentAPIToken = ""
	adminAPIToken := os.Getenv("ADMIN_API_TOKEN")
	if adminAPIToken == "" {
		log.Fatal("Please provide ADMIN_API_TOKEN")
	}
	ghostURL := os.Getenv("GHOST_URL")
	if ghostURL == "" {
		log.Fatal("Please provide GHOST_URL")
	}
	ghostAPI, err := ghost.New(ghostURL, contentAPIToken, adminAPIToken)
	if err != nil {
		log.WithError(err).Fatal("ghost.New() failed")
	}
	posts, err := ghostAPI.AdminGetPosts()
	if err != nil {
		log.WithError(err).Fatal("Fetching posts failed")
	}

	var pagesWithErrors = 0
	for _, post := range posts.Posts {
		var seoPage = seo.Page{
			Title:           post.Title,
			Text:            post.MobileDoc,
			MetaDescription: post.MetaDescription,
			MetaTitle:       post.MetaTitle,
			Excerpt:         post.Excerpt,
			FeatureImage:    post.FeatureImage,
		}

		seoErrors := seo.CheckPost(seoPage)
		if len(seoErrors) > 0 {
			fmt.Printf("%s -> %q\n", post.URL, post.Title)
			for _, err := range seoErrors {
				fmt.Printf("\t%s\n", err.Error())
			}
			pagesWithErrors++
		}
	}

	log.Infof("%d of %d pages have SEO errors", pagesWithErrors, len(posts.Posts))
}
