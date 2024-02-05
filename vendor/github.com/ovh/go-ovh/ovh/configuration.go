package ovh

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"gopkg.in/ini.v1"
)

var configPaths = []string{
	// System wide configuration
	"/etc/ovh.conf",
	// Configuration in user's home
	"~/.ovh.conf",
	// Configuration in local folder
	"./ovh.conf",
}

// currentUserHome attempts to get current user's home directory.
func currentUserHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		// Fallback by trying to read $HOME
		if userHome := os.Getenv("HOME"); userHome != "" {
			return userHome, nil
		}
		return "", err
	}

	return usr.HomeDir, nil
}

// configPaths returns configPaths, with ~/ prefix expanded.
func expandConfigPaths() []interface{} {
	paths := []interface{}{}

	// Will be initialized on first use
	var home string
	var homeErr error

	for _, path := range configPaths {
		if strings.HasPrefix(path, "~/") {
			// Find home if needed
			if home == "" && homeErr == nil {
				home, homeErr = currentUserHome()
			}
			// Ignore file in HOME if we cannot find it
			if homeErr != nil {
				continue
			}

			path = home + path[1:]
		}

		paths = append(paths, path)
	}

	return paths
}

// loadINI builds a ini.File from the configuration paths provided in configPaths.
// It's a helper for loadConfig.
func loadINI() (*ini.File, error) {
	paths := expandConfigPaths()
	if len(paths) == 0 {
		return ini.Empty(), nil
	}

	return ini.LooseLoad(paths[0], paths[1:]...)
}

// loadConfig loads client configuration from params, environments or configuration
// files (by order of decreasing precedence).
//
// loadConfig will check OVH_CONSUMER_KEY, OVH_APPLICATION_KEY, OVH_APPLICATION_SECRET
// and OVH_ENDPOINT environment variables. If any is present, it will take precedence
// over any configuration from file.
//
// Configuration files are ini files. They share the same format as python-ovh,
// node-ovh, php-ovh and all other wrappers. If any wrapper is configured, all
// can re-use the same configuration. loadConfig will check for configuration in:
//
// - ./ovh.conf
// - $HOME/.ovh.conf
// - /etc/ovh.conf
func (c *Client) loadConfig(endpointName string) error {
	if strings.HasSuffix(endpointName, "/") {
		return fmt.Errorf("endpoint name cannot have a tailing slash")
	}

	// Load configuration files by order of increasing priority. All configuration
	// files are optional. Only load file from user home if home could be resolve
	cfg, err := loadINI()
	if err != nil {
		return fmt.Errorf("cannot load configuration: %w", err)
	}

	// Canonicalize configuration
	if endpointName == "" {
		endpointName = getConfigValue(cfg, "default", "endpoint", "ovh-eu")
	}

	if c.AppKey == "" {
		c.AppKey = getConfigValue(cfg, endpointName, "application_key", "")
	}

	if c.AppSecret == "" {
		c.AppSecret = getConfigValue(cfg, endpointName, "application_secret", "")
	}

	if c.ConsumerKey == "" {
		c.ConsumerKey = getConfigValue(cfg, endpointName, "consumer_key", "")
	}

	// Load real endpoint URL by name. If endpoint contains a '/', consider it as a URL
	if strings.Contains(endpointName, "/") {
		c.endpoint = endpointName
	} else {
		c.endpoint = Endpoints[endpointName]
	}

	// If we still have no valid endpoint, AppKey or AppSecret, return an error
	if c.endpoint == "" {
		return fmt.Errorf("unknown endpoint '%s', consider checking 'Endpoints' list of using an URL", endpointName)
	}
	if c.AppKey == "" {
		return fmt.Errorf("missing application key, please check your configuration or consult the documentation to create one")
	}
	if c.AppSecret == "" {
		return fmt.Errorf("missing application secret, please check your configuration or consult the documentation to create one")
	}

	return nil
}

// getConfigValue returns the value of OVH_<NAME> or "name" value from "section". If
// the value could not be read from either env or any configuration files, return 'def'
func getConfigValue(cfg *ini.File, section, name, def string) string {
	// Attempt to load from environment
	fromEnv := os.Getenv("OVH_" + strings.ToUpper(name))
	if len(fromEnv) > 0 {
		return fromEnv
	}

	// Attempt to load from configuration
	fromSection := cfg.Section(section)
	if fromSection == nil {
		return def
	}

	fromSectionKey := fromSection.Key(name)
	if fromSectionKey == nil {
		return def
	}
	return fromSectionKey.String()
}
