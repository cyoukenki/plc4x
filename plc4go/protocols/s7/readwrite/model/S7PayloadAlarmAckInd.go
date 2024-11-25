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

// S7PayloadAlarmAckInd is the corresponding interface of S7PayloadAlarmAckInd
type S7PayloadAlarmAckInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() AlarmMessageAckPushType
}

// S7PayloadAlarmAckIndExactly can be used when we want exactly this type and not a type which fulfills S7PayloadAlarmAckInd.
// This is useful for switch cases.
type S7PayloadAlarmAckIndExactly interface {
	S7PayloadAlarmAckInd
	isS7PayloadAlarmAckInd() bool
}

// _S7PayloadAlarmAckInd is the data-structure of this message
type _S7PayloadAlarmAckInd struct {
	*_S7PayloadUserDataItem
	AlarmMessage AlarmMessageAckPushType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadAlarmAckInd) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadAlarmAckInd) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadAlarmAckInd) GetCpuSubfunction() uint8 {
	return 0x0c
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadAlarmAckInd) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadAlarmAckInd) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadAlarmAckInd) GetAlarmMessage() AlarmMessageAckPushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadAlarmAckInd factory function for _S7PayloadAlarmAckInd
func NewS7PayloadAlarmAckInd(alarmMessage AlarmMessageAckPushType, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadAlarmAckInd {
	_result := &_S7PayloadAlarmAckInd{
		AlarmMessage:           alarmMessage,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadAlarmAckInd(structType any) S7PayloadAlarmAckInd {
	if casted, ok := structType.(S7PayloadAlarmAckInd); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadAlarmAckInd); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadAlarmAckInd) GetTypeName() string {
	return "S7PayloadAlarmAckInd"
}

func (m *_S7PayloadAlarmAckInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadAlarmAckInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadAlarmAckIndParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadAlarmAckInd, error) {
	return S7PayloadAlarmAckIndParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadAlarmAckIndParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadAlarmAckInd, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadAlarmAckInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadAlarmAckInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmMessage")
	}
	_alarmMessage, _alarmMessageErr := AlarmMessageAckPushTypeParseWithBuffer(ctx, readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field of S7PayloadAlarmAckInd")
	}
	alarmMessage := _alarmMessage.(AlarmMessageAckPushType)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmMessage")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadAlarmAckInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadAlarmAckInd")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadAlarmAckInd{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		AlarmMessage:           alarmMessage,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadAlarmAckInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadAlarmAckInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarmAckInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadAlarmAckInd")
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmMessage")
		}
		_alarmMessageErr := writeBuffer.WriteSerializable(ctx, m.GetAlarmMessage())
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmMessage")
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarmAckInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadAlarmAckInd")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadAlarmAckInd) isS7PayloadAlarmAckInd() bool {
	return true
}

func (m *_S7PayloadAlarmAckInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}