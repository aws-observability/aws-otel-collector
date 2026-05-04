package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const databasePath = "/v2/databases"

// DatabaseService is the interface to interact with the Database endpoints on the Vultr API
// Link: https://www.vultr.com/api/#tag/managed-databases
type DatabaseService interface {
	ListPlans(ctx context.Context, options *DBPlanListOptions) ([]DatabasePlan, *Meta, *http.Response, error)

	List(ctx context.Context, options *DBListOptions) ([]Database, *Meta, *http.Response, error)
	Create(ctx context.Context, databaseReq *DatabaseCreateReq) (*Database, *http.Response, error)
	Get(ctx context.Context, databaseID string) (*Database, *http.Response, error)
	Update(ctx context.Context, databaseID string, databaseReq *DatabaseUpdateReq) (*Database, *http.Response, error)
	Delete(ctx context.Context, databaseID string) error

	GetUsage(ctx context.Context, databaseID string) (*DatabaseUsage, *http.Response, error)

	ListUsers(ctx context.Context, databaseID string) ([]DatabaseUser, *Meta, *http.Response, error)
	CreateUser(ctx context.Context, databaseID string, databaseUserReq *DatabaseUserCreateReq) (*DatabaseUser, *http.Response, error)
	GetUser(ctx context.Context, databaseID string, username string) (*DatabaseUser, *http.Response, error)
	UpdateUser(ctx context.Context, databaseID string, username string, databaseUserReq *DatabaseUserUpdateReq) (*DatabaseUser, *http.Response, error) //nolint:lll
	DeleteUser(ctx context.Context, databaseID string, username string) error
	UpdateUserACL(ctx context.Context, databaseID string, username string, databaseUserACLReq *DatabaseUserACLReq) (*DatabaseUser, *http.Response, error) //nolint:lll

	ListDBs(ctx context.Context, databaseID string) ([]DatabaseDB, *Meta, *http.Response, error)
	CreateDB(ctx context.Context, databaseID string, databaseDBReq *DatabaseDBCreateReq) (*DatabaseDB, *http.Response, error)
	GetDB(ctx context.Context, databaseID string, dbname string) (*DatabaseDB, *http.Response, error)
	DeleteDB(ctx context.Context, databaseID string, dbname string) error

	ListTopics(ctx context.Context, databaseID string) ([]DatabaseTopic, *Meta, *http.Response, error)
	CreateTopic(ctx context.Context, databaseID string, databaseTopicReq *DatabaseTopicCreateReq) (*DatabaseTopic, *http.Response, error)
	GetTopic(ctx context.Context, databaseID string, topicName string) (*DatabaseTopic, *http.Response, error)
	UpdateTopic(ctx context.Context, databaseID string, topicName string, databaseTopicReq *DatabaseTopicUpdateReq) (*DatabaseTopic, *http.Response, error) //nolint:lll
	DeleteTopic(ctx context.Context, databaseID string, topicName string) error

	ListQuotas(ctx context.Context, databaseID string) ([]DatabaseQuota, *Meta, *http.Response, error)
	CreateQuota(ctx context.Context, databaseID string, databaseQuotaReq *DatabaseQuotaCreateReq) (*DatabaseQuota, *http.Response, error)
	GetQuota(ctx context.Context, databaseID string, clientID, username string) (*DatabaseQuota, *http.Response, error)
	UpdateQuota(ctx context.Context, databaseID string, clientID, username string, databaseQuotaReq *DatabaseQuotaUpdateReq) (*DatabaseQuota, *http.Response, error) //nolint:lll
	DeleteQuota(ctx context.Context, databaseID string, clientID, username string) error

	ListMaintenanceUpdates(ctx context.Context, databaseID string) ([]string, *http.Response, error)
	StartMaintenance(ctx context.Context, databaseID string) (string, *http.Response, error)

	ListAvailableConnectors(ctx context.Context, databaseID string) ([]DatabaseAvailableConnector, *http.Response, error)
	GetConnectorConfigurationSchema(ctx context.Context, databaseID string, connectorClass string) ([]DatabaseConnectorConfigurationOption, *http.Response, error) //nolint:lll
	ListConnectors(ctx context.Context, databaseID string) ([]DatabaseConnector, *Meta, *http.Response, error)
	CreateConnector(ctx context.Context, databaseID string, databaseConnectorReq *DatabaseConnectorCreateReq) (*DatabaseConnector, *http.Response, error) //nolint:lll
	GetConnector(ctx context.Context, databaseID string, connectorName string) (*DatabaseConnector, *http.Response, error)
	UpdateConnector(ctx context.Context, databaseID string, connectorName string, databaseConnectorReq *DatabaseConnectorUpdateReq) (*DatabaseConnector, *http.Response, error) //nolint:lll
	DeleteConnector(ctx context.Context, databaseID string, connectorName string) error
	GetConnectorStatus(ctx context.Context, databaseID string, connectorName string) (*DatabaseConnectorStatus, *http.Response, error)
	RestartConnector(ctx context.Context, databaseID string, connectorName string) error
	PauseConnector(ctx context.Context, databaseID string, connectorName string) error
	ResumeConnector(ctx context.Context, databaseID string, connectorName string) error
	RestartConnectorTask(ctx context.Context, databaseID string, connectorName string, taskID int) error

	ListServiceAlerts(ctx context.Context, databaseID string, databaseAlertsReq *DatabaseListAlertsReq) ([]DatabaseAlert, *http.Response, error) //nolint:lll

	GetMigrationStatus(ctx context.Context, databaseID string) (*DatabaseMigration, *http.Response, error)
	StartMigration(ctx context.Context, databaseID string, databaseMigrationReq *DatabaseMigrationStartReq) (*DatabaseMigration, *http.Response, error) //nolint:lll
	DetachMigration(ctx context.Context, databaseID string) error

	AddReadOnlyReplica(ctx context.Context, databaseID string, databaseReplicaReq *DatabaseAddReplicaReq) (*Database, *http.Response, error)
	PromoteReadReplica(ctx context.Context, databaseID string) error

	GetBackupInformation(ctx context.Context, databaseID string) (*DatabaseBackups, *http.Response, error)
	RestoreFromBackup(ctx context.Context, databaseID string, databaseRestoreReq *DatabaseBackupRestoreReq) (*Database, *http.Response, error)
	Fork(ctx context.Context, databaseID string, databaseForkReq *DatabaseForkReq) (*Database, *http.Response, error)

	ListConnectionPools(ctx context.Context, databaseID string) (*DatabaseConnections, []DatabaseConnectionPool, *Meta, *http.Response, error)
	CreateConnectionPool(ctx context.Context, databaseID string, databaseConnectionPoolReq *DatabaseConnectionPoolCreateReq) (*DatabaseConnectionPool, *http.Response, error) //nolint:lll
	GetConnectionPool(ctx context.Context, databaseID string, poolName string) (*DatabaseConnectionPool, *http.Response, error)
	UpdateConnectionPool(ctx context.Context, databaseID string, poolName string, databaseConnectionPoolReq *DatabaseConnectionPoolUpdateReq) (*DatabaseConnectionPool, *http.Response, error) //nolint:lll
	DeleteConnectionPool(ctx context.Context, databaseID string, poolName string) error

	ListAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseAdvancedOptions, []AvailableOption, *http.Response, error)
	UpdateAdvancedOptions(ctx context.Context, databaseID string, databaseAdvancedOptionsReq *DatabaseAdvancedOptions) (*DatabaseAdvancedOptions, []AvailableOption, *http.Response, error)                                                         //nolint:lll
	ListKafkaRESTAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseKafkaRESTAdvancedOptions, []AvailableOption, *http.Response, error)                                                                                              //nolint:lll
	UpdateKafkaRESTAdvancedOptions(ctx context.Context, databaseID string, databaseKafkaRESTAdvancedOptionsReq *DatabaseKafkaRESTAdvancedOptions) (*DatabaseKafkaRESTAdvancedOptions, []AvailableOption, *http.Response, error)                     //nolint:lll
	ListSchemaRegistryAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseSchemaRegistryAdvancedOptions, []AvailableOption, *http.Response, error)                                                                                    //nolint:lll
	UpdateSchemaRegistryAdvancedOptions(ctx context.Context, databaseID string, databaseSchemaRegistryAdvancedOptionsReq *DatabaseSchemaRegistryAdvancedOptions) (*DatabaseSchemaRegistryAdvancedOptions, []AvailableOption, *http.Response, error) //nolint:lll
	ListKafkaConnectAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseKafkaConnectAdvancedOptions, []AvailableOption, *http.Response, error)                                                                                        //nolint:lll
	UpdateKafkaConnectAdvancedOptions(ctx context.Context, databaseID string, databaseKafkaConnectAdvancedOptionsReq *DatabaseKafkaConnectAdvancedOptions) (*DatabaseKafkaConnectAdvancedOptions, []AvailableOption, *http.Response, error)         //nolint:lll

	ListAvailableVersions(ctx context.Context, databaseID string) ([]string, *http.Response, error)
	StartVersionUpgrade(ctx context.Context, databaseID string, databaseVersionUpgradeReq *DatabaseVersionUpgradeReq) (string, *http.Response, error) //nolint:lll
}

// DatabaseServiceHandler handles interaction with the server methods for the Vultr API
type DatabaseServiceHandler struct {
	client *Client
}

// DBPlanListOptions handles GET request parameters for the ListPlans endpoint
type DBPlanListOptions struct {
	Engine string `url:"engine,omitempty"`
	Nodes  int    `url:"nodes,omitempty"`
	Region string `url:"region,omitempty"`
}

