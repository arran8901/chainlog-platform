/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../contract/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { SmartContractWithAddress } from "../contract/smart_contract";

export const protobufPackage = "arran8901.chainlogplatform.contract";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryContractCodeRequest {
  contractAddress: string;
}

export interface QueryContractCodeResponse {
  code: string;
  dynamicKb: string;
}

export interface QueryAllContractsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllContractsResponse {
  contracts: SmartContractWithAddress[];
  pagination: PageResponse | undefined;
}

export interface QueryQueryContractRequest {
  contractAddress: string;
  query: string;
  nDerivations: string;
}

export interface QueryQueryContractResponse {}

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

const baseQueryContractCodeRequest: object = { contractAddress: "" };

export const QueryContractCodeRequest = {
  encode(
    message: QueryContractCodeRequest,
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
  ): QueryContractCodeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryContractCodeRequest,
    } as QueryContractCodeRequest;
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

  fromJSON(object: any): QueryContractCodeRequest {
    const message = {
      ...baseQueryContractCodeRequest,
    } as QueryContractCodeRequest;
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

  toJSON(message: QueryContractCodeRequest): unknown {
    const obj: any = {};
    message.contractAddress !== undefined &&
      (obj.contractAddress = message.contractAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryContractCodeRequest>
  ): QueryContractCodeRequest {
    const message = {
      ...baseQueryContractCodeRequest,
    } as QueryContractCodeRequest;
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

const baseQueryContractCodeResponse: object = { code: "", dynamicKb: "" };

export const QueryContractCodeResponse = {
  encode(
    message: QueryContractCodeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    if (message.dynamicKb !== "") {
      writer.uint32(18).string(message.dynamicKb);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryContractCodeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryContractCodeResponse,
    } as QueryContractCodeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        case 2:
          message.dynamicKb = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryContractCodeResponse {
    const message = {
      ...baseQueryContractCodeResponse,
    } as QueryContractCodeResponse;
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

  toJSON(message: QueryContractCodeResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.dynamicKb !== undefined && (obj.dynamicKb = message.dynamicKb);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryContractCodeResponse>
  ): QueryContractCodeResponse {
    const message = {
      ...baseQueryContractCodeResponse,
    } as QueryContractCodeResponse;
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

const baseQueryAllContractsRequest: object = {};

export const QueryAllContractsRequest = {
  encode(
    message: QueryAllContractsRequest,
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
  ): QueryAllContractsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
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

  fromJSON(object: any): QueryAllContractsRequest {
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractsRequest>
  ): QueryAllContractsRequest {
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllContractsResponse: object = {};

export const QueryAllContractsResponse = {
  encode(
    message: QueryAllContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.contracts) {
      SmartContractWithAddress.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contracts.push(
            SmartContractWithAddress.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllContractsResponse {
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    if (object.contracts !== undefined && object.contracts !== null) {
      for (const e of object.contracts) {
        message.contracts.push(SmartContractWithAddress.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractsResponse): unknown {
    const obj: any = {};
    if (message.contracts) {
      obj.contracts = message.contracts.map((e) =>
        e ? SmartContractWithAddress.toJSON(e) : undefined
      );
    } else {
      obj.contracts = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractsResponse>
  ): QueryAllContractsResponse {
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    if (object.contracts !== undefined && object.contracts !== null) {
      for (const e of object.contracts) {
        message.contracts.push(SmartContractWithAddress.fromPartial(e));
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

const baseQueryQueryContractRequest: object = {
  contractAddress: "",
  query: "",
  nDerivations: "",
};

export const QueryQueryContractRequest = {
  encode(
    message: QueryQueryContractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.contractAddress !== "") {
      writer.uint32(10).string(message.contractAddress);
    }
    if (message.query !== "") {
      writer.uint32(18).string(message.query);
    }
    if (message.nDerivations !== "") {
      writer.uint32(26).string(message.nDerivations);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQueryContractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQueryContractRequest,
    } as QueryQueryContractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contractAddress = reader.string();
          break;
        case 2:
          message.query = reader.string();
          break;
        case 3:
          message.nDerivations = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryQueryContractRequest {
    const message = {
      ...baseQueryQueryContractRequest,
    } as QueryQueryContractRequest;
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = String(object.contractAddress);
    } else {
      message.contractAddress = "";
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = String(object.query);
    } else {
      message.query = "";
    }
    if (object.nDerivations !== undefined && object.nDerivations !== null) {
      message.nDerivations = String(object.nDerivations);
    } else {
      message.nDerivations = "";
    }
    return message;
  },

  toJSON(message: QueryQueryContractRequest): unknown {
    const obj: any = {};
    message.contractAddress !== undefined &&
      (obj.contractAddress = message.contractAddress);
    message.query !== undefined && (obj.query = message.query);
    message.nDerivations !== undefined &&
      (obj.nDerivations = message.nDerivations);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryQueryContractRequest>
  ): QueryQueryContractRequest {
    const message = {
      ...baseQueryQueryContractRequest,
    } as QueryQueryContractRequest;
    if (
      object.contractAddress !== undefined &&
      object.contractAddress !== null
    ) {
      message.contractAddress = object.contractAddress;
    } else {
      message.contractAddress = "";
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = object.query;
    } else {
      message.query = "";
    }
    if (object.nDerivations !== undefined && object.nDerivations !== null) {
      message.nDerivations = object.nDerivations;
    } else {
      message.nDerivations = "";
    }
    return message;
  },
};

const baseQueryQueryContractResponse: object = {};

export const QueryQueryContractResponse = {
  encode(
    _: QueryQueryContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryQueryContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryQueryContractResponse,
    } as QueryQueryContractResponse;
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

  fromJSON(_: any): QueryQueryContractResponse {
    const message = {
      ...baseQueryQueryContractResponse,
    } as QueryQueryContractResponse;
    return message;
  },

  toJSON(_: QueryQueryContractResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryQueryContractResponse>
  ): QueryQueryContractResponse {
    const message = {
      ...baseQueryQueryContractResponse,
    } as QueryQueryContractResponse;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of ContractCode items. */
  ContractCode(
    request: QueryContractCodeRequest
  ): Promise<QueryContractCodeResponse>;
  /** Queries a list of AllContracts items. */
  AllContracts(
    request: QueryAllContractsRequest
  ): Promise<QueryAllContractsResponse>;
  /** Queries a list of QueryContract items. */
  QueryContract(
    request: QueryQueryContractRequest
  ): Promise<QueryQueryContractResponse>;
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

  ContractCode(
    request: QueryContractCodeRequest
  ): Promise<QueryContractCodeResponse> {
    const data = QueryContractCodeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "ContractCode",
      data
    );
    return promise.then((data) =>
      QueryContractCodeResponse.decode(new Reader(data))
    );
  }

  AllContracts(
    request: QueryAllContractsRequest
  ): Promise<QueryAllContractsResponse> {
    const data = QueryAllContractsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "AllContracts",
      data
    );
    return promise.then((data) =>
      QueryAllContractsResponse.decode(new Reader(data))
    );
  }

  QueryContract(
    request: QueryQueryContractRequest
  ): Promise<QueryQueryContractResponse> {
    const data = QueryQueryContractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "arran8901.chainlogplatform.contract.Query",
      "QueryContract",
      data
    );
    return promise.then((data) =>
      QueryQueryContractResponse.decode(new Reader(data))
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
