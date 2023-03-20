/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.notification.v1";

export interface CreateRegistrationTokenRequest {
  token: string;
}

export interface CreateRegistrationTokenResponse {
}

function createBaseCreateRegistrationTokenRequest(): CreateRegistrationTokenRequest {
  return { token: "" };
}

export const CreateRegistrationTokenRequest = {
  encode(message: CreateRegistrationTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateRegistrationTokenRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRegistrationTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateRegistrationTokenRequest {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: CreateRegistrationTokenRequest): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateRegistrationTokenRequest>, I>>(base?: I): CreateRegistrationTokenRequest {
    return CreateRegistrationTokenRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateRegistrationTokenRequest>, I>>(
    object: I,
  ): CreateRegistrationTokenRequest {
    const message = createBaseCreateRegistrationTokenRequest();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseCreateRegistrationTokenResponse(): CreateRegistrationTokenResponse {
  return {};
}

export const CreateRegistrationTokenResponse = {
  encode(_: CreateRegistrationTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateRegistrationTokenResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRegistrationTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CreateRegistrationTokenResponse {
    return {};
  },

  toJSON(_: CreateRegistrationTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateRegistrationTokenResponse>, I>>(base?: I): CreateRegistrationTokenResponse {
    return CreateRegistrationTokenResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateRegistrationTokenResponse>, I>>(_: I): CreateRegistrationTokenResponse {
    const message = createBaseCreateRegistrationTokenResponse();
    return message;
  },
};

export interface NotificationService {
  CreateRegistrationToken(request: CreateRegistrationTokenRequest): Promise<CreateRegistrationTokenResponse>;
}

export class NotificationServiceClientImpl implements NotificationService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.notification.v1.NotificationService";
    this.rpc = rpc;
    this.CreateRegistrationToken = this.CreateRegistrationToken.bind(this);
  }
  CreateRegistrationToken(request: CreateRegistrationTokenRequest): Promise<CreateRegistrationTokenResponse> {
    const data = CreateRegistrationTokenRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateRegistrationToken", data);
    return promise.then((data) => CreateRegistrationTokenResponse.decode(_m0.Reader.create(data)));
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
