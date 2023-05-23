package machineid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// run wraps `exec.Command` with easy access to stdout and stderr.
func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}

// protect calculates HMAC-SHA256 of the application ID, keyed by the machine ID and returns a hex-encoded string.
func protect(appID, id string) string {
	mac := hmac.New(sha256.New, []byte(id))
	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil))
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func writeFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// readFirstFile tries all the pathnames listed and returns the contents of the first readable file.
func readFirstFile(pathnames []string) ([]byte, error) {
	contents := []byte{}
	var err error
	for _, pathname := range pathnames {
		if pathname != "" {
			contents, err = readFile(pathname)
			if err == nil {
				return contents, nil
			}
		}
	}
	return contents, err
}

// writeFirstFile writes to the first file that "works" between all pathnames listed.
func writeFirstFile(pathnames []string, data []byte) error {
	var err error
	for _, pathname := range pathnames {
		if pathname != "" {
			err = writeFile(pathname, data)
			if err == nil {
				return nil
			}
		}
	}
	return err
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
