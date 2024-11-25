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

package tests

import (
	"github.com/apache/plc4x/plc4go/spi/options"
	"testing"

	eipIO "github.com/apache/plc4x/plc4go/protocols/eip/readwrite"
	"github.com/apache/plc4x/plc4go/spi/testutils"
)

func TestEipParserSerializerBigEndian(t *testing.T) {
	testutils.RunParserSerializerTestsuite(
		t,
		"assets/testing/protocols/eip/ParserSerializerTestsuiteBigEndian.xml",
		eipIO.EipParserHelper{},
		options.WithCustomLogger(testutils.ProduceTestingLogger(t)),
	)
}

func TestEipParserSerializerLittleEndian(t *testing.T) {
	testutils.RunParserSerializerTestsuite(
		t,
		"assets/testing/protocols/eip/ParserSerializerTestsuiteLittleEndian.xml",
		eipIO.EipParserHelper{},
		options.WithCustomLogger(testutils.ProduceTestingLogger(t)),
	)
}