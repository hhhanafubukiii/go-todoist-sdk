package gotodoistsdk

const (
	BASE_URL      = "https://api.todoist.com"
	AUTH_BASE_URL = "https://todoist.com"
	SYNC_VERSION  = "v9"
	REST_VERSION  = "v2"

	TASKS_ENDPOINT         = "tasks"
	PROJECTS_ENDPOINT      = "projects"
	COLLABORATORS_ENDPOINT = "collaborators"
	SECTIONS_ENDPOINT      = "sections"
	COMMENTS_ENDPOINT      = "comments"
	LABELS_ENDPOINT        = "labels"
	SHARED_LABELS_ENDPOINT = "labels/shared"
	QUICK_ADD_ENDPOINT     = "quick/add"

	AUTHORIZE_ENDPOINT    = "oauth/authorize"
	TOKEN_ENDPOINT        = "oauth/access_token"
	REVOKE_TOKEN_ENDPOINT = "access_tokens/revoke"

	COMPLETED_ITEMS_ENDPOINT = "archive/items"
)
