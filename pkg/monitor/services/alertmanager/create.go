/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package alertmanager

import (
	"bytes"
	"context"

	"github.com/pkg/errors"
	alertconfig "github.com/prometheus/alertmanager/config"
	"tkestack.io/tke/pkg/util/log"
)

func (h *processor) Create(ctx context.Context, clusterName string, alertValue string, entity *alertconfig.Route) error {
	h.Lock()
	defer h.Unlock()

	if clusterName == "" {
		return errors.New("empty clusterName")
	}

	if alertValue == "" {
		return errors.New("empty alertValue")
	}

	routeOp, err := h.loadConfig(ctx, clusterName)
	if err != nil {
		return errors.Wrapf(err, "route operator not found")
	}

	log.Infof("Start to add route %s", alertValue)

	_, err = routeOp.InsertRoute(entity)
	if err != nil {
		return errors.Wrapf(err, "failed to insert route %s", alertValue)
	}

	output := bytes.NewBufferString("")
	err = routeOp.Save(output)
	if err != nil {
		return errors.Wrapf(err, "failed to save")
	}

	err = h.saveConfig(ctx, clusterName, output.String())
	if err != nil {
		return errors.Wrapf(err, "failed to save configmap")
	}

	return nil
}
