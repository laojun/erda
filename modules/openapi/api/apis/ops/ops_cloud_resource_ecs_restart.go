// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ops

import (
	"strings"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var OPS_CLOUD_RESOURCE_ECS_RESTART = apis.ApiSpec{
	Path:         "/api/cloud-ecs/actions/restart",
	BackendPath:  "/api/cloud-ecs/actions/restart",
	Host:         "ops.marathon.l4lb.thisdcos.directory:9027",
	Scheme:       "http",
	Method:       "POST",
	CheckLogin:   true,
	RequestType:  apistructs.HandleCloudResourceEcsRequest{},
	ResponseType: apistructs.HandleCloudResourceECSResponse{},
	Doc:          "重启ECS实例",
	Audit: func(ctx *spec.AuditContext) error {
		var request apistructs.HandleCloudResourceEcsRequest
		if err := ctx.BindRequestData(&request); err != nil {
			return err
		}

		if request.Vendor == "" || request.Vendor == "aliyun" {
			request.Vendor = "alicloud"
		}

		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.OrgScope,
			ScopeID:      uint64(ctx.OrgID),
			TemplateName: apistructs.EcsRestartTemplate,
			Context: map[string]interface{}{
				"vendor":      request.Vendor,
				"region":      request.Region,
				"instanceIds": strings.Join(request.InstanceIds, ","),
			},
		})
	},
}
