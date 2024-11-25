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

public enum BACnetEventState {
  NORMAL((int) 0),
  FAULT((int) 1),
  OFFNORMAL((int) 2),
  HIGH_LIMIT((int) 3),
  LOW_LIMIT((int) 4),
  LIFE_SAVETY_ALARM((int) 5),
  VENDOR_PROPRIETARY_VALUE((int) 0xFFFF);
  private static final Map<Integer, BACnetEventState> map;

  static {
    map = new HashMap<>();
    for (BACnetEventState value : BACnetEventState.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  BACnetEventState(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static BACnetEventState enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
