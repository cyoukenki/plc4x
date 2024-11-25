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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// CreateMonitoredItemsRequest is the corresponding interface of CreateMonitoredItemsRequest
type CreateMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetNoOfItemsToCreate returns NoOfItemsToCreate (property field)
	GetNoOfItemsToCreate() int32
	// GetItemsToCreate returns ItemsToCreate (property field)
	GetItemsToCreate() []ExtensionObjectDefinition
}

// CreateMonitoredItemsRequestExactly can be used when we want exactly this type and not a type which fulfills CreateMonitoredItemsRequest.
// This is useful for switch cases.
type CreateMonitoredItemsRequestExactly interface {
	CreateMonitoredItemsRequest
	isCreateMonitoredItemsRequest() bool
}

// _CreateMonitoredItemsRequest is the data-structure of this message
type _CreateMonitoredItemsRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader      ExtensionObjectDefinition
	SubscriptionId     uint32
	TimestampsToReturn TimestampsToReturn
	NoOfItemsToCreate  int32
	ItemsToCreate      []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetIdentifier() string {
	return "751"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateMonitoredItemsRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CreateMonitoredItemsRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_CreateMonitoredItemsRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_CreateMonitoredItemsRequest) GetNoOfItemsToCreate() int32 {
	return m.NoOfItemsToCreate
}

func (m *_CreateMonitoredItemsRequest) GetItemsToCreate() []ExtensionObjectDefinition {
	return m.ItemsToCreate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateMonitoredItemsRequest factory function for _CreateMonitoredItemsRequest
func NewCreateMonitoredItemsRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, timestampsToReturn TimestampsToReturn, noOfItemsToCreate int32, itemsToCreate []ExtensionObjectDefinition) *_CreateMonitoredItemsRequest {
	_result := &_CreateMonitoredItemsRequest{
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		TimestampsToReturn:         timestampsToReturn,
		NoOfItemsToCreate:          noOfItemsToCreate,
		ItemsToCreate:              itemsToCreate,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateMonitoredItemsRequest(structType any) CreateMonitoredItemsRequest {
	if casted, ok := structType.(CreateMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateMonitoredItemsRequest) GetTypeName() string {
	return "CreateMonitoredItemsRequest"
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Simple field (noOfItemsToCreate)
	lengthInBits += 32

	// Array field
	if len(m.ItemsToCreate) > 0 {
		for _curItem, element := range m.ItemsToCreate {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ItemsToCreate), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateMonitoredItemsRequestParse(ctx context.Context, theBytes []byte, identifier string) (CreateMonitoredItemsRequest, error) {
	return CreateMonitoredItemsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CreateMonitoredItemsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CreateMonitoredItemsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of CreateMonitoredItemsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (subscriptionId)
	_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of CreateMonitoredItemsRequest")
	}
	subscriptionId := _subscriptionId

	// Simple Field (timestampsToReturn)
	if pullErr := readBuffer.PullContext("timestampsToReturn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestampsToReturn")
	}
	_timestampsToReturn, _timestampsToReturnErr := TimestampsToReturnParseWithBuffer(ctx, readBuffer)
	if _timestampsToReturnErr != nil {
		return nil, errors.Wrap(_timestampsToReturnErr, "Error parsing 'timestampsToReturn' field of CreateMonitoredItemsRequest")
	}
	timestampsToReturn := _timestampsToReturn
	if closeErr := readBuffer.CloseContext("timestampsToReturn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestampsToReturn")
	}

	// Simple Field (noOfItemsToCreate)
	_noOfItemsToCreate, _noOfItemsToCreateErr := readBuffer.ReadInt32("noOfItemsToCreate", 32)
	if _noOfItemsToCreateErr != nil {
		return nil, errors.Wrap(_noOfItemsToCreateErr, "Error parsing 'noOfItemsToCreate' field of CreateMonitoredItemsRequest")
	}
	noOfItemsToCreate := _noOfItemsToCreate

	// Array field (itemsToCreate)
	if pullErr := readBuffer.PullContext("itemsToCreate", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for itemsToCreate")
	}
	// Count array
	itemsToCreate := make([]ExtensionObjectDefinition, utils.Max(noOfItemsToCreate, 0))
	// This happens when the size is set conditional to 0
	if len(itemsToCreate) == 0 {
		itemsToCreate = nil
	}
	{
		_numItems := uint16(utils.Max(noOfItemsToCreate, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "745")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'itemsToCreate' field of CreateMonitoredItemsRequest")
			}
			itemsToCreate[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("itemsToCreate", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for itemsToCreate")
	}

	if closeErr := readBuffer.CloseContext("CreateMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateMonitoredItemsRequest")
	}

	// Create a partially initialized instance
	_child := &_CreateMonitoredItemsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		TimestampsToReturn:         timestampsToReturn,
		NoOfItemsToCreate:          noOfItemsToCreate,
		ItemsToCreate:              itemsToCreate,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CreateMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateMonitoredItemsRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (subscriptionId)
		subscriptionId := uint32(m.GetSubscriptionId())
		_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, uint32((subscriptionId)))
		if _subscriptionIdErr != nil {
			return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
		}

		// Simple Field (timestampsToReturn)
		if pushErr := writeBuffer.PushContext("timestampsToReturn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestampsToReturn")
		}
		_timestampsToReturnErr := writeBuffer.WriteSerializable(ctx, m.GetTimestampsToReturn())
		if popErr := writeBuffer.PopContext("timestampsToReturn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestampsToReturn")
		}
		if _timestampsToReturnErr != nil {
			return errors.Wrap(_timestampsToReturnErr, "Error serializing 'timestampsToReturn' field")
		}

		// Simple Field (noOfItemsToCreate)
		noOfItemsToCreate := int32(m.GetNoOfItemsToCreate())
		_noOfItemsToCreateErr := writeBuffer.WriteInt32("noOfItemsToCreate", 32, int32((noOfItemsToCreate)))
		if _noOfItemsToCreateErr != nil {
			return errors.Wrap(_noOfItemsToCreateErr, "Error serializing 'noOfItemsToCreate' field")
		}

		// Array Field (itemsToCreate)
		if pushErr := writeBuffer.PushContext("itemsToCreate", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for itemsToCreate")
		}
		for _curItem, _element := range m.GetItemsToCreate() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItemsToCreate()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'itemsToCreate' field")
			}
		}
		if popErr := writeBuffer.PopContext("itemsToCreate", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for itemsToCreate")
		}

		if popErr := writeBuffer.PopContext("CreateMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateMonitoredItemsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateMonitoredItemsRequest) isCreateMonitoredItemsRequest() bool {
	return true
}

func (m *_CreateMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}