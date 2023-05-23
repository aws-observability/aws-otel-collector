// +build linux

package machineid

import (
	"os"
	"path"
)

const (
	// the environment variable name pointing to the machine id pathname
	ENV_VARNAME = "MACHINE_ID_FILE"

	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"

	// this returns a random uuid each time it's read
	linuxRandomUuid = "/proc/sys/kernel/random/uuid"
)

// machineID returns the uuid specified in the "canonical" locations. If not such value is found, one is generated and persisted.
// The machine id is looked in:
//   - the file pointed by the `MACHINE_ID_FILE` env var
//   - `/var/lib/dbus/machine-id`
//   - `/etc/machine-id`
//   - `$HOME/.config/machine-id`
//
// If no such file is found, a random uuid is generated and persisted in the first
// writable file among `$MACHINE_ID_FILE`, `/var/lib/dbus/machine-id`, `/etc/machine-id`, `$HOME/.config/machine-id`.
//
// If there is an error reading _all_ the files an empty string is returned.
// The logic implemented is a variation of the one described in https://github.com/denisbrodbeck/machineid/issues/5#issuecomment-523803164
// See also https://unix.stackexchange.com/questions/144812/generate-consistent-machine-unique-id
func machineID() (string, error) {
	env_pathname := os.Getenv(ENV_VARNAME)

	home := os.Getenv("HOME")
	userMachineId := path.Join(home, ".config", "machine-id")

	id, err := readFirstFile([]string{
		env_pathname, dbusPath, dbusPathEtc, userMachineId,
	})
	if err != nil {
		id, err = readFile(linuxRandomUuid)
		if err == nil {
			writeFirstFile([]string{
				env_pathname, dbusPathEtc, dbusPath, userMachineId,
			}, id)
		}
	}
	if err != nil {
		return "", err
	}
	return trim(string(id)), nil
}
