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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetConfirmedServiceChoice {
  ACKNOWLEDGE_ALARM((short) 0x00),
  CONFIRMED_COV_NOTIFICATION((short) 0x01),
  CONFIRMED_COV_NOTIFICATION_MULTIPLE((short) 0x1F),
  CONFIRMED_EVENT_NOTIFICATION((short) 0x02),
  GET_ALARM_SUMMARY((short) 0x03),
  GET_ENROLLMENT_SUMMARY((short) 0x04),
  GET_EVENT_INFORMATION((short) 0x1D),
  LIFE_SAFETY_OPERATION((short) 0x1B),
  SUBSCRIBE_COV((short) 0x05),
  SUBSCRIBE_COV_PROPERTY((short) 0x1C),
  SUBSCRIBE_COV_PROPERTY_MULTIPLE((short) 0x1E),
  ATOMIC_READ_FILE((short) 0x06),
  ATOMIC_WRITE_FILE((short) 0x07),
  ADD_LIST_ELEMENT((short) 0x08),
  REMOVE_LIST_ELEMENT((short) 0x09),
  CREATE_OBJECT((short) 0x0A),
  DELETE_OBJECT((short) 0x0B),
  READ_PROPERTY((short) 0x0C),
  READ_PROPERTY_MULTIPLE((short) 0x0E),
  READ_RANGE((short) 0x1A),
  WRITE_PROPERTY((short) 0x0F),
  WRITE_PROPERTY_MULTIPLE((short) 0x10),
  DEVICE_COMMUNICATION_CONTROL((short) 0x11),
  CONFIRMED_PRIVATE_TRANSFER((short) 0x12),
  CONFIRMED_TEXT_MESSAGE((short) 0x13),
  REINITIALIZE_DEVICE((short) 0x14),
  VT_OPEN((short) 0x15),
  VT_CLOSE((short) 0x16),
  VT_DATA((short) 0x17),
  AUTHENTICATE((short) 0x18),
  REQUEST_KEY((short) 0x19),
  READ_PROPERTY_CONDITIONAL((short) 0x0D);
  private static final Map<Short, BACnetConfirmedServiceChoice> map;

  static {
    map = new HashMap<>();
    for (BACnetConfirmedServiceChoice value : BACnetConfirmedServiceChoice.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  BACnetConfirmedServiceChoice(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static BACnetConfirmedServiceChoice enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}