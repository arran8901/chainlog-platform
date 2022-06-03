/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "arran8901.chainlogplatform.contract";

export interface MsgCreateContract {
  creator: string;
  code: string;
  value: string;
}

export interface MsgCreateContractResponse {
  contractAddress: string;
}

export interface MsgCallContract {
  creator: string;
  contractAddress: string;
  value: string;
  messageTerm: string;
}

export interface MsgCallContractResponse {}

const baseMsgCreateContract: object = { creator: "", code: "", value: "" };

export const MsgCreateContract = {
  encode(message: MsgCreateContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.code = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: MsgCreateContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.code !== undefined && (obj.code = message.code);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateContract>): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseMsgCreateContractResponse: object = { contractAddress: "" };

export const MsgCreateContractResponse = {
  encode(
    message: MsgCreateContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.contractAddress !== "") {
      writer.uint32(10).string(message.contractAddress);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contractAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = String(object.contractAddress);
    } else {
      message.contractAddress = "";
    }
    return message;
  },

  toJSON(message: MsgCreateContractResponse): unknown {
    const obj: any = {};
    message.contractAddress !== undefined &&
      (obj.contractAddress = message.contractAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateContractResponse>
  ): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = object.contractAddress;
    } else {
      message.contractAddress = "";
    }
    return message;
  },
};

const baseMsgCallContract: object = {
  creator: "",
  contractAddress: "",
  value: "",
  messageTerm: "",
};

export const MsgCallContract = {
  encode(message: MsgCallContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.contractAddress !== "") {
      writer.uint32(18).string(message.contractAddress);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    if (message.messageTerm !== "") {
      writer.uint32(34).string(message.messageTerm);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCallContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCallContract } as MsgCallContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.contractAddress = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          message.messageTerm = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCallContract {
    const message = { ...baseMsgCallContract } as MsgCallContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = String(object.contractAddress);
    } else {
      message.contractAddress = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.messageTerm !== undefined && object.messageTerm !== null) {
      message.messageTerm = String(object.messageTerm);
    } else {
      message.messageTerm = "";
    }
    return message;
  },

  toJSON(message: MsgCallContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.contractAddress !== undefined &&
      (obj.contractAddress = message.contractAddress);
    message.value !== undefined && (obj.value = message.value);
    message.messageTerm !== undefined &&
      (obj.messageTerm = message.messageTerm);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCallContract>): MsgCallContract {
    const message = { ...baseMsgCallContract } as MsgCallContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = object.contractAddress;
    } else {
      message.contractAddress = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.messageTerm !== undefined && object.messageTerm !== null) {
      message.messageTerm = object.messageTerm;
    } else {
      message.messageTerm = "";
    }
    return message;
  },
};

const baseMsgCallContractResponse: object = {};

export const MsgCallContractResponse = {
  encode(_: MsgCallContractResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCallContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCallContractResponse,
    } as MsgCallContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCallContractResponse {
    const message = {
      ...baseMsgCallContractResponse,
    } as MsgCallContractResponse;
    return message;
  },

  toJSON(_: MsgCallContractResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCallContractResponse>
  ): MsgCallContractResponse {
    const message = {
      ...baseMsgCallContractResponse,
    } as MsgCallContractResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CallContract(request: MsgCallContract): Promise<MsgCallContractResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse> {
    const data = MsgCreateContract.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Msg",
      "CreateContract",
      data
    );
    return promise.then((data) =>
      MsgCreateContractResponse.decode(new Reader(data))
    );
  }

  CallContract(request: MsgCallContract): Promise<MsgCallContractResponse> {
    const data = MsgCallContract.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Msg",
      "CallContract",
      data
    );
    return promise.then((data) =>
      MsgCallContractResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
