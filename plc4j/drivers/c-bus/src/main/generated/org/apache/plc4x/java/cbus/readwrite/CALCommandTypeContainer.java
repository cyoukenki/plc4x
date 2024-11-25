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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum CALCommandTypeContainer {
  CALCommandReset((short) 0x08, (byte) 0, CALCommandType.RESET),
  CALCommandRecall((short) 0x1A, (byte) 2, CALCommandType.RECALL),
  CALCommandIdentify((short) 0x21, (byte) 1, CALCommandType.IDENTIFY),
  CALCommandGetStatus((short) 0x2A, (byte) 2, CALCommandType.GET_STATUS),
  CALCommandAcknowledge((short) 0x32, (byte) 2, CALCommandType.ACKNOWLEDGE),
  CALCommandReply_0Bytes((short) 0x80, (byte) 0, CALCommandType.REPLY),
  CALCommandReply_1Bytes((short) 0x81, (byte) 1, CALCommandType.REPLY),
  CALCommandReply_2Bytes((short) 0x82, (byte) 2, CALCommandType.REPLY),
  CALCommandReply_3Bytes((short) 0x83, (byte) 3, CALCommandType.REPLY),
  CALCommandReply_4Bytes((short) 0x84, (byte) 4, CALCommandType.REPLY),
  CALCommandReply_5Bytes((short) 0x85, (byte) 5, CALCommandType.REPLY),
  CALCommandReply_6Bytes((short) 0x86, (byte) 6, CALCommandType.REPLY),
  CALCommandReply_7Bytes((short) 0x87, (byte) 7, CALCommandType.REPLY),
  CALCommandReply_8Bytes((short) 0x88, (byte) 8, CALCommandType.REPLY),
  CALCommandReply_9Bytes((short) 0x89, (byte) 9, CALCommandType.REPLY),
  CALCommandReply_10Bytes((short) 0x8A, (byte) 10, CALCommandType.REPLY),
  CALCommandReply_11Bytes((short) 0x8B, (byte) 11, CALCommandType.REPLY),
  CALCommandReply_12Bytes((short) 0x8C, (byte) 12, CALCommandType.REPLY),
  CALCommandReply_13Bytes((short) 0x8D, (byte) 13, CALCommandType.REPLY),
  CALCommandReply_14Bytes((short) 0x8E, (byte) 14, CALCommandType.REPLY),
  CALCommandReply_15Bytes((short) 0x8F, (byte) 15, CALCommandType.REPLY),
  CALCommandReply_16Bytes((short) 0x90, (byte) 16, CALCommandType.REPLY),
  CALCommandReply_17Bytes((short) 0x91, (byte) 17, CALCommandType.REPLY),
  CALCommandReply_18Bytes((short) 0x92, (byte) 18, CALCommandType.REPLY),
  CALCommandReply_19Bytes((short) 0x93, (byte) 19, CALCommandType.REPLY),
  CALCommandReply_20Bytes((short) 0x94, (byte) 20, CALCommandType.REPLY),
  CALCommandReply_21Bytes((short) 0x95, (byte) 21, CALCommandType.REPLY),
  CALCommandReply_22Bytes((short) 0x96, (byte) 22, CALCommandType.REPLY),
  CALCommandReply_23Bytes((short) 0x97, (byte) 23, CALCommandType.REPLY),
  CALCommandReply_24Bytes((short) 0x98, (byte) 24, CALCommandType.REPLY),
  CALCommandReply_25Bytes((short) 0x99, (byte) 25, CALCommandType.REPLY),
  CALCommandReply_26Bytes((short) 0x9A, (byte) 26, CALCommandType.REPLY),
  CALCommandReply_27Bytes((short) 0x9B, (byte) 27, CALCommandType.REPLY),
  CALCommandReply_28Bytes((short) 0x9C, (byte) 28, CALCommandType.REPLY),
  CALCommandReply_29Bytes((short) 0x9D, (byte) 29, CALCommandType.REPLY),
  CALCommandReply_30Bytes((short) 0x9E, (byte) 30, CALCommandType.REPLY),
  CALCommandReply_31Bytes((short) 0x9F, (byte) 31, CALCommandType.REPLY),
  CALCommandWrite_0Bytes((short) 0xA0, (byte) 0, CALCommandType.WRITE),
  CALCommandWrite_1Bytes((short) 0xA1, (byte) 1, CALCommandType.WRITE),
  CALCommandWrite_2Bytes((short) 0xA2, (byte) 2, CALCommandType.WRITE),
  CALCommandWrite_3Bytes((short) 0xA3, (byte) 3, CALCommandType.WRITE),
  CALCommandWrite_4Bytes((short) 0xA4, (byte) 4, CALCommandType.WRITE),
  CALCommandWrite_5Bytes((short) 0xA5, (byte) 5, CALCommandType.WRITE),
  CALCommandWrite_6Bytes((short) 0xA6, (byte) 6, CALCommandType.WRITE),
  CALCommandWrite_7Bytes((short) 0xA7, (byte) 7, CALCommandType.WRITE),
  CALCommandWrite_8Bytes((short) 0xA8, (byte) 8, CALCommandType.WRITE),
  CALCommandWrite_9Bytes((short) 0xA9, (byte) 9, CALCommandType.WRITE),
  CALCommandWrite_10Bytes((short) 0xAA, (byte) 10, CALCommandType.WRITE),
  CALCommandWrite_11Bytes((short) 0xAB, (byte) 11, CALCommandType.WRITE),
  CALCommandWrite_12Bytes((short) 0xAC, (byte) 12, CALCommandType.WRITE),
  CALCommandWrite_13Bytes((short) 0xAD, (byte) 13, CALCommandType.WRITE),
  CALCommandWrite_14Bytes((short) 0xAE, (byte) 14, CALCommandType.WRITE),
  CALCommandWrite_15Bytes((short) 0xAF, (byte) 15, CALCommandType.WRITE),
  CALCommandStatus_0Bytes((short) 0xC0, (byte) 0, CALCommandType.STATUS),
  CALCommandStatus_1Bytes((short) 0xC1, (byte) 1, CALCommandType.STATUS),
  CALCommandStatus_2Bytes((short) 0xC2, (byte) 2, CALCommandType.STATUS),
  CALCommandStatus_3Bytes((short) 0xC3, (byte) 3, CALCommandType.STATUS),
  CALCommandStatus_4Bytes((short) 0xC4, (byte) 4, CALCommandType.STATUS),
  CALCommandStatus_5Bytes((short) 0xC5, (byte) 5, CALCommandType.STATUS),
  CALCommandStatus_6Bytes((short) 0xC6, (byte) 6, CALCommandType.STATUS),
  CALCommandStatus_7Bytes((short) 0xC7, (byte) 7, CALCommandType.STATUS),
  CALCommandStatus_8Bytes((short) 0xC8, (byte) 8, CALCommandType.STATUS),
  CALCommandStatus_9Bytes((short) 0xC9, (byte) 9, CALCommandType.STATUS),
  CALCommandStatus_10Bytes((short) 0xCA, (byte) 10, CALCommandType.STATUS),
  CALCommandStatus_11Bytes((short) 0xCB, (byte) 11, CALCommandType.STATUS),
  CALCommandStatus_12Bytes((short) 0xCC, (byte) 12, CALCommandType.STATUS),
  CALCommandStatus_13Bytes((short) 0xCD, (byte) 13, CALCommandType.STATUS),
  CALCommandStatus_14Bytes((short) 0xCE, (byte) 14, CALCommandType.STATUS),
  CALCommandStatus_15Bytes((short) 0xCF, (byte) 15, CALCommandType.STATUS),
  CALCommandStatus_16Bytes((short) 0xD0, (byte) 16, CALCommandType.STATUS),
  CALCommandStatus_17Bytes((short) 0xD1, (byte) 17, CALCommandType.STATUS),
  CALCommandStatus_18Bytes((short) 0xD2, (byte) 18, CALCommandType.STATUS),
  CALCommandStatus_19Bytes((short) 0xD3, (byte) 19, CALCommandType.STATUS),
  CALCommandStatus_20Bytes((short) 0xD4, (byte) 20, CALCommandType.STATUS),
  CALCommandStatus_21Bytes((short) 0xD5, (byte) 21, CALCommandType.STATUS),
  CALCommandStatus_22Bytes((short) 0xD6, (byte) 22, CALCommandType.STATUS),
  CALCommandStatus_23Bytes((short) 0xD7, (byte) 23, CALCommandType.STATUS),
  CALCommandStatus_24Bytes((short) 0xD8, (byte) 24, CALCommandType.STATUS),
  CALCommandStatus_25Bytes((short) 0xD9, (byte) 25, CALCommandType.STATUS),
  CALCommandStatus_26Bytes((short) 0xDA, (byte) 26, CALCommandType.STATUS),
  CALCommandStatus_27Bytes((short) 0xDB, (byte) 27, CALCommandType.STATUS),
  CALCommandStatus_28Bytes((short) 0xDC, (byte) 28, CALCommandType.STATUS),
  CALCommandStatus_29Bytes((short) 0xDD, (byte) 29, CALCommandType.STATUS),
  CALCommandStatus_30Bytes((short) 0xDE, (byte) 30, CALCommandType.STATUS),
  CALCommandStatus_31Bytes((short) 0xDF, (byte) 31, CALCommandType.STATUS),
  CALCommandStatusExtended_0Bytes((short) 0xE0, (byte) 0, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_1Bytes((short) 0xE1, (byte) 1, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_2Bytes((short) 0xE2, (byte) 2, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_3Bytes((short) 0xE3, (byte) 3, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_4Bytes((short) 0xE4, (byte) 4, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_5Bytes((short) 0xE5, (byte) 5, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_6Bytes((short) 0xE6, (byte) 6, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_7Bytes((short) 0xE7, (byte) 7, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_8Bytes((short) 0xE8, (byte) 8, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_9Bytes((short) 0xE9, (byte) 9, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_10Bytes((short) 0xEA, (byte) 10, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_11Bytes((short) 0xEB, (byte) 11, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_12Bytes((short) 0xEC, (byte) 12, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_13Bytes((short) 0xED, (byte) 13, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_14Bytes((short) 0xEE, (byte) 14, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_15Bytes((short) 0xEF, (byte) 15, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_16Bytes((short) 0xF0, (byte) 16, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_17Bytes((short) 0xF1, (byte) 17, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_18Bytes((short) 0xF2, (byte) 18, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_19Bytes((short) 0xF3, (byte) 19, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_20Bytes((short) 0xF4, (byte) 20, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_21Bytes((short) 0xF5, (byte) 21, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_22Bytes((short) 0xF6, (byte) 22, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_23Bytes((short) 0xF7, (byte) 23, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_24Bytes((short) 0xF8, (byte) 24, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_25Bytes((short) 0xF9, (byte) 25, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_26Bytes((short) 0xFA, (byte) 26, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_27Bytes((short) 0xFB, (byte) 27, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_28Bytes((short) 0xFC, (byte) 28, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_29Bytes((short) 0xFD, (byte) 29, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_30Bytes((short) 0xFE, (byte) 30, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_31Bytes((short) 0xFF, (byte) 31, CALCommandType.STATUS_EXTENDED);
  private static final Map<Short, CALCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (CALCommandTypeContainer value : CALCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final byte numBytes;
  private final CALCommandType commandType;

  CALCommandTypeContainer(short value, byte numBytes, CALCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public byte getNumBytes() {
    return numBytes;
  }

  public static CALCommandTypeContainer firstEnumForFieldNumBytes(byte fieldValue) {
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<CALCommandTypeContainer> enumsForFieldNumBytes(byte fieldValue) {
    List<CALCommandTypeContainer> _values = new ArrayList<>();
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public CALCommandType getCommandType() {
    return commandType;
  }

  public static CALCommandTypeContainer firstEnumForFieldCommandType(CALCommandType fieldValue) {
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<CALCommandTypeContainer> enumsForFieldCommandType(CALCommandType fieldValue) {
    List<CALCommandTypeContainer> _values = new ArrayList<>();
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static CALCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}