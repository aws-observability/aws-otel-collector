# Change Log
## [3.28.1](https://github.com/vultr/govultr/compare/v3.28.0...v3.28.1) (2026-02-25)
### Enhancements
* Block Storage: Add missing snapshot ID create param [PR 442](https://github.com/vultr/govultr/pull/442)

## [3.28.0](https://github.com/vultr/govultr/compare/v3.27.0...v3.28.0) (2026-02-25)
### Enhancements
* Block Storage: Add fields for bootable block creation [PR 437](https://github.com/vultr/govultr/pull/437)
* Instances: Add fields for bootable block storage creation [PR 437](https://github.com/vultr/govultr/pull/437)
* Block Storage: Add additional bootable instance fields [PR 440](https://github.com/vultr/govultr/pull/440)
* Database: Add pending charges field [PR 440](https://github.com/vultr/govultr/pull/440)
* Regions: Add connectivity field [PR 440](https://github.com/vultr/govultr/pull/440)

### Dependencies
* Bump github.com/google/go-querystring from 1.1.0 to 1.2.0 [PR 426](https://github.com/vultr/govultr/pull/426)

### Automation
* Add golangci-lint exclusion for gosec g117 [PR 438](https://github.com/vultr/govultr/pull/438)

## [3.27.0](https://github.com/vultr/govultr/compare/v3.26.1...v3.27.0) (2026-02-05)
### Enhancements
* Kubernetes: Add cluster OIDC config fields [PR 432](https://github.com/vultr/govultr/pull/432)
* VPC: Add NAT gateways [PR 433](https://github.com/vultr/govultr/pull/433)

## [3.26.1](https://github.com/vultr/govultr/compare/v3.26.0...v3.26.1) (2025-12-18)
### Enhancements
* User: Add service user create parameter [PR 418](https://github.com/vultr/govultr/pull/418)

## [3.26.0](https://github.com/vultr/govultr/compare/v3.25.0...3.26.0) (2025-12-11)
### Enhancements
* Kubernetes: Add specialized taints & labels functions [PR 415](https://github.com/vultr/govultr/pull/415)

### Clean Up
* Kubernetes: Resolve lint issues [PR 419](https://github.com/vultr/govultr/pull/419)
* Domains: Resolve lint issues [PR 420](https://github.com/vultr/govultr/pull/420)

### Bug Fix
* Domains: Make request structs omit empty name field [PR 401](https://github.com/vultr/govultr/pull/401)

### New Contributors
* @pnx made their first contribution in [PR 401](https://github.com/vultr/govultr/pull/401)

## [3.25.0](https://github.com/vultr/govultr/compare/v3.24.0...v3.25.0) (2025-11-21)
### Enhancement
* Bare Metal Server: Add vpc_only field [PR 404](https://github.com/vultr/govultr/pull/404)
* Instance: Add vpc_only field [PR 404](https://github.com/vultr/govultr/pull/404)
* Container Registry: Add support for artifact, robot, replication and retention endpoints [PR 406](https://github.com/vultr/govultr/pull/406)

### New Contributors
* @nathangoulding made their first contribution in [PR 404](https://github.com/vultr/govultr/pull/404)

## [3.24.0](https://github.com/vultr/govultr/compare/v3.23.0...v3.24.0) (2025-09-19)
### Enhancements
* Logs: Added to library [PR 399](https://github.com/vultr/govultr/pull/399)

### Automation
* Migrate golangci-lint config to v2 [PR 398](https://github.com/vultr/govultr/pull/398)

## [3.23.0](https://github.com/vultr/govultr/compare/v3.22.1...v3.23.0) (2025-08-27)
### Enhancements
* Bare Metal Server: Add snapshot ID field [PR 396](https://github.com/vultr/govultr/pull/396)
* Instance: Add snapshot ID field [PR 396](https://github.com/vultr/govultr/pull/396)
* Database: Add CA certificate field [PR 396](https://github.com/vultr/govultr/pull/396)

## [3.22.1](https://github.com/vultr/govultr/compare/v3.22.0...v3.22.1) (2025-08-11)
### Enhancements
* Snapshots: Add support for CreateFromURL UEFI field [PR 390](https://github.com/vultr/govultr/pull/390)

### New Contributors
* @mmonaco made their first contribution in [PR 390](https://github.com/vultr/govultr/pull/390)

## [3.22.0](https://github.com/vultr/govultr/compare/v3.21.1...v3.22.0) (2025-08-11)
### Enhancements
* Kubernetes: Add node pool user_data to create/update structs [PR 391](https://github.com/vultr/govultr/pull/391)

## [3.21.1](https://github.com/vultr/govultr/compare/v3.21.0...v3.21.1) (2025-07-16
### Bug Fixes
* Load Balancer: Add missing auto SSL domain field [PR 388](https://github.com/vultr/govultr/pull/388)

## [3.21.0](https://github.com/vultr/govultr/compare/v3.20.0...v3.21.0) (2025-06-20)
### Enhancements
* Database: Add Kafka additional features [PR 386](https://github.com/vultr/govultr/pull/386)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.7 to 0.7.8 [PR 385](https://github.com/vultr/govultr/pull/385)

### Automation
* Update goreleaser to lock to v2 [PR 384](https://github.com/vultr/govultr/pull/384)

## [3.20.0](https://github.com/vultr/govultr/compare/v3.19.1...v3.20.0) (2025-05-06)
### Enhancements
* Database: Add backup schedule [PR 382](https://github.com/vultr/govultr/pull/382)

## [3.19.1](https://github.com/vultr/govultr/compare/v3.19.0...v3.19.1) (2025-04-08)
### Bug Fixes
* Kubernetes: Allow node pool update to properly delete labels [PR 380](https://github.com/vultr/govultr/pull/380)

## [3.19.0](https://github.com/vultr/govultr/compare/v3.18.0...v3.19.0) (2025-04-08)
### Enhancements
* Kubernetes: Add VPC ID create param [PR 378](https://github.com/vultr/govultr/pull/378)

### Bug Fixes
* Bare Metal Server: Fix URI of VPC detach [PR 377](https://github.com/vultr/govultr/pull/377)

## [3.18.0](https://github.com/vultr/govultr/compare/v3.17.0...v3.18.0) (2025-03-28)
### Enhancements
* Bare Metal: Add iPXE URL to bare metal [PR 372](https://github.com/vultr/govultr/pull/372)
* Govultr: Increase default timeouts [PR 373](https://github.com/vultr/govultr/pull/373)

## [3.17.0](https://github.com/vultr/govultr/compare/v3.16.1...v3.17.0) (2025-03-13)
### Enhancements
* Kubernetes: Add support for taints to node pools [PR 363](https://github.com/vultr/govultr/pull/363)

### New Contributors
* @biondizzle made their first contribution in [PR 363](https://github.com/vultr/govultr/pull/363)

## [3.16.1](https://github.com/vultr/govultr/compare/v3.16.0...v3.16.1) (2025-03-10)
### Bug Fixes
* Object Storage: Make create & update requests use request struct [PR 369](https://github.com/vultr/govultr/pull/369)

## [3.16.0](https://github.com/vultr/govultr/compare/v3.15.0...v3.16.0) (2025-03-10)
### Enhancements
* Object Storage: Add tiers [PR 365](https://github.com/vultr/govultr/pull/365)

### Documentation
* Bare Metal: Add notice for deprecation of VPC2 [PR 364](https://github.com/vultr/govultr/pull/364)
* Instance: Add notice for deprecation of VPC2 [PR 364](https://github.com/vultr/govultr/pull/364)
* VPC2: Add notice for deprecation of VPC2[PR 364](https://github.com/vultr/govultr/pull/364)

### Dependencies
* Reset min go.mod go version to v1.23 [PR 367](https://github.com/vultr/govultr/pull/367)

## [3.15.0](https://github.com/vultr/govultr/compare/v3.14.1...v3.15.0) (2025-03-04)
### Enhancements
* Add functions for account bandwidth [PR 356](https://github.com/vultr/govultr/pull/356)
* Add functions for billing pending charges [PR 357](https://github.com/vultr/govultr/pull/357)

### Dependencies
* Update Go from v1.23 to v1.24 [PR 360](https://github.com/vultr/govultr/pull/360)

### Clean Up
* Fix whitespace [PR 359](https://github.com/vultr/govultr/pull/359)

### New Contributors
* @DazWilkin made their first contribution in [PR 356](https://github.com/vultr/govultr/pull/356)

## [3.14.1](https://github.com/vultr/govultr/compare/v3.14.0...v3.14.1) (2025-01-17)
### Bug Fixes
* Load Balancers: Add missing SSL fields and AutoSSL struct [PR 352](https://github.com/vultr/govultr/pull/352)

## [3.14.0](https://github.com/vultr/govultr/compare/v3.13.0...v3.14.0) (2025-01-15)
### Enhancements
* Load Balancers: Add support for AutoSSL and GlobalRegions [PR 350](https://github.com/vultr/govultr/pull/350)

### Clean Up
* Database: Remove deprecated Redis references [PR 349](https://github.com/vultr/govultr/pull/349)

## [3.13.0](https://github.com/vultr/govultr/compare/v3.12.0...v3.13.0) (2024-12-17)
### Enhancements
* VFS Storage: Add support for virtual file system storages [PR 347](https://github.com/vultr/govultr/pull/347)

## [3.12.0]( https://github.com/vultr/govultr/compare/v3.11.2...v3.12.0) (2024-11-22)
### Clean Up
* Database: Deprecate Redis-Named Fields [PR 345](https://github.com/vultr/govultr/pull/345)

## [3.11.2]( https://github.com/vultr/govultr/compare/v3.11.1...v3.11.2) (2024-10-30)
### Bug fixes
* Database: Fix Kafka advanced config field names [PR 343](https://github.com/vultr/govultr/pull/343)

## [3.11.1](https://github.com/vultr/govultr/compare/v3.11.0...v3.11.1) (2024-10-24)
### Bug Fixes
* Database: Fix quota create endpoint param for Kafka [PR 341](https://github.com/vultr/govultr/pull/341)

## [3.11.0](https://github.com/vultr/govultr/compare/v3.10.0...v3.11.0) (2024-10-22)
### Enhancements
* Database: Add support for managed Kafka [PR 337](https://github.com/vultr/govultr/pull/337)

### Automation
* Remove deprecated exportloopref linter [PR 338](https://github.com/vultr/govultr/pull/338)
* Add Github CODEOWNERS file [PR 339](https://github.com/vultr/govultr/pull/339)

## [3.10.0](https://github.com/vultr/govultr/compare/v3.9.1...v3.10.0) (2024-10-10)
### Enhancements
* Load Balancers: Add HTTP2/3 and timeout options [PR 325](https://github.com/vultr/govultr/pull/325)
* CDN: add vanity domain and SSL options [PR 326](https://github.com/vultr/govultr/pull/326)
* Instance: add user scheme field [PR 328](https://github.com/vultr/govultr/pull/328)
* Bare Metal: add user scheme field [PR 335](https://github.com/vultr/govultr/pull/335)
* Sub-accounts: Add support for sub-accounts [PR 329](https://github.com/vultr/govultr/pull/329)

### Bug Fixes
* CDN: Fix default empty struct fields and list files endpoint URL [PR 330](https://github.com/vultr/govultr/pull/330)
* Sub-accounts: Fix base struct on create [PR 331](https://github.com/vultr/govultr/pull/331)

### Clean Up
* Remove deprecated private network functionality [PR 324](https://github.com/vultr/govultr/pull/324)
* Bare Metal: remove deprecated tag fields [PR 327](https://github.com/vultr/govultr/pull/327)
* Firewall: remove deprecated type fields [PR 327](https://github.com/vultr/govultr/pull/327)
* Instance: remove deprecated tag fields [PR 327](https://github.com/vultr/govultr/pull/327)

### Dependencies & Automation
* Update go from v1.21 to v1.23 [PR 333](https://github.com/vultr/govultr/pull/333)
* Update github workflows to go 1.23 [PR 334](https://github.com/vultr/govultr/pull/334)
* Add CDN & Sub-account tests [PR 332](https://github.com/vultr/govultr/pull/332)

## [3.9.1](https://github.com/vultr/govultr/compare/v3.9.0...v3.9.1) (2024-08-13)
### Enhancements
* Database: Add support for Managed MySQL advanced config [PR 322](https://github.com/vultr/govultr/pull/322)

### Bug Fixes
* CDN: Fix http method used on push zone file delete [PR 321](https://github.com/vultr/govultr/pull/321)

## [3.9.0](https://github.com/vultr/govultr/compare/v3.8.1...v3.9.0) (2024-06-30)
### Enhancements
* CDN: Full support added [PR 319](https://github.com/vultr/govultr/pull/319)

## [3.8.1](https://github.com/vultr/govultr/compare/v3.8.0...v3.8.1) (2023-06-06)
### Automation
* Update goreleaser github action from v2 to v6 [PR 317](https://github.com/vultr/govultr/pull/317)

## [3.8.0](https://github.com/vultr/govultr/compare/v3.7.0...v3.8.0) (2023-06-06)
### Enhancements
* Inference: add support for serverless inference endpoints [PR 315](https://github.com/vultr/govultr/pull/315)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.6 to 0.7.7 [PR 312](https://github.com/vultr/govultr/pull/312)

### Automation
* Update golangci-lint and enable most go-vet checks [PR 314](https://github.com/vultr/govultr/pull/314)

## [3.7.0](https://github.com/vultr/govultr/compare/v3.6.4...v3.7.0) (2024-05-28)
### Enhancements
* Bare Metal: Add MdiskMode to BareMetalCreate and BareMetalUpdate structs [PR 310](https://github.com/vultr/govultr/pull/310)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.5 to 0.7.6 [PR 308](https://github.com/vultr/govultr/pull/308)
* Update Go from v1.20 to v1.21 [PR 309](https://github.com/vultr/govultr/pull/309)

### Automation
* Update and fix mattermost notifications [PR 305](https://github.com/vultr/govultr/pull/305)
* Fix mattermost notifications [PR 307](https://github.com/vultr/govultr/pull/307)

### New Contributors
* @fjoenichols made their first contribution in [PR 310](https://github.com/vultr/govultr/pull/310)

## [3.6.4](https://github.com/vultr/govultr/compare/v3.6.3...v3.6.4) (2024-03-07)
### Enhancement
* Kubernetes: add labels to the node pool update request struct [PR 302](https://github.com/vultr/govultr/pull/302)

## [3.6.3](https://github.com/vultr/govultr/compare/v3.6.2...v3.6.3) (2024-02-29)
### Bug Fixes
* All: add a json struct tag for the meta links field [PR 298](https://github.com/vultr/govultr/pull/298)
* Startup Scripts: allow empty values in create/update request struct [PR 299](https://github.com/vultr/govultr/pull/299)

### Automation
* Update codeql runner from v1 to v2 [PR 300](https://github.com/vultr/govultr/pull/300)

## [v3.6.2](https://github.com/vultr/govultr/compare/v3.6.1...v3.6.2) (2024-02-20)
### Enhancement
* Kubernetes: add support for nodepool labels [PR 296](https://github.com/vultr/govultr/pull/296)

## [v3.6.1](https://github.com/vultr/govultr/compare/v3.6.0...v3.6.1) (2024-01-16)
### Enhancement
* Bare Metal: add functions to allow VPC [PR 293](https://github.com/vultr/govultr/pull/293)

## [v3.6.0](https://github.com/vultr/govultr/compare/v3.5.0...v3.6.0) (2023-12-15)
### Enhancement
* Marketplace: add support for new API route [PR 291](https://github.com/vultr/govultr/pull/291)
* Bare Metal: add marketplace app variables [PR 291](https://github.com/vultr/govultr/pull/291)
* Instance: add marketplace app variables [PR 291](https://github.com/vultr/govultr/pull/291)

### Documentation
* Update the README with non-auth client examples [PR 290](https://github.com/vultr/govultr/pull/290)

## [v3.5.0](https://github.com/vultr/govultr/compare/v3.4.1...v3.5.0) (2023-12-01)
### Enhancement
* Instance: Add disable IPv4 option create request [PR 287](https://github.com/vultr/govultr/pull/287)
* Database: Add user access control for Redis [PR 288](https://github.com/vultr/govultr/pull/288)

## [v3.4.1](https://github.com/vultr/govultr/compare/v3.4.0...v3.4.1) (2023-11-17)
### Enhancement
* Database: Add support for usage endpoint [PR 282](https://github.com/vultr/govultr/pull/282)

### Bug Fix
* Container Registry: minor API fixes [PR 284](https://github.com/vultr/govultr/pull/284)

## [v3.4.0](https://github.com/vultr/govultr/compare/v3.3.4...v3.4.0) (2023-11-10)
### Enhancements
* Database: Support read replica promotion [PR 276](https://github.com/vultr/govultr/pull/276)
* Kubernetes: Add managed firewall support [PR 277](https://github.com/vultr/govultr/pull/277)
* Container Registry: Add support for container registry operations [PR 278](https://github.com/vultr/govultr/pull/278)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.4 to 0.7.5 [PR 280](https://github.com/vultr/govultr/pull/280)

### New Contributors
* @Byteflux made their first contribution in [PR 277](https://github.com/vultr/govultr/pull/277)

## [v3.3.4](https://github.com/vultr/govultr/compare/v3.3.3...v3.3.4) (2023-10-30)
### Enhancements
* Database: Add support for FerretDB [PR 272](https://github.com/vultr/govultr/pull/272)
* Kubernetes: Add HA control planes support [PR 273](https://github.com/vultr/govultr/pull/273)

## [v3.3.3](https://github.com/vultr/govultr/compare/v3.3.2...v3.3.3) (2023-10-24)
### Bug Fixes
* Database: Change VPCID to pointer for empty & nil values in updates [PR 270](https://github.com/vultr/govultr/pull/270)

## [v3.3.2](https://github.com/vultr/govultr/compare/v3.3.1...v3.3.2) (2023-10-23)
### Enhancements
* General: Remove references to deprecated V1 API [PR 266](https://github.com/vultr/govultr/pull/266)
* Database: Add support for public/private hostnames [PR 268](https://github.com/vultr/govultr/pull/268)

## [v3.3.1](https://github.com/vultr/govultr/compare/v3.3.0...v3.3.1) (2023-08-18)
### Enhancements
* VPC2: Add nodes endpoints [PR 263](https://github.com/vultr/govultr/pull/263)

## [v3.3.0](https://github.com/vultr/govultr/compare/v3.2.0...v3.3.0) (2023-08-10)
### Enhancements
* Add VPC2 [PR 261](https://github.com/vultr/govultr/pull/261)
* Bare Metal/Instances: Add support for VPC 2.0 [PR 261](https://github.com/vultr/govultr/pull/261)

### New Contributors
* @ogawa0071 made their first contribution in [PR 261](https://github.com/vultr/govultr/pull/261)

## [v3.2.0](https://github.com/vultr/govultr/compare/v3.1.0...v3.2.0) (2023-07-24)
### Enhancements
* Database: add support for DBaaS VPC networks [PR 255](https://github.com/vultr/govultr/pull/255)
* Implement stricter golangci-lint configurations [PR 259](https://github.com/vultr/govultr/pull/259)

## [v3.1.0](https://github.com/vultr/govultr/compare/v3.0.3...v3.1.0) (2023-07-13)
### Enhancements
* Loadbalancers: add support for multi-nodes in [PR 250](https://github.com/vultr/govultr/pull/250)

### New Contributors
* @happytreees made their first contribution in [PR 250](https://github.com/vultr/govultr/pull/250)

## [v3.0.3](https://github.com/vultr/govultr/compare/v3.0.2...v3.0.3) (2023-06-07)
### Enhancements
* Remove unused parameters from database update request [247](https://github.com/vultr/govultr/pull/247)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.2 to 0.7.4 [248](https://github.com/vultr/govultr/pull/248)

## [v3.0.2](https://github.com/vultr/govultr/compare/v3.0.1...v3.0.2) (2023-03-31)
### Bug fixes
* Allow empty password parameter on DBaaS user update [244](https://github.com/vultr/govultr/pull/244)

## [v3.0.1](https://github.com/vultr/govultr/compare/v2.17.2...v3.0.1) (2023-03-20)

### Enhancements
* Add golangci and update go version in workflows [235](https://github.com/vultr/govultr/pull/235)
* Fix context error in govultr test [239](https://github.com/vultr/govultr/pull/239)
* Add resource response functionality [240](https://github.com/vultr/govultr/pull/240)
* Add support for Vultr Managed Databases [238](https://github.com/vultr/govultr/pull/238)

### Dependencies
* Bump github.com/hashicorp/go-retryablehttp from 0.7.1 to 0.7.2 [236](https://github.com/vultr/govultr/pull/236)
* Update Go to v1.20 [241](https://github.com/vultr/govultr/pull/241)
* Update go to v1.19 [234](https://github.com/vultr/govultr/pull/234)

### New Contributors
* @mondragonfx made their first contribution in [240](https://github.com/vultr/govultr/pull/240)
* @christhemorse made their first contribution in [238](https://github.com/vultr/govultr/pull/238)

## [v2.17.2](https://github.com/vultr/govultr/compare/v2.17.1...v2.17.2) (2022-06-13)

### Enhancement
* Reserved IP: Add support for updating label [227](https://github.com/vultr/govultr/pull/227)

## [v2.17.1](https://github.com/vultr/govultr/compare/v2.17.0...v2.17.1) (2022-06-02)
* Plans: Add GPU specific fields to plan struct [224](https://github.com/vultr/govultr/pull/224)

## [v2.17.0](https://github.com/vultr/govultr/compare/v2.16.0..v2.17.0) (2022-05-17)

### Enhancement
* Kubernetes: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)
* Baremetal: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)
* Instance: allow `tag` update to delete existing value [222](https://github.com/vultr/govultr/pull/222)

### Bug fixes
* Kubernetes: fix data type for `auto_scaler` to avoid sending null values in requests when not set [222](https://github.com/vultr/govultr/pull/222)

### Breaking Change
* Kubernetes: change data type for `Tag` in node pool update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Kubernetes: change data type for `AutoScaler` in node pool update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Baremetal: change data type for `Tag` in update requirements struct [222](https://github.com/vultr/govultr/pull/222)
* Instance: change data type for `Tag` in update requirements struct [222](https://github.com/vultr/govultr/pull/222)


## [v2.16.0](https://github.com/vultr/govultr/compare/v2.15.1..v2.16.0) (2022-05-04)

### Enhancement
* Kubernetes: added auto scaler options to node pools [215](https://github.com/vultr/govultr/pull/215)
* Firewall rules: added new field `ip_type` in get/list responses to be consistent with the create calls [216](https://github.com/vultr/govultr/pull/216)
* Kubernetes: Upgrade support [217](https://github.com/vultr/govultr/pull/217)
* Baremetal: Added support for new `tags` field. This field allows multiple string tags to be associated with an instance [218](https://github.com/vultr/govultr/pull/218)
* Instance: Added support for new `tags` field. This field allows multiple string tags to be associated with an instance [218](https://github.com/vultr/govultr/pull/218)

### Deprecations
* Instance: The `tag` field has been deprecated in favor for `tags` [218](https://github.com/vultr/govultr/pull/218)
* Baremetal: The `tag` field has been deprecated in favor for `tags` [218](https://github.com/vultr/govultr/pull/218)
* Firewall rules: The `type` field has been deprecated in favor for `ip_type` [216](https://github.com/vultr/govultr/pull/216)

### Dependency Update
* Bump github.com/hashicorp/go-retryablehttp from 0.7.0 to 0.7.1 [214](https://github.com/vultr/govultr/pull/214)

## [v2.15.1](https://github.com/vultr/govultr/compare/v2.15.0..v2.15.1) (2022-04-12)
### Bug fixes
* Block : add `omityempty` to `block_type` to prevent deploy issues [212](https://github.com/vultr/govultr/pull/212)

## [v2.15.0](https://github.com/vultr/govultr/compare/v2.14.2..v2.15.0) (2022-04-12)
### Enhancement
* Block : New optional field `block_type`. This new field is currently optional but may become required at a later release [209](https://github.com/vultr/govultr/pull/209)
* VPC : New API endpoints that will be replacing `network` [210](https://github.com/vultr/govultr/pull/210)
* Updated Go version from 1.16 to 1.17 [208](https://github.com/vultr/govultr/pull/208)

### Deprecations
* Network : The network resource and all related private network fields on structs are deprecated. You should now be using the VPC provided replacements [210](https://github.com/vultr/govultr/pull/210)

## [v2.14.2](https://github.com/vultr/govultr/compare/v2.14.1..v2.14.2) (2022-03-23)
### Bug Fix
* Instances : restore support requestBody [206](https://github.com/vultr/govultr/pull/206) Thanks @andrake81

## [v2.14.1](https://github.com/vultr/govultr/compare/v2.14.0..v2.14.1) (2022-02-02)
### Enhancement
* Improved retry error response [204](https://github.com/vultr/govultr/pull/204)

## [v2.14.0](https://github.com/vultr/govultr/compare/v2.13.0..v2.14.0) (2022-01-21)
### Enhancement
* ListOptions : [Added query param Region](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [200](https://github.com/vultr/govultr/pull/200)
* ListOptions : [Added query param Description](https://www.vultr.com/api/#operation/list-snapshots) that can be used with `Snapshot.List`  [202](https://github.com/vultr/govultr/pull/202)
* Snapshot : `CreateFromURL` has new optional field called `description` which lets you set a custom description [202](https://github.com/vultr/govultr/pull/202)

## [v2.13.0](https://github.com/vultr/govultr/compare/v2.12.0..v2.13.0) (2022-01-05)
### Enhancement
* ListOptions : [Added query params](https://www.vultr.com/api/#operation/list-instances) that can be used with `Instance.List`  [197](https://github.com/vultr/govultr/pull/197)

## [v2.12.0](https://github.com/vultr/govultr/compare/v2.11.1..v2.12.0) (2021-12-01)
### Breaking Changes
* Plans : Changed `MonthlyCost` from `int` to `float32` [192](https://github.com/vultr/govultr/pull/192)

## [v2.11.1](https://github.com/vultr/govultr/compare/v2.11.0..v2.11.1) (2021-11-26)
### Bug fixes
* LoadBalancers : Fixed SSL struct json params to the proper API fields [189](https://github.com/vultr/govultr/pull/189)

## [v2.11.0](https://github.com/vultr/govultr/compare/v2.10.0..v2.11.0) (2021-11-18)
### Breaking Changes
* Instances : Update call will now return `*Instance` in addition to `error` [185](https://github.com/vultr/govultr/pull/185)
* Instances : Reinstall call now allows changing of hostname and also returns `*Instance` in addition to `error` [181](https://github.com/vultr/govultr/pull/181)

### Enhancement
* Instances : The hostname of the instance is now returned in any call that returns Instance data [187](https://github.com/vultr/govultr/pull/187)
* Domains : There is a new field called `dns_sec` which will return `enabled` or `disabled` depending on how your domain is configured [184](https://github.com/vultr/govultr/pull/184)

## [v2.10.0](https://github.com/vultr/govultr/compare/v2.9.2..v2.10.0) (2021-11-03)
### Enhancement
* Billing : Added support for billing [178](https://github.com/vultr/govultr/pull/178)

## [v2.9.2](https://github.com/vultr/govultr/compare/v2.9.1..v2.9.2) (2021-10-20)
### Change
* Iso : Changed `client` field to be unexported [168](https://github.com/vultr/govultr/pull/168)
* Snapshot : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Plans : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)
* Regions : Changed `client` field to be unexported  [168](https://github.com/vultr/govultr/pull/168)

## [v2.9.1](https://github.com/vultr/govultr/compare/v2.9.0..v2.9.1) (2021-10-18)
### Enhancement
* Kubernetes : Added `Tag` support for nodepools [164](https://github.com/vultr/govultr/pull/164)

## [v2.9.0](https://github.com/vultr/govultr/compare/v2.8.1..v2.9.0) (2021-09-27)
### Breaking Change
* Kubernetes : PlanID is now Plan and Count is now NodeQuantity to follow API pattern [161](https://github.com/vultr/govultr/pull/161)

### Enhancement
* Snapshots : Add compressed size field [162](https://github.com/vultr/govultr/pull/162)

## [v2.8.1](https://github.com/vultr/govultr/compare/v2.8.0..v2.8.1) (2021-08-31)
### Enhancement
* Kubernetes : Add support for deletion with resources [159](https://github.com/vultr/govultr/pull/159)
* Kubernetes : Add support for getting available versions[159](https://github.com/vultr/govultr/pull/159)

### Dependency Update
* Bump Go version to 1.16 [158](https://github.com/vultr/govultr/pull/158)

## [v2.8.0](https://github.com/vultr/govultr/compare/v2.7.1..v2.8.0) (2021-08-18)
### Enhancement
* Added support for Vultr Kubernetes Engine [156](https://github.com/vultr/govultr/pull/156)

## [v2.7.1](https://github.com/vultr/govultr/compare/v2.7.0..v2.7.1) (2021-07-15)
### Enhancement
* BareMetal : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)
* Instances : Add support for `image_id` on update [152](https://github.com/vultr/govultr/pull/152)

## [v2.7.0](https://github.com/vultr/govultr/compare/v2.6.0..v2.7.0) (2021-07-14)
### Enhancement
* BareMetal : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Instances : Add support for `image_id` [150](https://github.com/vultr/govultr/pull/150)
* Applications : added support for marketplace applications [150](https://github.com/vultr/govultr/pull/150)

## [v2.6.0](https://github.com/vultr/govultr/compare/v2.5.1..v2.6.0) (2021-07-02)
### Enhancement
* BareMetal : Add support for `persistent_pxe` [148](https://github.com/vultr/govultr/pull/148)

## [v2.5.1](https://github.com/vultr/govultr/compare/v2.5.0..v2.5.1) (2021-05-10)
### Bug fix
* Instances : BackupScheduleReq change DOW + Hour to pointers  [145](https://github.com/vultr/govultr/pull/145)

## [v2.5.0](https://github.com/vultr/govultr/compare/v2.4.2..v2.5.0) (2021-05-06)
### Enhancement
* LoadBalancers : New Features and endpoints [143](https://github.com/vultr/govultr/pull/143)
  * Ability to attach private networks
  * Ability to set firewalls
  * Get Firewall Rules
  * List Firewall Rules

## [v2.4.2](https://github.com/vultr/govultr/compare/v2.4.1..v2.4.2) (2021-05-03)
### Bug fix
* Instances : ListPrivateNetworks missing paging ability [140](https://github.com/vultr/govultr/pull/140)

## [v2.4.1](https://github.com/vultr/govultr/compare/v2.4.0..v2.4.1) (2021-05-03)
### Dependency updates
* Bump github.com/hashicorp/go-retryablehttp from 0.6.8 to 0.7.0 [138](https://github.com/vultr/govultr/pull/138)
* Bump github.com/google/go-querystring from 1.0.0 to 1.1.0 [137](https://github.com/vultr/govultr/pull/137)

## [v2.4.0](https://github.com/vultr/govultr/compare/v2.3.2..v2.4.0) (2021-02-11)
### Enhancement
* Block Storage - add `mount_id` field to BlockStorage struct [131](https://github.com/vultr/govultr/pull/131)
* Plans - add `disk_count` field to Plan and BareMetalPlan struct [130](https://github.com/vultr/govultr/pull/130)

## [v2.3.2](https://github.com/vultr/govultr/compare/v2.3.1..v2.3.2) (2021-01-06)
### Bug Fix
* Instances - Fixed DetachPrivateNetwork which had wrong URI [122](https://github.com/vultr/govultr/pull/122)

## [v2.3.1](https://github.com/vultr/govultr/compare/v2.3.0..v2.3.1) (2021-01-04)
### Bug Fix
* Domain Record - removed `omitempty` on `name` field in `DomainRecordReq` to allow creation of nameless records. [120](https://github.com/vultr/govultr/pull/120)

## [v2.3.0](https://github.com/vultr/govultr/compare/v2.2.0..v2.3.0) (2020-12-17)
### Enhancement
* Bare Metal - Start call added [118](https://github.com/vultr/govultr/pull/118)

## [v2.2.0](https://github.com/vultr/govultr/compare/v2.1.0..v2.2.0) (2020-12-07)
### Breaking Change
* All bools have been updated to pointers to avoid issues where false values not being sent in request. Thanks @Static-Flow [115](https://github.com/vultr/govultr/pull/115)

## [v2.1.0](https://github.com/vultr/govultr/compare/v2.0.0..v2.1.0) (2020-11-30)
### Bug fixes
* ReservedIP - Attach call creates proper json. [112](https://github.com/vultr/govultr/pull/112)
* User - APIEnabled takes pointer of bool [112](https://github.com/vultr/govultr/pull/112)

## v2.0.0 (2020-11-20)
### Initial Release
* GoVultr v2.0.0 Release - Uses Vultr API v2.
* GoVultr v1.0.0 is now on [branch v1](https://github.com/vultr/govultr/tree/v1)
