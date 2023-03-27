/* eslint-disable */
import * as Long from "long";
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
  locationName: string;
  lastUpdated: number;
}

export interface UpdateGhostModeRequest {
  ghostMode: boolean;
}

export interface UpdateGhostModeResponse {
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendLocationRequest();
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendLocationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.items.push(Location.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
  return {
    latitude: 0,
    longitude: 0,
    userId: "",
    name: "",
    username: "",
    picture: "",
    verified: false,
    locationName: "",
    lastUpdated: 0,
  };
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
    if (message.locationName !== "") {
      writer.uint32(66).string(message.locationName);
    }
    if (message.lastUpdated !== 0) {
      writer.uint32(72).int64(message.lastUpdated);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Location {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 9) {
            break;
          }

          message.latitude = reader.double();
          continue;
        case 2:
          if (tag != 17) {
            break;
          }

          message.longitude = reader.double();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.name = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.username = reader.string();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.verified = reader.bool();
          continue;
        case 8:
          if (tag != 66) {
            break;
          }

          message.locationName = reader.string();
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.lastUpdated = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      lastUpdated: isSet(object.lastUpdated) ? Number(object.lastUpdated) : 0,
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
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.lastUpdated !== undefined && (obj.lastUpdated = Math.round(message.lastUpdated));
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
    message.locationName = object.locationName ?? "";
    message.lastUpdated = object.lastUpdated ?? 0;
    return message;
  },
};

function createBaseUpdateGhostModeRequest(): UpdateGhostModeRequest {
  return { ghostMode: false };
}

export const UpdateGhostModeRequest = {
  encode(message: UpdateGhostModeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ghostMode === true) {
      writer.uint32(8).bool(message.ghostMode);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateGhostModeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateGhostModeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.ghostMode = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateGhostModeRequest {
    return { ghostMode: isSet(object.ghostMode) ? Boolean(object.ghostMode) : false };
  },

  toJSON(message: UpdateGhostModeRequest): unknown {
    const obj: any = {};
    message.ghostMode !== undefined && (obj.ghostMode = message.ghostMode);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateGhostModeRequest>, I>>(base?: I): UpdateGhostModeRequest {
    return UpdateGhostModeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateGhostModeRequest>, I>>(object: I): UpdateGhostModeRequest {
    const message = createBaseUpdateGhostModeRequest();
    message.ghostMode = object.ghostMode ?? false;
    return message;
  },
};

function createBaseUpdateGhostModeResponse(): UpdateGhostModeResponse {
  return {};
}

export const UpdateGhostModeResponse = {
  encode(_: UpdateGhostModeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateGhostModeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateGhostModeResponse();
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

  fromJSON(_: any): UpdateGhostModeResponse {
    return {};
  },

  toJSON(_: UpdateGhostModeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateGhostModeResponse>, I>>(base?: I): UpdateGhostModeResponse {
    return UpdateGhostModeResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateGhostModeResponse>, I>>(_: I): UpdateGhostModeResponse {
    const message = createBaseUpdateGhostModeResponse();
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateLocationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 2:
          if (tag != 17) {
            break;
          }

          message.latitude = reader.double();
          continue;
        case 3:
          if (tag != 25) {
            break;
          }

          message.longitude = reader.double();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateLocationResponse();
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
  UpdateGhostMode(request: UpdateGhostModeRequest): Promise<UpdateGhostModeResponse>;
  ListFriendLocation(request: ListFriendLocationRequest): Promise<ListFriendLocationResponse>;
}

export class LocationServiceClientImpl implements LocationService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.location.v1.LocationService";
    this.rpc = rpc;
    this.UpdateLocation = this.UpdateLocation.bind(this);
    this.UpdateGhostMode = this.UpdateGhostMode.bind(this);
    this.ListFriendLocation = this.ListFriendLocation.bind(this);
  }
  UpdateLocation(request: UpdateLocationRequest): Promise<UpdateLocationResponse> {
    const data = UpdateLocationRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateLocation", data);
    return promise.then((data) => UpdateLocationResponse.decode(_m0.Reader.create(data)));
  }

  UpdateGhostMode(request: UpdateGhostModeRequest): Promise<UpdateGhostModeResponse> {
    const data = UpdateGhostModeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateGhostMode", data);
    return promise.then((data) => UpdateGhostModeResponse.decode(_m0.Reader.create(data)));
  }

  ListFriendLocation(request: ListFriendLocationRequest): Promise<ListFriendLocationResponse> {
    const data = ListFriendLocationRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFriendLocation", data);
    return promise.then((data) => ListFriendLocationResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