// DatabasePlan represents a Managed Database plan
type DatabasePlan struct {
	ID               string           `json:"id"`
	NumberOfNodes    int              `json:"number_of_nodes"`
	Type             string           `json:"type"`
	VCPUCount        int              `json:"vcpu_count"`
	RAM              int              `json:"ram"`
	Disk             int              `json:"disk"`
	MonthlyCost      int              `json:"monthly_cost"`
	SupportedEngines SupportedEngines `json:"supported_engines"`
	MaxConnections   *MaxConnections  `json:"max_connections,omitempty"`
	Locations        []string         `json:"locations"`
}

// SupportedEngines represents an object containing supported database engine types for Managed Database plans
type SupportedEngines struct {
	MySQL  *bool `json:"mysql"`
	PG     *bool `json:"pg"`
	Valkey *bool `json:"valkey"`
	Kafka  *bool `json:"kafka"`
}

// MaxConnections represents an object containing the maximum number of connections by engine type for Managed Database plans
type MaxConnections struct {
	MySQL int `json:"mysql,omitempty"`
	PG    int `json:"pg,omitempty"`
}

// databasePlansBase holds the entire ListPlans API response
type databasePlansBase struct {
	DatabasePlans []DatabasePlan `json:"plans"`
	Meta          *Meta          `json:"meta"`
}

// DBListOptions handles GET request parameters for the List endpoint
type DBListOptions struct {
	Label  string `url:"label,omitempty"`
	Tag    string `url:"tag,omitempty"`
	Region string `url:"region,omitempty"`
}

// Database represents a Managed Database subscription
type Database struct {
	ID                     string               `json:"id"`
	DateCreated            string               `json:"date_created"`
	Plan                   string               `json:"plan"`
	PlanDisk               int                  `json:"plan_disk"`
	PlanRAM                int                  `json:"plan_ram"`
	PlanVCPUs              int                  `json:"plan_vcpus"`
	PlanReplicas           *int                 `json:"plan_replicas,omitempty"`
	PlanBrokers            int                  `json:"plan_brokers,omitempty"`
	Region                 string               `json:"region"`
	DatabaseEngine         string               `json:"database_engine"`
	DatabaseEngineVersion  string               `json:"database_engine_version"`
	VPCID                  string               `json:"vpc_id"`
	Status                 string               `json:"status"`
	Label                  string               `json:"label"`
	Tag                    string               `json:"tag"`
	PendingCharges         float32              `json:"pending_charges"`
	DBName                 string               `json:"dbname,omitempty"`
	FerretDBCredentials    *FerretDBCredentials `json:"ferretdb_credentials,omitempty"`
	Host                   string               `json:"host"`
	PublicHost             string               `json:"public_host,omitempty"`
	Port                   string               `json:"port"`
	SASLPort               string               `json:"sasl_port,omitempty"`
	EnableKafkaREST        *bool                `json:"enable_kafka_rest,omitempty"`
	KafkaRESTURI           string               `json:"kafka_rest_uri,omitempty"`
	EnableSchemaRegistry   *bool                `json:"enable_schema_registry,omitempty"`
	SchemaRegistryURI      string               `json:"schema_registry_uri,omitempty"`
	EnableKafkaConnect     *bool                `json:"enable_kafka_connect,omitempty"`
	User                   string               `json:"user"`
	Password               string               `json:"password"`
	AccessKey              string               `json:"access_key,omitempty"`
	AccessCert             string               `json:"access_cert,omitempty"`
	MaintenanceDOW         string               `json:"maintenance_dow"`
	MaintenanceTime        string               `json:"maintenance_time"`
	BackupHour             *string              `json:"backup_hour,omitempty"`
	BackupMinute           *string              `json:"backup_minute,omitempty"`
	LatestBackup           string               `json:"latest_backup"`
	TrustedIPs             []string             `json:"trusted_ips"`
	CACertificate          string               `json:"ca_certificate"`
	MySQLSQLModes          []string             `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool                `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool                `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int                  `json:"mysql_long_query_time,omitempty"`
	PGAvailableExtensions  []PGExtension        `json:"pg_available_extensions,omitempty"`
	EvictionPolicy         string               `json:"eviction_policy,omitempty"`
	ClusterTimeZone        string               `json:"cluster_time_zone,omitempty"`
	ReadReplicas           []Database           `json:"read_replicas,omitempty"`
}

// FerretDBCredentials represents connection details and IP address information for FerretDB engine type subscriptions
type FerretDBCredentials struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	PublicIP  string `json:"public_ip"`
	PrivateIP string `json:"private_ip,omitempty"`
}

// PGExtension represents an object containing extension name and version information
type PGExtension struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}

// databasesBase holds the entire List API response
type databasesBase struct {
	Databases []Database `json:"databases"`
	Meta      *Meta      `json:"meta"`
}

// databaseBase holds the entire Get API response
type databaseBase struct {
	Database *Database `json:"database"`
}

// DatabaseCreateReq struct used to create a database
type DatabaseCreateReq struct {
	DatabaseEngine         string   `json:"database_engine,omitempty"`
	DatabaseEngineVersion  string   `json:"database_engine_version,omitempty"`
	Region                 string   `json:"region,omitempty"`
	Plan                   string   `json:"plan,omitempty"`
	Label                  string   `json:"label,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
	VPCID                  string   `json:"vpc_id,omitempty"`
	MaintenanceDOW         string   `json:"maintenance_dow,omitempty"`
	MaintenanceTime        string   `json:"maintenance_time,omitempty"`
	BackupHour             *string  `json:"backup_hour,omitempty"`
	BackupMinute           *string  `json:"backup_minute,omitempty"`
	TrustedIPs             []string `json:"trusted_ips,omitempty"`
	MySQLSQLModes          []string `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool    `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool    `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int      `json:"mysql_long_query_time,omitempty"`
	EvictionPolicy         string   `json:"eviction_policy,omitempty"`
	EnableKafkaREST        *bool    `json:"enable_kafka_rest,omitempty"`
	EnableSchemaRegistry   *bool    `json:"enable_schema_registry,omitempty"`
	EnableKafkaConnect     *bool    `json:"enable_kafka_connect,omitempty"`
}

// DatabaseUpdateReq struct used to update a database
type DatabaseUpdateReq struct {
	Region                 string   `json:"region,omitempty"`
	Plan                   string   `json:"plan,omitempty"`
	Label                  string   `json:"label,omitempty"`
	Tag                    string   `json:"tag,omitempty"`
	VPCID                  *string  `json:"vpc_id,omitempty"`
	MaintenanceDOW         string   `json:"maintenance_dow,omitempty"`
	MaintenanceTime        string   `json:"maintenance_time,omitempty"`
	BackupHour             *string  `json:"backup_hour,omitempty"`
	BackupMinute           *string  `json:"backup_minute,omitempty"`
	ClusterTimeZone        string   `json:"cluster_time_zone,omitempty"`
	TrustedIPs             []string `json:"trusted_ips,omitempty"`
	MySQLSQLModes          []string `json:"mysql_sql_modes,omitempty"`
	MySQLRequirePrimaryKey *bool    `json:"mysql_require_primary_key,omitempty"`
	MySQLSlowQueryLog      *bool    `json:"mysql_slow_query_log,omitempty"`
	MySQLLongQueryTime     int      `json:"mysql_long_query_time,omitempty"`
	EvictionPolicy         string   `json:"eviction_policy,omitempty"`
	EnableKafkaREST        *bool    `json:"enable_kafka_rest,omitempty"`
	EnableSchemaRegistry   *bool    `json:"enable_schema_registry,omitempty"`
	EnableKafkaConnect     *bool    `json:"enable_kafka_connect,omitempty"`
}

// DatabaseUsage represents disk, memory, and CPU usage for a Managed Database
type DatabaseUsage struct {
	Disk   DatabaseDiskUsage   `json:"disk"`
	Memory DatabaseMemoryUsage `json:"memory"`
	CPU    DatabaseCPUUsage    `json:"cpu"`
}

// DatabaseDiskUsage represents disk usage details for a Managed Database
type DatabaseDiskUsage struct {
	CurrentGB  float32 `json:"current_gb"`
	MaxGB      int     `json:"max_gb"`
	Percentage float32 `json:"percentage"`
}

// DatabaseMemoryUsage represents memory usage details for a Managed Database
type DatabaseMemoryUsage struct {
	CurrentMB  float32 `json:"current_mb"`
	MaxMB      int     `json:"max_mb"`
	Percentage float32 `json:"percentage"`
}

// DatabaseCPUUsage represents average CPU usage for a Managed Database
type DatabaseCPUUsage struct {
	Percentage float32 `json:"percentage"`
}

// databaseUsageBase represents a usage details API response for a Managed Database
type databaseUsageBase struct {
	Usage *DatabaseUsage `json:"usage"`
}

// DatabaseUser represents a user within a Managed Database cluster
type DatabaseUser struct {
	Username      string           `json:"username"`
	Password      string           `json:"password"`
	Encryption    string           `json:"encryption,omitempty"`
	AccessControl *DatabaseUserACL `json:"access_control,omitempty"`
	Permission    string           `json:"permission,omitempty"`
	AccessKey     string           `json:"access_key,omitempty"`
	AccessCert    string           `json:"access_cert,omitempty"`
}

// DatabaseUserACL represents an access control configuration for a user within a Valkey Managed Database cluster
type DatabaseUserACL struct {
	ACLCategories []string `json:"acl_categories"`
	ACLChannels   []string `json:"acl_channels"`
	ACLCommands   []string `json:"acl_commands"`
	ACLKeys       []string `json:"acl_keys"`
}

// DatabaseUserACLReq represents input for updating a user's access control within a Managed Database cluster
type DatabaseUserACLReq struct {
	ACLCategories *[]string `json:"acl_categories,omitempty"`
	ACLChannels   *[]string `json:"acl_channels,omitempty"`
	ACLCommands   *[]string `json:"acl_commands,omitempty"`
	ACLKeys       *[]string `json:"acl_keys,omitempty"`
	Permission    string    `json:"permission,omitempty"`
}

// databaseUserBase holds the API response for retrieving a single database user within a Managed Database
type databaseUserBase struct {
	DatabaseUser *DatabaseUser `json:"user"`
}

// databaseUsersBase holds the API response for retrieving a list of database users within a Managed Database
type databaseUsersBase struct {
	DatabaseUsers []DatabaseUser `json:"users"`
	Meta          *Meta          `json:"meta"`
}

// DatabaseUserCreateReq struct used to create a user within a Managed Database
type DatabaseUserCreateReq struct {
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Permission string `json:"permission,omitempty"`
}

// DatabaseUserUpdateReq struct used to update a user within a Managed Database
type DatabaseUserUpdateReq struct {
	Password string `json:"password"`
}

// DatabaseDB represents a logical database within a Managed Database cluster
type DatabaseDB struct {
	Name string `json:"name"`
}

// databaseDBBase holds the API response for retrieving a single logical database within a Managed Database
type databaseDBBase struct {
	DatabaseDB *DatabaseDB `json:"db"`
}

// databaseDBsBase holds the API response for retrieving a list of logical databases within a Managed Database
type databaseDBsBase struct {
	DatabaseDBs []DatabaseDB `json:"dbs"`
	Meta        *Meta        `json:"meta"`
}

// DatabaseDBCreateReq struct used to create a logical database within a Managed Database
type DatabaseDBCreateReq struct {
	Name string `json:"name"`
}

// DatabaseTopic represents a Kafka topic within a Managed Database cluster
type DatabaseTopic struct {
	Name           string `json:"name"`
	Partitions     int    `json:"partitions"`
	Replication    int    `json:"replication"`
	RetentionHours int    `json:"retention_hours"`
	RetentionBytes int    `json:"retention_bytes"`
}

// databaseTopicBase holds the API response for retrieving a single Kafka topic within a Managed Database
type databaseTopicBase struct {
	DatabaseTopic *DatabaseTopic `json:"topic"`
}

// databaseTopicsBase holds the API response for retrieving a list of Kafka topics within a Managed Database
type databaseTopicsBase struct {
	DatabaseTopics []DatabaseTopic `json:"topics"`
	Meta           *Meta           `json:"meta"`
}

// DatabaseTopicCreateReq struct used to create a Kafka topic within a Managed Database
type DatabaseTopicCreateReq struct {
	Name           string `json:"name"`
	Partitions     int    `json:"partitions"`
	Replication    int    `json:"replication"`
	RetentionHours int    `json:"retention_hours"`
	RetentionBytes int    `json:"retention_bytes"`
}

// DatabaseTopicUpdateReq struct used to update a Kafka topic within a Managed Database
type DatabaseTopicUpdateReq struct {
	Partitions     int `json:"partitions"`
	Replication    int `json:"replication"`
	RetentionHours int `json:"retention_hours"`
	RetentionBytes int `json:"retention_bytes"`
}

// DatabaseQuota represents a Kafka quota within a Managed Database cluster
type DatabaseQuota struct {
	ClientID          string `json:"client_id"`
	User              string `json:"user"`
	ConsumerByteRate  int    `json:"consumer_byte_rate"`
	ProducerByteRate  int    `json:"producer_byte_rate"`
	RequestPercentage int    `json:"request_percentage"`
}

// databaseQuotaBase holds the API response for retrieving a single Kafka quota within a Managed Database
type databaseQuotaBase struct {
	DatabaseQuota *DatabaseQuota `json:"quota"`
}

// databaseQuotasBase holds the API response for retrieving a list of Kafka quotas within a Managed Database
type databaseQuotasBase struct {
	DatabaseQuotas []DatabaseQuota `json:"quotas"`
	Meta           *Meta           `json:"meta"`
}

// DatabaseQuotaCreateReq struct used to create a Kafka quota within a Managed Database
type DatabaseQuotaCreateReq struct {
	ClientID          string `json:"client_id"`
	User              string `json:"user"`
	ConsumerByteRate  int    `json:"consumer_byte_rate"`
	ProducerByteRate  int    `json:"producer_byte_rate"`
	RequestPercentage int    `json:"request_percentage"`
}

// DatabaseQuotaUpdateReq struct used to update a Kafka quota within a Managed Database
type DatabaseQuotaUpdateReq struct {
	ConsumerByteRate  int `json:"consumer_byte_rate"`
	ProducerByteRate  int `json:"producer_byte_rate"`
	RequestPercentage int `json:"request_percentage"`
}

// DatabaseAvailableConnector represents an available Kafka connector within a Managed Database cluster
type DatabaseAvailableConnector struct {
	Class   string `json:"class"`
	Title   string `json:"title"`
	Version string `json:"version"`
	Type    string `json:"type"`
	DocURL  string `json:"doc_url"`
}

// databaseAvailableConnectorsBase holds the API response for retrieving a list of available Kafka connectors within a Managed Database
type databaseAvailableConnectorsBase struct {
	DatabaseAvailableConnectors []DatabaseAvailableConnector `json:"available_connectors"`
}

// DatabaseConnectorConfigurationOption represents a configuration option for a Kafka connector within a Managed Database cluster
type DatabaseConnectorConfigurationOption struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
}

// databaseConnectorConfigBase holds the API response for retrieving a configuration schema for a Kafka connector within a Managed Database
type databaseConnectorConfigBase struct {
	ConfigurationSchema []DatabaseConnectorConfigurationOption `json:"configuration_schema"`
}

// DatabaseConnector represents a Kafka connector within a Managed Database cluster
type DatabaseConnector struct {
	Name   string                 `json:"name"`
	Class  string                 `json:"class"`
	Topics string                 `json:"topics"`
	Config map[string]interface{} `json:"config"`
}

// databaseConnectorBase holds the API response for retrieving a single Kafka connector within a Managed Database
type databaseConnectorBase struct {
	DatabaseConnector *DatabaseConnector `json:"connector"`
}

// databaseConnectorsBase holds the API response for retrieving a list of Kafka connectors within a Managed Database
type databaseConnectorsBase struct {
	DatabaseConnectors []DatabaseConnector `json:"connectors"`
	Meta               *Meta               `json:"meta"`
}

// DatabaseConnectorCreateReq struct used to create a Kafka connector within a Managed Database
type DatabaseConnectorCreateReq struct {
	Name   string                 `json:"name"`
	Class  string                 `json:"class"`
	Topics string                 `json:"topics"`
	Config map[string]interface{} `json:"config,omitempty"`
}

// DatabaseConnectorUpdateReq struct used to update a Kafka connector within a Managed Database
type DatabaseConnectorUpdateReq struct {
	Topics string                 `json:"topics,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
}

