package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type PostgresDatabaseTarget string

const (
	PostgresDatabaseTargetPrimary   PostgresDatabaseTarget = "primary"
	PostgresDatabaseTargetSecondary PostgresDatabaseTarget = "secondary"
)

type PostgresCommitType string

const (
	PostgresCommitTrue        PostgresCommitType = "true"
	PostgresCommitFalse       PostgresCommitType = "false"
	PostgresCommitLocal       PostgresCommitType = "local"
	PostgresCommitRemoteWrite PostgresCommitType = "remote_write"
	PostgresCommitRemoteApply PostgresCommitType = "remote_apply"
)

type PostgresReplicationType string

const (
	PostgresReplicationNone      PostgresReplicationType = "none"
	PostgresReplicationAsynch    PostgresReplicationType = "asynch"
	PostgresReplicationSemiSynch PostgresReplicationType = "semi_synch"
)

// A PostgresDatabase is an instance of Linode Postgres Managed Databases
type PostgresDatabase struct {
	ID        int            `json:"id"`
	Status    DatabaseStatus `json:"status"`
	Label     string         `json:"label"`
	Region    string         `json:"region"`
	Type      string         `json:"type"`
	Engine    string         `json:"engine"`
	Version   string         `json:"version"`
	AllowList []string       `json:"allow_list"`
	Port      int            `json:"port"`

	ClusterSize int              `json:"cluster_size"`
	Platform    DatabasePlatform `json:"platform"`

	// Members has dynamic keys so it is a map
	Members map[string]DatabaseMemberType `json:"members"`

	// Deprecated: ReplicationCommitType is a deprecated property, as it is no longer supported in DBaaS V2.
	ReplicationCommitType PostgresCommitType `json:"replication_commit_type"`
	// Deprecated: ReplicationType is a deprecated property, as it is no longer supported in DBaaS V2.
	ReplicationType PostgresReplicationType `json:"replication_type"`
	// Deprecated: SSLConnection is a deprecated property, as it is no longer supported in DBaaS V2.
	SSLConnection bool `json:"ssl_connection"`
	// Deprecated: Encrypted is a deprecated property, as it is no longer supported in DBaaS V2.
	Encrypted bool `json:"encrypted"`

	Hosts             DatabaseHost              `json:"hosts"`
	Updates           DatabaseMaintenanceWindow `json:"updates"`
	Created           *time.Time                `json:"-"`
	Updated           *time.Time                `json:"-"`
	Fork              *DatabaseFork             `json:"fork"`
	OldestRestoreTime *time.Time                `json:"-"`
	UsedDiskSizeGB    int                       `json:"used_disk_size_gb"`
	TotalDiskSizeGB   int                       `json:"total_disk_size_gb"`
}

func (d *PostgresDatabase) UnmarshalJSON(b []byte) error {
	type Mask PostgresDatabase

	p := struct {
		*Mask
		Created           *parseabletime.ParseableTime `json:"created"`
		Updated           *parseabletime.ParseableTime `json:"updated"`
		OldestRestoreTime *parseabletime.ParseableTime `json:"oldest_restore_time"`
	}{
		Mask: (*Mask)(d),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	d.Created = (*time.Time)(p.Created)
	d.Updated = (*time.Time)(p.Updated)
	d.OldestRestoreTime = (*time.Time)(p.OldestRestoreTime)
	return nil
}

// PostgresCreateOptions fields are used when creating a new Postgres Database
type PostgresCreateOptions struct {
	Label       string   `json:"label"`
	Region      string   `json:"region"`
	Type        string   `json:"type"`
	Engine      string   `json:"engine"`
	AllowList   []string `json:"allow_list,omitempty"`
	ClusterSize int      `json:"cluster_size,omitempty"`

	// Deprecated: Encrypted is a deprecated property, as it is no longer supported in DBaaS V2.
	Encrypted bool `json:"encrypted,omitempty"`
	// Deprecated: SSLConnection is a deprecated property, as it is no longer supported in DBaaS V2.
	SSLConnection bool `json:"ssl_connection,omitempty"`
	// Deprecated: ReplicationType is a deprecated property, as it is no longer supported in DBaaS V2.
	ReplicationType PostgresReplicationType `json:"replication_type,omitempty"`
	// Deprecated: ReplicationCommitType is a deprecated property, as it is no longer supported in DBaaS V2.
	ReplicationCommitType PostgresCommitType `json:"replication_commit_type,omitempty"`

	Fork *DatabaseFork `json:"fork,omitempty"`
}

// PostgresUpdateOptions fields are used when altering the existing Postgres Database
type PostgresUpdateOptions struct {
	Label       string                     `json:"label,omitempty"`
	AllowList   *[]string                  `json:"allow_list,omitempty"`
	Updates     *DatabaseMaintenanceWindow `json:"updates,omitempty"`
	Type        string                     `json:"type,omitempty"`
	ClusterSize int                        `json:"cluster_size,omitempty"`
	Version     string                     `json:"version,omitempty"`
}

// PostgresDatabaseSSL is the SSL Certificate to access the Linode Managed Postgres Database
type PostgresDatabaseSSL struct {
	CACertificate []byte `json:"ca_certificate"`
}

// PostgresDatabaseCredential is the Root Credentials to access the Linode Managed Database
type PostgresDatabaseCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ListPostgresDatabases lists all Postgres Databases associated with the account
func (c *Client) ListPostgresDatabases(ctx context.Context, opts *ListOptions) ([]PostgresDatabase, error) {
	return getPaginatedResults[PostgresDatabase](ctx, c, "databases/postgresql/instances", opts)
}

// PostgresDatabaseBackup is information for interacting with a backup for the existing Postgres Database
// Deprecated: PostgresDatabaseBackup is a deprecated struct, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
type PostgresDatabaseBackup struct {
	ID      int        `json:"id"`
	Label   string     `json:"label"`
	Type    string     `json:"type"`
	Created *time.Time `json:"-"`
}

func (d *PostgresDatabaseBackup) UnmarshalJSON(b []byte) error {
	type Mask PostgresDatabaseBackup

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(d),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	d.Created = (*time.Time)(p.Created)
	return nil
}

// PostgresBackupCreateOptions are options used for CreatePostgresDatabaseBackup(...)
// Deprecated: PostgresBackupCreateOptions is a deprecated struct, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
type PostgresBackupCreateOptions struct {
	Label  string                 `json:"label"`
	Target PostgresDatabaseTarget `json:"target"`
}

// ListPostgresDatabaseBackups lists all Postgres Database Backups associated with the given Postgres Database
// Deprecated: ListPostgresDatabaseBackups is a deprecated method, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
func (c *Client) ListPostgresDatabaseBackups(ctx context.Context, databaseID int, opts *ListOptions) ([]PostgresDatabaseBackup, error) {
	return getPaginatedResults[PostgresDatabaseBackup](ctx, c, formatAPIPath("databases/postgresql/instances/%d/backups", databaseID), opts)
}

// GetPostgresDatabase returns a single Postgres Database matching the id
func (c *Client) GetPostgresDatabase(ctx context.Context, databaseID int) (*PostgresDatabase, error) {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doGETRequest[PostgresDatabase](ctx, c, e)
}

// CreatePostgresDatabase creates a new Postgres Database using the createOpts as configuration, returns the new Postgres Database
func (c *Client) CreatePostgresDatabase(ctx context.Context, opts PostgresCreateOptions) (*PostgresDatabase, error) {
	return doPOSTRequest[PostgresDatabase](ctx, c, "databases/postgresql/instances", opts)
}

// DeletePostgresDatabase deletes an existing Postgres Database with the given id
func (c *Client) DeletePostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doDELETERequest(ctx, c, e)
}

