/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../contract/params";
import { SmartContract } from "../contract/smart_contract";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "arran8901.chainlogplatform.contract";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSmartContractRequest {
  address: string;
}

export interface QueryGetSmartContractResponse {
  smartContract: SmartContract | undefined;
}

export interface QueryAllSmartContractRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSmartContractResponse {
  smartContract: SmartContract[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetSmartContractRequest: object = { address: "" };

export const QueryGetSmartContractRequest = {
  encode(
    message: QueryGetSmartContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSmartContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSmartContractRequest,
    } as QueryGetSmartContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSmartContractRequest {
    const message = {
      ...baseQueryGetSmartContractRequest,
    } as QueryGetSmartContractRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetSmartContractRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSmartContractRequest>
  ): QueryGetSmartContractRequest {
    const message = {
      ...baseQueryGetSmartContractRequest,
    } as QueryGetSmartContractRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetSmartContractResponse: object = {};

export const QueryGetSmartContractResponse = {
  encode(
    message: QueryGetSmartContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.smartContract !== undefined) {
      SmartContract.encode(
        message.smartContract,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSmartContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSmartContractResponse,
    } as QueryGetSmartContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.smartContract = SmartContract.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSmartContractResponse {
    const message = {
      ...baseQueryGetSmartContractResponse,
    } as QueryGetSmartContractResponse;
    if (object.smartContract !== undefined && object.smartContract !== null) {
      message.smartContract = SmartContract.fromJSON(object.smartContract);
    } else {
      message.smartContract = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSmartContractResponse): unknown {
    const obj: any = {};
    message.smartContract !== undefined &&
      (obj.smartContract = message.smartContract
        ? SmartContract.toJSON(message.smartContract)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSmartContractResponse>
  ): QueryGetSmartContractResponse {
    const message = {
      ...baseQueryGetSmartContractResponse,
    } as QueryGetSmartContractResponse;
    if (object.smartContract !== undefined && object.smartContract !== null) {
      message.smartContract = SmartContract.fromPartial(object.smartContract);
    } else {
      message.smartContract = undefined;
    }
    return message;
  },
};

const baseQueryAllSmartContractRequest: object = {};

export const QueryAllSmartContractRequest = {
  encode(
    message: QueryAllSmartContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSmartContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSmartContractRequest,
    } as QueryAllSmartContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSmartContractRequest {
    const message = {
      ...baseQueryAllSmartContractRequest,
    } as QueryAllSmartContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSmartContractRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSmartContractRequest>
  ): QueryAllSmartContractRequest {
    const message = {
      ...baseQueryAllSmartContractRequest,
    } as QueryAllSmartContractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSmartContractResponse: object = {};

export const QueryAllSmartContractResponse = {
  encode(
    message: QueryAllSmartContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.smartContract) {
      SmartContract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSmartContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSmartContractResponse,
    } as QueryAllSmartContractResponse;
    message.smartContract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.smartContract.push(
            SmartContract.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSmartContractResponse {
    const message = {
      ...baseQueryAllSmartContractResponse,
    } as QueryAllSmartContractResponse;
    message.smartContract = [];
    if (object.smartContract !== undefined && object.smartContract !== null) {
      for (const e of object.smartContract) {
        message.smartContract.push(SmartContract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSmartContractResponse): unknown {
    const obj: any = {};
    if (message.smartContract) {
      obj.smartContract = message.smartContract.map((e) =>
        e ? SmartContract.toJSON(e) : undefined
      );
    } else {
      obj.smartContract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSmartContractResponse>
  ): QueryAllSmartContractResponse {
    const message = {
      ...baseQueryAllSmartContractResponse,
    } as QueryAllSmartContractResponse;
    message.smartContract = [];
    if (object.smartContract !== undefined && object.smartContract !== null) {
      for (const e of object.smartContract) {
        message.smartContract.push(SmartContract.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SmartContract by index. */
  SmartContract(
    request: QueryGetSmartContractRequest
  ): Promise<QueryGetSmartContractResponse>;
  /** Queries a list of SmartContract items. */
  SmartContractAll(
    request: QueryAllSmartContractRequest
  ): Promise<QueryAllSmartContractResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SmartContract(
    request: QueryGetSmartContractRequest
  ): Promise<QueryGetSmartContractResponse> {
    const data = QueryGetSmartContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "SmartContract",
      data
    );
    return promise.then((data) =>
      QueryGetSmartContractResponse.decode(new Reader(data))
    );
  }

  SmartContractAll(
    request: QueryAllSmartContractRequest
  ): Promise<QueryAllSmartContractResponse> {
    const data = QueryAllSmartContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "SmartContractAll",
      data
    );
    return promise.then((data) =>
      QueryAllSmartContractResponse.decode(new Reader(data))
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
