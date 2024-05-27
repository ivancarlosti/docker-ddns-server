package model

import (
	"time"

	"gorm.io/gorm"
)

// UpdateHost updates all fields of a host entry
// and sets a new LastUpdate date.
func (h *Host) UpdateHost(updateHost *Host) (updateRecord bool) {
	updateRecord = false
	if h.Ip != updateHost.Ip || h.Ttl != updateHost.Ttl {
		updateRecord = true
		h.LastUpdate = time.Now()
	}

	h.Ip = updateHost.Ip
	h.Ttl = updateHost.Ttl
	h.UserName = updateHost.UserName
	h.Password = updateHost.Password

	return
}
