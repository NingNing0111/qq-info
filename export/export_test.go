package export

import (
	"testing"
)

func Test_exportData(t *testing.T) {
	exportFriendData()
}

func Test_exportGroupData(t *testing.T) {
	exportGroupData()
}

func Test_exportGroupMemberData(t *testing.T) {
	exportGroupUserData()
}

func Test_run(t *testing.T) {
	Run()
}
