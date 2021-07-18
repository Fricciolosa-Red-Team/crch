package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	output := GetTargets()
	for _, elem := range output {
		fmt.Println(elem)
	}
}

type Target struct {
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Bounty  bool     `json:"bounty"`
	Domains []string `json:"domains"`
}

type Programs struct {
	Targets []Target `json:"programs"`
}

func GetTargets() []string {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	var results Programs
	url := "https://raw.githubusercontent.com/projectdiscovery/public-bugbounty-programs/master/chaos-bugbounty-list.json"
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		fmt.Println(err)
		return []string{}
	}

	var output []string

	for _, res := range results.Targets {
		if res.Bounty {
			if strings.Contains(res.Url, "hackerone") || strings.Contains(res.Url, "bugcrowd") ||
				strings.Contains(res.Url, "intigriti") || strings.Contains(res.Url, "yeswehack") {
				output = append(output, cleanForbidden(res.Domains)...)
			}
		}
	}
	return output
}

func cleanForbidden(domains []string) []string {
	var forbidden = []string{"tweakblogs", ".onion"}
	var forbiddens []string
	for _, domain := range domains {
		for _, forb := range forbidden {
			if strings.Contains(domain, forb) {
				forbiddens = append(forbiddens, domain)
			}
		}
	}
	return difference(domains, forbiddens)
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
