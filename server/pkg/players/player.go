package players

type playerInfo struct {
	Address      string
	Name         string
	CurrentStore uint8
}

type Players struct {
	PlayerOne playerInfo
	PlayerTwo playerInfo
}