// DatabaseConnector represents a Kafka connector status within a Managed Database cluster
type DatabaseConnectorStatus struct {
	State string                  `json:"state"`
	Tasks []DatabaseConnectorTask `json:"tasks"`
}

// DatabaseConnectorTask represents a Kafka connector task within a Managed Database cluster
type DatabaseConnectorTask struct {
	ID    int    `json:"id"`
	State string `json:"state"`
	Trace string `json:"trace"`
}

// databaseConnectorStatusBase holds the API response for retrieving a Kafka connector status within a Managed Database
type databaseConnectorStatusBase struct {
	ConnectorStatus *DatabaseConnectorStatus `json:"connector_status"`
}

// databaseUpdatesBase holds the API response for retrieving a list of available maintenance updates within a Managed Database
type databaseUpdatesBase struct {
	AvailableUpdates []string `json:"available_updates"`
}

// databaseMessage is a bsic object holding a return message for certain API endpoints
type databaseMessage struct {
	Message string `json:"message"`
}

// DatabaseAlert represents a service alert for a Managed Database cluster
type DatabaseAlert struct {
	Timestamp            string `json:"timestamp"`
	MessageType          string `json:"message_type"`
	Description          string `json:"description"`
	Recommendation       string `json:"recommendation,omitempty"`
	MaintenanceScheduled string `json:"maintenance_scheduled,omitempty"`
	ResourceType         string `json:"resource_type,omitempty"`
	TableCount           int    `json:"table_count,omitempty"`
}

// databaseAlertsBase holds the API response for querying service alerts within a Managed Database
type databaseAlertsBase struct {
	DatabaseAlerts []DatabaseAlert `json:"alerts"`
}

// DatabaseListAlertsReq struct used to query service alerts for a Managed Database
type DatabaseListAlertsReq struct {
	Period string `json:"period"`
}

// DatabaseMigration represents migration details for a Managed Database cluster
type DatabaseMigration struct {
	Status      string              `json:"status"`
	Method      string              `json:"method,omitempty"`
	Error       string              `json:"error,omitempty"`
	Credentials DatabaseCredentials `json:"credentials"`
}

