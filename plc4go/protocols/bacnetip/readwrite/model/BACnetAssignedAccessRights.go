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

// BACnetAssignedAccessRights is the corresponding interface of BACnetAssignedAccessRights
type BACnetAssignedAccessRights interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAssignedAccessRights returns AssignedAccessRights (property field)
	GetAssignedAccessRights() BACnetDeviceObjectReferenceEnclosed
	// GetEnable returns Enable (property field)
	GetEnable() BACnetContextTagBoolean
}

// BACnetAssignedAccessRightsExactly can be used when we want exactly this type and not a type which fulfills BACnetAssignedAccessRights.
// This is useful for switch cases.
type BACnetAssignedAccessRightsExactly interface {
	BACnetAssignedAccessRights
	isBACnetAssignedAccessRights() bool
}

// _BACnetAssignedAccessRights is the data-structure of this message
type _BACnetAssignedAccessRights struct {
	AssignedAccessRights BACnetDeviceObjectReferenceEnclosed
	Enable               BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedAccessRights) GetAssignedAccessRights() BACnetDeviceObjectReferenceEnclosed {
	return m.AssignedAccessRights
}

func (m *_BACnetAssignedAccessRights) GetEnable() BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAssignedAccessRights factory function for _BACnetAssignedAccessRights
func NewBACnetAssignedAccessRights(assignedAccessRights BACnetDeviceObjectReferenceEnclosed, enable BACnetContextTagBoolean) *_BACnetAssignedAccessRights {
	return &_BACnetAssignedAccessRights{AssignedAccessRights: assignedAccessRights, Enable: enable}
}

// Deprecated: use the interface for direct cast
func CastBACnetAssignedAccessRights(structType any) BACnetAssignedAccessRights {
	if casted, ok := structType.(BACnetAssignedAccessRights); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedAccessRights); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedAccessRights) GetTypeName() string {
	return "BACnetAssignedAccessRights"
}

func (m *_BACnetAssignedAccessRights) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (assignedAccessRights)
	lengthInBits += m.AssignedAccessRights.GetLengthInBits(ctx)

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedAccessRights) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedAccessRightsParse(ctx context.Context, theBytes []byte) (BACnetAssignedAccessRights, error) {
	return BACnetAssignedAccessRightsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAssignedAccessRightsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedAccessRights, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAssignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedAccessRights")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (assignedAccessRights)
	if pullErr := readBuffer.PullContext("assignedAccessRights"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for assignedAccessRights")
	}
	_assignedAccessRights, _assignedAccessRightsErr := BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _assignedAccessRightsErr != nil {
		return nil, errors.Wrap(_assignedAccessRightsErr, "Error parsing 'assignedAccessRights' field of BACnetAssignedAccessRights")
	}
	assignedAccessRights := _assignedAccessRights.(BACnetDeviceObjectReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("assignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for assignedAccessRights")
	}

	// Simple Field (enable)
	if pullErr := readBuffer.PullContext("enable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enable")
	}
	_enable, _enableErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field of BACnetAssignedAccessRights")
	}
	enable := _enable.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("enable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enable")
	}

	if closeErr := readBuffer.CloseContext("BACnetAssignedAccessRights"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedAccessRights")
	}

	// Create the instance
	return &_BACnetAssignedAccessRights{
		AssignedAccessRights: assignedAccessRights,
		Enable:               enable,
	}, nil
}

func (m *_BACnetAssignedAccessRights) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedAccessRights) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedAccessRights"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedAccessRights")
	}

	// Simple Field (assignedAccessRights)
	if pushErr := writeBuffer.PushContext("assignedAccessRights"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for assignedAccessRights")
	}
	_assignedAccessRightsErr := writeBuffer.WriteSerializable(ctx, m.GetAssignedAccessRights())
	if popErr := writeBuffer.PopContext("assignedAccessRights"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for assignedAccessRights")
	}
	if _assignedAccessRightsErr != nil {
		return errors.Wrap(_assignedAccessRightsErr, "Error serializing 'assignedAccessRights' field")
	}

	// Simple Field (enable)
	if pushErr := writeBuffer.PushContext("enable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for enable")
	}
	_enableErr := writeBuffer.WriteSerializable(ctx, m.GetEnable())
	if popErr := writeBuffer.PopContext("enable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for enable")
	}
	if _enableErr != nil {
		return errors.Wrap(_enableErr, "Error serializing 'enable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedAccessRights"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedAccessRights")
	}
	return nil
}

func (m *_BACnetAssignedAccessRights) isBACnetAssignedAccessRights() bool {
	return true
}

func (m *_BACnetAssignedAccessRights) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}