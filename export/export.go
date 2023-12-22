package export

import (
	"log"
	"strconv"

	"github.com/ningning0111/api"
	"github.com/ningning0111/config"
	"github.com/ningning0111/util"
)

var conf = config.GetAppConf()

func Run() {
	exportFriendData()
	exportGroupData()
	exportGroupUserData()
}

func exportGroupData() {
	data := api.Get_group_list()
	w := util.OpenCsvFileWriter(conf.GroupFileName)
	cnt := 0
	w.Write([]string{"序号", "群昵称", "群号", "群人数", "最大群人数"})
	for _, groupInfo := range data {
		cnt++
		line := []string{
			strconv.FormatInt(int64(cnt), 10),
			groupInfo.Name,
			strconv.FormatInt(int64(groupInfo.QQ), 10),
			strconv.FormatInt(int64(groupInfo.MemberCount), 10),
			strconv.FormatInt(int64(groupInfo.MaxMemberCount), 10),
		}
		err := w.Write(line)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func exportFriendData() {
	data := api.Get_friend_list()
	w := util.OpenCsvFileWriter(conf.FriendFileName)
	cnt := 0
	w.Write([]string{"序号", "好友昵称", "好友备注", "QQ号"})
	for _, userInfo := range data {
		cnt++
		line := []string{
			strconv.FormatInt(int64(cnt), 10),
			userInfo.Nickname,
			userInfo.Remark,
			strconv.FormatInt(userInfo.QQ, 10),
		}
		err := w.Write(line)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func exportGroupUserData() {
	groupData := api.Get_group_list()
	w := util.OpenCsvFileWriter(conf.GroupUserFileName)
	cnt := 0
	w.Write([]string{"序号", "QQ号", "群员昵称", "性别", "群权限", "群内备注"})
	for _, group := range groupData {
		memberList := api.Get_group_member_info(group.QQ)
		for _, member := range memberList {
			cnt++
			if member.Card == "" {
				member.Card = "无"
			}
			if member.Sex == "unknown" {
				member.Sex = "未知"
			}
			line := []string{
				strconv.FormatInt(int64(cnt), 10),
				strconv.FormatInt(member.QQ, 10),
				member.Nickname,
				member.Sex,
				member.Role,
				member.Card,
			}
			err := w.Write(line)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}
	}
}
