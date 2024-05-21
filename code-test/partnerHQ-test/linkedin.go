package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

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
	csvCntn, err := os.ReadFile(csv)
	if err != nil {
		log.Printf("failed to read connections from file: %s", csv)
		return err
	}
	lines := strings.Split(string(csvCntn), "\n")
	errs := []error{}
	for _, line := range lines {
		conn, err := unmarshalLine([]byte(line))
		if err != nil {
			errs = append(errs, err)
		}
		lk.Connections = append(lk.Connections, *conn)
	}

	// TODO: we may do the enrich via multithread
	for _, conn := range lk.Connections {
		if !isRequiredTitle(conn.Position) {
			continue
		}
		p, err := lk.enrichProfile(conn.URL)
		if err != nil {
			errs = append(errs, err)
		}
		lk.Profiles[conn.URL] = *p
	}

	// retrieve myself's profile
	myprofile, err := lk.enrichProfile(lk.URL)
	if err != nil {
		errs = append(errs, err)
	}
	lk.Profiles[lk.URL] = *myprofile

	for url := range lk.Profiles {
		if url == lk.URL {
			continue
		}
		lk.Scores[url], err = lk.scoreRelation(url)
		if err != nil {
			errs = append(errs, err)
		}
	}
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
	profileBytes, err := LinkedInProfileAPIGet(url)
	if err != nil {
		return nil, err
	}

	var profile EnrichedProfile
	err = json.Unmarshal(profileBytes, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// scoreRelation Calculate the relationship between user and the given connection
//
//	take in the connection’s information
//	return a score from 0-10 based on the strength of their relationship, saved in LinkedIn.Scores
func (lk *LinkedIn) scoreRelation(url string) (int, error) {
	var prof EnrichedProfile
	prof, f := lk.Profiles[url]
	if !f {
		return -1, fmt.Errorf("not found profile for %s", url)
	}

	// TODO: should be optimized by sorting experiences
	scores := 0
	myprof := lk.Profiles[lk.URL]
	for _, myexpr := range myprof.Experiences {
		for _, expr := range prof.Experiences {
			scores += getOverlapPercentage(myexpr.StartsAt, myexpr.EndsAt, expr.StartsAt, expr.EndsAt)
		}
	}
	return scores, nil
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
	for _, rt := range RequiredTitlesRegex {
		l.titleRegexs = append(l.titleRegexs, *regexp.MustCompile(rt))
	}
	return &l, nil
}
