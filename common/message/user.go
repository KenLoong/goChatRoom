package message

type User struct{
	//为了序列化和反序列化成功
	//我们需要保证用户信息的json字符串的key 和 结构体的字段对应的 tag 名字一致
	UserId int			`json:"userId"`
	UserPwd string  	`json:"userPwd"`
	UserName string  	`json:"userName"`
	UserStatus int 		`json:"userStatus"`
}