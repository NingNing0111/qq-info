package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ningning0111/config"
)

var conf = config.GetAppConf()

func Get_friend_list() []UserInfo {
	friendList := post(GetFriendList, nil)
	var rep FriendListResponse
	_ = parseJson(friendList, &rep)
	if rep.Status == "ok" {
		return rep.Data
	}
	return nil
}

func Get_group_member_info(qq int64) []UserInfoGroup {
	formValues := url.Values{}
	formValues.Add("group_id", strconv.FormatInt(int64(qq), 10))
	groupMembers := post(GetGroupMemberList, &formValues)
	var rep GroupMembersResponse
	_ = parseJson(groupMembers, &rep)
	if rep.Status == "ok" {
		return rep.Data
	}
	return nil
}

func Get_group_list() []GroupInfo {
	groupListJson := post(GetGroupList, nil)
	var rep GroupListResponse
	_ = parseJson(groupListJson, &rep)
	if rep.Status == "ok" {
		return rep.Data
	}
	return nil
}

func post(a Api, values *url.Values) string {
	u := buildUrl(a)
	if values == nil {
		values = &url.Values{}
	}

	resp, err := http.PostForm(u, *values)
	if err != nil {
		log.Println("Network error.", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read exception.", err.Error())
	}
	return string(body)
}

func buildUrl(s Api) string {
	return conf.CqUrl + string(s)
}

func parseJson(s string, obj any) error {
	b := []byte(s)
	err := json.Unmarshal(b, &obj)
	if err != nil {
		log.Println("Json string parsing exception.", err.Error())
		return err
	}
	return nil
}
