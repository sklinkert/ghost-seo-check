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
		//if post.MetaDescription == ""	{
		//	if strings.Contains(post.Title, "Woche") {
		//		continue
		//	}
		//		post.MetaDescription = strings.Title(post.Tags[0].Name) + ""
		//		log.Infof("New: %s",post.MetaDescription)
		//
		//		updatedPost := ghost.Post {
		//			ID: post.ID,
		//			MetaDescription: post.MetaDescription,
		//			UpdatedAt: post.UpdatedAt,
		//			MobileDoc: post.MobileDoc,
		//			PublishedAt: post.PublishedAt,
		//		}
		//		pagesWithErrors++
		//		if err := ghostAPI.AdminUpdatePost(updatedPost); err != nil {
		//			log.WithError(err).Fatal("update")
		//	}
		//}

		var seoPage = seo.Page{
			Title:           post.Title,
			Text:            post.MobileDoc,
			MetaDescription: post.MetaDescription,
			MetaTitle:       post.MetaTitle,
			Excerpt:         post.Excerpt,
			FeatureImage:    post.FeatureImage,
		}

		for _, tag := range post.Tags {
			seoPage.Tags = append(seoPage.Tags, tag.Name)
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

	percent := 100.0 / float64(len(posts.Posts)) * float64(pagesWithErrors)
	log.Infof("%d (%.2f%%) out of %d pages have SEO errors", pagesWithErrors, percent, len(posts.Posts))
}
