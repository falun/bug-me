package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

func main() {
	username := ""
	pat := ""
	flag.StringVar(&username, "u", "", "github user name for backing item store")
	flag.StringVar(&pat, "p", "", "PAT for github user")
	flag.Parse()

	if pat == "" || username == "" {
		return
	}

	bat := github.BasicAuthTransport{username, pat, "", nil}
	client := github.NewClient(bat.Client())

	// ([]Membership, *Response, error)
	memberships, _, err := client.Organizations.ListOrgMemberships(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Memberships:")
	for _, m := range memberships {
		st := "<unset>"
		if m.State != nil {
			st = *m.State
		}

		rl := "<unset>"
		if m.Role != nil {
			rl = *m.Role
		}

		/*
			ttl := "<?>"
			switch {
			case m.Organization.Name != nil:
				ttl = *m.Organization.Name
			case m.Organization.Company != nil:
				ttl = "company - " + *m.Organization.Company
			}
		*/

		// (*User, *Response, error)
		u, _, e := client.Users.Get(*m.Organization.Login)
		if e != nil {
			log.Fatal(e)
		}

		b, e := json.MarshalIndent(u, "    ", "  ")
		if e != nil {
			log.Fatal(e)
		}

		fmt.Printf("  %s - %s:\n    %s\n", st, rl, string(b))
	}
}
