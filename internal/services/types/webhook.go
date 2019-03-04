// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package types

type WebhookEvent string

const (
	WebhookEventPush        WebhookEvent = "push"
	WebhookEventTag         WebhookEvent = "tag"
	WebhookEventPullRequest WebhookEvent = "pull_request"
)

type RunType string

const (
	RunTypeProject RunType = "project"
	RunTypeUser    RunType = "user"
)

type WebhookData struct {
	Event     WebhookEvent `json:"event,omitempty"`
	ProjectID string       `json:"project_id,omitempty"`

	CompareLink  string `json:"compare_link,omitempty"`   // Pimray link to source. It can be the commit
	CommitLink   string `json:"commit_link,omitempty"`    // Pimray link to source. It can be the commit
	CommitSHA    string `json:"commit_sha,omitempty"`     // commit SHA (SHA1 but also future SHA like SHA256)
	OldCommitSHA string `json:"old_commit_sha,omitempty"` // commit SHA of the head before this push
	Ref          string `json:"ref,omitempty"`            // Ref containing the commit SHA
	Message      string `json:"message,omitempty"`        // Message to use (Push last commit message summary, PR title, Tag message etc...)
	Sender       string `json:"sender,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Branch     string `json:"branch,omitempty"`
	BranchLink string `json:"branch_link,omitempty"`

	Tag     string `json:"tag,omitempty"`
	TagLink string `json:"tag_link,omitempty"`

	// use a string if on some platform (current or future) some PRs id will not be numbers
	PullRequestID   string `json:"pull_request_id,omitempty"`
	PullRequestLink string `json:"link,omitempty"` // Link to pull request

	Repo WebhookDataRepo `json:"repo,omitempty"`
}

type WebhookDataRepo struct {
	Name     string `json:"name,omitempty"`
	Owner    string `json:"owner,omitempty"`
	FullName string `json:"full_name,omitempty"`
	RepoURL  string `json:"repo_url,omitempty"`
}