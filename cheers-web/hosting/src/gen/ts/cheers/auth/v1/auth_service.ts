/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "../../../google/protobuf/timestamp";

export const protobufPackage = "cheers.auth.v1";

export interface CreateModeratorRequest {
  userId: string;
}

export interface CreateModeratorResponse {
}

export interface CreateBusinessAccountRequest {
  userId: string;
}

export interface CreateBusinessAccountResponse {
}

function createBaseCreateModeratorRequest(): CreateModeratorRequest {
  return { userId: "" };
}

export const CreateModeratorRequest = {
  encode(message: CreateModeratorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateModeratorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateModeratorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateModeratorRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: CreateModeratorRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateModeratorRequest>, I>>(object: I): CreateModeratorRequest {
    const message = createBaseCreateModeratorRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseCreateModeratorResponse(): CreateModeratorResponse {
  return {};
}

export const CreateModeratorResponse = {
  encode(_: CreateModeratorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateModeratorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateModeratorResponse();
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

  fromJSON(_: any): CreateModeratorResponse {
    return {};
  },

  toJSON(_: CreateModeratorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateModeratorResponse>, I>>(_: I): CreateModeratorResponse {
    const message = createBaseCreateModeratorResponse();
    return message;
  },
};

function createBaseCreateBusinessAccountRequest(): CreateBusinessAccountRequest {
  return { userId: "" };
}

export const CreateBusinessAccountRequest = {
  encode(message: CreateBusinessAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBusinessAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBusinessAccountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateBusinessAccountRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: CreateBusinessAccountRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBusinessAccountRequest>, I>>(object: I): CreateBusinessAccountRequest {
    const message = createBaseCreateBusinessAccountRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseCreateBusinessAccountResponse(): CreateBusinessAccountResponse {
  return {};
}

export const CreateBusinessAccountResponse = {
  encode(_: CreateBusinessAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBusinessAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBusinessAccountResponse();
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

  fromJSON(_: any): CreateBusinessAccountResponse {
    return {};
  },

  toJSON(_: CreateBusinessAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBusinessAccountResponse>, I>>(_: I): CreateBusinessAccountResponse {
    const message = createBaseCreateBusinessAccountResponse();
    return message;
  },
};

export interface AuthService {
  CreateModerator(request: CreateModeratorRequest): Promise<CreateModeratorResponse>;
  CreateBusinessAccount(request: CreateBusinessAccountRequest): Promise<CreateBusinessAccountResponse>;
}

export class AuthServiceClientImpl implements AuthService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.auth.v1.AuthService";
    this.rpc = rpc;
    this.CreateModerator = this.CreateModerator.bind(this);
    this.CreateBusinessAccount = this.CreateBusinessAccount.bind(this);
  }
  CreateModerator(request: CreateModeratorRequest): Promise<CreateModeratorResponse> {
    const data = CreateModeratorRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateModerator", data);
    return promise.then((data) => CreateModeratorResponse.decode(new _m0.Reader(data)));
  }

  CreateBusinessAccount(request: CreateBusinessAccountRequest): Promise<CreateBusinessAccountResponse> {
    const data = CreateBusinessAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateBusinessAccount", data);
    return promise.then((data) => CreateBusinessAccountResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
