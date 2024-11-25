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

// BrowsePathResult is the corresponding interface of BrowsePathResult
type BrowsePathResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfTargets returns NoOfTargets (property field)
	GetNoOfTargets() int32
	// GetTargets returns Targets (property field)
	GetTargets() []ExtensionObjectDefinition
}

// BrowsePathResultExactly can be used when we want exactly this type and not a type which fulfills BrowsePathResult.
// This is useful for switch cases.
type BrowsePathResultExactly interface {
	BrowsePathResult
	isBrowsePathResult() bool
}

// _BrowsePathResult is the data-structure of this message
type _BrowsePathResult struct {
	*_ExtensionObjectDefinition
	StatusCode  StatusCode
	NoOfTargets int32
	Targets     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePathResult) GetIdentifier() string {
	return "551"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePathResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowsePathResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePathResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_BrowsePathResult) GetNoOfTargets() int32 {
	return m.NoOfTargets
}

func (m *_BrowsePathResult) GetTargets() []ExtensionObjectDefinition {
	return m.Targets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowsePathResult factory function for _BrowsePathResult
func NewBrowsePathResult(statusCode StatusCode, noOfTargets int32, targets []ExtensionObjectDefinition) *_BrowsePathResult {
	_result := &_BrowsePathResult{
		StatusCode:                 statusCode,
		NoOfTargets:                noOfTargets,
		Targets:                    targets,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowsePathResult(structType any) BrowsePathResult {
	if casted, ok := structType.(BrowsePathResult); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePathResult); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePathResult) GetTypeName() string {
	return "BrowsePathResult"
}

func (m *_BrowsePathResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfTargets)
	lengthInBits += 32

	// Array field
	if len(m.Targets) > 0 {
		for _curItem, element := range m.Targets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Targets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowsePathResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowsePathResultParse(ctx context.Context, theBytes []byte, identifier string) (BrowsePathResult, error) {
	return BrowsePathResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowsePathResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowsePathResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BrowsePathResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePathResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of BrowsePathResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (noOfTargets)
	_noOfTargets, _noOfTargetsErr := readBuffer.ReadInt32("noOfTargets", 32)
	if _noOfTargetsErr != nil {
		return nil, errors.Wrap(_noOfTargetsErr, "Error parsing 'noOfTargets' field of BrowsePathResult")
	}
	noOfTargets := _noOfTargets

	// Array field (targets)
	if pullErr := readBuffer.PullContext("targets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targets")
	}
	// Count array
	targets := make([]ExtensionObjectDefinition, utils.Max(noOfTargets, 0))
	// This happens when the size is set conditional to 0
	if len(targets) == 0 {
		targets = nil
	}
	{
		_numItems := uint16(utils.Max(noOfTargets, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "548")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'targets' field of BrowsePathResult")
			}
			targets[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("targets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targets")
	}

	if closeErr := readBuffer.CloseContext("BrowsePathResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePathResult")
	}

	// Create a partially initialized instance
	_child := &_BrowsePathResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		NoOfTargets:                noOfTargets,
		Targets:                    targets,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowsePathResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePathResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePathResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePathResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (noOfTargets)
		noOfTargets := int32(m.GetNoOfTargets())
		_noOfTargetsErr := writeBuffer.WriteInt32("noOfTargets", 32, int32((noOfTargets)))
		if _noOfTargetsErr != nil {
			return errors.Wrap(_noOfTargetsErr, "Error serializing 'noOfTargets' field")
		}

		// Array Field (targets)
		if pushErr := writeBuffer.PushContext("targets", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targets")
		}
		for _curItem, _element := range m.GetTargets() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetTargets()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'targets' field")
			}
		}
		if popErr := writeBuffer.PopContext("targets", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targets")
		}

		if popErr := writeBuffer.PopContext("BrowsePathResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePathResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePathResult) isBrowsePathResult() bool {
	return true
}

func (m *_BrowsePathResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}