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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestGetEventInformation is the corresponding interface of BACnetConfirmedServiceRequestGetEventInformation
type BACnetConfirmedServiceRequestGetEventInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetLastReceivedObjectIdentifier returns LastReceivedObjectIdentifier (property field)
	GetLastReceivedObjectIdentifier() BACnetContextTagObjectIdentifier
}

// BACnetConfirmedServiceRequestGetEventInformationExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestGetEventInformation.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestGetEventInformationExactly interface {
	BACnetConfirmedServiceRequestGetEventInformation
	isBACnetConfirmedServiceRequestGetEventInformation() bool
}

// _BACnetConfirmedServiceRequestGetEventInformation is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEventInformation struct {
	*_BACnetConfirmedServiceRequest
	LastReceivedObjectIdentifier BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLastReceivedObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.LastReceivedObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEventInformation factory function for _BACnetConfirmedServiceRequestGetEventInformation
func NewBACnetConfirmedServiceRequestGetEventInformation(lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestGetEventInformation {
	_result := &_BACnetConfirmedServiceRequestGetEventInformation{
		LastReceivedObjectIdentifier:   lastReceivedObjectIdentifier,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEventInformation(structType any) BACnetConfirmedServiceRequestGetEventInformation {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEventInformation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEventInformation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEventInformation"
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (lastReceivedObjectIdentifier)
	if m.LastReceivedObjectIdentifier != nil {
		lengthInBits += m.LastReceivedObjectIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEventInformationParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestGetEventInformation, error) {
	return BACnetConfirmedServiceRequestGetEventInformationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestGetEventInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestGetEventInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEventInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEventInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (lastReceivedObjectIdentifier) (Can be skipped, if a given expression evaluates to false)
	var lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("lastReceivedObjectIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for lastReceivedObjectIdentifier")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'lastReceivedObjectIdentifier' field of BACnetConfirmedServiceRequestGetEventInformation")
		default:
			lastReceivedObjectIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("lastReceivedObjectIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for lastReceivedObjectIdentifier")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEventInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEventInformation")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestGetEventInformation{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		LastReceivedObjectIdentifier: lastReceivedObjectIdentifier,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEventInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEventInformation")
		}

		// Optional Field (lastReceivedObjectIdentifier) (Can be skipped, if the value is null)
		var lastReceivedObjectIdentifier BACnetContextTagObjectIdentifier = nil
		if m.GetLastReceivedObjectIdentifier() != nil {
			if pushErr := writeBuffer.PushContext("lastReceivedObjectIdentifier"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for lastReceivedObjectIdentifier")
			}
			lastReceivedObjectIdentifier = m.GetLastReceivedObjectIdentifier()
			_lastReceivedObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, lastReceivedObjectIdentifier)
			if popErr := writeBuffer.PopContext("lastReceivedObjectIdentifier"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for lastReceivedObjectIdentifier")
			}
			if _lastReceivedObjectIdentifierErr != nil {
				return errors.Wrap(_lastReceivedObjectIdentifierErr, "Error serializing 'lastReceivedObjectIdentifier' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEventInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEventInformation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) isBACnetConfirmedServiceRequestGetEventInformation() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}