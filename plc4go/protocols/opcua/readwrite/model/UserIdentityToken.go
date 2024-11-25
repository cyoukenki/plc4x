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

// UserIdentityToken is the corresponding interface of UserIdentityToken
type UserIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetUserIdentityTokenDefinition returns UserIdentityTokenDefinition (property field)
	GetUserIdentityTokenDefinition() UserIdentityTokenDefinition
}

// UserIdentityTokenExactly can be used when we want exactly this type and not a type which fulfills UserIdentityToken.
// This is useful for switch cases.
type UserIdentityTokenExactly interface {
	UserIdentityToken
	isUserIdentityToken() bool
}

// _UserIdentityToken is the data-structure of this message
type _UserIdentityToken struct {
	*_ExtensionObjectDefinition
	PolicyId                    PascalString
	UserIdentityTokenDefinition UserIdentityTokenDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserIdentityToken) GetIdentifier() string {
	return "316"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserIdentityToken) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_UserIdentityToken) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserIdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

func (m *_UserIdentityToken) GetUserIdentityTokenDefinition() UserIdentityTokenDefinition {
	return m.UserIdentityTokenDefinition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUserIdentityToken factory function for _UserIdentityToken
func NewUserIdentityToken(policyId PascalString, userIdentityTokenDefinition UserIdentityTokenDefinition) *_UserIdentityToken {
	_result := &_UserIdentityToken{
		PolicyId:                    policyId,
		UserIdentityTokenDefinition: userIdentityTokenDefinition,
		_ExtensionObjectDefinition:  NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUserIdentityToken(structType any) UserIdentityToken {
	if casted, ok := structType.(UserIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*UserIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_UserIdentityToken) GetTypeName() string {
	return "UserIdentityToken"
}

func (m *_UserIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (policyLength)
	lengthInBits += 32

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	// Simple field (userIdentityTokenDefinition)
	lengthInBits += m.UserIdentityTokenDefinition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UserIdentityTokenParse(ctx context.Context, theBytes []byte, identifier string) (UserIdentityToken, error) {
	return UserIdentityTokenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func UserIdentityTokenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (UserIdentityToken, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("UserIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (policyLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	policyLength, _policyLengthErr := readBuffer.ReadInt32("policyLength", 32)
	_ = policyLength
	if _policyLengthErr != nil {
		return nil, errors.Wrap(_policyLengthErr, "Error parsing 'policyLength' field of UserIdentityToken")
	}

	// Simple Field (policyId)
	if pullErr := readBuffer.PullContext("policyId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for policyId")
	}
	_policyId, _policyIdErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _policyIdErr != nil {
		return nil, errors.Wrap(_policyIdErr, "Error parsing 'policyId' field of UserIdentityToken")
	}
	policyId := _policyId.(PascalString)
	if closeErr := readBuffer.CloseContext("policyId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for policyId")
	}

	// Simple Field (userIdentityTokenDefinition)
	if pullErr := readBuffer.PullContext("userIdentityTokenDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userIdentityTokenDefinition")
	}
	_userIdentityTokenDefinition, _userIdentityTokenDefinitionErr := UserIdentityTokenDefinitionParseWithBuffer(ctx, readBuffer, string(policyId.GetStringValue()))
	if _userIdentityTokenDefinitionErr != nil {
		return nil, errors.Wrap(_userIdentityTokenDefinitionErr, "Error parsing 'userIdentityTokenDefinition' field of UserIdentityToken")
	}
	userIdentityTokenDefinition := _userIdentityTokenDefinition.(UserIdentityTokenDefinition)
	if closeErr := readBuffer.CloseContext("userIdentityTokenDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userIdentityTokenDefinition")
	}

	if closeErr := readBuffer.CloseContext("UserIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserIdentityToken")
	}

	// Create a partially initialized instance
	_child := &_UserIdentityToken{
		_ExtensionObjectDefinition:  &_ExtensionObjectDefinition{},
		PolicyId:                    policyId,
		UserIdentityTokenDefinition: userIdentityTokenDefinition,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_UserIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserIdentityToken")
		}

		// Implicit Field (policyLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		policyLength := int32(int32(m.GetPolicyId().GetLengthInBytes(ctx)) + int32(m.GetUserIdentityTokenDefinition().GetLengthInBytes(ctx)))
		_policyLengthErr := writeBuffer.WriteInt32("policyLength", 32, int32((policyLength)))
		if _policyLengthErr != nil {
			return errors.Wrap(_policyLengthErr, "Error serializing 'policyLength' field")
		}

		// Simple Field (policyId)
		if pushErr := writeBuffer.PushContext("policyId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for policyId")
		}
		_policyIdErr := writeBuffer.WriteSerializable(ctx, m.GetPolicyId())
		if popErr := writeBuffer.PopContext("policyId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for policyId")
		}
		if _policyIdErr != nil {
			return errors.Wrap(_policyIdErr, "Error serializing 'policyId' field")
		}

		// Simple Field (userIdentityTokenDefinition)
		if pushErr := writeBuffer.PushContext("userIdentityTokenDefinition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userIdentityTokenDefinition")
		}
		_userIdentityTokenDefinitionErr := writeBuffer.WriteSerializable(ctx, m.GetUserIdentityTokenDefinition())
		if popErr := writeBuffer.PopContext("userIdentityTokenDefinition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userIdentityTokenDefinition")
		}
		if _userIdentityTokenDefinitionErr != nil {
			return errors.Wrap(_userIdentityTokenDefinitionErr, "Error serializing 'userIdentityTokenDefinition' field")
		}

		if popErr := writeBuffer.PopContext("UserIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserIdentityToken")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserIdentityToken) isUserIdentityToken() bool {
	return true
}

func (m *_UserIdentityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