// UpdatePostgresDatabase updates the given Postgres Database with the provided opts, returns the PostgresDatabase with the new settings
func (c *Client) UpdatePostgresDatabase(ctx context.Context, databaseID int, opts PostgresUpdateOptions) (*PostgresDatabase, error) {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doPUTRequest[PostgresDatabase](ctx, c, e, opts)
}

// PatchPostgresDatabase applies security patches and updates to the underlying operating system of the Managed Postgres Database
func (c *Client) PatchPostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/patch", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// GetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database
func (c *Client) GetPostgresDatabaseCredentials(ctx context.Context, databaseID int) (*PostgresDatabaseCredential, error) {
	e := formatAPIPath("databases/postgresql/instances/%d/credentials", databaseID)
	return doGETRequest[PostgresDatabaseCredential](ctx, c, e)
}

// ResetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database (may take a few seconds to work)
func (c *Client) ResetPostgresDatabaseCredentials(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/credentials/reset", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// GetPostgresDatabaseSSL returns the SSL Certificate for the given Postgres Database
func (c *Client) GetPostgresDatabaseSSL(ctx context.Context, databaseID int) (*PostgresDatabaseSSL, error) {
	e := formatAPIPath("databases/postgresql/instances/%d/ssl", databaseID)
	return doGETRequest[PostgresDatabaseSSL](ctx, c, e)
}

// GetPostgresDatabaseBackup returns a specific Postgres Database Backup with the given ids
// Deprecated: GetPostgresDatabaseBackup is a deprecated method, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
func (c *Client) GetPostgresDatabaseBackup(ctx context.Context, databaseID int, backupID int) (*PostgresDatabaseBackup, error) {
	e := formatAPIPath("databases/postgresql/instances/%d/backups/%d", databaseID, backupID)
	return doGETRequest[PostgresDatabaseBackup](ctx, c, e)
}

// RestorePostgresDatabaseBackup returns the given Postgres Database with the given Backup
// Deprecated: RestorePostgresDatabaseBackup is a deprecated method, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
func (c *Client) RestorePostgresDatabaseBackup(ctx context.Context, databaseID int, backupID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/backups/%d/restore", databaseID, backupID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// CreatePostgresDatabaseBackup creates a snapshot for the given Postgres database
// Deprecated: CreatePostgresDatabaseBackup is a deprecated method, as the backup endpoints are no longer supported in DBaaS V2.
// In DBaaS V2, databases can be backed up via database forking.
func (c *Client) CreatePostgresDatabaseBackup(ctx context.Context, databaseID int, opts PostgresBackupCreateOptions) error {
	e := formatAPIPath("databases/postgresql/instances/%d/backups", databaseID)
	return doPOSTRequestNoResponseBody(ctx, c, e, opts)
}

// SuspendPostgresDatabase suspends a PostgreSQL Managed Database, releasing idle resources and keeping only necessary data.
// All service data is lost if there are no backups available.
func (c *Client) SuspendPostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/suspend", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// ResumePostgresDatabase resumes a suspended PostgreSQL Managed Database
func (c *Client) ResumePostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/resume", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}
