package interfaces

import "Catalog/core"

type ISocksRepository interface {
	CreateSock(sock core.Sock) bool
	GetAllSocks() []*core.Sock
	GetSockById(sock int) *core.Sock
	DeleteSock(sock core.Sock) bool
	UpdateSock(sock core.Sock) bool
}
