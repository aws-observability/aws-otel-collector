# CHANGELOG

## 2.52.0/2025-12-17

### Changed
* Mark Incident Impact Endpoints stable [#3530](https://github.com/DataDog/datadog-api-client-go/pull/3530)
* Vulnerability Management - Update ListVulnerabilities endpoint query params and response schema [#3521](https://github.com/DataDog/datadog-api-client-go/pull/3521)
* Update specs for AWS account updates and creation for CCM configs [#3514](https://github.com/DataDog/datadog-api-client-go/pull/3514)
* add suppression version history [#3490](https://github.com/DataDog/datadog-api-client-go/pull/3490)
* Add processors groups to Observability Pipelines [#3450](https://github.com/DataDog/datadog-api-client-go/pull/3450)

### Added
* Update security finding triage specs [#3522](https://github.com/DataDog/datadog-api-client-go/pull/3522)
* Add routes for managing On-Call user notification channels [#3517](https://github.com/DataDog/datadog-api-client-go/pull/3517)
* Add host field to Post an event V2 API [#3515](https://github.com/DataDog/datadog-api-client-go/pull/3515)
* Add `GET /api/v2/apm/services` endpoint to public documentation [#3512](https://github.com/DataDog/datadog-api-client-go/pull/3512)
* Dashboards - Add semantic_mode support to FormulaAndFunctionMetricQueryDefinition [#3506](https://github.com/DataDog/datadog-api-client-go/pull/3506)
* Cloud SIEM - Add instantaneousBaseline feature parameter. [#3503](https://github.com/DataDog/datadog-api-client-go/pull/3503)
* Add new fields to usage metering api spec [#3501](https://github.com/DataDog/datadog-api-client-go/pull/3501)
* update geomap widget definition about `conditional_format` and `text_formats` and view focus [#3495](https://github.com/DataDog/datadog-api-client-go/pull/3495)
* Add new endpoint for listing rules for a gate [#3481](https://github.com/DataDog/datadog-api-client-go/pull/3481)

### Removed
* Tag security findings Jira endpoints as unstable [#3510](https://github.com/DataDog/datadog-api-client-go/pull/3510)

### Fixed
* Add field attribute to the Workload Protection hash action [#3487](https://github.com/DataDog/datadog-api-client-go/pull/3487)

## 2.51.0/2025-12-08

### Added
* On-Call Add positioned schedule policy target [#3491](https://github.com/DataDog/datadog-api-client-go/pull/3491)
* Introduced new APIs to manage team hierarchy links [#3482](https://github.com/DataDog/datadog-api-client-go/pull/3482)
* Add Row Update Endpoints to Reference Tables API spec [#3472](https://github.com/DataDog/datadog-api-client-go/pull/3472)
* Add incident management seats  to spec [#3469](https://github.com/DataDog/datadog-api-client-go/pull/3469)
* Support provisioning teams from external sources [#3467](https://github.com/DataDog/datadog-api-client-go/pull/3467)
* security_monitoring - Add signalOutput field to ThreatHuntingJobResponseAttributes schema [#3465](https://github.com/DataDog/datadog-api-client-go/pull/3465)
* Add filter.scope to Monitor Notification Rules [#3462](https://github.com/DataDog/datadog-api-client-go/pull/3462)
* Add Support for Monitor Assets [#3452](https://github.com/DataDog/datadog-api-client-go/pull/3452)
* Add api specs for deployment gates [#3412](https://github.com/DataDog/datadog-api-client-go/pull/3412)

### Changed
* Add Security Finding Ticketing endpoints [#3485](https://github.com/DataDog/datadog-api-client-go/pull/3485)
* Flatten file_metadata response schema to avoid OneOf unmarshaling issues [#3471](https://github.com/DataDog/datadog-api-client-go/pull/3471)

### Fixed
* obs_pipelines: make google auth optional [#3476](https://github.com/DataDog/datadog-api-client-go/pull/3476)

## 2.50.0/2025-11-14

### Added
* Add suppression tags [#3456](https://github.com/DataDog/datadog-api-client-go/pull/3456)
* Add Team Connection API Documentation [#3454](https://github.com/DataDog/datadog-api-client-go/pull/3454)
* Add new summary keys for new standalone billing dimensions [#3451](https://github.com/DataDog/datadog-api-client-go/pull/3451)
* Add Bits AI Investigations and On-Call to  API specs [#3447](https://github.com/DataDog/datadog-api-client-go/pull/3447)
* Add `PreviewCatalogEntities` [#3444](https://github.com/DataDog/datadog-api-client-go/pull/3444)
* Sync 'audience_management.yaml' files with backend [#3441](https://github.com/DataDog/datadog-api-client-go/pull/3441)
* Dashboards - Add on_call_events datasources [#3440](https://github.com/DataDog/datadog-api-client-go/pull/3440)
* Add last_login_time to Users v2 API [#3399](https://github.com/DataDog/datadog-api-client-go/pull/3399)

### Deprecated
* [api-spec] Mark PATCH /api/v2/incidents/incident_id/attachments endpoint as deprecated [#3453](https://github.com/DataDog/datadog-api-client-go/pull/3453)
* [METEXP-2068] Deprecate api/v1/search Endpoint [#3448](https://github.com/DataDog/datadog-api-client-go/pull/3448)

## 2.49.0/2025-10-30

### Added
* 📝 [LOGSAC-1298] Add logs-pipeline type to restriction policy OpenAPI spec [#3434](https://github.com/DataDog/datadog-api-client-go/pull/3434)
* Security Monitoring - Update Signal Archive Reasons [#3433](https://github.com/DataDog/datadog-api-client-go/pull/3433)
* Add endpoints for Software Composition Analysis [#3430](https://github.com/DataDog/datadog-api-client-go/pull/3430)
* Add New Serverless Summary Entries to Api Spec [#3418](https://github.com/DataDog/datadog-api-client-go/pull/3418)
* Add metric namespace config filters to V2 GCP API [#3417](https://github.com/DataDog/datadog-api-client-go/pull/3417)
* Add specs for v2 eventbridge API [#3414](https://github.com/DataDog/datadog-api-client-go/pull/3414)
* Add support for `Schema Processor` in `Logs Pipelines` [#3411](https://github.com/DataDog/datadog-api-client-go/pull/3411)
* Add Static Analysis Rules Endpoints [#3395](https://github.com/DataDog/datadog-api-client-go/pull/3395)

### Changed
* Rename historical job API endpoints to threat hunting [#3431](https://github.com/DataDog/datadog-api-client-go/pull/3431)

## 2.48.0/2025-10-23

### Changed
* Include mention to new scanned-assets-metadata endpoint on container images v1 endpoint [#3410](https://github.com/DataDog/datadog-api-client-go/pull/3410)
* Include mention to new scanned-assets-metadata endpoint on hosts v1 endpoint [#3409](https://github.com/DataDog/datadog-api-client-go/pull/3409)
* security_monitoring - Add indexes to deprecate index in ruleQuery [#3402](https://github.com/DataDog/datadog-api-client-go/pull/3402)
* Add support for vulnerability management - Add ListScannedAssetsMetadata new endpoint and update existing ones [#3400](https://github.com/DataDog/datadog-api-client-go/pull/3400)
* Update description, operationId and examples for tag pipeline and custom allocation rules [#3397](https://github.com/DataDog/datadog-api-client-go/pull/3397)

### Fixed
* Update the summary name for get a tag pipeline ruleset. [#3405](https://github.com/DataDog/datadog-api-client-go/pull/3405)

### Added
* Add new DeleteAssignee endpoint to Error Tracking APIs [#3404](https://github.com/DataDog/datadog-api-client-go/pull/3404)
* document agentless GCP scan options CRUD endpoints [#3401](https://github.com/DataDog/datadog-api-client-go/pull/3401)
* Document `/api/v2/roles/templates`  [#3390](https://github.com/DataDog/datadog-api-client-go/pull/3390)
* Add Reference Tables API spec [#3389](https://github.com/DataDog/datadog-api-client-go/pull/3389)
* Add blockedRequestPatterns to synthetics browser test options [#3383](https://github.com/DataDog/datadog-api-client-go/pull/3383)
* Add `BulkDeleteDatastoreItems` to Datastore API spec [#3382](https://github.com/DataDog/datadog-api-client-go/pull/3382)
* Add some missing Workload Protection agent rule fields [#3381](https://github.com/DataDog/datadog-api-client-go/pull/3381)
* Add AzureScanOptions to agentless scanning API [#3379](https://github.com/DataDog/datadog-api-client-go/pull/3379)
* Add conditional recipients to notification rule [#3375](https://github.com/DataDog/datadog-api-client-go/pull/3375)
* Update ci_app description with max 1 year event run duration restriction [#3358](https://github.com/DataDog/datadog-api-client-go/pull/3358)
* Documenting the new Flaky Test Management API endpoint for public beta [#3347](https://github.com/DataDog/datadog-api-client-go/pull/3347)
* Document multiple case-management endpoints [#3273](https://github.com/DataDog/datadog-api-client-go/pull/3273)

## 2.47.0/2025-09-30

### Added
* Add API Key ID to rum application response [#3377](https://github.com/DataDog/datadog-api-client-go/pull/3377)
* Add suppression list query string parameter [#3376](https://github.com/DataDog/datadog-api-client-go/pull/3376)
* Add datastore trigger to workflows public API  [#3373](https://github.com/DataDog/datadog-api-client-go/pull/3373)
* Add Google PubSub destination to the Observability Pipelines API [#3364](https://github.com/DataDog/datadog-api-client-go/pull/3364)
* Add API spec for AWS Integrations standard and resource collection IAM permissions [#3362](https://github.com/DataDog/datadog-api-client-go/pull/3362)
* Publish new incident impact APIs [#3355](https://github.com/DataDog/datadog-api-client-go/pull/3355)
* Add product analytics to FormulaAndFunctionEventsDataSource [#3354](https://github.com/DataDog/datadog-api-client-go/pull/3354)
* Add sequence detection to security monitoring rules [#3348](https://github.com/DataDog/datadog-api-client-go/pull/3348)
* Update ci_app description with max 1 year event run duration restriction [#3333](https://github.com/DataDog/datadog-api-client-go/pull/3333)
* Add Public Delete Dora Events Endpoints [#3330](https://github.com/DataDog/datadog-api-client-go/pull/3330)

### Fixed
* Remove any references to synthetics test suites [#3368](https://github.com/DataDog/datadog-api-client-go/pull/3368)

### Changed
* Add tag pipeline, custom allocation rule and get cloud account by id for AWS/Azure/GCP configs APIs [#3353](https://github.com/DataDog/datadog-api-client-go/pull/3353)

## 2.46.0/2025-09-15

### Added
* Add Query Parameters to ListOrgConnections Endpoint [#3346](https://github.com/DataDog/datadog-api-client-go/pull/3346)
* Add Incident Notification Rules Public Spec [#3340](https://github.com/DataDog/datadog-api-client-go/pull/3340)
* Update v1 and v2 GCP API specs to support `monitored_resource_configs` [#3338](https://github.com/DataDog/datadog-api-client-go/pull/3338)
* Add action datastore API [#3315](https://github.com/DataDog/datadog-api-client-go/pull/3315)
* Security Monitoring - Make hasOptionalGroupByFields updatable [#3274](https://github.com/DataDog/datadog-api-client-go/pull/3274)

### Deprecated
* Promote unstable aws v2 APIs and deprecate v1 [#3337](https://github.com/DataDog/datadog-api-client-go/pull/3337)

### Changed
* Allow to send batches of events in pipelines API [#3319](https://github.com/DataDog/datadog-api-client-go/pull/3319)

## 2.45.0/2025-09-09

### Added
* Add Incident Notification Template Public Docs [#3332](https://github.com/DataDog/datadog-api-client-go/pull/3332)
* Add Cross Org API to Open API specs [#3331](https://github.com/DataDog/datadog-api-client-go/pull/3331)
* Add readonly ID of synthetics test steps [#3326](https://github.com/DataDog/datadog-api-client-go/pull/3326)
* Create Cloud SIEM histsignals endpoints [#3325](https://github.com/DataDog/datadog-api-client-go/pull/3325)
* Security Monitoring - Validation Endpoint for Suppressions [#3322](https://github.com/DataDog/datadog-api-client-go/pull/3322)
* Update Get All Notification Rules API docs to include pagination, sorting, and filtering params [#3320](https://github.com/DataDog/datadog-api-client-go/pull/3320)
* Security Monitoring - Related Suppressions for a Rule [#3318](https://github.com/DataDog/datadog-api-client-go/pull/3318)
* Extend Widget time schema with support for hide_incomplete_cost_data [#3307](https://github.com/DataDog/datadog-api-client-go/pull/3307)
* Add SDS rule `should_save_match` field [#3305](https://github.com/DataDog/datadog-api-client-go/pull/3305)
* Add spec for Agentless GetAwsScanOptions [#3302](https://github.com/DataDog/datadog-api-client-go/pull/3302)
* Add Cross Org API to Open API specs [#3300](https://github.com/DataDog/datadog-api-client-go/pull/3300)
* Add DNAP Spark Pod Autosizing service to API client [#3296](https://github.com/DataDog/datadog-api-client-go/pull/3296)
* Add version parameter to synthetic test trigger ci endpoint [#3295](https://github.com/DataDog/datadog-api-client-go/pull/3295)
* Document Error Tracking public APIs [#3292](https://github.com/DataDog/datadog-api-client-go/pull/3292)
* Add docs for 404 not found error in cost-onboarding-api [#3287](https://github.com/DataDog/datadog-api-client-go/pull/3287)

### Fixed
* Security Monitoring - Fix payload of Validation Endpoint for Suppressions [#3327](https://github.com/DataDog/datadog-api-client-go/pull/3327)
* [CCA-938][CCC-883] Audit existing CCM endpoints in OpenAPI spec [#3283](https://github.com/DataDog/datadog-api-client-go/pull/3283)
* Add enum Dataset type to Dataset API spec [#3281](https://github.com/DataDog/datadog-api-client-go/pull/3281)

### Changed
* Update public cost permissions [#3304](https://github.com/DataDog/datadog-api-client-go/pull/3304)
* Add Product Scales support to RUM v2 Applications API [#3285](https://github.com/DataDog/datadog-api-client-go/pull/3285)

## 2.44.0/2025-08-12

### Added
* Add Flex_Logs_Compute_XL to API Spec [#3266](https://github.com/DataDog/datadog-api-client-go/pull/3266)
* Support Host and IaC finding types in security notifications  [#3265](https://github.com/DataDog/datadog-api-client-go/pull/3265)
* New keys for summary public endpoint for Event Management Correlation product [#3261](https://github.com/DataDog/datadog-api-client-go/pull/3261)
* Add log autosubscription tag filters config to aws v2 API [#3257](https://github.com/DataDog/datadog-api-client-go/pull/3257)
* Extended List Findings API to expose resource related Private IP Addresses to details [#3250](https://github.com/DataDog/datadog-api-client-go/pull/3250)
* update metrics.yaml for ListMetricAssets and include Dashboard info [#3245](https://github.com/DataDog/datadog-api-client-go/pull/3245)
* Support Cloud SIEM scheduled rules in API client [#3242](https://github.com/DataDog/datadog-api-client-go/pull/3242)
* Uncomment edit dataset block, add dataset limitations into endpoint descriptions  [#3240](https://github.com/DataDog/datadog-api-client-go/pull/3240)
* Add `text` field in synthetics search endpoint [#3239](https://github.com/DataDog/datadog-api-client-go/pull/3239)
* Adding all action connection types to public API [#3238](https://github.com/DataDog/datadog-api-client-go/pull/3238)
* Document case management attributes endpoints [#3236](https://github.com/DataDog/datadog-api-client-go/pull/3236)
* add AP2 endpoint for On-Call Paging [#3232](https://github.com/DataDog/datadog-api-client-go/pull/3232)
* Flag IP case action [#3230](https://github.com/DataDog/datadog-api-client-go/pull/3230)
* Add DNS specs for Cloud Network Monitoring API [#3228](https://github.com/DataDog/datadog-api-client-go/pull/3228)
* Adding Datadog Connection to Connection API [#3222](https://github.com/DataDog/datadog-api-client-go/pull/3222)

### Fixed
* Split Dataset into separate request and response objects, mark unstable [#3249](https://github.com/DataDog/datadog-api-client-go/pull/3249)
* Disables some tests to avoid fails as the service is disabled [#3244](https://github.com/DataDog/datadog-api-client-go/pull/3244)
* OP make 'support_rules' field in parse_grok processor optional [#3233](https://github.com/DataDog/datadog-api-client-go/pull/3233)

### Deprecated
* Deprecate signals triage v1 endpoints [#3246](https://github.com/DataDog/datadog-api-client-go/pull/3246)

## 2.43.0/2025-07-14

### Added
* Add Datasets API to Open API Spec  [#3210](https://github.com/DataDog/datadog-api-client-go/pull/3210)
* Add support for vulnerability management - GetSBOMsList new endpoint and update existing ones [#3209](https://github.com/DataDog/datadog-api-client-go/pull/3209)
* Add spreadsheet to restriction_policy specs [#3203](https://github.com/DataDog/datadog-api-client-go/pull/3203)
* Adds missing /api/v1/synthetics/tests/search spec [#3196](https://github.com/DataDog/datadog-api-client-go/pull/3196)
* Add API spec for AWS Integrations IAM permissions [#3190](https://github.com/DataDog/datadog-api-client-go/pull/3190)
* New keys added for multiple products [#3188](https://github.com/DataDog/datadog-api-client-go/pull/3188)
* Add extractFromEmailBody step for synthetics browser test [#3185](https://github.com/DataDog/datadog-api-client-go/pull/3185)
* Add support for `Array Processor` in `Logs Pipelines` [#3183](https://github.com/DataDog/datadog-api-client-go/pull/3183)
* Document Cloud Cost Management GCP endpoints publicly [#3132](https://github.com/DataDog/datadog-api-client-go/pull/3132)

### Changed
* Update template variable schemas with type field in dashboards and shared dashboards endpoints for group by template variables [#3184](https://github.com/DataDog/datadog-api-client-go/pull/3184)

## 2.42.0/2025-06-30

### Fixed
* Synthetics mobile test `message` field is now required [#3182](https://github.com/DataDog/datadog-api-client-go/pull/3182)
* Make dns port be string and number [#3166](https://github.com/DataDog/datadog-api-client-go/pull/3166)

### Security
* Remove caseIndex from historical jobs api spec [#3181](https://github.com/DataDog/datadog-api-client-go/pull/3181)

### Changed
* Update events intake specs for v2 Events post endpoint [#3177](https://github.com/DataDog/datadog-api-client-go/pull/3177)

### Added
* Update Incident API specs to include `is_test` in `POST /incidents` and incidents response [#3176](https://github.com/DataDog/datadog-api-client-go/pull/3176)
* Add App Key Registration API  [#3170](https://github.com/DataDog/datadog-api-client-go/pull/3170)
* Add Monitor Template API [#3110](https://github.com/DataDog/datadog-api-client-go/pull/3110)

### Deprecated
* Deprecate SLO metadata fields in api spec [#3137](https://github.com/DataDog/datadog-api-client-go/pull/3137)

## 2.41.0/2025-06-24

### Fixed
* Fix basic auth requirements [#3163](https://github.com/DataDog/datadog-api-client-go/pull/3163)

### Added
* Microsoft Sentinel Public API support [#3161](https://github.com/DataDog/datadog-api-client-go/pull/3161)
* Add the AP2 datacenter. [#3159](https://github.com/DataDog/datadog-api-client-go/pull/3159)

## 2.40.0/2025-06-23

### Fixed
* Fix basic auth requirements [#3163](https://github.com/DataDog/datadog-api-client-go/pull/3163)
* Add support for the api_security detection rule type [#3158](https://github.com/DataDog/datadog-api-client-go/pull/3158)

### Added
* Microsoft Sentinel Public API support [#3161](https://github.com/DataDog/datadog-api-client-go/pull/3161)
* Add custom fields to Rule update/validate API public documentation. [#3150](https://github.com/DataDog/datadog-api-client-go/pull/3150)
* Add hash field to actions in CWS agent rules [#3147](https://github.com/DataDog/datadog-api-client-go/pull/3147)
* SDCD-1142: adding `custom_tags` optional attribute to DORA API spec [#3134](https://github.com/DataDog/datadog-api-client-go/pull/3134)
* Add sampling fields to SDS spec [#3129](https://github.com/DataDog/datadog-api-client-go/pull/3129)
* Add API spec for team hierarchy APIs [#3105](https://github.com/DataDog/datadog-api-client-go/pull/3105)

### Changed
* Update events intake specs for v2 Events post endpoint [#3144](https://github.com/DataDog/datadog-api-client-go/pull/3144)

## 2.39.0/2025-06-16

### Fixed
* sanitize configuration keys correctly [#3143](https://github.com/DataDog/datadog-api-client-go/pull/3143)

### Changed
* Add billing read permission [#3139](https://github.com/DataDog/datadog-api-client-go/pull/3139)
* Update DORA endpoints [#3119](https://github.com/DataDog/datadog-api-client-go/pull/3119)

### Added
* Add `form` field for `multipart/form-data` HTTP API tests [#3135](https://github.com/DataDog/datadog-api-client-go/pull/3135)
* Add new endpoint to upsert/list/delete custom kinds [#3128](https://github.com/DataDog/datadog-api-client-go/pull/3128)
* Add spec for team on-call endpoint [#3126](https://github.com/DataDog/datadog-api-client-go/pull/3126)
* Add support for all subtypes in multistep steps [#3100](https://github.com/DataDog/datadog-api-client-go/pull/3100)
* Exposing set action on Terraform V2 [#3099](https://github.com/DataDog/datadog-api-client-go/pull/3099)
* Added new optional field definition to include more detail in findings for '/api/v2/posture_management/findings'  [#3096](https://github.com/DataDog/datadog-api-client-go/pull/3096)
* Add monitor draft status field [#3095](https://github.com/DataDog/datadog-api-client-go/pull/3095)
* Add rum application to restriction policy [#2984](https://github.com/DataDog/datadog-api-client-go/pull/2984)

## 2.38.0/2025-05-28

### Fixed
* add `include` parameter to On-Call team rules test [#3111](https://github.com/DataDog/datadog-api-client-go/pull/3111)
* fix On-Call spec [#3102](https://github.com/DataDog/datadog-api-client-go/pull/3102)
* Make assertion target be int or string [#3093](https://github.com/DataDog/datadog-api-client-go/pull/3093)
* Fix incorrect pattern for url [#3085](https://github.com/DataDog/datadog-api-client-go/pull/3085)
* Make metadata optional for GCS destination [#3075](https://github.com/DataDog/datadog-api-client-go/pull/3075)
* Remove isReadOnly default when creating dashboards [#3074](https://github.com/DataDog/datadog-api-client-go/pull/3074)

### Added
* Add support for Datadog Events as a data source for rules [#3106](https://github.com/DataDog/datadog-api-client-go/pull/3106)
* Add public APIs to search DORA events [#3103](https://github.com/DataDog/datadog-api-client-go/pull/3103)
* Adding endpoints to manage Resource Evaluation Filters [#3091](https://github.com/DataDog/datadog-api-client-go/pull/3091)
* Add Sev0 as a supported incident severity [#3087](https://github.com/DataDog/datadog-api-client-go/pull/3087)
* Share timerestriction object [#3084](https://github.com/DataDog/datadog-api-client-go/pull/3084)
* add On-Call Paging spec [#3078](https://github.com/DataDog/datadog-api-client-go/pull/3078)
* Add pagination method for NDM ListDevices. [#3072](https://github.com/DataDog/datadog-api-client-go/pull/3072)
* Add On-Call Team Rules [#3070](https://github.com/DataDog/datadog-api-client-go/pull/3070)

## 2.37.1 / 2025-04-14

### Fixed
* Change `type` to enum to discriminate included items in the response of `ListCatalogEntity` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2991
* Deprecate options from logs aggregate API public spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3001
* set global headers earlier when preparing request by @amaskara-dd in https://github.com/DataDog/datadog-api-client-go/pull/3000
* change a category in enum for datadog_appsec_waf_custom_rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2993
### Added
* Add datasource to job definition for security monitoring  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2985
* Include new rum types in Usage_metering Yaml by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2989
* Adding new UT apm_error_events keys in summary endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2982
* Add more triggers for workflow automation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2994
* Add specs for Cloud Network Monitoring API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3007
* Add more Security Monitoring Data Source enum values by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2999
* Add componentOf field to Service, Queue, and Datastore V3 Software Catalog definitions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3003
* Add 'mute_buttons' argument to slack channel definition by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3009
* Add Observability Pipelines API  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3008
* add rum slo bugfix by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3018
* Add trace_rate support to APM retention filter APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3023
* Update NDM GetInterfaces documentation to add ip_addresses attribute by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3022
* Add assertRequests browser step type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3027
* Add user behavior case actions in API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3026
### Changed
* Remove OpenAPI enum enforcement of Service Definition v2dot2 type field from service definition endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2986
* Add on-call schedules endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/3013


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.36.1...v2.37.1

## 2.36.1 / 2025-03-11

### Changed
* Remove meta from RUM retention filters APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2977


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.36.0...v2.36.1

## 2.36.0 / 2025-03-11

### Fixed
* Remove `javascript` browser variable type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2920
* Additional rules to inject openapi type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2940
* Fix `ListCatalogEntity` pagination endpoint to use correct offset value by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2959
### Added
* add new related_assets filter query parameter to the get a list of metrics V2 API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2886
* Add actions and groupSignalsBy field in detection rules API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2922
* Add Workflows CRUD Public API Endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2912
* Add endpoint to retrieve Security Monitoring rule version history by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2893
* Adds override_existing_configurations and include_actively_queried_configurations to bulk tag config endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2930
* Add `number_format` to each formula in widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2933
* Add `trend` support for `cell_display_mode` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2929
* Add support for span id remapper in logs pipelines processors by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2942
* Add evaluation_window and keep_alive for Security monitoring rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2943
* Add `extractedValuesFromScript` to multistep API tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2947
* Update timezone for cumulative window by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2946
* Document Agentless AWS scan options routes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2955
* Create types for app builder queries explicitly, remove experimental flag by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2931
* Document Agentless AWS on demand routes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2960
* Add quality_issues to monitor schema on monitor search API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2962
* Introduce public v2 endpoints for Application Security by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2948
* Add delete log index to public API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2961
* Add v2 endpoints for RUM retention filters. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2957
* Added storage class information to the S3 archive destination by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2968
### Changed
* Revert GetSBOM to `x-unstable` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2924
* Update documentation with account filtering info for aws_cur_config endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2936
* Update sharing APIs to match server by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2945
* Update Vulnerabilities endpoints documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2950
### Deprecated
* Deprecate API management endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2935

## New Contributors
* @ksepehr made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2941

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.35.0...v2.36.0

## 2.35.0 / 2025-02-05

### Fixed
* Modify owner properties to be a string by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2913
### Added
* Add UT breakdown for fargate_container_profiler billing dimension by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2842
* Add synthetics browser step public_id field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2709
* Add support for vulnerability management  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2843
* add start_date to suppression APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2810
* Add CSM Coverage Analysis API specs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2849
* Add allow_self_lockout to documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2858
* Ephemeral Infra_host new keys in summary endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2860
* Update app builder API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2854
* Add meta and source fields to JSONAPIErrorItem by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2853
* Add CSM Agentless Read Endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2850
* Update rum doc to include new usage types by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2863
* add cost monitor type to API Spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2874
* Add Action Connection API for Workflow Automation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2861
* Add `type` in Data Deletion API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2878
* Add `provider_name` attribute to pipelines API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2865
* Add support for vulnerability management - GetSBOM new endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2882
* Remove preview status for GetBillingDimensionMapping endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2892
* Add encryption field to logs archive destination by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2889
* Add tags and description to logs pipelines by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2895
* Publish security notification rules API endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2887
* Publish app builder API documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2867
* update public document with configuration event type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2903
* Add support for Entity kind API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2908
* Rename `embeddedQueries` attribute to `queries` in app builder api by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2909
### Changed
* Fix specification for Azure metric filtering by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2846
* Change allow_self_lockout from string to bool by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2862
* remove flag Beta for cost-by-tag endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2900
* Added Support for Workflow Webhooks Public API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2904
* Vulnerabilities endpoints GA - Remove `x-unstable` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2910


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.34.0...v2.35.0

## 2.34.0 / 2024-12-17

### Added
* Create AWS Integrations v2 API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2516
* Add step_functions as valid enum for v1 AWS tag filter spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2788
* Fix authz scope descriptions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2799
* Updated OpenAPI logs_pattern_query to support Patterns for any attribute by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2797
* Add API specification for events intake v2 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2731
* Data Deletion Endpoints Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2776
* Add `exitIfSucceed` to multistep API tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2825
* Security Monitoring Rule - Add the updatedAt field in the SecurityMonitoringStandardRuleResponse by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2801
* add docs for pagination in /api/v2/metrics endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2701
* Add daily as a valid enum for SLOReportInterval by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2835
* Add new product Code Security host for summary endpoint and UA endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2803
* Add CSM Agents Read Endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2833
* Add app builder API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2828
### Changed
* Remove mobile device ids and make all device ids simple string by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2794
* Remove support for `namespace_filters.include/exclude_all` in v2 AWS Integrations API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2796
* Add running pipelines on custom pipelines API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2791
### Removed
* Remove unnecessary field in list stream column config by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2836
### Deprecated
* Remove `/api/v2/cost/enabled` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2822

## New Contributors
* @bthuilot made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2831

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.33.0...v2.34.0

## 2.33.0 / 2024-11-12

### Added
* Adds accepted reasons for archiving signal by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2681
* Add usage type breakdown for error tracking billing dimension by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2768
* Add Historical Job endpoints to Datadog API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2765
* Add new keys for CWS Fargate Task in summary usage and usage attribution endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2773
* Add missing measures for SLOs data source by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2774
* Downgrade noisy unstable operation log to debug by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2780


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.32.0...v2.33.0

## 2.32.0 / 2024-11-07

### Fixed
* Fix Toplist widget's stacked display style - remove legend as required field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2710
* Remove user fields that are unsupported by the Incidents API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2721
* Fix Synthetics batch status by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2760
### Added
* Add MSTeams integration metadata info by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2723
* Update GCP API Spec to support `is_resource_change_collection_enabled` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2727
* Add vulnerability type to Findings API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2726
* Update Documentation for Data Stream Monitoring by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2733
* Add LLM Observability to ListStreamSource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2732
* Add synthetics stepDetail.allowFailure and stepDetail.failure by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2741
* Integrate incident types into Incidents API documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2725
* Add `use_recommended_keywords` attribute to sensitive data scanner rule spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2752
* Add domain allowlist endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2749
* Add v2 endpoints for RUM custom metrics. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2748
* Documentation for beta /v2/usage/billing_dimension_mapping by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2671
* Add `alwaysExecute` and `exitIfSucceed` to Synthetics steps by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2761
* Add metric_namespace_configs to GCP v2 API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2763
### Changed
* Edit Naming for v2 Microsoft Teams Integration Endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2722
* Change the mobile device ids from enum to string by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2720
* Mark Cost Attribution end_month parameter as not required by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2706
* Allow for any type for additionalProperties in HTTPLogItem by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2742
* Make some amendments to the new mobiles schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2740
* Make value be oneOf number or string by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2753
* Add examples for resources for Cloudflare by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2747
### Removed
* Remove deprecated estimated usage types for usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2745
### Deprecated
* Deprecate two sds metadata fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2739
* Delete `api/v2/cost/aws_related_accounts` from spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2754
* Deprecate `api/v2/cost/enabled` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2756


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.31.0...v2.32.0

## 2.31.0 / 2024-10-02

### Fixed
* change schema used in FastlyServicesResponse by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2700
### Added
* Add new synthetics HTTP javascript assertion by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2616
* Dashboards - Toplist widget style - Add palette by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2668
* Allow Table Widget requests to specify text replace formatting in dashboards by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2669
* Add documentation for Data Jobs Monitoring summary keys by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2672
* Update estimate docs with realtime changes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2704
* Ensure clients can handle empty oneOf objects by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2702
* Add referenceTables field to security monitoring endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2697
* Add UA documentation for new DJM usage_type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2698
* Add v2 endpoints for MS Teams Integration by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2707
* Add documention for OCI Integration by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2713
* Add schema for mobile test by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2682
* Add Synthetics endpoint to fetch uptimes in API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2661
### Changed
* Split the synthetics request port field into a oneOf by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2678
* Remove unused field `color` in `TeamUpdateAttributes` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2674
* Powerpack add support for prefix and available values by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2683
* bump go version to 1.22 by @amaskara-dd in https://github.com/DataDog/datadog-api-client-go/pull/2692
* Update v2 metrics list endpoint filter by metric type to use metric type category by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2705


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.30.0...v2.31.0

## 2.30.0 / 2024-09-04

### Fixed
* Add `is_totp` and `is_fido` to Synthetic global variables by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2582
### Added
* Add `api_key` and `name` to `CloudflareAccountResponseAttributes`. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2631
* Add `api_key` and `name` to `FastlyAccountUpdateRequestAttributes`. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2632
* Add `opsgenie_api_key` to `OpsgenieServiceResponseAttributes`. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2633
* Add `category` and `remote_config_read_enabled` to `APIKeyCreateAttributes`, and add `LeakedKey`. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2634
* Allow 4 group-bys for pattern viz by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2642
* add url attribute to metrics assets v2 api by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2644
* Add editable field to suppression rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2545
* Add `num_flex_logs_retention_days` field to logs_indexes api spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2645
* Software catalog openapi spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2641
### Changed
* allow variables in port by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2624
* Fix VFTs and extracted local variables enum types by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2604
* Changed Widget time schema to add support for new fixed_span and live_span object by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2629
### Deprecated
* mark groupby_simple_monitor as deprecated by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2658


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.29.0...v2.30.0

## 2.29.0 / 2024-08-12

### Fixed
* Add `409 Conflict` to `CreateGlobalVariable` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2586
### Added
* Support `incident_analytics` enum in dashboard widget `FormulaAndFunctionEventsDataSource` data sources by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2594
* update usage summary API docs for partner program by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2589
* update historical_cost and projected_cost for partner program by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2579
* Add custom cost endpoints to public API documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2595
* Update documentation for Cloud SIEM Analyzed Logs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2597
* Update documentation for App Sec SCA by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2584
* Add trigger API documentation for workflow automation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2558
* Add PUT endpoint to scorecards APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2605
* Add json tag to `AdditionalProperties` by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2613
* Documentation for new device tags endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2599
* Update documentation for Flex Logs Starter by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2614
### Changed
* add mfa_enabled field and change created_at type to datetime by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2615


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.28.0...v2.29.0

## 2.28.0 / 2024-07-22

### Fixed
* fix monitor enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2569
* dashboards add support for time-slice SLOs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2570
* Make modified by field nullable for get all API keys by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2576
### Added
* add cross org uuids to timeseries query by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2535
* Add network performance monitor type to API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2546
* Document `force_delete_dependencies` for synthetics test deletion by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2540
* Support metric filtering in integration azure GET, PUT APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2551
* add enableProfiling and enableSecurityTesting options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2562
* Add convert rule JSON to terraform to Datadog API Spec. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2537
* add changes for datadog partner program to estimated cost and billable usage APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2542
* Add type as a required field for the different basic auth types by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2553
* Adding Network Device Monitoring API Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2548
* Security Monitoring - Support anomaly threshold detection method by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2567
* update hourly usage API docs for partner program by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2577
* Add resource_type query param to authn mapping spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2515
* Add rum stream to API definition by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2581
### Deprecated
* Deprecate `ListAWSRelatedAccounts` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2585

## New Contributors
* @amaskara-dd made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2573

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.27.0...v2.28.0

## 2.27.0 / 2024-07-01

### Fixed
* Security Monitoring - Define specific payload for rule validation/testing by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2513
* Remove the maximum limitation for the synthetics renotify_interval monitor option by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2517
* Add bodyHash as a synthetics assertion type. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2518
* Add missing attributes envelope in ListAPIs response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2526
### Added
* Allow the usage of the filters field when creating an agent rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2492
* Add tileDef sort attribute by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2494
* Add Security Monitoring rule test endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2509
* Add originalFileName field to the SyntheticsTestRequestBodyFile definition by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2514
* Add support for API management ListAPIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2521
* Add elementsOperator to json path assertion for synthetic HTTP tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2532
* Add /api/v2/org_configs specs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2531
* Update docs for RU Rollout New and Deprecated Keys planned for Oct 1st by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2544
* Add option for wait step in multistep api tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2543
### Changed
* Monitor priority can have custom ranges and be null by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2541

## New Contributors
* @tim-chaplin-dd made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2450
* @jack-edmonds-dd made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2536

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.26.0...v2.27.0

## 2.26.0 / 2024-05-21

### Fixed
* fix case search documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2469
### Added
* Add support variablesFromScript in Synthetics API test by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2471
* Add JSONSchema assertion support to API and multistep tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2448
* add 1 day logs to usage api docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2477
* Update UserTeamIncluded to include teams by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2482
* Security Monitoring - Make Default Tags available in the response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2491
* Add flex logs storage tier by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2493
### [**Breaking**]Changed
* Rename the Cloud Workload Security tag to CSM Threats by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2481


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.25.0...v2.26.0

## 2.25.0 / 2024-04-11

### Fixed
* Update Cleanup script to use GCP STS endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2423
* Add include data to get team memberships response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2407
### Added
* Add `ci-pipeline-fingerprints` field in v2.2 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2432
* Add validation endpoint for Security Monitoring Rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2453
* Add UA documentation for online_archive and incident_management by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2457
* Mark `unit` as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2459
* Add query_interval_seconds to time-slice SLO condition parameters by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2461
* Support providing files for the file upload feature when creating a Synthetic API test by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2460
* Adding SLO Reporting API Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2426
* Security Monitoring Suppression - Add data_exclusion_query field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2465
* aws api adding extended and deprecating old resource collection field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2463
### Changed
* Add Team relationship to AuthNMappings by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2382
### Deprecated
* Remove deprecated /api/v1/usage/attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2443
* Deprecate legacy hourly usage metering endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2439


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.24.0...v2.25.0

## 2.24.0 / 2024-03-13

### Fixed
* Disable additionalProperties for Downtime Schedule UpdateRequest oneOfs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2390
* Fix ListServiceDefinitions pagination information by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2416
### Added
* Adds support for `ListMetricAssets` endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2404
* Add support for new CRUD agent rules endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2410
* Add documentation for workflow usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2418
* Add Custom Destinations Public API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2422
### Deprecated
* Deprecate the pattern property for SDS Standard Pattern Attributes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2411
* Deprecate Incident Services endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2424

## New Contributors
* @antonio-ramadas-dd made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/2421

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.23.0...v2.24.0

## 2.23.0 / 2024-02-27

### Fixed
* Move under common tag Case Management by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2385
* Include user data with team membership resource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2380
### Added
* Case Management Public API documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2364
* Make grpc steps available for synthetics api multisteps tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2384
* Add cloud run filter to GCP v1 and v2 spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2361
* add ASM serverless to usage metering API docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2393
* Add new products to usage API docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2396
### Changed
* Update spec for DORA Metrics Incident endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2381


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.22.0...v2.23.0

## 2.22.0 / 2024-02-06

### Fixed
* Add test support for file parameters by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2356
* Security Monitoring Suppressions - Make expiration date nullable in update payload by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2365
### Added
* Security Monitoring - Add API support for suppression rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2351
* Document support for BYDAY in SLO corrections by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2357
* Add missing optional field env in DORA API endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2358
* Add compressedProtoFile field to SyntheticsTestRequest by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2355
* Add daily limit reset options to logs indexes api by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2359
* Add support for API management API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2354
* Add pagination helper for team memberships by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2367
* Increase limit on allowed number of graphs in split graph widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2372
### Deprecated
* Mark dashboard 'is_read_only' and 'restricted_roles' properties as deprecated by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2343


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.21.0...v2.22.0

## 2.21.0 / 2024-01-10

### Added
* Add priority field to SDS rule and standard-pattern by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2317
* Document new `resource_collection` and `is_security_command_center_enabled` fields in GCP APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2318
* Add SAML attributes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2311
* Security Monitoring - Support custom third party rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2337
* Add public API support for time-slice SLOs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2328
* Add included_keyword_configuration field to SDS rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2340
* Update Documentation for APM DevSecOps by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2341
### Changed
* Change binary to use io.Reader by @therve in https://github.com/DataDog/datadog-api-client-go/pull/2329
* Mark v1 downtime endpoints as deprecated by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2334
* Adding Cloud Cost Management API Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2326
* Add support for Cloudflare API `zone` and `resource` fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2339


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.20.0...v2.21.0

## 2.20.0 / 2023-12-12

### Fixed
* Fix Powerpack schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2262
### Added
* Add support for projected-cost endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2245
* Document missing incident fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2268
* Add active billing dimensions to usage metering by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2288
* Add Cost Attribution To Usage Metering Public Beta Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2273
* Update spec to include new DORA API endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2287
* Add support to patch Synthetics test with partial data using JSON Patch by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2281
* Document new api/app key schemas by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2267
* Document new field `filters` for `CloudWorkloadSecurityAgentRule` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2289
* Live and historical custom timeseries docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2279
* Add week_to_date and month_to_date to widget livespan by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2282
* Document `included_keywords` in `ListStandardPatterns` response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2292
* Document fields `remote_config_read_enabled` and `category` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2290
* Update Azure Spec to include Resource Collection by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2298
* Allow creation of Application Security detection rules from the v2 API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2300
* Add Okta Integration APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2280
* Remove unstable flag for Events v2 api by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2306


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.19.0...v2.20.0

## 2.19.0 / 2023-11-15

### Fixed
* Remove notify_no_data default by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2242
* Fix SecurityMonitoringSignalAttribute field name by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2249
* Fix typo in service definition field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2266
### Added
* Update documentation for Cloud SIEM by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2234
* Add containers API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2235
* Add serverless apm to usage attribution api by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2240
* Document missing parameters by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2198
* Powerpack Live Span Support by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2247
* Add Amazon EventBridge endpoints to AWS Integration API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2236
* Add Container App filters to Azure API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2217
* Add UUID format support by @HantingZhang2 in https://github.com/DataDog/datadog-api-client-go/pull/2253
* Add new UA products to usage metering docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2251
* Cleanup linter warnings by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2258
* Add scorecards endpoints  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2259
* Document top list widget style by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2243
* Add optional group-bys support to security signals by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2260
### Changed
* Add Beta Banner to Send Pipeline Events Endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2248
* Remove endpoint for mute or unmute a finding and add support for bulk mute findings endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2244
* Place `goccy/go-json` behind built tags and revert default encoder to `encoder/json` package by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2270


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.18.0...v2.19.0

## 2.18.0 / 2023-10-16

### Fixed
* Fix schema for query scalar API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2187
* Remove the application key from CreateCIAppPipelineEvent endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2197
* Document 403 on team endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2206
* Powerpack improve group_widget object by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2218
* Remove escalation message default by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2232
### Added
* Update v1 monitor api docs to exclude downtimes v2 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2178
* Add timing scope for response time assertions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2170
* Add Formula and Function query support to heatmap widgets by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2181
* Add synthetics mobile application testing to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2174
* Add split graph widget to dashboard schema  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2175
* Update public docs for CSM Enterprise and CSPM by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2183
* Add serverless apps to usage and usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2202
* Add Network Device Monitoring Netflow to usage by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2212
* Add Powerpacks endpoints to public api spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2184
* Add account-tags to GCP Service Account Attributes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2196
* Add powerpack widget to dashboard schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2213
* Add custom schedule to monitor scheduling options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2199
* Service Catalog support service definition schema v2.2 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2222
* Powerpack pagination and test fixes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2226
* Add support for container images endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2225
* Add global IP ranges to spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1878
### Changed
* Add APM retention filter api documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2172
* Update request requirements of CI Visibility public pipelines write API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2208
* Add get APM retention filter endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2210
* Remove beta label notice on create pipeline API endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2224


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.17.0...v2.18.0

## 2.17.0 / 2023-09-12

### Fixed
* Fix downtimes monitor relationship id schema type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2168
### Added
* Add trace_stream to dashboard ListStreamSource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2139
* Add pagination extension to SLO corrections by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2151
* Adding aas count to the documentation for summary and hourly usage endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2155
* Add pagination extension to SLOs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2152
* Add pagination extension to monitors by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2153
* Add pagination extension to synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2159
* Add 'style' to sunburst requests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2158
* Add pagination extension to notebook by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2161
* Add support for dashboard listing pagination parameters by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2165
* Add pagination parameters to downtimes listing by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2167
* Add pagination extension to user list by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2166
* Add pagination extension to team listing by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2169
* Remove private beta for Downtimes v2 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2163


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.16.0...v2.17.0

## 2.16.0 / 2023-08-23

### Fixed
* Handle `{}` and bool value for additionalProperties by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2104
* Update team schemas by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2103
* Mark downtime v2 start response as required by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2116
* Document new properties and fix security monitoring schemas by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2102
* Add missing CI App fields `page` and `test_level` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2077
* Fix `unparsedObject` deserialization for lists by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2125
* Fix unparsedObject early exit by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2119
* Cleanup UnmarshalJSON in models by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2127
### Added
* Update stated limit for api/v2/metrics from 14 days to 30 days by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2110
* Add missing sensitive data scanner fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2099
* Add Workflow Executions to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2112
* Add missing `type` field for OnDemandConcurrencyCap response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2101
* Add CI Visibility Intelligent Test Runner to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2111
* Add custom_links to distribution widget schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2120
* Add usage field `region` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2121
* Add `message` field to audit logs response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2123
* Add `tags` field to dashboard list response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2122
* API specs for user team memberships by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2135
* Document `EQUAL` comparator by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2138
* Add persistCookies option synthetics test request by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2137
* Expose sds_scanned_bytes_usage in usage attribution API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2141
* Add support in azure integration endpoint for app service plan filters/cspm/custom metrics by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2136
* Add APM and USM usage attribution type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2133
* Document new attributes for team models by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2130
### Changed
* Update dependencies by @therve in https://github.com/DataDog/datadog-api-client-go/pull/2090
* Bump go to `1.19` by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2128


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.15.0...v2.16.0

## 2.15.0 / 2023-07-20

### Fixed
* Spans API docs update by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2070
* Fix filter indexes parameter in logs search by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2075
* Fix nullable `enum` default value rendering by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2081
* Remove unused nullable models by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2085
* Fix Spans endpoint schemas by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2095
### Added
* Add support for geomap widget using response_type `event_list` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2054
* Add support for the spans API endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2060
* Add a new field additional_query_filters to formula and function slo query by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2071
* Add support for `enable_custom_metrics` in Confluent Account by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2068
* Add missing `id` attribute for Confluent Account Response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2074
* Fix downtimes v2 schema and add missing field `canceled` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2080
* Add cloud_cost data source and query definition to dashboards by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2082
* Add missing cloud workload security fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2089
* Add `integration_id` field for dashboard list item by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2094
* Add events response fields `message` and `status` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2091
* Add missing `GetRUMApplications` response field `id` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2092
* Add missing service definition fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2093
* Add overlay type to Dashboards WidgetDisplayType by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2096
* Update IP ranges with remote configuration section by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2098
* Add missing `relationships` to UsersInvitations response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2100
* Added optional field filters when creating a cloud configuration rule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2105
### Changed
* Use alternative go parser by @therve in https://github.com/DataDog/datadog-api-client-go/pull/2061
* Add downtime v2 API in private beta by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2026
* Mark `access_role` as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2078


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.14.0...v2.15.0

## 2.14.0 / 2023-06-27

### Fixed
* Mark `restricted_roles` as nullable in monitor update request by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2029
* Mark additional usage fields as `nullable` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2037
* Updated findings api error responses by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2016
* Mark usage metering field `lines_indexed` as `nullable` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2043
* Update dashboard widget axis field descriptions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2028
* Fix `CreateGCPSTSAccount` return code and update tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2049
* Add support for nullable primitive lists by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2042
* Fix `CreateGCPSTSAccount` response status code by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2052
* Add missing descriptions for authorization scopes in public docs  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2050
* Update CI Visibility pipelines write API endpoint fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2056
* Remove required unmarshal in models by @therve in https://github.com/DataDog/datadog-api-client-go/pull/2057
### Added
* Add support for mute findings endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2000
* Expose `database-monitoring` monitor type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1885
* Add endpoint to get Synthetics default locations by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2036
* Add usage metering RUM Roku fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2044
* Add usage metering fields for AWS and Azure cloud cost management by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2046
* Add support for CI Visibility create pipeline events endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2045
* Add isUndefined synthetics assertion operator by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2048
* Add missing Synthetics and Metrics Scope descriptions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2051
### Changed
* Team name and handle length updates by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2024
### Deprecated
* mark v1 GCP APIs as deprecated by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2039


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.13.0...v2.14.0

## 2.13.0 / 2023-05-31

### Fixed
* Mark usage fields as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1967
* Properly mark usage fields as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1983
* Remove read only attributes from team create and update by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1987
* Fix optional body setting in requests by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/2007
### Added
* Expose `include_breakdown` param for v2 hourly_usage by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1960
* Add support for deserializing `additionalProperties` in GO client by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1951
* Add new grpc assertions for Synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1966
* add additional_query_filters to slo widget  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1976
* Add `customer_impact_scope` to fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1974
* Add notify_end_states and notify_end_types options to downtime by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1978
* Add snapshot timestamp to GetFinding by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1989
* Support schema version parameter in Get and List Service Definition endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1991
* Add Application Vulnerability Management to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1992
* Add formula and function slo query to dash widgets by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1997
* Add secure field to Synthetics Browser Test variables and update docs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1996
* Add MatchingDowntime to monitor schema and with_downtimes parameter to GetMonitor by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2004
* Add auth scopes for the `service_definition` endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2009
* Update documentation for observability pipeline bytes usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2001
* Add option to obfuscate extracted values from Synthetics multistep tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2002
* Add support for GCP STS endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1995
* Add `sort` field to List Stream Widget's request query by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2020
### Changed
* Update spec to change findings limit and security monitoring menu order by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1962
* Require teams_manage scope for creating and deleting teams by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1977
* Update team name and handle length restrictions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/2021
### Deprecated
* Deprecate note for Incident Teams endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1982


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.12.0...v2.13.0

## 2.12.0 / 2023-04-18

### Fixed
* Fix application_security_host_top99p usage field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1916
* Mark `resource_type` attribute as required for Confluent Account by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1931
* Fix spec errors caught with prism validation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1945
* Fix spans/logs custom metrics delete operation responses by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1953
### Added
* Add support for Incident Todo APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1900
* Add supported relations in restriction policy  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1910
* Add parameter to downtime API for returning creator info by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1914
* Publish the new ingested timeseries metrics for usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1897
* Add tags field to dashboard API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1907
* Add pagination support to SearchIncidents by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1923
* Add service catalog v2.1 schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1920
* Add team API specs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1917
* Add spans metrics API endpoints specification by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1924
* Add universal service monitoring to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1927
* Add a new contact type in service catalog api for schema v2 and v2.1 by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1930
* Add pagination support for the GET service_definitions endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1935
* Publish logs forwarding fields in summary usage API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1934
* Add compressedJsonDescriptor to Synthetics gRPC tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1932
* Add region field and note about multiregion start by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1918
* Add AP1 support by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1941
* Add support for shared dashboards endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1928
### Deprecated
* Deprecate audit logs usage by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1943


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.11.0...v2.12.0

## 2.11.0 / 2023-03-14

### Added
* Add SLO to GRACE API spec by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1905
* Add retry logic when rate limit by @HantingZhang2 in https://github.com/DataDog/datadog-api-client-go/pull/1896
* Add audit trail to usage metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1902

## New Contributors
* @HantingZhang2 made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1896

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.10.0...v2.11.0

## 2.10.0 / 2023-03-07

### Added
* Add restriction policy APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1866
* Support RUM data source in Query API and fix aggregators by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1874
* Add endpoint to get and set on demand concurrency cap for Synthetics by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1870
* Publish IP allowlist APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1859
* Expose Flutter fields to rum product in the meter usage API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1881
* Add profiled fargate tasks to usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1886
* Add cipipeline stream to ListStreamSource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1847
* Add application_security to security monitoring rule type enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1884
* Add `ci_pipelines` enum to `FormulaAndFunctionEventsDataSource` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1894
* Add citest stream to ListStreamSource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1889
* Add `logs_issue_stream` enum to `ListStreamSource` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1895
* Add support for Incident Integration Metadata APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1887
### Changed
* Refactor api request execution by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1879
### Security
* Resolve Dependabot alert by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1880


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.9.0...v2.10.0

## 2.9.0 / 2023-02-15

### Fixed
* Set hosts versions as `type any` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1868
### Added
* Add orchestrator section in IP ranges by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1860
* Add Cloud Cost Management fields to Usage Metering endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1864
* Add cloud-cost as a supported query data source by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1856
* Add Cloud Cost And Container Excl Agent Usage Fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1869
* Add SLO status and error budget remaining to search API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1873
* Add `sort` field to SLOListWidgetQuery by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1871


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.8.0...v2.9.0

## 2.8.0 / 2023-02-08

### Fixed
* Mark timeseries values as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1819
* Fix path constraints check by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1833
* Add namespaces attribute and rename excluded_attributes in SDS Public API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1834
* Set macV as `type any` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1851
### Added
* Add httpVersion option to Synthetics API tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1824
* Add deprecationDate to security monitoring rule response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1825
* Add new group by configuration to list stream widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1821
* Add synthetics advanced scheduling by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1810
* Add notification preset enum field to monitor options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1829
* Add support for Cloudflare integration API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1828
* Add support for Fastly account API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1750
* Add monitor configuration policies by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1832
* Support is_cspm_enabled field in GCP integrations by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1843
* Add run workflow widget to dashboard schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1837
* Add new missing enum values for `aggregation` and `detectionMethod` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1854
* Add region to estimated cost and historical cost response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1855
* Add Usage Metering container_excl_agent_usage fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1857
* Add event_stream fields to dashboard list stream widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1858
### Changed
* Move Service account create from users to service accounts by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1841


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.7.0...v2.8.0

## 2.7.0 / 2023-01-11

### Fixed
* Update CI Visibility types of BucketResponse schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1805
* Fix logs aggregate integer facets by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1812
### Added
* Add support for query scalar and timeseries endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1798
* Add estimated rum sessions usage types to UA enums by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1791
* Update API spec to allow primary timeframe, target, and warning by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1808
* Add Usage Metering Cont Usage fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1809
* Add secure field to synthetics config variables by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1811
* Add Support for Incident Management Search API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1769
* Add TOTP parameters to Synthetics test options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1815
### Changed
* Remove indexed logs from Usage Attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1814
* Remove pagination parameter from CI visibility aggregate endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1818


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.6.1...v2.7.0

## 2.6.1 / 2022-12-20

### Fixed
* Remove incorrect required fields from CloudConfigurationComplianceRuleOptions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1801


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.6.0...v2.6.1

## 2.6.0 / 2022-12-20

### Fixed
* Fix service catalog schema change by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1749
* Add missing response fields to MTD usage attribution endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1776
* Fix missing field in Synthetics tests authentication configuration by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1782
* Mark `hosts` response version fields as nullable by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1778
### Added
* Add fields for CSPM GCP usage by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1756
* Introduce `PaginationResult` type to return errors in WithPagination methods by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1755
* Add offset and limit parameter to SLO correction API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1759
* Add documentation for Logs Pipelines ReferenceTableLogsLookupProcessor  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1758
* Adding new field for the usage metering infra hosts by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1761
* Add `include_percentiles` field in Logs Custom Metrics by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1740
* Add OAuth support for Synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1764
* Add new billable summary fields by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1774
* RUM Applications Management API add client_token by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1765
* Support GRPC unary calls in Synthetics by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1772
* Add style object to dashboard widget formulas by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1787
* Add enable_samples monitor option by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1789
* Update security_monitoring endpoints for cloud_configuration rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1790
* Add support for sensitive data scanner APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1784
* Add synthetics_parallel_testing to Usage Metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1792
* Synthetics add pagination params to get all tests endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1793


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.5.0...v2.6.0

## 2.5.0 / 2022-11-09

### Added
* Add support for CI Visibility API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1725
* Add support for querying logs in Online Archives by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1711
* Add new SDS fields to usage API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1726
* Remove Beta status for SLO history endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1728
* Update formula and function monitor enum datasource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1733
* Add scheduling_options to monitor definition by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1735
* Appsec Fargate Public Documentation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1737
* Adds noScreenshot to SyntheticsStep by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1739
* Add support for xpath assertions in synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1715
* Add bodyType to Synthetics request by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1743

## New Contributors
* @bripkens made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1734

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.4.0...v2.5.0

## 2.4.0 / 2022-10-24

### Fixed
* Fix SearchSLO response structure by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1691
* Add Default Rule ID in SignalRuleResponseQuery by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1705
* Remove incident's resolved attribute from update requests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1707
* Fix event monitor created_at by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1710
* Fix spectral rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1713
### Added
* Add support for incident attachment APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1685
* Add notify_by monitor option by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1699
* Add support for service definitions APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1714
* Add support for confluent cloud integration by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1720
### Deprecated
* Add a note for deprecated APIs and models by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1704
* Deprecate metric field of Security Monitoring Rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1723


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.3.1...v2.4.0

## 2.3.1 / 2022-09-28

### Fixed
* Refactor RuleQuery models by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1689


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.3.0...v2.3.1

## 2.3.0 / 2022-09-27

### Added
* Add ListActiveConfigurations endpoint and add new filter[queried] param to list tag configurations endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1619
* Add doesNotExist to synthetics operator enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1659
* Add TopologyMapWidget to dashboard schema by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1662
* Add Overall Status support to SLO Search API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1664
* Add APM Fargate to Usage Metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1658
* [dashboards] Add support for template variable multiselect by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1668
* Add storage option to widget query definitions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1673
* Add support for retrieving a security signal by ID by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1681
* Add support for signal correlation API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1674
* Add support for SLO List widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1680
* Add new historical_cost endpoint, and update estimate_cost by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1672

## New Contributors
* @nkzou made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1657
* @dependabot made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1676

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.2.0...v2.3.0

## 2.2.0 / 2022-08-31

### Added
* add priority parameters for dashboard monitor summary widget by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1631
* Add `logs_pattern_stream` to `list_stream` widget source by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1633
* Add group_retention_duration and on_missing_data monitor options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1626
* Expose CSPM aws host count in Usage Metering API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1645
* Add estimated ingested logs attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1613
* Add org region to usage summary and billable usage summary by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1639
* add compression methods to metric payloads by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1610
* Add role relationships to RoleUpdateData by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1642
* Add `ci_tests` enum to FormulaAndFunctionEventsDataSource by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1641
* Add missing options and request option to synthetics test by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1632
* Add support for global variable from multistep synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1598
### Changed
* update deprecated usage attribution API docs to direct users to migra… by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1634
* [Synthetics] remove started form eventType enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1636

## New Contributors
* @ganeshkumarsv made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1592

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.1.0...v2.2.0

## 2.1.0 / 2022-08-10

### Fixed
* Update Pagerduty operation `DeletePagerDutyIntegrationService` response status code by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1612
### Added
* Add support for digest auth in synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1622
* Add support for RUM application endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1623
### Changed
* Refactor package names by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1624


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v2.0.0...v2.1.0

## 2.0.1 / 2022-08-10

**_NOTE:_** Version used to retract v2.0.0 and v2.0.1. DO NOT USE

## 2.0.0 / 2022-08-01

**_NOTE:_** Premature major version v2 release. DO NOT USE

### Added
* Add support for Events V2 endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1551
* [RQ-2492]: Add custom_events to list of product families in hourly-usage api. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1608
* Re-introduce Estimated Cost API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1614
### Changed
* Create a `common` shared package by @skarimo in https://github.com/DataDog/datadog-api-client-go/pull/1588


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.16.0...v2.0.0

## 1.16.0 / 2022-07-20

### Fixed
* Add synthetics results api replay only tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1595
### Added
* Add estimated ingested spans to usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1581
* Add v2 Security monitoring signals triage operations. by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1571
* docs(dataviz): update Treemap widget definition with deprecated properties + updated description [VIZZ-2305] by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1596
* Add hourly usage v2 endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1577
* Add metrics field in the RuleQuery by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1578
### Changed
* Add description of metric type enums by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1583
* remove x-unstable property for usage attribution endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1594


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.15.1...v1.16.0

## 1.15.1 / 2022-07-11

### Fixed
* Allow compilation without cgo by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1585


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.15.0...v1.15.1

## 1.15.0 / 2022-07-04

### Fixed
* Fix serialization of arrays by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1508
* AuthN Mapping spec cleanup to match implementation by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1507
* Fix additionalProperties on SyntheticsAPITestResultData by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1533
* Fix synthetics vitals type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1534
* Remove include_percentiles default by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1559
* Mark message as required for Synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1553
* Don't store decode errors by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1572
### Added
* Add `ci-tests` monitor type by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1492
* Add RUM settings schema to synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1500
* Add v1 signal triage endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1496
* Add connection to synthetics assertion type enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1512
* Add grpc subtype to synthetics tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1486
* Add support for `zstd1` Content-Encoding by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1448
* Add include descendants to monthly and hourly usage attribution APIs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1530
* Add v2 endpoints for Opsgenie Integration by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1520
* Add distribution points intake endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1518
* Add height and width params to graph snapshot by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1542
* Add support for defining histogram requests in Distribution widgets by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1527
* Add DowngradeOrg endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1544
* Add new options for new value detection type on security monitoring rules by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1537
* Add ci execution rule in Synthetics options by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1538
* Add SLO Search API endpoint  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1451
* New usage metering endpoint for estimated cost by org by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1562
* Add estimated indexed spans usage attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1561
* Handle raw  json for additionalProperties in typescript  by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1535
* Add Application Security Monitoring Hosts Attribution by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1563
* Add support for security monitoring rule dynamic criticality by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1548
* Update IP ranges with synthetics private locations section by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1567
* Add new products to billable summary by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1570
* Update usage attribution enums by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1564
### Changed
* Remove unstable marker from SLO corrections API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1519
* Remove unstable/beta note since Metrics Without Limits is GA by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1526
* Remove unstable marker on security list signal endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1540
* Update metric intake v2 accept response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1568

## New Contributors
* @jybp made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1557

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.14.0...v1.15.0

## 1.14.0 / 2022-05-18

### Fixed
* Normalize format of date-time fields by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1472
* Remove unused pararameter from authn mapping by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1487
### Added
* Add Usage API endpoint for observability-pipelines and add properties to v1 GetUsageSummary by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1422
* Add Historical Chargeback Summary endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1469
* Add `...WithPagination` helper methods by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1468
* Expose v2 usage endpoint for application security monitoring by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1479
* Add `rehydration_max_scan_size_in_gb` field to Logs Archives by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1475
* Add `mute_first_recovery_notification` option to downtime by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1444
* Add lambda traced invocations usage endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1484
* Expose new usage field for react sessions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1488
* Add missing option and enum value for SecurityMonitoringRule by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1493
* Adds docs for metric estimate endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1476
* Allow additional log attributes by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1467
* Add v2 endpoint for submitting series by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1409
### Deprecated
* Deprecate old usage apis by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1490


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.13.0...v1.14.0

## 1.13.0 / 2022-04-20

### Fixed
* Set correct type for `tags` property by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1452
### Added
* Add `restricted_roles` to Synthetics tests and private locations by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1443
* Add v2 SAML config IdP Metadata upload endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1450
* Support pagination in Python by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1457
### Changed
* Remove references to optional arrays by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1454
### Removed
* [dashboards] Removed `issue_stream` type from `ListStreamSource` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1446


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.12.0...v1.13.0

## 1.12.0 / 2022-04-06

### Fixed
* Fix required nullable fields by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1428
* Propagate unparsed objects by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1425
* Make type optional for synthetics basic auth model by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1430
### Added
* Add aggregate endpoint for RUM by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1417
* Add  `median` aggregation functions to RUM and logs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1427
* Add endpoint for validation of existing monitors by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1429
* Create new ListStreamSource types in order to deprecate ISSUE_STREAM by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1439
* [Query Value Widget] Add the timeseries background by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1426
### Changed
* Cleanup unused modules and functions by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1420
### Removed
* Remove `lambda_usage` and `lambda_percentage` from usage API by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1411

## New Contributors
* @Stoovles made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1423

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.11.0...v1.12.0

## 1.11.0 / 2022-03-28

### Fixed
* Fix org name maximum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1393
* Use `$ref` names for `oneOf` attribute names by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1397
* Fix pagination for top avg metrics endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1400
### Added
* [RUM] Add search endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1360
* Add support for getting online archive usage by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1372
* Add endpoint for retrieving audit logs by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1371
* Add support for Error Tracking monitors by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1395
* Add support for `ci-pipelines` monitor using Formulas and Functions by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1405
### Changed
* Use Python port of code generator by @jirikuncar in https://github.com/DataDog/datadog-api-client-go/pull/1376
### Deprecated
* [monitors] Deprecate `locked` property and clarify documentation for `restricted_roles` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1387

## New Contributors
* @juan-fernandez made their first contribution in https://github.com/DataDog/datadog-api-client-go/pull/1379

**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.10.0...v1.11.0

## 1.10.0 / 2022-03-03

### Fixed
* Fix event intake response by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1350
* Fix type for `date` field in `LogsByRetentionMonthlyUsage` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1362
### Added
* [Synthetics] Add missing option for SSL tests by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1354
* Add impossible travel detection method by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1357
* Add CI App usage endpoint and usage summary columns by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1361


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.9.0...v1.10.0

## 1.9.0 / 2022-02-18

### Fixed
* Add missing type to `CloudWorkloadSecurityAgentRuleAttributes` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1339
* Add missing type to enum by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1342
* Add nullable user relationships to incidents and use this relationship schema for `commander_user` by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1311
### Added
* Add organization metadata to additional Usage API responses by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1317
* Add support for formula and function in monitors by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1321
* Add endpoint for managing SAML AuthN mappings by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1319
* [Synthetics] Add `isCritical` to browser test steps by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1327
* Add metrics bulk-config endpoint by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1324
* Add support for "estimated usage attribution" by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1335
* Add org metadata for all hourly usage endpoints by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1341
### Changed
* Remove default nullable models by @therve in https://github.com/DataDog/datadog-api-client-go/pull/1312
* Add CSPM usage fields and change properties to nullable doubles by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1307
* Add synthetics test result failure field by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1318
* Fix funnel steps definition by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1330
* Extract incident meta object by @api-clients-generation-pipeline in https://github.com/DataDog/datadog-api-client-go/pull/1333


**Full Changelog**: https://github.com/DataDog/datadog-api-client-go/compare/v1.8.0...v1.9.0

## 1.8.0 / 2022-01-18

* [Added] Add `filter[deleted]` parameter for searching recently deleted dashboards. See [#1296](https://github.com/DataDog/datadog-api-client-go/pull/1296).
* [Added] Add support for authentication and proxy options in Synthetics. See [#1267](https://github.com/DataDog/datadog-api-client-go/pull/1267).
* [Added] Support formulas and functions in Treemap Widget. See [#1291](https://github.com/DataDog/datadog-api-client-go/pull/1291).
* [Added] Add Cloud Workload Security Agent Rules API. See [#1282](https://github.com/DataDog/datadog-api-client-go/pull/1282).
* [Added] Add `offset` and `limit` parameters to usage listing endpoint. See [#1285](https://github.com/DataDog/datadog-api-client-go/pull/1285).
* [Added] Add monthly usage attribution API spec. See [#1274](https://github.com/DataDog/datadog-api-client-go/pull/1274).
* [Added] Add missing hosts metadata fields. See [#1269](https://github.com/DataDog/datadog-api-client-go/pull/1269).
* [Added] Add `replay_session_count ` and update documentation for `rum_session_count`. See [#1284](https://github.com/DataDog/datadog-api-client-go/pull/1284).
* [Added] Add retry options for a step in Synthetics multistep test. See [#1277](https://github.com/DataDog/datadog-api-client-go/pull/1277).
* [Added] Document `author_name` in dashboard response. See [#1275](https://github.com/DataDog/datadog-api-client-go/pull/1275).
* [Added] Add organization metadata for RUM sessions usage and expose `rum_browser_and_mobile_session_count`. See [#1270](https://github.com/DataDog/datadog-api-client-go/pull/1270).
* [Added] Add endpoint to retrieve hourly usage attribution. See [#1249](https://github.com/DataDog/datadog-api-client-go/pull/1249).
* [Added] Add support for scoped application keys. See [#1234](https://github.com/DataDog/datadog-api-client-go/pull/1234).
* [Added] Add endpoint for cloning roles. See [#1258](https://github.com/DataDog/datadog-api-client-go/pull/1258).
* [Added] Add organization metadata for audit logs, CWS, CSPM, DBM. See [#1264](https://github.com/DataDog/datadog-api-client-go/pull/1264).
* [Added] Add `ci-pipelines alert` to monitors enum. See [#1255](https://github.com/DataDog/datadog-api-client-go/pull/1255).
* [Added] Add support for sunburst widget in dashboard. See [#1262](https://github.com/DataDog/datadog-api-client-go/pull/1262).
* [Fixed] Clarify required fields for `SyntheticsAPIStep`, `SyntheticsAPITest`, and `SyntheticsBrowserTest`. See [#1202](https://github.com/DataDog/datadog-api-client-go/pull/1202).
* [Fixed] Fixes to Cloud Workload Security API. See [#1294](https://github.com/DataDog/datadog-api-client-go/pull/1294).
* [Fixed] Make downtime weekdays nullable. See [#1279](https://github.com/DataDog/datadog-api-client-go/pull/1279).
* [Fixed] Fix a typo in an incident field attribute description. See [#1240](https://github.com/DataDog/datadog-api-client-go/pull/1240).
* [Fixed] Fix `SecurityMonitoringSignal.attributes.tags` type. See [#1243](https://github.com/DataDog/datadog-api-client-go/pull/1243).
* [Changed] Remove read only fields in `EventCreateRequest`. See [#1292](https://github.com/DataDog/datadog-api-client-go/pull/1292).
* [Changed] Change pagination arguments for querying usage attribution. See [#1273](https://github.com/DataDog/datadog-api-client-go/pull/1273).
* [Deprecated] Remove session counts from RUM units response. See [#1252](https://github.com/DataDog/datadog-api-client-go/pull/1252).
* [Removed] Remove deprecated AgentRule field in Security Rules API for CWS. See [#1268](https://github.com/DataDog/datadog-api-client-go/pull/1268).

## 1.7.0 / 2021-12-09

* [Added] Add Limit Note for Hourly Requests. See [#1230](https://github.com/DataDog/datadog-api-client-go/pull/1230).
* [Added] Expose estimated logs usage in Usage Attribution API. See [#1231](https://github.com/DataDog/datadog-api-client-go/pull/1231).
* [Added] Add endpoint to get corrections applied to an SLO. See [#1221](https://github.com/DataDog/datadog-api-client-go/pull/1221).
* [Added] Expose `public_id` and `org_name` in Usage API response. See [#1224](https://github.com/DataDog/datadog-api-client-go/pull/1224).
* [Added] Document query in `MonitorSearchResult`. See [#1222](https://github.com/DataDog/datadog-api-client-go/pull/1222).
* [Added] Add 429 error responses. See [#1208](https://github.com/DataDog/datadog-api-client-go/pull/1208).
* [Added] Add support for profiled Fargate tasks in Usage API. See [#1205](https://github.com/DataDog/datadog-api-client-go/pull/1205).
* [Added] Add support for `websocket` synthetics tests. See [#1206](https://github.com/DataDog/datadog-api-client-go/pull/1206).
* [Added] [Synthetics] Add support for UDP API tests. See [#1197](https://github.com/DataDog/datadog-api-client-go/pull/1197).
* [Added] Add trigger synthetics tests endpoint. See [#1173](https://github.com/DataDog/datadog-api-client-go/pull/1173).
* [Added] Add RUM Units to usage metering API. See [#1188](https://github.com/DataDog/datadog-api-client-go/pull/1188).
* [Added] [dashboards formulas and functions] Add formulas and functions support to change widget. See [#1204](https://github.com/DataDog/datadog-api-client-go/pull/1204).
* [Fixed] Be more resilient to plain text errors. See [#1227](https://github.com/DataDog/datadog-api-client-go/pull/1227).
* [Fixed] Fix monitor `timeout_h` example and limits. See [#1219](https://github.com/DataDog/datadog-api-client-go/pull/1219).
* [Fixed] Remove event title length constraint. See [#1215](https://github.com/DataDog/datadog-api-client-go/pull/1215).
* [Fixed] Mark `batch_id` in Synthetics Trigger CI response as nullable. See [#1210](https://github.com/DataDog/datadog-api-client-go/pull/1210).
* [Fixed] SLO Correction attributes `rrule` and `duration` can be nullable. See [#1200](https://github.com/DataDog/datadog-api-client-go/pull/1200).
* [Fixed] Change `UsageNetworkFlowsHour.indexed_event_count` to match actual API. See [#1196](https://github.com/DataDog/datadog-api-client-go/pull/1196).
* [Fixed] Fix type for `ratio_in_month` in usage metering. See [#1183](https://github.com/DataDog/datadog-api-client-go/pull/1183).
* [Changed] [Synthetics] Fix required target in assertions and type in step results. See [#1201](https://github.com/DataDog/datadog-api-client-go/pull/1201).

## 1.6.0 / 2021-11-09

* [Added] Add support for Azure `automute` option. See [#1179](https://github.com/DataDog/datadog-api-client-go/pull/1179).
* [Added] Add v2 intake endpoint. See [#1172](https://github.com/DataDog/datadog-api-client-go/pull/1172).
* [Added] Add support for RRULE fields in SLO corrections. See [#1126](https://github.com/DataDog/datadog-api-client-go/pull/1126).
* [Added] Add aggregations attribute to v2 metric tag configuration. See [#1101](https://github.com/DataDog/datadog-api-client-go/pull/1101).
* [Added] Add `apm_stats_query` property to `DistributionWidgetRequest`. See [#1161](https://github.com/DataDog/datadog-api-client-go/pull/1161).
* [Fixed] Use plural form for dbm hosts usage properties. See [#1141](https://github.com/DataDog/datadog-api-client-go/pull/1141).
* [Changed] Update Synthetics CI test metadata. See [#1140](https://github.com/DataDog/datadog-api-client-go/pull/1140).
* [Deprecated] Update property descriptions for Dashboard RBAC release. See [#1171](https://github.com/DataDog/datadog-api-client-go/pull/1171).

## 1.5.0 / 2021-10-18

* [Added] Add support for funnel widget in dashboards. See [#1115](https://github.com/DataDog/datadog-api-client-go/pull/1115).
* [Added] Add information about creator to Synthetics tests details. See [#1122](https://github.com/DataDog/datadog-api-client-go/pull/1122).
* [Added] Add support for gzip and deflate encoding. See [#1119](https://github.com/DataDog/datadog-api-client-go/pull/1119).
* [Added] Add support for formulas and functions in the Scatterplot Widget for dashboards. See [#1113](https://github.com/DataDog/datadog-api-client-go/pull/1113).
* [Added] Document encoding in metrics intake. See [#1131](https://github.com/DataDog/datadog-api-client-go/pull/1131).
* [Added] Add `servername` property to SSL Synthetics tests request. See [#1130](https://github.com/DataDog/datadog-api-client-go/pull/1130).
* [Added] Add `renotify_occurrences` and `renotify_statuses` monitor options. See [#1143](https://github.com/DataDog/datadog-api-client-go/pull/1143).
* [Added] Add `type` and `is_template` properties to notebooks. See [#1146](https://github.com/DataDog/datadog-api-client-go/pull/1146).
* [Added] [Synthetics] Add endpoint to get details of a batch. See [#1090](https://github.com/DataDog/datadog-api-client-go/pull/1090).
* [Added] Add SDS to usage metering endpoint. See [#1153](https://github.com/DataDog/datadog-api-client-go/pull/1153).
* [Added] Add `metrics_collection_enabled`, `cspm_resource_collection_enabled ` and `resource_collection_enabled` to AWS integration request. See [#1150](https://github.com/DataDog/datadog-api-client-go/pull/1150).
* [Fixed] Fix typo in usage attribution field names for profiled containers. See [#1123](https://github.com/DataDog/datadog-api-client-go/pull/1123).
* [Fixed] Make sure that OpenAPI definition are valid with real server responses. See [#1121](https://github.com/DataDog/datadog-api-client-go/pull/1121).
* [Fixed] Fix incidents schemas. See [#1128](https://github.com/DataDog/datadog-api-client-go/pull/1128).
* [Fixed] `IncidentFieldAttributesMultipleValue` can be nullable. See [#1129](https://github.com/DataDog/datadog-api-client-go/pull/1129).
* [Fixed] Allow nullable date in notebook cells. See [#1134](https://github.com/DataDog/datadog-api-client-go/pull/1134).
* [Fixed] Fix go handling of nullable enums. See [#1152](https://github.com/DataDog/datadog-api-client-go/pull/1152).
* [Fixed] Remove event title length constraint. See [#1124](https://github.com/DataDog/datadog-api-client-go/pull/1124).
* [Fixed] Make monitor properties `priority` and `restricted_roles` nullable. See [#1158](https://github.com/DataDog/datadog-api-client-go/pull/1158).
* [Changed] Use AVG aggregation function for DBM queries. See [#1118](https://github.com/DataDog/datadog-api-client-go/pull/1118).
* [Changed] Enable compression in responses. See [#1142](https://github.com/DataDog/datadog-api-client-go/pull/1142).

## 1.4.0 / 2021-09-15

* [Added] Added `available_values` property to template variables schema. See [#1089](https://github.com/DataDog/datadog-api-client-go/pull/1089).
* [Added] Add `follow_redirects` options to test request in Synthetics. See [#1096](https://github.com/DataDog/datadog-api-client-go/pull/1096).
* [Added] ApmDependencyStatsQuery for formulas and functions dashboard widgets. See [#1103](https://github.com/DataDog/datadog-api-client-go/pull/1103).
* [Added] Add formula and function APM resource stats query definition for dashboards. See [#1104](https://github.com/DataDog/datadog-api-client-go/pull/1104).
* [Fixed] Fix SLO history error response type for overall errors. See [#1095](https://github.com/DataDog/datadog-api-client-go/pull/1095).
* [Fixed] Mark SLO Correction Type as required. See [#1093](https://github.com/DataDog/datadog-api-client-go/pull/1093).
* [Fixed] Make the `name` property required for APM Dependency Stat Query. See [#1110](https://github.com/DataDog/datadog-api-client-go/pull/1110).
* [Changed] Fix SLO history schema for groups and monitors fields. See [#1099](https://github.com/DataDog/datadog-api-client-go/pull/1099).
* [Changed] Remove metadata from required list for metric SLO history endpoint. See [#1102](https://github.com/DataDog/datadog-api-client-go/pull/1102).

## 1.3.0 / 2021-08-26

* [Added] Add config variables to Synthetics browser test config. See [#1086](https://github.com/DataDog/datadog-api-client-go/pull/1086).
* [Added] Add DBM usage endpoint. See [#1068](https://github.com/DataDog/datadog-api-client-go/pull/1068).
* [Added] Add `audit alert` monitor type. See [#1081](https://github.com/DataDog/datadog-api-client-go/pull/1081).
* [Added] Add `batch_id` to the synthetics trigger endpoint response. See [#1079](https://github.com/DataDog/datadog-api-client-go/pull/1079).
* [Added] Adding support for security monitoring rule `type` property. See [#1065](https://github.com/DataDog/datadog-api-client-go/pull/1065).
* [Added] Add events data source to Dashboard widgets. See [#1067](https://github.com/DataDog/datadog-api-client-go/pull/1067).
* [Added] Add restricted roles for Synthetics global variables. See [#1072](https://github.com/DataDog/datadog-api-client-go/pull/1072).
* [Added] Webhooks integration SDK. See [#1071](https://github.com/DataDog/datadog-api-client-go/pull/1071).
* [Added] Add missing synthetics variable parser type `x_path`. See [#1070](https://github.com/DataDog/datadog-api-client-go/pull/1070).
* [Added] Add `audit_stream` to `ListStreamSource`. See [#1056](https://github.com/DataDog/datadog-api-client-go/pull/1056).
* [Added] Add percentile to dashboard `WidgetAggregator` schema. See [#1051](https://github.com/DataDog/datadog-api-client-go/pull/1051).
* [Added] Add `id_str` property to Event response. See [#1059](https://github.com/DataDog/datadog-api-client-go/pull/1059).
* [Added] Add edge to Synthetics devices. See [#1063](https://github.com/DataDog/datadog-api-client-go/pull/1063).
* [Added] Add endpoints to manage Service Accounts v2. See [#1043](https://github.com/DataDog/datadog-api-client-go/pull/1043).
* [Added] Add `new_group_delay` and deprecate `new_host_delay` monitor properties. See [#1055](https://github.com/DataDog/datadog-api-client-go/pull/1055).
* [Added] Add `include_descendants` param to usage attribution API. See [#1062](https://github.com/DataDog/datadog-api-client-go/pull/1062).
* [Added] Improve resiliency of go SDK when deserializing enums/oneOfs. See [#1028](https://github.com/DataDog/datadog-api-client-go/pull/1028).
* [Added] Add `ContainsUnparsedObject` utility method to check if an object wasn't fully deserialized. See [#1073](https://github.com/DataDog/datadog-api-client-go/pull/1073) and [#1077](https://github.com/DataDog/datadog-api-client-go/pull/1077).
* [Added] Add support for list widget in dashboards. See [#1023](https://github.com/DataDog/datadog-api-client-go/pull/1023).
* [Added] Extend table widget requests to support formulas and functions in dashboards. See [#1046](https://github.com/DataDog/datadog-api-client-go/pull/1046).
* [Added] Add CSPM to usage attribution. See [#1037](https://github.com/DataDog/datadog-api-client-go/pull/1037).
* [Added] Add support for dashboard bulk delete and restore endpoints. See [#1020](https://github.com/DataDog/datadog-api-client-go/pull/1020).
* [Added] Add support for audit logs data source in dashboards. See [#1041](https://github.com/DataDog/datadog-api-client-go/pull/1041).
* [Added] Add `allow_insecure` option for multistep steps in Synthetics. See [#1031](https://github.com/DataDog/datadog-api-client-go/pull/1031).
* [Fixed] Make SLO history metadata unit nullable. See [#1078](https://github.com/DataDog/datadog-api-client-go/pull/1078).
* [Fixed] Minor fixes of the incident schema. See [#1074](https://github.com/DataDog/datadog-api-client-go/pull/1074).
* [Fixed] Fix serialization of query metrics response containing nullable points. See [#1034](https://github.com/DataDog/datadog-api-client-go/pull/1034).
* [Fixed] Fix `status` property name for browser error status in Synthetics. See [#1036](https://github.com/DataDog/datadog-api-client-go/pull/1036).
* [Changed] Add separate schema for deleting AWS account. See [#1030](https://github.com/DataDog/datadog-api-client-go/pull/1030).
* [Removed] Remove deprecated endpoints `/api/v1/usage/traces` and `/api/v1/usage/tracing-without-limits`. See [#1038](https://github.com/DataDog/datadog-api-client-go/pull/1038).

## 1.2.0 / 2021-07-09

* [Added] Add support for `GET /api/v2/application_keys/{app_key_id}`. See [#1021](https://github.com/DataDog/datadog-api-client-go/pull/1021).
* [Added] Add `meta` property with pagination info to SLOCorrectionList endpoint response. See [#1018](https://github.com/DataDog/datadog-api-client-go/pull/1018).
* [Added] Add support for treemap widget. See [#1013](https://github.com/DataDog/datadog-api-client-go/pull/1013).
* [Added] Add missing properties `query_index` and `tag_set` to `MetricsQueryMetadata`. See [#979](https://github.com/DataDog/datadog-api-client-go/pull/979).
* [Fixed] Remove US only constraint for AWS tag filtering. See [#1007](https://github.com/DataDog/datadog-api-client-go/pull/1007).
* [Fixed] Add BDD tests to synthetics. See [#1006](https://github.com/DataDog/datadog-api-client-go/pull/1006).
* [Fixed] Fix response of security filter delete. See [#1002](https://github.com/DataDog/datadog-api-client-go/pull/1002).
* [Fixed] Handle null in query metrics unit. See [#1001](https://github.com/DataDog/datadog-api-client-go/pull/1001).
* [Changed] Remove Synthetics tick interval enum. See [#1005](https://github.com/DataDog/datadog-api-client-go/pull/1005).

## 1.1.0 / 2021-06-16

* [Added] Add missing fields `hasExtendedTitle`, `type`, `version` and `updateAuthorId` for Security Monitoring Rule endpoints. See [#998](https://github.com/DataDog/datadog-api-client-go/pull/998).
* [Added] Dashboard RBAC role support. See [#993](https://github.com/DataDog/datadog-api-client-go/pull/993).
* [Fixed] Fix go JSON struct. See [#992](https://github.com/DataDog/datadog-api-client-go/pull/992).

## 1.0.0 / 2021-06-10

* [Added] Add missing fields in usage billable summary keys. See [#987](https://github.com/DataDog/datadog-api-client-go/pull/987).
* [Added] Add monitor name and priority options. See [#984](https://github.com/DataDog/datadog-api-client-go/pull/984).
* [Added] Add endpoint to list Synthetics global variables. See [#965](https://github.com/DataDog/datadog-api-client-go/pull/965).
* [Added] Add monitors search endpoints. See [#959](https://github.com/DataDog/datadog-api-client-go/pull/959).
* [Added] Add CWS to usage metering endpoint. See [#964](https://github.com/DataDog/datadog-api-client-go/pull/964).
* [Added] Add `tag_config_source` to usage attribution response. See [#952](https://github.com/DataDog/datadog-api-client-go/pull/952).
* [Added] Add audit logs to usage endpoints. See [#978](https://github.com/DataDog/datadog-api-client-go/pull/978).
* [Fixed] Make `assertions` field optional for multistep synthetics tests, and add `global` config variable type. See [#961](https://github.com/DataDog/datadog-api-client-go/pull/961).
* [Fixed] Fix type of day/month response attribute in custom metrics usage. See [#981](https://github.com/DataDog/datadog-api-client-go/pull/981).
* [Fixed] Properly mark monitor required fields. See [#950](https://github.com/DataDog/datadog-api-client-go/pull/950).
* [Changed] Rename `compliance` to `CSPM` in usage endpoint. See [#978](https://github.com/DataDog/datadog-api-client-go/pull/978).
* [Changed] Rename `incident_integration_metadata` to `incident_integrations` to match API. See [#944](https://github.com/DataDog/datadog-api-client-go/pull/944).

## 1.0.0-beta.22 / 2021-05-17

* [Added] Add endpoints to configure Security Filters. See [#938](https://github.com/DataDog/datadog-api-client-go/pull/938).
* [Added] Add `active_child` nested downtime object to `Downtime` component for downtime APIs. See [#930](https://github.com/DataDog/datadog-api-client-go/pull/930).
* [Changed] Change Dashboard WidgetCustomLink properties. See [#937](https://github.com/DataDog/datadog-api-client-go/pull/937).
* [Changed] Make various fixes to synthetics models. See [#935](https://github.com/DataDog/datadog-api-client-go/pull/935).
* [Changed] Update usage attribute endpoint metadata fields. See [#932](https://github.com/DataDog/datadog-api-client-go/pull/932).

## 1.0.0-beta.21 / 2021-05-12

* [Added] Notebooks Public API Documentation. See [#926](https://github.com/DataDog/datadog-api-client-go/pull/926).
* [Added] Add `logs_by_retention` usage property and `GetUsageLogsByRetention` endpoint. See [#915](https://github.com/DataDog/datadog-api-client-go/pull/915).
* [Added] Add anomaly detection method to `SecurityMonitoringRuleDetectionMethod` enum. See [#914](https://github.com/DataDog/datadog-api-client-go/pull/914).
* [Added] Add `with_configured_alert_ids` parameter to get a SLO details endpoint. See [#910](https://github.com/DataDog/datadog-api-client-go/pull/910).
* [Added] Add `setCookie`, `dnsServerPort`,  `allowFailure ` and `isCritical` fields for Synthetics tests. See [#903](https://github.com/DataDog/datadog-api-client-go/pull/903).
* [Added] Add `metadata` property with pagination info to `SLOList` endpoint response. See [#899](https://github.com/DataDog/datadog-api-client-go/pull/899).
* [Added] Add new properties to group widget, note widget and image widget. See [#895](https://github.com/DataDog/datadog-api-client-go/pull/895).
* [Added] Add support for a `rate` metric type in manage metric tags v2 endpoint. See [#892](https://github.com/DataDog/datadog-api-client-go/pull/892).
* [Fixed] Handle typed nils for go client. See [#927](https://github.com/DataDog/datadog-api-client-go/pull/927).
* [Fixed] Remove default value of `is_column_break` layout property of dashboard. See [#925](https://github.com/DataDog/datadog-api-client-go/pull/925).
* [Changed] Enumerate accepted values for fields parameter in usage attr requests. See [#919](https://github.com/DataDog/datadog-api-client-go/pull/919).
* [Changed] Add frequency and remove request as required field from synthetics test. See [#916](https://github.com/DataDog/datadog-api-client-go/pull/916).

## 1.0.0-beta.20 / 2021-04-27

* [Added] Add support for ICMP Synthetics tests. See [#887](https://github.com/DataDog/datadog-api-client-go/pull/887).
* [Added] Add vSphere usage information. See [#880](https://github.com/DataDog/datadog-api-client-go/pull/880).
* [Added] Update properties for dashboard distribution widget. See [#877](https://github.com/DataDog/datadog-api-client-go/pull/877).
* [Added] Mark metric volumes and ingested tags endpoints as stable. See [#872](https://github.com/DataDog/datadog-api-client-go/pull/872).
* [Added] Add `filter[shared]` query parameter for searching dashboards. See [#860](https://github.com/DataDog/datadog-api-client-go/pull/860).
* [Added] Add profiling product fields in usage metering endpoint. See [#859](https://github.com/DataDog/datadog-api-client-go/pull/859).
* [Added] Add `title` and `background_color` properties to dashboard group widget. See [#858](https://github.com/DataDog/datadog-api-client-go/pull/858).
* [Changed] Use new model for Go client API. See [#885](https://github.com/DataDog/datadog-api-client-go/pull/885).
* [Removed] Remove deprecated Synthetics methods `CreateTest` and `UpdateTest`. See [#881](https://github.com/DataDog/datadog-api-client-go/pull/881).

## 1.0.0-beta.19 / 2021-04-14

* [Added] Add `reflow_type` property to dashboard object. See [#841](https://github.com/DataDog/datadog-api-client-go/pull/841).
* [Added] Add security track and formulas and functions support for geomap dashboard widget. See [#837](https://github.com/DataDog/datadog-api-client-go/pull/837).
* [Added] Generate intake endpoints. See [#834](https://github.com/DataDog/datadog-api-client-go/pull/834).
* [Added] Add endpoint for listing all downtimes for the specified monitor. See [#828](https://github.com/DataDog/datadog-api-client-go/pull/828).
* [Added] Add `modified_at` attribute to user response v2 schema. See [#817](https://github.com/DataDog/datadog-api-client-go/pull/817).
* [Added] Add default environment loading in clients. See [#816](https://github.com/DataDog/datadog-api-client-go/pull/816).
* [Added] Add `passed`, `noSavingResponseBody`, `noScreenshot`, and `disableCors` fields to Synthetics. See [#815](https://github.com/DataDog/datadog-api-client-go/pull/815).
* [Added] Add compliance usage endpoint and compliance host statistics. See [#814](https://github.com/DataDog/datadog-api-client-go/pull/814).
* [Added] Add tag filter options for `/api/v{1,2}/metrics`. See [#813](https://github.com/DataDog/datadog-api-client-go/pull/813).
* [Added] Add usage fields for Heroku and OpenTelemetry. See [#810](https://github.com/DataDog/datadog-api-client-go/pull/810).
* [Added] Add `global_time_target` field to SLO widget. See [#808](https://github.com/DataDog/datadog-api-client-go/pull/808).
* [Added] Add method to export an API test in Synthetics. See [#807](https://github.com/DataDog/datadog-api-client-go/pull/807).
* [Added] Add metadata to usage top average metrics response. See [#806](https://github.com/DataDog/datadog-api-client-go/pull/806).
* [Added] Add median as valid aggregator for formulas and functions. See [#800](https://github.com/DataDog/datadog-api-client-go/pull/800).
* [Fixed] Browser Test message required. See [#803](https://github.com/DataDog/datadog-api-client-go/pull/803).
* [Changed] Return correct object in `GetBrowserTest` endpoint. See [#827](https://github.com/DataDog/datadog-api-client-go/pull/827).
* [Changed] Add agent rules in security monitoring rules queries. See [#809](https://github.com/DataDog/datadog-api-client-go/pull/809).

## 1.0.0-beta.18 / 2021-03-22

* [Added] Add `legend_layout` and `legend_columns` to timeseries widget definition. See [#791](https://github.com/DataDog/datadog-api-client-go/pull/791).

## 1.0.0-beta.17 / 2021-03-15

* [Added] Add support for multistep tests in Synthetics. See [#775](https://github.com/DataDog/datadog-api-client-go/pull/775).
* [Added] Add core web vitals to synthetics browser test results. See [#771](https://github.com/DataDog/datadog-api-client-go/pull/771).
* [Added] Add v2 metric tags and metric volumes endpoints. See [#769](https://github.com/DataDog/datadog-api-client-go/pull/769).
* [Added] Add new endpoints for browser and API tests in Synthetics. See [#762](https://github.com/DataDog/datadog-api-client-go/pull/762).
* [Changed] Update response schema for service level objective operation `GetSLOHistory`. See [#784](https://github.com/DataDog/datadog-api-client-go/pull/784).
* [Changed] Make query name required in formulas and functions queries. See [#774](https://github.com/DataDog/datadog-api-client-go/pull/774).

## 1.0.0-beta.16 / 2021-03-02

* [Added] Add groupby_simple_monitor option to monitors. See [#758](https://github.com/DataDog/datadog-api-client-go/pull/758).
* [Added] Allow formula and functions in query value requests. See [#756](https://github.com/DataDog/datadog-api-client-go/pull/756).
* [Added] Allow formula and functions in toplist requests. See [#753](https://github.com/DataDog/datadog-api-client-go/pull/753).
* [Added] Add slack resource. See [#744](https://github.com/DataDog/datadog-api-client-go/pull/744).
* [Added] Add detectionMethod and newValueOptions fields to security monitoring rules. See [#739](https://github.com/DataDog/datadog-api-client-go/pull/739).
* [Added] Expose "event-v2 alert" monitor type. See [#738](https://github.com/DataDog/datadog-api-client-go/pull/738).
* [Added] Add new US3 region. See [#737](https://github.com/DataDog/datadog-api-client-go/pull/737).
* [Added] Add org_name field to usage attribution response. See [#736](https://github.com/DataDog/datadog-api-client-go/pull/736).
* [Added] Add profile_metrics_query properties to dashboard widget requests. See [#728](https://github.com/DataDog/datadog-api-client-go/pull/728).
* [Added] Add geomap widget to dashboards. See [#720](https://github.com/DataDog/datadog-api-client-go/pull/720).
* [Added] Add v2 API for metric tag configuration. See [#718](https://github.com/DataDog/datadog-api-client-go/pull/718).
* [Added] Add Lambda invocations usage to response. See [#716](https://github.com/DataDog/datadog-api-client-go/pull/716).
* [Added] Remove unstable flag for logs apis. See [#709](https://github.com/DataDog/datadog-api-client-go/pull/709).
* [Fixed] Add missing tlsVersion and minTlsVersion to Synthetics assertion types. See [#731](https://github.com/DataDog/datadog-api-client-go/pull/731).
* [Fixed] Change analyzed_spans to spans in dashboard. See [#711](https://github.com/DataDog/datadog-api-client-go/pull/711).
* [Changed] Rename objects throughout the code for consistency. See [#724](https://github.com/DataDog/datadog-api-client-go/pull/724).
* [Changed] Rename objects for formula and functions to be more generic. See [#747](https://github.com/DataDog/datadog-api-client-go/pull/747).

## v1.0.0-beta.15 / 2021-02-08

* [Added] Add restricted roles to monitor update. See [#691](https://github.com/DataDog/datadog-api-client-go/pull/691).
* [Added] Add endpoint for IoT billing usage. See [#684](https://github.com/DataDog/datadog-api-client-go/pull/684).
* [Added] Add query parameters for SLO search endpoint. See [#682](https://github.com/DataDog/datadog-api-client-go/pull/682).
* [Added] Add fields for formula and function query definition and widget formulas. See [#680](https://github.com/DataDog/datadog-api-client-go/pull/680).
* [Added] Add global_time to time_window SLO widget. See [#675](https://github.com/DataDog/datadog-api-client-go/pull/675).
* [Added] Update required fields in SLO correction create and update requests. See [#668](https://github.com/DataDog/datadog-api-client-go/pull/668).
* [Fixed] Fix AWS tag filter delete request. See [#701](https://github.com/DataDog/datadog-api-client-go/pull/701).
* [Fixed] Remove unnecessary field from TimeSeriesFormulaAndFunctionEventQuery. See [#700](https://github.com/DataDog/datadog-api-client-go/pull/700).
* [Fixed] Fix unit format in SLO history response. See [#695](https://github.com/DataDog/datadog-api-client-go/pull/695).
* [Fixed] Change group_by from object to list of objects. See [#694](https://github.com/DataDog/datadog-api-client-go/pull/694).
* [Fixed] Fix location of monitor restricted roles. See [#687](https://github.com/DataDog/datadog-api-client-go/pull/687).
* [Fixed] Fix paging parameter names for logs aggregate queries. See [#681](https://github.com/DataDog/datadog-api-client-go/pull/681).

## v1.0.0-beta.14 / 2021-01-19

* [Added] Add log index creation. See [#662](https://github.com/DataDog/datadog-api-client-go/pull/662).
* [Added] Add SLO Corrections. See [#654](https://github.com/DataDog/datadog-api-client-go/pull/654).
* [Added] Add new live and rehydrated logs breakdowns for Usage API. See [#652](https://github.com/DataDog/datadog-api-client-go/pull/652).
* [Added] Add support for Synthetics variables from test. See [#641](https://github.com/DataDog/datadog-api-client-go/pull/641).
* [Fixed] Add additionalProperties: false to synthetics target field. See [#657](https://github.com/DataDog/datadog-api-client-go/pull/657).
* [Fixed] Fix missing field for synthetics variables from test. See [#649](https://github.com/DataDog/datadog-api-client-go/pull/649).
* [Changed] Extract key sorting enum to a specific schema in key management endpoint. See [#646](https://github.com/DataDog/datadog-api-client-go/pull/646).
* [Changed] Extract enum to specific schema in incidents endpoint. See [#650](https://github.com/DataDog/datadog-api-client-go/pull/650).
* [Changed] Fix some integer/number formats in Logs and Synthetics endpoints. See [#658](https://github.com/DataDog/datadog-api-client-go/pull/658).

## 1.0.0-beta.13 / 2021-01-06

* [Added] Added filters to rule endpoints in security monitoring API. See [#632](https://github.com/DataDog/datadog-api-client-go/pull/632).
* [Added] Add Azure app services fields to usage v1 endpoints. See [#631](https://github.com/DataDog/datadog-api-client-go/pull/631).
* [Added] Add mobile RUM OS types usage fields. See [#629](https://github.com/DataDog/datadog-api-client-go/pull/629).
* [Added] Add config variables for synthetics API tests. See [#628](https://github.com/DataDog/datadog-api-client-go/pull/628).
* [Added] Add endpoints for the public API of Logs2Metrics. See [#626](https://github.com/DataDog/datadog-api-client-go/pull/626).
* [Added] Add endpoints for API Keys v2. See [#620](https://github.com/DataDog/datadog-api-client-go/pull/620).
* [Added] Add utils to validate and create valid enums. See [#617](https://github.com/DataDog/datadog-api-client-go/pull/617).
* [Added] Add javascript value to synthetics browser variable types. See [#616](https://github.com/DataDog/datadog-api-client-go/pull/616).
* [Added] Add synthetics assertion operator. See [#609](https://github.com/DataDog/datadog-api-client-go/pull/609).
* [Added] Application keys v2 API. See [#605](https://github.com/DataDog/datadog-api-client-go/pull/605).
* [Fixed] Redact auth methods from debug logs. See [#618](https://github.com/DataDog/datadog-api-client-go/pull/618).
* [Removed] Remove Synthetic resources property. See [#622](https://github.com/DataDog/datadog-api-client-go/pull/622).

## 1.0.0-beta.12 / 2020-12-07

* [Added] Mark Usage Attribution endpoint as public beta. See [#592](https://github.com/DataDog/datadog-api-client-go/pull/592).
* [Added] Add AWS filtering endpoints. See [#589](https://github.com/DataDog/datadog-api-client-go/pull/589).
* [Added] Add limit parameter for get usage top average metrics. See [#586](https://github.com/DataDog/datadog-api-client-go/pull/586).
* [Added] Add endpoint to fetch process summaries. See [#585](https://github.com/DataDog/datadog-api-client-go/pull/585).
* [Added] Add synthetics private location endpoints. See [#584](https://github.com/DataDog/datadog-api-client-go/pull/584).
* [Added] Add user_update, recommendation and snapshot as event alert types. See [#583](https://github.com/DataDog/datadog-api-client-go/pull/583).
* [Added] Add Usage Attribution endpoint. See [#582](https://github.com/DataDog/datadog-api-client-go/pull/582).
* [Added] Add new API for incident management usage. See [#578](https://github.com/DataDog/datadog-api-client-go/pull/578).
* [Added] Add the incident schema. See [#572](https://github.com/DataDog/datadog-api-client-go/pull/572).
* [Added] Add IP prefixes by location for synthetics endpoints. See [#565](https://github.com/DataDog/datadog-api-client-go/pull/565).
* [Added] Add filter parameter for listing teams and services. See [#564](https://github.com/DataDog/datadog-api-client-go/pull/564).
* [Added] Add restricted roles to monitor create and edit requests. See [#562](https://github.com/DataDog/datadog-api-client-go/pull/562).
* [Fixed] Quota & retention are now editable fields in log indexes. See [#568](https://github.com/DataDog/datadog-api-client-go/pull/568).
* [Changed] Mark request bodies as required or explicitly optional. See [#598](https://github.com/DataDog/datadog-api-client-go/pull/598).
* [Changed] Deprecate subscription and billing fields in create organization endpoint. See [#588](https://github.com/DataDog/datadog-api-client-go/pull/588).
* [Changed] Mark query field as optional when searching logs. See [#577](https://github.com/DataDog/datadog-api-client-go/pull/577).
* [Changed] Change event_query property to use log query definition in dashboard widgets. See [#573](https://github.com/DataDog/datadog-api-client-go/pull/573).
* [Changed] Rename tracing without limits and traces usage endpoints. See [#561](https://github.com/DataDog/datadog-api-client-go/pull/561).
* [Removed] Remove org_id parameter from Usage Attribution endpoint. See [#594](https://github.com/DataDog/datadog-api-client-go/pull/594).

## v1.0.0-beta.11 / 2020-11-06

* [Added] Add 3 new palettes to the conditional formatting options. See [#554](https://github.com/DataDog/datadog-api-client-go/pull/554).

## v1.0.0-beta.10 / 2020-11-02

* [Changed] Change teams and services objects names to be incident specific. See [#538](https://github.com/DataDog/datadog-api-client-go/pull/538).
* [Removed] Remove `require_full_window` client default value for monitors. See [#540](https://github.com/DataDog/datadog-api-client-go/pull/540).

## 1.0.0-beta.9 / 2020-10-27

* [Added] Add missing synthetics step types. See [#534](https://github.com/DataDog/datadog-api-client-go/pull/534).
* [Added] Add include_tags in logs archives. See [#530](https://github.com/DataDog/datadog-api-client-go/pull/530).
* [Added] Add dns server and client certificate support to synthetics tests. See [#523](https://github.com/DataDog/datadog-api-client-go/pull/523).
* [Added] Add rehydration_tags property to the logs archives. See [#513](https://github.com/DataDog/datadog-api-client-go/pull/513).
* [Added] Add endpoint to reorder Logs Archives. See [#505](https://github.com/DataDog/datadog-api-client-go/pull/505).
* [Added] Add has_search_bar and cell_display_mode properties to table widget definition. See [#502](https://github.com/DataDog/datadog-api-client-go/pull/502).
* [Added] Add target_format property to the Logs attribute remapper . See [#501](https://github.com/DataDog/datadog-api-client-go/pull/501).
* [Added] Add dual y-axis configuration to time-series widget in Dashboard. See [#498](https://github.com/DataDog/datadog-api-client-go/pull/498).
* [Added] Mark logs aggregate endpoint as stable. See [#496](https://github.com/DataDog/datadog-api-client-go/pull/496).
* [Added] Add endpoint to get a Synthetics global variable. See [#489](https://github.com/DataDog/datadog-api-client-go/pull/489).
* [Added] Add assertion types for DNS Synthetics tests. See [#486](https://github.com/DataDog/datadog-api-client-go/pull/486).
* [Added] Add DNS test type to Synthetics. See [#482](https://github.com/DataDog/datadog-api-client-go/pull/482).
* [Added] Add API endpoints for teams and services. See [#470](https://github.com/DataDog/datadog-api-client-go/pull/470).
* [Added] Add mobile_rum_session_count_sum property to usage responses. See [#469](https://github.com/DataDog/datadog-api-client-go/pull/469).
* [Fixed] Fix synthetics_check_id type in MonitorOptions. See [#526](https://github.com/DataDog/datadog-api-client-go/pull/526).
* [Fixed] Remove default for cell_display_mode in table widget. See [#519](https://github.com/DataDog/datadog-api-client-go/pull/519).
* [Fixed] Fix tags attribute type in event aggregation API. See [#463](https://github.com/DataDog/datadog-api-client-go/pull/463).
* [Changed] Change `columns` attribute type from string array to object array in APM stats query widget. See [#509](https://github.com/DataDog/datadog-api-client-go/pull/509).
* [Changed] Rename to ApmStats and add required properties. See [#490](https://github.com/DataDog/datadog-api-client-go/pull/490).
* [Changed] Remove unused `aggregation_key` and `related_event_id` properties from events responses. See [#480](https://github.com/DataDog/datadog-api-client-go/pull/480).
* [Changed] Define required fields for v2 requests. See [#475](https://github.com/DataDog/datadog-api-client-go/pull/475).
* [Changed] Mark required type fields in User and Roles API v2. See [#467](https://github.com/DataDog/datadog-api-client-go/pull/467).
* [Removed] Remove check_type parameter from ListTests endpoint. See [#465](https://github.com/DataDog/datadog-api-client-go/pull/465).

## v1.0.0-beta.8 / 2020-09-16

* [Added] Add `aggregation` and `metric` fields to `SecurityMonitoringRuleQuery`. See [#457](https://github.com/DataDog/datadog-api-client-go/pull/457).
* [Added] Add tracing without limits to usage API. See [#449](https://github.com/DataDog/datadog-api-client-go/pull/449).
* [Added] Add response codes for AWS API. See [#443](https://github.com/DataDog/datadog-api-client-go/pull/443).
* [Added] Add `custom_links` support for Dashboard widgets. See [#442](https://github.com/DataDog/datadog-api-client-go/pull/442).
* [Added] Add profiling to usage API. See [#436](https://github.com/DataDog/datadog-api-client-go/pull/436).
* [Added] Add synthetics CI endpoint. See [#429](https://github.com/DataDog/datadog-api-client-go/pull/429).
* [Added] Add APM resources data source to table widgets. See [#428](https://github.com/DataDog/datadog-api-client-go/pull/428).
* [Added] Add list API for security monitoring signals. See [#424](https://github.com/DataDog/datadog-api-client-go/pull/424).
* [Added] Add create, edit and delete endpoints for synthetics global variables. See [#421](https://github.com/DataDog/datadog-api-client-go/pull/421).
* [Added] Add monitor option `renotify_interval` to synthetics tests. See [#420](https://github.com/DataDog/datadog-api-client-go/pull/420).
* [Added] Add event aggregation v2 API. See [#419](https://github.com/DataDog/datadog-api-client-go/pull/419).
* [Added] Add Profiling Host to Usage endpoint. See [#417](https://github.com/DataDog/datadog-api-client-go/pull/417).
* [Added] Add `distinctFields` to `SecurityMonitoringRuleQuery`. See [#412](https://github.com/DataDog/datadog-api-client-go/pull/412).
* [Added] Add missing `security_query` on `QueryValueWidgetRequest`. See [#407](https://github.com/DataDog/datadog-api-client-go/pull/407).
* [Added] Enable security source for dashboards. See [#403](https://github.com/DataDog/datadog-api-client-go/pull/403).
* [Added] Add SLO alerts to monitor enum. See [#401](https://github.com/DataDog/datadog-api-client-go/pull/401).
* [Fixed] Add 200 response code to PATCH v2 users. See [#441](https://github.com/DataDog/datadog-api-client-go/pull/441).
* [Fixed] Fix hourly host usage descriptions. See [#438](https://github.com/DataDog/datadog-api-client-go/pull/438).
* [Fixed] Remove enum from `legend_size` widget attribute. See [#432](https://github.com/DataDog/datadog-api-client-go/pull/432).
* [Fixed] Fix content-type spelling errors. See [#423](https://github.com/DataDog/datadog-api-client-go/pull/423).
* [Fixed] Properly mark `status` and `query` field as required for creation of Security Monitoring rule. See [#422](https://github.com/DataDog/datadog-api-client-go/pull/422).
* [Fixed] Fix name of `isEnabled` parameter for Security Monitoring rule. See [#409](https://github.com/DataDog/datadog-api-client-go/pull/409).
* [Removed] Remove 204 response from PATCH v2 users. See [#446](https://github.com/DataDog/datadog-api-client-go/pull/446).

## v1.0.0-beta.7 / 2020-07-22

* [Added] Adding four usage attribution endpoints. See [#393](https://github.com/DataDog/datadog-api-client-go/pull/393).
* [Added] Fix documentation for `v1/hosts`. See [#383](https://github.com/DataDog/datadog-api-client-go/pull/383).
* [Changed] Update synthetics test to contain latest features. See [#375](https://github.com/DataDog/datadog-api-client-go/pull/375).
* [Added] Usage Billable Summary response. See [#368](https://github.com/DataDog/datadog-api-client-go/pull/368).
* [Added] Add Logs Search API v2. See [#365](https://github.com/DataDog/datadog-api-client-go/pull/365).
* [Fixed] RRULE property for Downtimes API. See [#364](https://github.com/DataDog/datadog-api-client-go/pull/364).
* [Deprecated] Dashboards List v1 has been deprecated. See [#363](https://github.com/DataDog/datadog-api-client-go/pull/363).

## v1.0.0-beta.6 / 2020-06-19

* [Fixed] Update enum of synthetics devices IDs to match API. See [#351](https://github.com/DataDog/datadog-api-client-go/pull/351).

## v1.0.0-beta.5 / 2020-06-19

* [Added] Update to the latest openapi-generator 5 snapshot. See [#338](https://github.com/DataDog/datadog-api-client-go/pull/338).
* [Added] Add synthetics location endpoint. See [#334](https://github.com/DataDog/datadog-api-client-go/pull/334).
* [Fixed] Widget legend size can also be "0". See [#336](https://github.com/DataDog/datadog-api-client-go/pull/336).
* [Fixed] Log Index as an optional parameter (default to "*") for List Queries. See [#335](https://github.com/DataDog/datadog-api-client-go/pull/335).
* [Changed] Rename payload objects to request for `users` v2 API. See [#346](https://github.com/DataDog/datadog-api-client-go/pull/346).
  * This change includes backwards incompatible changes when using the `users` v2 endpoint.
* [Changed] Split schema for roles API. See [#337](https://github.com/DataDog/datadog-api-client-go/pull/337).
  * This change includes backwards incompatible changes when using the `role` endpoint.

## v1.0.0-beta.4 / 2020-06-09

* [BREAKING] Add missing values to enums. See [#320](https://github.com/DataDog/datadog-api-client-go/pull/320).
    * This change includes backwards incompatible changes when using the `MonitorSummary` widget.
* [BREAKING] Split schemas from DashboardList v2. See [#318](https://github.com/DataDog/datadog-api-client-go/pull/318).
    * This change includes backwards incompatible changes when using corresponding endpoints methods.
* [BREAKING] Clean synthetics test CRUD endpoints. See [#317](https://github.com/DataDog/datadog-api-client-go/pull/317).
    * This change includes backwards incompatible changes when using corresponding endpoints methods.
* [Added] Add Logs Archives endpoints. See [#323](https://github.com/DataDog/datadog-api-client-go/pull/323).

## v1.0.0-beta.3 / 2020-05-21

* [BREAKING] Update to openapi-generator 5.0.0. See [#303](https://github.com/DataDog/datadog-api-client-go/pull/303).
    * This change includes backwards incompatible changes when using structs generated from `oneOf` schemas.
* [Added] Add SIEM and SNMP usage API. See [#309](https://github.com/DataDog/datadog-api-client-go/pull/309).
* [Added] Add security monitoring to clients. See [#304](https://github.com/DataDog/datadog-api-client-go/pull/304).
* [Added] Add /v1/validate endpoint. See [#290](https://github.com/DataDog/datadog-api-client-go/pull/290).
* [Added] Add generated_files file. See [#270](https://github.com/DataDog/datadog-api-client-go/pull/270).
* [Fixed] Add authentication to Go examples. See [#299](https://github.com/DataDog/datadog-api-client-go/pull/299).
* [Fixed] Add 422 error codes to users and roles v2 endpoints. See [#296](https://github.com/DataDog/datadog-api-client-go/pull/296).
* [Fixed] Update import in Go examples. See [#295](https://github.com/DataDog/datadog-api-client-go/pull/295).
* [Fixed] Check duplicate object definitions. See [#288](https://github.com/DataDog/datadog-api-client-go/pull/288).
* [Fixed] Mark unstable endpoints with beta note. See [#281](https://github.com/DataDog/datadog-api-client-go/pull/281).
* [Changed] Update ServiceLevelObjective schema names. See [#279](https://github.com/DataDog/datadog-api-client-go/pull/279).
* [Deprecated] Add deprecated fields `logset`, `count` and `start` to appropriate dashboard widgets. See [#285](https://github.com/DataDog/datadog-api-client-go/pull/285).

## v1.0.0-beta.2 / 2020-05-04

* [Added] Add RUM Monitor Type and update documentation. See [#273](https://github.com/DataDog/datadog-api-client-go/pull/273).
* [Added] Add Logs Pipeline Processor. See [#268](https://github.com/DataDog/datadog-api-client-go/pull/268).
* [Added] Add additional fields to synthetics test request. See [#262](https://github.com/DataDog/datadog-api-client-go/pull/262).
* [Added] Add Monitor Pagination. See [#253](https://github.com/DataDog/datadog-api-client-go/pull/253).
* [Fixed] Mark synthetics test request "method" and "url" as optional. See [#265](https://github.com/DataDog/datadog-api-client-go/pull/265).
* [Fixed] Update error responses for roles v2 endpoints. See [#248](https://github.com/DataDog/datadog-api-client-go/pull/248).
* [Fixed] Add missing ListSLO's 404 response. See [#245](https://github.com/DataDog/datadog-api-client-go/pull/245).
* [Removed] Remove Pagerduty endpoints from the client. See [#264](https://github.com/DataDog/datadog-api-client-go/pull/264).

## 1.0.0-beta.1 / 2020-04-22

* [Added] Initial beta release of the Datadog API Client
