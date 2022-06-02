/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "arran8901.chainlogplatform.contract";

export interface SmartContract {
  address: string;
  code: string;
  dynamicKb: string;
}

const baseSmartContract: object = { address: "", code: "", dynamicKb: "" };

export const SmartContract = {
  encode(message: SmartContract, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.dynamicKb !== "") {
      writer.uint32(26).string(message.dynamicKb);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SmartContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSmartContract } as SmartContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.code = reader.string();
          break;
        case 3:
          message.dynamicKb = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SmartContract {
    const message = { ...baseSmartContract } as SmartContract;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (object.dynamicKb !== undefined && object.dynamicKb !== null) {
      message.dynamicKb = String(object.dynamicKb);
    } else {
      message.dynamicKb = "";
    }
    return message;
  },

  toJSON(message: SmartContract): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.code !== undefined && (obj.code = message.code);
    message.dynamicKb !== undefined && (obj.dynamicKb = message.dynamicKb);
    return obj;
  },

  fromPartial(object: DeepPartial<SmartContract>): SmartContract {
    const message = { ...baseSmartContract } as SmartContract;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (object.dynamicKb !== undefined && object.dynamicKb !== null) {
      message.dynamicKb = object.dynamicKb;
    } else {
      message.dynamicKb = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
