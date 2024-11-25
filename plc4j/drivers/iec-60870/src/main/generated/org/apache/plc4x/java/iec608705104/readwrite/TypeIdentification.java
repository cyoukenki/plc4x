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
package org.apache.plc4x.java.iec608705104.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum TypeIdentification {
  NOT_USED((short) 0x00, (byte) 0),
  SINGLE_POINT_INFORMATION((short) 0x01, (byte) 0),
  SINGLE_POINT_INFORMATION_WITH_TIME_TAG((short) 0x02, (byte) 3),
  DOUBLE_POINT_INFORMATION((short) 0x03, (byte) 0),
  DOUBLE_POINT_INFORMATION_WITH_TIME_TAG((short) 0x04, (byte) 3),
  STEP_POSITION_INFORMATION((short) 0x05, (byte) 0),
  STEP_POSITION_INFORMATION_WITH_TIME_TAG((short) 0x06, (byte) 3),
  BITSTRING_OF_32_BIT((short) 0x07, (byte) 0),
  BITSTRING_OF_32_BIT_WITH_TIME_TAG((short) 0x08, (byte) 3),
  MEASURED_VALUE_NORMALISED_VALUE((short) 0x09, (byte) 0),
  MEASURED_VALUE_NORMALIZED_VALUE_WITH_TIME_TAG((short) 0x0A, (byte) 3),
  MEASURED_VALUE_SCALED_VALUE((short) 0x0B, (byte) 0),
  MEASURED_VALUE_SCALED_VALUE_WITH_TIME_TAG((short) 0x0C, (byte) 3),
  MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER((short) 0x0D, (byte) 0),
  MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER_WITH_TIME_TAG((short) 0x0E, (byte) 3),
  INTEGRATED_TOTALS((short) 0x0F, (byte) 0),
  INTEGRATED_TOTALS_WITH_TIME_TAG((short) 0x10, (byte) 3),
  EVENT_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG((short) 0x11, (byte) 3),
  PACKED_START_EVENTS_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG((short) 0x12, (byte) 3),
  PACKED_OUTPUT_CIRCUIT_INFORMATION_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG((short) 0x13, (byte) 3),
  PACKED_SINGLE_POINT_INFORMATION_WITH_STATUS_CHANGE_DETECTION((short) 0x14, (byte) 0),
  MEASURED_VALUE_NORMALIZED_VALUE_WITHOUT_QUALITY_DESCRIPTOR((short) 0x15, (byte) 0),
  SINGLE_POINT_INFORMATION_WITH_TIME_TAG_CP56TIME2A((short) 0x1E, (byte) 7),
  DOUBLE_POINT_INFORMATION_WITH_TIME_TAG_CP56TIME2A((short) 0x1F, (byte) 7),
  STEP_POSITION_INFORMATION_WITH_TIME_TAG_CP56TIME2A((short) 0x20, (byte) 7),
  BITSTRING_OF_32_BIT_WITH_TIME_TAG_CP56TIME2A((short) 0x21, (byte) 7),
  MEASURED_VALUE_NORMALISED_VALUE_WITH_TIME_TAG_CP56TIME2A((short) 0x22, (byte) 7),
  MEASURED_VALUE_SCALED_VALUE_WITH_TIME_TAG_CP56TIME2A((short) 0x23, (byte) 7),
  MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER_WITH_TIME_TAG_CP56TIME2A((short) 0x24, (byte) 7),
  INTEGRATED_TOTALS_WITH_TIME_TAG_CP56TIME2A((short) 0x25, (byte) 7),
  EVENT_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A((short) 0x26, (byte) 7),
  PACKED_START_EVENTS_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A((short) 0x27, (byte) 7),
  PACKED_OUTPUT_CIRCUIT_INFORMATION_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A(
      (short) 0x28, (byte) 7),
  SINGLE_COMMAND((short) 0x2D, (byte) 0),
  DOUBLE_COMMAND((short) 0x2E, (byte) 0),
  REGULATING_STEP_COMMAND((short) 0x2F, (byte) 0),
  SET_POINT_COMMAND_NORMALISED_VALUE((short) 0x30, (byte) 0),
  SET_POINT_COMMAND_SCALED_VALUE((short) 0x31, (byte) 0),
  SET_POINT_COMMAND_SHORT_FLOATING_POINT_NUMBER((short) 0x32, (byte) 0),
  BITSTRING_32_BIT_COMMAND((short) 0x33, (byte) 0),
  SINGLE_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x3A, (byte) 7),
  DOUBLE_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x3B, (byte) 7),
  REGULATING_STEP_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x3C, (byte) 7),
  MEASURED_VALUE_NORMALISED_VALUE_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x3D, (byte) 7),
  MEASURED_VALUE_SCALED_VALUE_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x3E, (byte) 7),
  MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER_COMMAND_WITH_TIME_TAG_CP56TIME2A(
      (short) 0x3F, (byte) 7),
  BITSTRING_OF_32_BIT_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x40, (byte) 7),
  END_OF_INITIALISATION((short) 0x46, (byte) 0),
  INTERROGATION_COMMAND((short) 0x64, (byte) 0),
  COUNTER_INTERROGATION_COMMAND((short) 0x65, (byte) 0),
  READ_COMMAND((short) 0x66, (byte) 0),
  CLOCK_SYNCHRONISATION_COMMAND((short) 0x67, (byte) 0),
  TEST_COMMAND((short) 0x68, (byte) 0),
  RESET_PROCESS_COMMAND((short) 0x69, (byte) 0),
  DELAY_ACQUISITION_COMMAND((short) 0x6A, (byte) 0),
  TEST_COMMAND_WITH_TIME_TAG_CP56TIME2A((short) 0x6B, (byte) 0),
  PARAMETER_OF_MEASURED_VALUES_NORMALIZED_VALUE((short) 0x6E, (byte) 0),
  PARAMETER_OF_MEASURED_VALUES_SCALED_VALUE((short) 0x6F, (byte) 0),
  PARAMETER_OF_MEASURED_VALUES_SHORT_FLOATING_POINT_NUMBER((short) 0x70, (byte) 0),
  PARAMETER_ACTIVATION((short) 0x71, (byte) 0),
  FILE_READY((short) 0x78, (byte) 0),
  SECTION_READY((short) 0x79, (byte) 0),
  CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION((short) 0x7A, (byte) 0),
  LAST_SECTION_LAST_SEGMENT((short) 0x7B, (byte) 0),
  ACK_FILE_ACK_SECTION((short) 0x7C, (byte) 0),
  SEGMENT((short) 0x7D, (byte) 0),
  DIRECTORY((short) 0x7E, (byte) 7);
  private static final Map<Short, TypeIdentification> map;

  static {
    map = new HashMap<>();
    for (TypeIdentification value : TypeIdentification.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final byte numTimeBytes;

  TypeIdentification(short value, byte numTimeBytes) {
    this.value = value;
    this.numTimeBytes = numTimeBytes;
  }

  public short getValue() {
    return value;
  }

  public byte getNumTimeBytes() {
    return numTimeBytes;
  }

  public static TypeIdentification firstEnumForFieldNumTimeBytes(byte fieldValue) {
    for (TypeIdentification _val : TypeIdentification.values()) {
      if (_val.getNumTimeBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TypeIdentification> enumsForFieldNumTimeBytes(byte fieldValue) {
    List<TypeIdentification> _values = new ArrayList<>();
    for (TypeIdentification _val : TypeIdentification.values()) {
      if (_val.getNumTimeBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static TypeIdentification enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}