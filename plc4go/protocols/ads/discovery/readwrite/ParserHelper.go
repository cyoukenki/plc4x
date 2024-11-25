/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package readwrite

import (
	"context"

	"github.com/apache/plc4x/plc4go/protocols/ads/discovery/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type AdsDiscoveryParserHelper struct {
}

func (m AdsDiscoveryParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (any, error) {
	switch typeName {
	case "AdsDiscovery":
		return model.AdsDiscoveryParseWithBuffer(context.Background(), io)
	case "AdsDiscoveryBlock":
		return model.AdsDiscoveryBlockParseWithBuffer(context.Background(), io)
	case "AdsDiscoveryConstants":
		return model.AdsDiscoveryConstantsParseWithBuffer(context.Background(), io)
	case "AmsNetId":
		return model.AmsNetIdParseWithBuffer(context.Background(), io)
	case "AmsString":
		return model.AmsStringParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}