// DatabaseCredentials represents migration credentials for migration within a Managed Database cluster
type DatabaseCredentials struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"database,omitempty"`
	IgnoredDatabases string `json:"ignored_databases,omitempty"`
	SSL              *bool  `json:"ssl"`
}

// databaseMigrationBase represents a migration status object API response for a Managed Database
type databaseMigrationBase struct {
	Migration *DatabaseMigration `json:"migration"`
}

// DatabaseMigrationStartReq struct used to start a migration for a Managed Database
type DatabaseMigrationStartReq struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"database,omitempty"`
	IgnoredDatabases string `json:"ignored_databases,omitempty"`
	SSL              *bool  `json:"ssl"`
}

// DatabaseAddReplicaReq struct used to add a read-only replica to a Managed Database
type DatabaseAddReplicaReq struct {
	Region string `json:"region,omitempty"`
	Label  string `json:"label,omitempty"`
}

// DatabaseBackups represents backup information for a Managed Database cluster
type DatabaseBackups struct {
	LatestBackup DatabaseBackup `json:"latest_backup,omitempty"`
	OldestBackup DatabaseBackup `json:"oldest_backup,omitempty"`
}

// DatabaseBackup represents individual backup details for a Managed Database cluster
type DatabaseBackup struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// DatabaseBackupRestoreReq struct used to restore the backup of a Managed Database to a new subscription
type DatabaseBackupRestoreReq struct {
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
	Date  string `json:"date,omitempty"`
	Time  string `json:"time,omitempty"`
}

// DatabaseForkReq struct used to fork a Managed Database to a new subscription from a backup
type DatabaseForkReq struct {
	Label  string `json:"label,omitempty"`
	Region string `json:"region,omitempty"`
	Plan   string `json:"plan,omitempty"`
	Type   string `json:"type,omitempty"`
	Date   string `json:"date,omitempty"`
	Time   string `json:"time,omitempty"`
}

// DatabaseConnectionPool represents a PostgreSQL connection pool within a Managed Database cluster
type DatabaseConnectionPool struct {
	Name     string `json:"name"`
	Database string `json:"database"`
	Username string `json:"username"`
	Mode     string `json:"mode"`
	Size     int    `json:"size"`
}

// DatabaseConnections represents a an object containing used and available connections for a PostgreSQL Managed Database cluster
type DatabaseConnections struct {
	Used      int `json:"used"`
	Available int `json:"available"`
	Max       int `json:"max"`
}

// databaseConnectionPoolBase represents the API response for retrieving a single connection pool for a PostgreSQL Managed Database
type databaseConnectionPoolBase struct {
	ConnectionPool *DatabaseConnectionPool `json:"connection_pool"`
}

// databaseConnectionPoolBase represents the API response for retrieving all connection pool information for a PostgreSQL Managed Database
type databaseConnectionPoolsBase struct {
	Connections     *DatabaseConnections     `json:"connections"`
	ConnectionPools []DatabaseConnectionPool `json:"connection_pools"`
	Meta            *Meta                    `json:"meta"`
}

// DatabaseConnectionPoolCreateReq struct used to create a connection pool within a PostgreSQL Managed Database
type DatabaseConnectionPoolCreateReq struct {
	Name     string `json:"name,omitempty"`
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Mode     string `json:"mode,omitempty"`
	Size     int    `json:"size,omitempty"`
}

// DatabaseConnectionPoolUpdateReq struct used to update a connection pool within a PostgreSQL Managed Database
type DatabaseConnectionPoolUpdateReq struct {
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Mode     string `json:"mode,omitempty"`
	Size     int    `json:"size,omitempty"`
}

// DatabaseAdvancedOptions represents user configurable advanced options within a Managed Database cluster
type DatabaseAdvancedOptions struct {
	AutovacuumAnalyzeScaleFactor                         float32 `json:"autovacuum_analyze_scale_factor,omitempty"`
	AutovacuumAnalyzeThreshold                           int     `json:"autovacuum_analyze_threshold,omitempty"`
	AutovacuumFreezeMaxAge                               int     `json:"autovacuum_freeze_max_age,omitempty"`
	AutovacuumMaxWorkers                                 int     `json:"autovacuum_max_workers,omitempty"`
	AutovacuumNaptime                                    int     `json:"autovacuum_naptime,omitempty"`
	AutovacuumVacuumCostDelay                            int     `json:"autovacuum_vacuum_cost_delay,omitempty"`
	AutovacuumVacuumCostLimit                            int     `json:"autovacuum_vacuum_cost_limit,omitempty"`
	AutovacuumVacuumScaleFactor                          float32 `json:"autovacuum_vacuum_scale_factor,omitempty"`
	AutovacuumVacuumThreshold                            int     `json:"autovacuum_vacuum_threshold,omitempty"`
	BGWRITERDelay                                        int     `json:"bgwriter_delay,omitempty"`
	BGWRITERFlushAFter                                   int     `json:"bgwriter_flush_after,omitempty"`
	BGWRITERLRUMaxPages                                  int     `json:"bgwriter_lru_maxpages,omitempty"`
	BGWRITERLRUMultiplier                                float32 `json:"bgwriter_lru_multiplier,omitempty"`
	DeadlockTimeout                                      int     `json:"deadlock_timeout,omitempty"`
	DefaultToastCompression                              string  `json:"default_toast_compression,omitempty"`
	IdleInTransactionSessionTimeout                      int     `json:"idle_in_transaction_session_timeout,omitempty"`
	Jit                                                  *bool   `json:"jit,omitempty"`
	LogAutovacuumMinDuration                             int     `json:"log_autovacuum_min_duration,omitempty"`
	LogErrorVerbosity                                    string  `json:"log_error_verbosity,omitempty"`
	LogLinePrefix                                        string  `json:"log_line_prefix,omitempty"`
	LogMinDurationStatement                              int     `json:"log_min_duration_statement,omitempty"`
	MaxFilesPerProcess                                   int     `json:"max_files_per_process,omitempty"`
	MaxLocksPerTransaction                               int     `json:"max_locks_per_transaction,omitempty"`
	MaxLogicalReplicationWorkers                         int     `json:"max_logical_replication_workers,omitempty"`
	MaxParallelWorkers                                   int     `json:"max_parallel_workers,omitempty"`
	MaxParallelWorkersPerGather                          int     `json:"max_parallel_workers_per_gather,omitempty"`
	MaxPredLocksPerTransaction                           int     `json:"max_pred_locks_per_transaction,omitempty"`
	MaxPreparedTransactions                              int     `json:"max_prepared_transactions,omitempty"`
	MaxReplicationSlots                                  int     `json:"max_replication_slots,omitempty"`
	MaxStackDepth                                        int     `json:"max_stack_depth,omitempty"`
	MaxStandbyArchiveDelay                               int     `json:"max_standby_archive_delay,omitempty"`
	MaxStandbyStreamingDelay                             int     `json:"max_standby_streaming_delay,omitempty"`
	MaxWalSenders                                        int     `json:"max_wal_senders,omitempty"`
	MaxWorkerProcesses                                   int     `json:"max_worker_processes,omitempty"`
	PGPartmanBGWInterval                                 int     `json:"pg_partman_bgw.interval,omitempty"`
	PGPartmanBGWRole                                     string  `json:"pg_partman_bgw.role,omitempty"`
	PGStateStatementsTrack                               string  `json:"pg_stat_statements.track,omitempty"`
	TempFileLimit                                        int     `json:"temp_file_limit,omitempty"`
	TrackActivityQuerySize                               int     `json:"track_activity_query_size,omitempty"`
	TrackCommitTimestamp                                 string  `json:"track_commit_timestamp,omitempty"`
	TrackFunctions                                       string  `json:"track_functions,omitempty"`
	TrackIOTiming                                        string  `json:"track_io_timing,omitempty"`
	WALSenderTImeout                                     int     `json:"wal_sender_timeout,omitempty"`
	WALWriterDelay                                       int     `json:"wal_writer_delay,omitempty"`
	ConnectTimeout                                       int     `json:"connect_timeout,omitempty"`
	GroupConcatMaxLen                                    int     `json:"group_concat_max_len,omitempty"`
	InnoDBChangeBufferMaxSize                            int     `json:"innodb_change_buffer_max_size,omitempty"`
	InnoDBFlushNeighbors                                 int     `json:"innodb_flush_neighbors,omitempty"`
	InnoDBFTMinTokenSize                                 int     `json:"innodb_ft_min_token_size,omitempty"`
	InnoDBFTServerStopwordTable                          string  `json:"innodb_ft_server_stopword_table,omitempty"`
	InnoDBLockWaitTimeout                                int     `json:"innodb_lock_wait_timeout,omitempty"`
	InnoDBLogBufferSize                                  int     `json:"innodb_log_buffer_size,omitempty"`
	InnoDBOnlineAlterLogMaxSize                          int     `json:"innodb_online_alter_log_max_size,omitempty"`
	InnoDBPrintAllDeadlocks                              *bool   `json:"innodb_print_all_deadlocks,omitempty"`
	InnoDBReadIOThreads                                  int     `json:"innodb_read_io_threads,omitempty"`
	InnoDBRollbackOnTimeout                              *bool   `json:"innodb_rollback_on_timeout,omitempty"`
	InnoDBThreadConcurrency                              int     `json:"innodb_thread_concurrency,omitempty"`
	InnoDBWriteIOThreads                                 int     `json:"innodb_write_io_threads,omitempty"`
	InteractiveTimeout                                   int     `json:"interactive_timeout,omitempty"`
	InternalTmpMemStorageEngine                          string  `json:"internal_tmp_mem_storage_engine,omitempty"`
	MaxAllowedPacket                                     int     `json:"max_allowed_packet,omitempty"`
	MaxHeapTableSize                                     int     `json:"max_heap_table_size,omitempty"`
	NetBufferLength                                      int     `json:"net_buffer_length,omitempty"`
	NetReadTimeout                                       int     `json:"net_read_timeout,omitempty"`
	NetWriteTimeout                                      int     `json:"net_write_timeout,omitempty"`
	SortBufferSize                                       int     `json:"sort_buffer_size,omitempty"`
	TmpTableSize                                         int     `json:"tmp_table_size,omitempty"`
	WaitTimeout                                          int     `json:"wait_timeout,omitempty"`
	CompressionType                                      string  `json:"compression_type,omitempty"`
	GroupInitialRebalanceDelayMS                         int     `json:"group_initial_rebalance_delay_ms,omitempty"`
	GroupMinSessinTimeoutMS                              int     `json:"group_min_session_timeout_ms,omitempty"`
	GroupMaxSessionTimeoutMS                             int     `json:"group_max_session_timeout_ms,omitempty"`
	ConnectionsMaxIdleMS                                 int     `json:"connections_max_idle_ms,omitempty"`
	MaxIncrementalFetchSessionCacheSlots                 int     `json:"max_incremental_fetch_session_cache_slots,omitempty"`
	MessageMaxBytes                                      int     `json:"message_max_bytes,omitempty"`
	OffsetsRetentionMinutes                              int     `json:"offsets_retention_minutes,omitempty"`
	LogCleanerDeleteRetentionMS                          int     `json:"log_cleaner_delete_retention_ms,omitempty"`
	LogCleanerMinCleanableRatio                          float32 `json:"log_cleaner_min_cleanable_ratio,omitempty"`
	LogCleanerMaxCompactionLagMS                         int     `json:"log_cleaner_max_compaction_lag_ms,omitempty"`
	LogCleanerMinCompactionLagMS                         int     `json:"log_cleaner_min_compaction_lag_ms,omitempty"`
	LogCleanupPolicy                                     string  `json:"log_cleanup_policy,omitempty"`
	LogFlushIntervalMessages                             int     `json:"log_flush_interval_messages,omitempty"`
	LogFlushIntervalMS                                   int     `json:"log_flush_interval_ms,omitempty"`
	LogIndexIntervalBytes                                int     `json:"log_index_interval_bytes,omitempty"`
	LogIndexSizeMaxBytes                                 int     `json:"log_index_size_max_bytes,omitempty"`
	LogLocalRetentionMS                                  int     `json:"log_local_retention_ms,omitempty"`
	LogLocalRetentionBytes                               int     `json:"log_local_retention_bytes,omitempty"`
	LogMessageDownconversionEnable                       *bool   `json:"log_message_downconversion_enable,omitempty"`
	LogMessageTimestampType                              string  `json:"log_message_timestamp_type,omitempty"`
	LogMessageTimestampDifferenceMaxMS                   int     `json:"log_message_timestamp_difference_max_ms,omitempty"`
	LogPreallocate                                       *bool   `json:"log_preallocate,omitempty"`
	LogRetentionBytes                                    int     `json:"log_retention_bytes,omitempty"`
	LogRetentionHours                                    int     `json:"log_retention_hours,omitempty"`
	LogRetentionMS                                       int     `json:"log_retention_ms,omitempty"`
	LogRollJitterMS                                      int     `json:"log_roll_jitter_ms,omitempty"`
	LogRollMS                                            int     `json:"log_roll_ms,omitempty"`
	LogSegmentBytes                                      int     `json:"log_segment_bytes,omitempty"`
	LogSegmentDeleteDelayMS                              int     `json:"log_segment_delete_delay_ms,omitempty"`
	AutoCreateTopicsEnable                               *bool   `json:"auto_create_topics_enable,omitempty"`
	MinInsyncReplicas                                    int     `json:"min_insync_replicas,omitempty"`
	NumPartitions                                        int     `json:"num_partitions,omitempty"`
	DefaultReplicationFactor                             int     `json:"default_replication_factor,omitempty"`
	ReplicaFetchMaxBytes                                 int     `json:"replica_fetch_max_bytes,omitempty"`
	ReplicaFetchResponseMaxBytes                         int     `json:"replica_fetch_response_max_bytes,omitempty"`
	MaxConnectionsPerIP                                  int     `json:"max_connections_per_ip,omitempty"`
	ProducerPurgatoryPurgeIntervalRequests               int     `json:"producer_purgatory_purge_interval_requests,omitempty"`
	SASLOauthbearerExpectedAudience                      string  `json:"sasl_oauthbearer_expected_audience,omitempty"`
	SASLOauthbearerExpectedIssuer                        string  `json:"sasl_oauthbearer_expected_issuer,omitempty"`
	SASLOauthbearerJWKSEndpointURL                       string  `json:"sasl_oauthbearer_jwks_endpoint_url,omitempty"`
	SASLOauthbearerSubClaimName                          string  `json:"sasl_oauthbearer_sub_claim_name,omitempty"`
	SocketRequestMaxBytes                                int     `json:"socket_request_max_bytes,omitempty"`
	TransactionStateLogSegmentBytes                      int     `json:"transaction_state_log_segment_bytes,omitempty"`
	TransactionRemoveExpiredTransactionCleanupIntervalMS int     `json:"transaction_remove_expired_transaction_cleanup_interval_ms,omitempty"`
	TransactionPartitionVerificationEnable               *bool   `json:"transaction_partition_verification_enable,omitempty"`
}

