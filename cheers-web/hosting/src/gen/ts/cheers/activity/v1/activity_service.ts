/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.activity.v1";

export interface ListActivityRequest {
  userId: string;
}

export interface ListActivityResponse {
  activities: Activity[];
}

export interface Activity {
  id: string;
  type: Activity_ActivityType;
  text: string;
  picture: string;
  userId: string;
  timestamp: number;
  mediaPicture: string;
  mediaId: string;
}

export enum Activity_ActivityType {
  LIKE_POST = 0,
  LIKE_STORY = 1,
  FOLLOW = 2,
  COMMENT_POST = 3,
  MENTION_POST = 4,
  UNRECOGNIZED = -1,
}

export function activity_ActivityTypeFromJSON(object: any): Activity_ActivityType {
  switch (object) {
    case 0:
    case "LIKE_POST":
      return Activity_ActivityType.LIKE_POST;
    case 1:
    case "LIKE_STORY":
      return Activity_ActivityType.LIKE_STORY;
    case 2:
    case "FOLLOW":
      return Activity_ActivityType.FOLLOW;
    case 3:
    case "COMMENT_POST":
      return Activity_ActivityType.COMMENT_POST;
    case 4:
    case "MENTION_POST":
      return Activity_ActivityType.MENTION_POST;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Activity_ActivityType.UNRECOGNIZED;
  }
}

export function activity_ActivityTypeToJSON(object: Activity_ActivityType): string {
  switch (object) {
    case Activity_ActivityType.LIKE_POST:
      return "LIKE_POST";
    case Activity_ActivityType.LIKE_STORY:
      return "LIKE_STORY";
    case Activity_ActivityType.FOLLOW:
      return "FOLLOW";
    case Activity_ActivityType.COMMENT_POST:
      return "COMMENT_POST";
    case Activity_ActivityType.MENTION_POST:
      return "MENTION_POST";
    case Activity_ActivityType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseListActivityRequest(): ListActivityRequest {
  return { userId: "" };
}

export const ListActivityRequest = {
  encode(message: ListActivityRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListActivityRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListActivityRequest();
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

  fromJSON(object: any): ListActivityRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: ListActivityRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListActivityRequest>, I>>(object: I): ListActivityRequest {
    const message = createBaseListActivityRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListActivityResponse(): ListActivityResponse {
  return { activities: [] };
}

export const ListActivityResponse = {
  encode(message: ListActivityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.activities) {
      Activity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListActivityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListActivityResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.activities.push(Activity.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListActivityResponse {
    return {
      activities: Array.isArray(object?.activities) ? object.activities.map((e: any) => Activity.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListActivityResponse): unknown {
    const obj: any = {};
    if (message.activities) {
      obj.activities = message.activities.map((e) => e ? Activity.toJSON(e) : undefined);
    } else {
      obj.activities = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListActivityResponse>, I>>(object: I): ListActivityResponse {
    const message = createBaseListActivityResponse();
    message.activities = object.activities?.map((e) => Activity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseActivity(): Activity {
  return { id: "", type: 0, text: "", picture: "", userId: "", timestamp: 0, mediaPicture: "", mediaId: "" };
}

export const Activity = {
  encode(message: Activity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(66).string(message.id);
    }
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    if (message.picture !== "") {
      writer.uint32(26).string(message.picture);
    }
    if (message.userId !== "") {
      writer.uint32(34).string(message.userId);
    }
    if (message.timestamp !== 0) {
      writer.uint32(40).int64(message.timestamp);
    }
    if (message.mediaPicture !== "") {
      writer.uint32(50).string(message.mediaPicture);
    }
    if (message.mediaId !== "") {
      writer.uint32(58).string(message.mediaId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Activity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 8:
          message.id = reader.string();
          break;
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.text = reader.string();
          break;
        case 3:
          message.picture = reader.string();
          break;
        case 4:
          message.userId = reader.string();
          break;
        case 5:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.mediaPicture = reader.string();
          break;
        case 7:
          message.mediaId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Activity {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      type: isSet(object.type) ? activity_ActivityTypeFromJSON(object.type) : 0,
      text: isSet(object.text) ? String(object.text) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
      mediaPicture: isSet(object.mediaPicture) ? String(object.mediaPicture) : "",
      mediaId: isSet(object.mediaId) ? String(object.mediaId) : "",
    };
  },

  toJSON(message: Activity): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = activity_ActivityTypeToJSON(message.type));
    message.text !== undefined && (obj.text = message.text);
    message.picture !== undefined && (obj.picture = message.picture);
    message.userId !== undefined && (obj.userId = message.userId);
    message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
    message.mediaPicture !== undefined && (obj.mediaPicture = message.mediaPicture);
    message.mediaId !== undefined && (obj.mediaId = message.mediaId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Activity>, I>>(object: I): Activity {
    const message = createBaseActivity();
    message.id = object.id ?? "";
    message.type = object.type ?? 0;
    message.text = object.text ?? "";
    message.picture = object.picture ?? "";
    message.userId = object.userId ?? "";
    message.timestamp = object.timestamp ?? 0;
    message.mediaPicture = object.mediaPicture ?? "";
    message.mediaId = object.mediaId ?? "";
    return message;
  },
};

export interface ActivityService {
  ListActivity(request: ListActivityRequest): Promise<ListActivityResponse>;
}

export class ActivityServiceClientImpl implements ActivityService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.activity.v1.ActivityService";
    this.rpc = rpc;
    this.ListActivity = this.ListActivity.bind(this);
  }
  ListActivity(request: ListActivityRequest): Promise<ListActivityResponse> {
    const data = ListActivityRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListActivity", data);
    return promise.then((data) => ListActivityResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
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
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
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
