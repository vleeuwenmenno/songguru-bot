package notify

type NotifyID string

const (
	Legacy              NotifyID = "legacy"
	Welcome             NotifyID = "welcome"
	BotRoleCreateFailed NotifyID = "bot_role_create_failed"
)
