/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.type";

export enum UserEventType {
  FOLLOW = 0,
  STORY_LIKE = 1,
  POST_LIKE = 2,
  COMMENT = 3,
  MENTION = 4,
  CREATE_POST = 5,
  CREATE_EVENT = 6,
  CREATE_STORY = 7,
  UNRECOGNIZED = -1,
}

export function userEventTypeFromJSON(object: any): UserEventType {
  switch (object) {
    case 0:
    case "FOLLOW":
      return UserEventType.FOLLOW;
    case 1:
    case "STORY_LIKE":
      return UserEventType.STORY_LIKE;
    case 2:
    case "POST_LIKE":
      return UserEventType.POST_LIKE;
    case 3:
    case "COMMENT":
      return UserEventType.COMMENT;
    case 4:
    case "MENTION":
      return UserEventType.MENTION;
    case 5:
    case "CREATE_POST":
      return UserEventType.CREATE_POST;
    case 6:
    case "CREATE_EVENT":
      return UserEventType.CREATE_EVENT;
    case 7:
    case "CREATE_STORY":
      return UserEventType.CREATE_STORY;
    case -1:
    case "UNRECOGNIZED":
    default:
      return UserEventType.UNRECOGNIZED;
  }
}

export function userEventTypeToJSON(object: UserEventType): string {
  switch (object) {
    case UserEventType.FOLLOW:
      return "FOLLOW";
    case UserEventType.STORY_LIKE:
      return "STORY_LIKE";
    case UserEventType.POST_LIKE:
      return "POST_LIKE";
    case UserEventType.COMMENT:
      return "COMMENT";
    case UserEventType.MENTION:
      return "MENTION";
    case UserEventType.CREATE_POST:
      return "CREATE_POST";
    case UserEventType.CREATE_EVENT:
      return "CREATE_EVENT";
    case UserEventType.CREATE_STORY:
      return "CREATE_STORY";
    case UserEventType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface UserEvent {
  type: UserEventType;
  userId: string;
  time: number;
  otherUserId: string;
  postId: string;
  eventId: string;
  storyId: string;
}

function createBaseUserEvent(): UserEvent {
  return { type: 0, userId: "", time: 0, otherUserId: "", postId: "", eventId: "", storyId: "" };
}

export const UserEvent = {
  encode(message: UserEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.time !== 0) {
      writer.uint32(24).int64(message.time);
    }
    if (message.otherUserId !== "") {
      writer.uint32(34).string(message.otherUserId);
    }
    if (message.postId !== "") {
      writer.uint32(42).string(message.postId);
    }
    if (message.eventId !== "") {
      writer.uint32(50).string(message.eventId);
    }
    if (message.storyId !== "") {
      writer.uint32(58).string(message.storyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.time = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.otherUserId = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.postId = reader.string();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.eventId = reader.string();
          continue;
        case 7:
          if (tag != 58) {
            break;
          }

          message.storyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserEvent {
    return {
      type: isSet(object.type) ? userEventTypeFromJSON(object.type) : 0,
      userId: isSet(object.userId) ? String(object.userId) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
      otherUserId: isSet(object.otherUserId) ? String(object.otherUserId) : "",
      postId: isSet(object.postId) ? String(object.postId) : "",
      eventId: isSet(object.eventId) ? String(object.eventId) : "",
      storyId: isSet(object.storyId) ? String(object.storyId) : "",
    };
  },

  toJSON(message: UserEvent): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = userEventTypeToJSON(message.type));
    message.userId !== undefined && (obj.userId = message.userId);
    message.time !== undefined && (obj.time = Math.round(message.time));
    message.otherUserId !== undefined && (obj.otherUserId = message.otherUserId);
    message.postId !== undefined && (obj.postId = message.postId);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.storyId !== undefined && (obj.storyId = message.storyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserEvent>, I>>(base?: I): UserEvent {
    return UserEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserEvent>, I>>(object: I): UserEvent {
    const message = createBaseUserEvent();
    message.type = object.type ?? 0;
    message.userId = object.userId ?? "";
    message.time = object.time ?? 0;
    message.otherUserId = object.otherUserId ?? "";
    message.postId = object.postId ?? "";
    message.eventId = object.eventId ?? "";
    message.storyId = object.storyId ?? "";
    return message;
  },
};

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