// AvailableOption represents an available advanced configuration option for a Managed Database cluster
type AvailableOption struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Enumerals []string `json:"enumerals,omitempty"`
	MinValue  *float32 `json:"min_value,omitempty"`
	MaxValue  *float32 `json:"max_value,omitempty"`
	AltValues []int    `json:"alt_values,omitempty"`
	Units     string   `json:"units,omitempty"`
}

// databaseAdvancedOptionsBase represents the API response for advanced configuration options for a Managed Database
type databaseAdvancedOptionsBase struct {
	ConfiguredOptions *DatabaseAdvancedOptions `json:"configured_options"`
	AvailableOptions  []AvailableOption        `json:"available_options"`
}

// DatabaseKafkaRESTAdvancedOptions represents user configurable Kafka REST advanced options within a Managed Database cluster
type DatabaseKafkaRESTAdvancedOptions struct {
	ProducerAcks              string `json:"producer_acks,omitempty"`
	ProducerCompressionType   string `json:"producer_compression_type,omitempty"`
	ProducerLingerMS          int    `json:"producer_linger_ms,omitempty"`
	ProducerMaxRequestSize    int    `json:"producer_max_request_size,omitempty"`
	ConsumerEnableAutoCommit  *bool  `json:"consumer_enable_auto_commit,omitempty"`
	ConsumerRequestMaxBytes   int    `json:"consumer_request_max_bytes,omitempty"`
	ConsumerRequestTimeoutMS  int    `json:"consumer_request_timeout_ms,omitempty"`
	NameStrategy              string `json:"name_strategy,omitempty"`
	NameStrategyValidation    *bool  `json:"name_strategy_validation,omitempty"`
	SimpleConsumerPoolSizeMax int    `json:"simpleconsumer_pool_size_max,omitempty"`
}

// databaseKafkaRESTAdvancedOptionsBase represents the API response for Kafka REST advanced configuration options for a Managed Database
type databaseKafkaRESTAdvancedOptionsBase struct {
	ConfiguredOptions *DatabaseKafkaRESTAdvancedOptions `json:"configured_options"`
	AvailableOptions  []AvailableOption                 `json:"available_options"`
}

// DatabaseSchemaRegistryAdvancedOptions represents user configurable Schema Registry advanced options within a Managed Database cluster
type DatabaseSchemaRegistryAdvancedOptions struct {
	LeaderEligibility       *bool `json:"leader_eligibility,omitempty"`
	SchemaReaderStrictMode  *bool `json:"schema_reader_strict_mode,omitempty"`
	RetriableErrorsSilenced *bool `json:"retriable_errors_silenced,omitempty"`
}

// databaseSchemaRegistryAdvancedOptionsBase represents the API response for Schema Registry advanced options for a Managed Database
type databaseSchemaRegistryAdvancedOptionsBase struct {
	ConfiguredOptions *DatabaseSchemaRegistryAdvancedOptions `json:"configured_options"`
	AvailableOptions  []AvailableOption                      `json:"available_options"`
}

// DatabaseKafkaConnectAdvancedOptions represents user configurable Kafka Connect advanced options within a Managed Database cluster
type DatabaseKafkaConnectAdvancedOptions struct {
	ConnectorClientConfigOverridePolicy string `json:"connector_client_config_override_policy,omitempty"`
	ConsumerAutoOffsetReset             string `json:"consumer_auto_offset_reset,omitempty"`
	ConsumerFetchMaxBytes               int    `json:"consumer_fetch_max_bytes,omitempty"`
	ConsumerIsolationLevel              string `json:"consumer_isolation_level,omitempty"`
	ConsumerMaxPartitionFetchBytes      int    `json:"consumer_max_partition_fetch_bytes,omitempty"`
	ConsumerMaxPollIntervalMS           int    `json:"consumer_max_poll_interval_ms,omitempty"`
	ConsumerMaxPollRecords              int    `json:"consumer_max_poll_records,omitempty"`
	OffsetFlushIntervalMS               int    `json:"offset_flush_interval_ms,omitempty"`
	OffsetFlushTimeoutMS                int    `json:"offset_flush_timeout_ms,omitempty"`
	ProducerBatchSize                   int    `json:"producer_batch_size,omitempty"`
	ProducerBufferMemory                int    `json:"producer_buffer_memory,omitempty"`
	ProducerCompressionType             string `json:"producer_compression_type,omitempty"`
	ProducerLingerMS                    int    `json:"producer_linger_ms,omitempty"`
	ProducerMaxRequestSize              int    `json:"producer_max_request_size,omitempty"`
	ScheduledRebalanceMaxDelayMS        int    `json:"scheduled_rebalance_max_delay_ms,omitempty"`
	SessionTimeoutMS                    int    `json:"session_timeout_ms,omitempty"`
}

