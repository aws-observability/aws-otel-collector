# Introduction

Dependabot usually create several PRs at the same time. Merging one by one might
be expensive, because for each merge we trigger our CI workflow.

Based on this, we are proposing the usage of a tool to merge all open dependabot
prs for golang dependencies updates into a single PR.

Requirements:
* Have the gh cli installed in your workstation.
* Start with a clean repository in the main branch.

The workflow is the following:

* Close all the dependabot PRs that should not be merged.
* Execute the script ./tools/workflow/dependabot-pr/dependabot-pr.sh
  * This will create a PR with the following title: `dependabot updates <date>`
  * This pr can be reviewed and merged as if it was submitted by a team member.


It is possible that some dependencies that should not be downgraded are downgraded or that the build fails. In this case you can fix the go.mod files manually and
run `make gomod-tidy` before running the workflow again.
