/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Post } from "../../type/post/post";

export const protobufPackage = "cheers.post.v1";

export interface PostEvent {
  userId: string;
  post: Post | undefined;
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

function createBasePostEvent(): PostEvent {
  return { userId: "", post: undefined, type: 0 };
}

export const PostEvent = {
  encode(message: PostEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(18).fork()).ldelim();
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
          message.post = Post.decode(reader, reader.uint32());
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
      post: isSet(object.post) ? Post.fromJSON(object.post) : undefined,
      type: isSet(object.type) ? postEvent_EventTypeFromJSON(object.type) : 0,
    };
  },

  toJSON(message: PostEvent): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.type !== undefined && (obj.type = postEvent_EventTypeToJSON(message.type));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PostEvent>, I>>(object: I): PostEvent {
    const message = createBasePostEvent();
    message.userId = object.userId ?? "";
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.type = object.type ?? 0;
    return message;
  },
};

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