// databaseKafkaConnectAdvancedOptionsBase represents the API response for Kafka Connect advanced options for a Managed Database
type databaseKafkaConnectAdvancedOptionsBase struct {
	ConfiguredOptions *DatabaseKafkaConnectAdvancedOptions `json:"configured_options"`
	AvailableOptions  []AvailableOption                    `json:"available_options"`
}

// DatabaseAvailableVersions represents available versions upgrades for a Managed Database cluster
type DatabaseAvailableVersions struct {
	AvailableVersions []string `json:"available_versions"`
}

// DatabaseVersionUpgradeReq struct used to initiate a version upgrade for a PostgreSQL Managed Database
type DatabaseVersionUpgradeReq struct {
	Version string `json:"version,omitempty"`
}

// ListPlans retrieves all database plans
func (d *DatabaseServiceHandler) ListPlans(ctx context.Context, options *DBPlanListOptions) ([]DatabasePlan, *Meta, *http.Response, error) {
	uri := fmt.Sprintf("%s/plans", databasePath)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	databasePlans := new(databasePlansBase)
	resp, err := d.client.DoWithContext(ctx, req, databasePlans)
	if err != nil {
		return nil, nil, nil, err
	}

	return databasePlans.DatabasePlans, databasePlans.Meta, resp, nil
}

// List retrieves all databases on your account
func (d *DatabaseServiceHandler) List(ctx context.Context, options *DBListOptions) ([]Database, *Meta, *http.Response, error) { //nolint:dupl,lll
	req, err := d.client.NewRequest(ctx, http.MethodGet, databasePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	databases := new(databasesBase)
	resp, err := d.client.DoWithContext(ctx, req, databases)
	if err != nil {
		return nil, nil, nil, err
	}

	return databases.Databases, databases.Meta, resp, nil
}

// Create will create a Managed Database with the given parameters
func (d *DatabaseServiceHandler) Create(ctx context.Context, databaseReq *DatabaseCreateReq) (*Database, *http.Response, error) {
	req, err := d.client.NewRequest(ctx, http.MethodPost, databasePath, databaseReq)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// Get will get a Managed Database with the given databaseID
func (d *DatabaseServiceHandler) Get(ctx context.Context, databaseID string) (*Database, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// Update will update a Managed Database with the given parameters
func (d *DatabaseServiceHandler) Update(ctx context.Context, databaseID string, databaseReq *DatabaseUpdateReq) (*Database, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseReq)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// Delete a Managed database, all data will be permanently lost
func (d *DatabaseServiceHandler) Delete(ctx context.Context, databaseID string) error {
	uri := fmt.Sprintf("%s/%s", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// GetUsage retrieves disk, memory, and CPU usage information for a Managed Database
func (d *DatabaseServiceHandler) GetUsage(ctx context.Context, databaseID string) (*DatabaseUsage, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/usage", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseUsage := new(databaseUsageBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUsage)
	if err != nil {
		return nil, nil, err
	}

	return databaseUsage.Usage, resp, nil
}

// ListUsers retrieves all database users on a Managed Database
func (d *DatabaseServiceHandler) ListUsers(ctx context.Context, databaseID string) ([]DatabaseUser, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/users", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseUsers := new(databaseUsersBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUsers)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseUsers.DatabaseUsers, databaseUsers.Meta, resp, nil
}

// CreateUser will create a user within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateUser(ctx context.Context, databaseID string, databaseUserReq *DatabaseUserCreateReq) (*DatabaseUser, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/users", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseUserReq)
	if err != nil {
		return nil, nil, err
	}

	databaseUser := new(databaseUserBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUser)
	if err != nil {
		return nil, nil, err
	}

	return databaseUser.DatabaseUser, resp, nil
}

// GetUser retrieves information on an individual user within a Managed Database based on a username and databaseID
func (d *DatabaseServiceHandler) GetUser(ctx context.Context, databaseID, username string) (*DatabaseUser, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseUser := new(databaseUserBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUser)
	if err != nil {
		return nil, nil, err
	}

	return databaseUser.DatabaseUser, resp, nil
}

// UpdateUser will update a user within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateUser(ctx context.Context, databaseID, username string, databaseUserReq *DatabaseUserUpdateReq) (*DatabaseUser, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseUserReq)
	if err != nil {
		return nil, nil, err
	}

	databaseUser := new(databaseUserBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUser)
	if err != nil {
		return nil, nil, err
	}

	return databaseUser.DatabaseUser, resp, nil
}

// DeleteUser will delete a user within a Managed Database
func (d *DatabaseServiceHandler) DeleteUser(ctx context.Context, databaseID, username string) error {
	uri := fmt.Sprintf("%s/%s/users/%s", databasePath, databaseID, username)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// UpdateUserACL will update a user's access control within the Valkey Managed Database
func (d *DatabaseServiceHandler) UpdateUserACL(ctx context.Context, databaseID, username string, databaseUserACLReq *DatabaseUserACLReq) (*DatabaseUser, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/users/%s/access-control", databasePath, databaseID, username)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseUserACLReq)
	if err != nil {
		return nil, nil, err
	}

	databaseUser := new(databaseUserBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUser)
	if err != nil {
		return nil, nil, err
	}

	return databaseUser.DatabaseUser, resp, nil
}

// ListDBs retrieves all logical databases on a Managed Database
func (d *DatabaseServiceHandler) ListDBs(ctx context.Context, databaseID string) ([]DatabaseDB, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/dbs", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseDBs := new(databaseDBsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseDBs)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseDBs.DatabaseDBs, databaseDBs.Meta, resp, nil
}

// CreateDB will create a logical database within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateDB(ctx context.Context, databaseID string, databaseDBReq *DatabaseDBCreateReq) (*DatabaseDB, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/dbs", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseDBReq)
	if err != nil {
		return nil, nil, err
	}

	databaseDB := new(databaseDBBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseDB)
	if err != nil {
		return nil, nil, err
	}

	return databaseDB.DatabaseDB, resp, nil
}

// GetDB retrieves information on an individual logical database within a Managed Database based on a dbname and databaseID
func (d *DatabaseServiceHandler) GetDB(ctx context.Context, databaseID, dbname string) (*DatabaseDB, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/dbs/%s", databasePath, databaseID, dbname)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseDB := new(databaseDBBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseDB)
	if err != nil {
		return nil, nil, err
	}

	return databaseDB.DatabaseDB, resp, nil
}

// DeleteDB will delete a user within a Managed Database
func (d *DatabaseServiceHandler) DeleteDB(ctx context.Context, databaseID, dbname string) error {
	uri := fmt.Sprintf("%s/%s/dbs/%s", databasePath, databaseID, dbname)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListTopics retrieves all Kafka topics on a Managed Database
func (d *DatabaseServiceHandler) ListTopics(ctx context.Context, databaseID string) ([]DatabaseTopic, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/topics", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseTopics := new(databaseTopicsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseTopics)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseTopics.DatabaseTopics, databaseTopics.Meta, resp, nil
}

// CreateTopic will create a Kafka topic within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateTopic(ctx context.Context, databaseID string, databaseTopicReq *DatabaseTopicCreateReq) (*DatabaseTopic, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/topics", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseTopicReq)
	if err != nil {
		return nil, nil, err
	}

	databaseTopic := new(databaseTopicBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseTopic)
	if err != nil {
		return nil, nil, err
	}

	return databaseTopic.DatabaseTopic, resp, nil
}

// GetTopic retrieves information on an individual Kafka topic within a Managed Database based on a topicName and databaseID
func (d *DatabaseServiceHandler) GetTopic(ctx context.Context, databaseID, topicName string) (*DatabaseTopic, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/topics/%s", databasePath, databaseID, topicName)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseTopic := new(databaseTopicBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseTopic)
	if err != nil {
		return nil, nil, err
	}

	return databaseTopic.DatabaseTopic, resp, nil
}

// UpdateTopic will update a Kafka topic within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateTopic(ctx context.Context, databaseID, topicName string, databaseTopicReq *DatabaseTopicUpdateReq) (*DatabaseTopic, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/topics/%s", databasePath, databaseID, topicName)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseTopicReq)
	if err != nil {
		return nil, nil, err
	}

	databaseTopic := new(databaseTopicBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseTopic)
	if err != nil {
		return nil, nil, err
	}

	return databaseTopic.DatabaseTopic, resp, nil
}

// DeleteTopic will delete a Kafka topic within a Managed Database
func (d *DatabaseServiceHandler) DeleteTopic(ctx context.Context, databaseID, topicName string) error {
	uri := fmt.Sprintf("%s/%s/topics/%s", databasePath, databaseID, topicName)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListQuotas retrieves all Kafka quotas on a Managed Database
func (d *DatabaseServiceHandler) ListQuotas(ctx context.Context, databaseID string) ([]DatabaseQuota, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/quotas", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseQuotas := new(databaseQuotasBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseQuotas)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseQuotas.DatabaseQuotas, databaseQuotas.Meta, resp, nil
}

// CreateQuota will create a Kafka quota within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateQuota(ctx context.Context, databaseID string, databaseQuotaReq *DatabaseQuotaCreateReq) (*DatabaseQuota, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/quotas", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseQuotaReq)
	if err != nil {
		return nil, nil, err
	}

	databaseQuota := new(databaseQuotaBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseQuota)
	if err != nil {
		return nil, nil, err
	}

	return databaseQuota.DatabaseQuota, resp, nil
}

// GetQuota retrieves information on an individual Kafka quota within a Managed Database based on a clientID and databaseID
func (d *DatabaseServiceHandler) GetQuota(ctx context.Context, databaseID, clientID, username string) (*DatabaseQuota, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/quotas/%s/%s", databasePath, databaseID, clientID, username)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseQuota := new(databaseQuotaBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseQuota)
	if err != nil {
		return nil, nil, err
	}

	return databaseQuota.DatabaseQuota, resp, nil
}

// UpdateQuota will update a Kafka quota within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateQuota(ctx context.Context, databaseID, clientID, username string, databaseQuotaReq *DatabaseQuotaUpdateReq) (*DatabaseQuota, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/quotas/%s/%s", databasePath, databaseID, clientID, username)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseQuotaReq)
	if err != nil {
		return nil, nil, err
	}

	databaseQuota := new(databaseQuotaBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseQuota)
	if err != nil {
		return nil, nil, err
	}

	return databaseQuota.DatabaseQuota, resp, nil
}

// DeleteQuota will delete a Kafka quota within a Managed Database
func (d *DatabaseServiceHandler) DeleteQuota(ctx context.Context, databaseID, clientID, username string) error {
	uri := fmt.Sprintf("%s/%s/quotas/%s/%s", databasePath, databaseID, clientID, username)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListMaintenanceUpdates retrieves all available maintenance updates for a Managed Database
func (d *DatabaseServiceHandler) ListMaintenanceUpdates(ctx context.Context, databaseID string) ([]string, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/maintenance", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseUpdates := new(databaseUpdatesBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseUpdates)
	if err != nil {
		return nil, nil, err
	}

	return databaseUpdates.AvailableUpdates, resp, nil
}

// StartMaintenance will start the maintenance update process for a Managed Database
func (d *DatabaseServiceHandler) StartMaintenance(ctx context.Context, databaseID string) (string, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/maintenance", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return "", nil, err
	}

	databaseUpdates := new(databaseMessage)
	resp, err := d.client.DoWithContext(ctx, req, databaseUpdates)
	if err != nil {
		return "", nil, err
	}

	return databaseUpdates.Message, resp, nil
}

// ListAvailableConnectors retrieves all available Kafka connectors for a Managed Database
func (d *DatabaseServiceHandler) ListAvailableConnectors(ctx context.Context, databaseID string) ([]DatabaseAvailableConnector, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/available-connectors", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseAvailableConnectors := new(databaseAvailableConnectorsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseAvailableConnectors)
	if err != nil {
		return nil, nil, err
	}

	return databaseAvailableConnectors.DatabaseAvailableConnectors, resp, nil
}

// GetConnectorConfigurationSchema retrieves all available configuration options for a Kafka connector on a Managed Database
func (d *DatabaseServiceHandler) GetConnectorConfigurationSchema(ctx context.Context, databaseID, connectorClass string) ([]DatabaseConnectorConfigurationOption, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/available-connectors/%s/configuration", databasePath, databaseID, connectorClass)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseConnectorConfigurationSchema := new(databaseConnectorConfigBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectorConfigurationSchema)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnectorConfigurationSchema.ConfigurationSchema, resp, nil
}

// ListConnectors retrieves all Kafka connectors on a Managed Database
func (d *DatabaseServiceHandler) ListConnectors(ctx context.Context, databaseID string) ([]DatabaseConnector, *Meta, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/connectors", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseConnectors := new(databaseConnectorsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectors)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseConnectors.DatabaseConnectors, databaseConnectors.Meta, resp, nil
}

// CreateConnector will create a Kafka connector within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateConnector(ctx context.Context, databaseID string, databaseConnectorReq *DatabaseConnectorCreateReq) (*DatabaseConnector, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connectors", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseConnectorReq)
	if err != nil {
		return nil, nil, err
	}

	databaseConnector := new(databaseConnectorBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnector)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnector.DatabaseConnector, resp, nil
}

