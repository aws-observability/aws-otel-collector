package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cenkalti/backoff"
	docker "github.com/docker/docker/client"
	"github.com/google/go-containerregistry/pkg/crane"
)

const (
	ghcr                 = "ghcr.io"
	gcr                  = "gcr.io"
	quay                 = "quay.io"
	defaultSleepDuration = 60 * time.Second
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

// GHCRTagsResponse contains the tags' information of the HTTP get response from ghcr.io.
type GHCRTagsResponse struct {
	Tags []string `json:"tags"`
}

// GCRTagsResponse contains the tags' information of the HTTP get response from gcr.io.
type GCRTagsResponse struct {
	Tags []string `json:"tags"`
}

// GHCRToken contains the necessary token information to get an HTTP response from ghcr.io
type GHCRToken struct {
	Token string
}

// QuayTagsResponse contains the tags' information of the HTTP get response from quay.io.
type QuayTagsResponse struct {
	Tags []RepositoryTag `json:"tags"`
}

// RepositoryTag holds the individual tag for the requested repository.
type RepositoryTag struct {
	Name string `json:"name"`
}

type mirror struct {
	ctx          context.Context
	dockerClient *docker.Client  // docker client used to pull, tag and push images
	ecrManager   *ecrManager     // ECR manager, used to ensure the ECR repository exist
	sourceRepo   Repository      // source repo to mirror from
	targetRepo   Repository      // target repo to mirror to
	remoteTags   []RepositoryTag // list of remote repository tags (post filtering)
}

func (m *mirror) setup(repos []Repository) error {
	m.sourceRepo = repos[0]
	m.targetRepo = repos[1]

	// Fetch remote tags from source repository.
	m.getRemoteTags()

	return nil
}

func (m *mirror) work() {
	if err := m.ecrManager.ensure(m.ctx, m.targetRepo.Name); err != nil {
		log.Fatalf("Failed to create ECR repo %s: %v", m.targetRepo.Name, err)
	}

	for _, tag := range m.remoteTags {
		source := fmt.Sprintf("%v:%v", m.sourceRepositoryFullName(), tag.Name)
		target := fmt.Sprintf("%v:%v", m.targetRepositoryName(), tag.Name)
		if err := crane.Copy(source, target); err != nil {
			log.Printf("Failed to copy docker image %s:%s: %v", m.sourceRepositoryFullName(), tag.Name, err)
			continue
		}

	}
}

func (m *mirror) getRemoteTags() {
	var url string
	switch m.sourceRepo.Host {
	case ghcr:
		url = fmt.Sprintf("https://ghcr.io/v2/%s/tags/list", m.sourceRepositoryName())
	case gcr:
		url = fmt.Sprintf("https://gcr.io/v2/%s/tags/list", m.sourceRepositoryName())
	case quay:
		url = fmt.Sprintf("https://quay.io/api/v1/repository/%s/tag", m.sourceRepositoryName())
	}

	if err := backoff.Retry(m.getTagResponseBackoff(url), backoffSettings); err != nil {
		log.Fatalf("Could not fetch tag information for %s: %v", url, err)
	}

	log.Printf("Finished scraping remote tags from %s", m.sourceRepositoryFullName())
}

func (m *mirror) sourceRepositoryName() string {
	return fmt.Sprintf("%s/%s", m.sourceRepo.Registry, m.sourceRepo.Name)
}

func (m *mirror) sourceRepositoryFullName() string {
	return fmt.Sprintf("%s/%s/%s", m.sourceRepo.Host, m.sourceRepo.Registry, m.sourceRepo.Name)
}

func (m *mirror) targetRepositoryName() string {
	return fmt.Sprintf("%s/%s", m.targetRepo.Registry, m.targetRepo.Name)
}

func (m *mirror) getTagResponse(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Sets the authorization header necessary for accessing ghcr.io
	if m.sourceRepo.Host == ghcr {
		tokenURL := fmt.Sprintf("https://ghcr.io/token?scope=repository:%s:pull", m.sourceRepositoryName())
		// G107: Potential HTTP request made with variable url
		// #nosec G107
		tokenRes, err := http.Get(tokenURL)
		if err != nil {
			return err
		}
		defer func() { _ = tokenRes.Body.Close() }()

		token := new(GHCRToken)
		err = json.NewDecoder(tokenRes.Body).Decode(token)
		if err != nil {
			return err
		}

		authToken := "Bearer " + token.Token
		req.Header.Set("Authorization", authToken)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Failed to get %s, retrying", url)
		return err
	} else if res.StatusCode == http.StatusTooManyRequests {
		sleepTime := getSleepTime(res.Header.Get("X-RateLimit-Reset"), time.Now())
		log.Printf("Rate limited on %s, sleeping for %s", url, sleepTime)
		time.Sleep(sleepTime)
		return fmt.Errorf("encounter error: too many requests")
	} else if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("get %s failed with %d, retrying", url, res.StatusCode)
	} else {
		defer func() { _ = res.Body.Close() }()

		// Decode the response and add the tags to remoteTags field of mirror struct.
		var allTags []RepositoryTag
		dc := json.NewDecoder(res.Body)

		switch m.sourceRepo.Host {
		case ghcr:
			var tags GHCRTagsResponse
			if err := dc.Decode(&tags); err != nil {
				return err
			}
			for _, tag := range tags.Tags {
				// Check if the ghcr image is on allowlist
				if tagInAllowlist(tag, m.sourceRepo.AllowedTags) {
					allTags = append(allTags, RepositoryTag{
						Name: tag,
					})
				}
			}
		case gcr:
			var tags GCRTagsResponse
			if err := dc.Decode(&tags); err != nil {
				return err
			}
			for _, tag := range tags.Tags {
				// Check if kube-rbac-proxy image is on allowlist
				if tagInAllowlist(tag, m.sourceRepo.AllowedTags) {
					allTags = append(allTags, RepositoryTag{
						Name: tag,
					})
				}
			}
		case quay:
			var tags QuayTagsResponse
			if err := dc.Decode(&tags); err != nil {
				return err
			}
			//			allTags = append(allTags, tags.Tags...)
			for _, tag := range tags.Tags {
				// Check if the kube-rbac-proxy image is on allowlist
				if tagInAllowlist(tag.Name, m.sourceRepo.AllowedTags) {
					allTags = append(allTags, RepositoryTag{
						Name: tag.Name,
					})
				}
			}
		}

		m.remoteTags = allTags

		return nil
	}
}

func (m *mirror) getTagResponseBackoff(url string) backoff.Operation {
	return func() error {
		return m.getTagResponse(url)
	}
}

func getSleepTime(rateLimitReset string, now time.Time) time.Duration {
	rateLimitResetInt, err := strconv.ParseInt(rateLimitReset, 10, 64)

	if err != nil {
		return defaultSleepDuration
	}

	sleepTime := time.Unix(rateLimitResetInt, 0)
	calculatedSleepTime := sleepTime.Sub(now)

	if calculatedSleepTime < (0 * time.Second) {
		return 0 * time.Second
	}

	return calculatedSleepTime
}

func tagInAllowlist(tag string, allowlist []string) bool {
	for _, allowed := range allowlist {
		if allowed == tag {
			return true
		}
	}
	return false
}
