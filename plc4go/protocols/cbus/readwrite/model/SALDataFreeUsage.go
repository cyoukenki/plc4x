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

// SALDataFreeUsage is the corresponding interface of SALDataFreeUsage
type SALDataFreeUsage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
}

// SALDataFreeUsageExactly can be used when we want exactly this type and not a type which fulfills SALDataFreeUsage.
// This is useful for switch cases.
type SALDataFreeUsageExactly interface {
	SALDataFreeUsage
	isSALDataFreeUsage() bool
}

// _SALDataFreeUsage is the data-structure of this message
type _SALDataFreeUsage struct {
	*_SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataFreeUsage) GetApplicationId() ApplicationId {
	return ApplicationId_FREE_USAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataFreeUsage) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataFreeUsage) GetParent() SALData {
	return m._SALData
}

// NewSALDataFreeUsage factory function for _SALDataFreeUsage
func NewSALDataFreeUsage(salData SALData) *_SALDataFreeUsage {
	_result := &_SALDataFreeUsage{
		_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataFreeUsage(structType any) SALDataFreeUsage {
	if casted, ok := structType.(SALDataFreeUsage); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataFreeUsage); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataFreeUsage) GetTypeName() string {
	return "SALDataFreeUsage"
}

func (m *_SALDataFreeUsage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SALDataFreeUsage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataFreeUsageParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataFreeUsage, error) {
	return SALDataFreeUsageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataFreeUsageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataFreeUsage, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SALDataFreeUsage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataFreeUsage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"FREE_USAGE Not yet implemented"})
	}

	if closeErr := readBuffer.CloseContext("SALDataFreeUsage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataFreeUsage")
	}

	// Create a partially initialized instance
	_child := &_SALDataFreeUsage{
		_SALData: &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataFreeUsage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataFreeUsage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataFreeUsage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataFreeUsage")
		}

		if popErr := writeBuffer.PopContext("SALDataFreeUsage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataFreeUsage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataFreeUsage) isSALDataFreeUsage() bool {
	return true
}

func (m *_SALDataFreeUsage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}