/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";
import { Comment } from "./comment_service";

export const protobufPackage = "cheers.comment.v1";

export interface CommentEvent {
  created?: CreatedComment | undefined;
  deleted?: DeletedComment | undefined;
}

export interface CreatedComment {
  comment: Comment | undefined;
  user: UserItem | undefined;
}

export interface DeletedComment {
  comment: Comment | undefined;
  user: UserItem | undefined;
}

function createBaseCommentEvent(): CommentEvent {
  return { created: undefined, deleted: undefined };
}

export const CommentEvent = {
  encode(message: CommentEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.created !== undefined) {
      CreatedComment.encode(message.created, writer.uint32(10).fork()).ldelim();
    }
    if (message.deleted !== undefined) {
      DeletedComment.encode(message.deleted, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CommentEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommentEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.created = CreatedComment.decode(reader, reader.uint32());
          break;
        case 2:
          message.deleted = DeletedComment.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CommentEvent {
    return {
      created: isSet(object.created) ? CreatedComment.fromJSON(object.created) : undefined,
      deleted: isSet(object.deleted) ? DeletedComment.fromJSON(object.deleted) : undefined,
    };
  },

  toJSON(message: CommentEvent): unknown {
    const obj: any = {};
    message.created !== undefined &&
      (obj.created = message.created ? CreatedComment.toJSON(message.created) : undefined);
    message.deleted !== undefined &&
      (obj.deleted = message.deleted ? DeletedComment.toJSON(message.deleted) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CommentEvent>, I>>(object: I): CommentEvent {
    const message = createBaseCommentEvent();
    message.created = (object.created !== undefined && object.created !== null)
      ? CreatedComment.fromPartial(object.created)
      : undefined;
    message.deleted = (object.deleted !== undefined && object.deleted !== null)
      ? DeletedComment.fromPartial(object.deleted)
      : undefined;
    return message;
  },
};

function createBaseCreatedComment(): CreatedComment {
  return { comment: undefined, user: undefined };
}

export const CreatedComment = {
  encode(message: CreatedComment, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.comment !== undefined) {
      Comment.encode(message.comment, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedComment {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedComment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.comment = Comment.decode(reader, reader.uint32());
          break;
        case 2:
          message.user = UserItem.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatedComment {
    return {
      comment: isSet(object.comment) ? Comment.fromJSON(object.comment) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreatedComment): unknown {
    const obj: any = {};
    message.comment !== undefined && (obj.comment = message.comment ? Comment.toJSON(message.comment) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatedComment>, I>>(object: I): CreatedComment {
    const message = createBaseCreatedComment();
    message.comment = (object.comment !== undefined && object.comment !== null)
      ? Comment.fromPartial(object.comment)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseDeletedComment(): DeletedComment {
  return { comment: undefined, user: undefined };
}

export const DeletedComment = {
  encode(message: DeletedComment, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.comment !== undefined) {
      Comment.encode(message.comment, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedComment {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedComment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.comment = Comment.decode(reader, reader.uint32());
          break;
        case 2:
          message.user = UserItem.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletedComment {
    return {
      comment: isSet(object.comment) ? Comment.fromJSON(object.comment) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: DeletedComment): unknown {
    const obj: any = {};
    message.comment !== undefined && (obj.comment = message.comment ? Comment.toJSON(message.comment) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletedComment>, I>>(object: I): DeletedComment {
    const message = createBaseDeletedComment();
    message.comment = (object.comment !== undefined && object.comment !== null)
      ? Comment.fromPartial(object.comment)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
