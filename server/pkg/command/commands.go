package command

type Command map[string]string

var (
	Commands = Command{
		"CMD_SEND_ROOMS":                       "->CMD_SEND_ROOMS<-",
		"CMD_SEND_ROOM_LIST_UPDATE_TO_CLIENTS": "->CMD_SEND_ROOM_LIST_UPDATE_TO_CLIENTS<-",
	}
)
