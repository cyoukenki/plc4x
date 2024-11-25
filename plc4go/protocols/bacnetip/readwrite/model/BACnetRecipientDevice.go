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

// BACnetRecipientDevice is the corresponding interface of BACnetRecipientDevice
type BACnetRecipientDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetRecipient
	// GetDeviceValue returns DeviceValue (property field)
	GetDeviceValue() BACnetContextTagObjectIdentifier
}

// BACnetRecipientDeviceExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientDevice.
// This is useful for switch cases.
type BACnetRecipientDeviceExactly interface {
	BACnetRecipientDevice
	isBACnetRecipientDevice() bool
}

// _BACnetRecipientDevice is the data-structure of this message
type _BACnetRecipientDevice struct {
	*_BACnetRecipient
	DeviceValue BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetRecipientDevice) InitializeParent(parent BACnetRecipient, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetRecipientDevice) GetParent() BACnetRecipient {
	return m._BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientDevice) GetDeviceValue() BACnetContextTagObjectIdentifier {
	return m.DeviceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientDevice factory function for _BACnetRecipientDevice
func NewBACnetRecipientDevice(deviceValue BACnetContextTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetRecipientDevice {
	_result := &_BACnetRecipientDevice{
		DeviceValue:      deviceValue,
		_BACnetRecipient: NewBACnetRecipient(peekedTagHeader),
	}
	_result._BACnetRecipient._BACnetRecipientChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientDevice(structType any) BACnetRecipientDevice {
	if casted, ok := structType.(BACnetRecipientDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientDevice) GetTypeName() string {
	return "BACnetRecipientDevice"
}

func (m *_BACnetRecipientDevice) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceValue)
	lengthInBits += m.DeviceValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientDeviceParse(ctx context.Context, theBytes []byte) (BACnetRecipientDevice, error) {
	return BACnetRecipientDeviceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientDeviceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientDevice, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetRecipientDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceValue)
	if pullErr := readBuffer.PullContext("deviceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceValue")
	}
	_deviceValue, _deviceValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _deviceValueErr != nil {
		return nil, errors.Wrap(_deviceValueErr, "Error parsing 'deviceValue' field of BACnetRecipientDevice")
	}
	deviceValue := _deviceValue.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("deviceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientDevice")
	}

	// Create a partially initialized instance
	_child := &_BACnetRecipientDevice{
		_BACnetRecipient: &_BACnetRecipient{},
		DeviceValue:      deviceValue,
	}
	_child._BACnetRecipient._BACnetRecipientChildRequirements = _child
	return _child, nil
}

func (m *_BACnetRecipientDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetRecipientDevice")
		}

		// Simple Field (deviceValue)
		if pushErr := writeBuffer.PushContext("deviceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceValue")
		}
		_deviceValueErr := writeBuffer.WriteSerializable(ctx, m.GetDeviceValue())
		if popErr := writeBuffer.PopContext("deviceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceValue")
		}
		if _deviceValueErr != nil {
			return errors.Wrap(_deviceValueErr, "Error serializing 'deviceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetRecipientDevice")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetRecipientDevice) isBACnetRecipientDevice() bool {
	return true
}

func (m *_BACnetRecipientDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}