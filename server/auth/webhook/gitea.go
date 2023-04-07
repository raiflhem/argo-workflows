package webhook

import (
	"net/http"

	"github.com/ergoapi/webhooks/gitea"
)

func giteaMatch(secret string, r *http.Request) bool {
	hook, err := gitea.New(gitea.Options.Secret(secret))
	if err != nil {
		return false
	}
	_, err = hook.Parse(r,
		gitea.HookEventCreate,
    gitea.HookEventDelete,
    gitea.HookEventFork,
    gitea.HookEventPush,
    gitea.HookEventIssues,
    gitea.HookEventIssueAssign,
    gitea.HookEventIssueLabel,
    gitea.HookEventIssueMilestone,
    gitea.HookEventIssueComment,
    gitea.HookEventPullRequest,
    gitea.HookEventPullRequestAssign,
    gitea.HookEventPullRequestLabel,
    gitea.HookEventPullRequestMilestone,
    gitea.HookEventPullRequestComment,
    gitea.HookEventPullRequestReviewApproved,
    gitea.HookEventPullRequestReviewRejected,
    gitea.HookEventPullRequestReviewComment,
    gitea.HookEventPullRequestSync,
    gitea.HookEventWiki,
    gitea.HookEventRepository,
    gitea.HookEventRelease,
    gitea.HookEventPackage,
	)
	return err == nil
}