// GetConnector retrieves information on an individual Kafka connector within a Managed Database based on a connectorName and databaseID
func (d *DatabaseServiceHandler) GetConnector(ctx context.Context, databaseID, connectorName string) (*DatabaseConnector, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connectors/%s", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseConnector := new(databaseConnectorBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnector)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnector.DatabaseConnector, resp, nil
}

// UpdateConnector will update a Kafka connector within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateConnector(ctx context.Context, databaseID, connectorName string, databaseConnectorReq *DatabaseConnectorUpdateReq) (*DatabaseConnector, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/connectors/%s", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseConnectorReq)
	if err != nil {
		return nil, nil, err
	}

	databaseConnector := new(databaseConnectorBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnector)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnector.DatabaseConnector, resp, nil
}

// DeleteConnector will delete a Kafka connector within a Managed Database
func (d *DatabaseServiceHandler) DeleteConnector(ctx context.Context, databaseID, connectorName string) error {
	uri := fmt.Sprintf("%s/%s/connectors/%s", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// GetConnectorStatus retrieves the status of a Kafka connector within a Managed Database based on a connectorName and databaseID
func (d *DatabaseServiceHandler) GetConnectorStatus(ctx context.Context, databaseID, connectorName string) (*DatabaseConnectorStatus, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connectors/%s/status", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseConnectorStatus := new(databaseConnectorStatusBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectorStatus)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnectorStatus.ConnectorStatus, resp, nil
}

// RestartConnector will restart a Kafka connector within a Managed Database
func (d *DatabaseServiceHandler) RestartConnector(ctx context.Context, databaseID, connectorName string) error {
	uri := fmt.Sprintf("%s/%s/connectors/%s/restart", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// PauseConnector will pause a Kafka connector within a Managed Database
func (d *DatabaseServiceHandler) PauseConnector(ctx context.Context, databaseID, connectorName string) error {
	uri := fmt.Sprintf("%s/%s/connectors/%s/pause", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ResumeConnector will resume a paused Kafka connector within a Managed Database
func (d *DatabaseServiceHandler) ResumeConnector(ctx context.Context, databaseID, connectorName string) error {
	uri := fmt.Sprintf("%s/%s/connectors/%s/resume", databasePath, databaseID, connectorName)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// RestartConnectorTask will restart a Kafka connector task within a Managed Database
func (d *DatabaseServiceHandler) RestartConnectorTask(ctx context.Context, databaseID, connectorName string, taskID int) error {
	uri := fmt.Sprintf("%s/%s/connectors/%s/tasks/%d/restart", databasePath, databaseID, connectorName, taskID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListServiceAlerts queries for service alerts for a Managed Database using the given parameters
func (d *DatabaseServiceHandler) ListServiceAlerts(ctx context.Context, databaseID string, databaseAlertsReq *DatabaseListAlertsReq) ([]DatabaseAlert, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/alerts", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseAlertsReq)
	if err != nil {
		return nil, nil, err
	}

	databaseAlerts := new(databaseAlertsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseAlerts)
	if err != nil {
		return nil, nil, err
	}

	return databaseAlerts.DatabaseAlerts, resp, nil
}

// GetMigrationStatus retrieves the migration status for a Managed Database
func (d *DatabaseServiceHandler) GetMigrationStatus(ctx context.Context, databaseID string) (*DatabaseMigration, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/migration", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseMigration := new(databaseMigrationBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseMigration)
	if err != nil {
		return nil, nil, err
	}

	return databaseMigration.Migration, resp, nil
}

// StartMigration will start a migration for a Managed Database using the given credentials
func (d *DatabaseServiceHandler) StartMigration(ctx context.Context, databaseID string, databaseMigrationReq *DatabaseMigrationStartReq) (*DatabaseMigration, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/migration", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseMigrationReq)
	if err != nil {
		return nil, nil, err
	}

	databaseMigration := new(databaseMigrationBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseMigration)
	if err != nil {
		return nil, nil, err
	}

	return databaseMigration.Migration, resp, nil
}

// DetachMigration will detach a migration from a Managed Database
func (d *DatabaseServiceHandler) DetachMigration(ctx context.Context, databaseID string) error {
	uri := fmt.Sprintf("%s/%s/migration", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// AddReadOnlyReplica will add a read-only replica node to a Managed Database with the given parameters
func (d *DatabaseServiceHandler) AddReadOnlyReplica(ctx context.Context, databaseID string, databaseReplicaReq *DatabaseAddReplicaReq) (*Database, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/read-replica", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseReplicaReq)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// PromoteReadReplica will promote a read-only replica to its own standalone Managed Database subscription
func (d *DatabaseServiceHandler) PromoteReadReplica(ctx context.Context, databaseID string) error {
	uri := fmt.Sprintf("%s/%s/promote-read-replica", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// GetBackupInformation retrieves backup information for a Managed Database
func (d *DatabaseServiceHandler) GetBackupInformation(ctx context.Context, databaseID string) (*DatabaseBackups, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/backups", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseBackups := new(DatabaseBackups)
	resp, err := d.client.DoWithContext(ctx, req, databaseBackups)
	if err != nil {
		return nil, nil, err
	}

	return databaseBackups, resp, nil
}

// RestoreFromBackup will create a new subscription of the same plan from a backup of a Managed Database using the given parameters
func (d *DatabaseServiceHandler) RestoreFromBackup(ctx context.Context, databaseID string, databaseRestoreReq *DatabaseBackupRestoreReq) (*Database, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/restore", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseRestoreReq)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// Fork will create a new subscription of any plan from a backup of a Managed Database using the given parameters
func (d *DatabaseServiceHandler) Fork(ctx context.Context, databaseID string, databaseForkReq *DatabaseForkReq) (*Database, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/fork", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseForkReq)
	if err != nil {
		return nil, nil, err
	}

	database := new(databaseBase)
	resp, err := d.client.DoWithContext(ctx, req, database)
	if err != nil {
		return nil, nil, err
	}

	return database.Database, resp, nil
}

// ListConnectionPools retrieves all connection pools within your PostgreSQL Managed Database
func (d *DatabaseServiceHandler) ListConnectionPools(ctx context.Context, databaseID string) (*DatabaseConnections, []DatabaseConnectionPool, *Meta, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connection-pools", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	databaseConnectionPools := new(databaseConnectionPoolsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectionPools)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return databaseConnectionPools.Connections, databaseConnectionPools.ConnectionPools, databaseConnectionPools.Meta, resp, nil
}

// CreateConnectionPool will create a connection pool within the PostgreSQL Managed Database with the given parameters
func (d *DatabaseServiceHandler) CreateConnectionPool(ctx context.Context, databaseID string, databaseConnectionPoolReq *DatabaseConnectionPoolCreateReq) (*DatabaseConnectionPool, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connection-pools", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseConnectionPoolReq)
	if err != nil {
		return nil, nil, err
	}

	databaseConnectionPool := new(databaseConnectionPoolBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectionPool)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnectionPool.ConnectionPool, resp, nil
}

// GetConnectionPool retrieves information on an individual connection pool
// within a PostgreSQL Managed Database based on a poolName and databaseID
func (d *DatabaseServiceHandler) GetConnectionPool(ctx context.Context, databaseID, poolName string) (*DatabaseConnectionPool, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/connection-pools/%s", databasePath, databaseID, poolName)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseConnectionPool := new(databaseConnectionPoolBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectionPool)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnectionPool.ConnectionPool, resp, nil
}

// UpdateConnectionPool will update a connection pool within the PostgreSQL Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateConnectionPool(ctx context.Context, databaseID, poolName string, databaseConnectionPoolReq *DatabaseConnectionPoolUpdateReq) (*DatabaseConnectionPool, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/connection-pools/%s", databasePath, databaseID, poolName)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseConnectionPoolReq)
	if err != nil {
		return nil, nil, err
	}

	databaseConnectionPool := new(databaseConnectionPoolBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseConnectionPool)
	if err != nil {
		return nil, nil, err
	}

	return databaseConnectionPool.ConnectionPool, resp, nil
}

// DeleteConnectionPool will delete a connection pool within a Managed Database
func (d *DatabaseServiceHandler) DeleteConnectionPool(ctx context.Context, databaseID, poolName string) error {
	uri := fmt.Sprintf("%s/%s/connection-pools/%s", databasePath, databaseID, poolName)

	req, err := d.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = d.client.DoWithContext(ctx, req, nil)
	return err
}

// ListAdvancedOptions retrieves all current and available advanced configuration options within a Managed Database
func (d *DatabaseServiceHandler) ListAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/advanced-options", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseAdvancedOptions := new(databaseAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseAdvancedOptions.ConfiguredOptions, databaseAdvancedOptions.AvailableOptions, resp, nil
}

// UpdateAdvancedOptions will update advanced configuration options within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateAdvancedOptions(ctx context.Context, databaseID string, databaseAdvancedOptionsReq *DatabaseAdvancedOptions) (*DatabaseAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/advanced-options", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseAdvancedOptionsReq)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseAdvancedOptions := new(databaseAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseAdvancedOptions.ConfiguredOptions, databaseAdvancedOptions.AvailableOptions, resp, nil
}

// ListKafkaRESTAdvancedOptions retrieves all current and available Kafka REST advanced configuration options within a Managed Database
func (d *DatabaseServiceHandler) ListKafkaRESTAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseKafkaRESTAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/advanced-options/kafka-rest", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseKafkaRESTAdvancedOptions := new(databaseKafkaRESTAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseKafkaRESTAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseKafkaRESTAdvancedOptions.ConfiguredOptions, databaseKafkaRESTAdvancedOptions.AvailableOptions, resp, nil
}

// UpdateKafkaRESTAdvancedOptions will update Kafka REST advanced configuration options within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateKafkaRESTAdvancedOptions(ctx context.Context, databaseID string, databaseKafkaRESTAdvancedOptionsReq *DatabaseKafkaRESTAdvancedOptions) (*DatabaseKafkaRESTAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/advanced-options/kafka-rest", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseKafkaRESTAdvancedOptionsReq)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseKafkaRESTAdvancedOptions := new(databaseKafkaRESTAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseKafkaRESTAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseKafkaRESTAdvancedOptions.ConfiguredOptions, databaseKafkaRESTAdvancedOptions.AvailableOptions, resp, nil
}

// ListSchemaRegistryAdvancedOptions retrieves all current and available Schema Registry advanced options within a Managed Database
func (d *DatabaseServiceHandler) ListSchemaRegistryAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseSchemaRegistryAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/advanced-options/schema-registry", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseSchemaRegistryAdvancedOptions := new(databaseSchemaRegistryAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseSchemaRegistryAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseSchemaRegistryAdvancedOptions.ConfiguredOptions, databaseSchemaRegistryAdvancedOptions.AvailableOptions, resp, nil
}

// UpdateSchemaRegistryAdvancedOptions will update Schema Registry advanced options within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateSchemaRegistryAdvancedOptions(ctx context.Context, databaseID string, databaseSchemaRegistryAdvancedOptionsReq *DatabaseSchemaRegistryAdvancedOptions) (*DatabaseSchemaRegistryAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/advanced-options/schema-registry", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseSchemaRegistryAdvancedOptionsReq)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseSchemaRegistryAdvancedOptions := new(databaseSchemaRegistryAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseSchemaRegistryAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseSchemaRegistryAdvancedOptions.ConfiguredOptions, databaseSchemaRegistryAdvancedOptions.AvailableOptions, resp, nil
}

// ListKafkaConnectAdvancedOptions retrieves all current and available Kafka Connect advanced options within a Managed Database
func (d *DatabaseServiceHandler) ListKafkaConnectAdvancedOptions(ctx context.Context, databaseID string) (*DatabaseKafkaConnectAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:dupl,lll
	uri := fmt.Sprintf("%s/%s/advanced-options/kafka-connect", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseKafkaConnectAdvancedOptions := new(databaseKafkaConnectAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseKafkaConnectAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseKafkaConnectAdvancedOptions.ConfiguredOptions, databaseKafkaConnectAdvancedOptions.AvailableOptions, resp, nil
}

// UpdateKafkaConnectAdvancedOptions will update Kafka Connect advanced options within a Managed Database with the given parameters
func (d *DatabaseServiceHandler) UpdateKafkaConnectAdvancedOptions(ctx context.Context, databaseID string, databaseKafkaConnectAdvancedOptionsReq *DatabaseKafkaConnectAdvancedOptions) (*DatabaseKafkaConnectAdvancedOptions, []AvailableOption, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/advanced-options/kafka-connect", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPut, uri, databaseKafkaConnectAdvancedOptionsReq)
	if err != nil {
		return nil, nil, nil, err
	}

	databaseKafkaConnectAdvancedOptions := new(databaseKafkaConnectAdvancedOptionsBase)
	resp, err := d.client.DoWithContext(ctx, req, databaseKafkaConnectAdvancedOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	return databaseKafkaConnectAdvancedOptions.ConfiguredOptions, databaseKafkaConnectAdvancedOptions.AvailableOptions, resp, nil
}

// ListAvailableVersions retrieves all available version upgrades for a Managed Database
func (d *DatabaseServiceHandler) ListAvailableVersions(ctx context.Context, databaseID string) ([]string, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/version-upgrade", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	databaseVersions := new(DatabaseAvailableVersions)
	resp, err := d.client.DoWithContext(ctx, req, databaseVersions)
	if err != nil {
		return nil, nil, err
	}

	return databaseVersions.AvailableVersions, resp, nil
}

// StartVersionUpgrade will start a migration for a Managed Database using the given credentials
func (d *DatabaseServiceHandler) StartVersionUpgrade(ctx context.Context, databaseID string, databaseVersionUpgradeReq *DatabaseVersionUpgradeReq) (string, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/version-upgrade", databasePath, databaseID)

	req, err := d.client.NewRequest(ctx, http.MethodPost, uri, databaseVersionUpgradeReq)
	if err != nil {
		return "", nil, err
	}

	databaseVersionUpgrade := new(databaseMessage)
	resp, err := d.client.DoWithContext(ctx, req, databaseVersionUpgrade)
	if err != nil {
		return "", nil, err
	}

	return databaseVersionUpgrade.Message, resp, nil
}
