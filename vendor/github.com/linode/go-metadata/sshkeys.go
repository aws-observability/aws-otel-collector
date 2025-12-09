package metadata

import "context"

// SSHKeysData contains information about SSH keys
// relevant to the current Linode instance.
type SSHKeysData struct {
	Users map[string][]string `json:"users"`
}

// GetSSHKeys gets all SSH keys for the current instance.
func (c *Client) GetSSHKeys(ctx context.Context) (*SSHKeysData, error) {
	req := c.R(ctx).SetResult(&SSHKeysData{})

	resp, err := coupleAPIErrors(req.Get("ssh-keys"))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*SSHKeysData), nil
}
