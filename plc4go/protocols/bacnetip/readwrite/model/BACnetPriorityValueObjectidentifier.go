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

// BACnetPriorityValueObjectidentifier is the corresponding interface of BACnetPriorityValueObjectidentifier
type BACnetPriorityValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetPriorityValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueObjectidentifier.
// This is useful for switch cases.
type BACnetPriorityValueObjectidentifierExactly interface {
	BACnetPriorityValueObjectidentifier
	isBACnetPriorityValueObjectidentifier() bool
}

// _BACnetPriorityValueObjectidentifier is the data-structure of this message
type _BACnetPriorityValueObjectidentifier struct {
	*_BACnetPriorityValue
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueObjectidentifier) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueObjectidentifier) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueObjectidentifier factory function for _BACnetPriorityValueObjectidentifier
func NewBACnetPriorityValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueObjectidentifier {
	_result := &_BACnetPriorityValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		_BACnetPriorityValue:  NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueObjectidentifier(structType any) BACnetPriorityValueObjectidentifier {
	if casted, ok := structType.(BACnetPriorityValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueObjectidentifier) GetTypeName() string {
	return "BACnetPriorityValueObjectidentifier"
}

func (m *_BACnetPriorityValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueObjectidentifierParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValueObjectidentifier, error) {
	return BACnetPriorityValueObjectidentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueObjectidentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPriorityValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectidentifierValue")
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field of BACnetPriorityValueObjectidentifier")
	}
	objectidentifierValue := _objectidentifierValue.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectidentifierValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueObjectidentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueObjectidentifier{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		ObjectidentifierValue: objectidentifierValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueObjectidentifier")
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectidentifierValue")
		}
		_objectidentifierValueErr := writeBuffer.WriteSerializable(ctx, m.GetObjectidentifierValue())
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectidentifierValue")
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueObjectidentifier")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueObjectidentifier) isBACnetPriorityValueObjectidentifier() bool {
	return true
}

func (m *_BACnetPriorityValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}