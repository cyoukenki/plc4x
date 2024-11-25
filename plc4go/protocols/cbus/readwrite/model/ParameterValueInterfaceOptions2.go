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

// ParameterValueInterfaceOptions2 is the corresponding interface of ParameterValueInterfaceOptions2
type ParameterValueInterfaceOptions2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() InterfaceOptions2
	// GetData returns Data (property field)
	GetData() []byte
}

// ParameterValueInterfaceOptions2Exactly can be used when we want exactly this type and not a type which fulfills ParameterValueInterfaceOptions2.
// This is useful for switch cases.
type ParameterValueInterfaceOptions2Exactly interface {
	ParameterValueInterfaceOptions2
	isParameterValueInterfaceOptions2() bool
}

// _ParameterValueInterfaceOptions2 is the data-structure of this message
type _ParameterValueInterfaceOptions2 struct {
	*_ParameterValue
	Value InterfaceOptions2
	Data  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueInterfaceOptions2) GetParameterType() ParameterType {
	return ParameterType_INTERFACE_OPTIONS_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueInterfaceOptions2) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueInterfaceOptions2) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueInterfaceOptions2) GetValue() InterfaceOptions2 {
	return m.Value
}

func (m *_ParameterValueInterfaceOptions2) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueInterfaceOptions2 factory function for _ParameterValueInterfaceOptions2
func NewParameterValueInterfaceOptions2(value InterfaceOptions2, data []byte, numBytes uint8) *_ParameterValueInterfaceOptions2 {
	_result := &_ParameterValueInterfaceOptions2{
		Value:           value,
		Data:            data,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueInterfaceOptions2(structType any) ParameterValueInterfaceOptions2 {
	if casted, ok := structType.(ParameterValueInterfaceOptions2); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueInterfaceOptions2); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueInterfaceOptions2) GetTypeName() string {
	return "ParameterValueInterfaceOptions2"
}

func (m *_ParameterValueInterfaceOptions2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueInterfaceOptions2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueInterfaceOptions2Parse(ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueInterfaceOptions2, error) {
	return ParameterValueInterfaceOptions2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueInterfaceOptions2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueInterfaceOptions2, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ParameterValueInterfaceOptions2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueInterfaceOptions2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{"InterfaceOptions2 has exactly one byte"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := InterfaceOptions2ParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueInterfaceOptions2")
	}
	value := _value.(InterfaceOptions2)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(numBytes) - uint16(uint16(1)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ParameterValueInterfaceOptions2")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueInterfaceOptions2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueInterfaceOptions2")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueInterfaceOptions2{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
		Data:  data,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueInterfaceOptions2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueInterfaceOptions2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueInterfaceOptions2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueInterfaceOptions2")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueInterfaceOptions2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueInterfaceOptions2")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueInterfaceOptions2) isParameterValueInterfaceOptions2() bool {
	return true
}

func (m *_ParameterValueInterfaceOptions2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
