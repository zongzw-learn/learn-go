package main

import (
	"errors"
	"fmt"
	"regexp"
)

// Relationship defines public interfaces
type Relationship interface {
	LoadRelations(csv string) error
	ShowRelationScores()
}

// LinkedInConnection the connection entity parsed from CSV
type LinkedInConnection struct {
	FirstName    string
	LastName     string
	URL          string
	EmailAddress string
	Company      string
	Position     string
	ConnectedOn  string
}

// EnrichedProfile the profile object containing related information retrieved from LinkedIn
type EnrichedProfile struct {
	FirstName    string
	LastName     string
	PhoneNumbers PhoneNumbers
	Emails       Emails
	Experiences  []Experience
}

// PhoneNumbers sub-struct embeded in EnrichedProfile
type PhoneNumbers struct {
	Personal map[string][]string
	Work     []string
}

// Emails sub-struct embeded in EnrichedProfile
type Emails struct {
	Personal []string
	Work     []string
}

// Experience sub-struct embeded in EnrichedProfile
type Experience struct {
	Company                   string
	CompanyFacebookProfileUrl string
	CompanyLinkedinProfileUrl string
	LogoUrl                   string
	EndsAt                    Date
	StartsAt                  Date
	Title                     string
}

// Date sub-struct embeded in Experience
type Date struct {
	Day, Month, Year int
}

type LinkedIn struct {
	URL         string
	Connections []LinkedInConnection
	Profiles    map[string]EnrichedProfile
	Scores      map[string]int
	titleRegexs []regexp.Regexp
}

// LoadRelations Load CSV items from filepath
func (lk *LinkedIn) LoadRelations(csv string) error {
	errs := []error{}
	return errors.Join(errs...)
}

// ShowRelationScores Show the scores for all given LinkedIn relations
func (lk *LinkedIn) ShowRelationScores() {
	fmt.Printf("Relation Scores for user %s: \n", lk.URL)
	for k, v := range lk.Scores {
		fmt.Printf("    %50s: %d\n", k, v)
	}
}

// enrichProfile Enrich the profiles for given linkedin relation
func (lk *LinkedIn) enrichProfile(url string) (*EnrichedProfile, error) {
	// call LinkedIn API to get the profile
	return &EnrichedProfile{}, nil
}

// scoreRelation Calculate the relationship between user and the given connection
//
//	take in the connection’s information
//	return a score from 0-10 based on the strength of their relationship, saved in LinkedIn.Scores
func (lk *LinkedIn) scoreRelation(url string) (int, error) {
	return 0, nil
}

// NewLinkedIn instanlize the LinkedIn object
func NewLinkedIn(url string) (Relationship, error) {
	l := LinkedIn{
		URL:         url,
		Connections: []LinkedInConnection{},
		Profiles:    map[string]EnrichedProfile{},
		Scores:      map[string]int{},
		titleRegexs: []regexp.Regexp{},
	}
	return &l, nil
}

var (
	_ Relationship = (*LinkedIn)(nil) // make sure LinkedIn type implement all Relationship interface.
)

func main() {
	var relation Relationship
	relation, _ = NewLinkedIn("my-linkedin-url")
	_ = relation.LoadRelations("my-csv-filepath")
	relation.ShowRelationScores()
}
