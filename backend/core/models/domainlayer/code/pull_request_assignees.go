/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package code

import (
	"github.com/apache/incubator-devlake/core/models/common"
)

type PullRequestAssignee struct {
	PullRequestId string `json:"id" gorm:"primaryKey;type:varchar(255);comment:This key is generated based on details from the original plugin"` // format: <Plugin>:<Entity>:<PK0>:<PK1>
	AssigneeId    int    `gorm:"primaryKey"`
	Name          string `gorm:"type:varchar(255)"`
	UserName      string `gorm:"type:varchar(255)"`

	common.NoPKModel
}

func (PullRequestAssignee) TableName() string {
	return "pull_request_assignees"
}