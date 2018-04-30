package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/blang/semver"
	"github.com/mitchellh/mapstructure"
)

// A RepoResponse is part of the Bitbucket API's /api/repositories response.
type RepoResponse struct {
	Name  string
	SCMID string
	Links Links
}

// A Links object contains links to Bitbucket API objects.
type Links struct {
	Clone []Link
	Self  []Link
}

// A Link object contains many kinds of links.
type Link struct {
	Name string
	Href string
}

// A Contributor is a single git author and a count of their contributions.
type Contributor struct {
	Name  string
	Count int
}

var (
	flagDebug     = flag.Bool("debug", false, "print debug info to stderr")
	flagInsecure  = flag.Bool("insecure-skip-verify-ssl", false, "do not validate SSL certificates")
	flagDeepClone = flag.Bool("force-deep-clone", false, "use a deep clone (slower but works on old git versions)")
	flagDays      = flag.Int("days", 90, "days over which to count contributors")
)

func main() {
	// Read user inputs
	flag.Parse()
	user := os.Getenv("BITBUCKET_USER")
	password := os.Getenv("BITBUCKET_PASSWORD")
	server, err := url.Parse(os.Getenv("BITBUCKET_URL"))
	if err != nil {
		panic(err)
	}
	days := strconv.Itoa(*flagDays)

	// Auto-detect git version
	shallowCloneVersion := semver.MustParse("2.11.0")
	v, err := exec.Command("git", "--version").Output()
	if err != nil {
		panic(err)
	}
	gitVersion, err := semver.Parse(strings.TrimSpace(strings.TrimPrefix(string(v), "git version ")))
	if err != nil {
		panic(err)
	}
	useDeepClone := gitVersion.LT(shallowCloneVersion) || *flagDeepClone
	debugf("Using deep clone: %t", useDeepClone)

	// Get all visible repositories
	reposURL, err := server.Parse("/rest/api/1.0/repos")
	if err != nil {
		panic(err)
	}
	reposPages, err := GetPaged(reposURL, user, password)
	if err != nil {
		panic(err)
	}
	var repos []RepoResponse
	for _, entry := range reposPages {
		var res RepoResponse
		err = mapstructure.Decode(entry, &res)
		if err != nil {
			panic(err)
		}
		repos = append(repos, res)
	}
	debugf("Repos: %#v", repos)

	// Analyze each repository
	contributors := make(map[string]int)
	for _, repo := range repos {
		// Get clone URL
		if repo.SCMID != "git" {
			warnf("Unsupported SCM type (%s) for repository %s", repo.SCMID, repo.Name)
			continue
		}
		var cloneURL *url.URL
		for _, link := range repo.Links.Clone {
			if link.Name == "http" || link.Name == "https" {
				cloneURL, err = url.Parse(link.Href)
				if err != nil {
					panic(err)
				}
				cloneURL.User = url.UserPassword(user, password)
				debugf("Clone URL: %#v", cloneURL.String())
			}
		}
		if cloneURL == nil {
			panic("could not find clone URL")
		}

		// Clone repository locally into temporary directory
		dir, err := ioutil.TempDir("", "fossa-contributor-count-")
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(dir)
		debugf("TempDir: %#v", dir)
		cloneArgs := []string{"clone", cloneURL.String(), dir}
		if !useDeepClone {
			cloneArgs = append(cloneArgs, fmt.Sprintf("--shallow-since=%s days ago", days))
		}
		_, err = exec.Command("git", cloneArgs...).Output()
		if err != nil {
			panic(err)
		}

		// Run `git shortlog` and parse output
		cmd := exec.Command(
			"git", "shortlog", "HEAD", "--summary", "--email", "--numbered", fmt.Sprintf("--since=%s days ago", days))
		cmd.Dir = dir
		shortlogBytes, err := cmd.Output()
		if err != nil {
			debugf("Stderr: %#v", string(err.(*exec.ExitError).Stderr))
			warnf("Warning: could not get contributors from repository %#v (it may be empty)", repo.Name)
			continue
		}
		shortlog := strings.TrimSpace(string(shortlogBytes))
		debugf("Shortlog: %#v", shortlog)
		for _, line := range strings.Split(shortlog, "\n") {
			r := regexp.MustCompile("\\s*([0-9]+)\\s+(.*?) <(.*?)>")
			matches := r.FindStringSubmatch(line)
			debugf("Line: %#v %#v", line, matches)
			contributor := matches[3]
			contributions, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			contributors[contributor] += contributions
		}
	}

	// Format output: sort contributors by contribution count
	output := fmt.Sprintf("Found %d contributors:", len(contributors))
	var contributorList []Contributor
	for contributor, contributions := range contributors {
		contributorList = append(contributorList, Contributor{
			Name:  contributor,
			Count: contributions,
		})
	}
	sort.Slice(contributorList, func(i, j int) bool {
		if contributorList[i].Count == contributorList[j].Count {
			return contributorList[i].Name < contributorList[j].Name
		}
		return contributorList[i].Count > contributorList[j].Count
	})
	for _, contributor := range contributorList {
		output += fmt.Sprintf("\n%4d %s", contributor.Count, contributor.Name)
	}
	fmt.Println(output)
}
