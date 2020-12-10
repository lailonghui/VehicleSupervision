package xid

import "github.com/dgryski/trifles/uuid"

// 获取xid
func GetXid() string {

	return uuid.UUIDv4()
}
