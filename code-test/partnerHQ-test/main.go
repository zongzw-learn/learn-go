package main

import (
	"flag"
	"log"
	"regexp"
)

// question link:
// 	https://manifold-core.notion.site/PartnerHQ-Technical-Screen-7baf1e34c72b44d3aa98a63eb6d83f7c

var (
	RequiredTitlesRegex = map[string]string{
		"VP":   "VP",
		"Head": "^Head of .*$",
		"CEO":  "CEO",
		"CTO":  "CTO",
		"COO":  "COO",
	}
	titleRegex = []regexp.Regexp{}

	_ Relationship = (*LinkedIn)(nil) // make sure LinkedIn type implement all Relationship interface.
)

func init() {
	for _, rt := range RequiredTitlesRegex {
		titleRegex = append(titleRegex, *regexp.MustCompile(rt))
	}
}

func main() {
	var myLinkedInUrl string
	var csvFilePath string
	flag.StringVar(&myLinkedInUrl, "linkedin-url", "", "the LinkedIn URL for scoring")
	flag.StringVar(&csvFilePath, "csv-filepath", "", "the csv filepath for searching")
	flag.Parse()

	lks, err := NewLinkedIn(myLinkedInUrl)
	if err != nil {
		log.Fatalf("failed to initialize linkedin instance: %s", err.Error())
	}
	lks.ShowRelationScores()
}

/*
Mocked LinkedIn APIs
*/

// LinkedInProfileAPIGet Get the very Profile Object identified by URL from LinkedIn
func LinkedInProfileAPIGet(url string) ([]byte, error) {
	return []byte(`{
		"first_name": "Stanley",
		"last_name": "Liu",
		"phone_numbers": {
			"personal": {
				"mobile": ["+17605835578"],
				"home": []
			},
			"work": []
		},
		"emails": {
			"personal": ["stanleykliu92@gmail.com", "stanleyliu@berkeley.edu"],
			"work": ["stan@trypartnerhq.com", "stan@notablehealth.com"]
		},
		"experiences": [
			{
				"company": "PartnerHQ",
				"company_facebook_profile_url": null,
				"company_linkedin_profile_url": "https://www.linkedin.com/company/trypartnerhq",
				"logo_url": "https://media.licdn.com/dms/image/C560BAQEYxazZM_hXgQ/company-logo_100_100/0/1634934418976?e=2147483647\u0026v=beta\u0026t=wI0YdMmxIctkzvnKxRfuAbT8h5eok_DlUqEph68J37s",
				"starts_at": {
					"day": 1,
					"month": 1,
					"year": 2024
				},
				"title": "Co-Founder"
			},
			{
				"company": "Notable Health",
				"company_facebook_profile_url": null,
				"company_linkedin_profile_url": "https://www.linkedin.com/company/notablehealth",
				"logo_url": "https://media.licdn.com/dms/image/C560BAQEYxazZM_hXgQ/company-logo_100_100/0/1634934418976?e=2147483647\u0026v=beta\u0026t=wI0YdMmxIctkzvnKxRfuAbT8h5eok_DlUqEph68J37s",
				"ends_at": {
					"day": 1,
					"month": 12,
					"year": 2023
				},
				"starts_at": {
					"day": 1,
					"month": 7,
					"year": 2018
				},
				"title": "Founding Engineer"
			}
		]
	}`), nil
}
