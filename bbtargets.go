// see https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	output := GetTargets()
	if len(output) == 0 {
		fmt.Println("[ ! ] Error while retrieving targets.")
		os.Exit(1)
	}
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
		return []string{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return []string{}
	}

	var output []string

	for _, res := range results.Targets {
		if res.Bounty {
			if strings.Contains(res.Url, "hackerone") || strings.Contains(res.Url, "bugcrowd") ||
				strings.Contains(res.Url, "intigriti") || strings.Contains(res.Url, "yeswehack") {
				output = append(output, cleanIgnored(res.Domains)...)
			}
		}
	}
	return output
}

func cleanIgnored(domains []string) []string {
	var ignored = readFile("ignored.txt")
	var ignoredsubs []string
	for _, domain := range domains {
		for _, forb := range ignored {
			if strings.Contains(domain, forb) {
				ignoredsubs = append(ignoredsubs, domain)
			}
		}
	}
	return difference(domains, ignoredsubs)
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

//readFile >
func readFile(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open %s ", inputFile)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	var dir = ""
	for scanner.Scan() {
		dir = scanner.Text()
		if len(dir) > 0 {
			text = append(text, dir)
		}
	}
	file.Close()
	text = removeDuplicateValues(text)
	return text
}

//removeDuplicateValues >
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
