/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.location.v1";

export interface ListFriendLocationRequest {
}

export interface ListFriendLocationResponse {
  items: Location[];
}

export interface Location {
  latitude: number;
  longitude: number;
  userId: string;
  name: string;
  username: string;
  picture: string;
  verified: boolean;
}

export interface UpdateLocationRequest {
  userId: string;
  /** The latitude in degrees. It must be in the range [-90.0, +90.0]. */
  latitude: number;
  /** The longitude in degrees. It must be in the range [-180.0, +180.0]. */
  longitude: number;
}

export interface UpdateLocationResponse {
}

function createBaseListFriendLocationRequest(): ListFriendLocationRequest {
  return {};
}

export const ListFriendLocationRequest = {
  encode(_: ListFriendLocationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendLocationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendLocationRequest();
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

  fromJSON(_: any): ListFriendLocationRequest {
    return {};
  },

  toJSON(_: ListFriendLocationRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendLocationRequest>, I>>(base?: I): ListFriendLocationRequest {
    return ListFriendLocationRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendLocationRequest>, I>>(_: I): ListFriendLocationRequest {
    const message = createBaseListFriendLocationRequest();
    return message;
  },
};

function createBaseListFriendLocationResponse(): ListFriendLocationResponse {
  return { items: [] };
}

export const ListFriendLocationResponse = {
  encode(message: ListFriendLocationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      Location.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendLocationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendLocationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(Location.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListFriendLocationResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => Location.fromJSON(e)) : [] };
  },

  toJSON(message: ListFriendLocationResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? Location.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendLocationResponse>, I>>(base?: I): ListFriendLocationResponse {
    return ListFriendLocationResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendLocationResponse>, I>>(object: I): ListFriendLocationResponse {
    const message = createBaseListFriendLocationResponse();
    message.items = object.items?.map((e) => Location.fromPartial(e)) || [];
    return message;
  },
};

function createBaseLocation(): Location {
  return { latitude: 0, longitude: 0, userId: "", name: "", username: "", picture: "", verified: false };
}

export const Location = {
  encode(message: Location, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.latitude !== 0) {
      writer.uint32(9).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(17).double(message.longitude);
    }
    if (message.userId !== "") {
      writer.uint32(26).string(message.userId);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.username !== "") {
      writer.uint32(42).string(message.username);
    }
    if (message.picture !== "") {
      writer.uint32(50).string(message.picture);
    }
    if (message.verified === true) {
      writer.uint32(56).bool(message.verified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Location {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.latitude = reader.double();
          break;
        case 2:
          message.longitude = reader.double();
          break;
        case 3:
          message.userId = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.username = reader.string();
          break;
        case 6:
          message.picture = reader.string();
          break;
        case 7:
          message.verified = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Location {
    return {
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
      userId: isSet(object.userId) ? String(object.userId) : "",
      name: isSet(object.name) ? String(object.name) : "",
      username: isSet(object.username) ? String(object.username) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      verified: isSet(object.verified) ? Boolean(object.verified) : false,
    };
  },

  toJSON(message: Location): unknown {
    const obj: any = {};
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    message.userId !== undefined && (obj.userId = message.userId);
    message.name !== undefined && (obj.name = message.name);
    message.username !== undefined && (obj.username = message.username);
    message.picture !== undefined && (obj.picture = message.picture);
    message.verified !== undefined && (obj.verified = message.verified);
    return obj;
  },

  create<I extends Exact<DeepPartial<Location>, I>>(base?: I): Location {
    return Location.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Location>, I>>(object: I): Location {
    const message = createBaseLocation();
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
    message.userId = object.userId ?? "";
    message.name = object.name ?? "";
    message.username = object.username ?? "";
    message.picture = object.picture ?? "";
    message.verified = object.verified ?? false;
    return message;
  },
};

function createBaseUpdateLocationRequest(): UpdateLocationRequest {
  return { userId: "", latitude: 0, longitude: 0 };
}

export const UpdateLocationRequest = {
  encode(message: UpdateLocationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.latitude !== 0) {
      writer.uint32(17).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(25).double(message.longitude);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateLocationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateLocationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        case 2:
          message.latitude = reader.double();
          break;
        case 3:
          message.longitude = reader.double();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateLocationRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
    };
  },

  toJSON(message: UpdateLocationRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateLocationRequest>, I>>(base?: I): UpdateLocationRequest {
    return UpdateLocationRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateLocationRequest>, I>>(object: I): UpdateLocationRequest {
    const message = createBaseUpdateLocationRequest();
    message.userId = object.userId ?? "";
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
    return message;
  },
};

function createBaseUpdateLocationResponse(): UpdateLocationResponse {
  return {};
}

export const UpdateLocationResponse = {
  encode(_: UpdateLocationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateLocationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateLocationResponse();
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

  fromJSON(_: any): UpdateLocationResponse {
    return {};
  },

  toJSON(_: UpdateLocationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateLocationResponse>, I>>(base?: I): UpdateLocationResponse {
    return UpdateLocationResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateLocationResponse>, I>>(_: I): UpdateLocationResponse {
    const message = createBaseUpdateLocationResponse();
    return message;
  },
};

export interface LocationService {
  UpdateLocation(request: UpdateLocationRequest): Promise<UpdateLocationResponse>;
  ListFriendLocation(request: ListFriendLocationRequest): Promise<ListFriendLocationResponse>;
}

export class LocationServiceClientImpl implements LocationService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.location.v1.LocationService";
    this.rpc = rpc;
    this.UpdateLocation = this.UpdateLocation.bind(this);
    this.ListFriendLocation = this.ListFriendLocation.bind(this);
  }
  UpdateLocation(request: UpdateLocationRequest): Promise<UpdateLocationResponse> {
    const data = UpdateLocationRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateLocation", data);
    return promise.then((data) => UpdateLocationResponse.decode(new _m0.Reader(data)));
  }

  ListFriendLocation(request: ListFriendLocationRequest): Promise<ListFriendLocationResponse> {
    const data = ListFriendLocationRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFriendLocation", data);
    return promise.then((data) => ListFriendLocationResponse.decode(new _m0.Reader(data)));
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
