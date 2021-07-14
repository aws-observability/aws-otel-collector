package main

import (
	"context"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/cenkalti/backoff"
	docker "github.com/fsouza/go-dockerclient"
	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	SourceRepos []Repository `yaml:"sourceRepos"`
	TargetRepos []Repository `yaml:"targetRepos"`
}

type Repository struct {
	Registry  string `yaml:"registry"`
	Name   	  string `yaml:"name"`
	Host      string `yaml:"host"`
}

// Create the Docker client which will be used to pull, retag and push images.
func createDockerClient() (*docker.Client, error) {
	client, err := docker.NewClientFromEnv()
	return client, err
}

// worker is the mirror working function.
func worker(wg *sync.WaitGroup, workerCh chan []Repository, dc *DockerClient, ecrm *ecrManager) {
	for {
		select {
		case repos := <-workerCh:
			log.Printf("Starting to mirror repo %s/%s/%s to repo %s/%s", repos[0].Host, repos[0].Registry, repos[0].Name, repos[1].Registry, repos[1].Name)
			m := mirror{
				dockerClient: dc,
				ecrManager:   ecrm,
			}
			if err := m.setup(repos); err != nil {
				log.Printf("Failed to setup mirror for repo %s/%s/%s: %v", repos[0].Host, repos[0].Registry, repos[0].Name, err)
				wg.Done()
				continue
			}

			m.work()
			log.Printf("From repo %s/%s/%s to repo %s/%s mirror completed", repos[0].Host, repos[0].Registry, repos[0].Name, repos[1].Registry, repos[1].Name)
			wg.Done()
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}

	log.Println("Creating Docker client")
	var client DockerClient
	client, err = createDockerClient()
	if err != nil {
		log.Fatalf("Could not create Docker client: %v", err)
	}

	info, err := client.Info()
	if err != nil {
		log.Fatalf("Could not get Docker info: %v", err)
	}
	log.Printf("Connected to Docker daemon: %s @ %s", info.Name, info.ServerVersion)

	log.Println("Creating AWS client")
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}
	// ECR public registries are only on us-east-1 region.
	cfg.Region = "us-east-1"

	ecrManager := &ecrManager{client: ecrpublic.NewFromConfig(cfg)}
	backoffSettings := backoff.NewExponentialBackOff()
	backoffSettings.InitialInterval = 1 * time.Second
	backoffSettings.MaxElapsedTime = 10 * time.Second
	notifyError := func(err error, d time.Duration) {
		log.Fatalf("%v (%s)", err, d.String())
	}

	// Build ECR registries cache.
	if err = backoff.RetryNotify(ecrManager.buildCacheBackoff(), backoffSettings, notifyError); err != nil {
		log.Fatalf("Could not build ECR cache: %v", err)
	}

	workersNum := runtime.NumCPU()

	workerCh := make(chan []Repository, 5)
	var wg sync.WaitGroup

	for i := 0; i < workersNum; i++ {
		go worker(&wg, workerCh, &client, ecrManager)
	}

	// Add source repo and target repo pair to worker channel.
	for i, _ := range config.SourceRepos {
		wg.Add(1)
		workerCh <- []Repository{config.SourceRepos[i], config.TargetRepos[i]}
	}

	wg.Wait()
	log.Println("Finished mirroring repositories")
}
