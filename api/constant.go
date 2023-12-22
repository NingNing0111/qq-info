package api

type Api string

const (
	GetGroupList       Api = "get_group_list"
	GetGroupMemberList Api = "get_group_member_list"
	GetFriendList      Api = "get_friend_list"
)

type UserInfo struct {
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	QQ       int64  `json:"user_id"`
}

type UserInfoGroup struct {
	Card     string `json:"card"`
	Nickname string `json:"nickname"`
	QQ       int64  `json:"user_id"`
	Sex      string `json:"sex"`
	Role     string `json:"role"`
}

type GroupInfo struct {
	QQ             int64  `json:"group_id"`
	Name           string `json:"group_name"`
	MaxMemberCount int    `json:"max_member_count"`
	MemberCount    int    `json:"member_count"`
}

type GroupListResponse struct {
	Status string      `json:"status"`
	Data   []GroupInfo `json:"data"`
}

type GroupMembersResponse struct {
	Status string          `json:"status"`
	Data   []UserInfoGroup `json:"data"`
}

type FriendListResponse struct {
	Status string     `json:"status"`
	Data   []UserInfo `json:"data"`
}
