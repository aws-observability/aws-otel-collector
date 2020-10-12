## Release a new version

This document outlines the steps to release a new version of AWS OTel Collector
you can skip it if you are not a developer.

### Prerequisites

Fork the repo into your own repo

### Step 1. Bump Version in forked repo

run below command to bump version and generate release note, this command will also create a git commit for the generated files including version file and release note

```
export RELEASE_VERSION={{replace with the version you want to release}}
export GITHUB_USER={{replace with the upstream username}}
export GITHUB_TOKEN={{replace with the github token which has read permission to the upstream repo}}
./tools/release/bump-version-and-create-release-note.sh
```

After the script is done, push the commited change, create pull request to the master branch of the upstream repo, please don't forget to attach a "bumpversion" label to this pr.
### Step 2. Merge the pull request

Please reach out to the approver of this repo to merge the pull request you just created in step1, after that a github workflow will be triggered to perform integ-test.

### Step 3. Wait until the integ-test and soaking test are done
Wait until the integ-test in the workflow is done, wait until the soaking test is done. it normally takes 7 days in soaking test 

### Step 4. Create a Tag to trigger release in the upstream repo

switch to the upstream repo
create a tag upon the commit of bumping version, the tag name should be same as the version

```
VERSION={"replace it with the version to release"}
GITHUB_SHA={"replace it with the git commit hash of the bumping version commit"}
git tag -a "${VERSION}" ${GITHUB_SHA} -m "release ${VERSION}"  
git push origin ${VERSION}
```

### Step 5. Check the release workflow

after the tag is pushed, the cd workflow on github will be triggered, please monitor the cd workflow until it finishes
