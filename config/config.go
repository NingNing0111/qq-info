package config

type AppConf struct {
	ExportPath        string
	FriendFileName    string
	GroupFileName     string
	GroupUserFileName string
	CqUrl             string
}

var appConf = AppConf{}

func init() {
	appConf = AppConf{
		ExportPath:        "./csv/", // 导出路径，相对于main.go
		FriendFileName:    "好友列表.csv",
		GroupFileName:     "群列表.csv",
		GroupUserFileName: "群成员列表.csv",
		CqUrl:             "http://127.0.0.1:5700/",
	}
}

func GetAppConf() AppConf {
	return appConf
}
