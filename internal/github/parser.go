package github

import (
	"encoding/json"
	"time"
)

func ParseActivity(data []byte) ([]string, error) {
	var rawEvents []map[string]interface{}
	err := json.Unmarshal(data, &rawEvents)
	if err != nil {
		return nil, err
	}

	var activities []string

	for _, event := range rawEvents {
		if text := RenderEvent(event); text != "" {
			activities = append(activities, text)
		}
	}

	return activities, nil
}

func timeFormatting(t string) string {
	tf, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return ""
	}

	userFormat := tf.Format("02/01/2006 15:04")
	return userFormat
}

func RenderEvent(data map[string]interface{}) string {
	createdAt, _ := data["created_at"].(string)
	time := timeFormatting(createdAt)

	eventType, _ := data["type"].(string)

	switch eventType {
	case "CommitCommentEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		comment, _ := payload["comment"].(map[string]interface{})
		htmlURL, _ := comment["html_url"].(string)

		if action != "" && htmlURL != "" {
			return action + " a commit comment at" + htmlURL + " at " + time
		}
	case "CreateEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		payload, _ := data["payload"].(map[string]interface{})
		refType, _ := payload["ref_type"].(string)

		if repoName != "" && refType != "" {
			return "created a " + refType + " in " + repoName + " at " + time
		}

	case "DeleteEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		payload, _ := data["payload"].(map[string]interface{})
		refType, _ := payload["ref_type"].(string)

		if repoName != "" && refType != "" {
			return "deleted a " + refType + " in " + repoName + " at " + time
		}

	case "DiscussionEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" && action != "" {
			return action + " a discusison in  " + repoName + " at " + time
		}

	case "ForkEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" && action != "" {
			return action + " a repository in  " + repoName + " at " + time
		}

	case "GollumEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" {
			return "created a wiki page for  " + repoName + " at " + time
		}

	case "IssueCommentEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		issue, _ := payload["issue"].(map[string]interface{})
		issueURL, _ := issue["url"].(string)

		if issueURL != "" && action != "" {
			return action + " an issue comment at  " + issueURL + " at " + time
		}

	case "MemberEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		member, _ := payload["member"].(map[string]interface{})
		memberName, _ := member["name"].(string)

		if repoName != "" && action != "" && memberName != "" {
			return action + " a  " + memberName + " as collaborator in " + repoName + " at " + time
		}

	case "PublicEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" {
			return " make public " + repoName + " at " + time
		}

	case "PullRequestEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		pullRequest, _ := payload["pull_request"].(map[string]interface{})
		pullRequestURL, _ := pullRequest["url"].(string)

		if action != "" && pullRequestURL != "" {
			return action + " a pull request at " + pullRequestURL + " at " + time
		}

	case "PullRequestReviewEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		pullRequest, _ := payload["pull_request"].(map[string]interface{})
		pullRequestURL, _ := pullRequest["url"].(string)

		if action != "" && pullRequestURL != "" {
			return action + " a pull request review at " + pullRequestURL + " at " + time
		}

	case "PullRequestReviewCommentEvent":
		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		pullRequest, _ := payload["pull_request"].(map[string]interface{})
		pullRequestURL, _ := pullRequest["url"].(string)

		if action != "" && pullRequestURL != "" {
			return action + " a pull request review comment at " + pullRequestURL + " at " + time
		}

	case "PushEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" {
			return "pushed a commit to " + repoName + " at " + time
		}

	case "ReleaseEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		payload, _ := data["payload"].(map[string]interface{})
		action, _ := payload["action"].(string)

		if repoName != "" && action != "" {
			return action + " a release at " + repoName + " at " + time
		}

	case "WatchEvent":
		repo, _ := data["repo"].(map[string]interface{})
		repoName, _ := repo["name"].(string)

		if repoName != "" {
			return "starred a " + repoName + " at " + time
		}
	}

	return ""
}
