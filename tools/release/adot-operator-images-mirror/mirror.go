package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cenkalti/backoff"
	docker "github.com/fsouza/go-dockerclient"
)

const (
	ecrPublic = "public.ecr.aws"
	quay = "quay.io"
	gcr = "gcr.io"
	defaultSleepDuration = 60 * time.Second
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

// QuayTagsResponse contains the tags information of the HTTP get response from quay.io.
type QuayTagsResponse struct {
	Tags []RepositoryTag `json:"tags"`
}

// GCRTagsResponse contains the tags information of the HTTP get response from gcr.io.
type GCRTagsResponse struct {
	Tags []string `json:"tags"`
}

// RepositoryTag holds the individual tag for the requested repository.
type RepositoryTag struct {
	Name        string    `json:"name"`
}

type DockerClient interface {
	Info() (*docker.DockerInfo, error)
	TagImage(string, docker.TagImageOptions) error
	PullImage(docker.PullImageOptions, docker.AuthConfiguration) error
	PushImage(docker.PushImageOptions, docker.AuthConfiguration) error
	RemoveImage(string) error
}

type mirror struct {
	dockerClient *DockerClient   // docker client used to pull, tag and push images
	ecrManager   *ecrManager     // ECR manager, used to ensure the ECR repository exist
	sourceRepo   Repository      // source repo to mirror from
	targetRepo   Repository      // target repo to mirror to
	remoteTags   []RepositoryTag // list of remote repository tags (post filtering)
}

func (m *mirror) setup(repos []Repository) (err error) {
	m.sourceRepo = repos[0]
	m.targetRepo = repos[1]

	// Fetch remote tags from source repository.
	m.getRemoteTags()

	// Filter the tags we got.
	m.filterTags()

	return nil
}

func (m *mirror) pullImage(tag string) error {
	pullOptions := docker.PullImageOptions{
		Repository:        m.sourceRepositoryFullName(),
		Tag:               tag,
		InactivityTimeout: 1 * time.Minute,
	}

	return (*m.dockerClient).PullImage(pullOptions, docker.AuthConfiguration{})
}

func (m *mirror) tagImage(tag string) error {
	tagOptions := docker.TagImageOptions{
		Repo:  m.targetRepositoryName(),
		Tag:   tag,
		Force: true,
	}

	return (*m.dockerClient).TagImage(m.sourceRepositoryFullName() + ":" + tag, tagOptions)
}

func (m *mirror) pushImage(tag string) error {
	pushOptions := docker.PushImageOptions{
		Name:              m.targetRepositoryName(),
		Registry:          m.targetRepo.Registry,
		Tag:               tag,
		InactivityTimeout: 1 * time.Minute,
	}

	// Get ecrpublic credentials.
	creds, err := getDockerCredentials(ecrPublic)
	if err != nil {
		return err
	}

	return (*m.dockerClient).PushImage(pushOptions, *creds)
}

func (m *mirror) work() {
	if err := m.ecrManager.ensure(m.targetRepo.Name); err != nil {
		log.Fatalf("Failed to create ECR repo %s: %v", m.targetRepo.Name, err)
	}

	for _, tag := range m.remoteTags {
		if err := m.pullImage(tag.Name); err != nil {
			log.Printf("Failed to pull docker image %s:%s: %v", m.sourceRepositoryFullName(), tag.Name, err)
			continue
		}

		if err := m.tagImage(tag.Name); err != nil {
			log.Printf("Failed to retag docker image %s:%s: %v", m.sourceRepositoryFullName(), tag.Name, err)
			continue
		}

		if err := m.pushImage(tag.Name); err != nil {
			log.Printf("Failed to push retagged image %s:%s: %v", m.targetRepositoryName(), tag.Name, err)
			continue
		}
	}
}

func (m *mirror) getRemoteTags() {
	var url string
	switch m.sourceRepo.Host {
	case quay:
		url = fmt.Sprintf("https://quay.io/api/v1/repository/%s/tag", m.sourceRepositoryName())
	case gcr:
		url = fmt.Sprintf("https://gcr.io/v2/%s/tags/list", m.sourceRepositoryName())
	}

	if err := backoff.Retry(m.getTagResponseBackoff(url), backoffSettings); err != nil {
		log.Fatalf("Could not fetch tag information for %s: %v", url, err)
	}

	log.Printf("Finished scraping remote tags from %s", m.sourceRepositoryFullName())
}

func (m *mirror) filterTags() {
	tmp := make([]RepositoryTag, 0)

	for _, remoteTag := range m.remoteTags {
		// We only keep the tags starting with "v" or with the name "latest".
		if string(remoteTag.Name[0]) == "v" || remoteTag.Name == "latest" {
			tmp = append(tmp, remoteTag)
		}
	}

	m.remoteTags = tmp
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Failed to get %s, retrying", url)
		return err
	} else if res.StatusCode == 429 {
		sleepTime := getSleepTime(res.Header.Get("X-RateLimit-Reset"), time.Now())
		log.Printf("Rate limited on %s, sleeping for %s", url, sleepTime)
		time.Sleep(sleepTime)
		return fmt.Errorf("encouter error: too many requests")
	} else if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("get %s failed with %d, retrying", url, res.StatusCode)
	} else {
		defer res.Body.Close()

		// Decode the response and add the tags to remoteTags field of mirror struct.
		var allTags []RepositoryTag
		dc := json.NewDecoder(res.Body)

		switch m.sourceRepo.Host {
		case quay:
			var tags QuayTagsResponse
			if err := dc.Decode(&tags); err != nil {
				return err
			}
			allTags = append(allTags, tags.Tags...)
		case gcr:
			var tags GCRTagsResponse
			if err := dc.Decode(&tags); err != nil {
				return err
			}
			for _, tag := range tags.Tags {
				allTags = append(allTags, RepositoryTag{
					Name: tag,
				})
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
