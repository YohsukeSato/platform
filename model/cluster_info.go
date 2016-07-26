// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ClusterInfo struct {
	Id                 string `json:"id"`
	InternodeUrl       string `json:"internode_url"`
	Hostname           string `json:"hostname"`
	LastSuccessfulPing int64  `json:"last_ping"`
	IsAlive            bool   `json:"is_alive"`
}

func (me *ClusterInfo) ToJson() string {
	b, err := json.Marshal(me)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ClusterInfoFromJson(data io.Reader) *ClusterInfo {
	decoder := json.NewDecoder(data)
	var me ClusterInfo
	err := decoder.Decode(&me)
	if err == nil {
		return &me
	} else {
		return nil
	}
}

func (me *ClusterInfo) HaveEstablishedInitialContact() bool {
	if me.Id != "" {
		return true
	}

	return false
}
