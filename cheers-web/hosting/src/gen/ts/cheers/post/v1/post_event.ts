/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.post.v1";

export interface PostEvent {
  userId: string;
  post: PostEvent_Post | undefined;
  type: PostEvent_EventType;
}

export enum PostEvent_EventType {
  CREATE = 0,
  UPDATE = 1,
  DELETE = 2,
  LIKE = 3,
  COMMENT = 4,
  UNRECOGNIZED = -1,
}

export function postEvent_EventTypeFromJSON(object: any): PostEvent_EventType {
  switch (object) {
    case 0:
    case "CREATE":
      return PostEvent_EventType.CREATE;
    case 1:
    case "UPDATE":
      return PostEvent_EventType.UPDATE;
    case 2:
    case "DELETE":
      return PostEvent_EventType.DELETE;
    case 3:
    case "LIKE":
      return PostEvent_EventType.LIKE;
    case 4:
    case "COMMENT":
      return PostEvent_EventType.COMMENT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PostEvent_EventType.UNRECOGNIZED;
  }
}

export function postEvent_EventTypeToJSON(object: PostEvent_EventType): string {
  switch (object) {
    case PostEvent_EventType.CREATE:
      return "CREATE";
    case PostEvent_EventType.UPDATE:
      return "UPDATE";
    case PostEvent_EventType.DELETE:
      return "DELETE";
    case PostEvent_EventType.LIKE:
      return "LIKE";
    case PostEvent_EventType.COMMENT:
      return "COMMENT";
    case PostEvent_EventType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface PostEvent_Post {
  id: string;
  creatorId: string;
  caption: string;
  address: string;
  photos: string[];
  isCommentEnabled: boolean;
  locationName: string;
  drink: string;
  drunkenness: number;
  commentsEnabled: boolean;
  shareEnabled: boolean;
}

function createBasePostEvent(): PostEvent {
  return { userId: "", post: undefined, type: 0 };
}

export const PostEvent = {
  encode(message: PostEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.post !== undefined) {
      PostEvent_Post.encode(message.post, writer.uint32(18).fork()).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PostEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePostEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        case 2:
          message.post = PostEvent_Post.decode(reader, reader.uint32());
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PostEvent {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      post: isSet(object.post) ? PostEvent_Post.fromJSON(object.post) : undefined,
      type: isSet(object.type) ? postEvent_EventTypeFromJSON(object.type) : 0,
    };
  },

  toJSON(message: PostEvent): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.post !== undefined && (obj.post = message.post ? PostEvent_Post.toJSON(message.post) : undefined);
    message.type !== undefined && (obj.type = postEvent_EventTypeToJSON(message.type));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PostEvent>, I>>(object: I): PostEvent {
    const message = createBasePostEvent();
    message.userId = object.userId ?? "";
    message.post = (object.post !== undefined && object.post !== null)
      ? PostEvent_Post.fromPartial(object.post)
      : undefined;
    message.type = object.type ?? 0;
    return message;
  },
};

function createBasePostEvent_Post(): PostEvent_Post {
  return {
    id: "",
    creatorId: "",
    caption: "",
    address: "",
    photos: [],
    isCommentEnabled: false,
    locationName: "",
    drink: "",
    drunkenness: 0,
    commentsEnabled: false,
    shareEnabled: false,
  };
}

export const PostEvent_Post = {
  encode(message: PostEvent_Post, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creatorId !== "") {
      writer.uint32(18).string(message.creatorId);
    }
    if (message.caption !== "") {
      writer.uint32(26).string(message.caption);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    for (const v of message.photos) {
      writer.uint32(50).string(v!);
    }
    if (message.isCommentEnabled === true) {
      writer.uint32(56).bool(message.isCommentEnabled);
    }
    if (message.locationName !== "") {
      writer.uint32(66).string(message.locationName);
    }
    if (message.drink !== "") {
      writer.uint32(74).string(message.drink);
    }
    if (message.drunkenness !== 0) {
      writer.uint32(80).int64(message.drunkenness);
    }
    if (message.commentsEnabled === true) {
      writer.uint32(104).bool(message.commentsEnabled);
    }
    if (message.shareEnabled === true) {
      writer.uint32(112).bool(message.shareEnabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PostEvent_Post {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePostEvent_Post();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.creatorId = reader.string();
          break;
        case 3:
          message.caption = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 6:
          message.photos.push(reader.string());
          break;
        case 7:
          message.isCommentEnabled = reader.bool();
          break;
        case 8:
          message.locationName = reader.string();
          break;
        case 9:
          message.drink = reader.string();
          break;
        case 10:
          message.drunkenness = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.commentsEnabled = reader.bool();
          break;
        case 14:
          message.shareEnabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PostEvent_Post {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      creatorId: isSet(object.creatorId) ? String(object.creatorId) : "",
      caption: isSet(object.caption) ? String(object.caption) : "",
      address: isSet(object.address) ? String(object.address) : "",
      photos: Array.isArray(object?.photos) ? object.photos.map((e: any) => String(e)) : [],
      isCommentEnabled: isSet(object.isCommentEnabled) ? Boolean(object.isCommentEnabled) : false,
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      drink: isSet(object.drink) ? String(object.drink) : "",
      drunkenness: isSet(object.drunkenness) ? Number(object.drunkenness) : 0,
      commentsEnabled: isSet(object.commentsEnabled) ? Boolean(object.commentsEnabled) : false,
      shareEnabled: isSet(object.shareEnabled) ? Boolean(object.shareEnabled) : false,
    };
  },

  toJSON(message: PostEvent_Post): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creatorId !== undefined && (obj.creatorId = message.creatorId);
    message.caption !== undefined && (obj.caption = message.caption);
    message.address !== undefined && (obj.address = message.address);
    if (message.photos) {
      obj.photos = message.photos.map((e) => e);
    } else {
      obj.photos = [];
    }
    message.isCommentEnabled !== undefined && (obj.isCommentEnabled = message.isCommentEnabled);
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.drink !== undefined && (obj.drink = message.drink);
    message.drunkenness !== undefined && (obj.drunkenness = Math.round(message.drunkenness));
    message.commentsEnabled !== undefined && (obj.commentsEnabled = message.commentsEnabled);
    message.shareEnabled !== undefined && (obj.shareEnabled = message.shareEnabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PostEvent_Post>, I>>(object: I): PostEvent_Post {
    const message = createBasePostEvent_Post();
    message.id = object.id ?? "";
    message.creatorId = object.creatorId ?? "";
    message.caption = object.caption ?? "";
    message.address = object.address ?? "";
    message.photos = object.photos?.map((e) => e) || [];
    message.isCommentEnabled = object.isCommentEnabled ?? false;
    message.locationName = object.locationName ?? "";
    message.drink = object.drink ?? "";
    message.drunkenness = object.drunkenness ?? 0;
    message.commentsEnabled = object.commentsEnabled ?? false;
    message.shareEnabled = object.shareEnabled ?? false;
    return message;
  },
};

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
