/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"net/http"

	frtypes "configcenter/src/framework/core/types"
	"configcenter/src/scene_server/topo_server/core/types"
)

func init() {
	apiInst.initFuncs = append(apiInst.initFuncs, apiInst.initObject)
}

func (cli *topoAPI) initObject() {
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/object/batch", HandlerFunc: cli.CreateObjectBatch})
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/object/search/batch", HandlerFunc: cli.SearchObjectBatch})
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/object", HandlerFunc: cli.CreateObject})
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/objects", HandlerFunc: cli.SearchObject})
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/objects/topo", HandlerFunc: cli.SearchObjectTopo})
	cli.actions = append(cli.actions, action{Method: http.MethodPut, Path: "/object/{id}", HandlerFunc: cli.UpdateObject})
	cli.actions = append(cli.actions, action{Method: http.MethodDelete, Path: "/object/{id}", HandlerFunc: cli.DeleteObject})
}

// CreateObjectBatch batch to create some objects
func (cli *topoAPI) CreateObjectBatch(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// SearchObjectBatch batch to search some objects
func (cli *topoAPI) SearchObjectBatch(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// CreateObject create a new object
func (cli *topoAPI) CreateObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// SearchObject search some objects by condition
func (cli *topoAPI) SearchObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// SearchObjectTopo search the object topo
func (cli *topoAPI) SearchObjectTopo(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// UpdateObject update the object
func (cli *topoAPI) UpdateObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}

// DeleteObject delete the object
func (cli *topoAPI) DeleteObject(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (frtypes.MapStr, error) {
	return nil, nil
}
