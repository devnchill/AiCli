package types

type MessageRole int

const (
	RoleUSER MessageRole = iota
	RoleLLM
	RoleSYSTEM
)

type Message struct {
	Role    MessageRole
	Content string
}
