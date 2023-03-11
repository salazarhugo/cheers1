/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";
import { Comment } from "./comment_service";

export const protobufPackage = "cheers.comment.v1";

export interface CommentEvent {
  created?: CreatedComment | undefined;
  deleted?: DeletedComment | undefined;
  liked?: LikedComment | undefined;
}

export interface LikedComment {
  commentId: string;
  userId: string;
}

export interface CreatedComment {
  comment: Comment | undefined;
  user: UserItem | undefined;
  mentions: string[];
}

export interface DeletedComment {
  comment: Comment | undefined;
  user: UserItem | undefined;
}

function createBaseCommentEvent(): CommentEvent {
  return { created: undefined, deleted: undefined, liked: undefined };
}

export const CommentEvent = {
  encode(message: CommentEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.created !== undefined) {
      CreatedComment.encode(message.created, writer.uint32(10).fork()).ldelim();
    }
    if (message.deleted !== undefined) {
      DeletedComment.encode(message.deleted, writer.uint32(18).fork()).ldelim();
    }
    if (message.liked !== undefined) {
      LikedComment.encode(message.liked, writer.uint32(26).fork()).ldelim();
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
        case 3:
          message.liked = LikedComment.decode(reader, reader.uint32());
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
      liked: isSet(object.liked) ? LikedComment.fromJSON(object.liked) : undefined,
    };
  },

  toJSON(message: CommentEvent): unknown {
    const obj: any = {};
    message.created !== undefined &&
      (obj.created = message.created ? CreatedComment.toJSON(message.created) : undefined);
    message.deleted !== undefined &&
      (obj.deleted = message.deleted ? DeletedComment.toJSON(message.deleted) : undefined);
    message.liked !== undefined && (obj.liked = message.liked ? LikedComment.toJSON(message.liked) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CommentEvent>, I>>(base?: I): CommentEvent {
    return CommentEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommentEvent>, I>>(object: I): CommentEvent {
    const message = createBaseCommentEvent();
    message.created = (object.created !== undefined && object.created !== null)
      ? CreatedComment.fromPartial(object.created)
      : undefined;
    message.deleted = (object.deleted !== undefined && object.deleted !== null)
      ? DeletedComment.fromPartial(object.deleted)
      : undefined;
    message.liked = (object.liked !== undefined && object.liked !== null)
      ? LikedComment.fromPartial(object.liked)
      : undefined;
    return message;
  },
};

function createBaseLikedComment(): LikedComment {
  return { commentId: "", userId: "" };
}

export const LikedComment = {
  encode(message: LikedComment, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commentId !== "") {
      writer.uint32(10).string(message.commentId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikedComment {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikedComment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commentId = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikedComment {
    return {
      commentId: isSet(object.commentId) ? String(object.commentId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: LikedComment): unknown {
    const obj: any = {};
    message.commentId !== undefined && (obj.commentId = message.commentId);
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<LikedComment>, I>>(base?: I): LikedComment {
    return LikedComment.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LikedComment>, I>>(object: I): LikedComment {
    const message = createBaseLikedComment();
    message.commentId = object.commentId ?? "";
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseCreatedComment(): CreatedComment {
  return { comment: undefined, user: undefined, mentions: [] };
}

export const CreatedComment = {
  encode(message: CreatedComment, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.comment !== undefined) {
      Comment.encode(message.comment, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.mentions) {
      writer.uint32(26).string(v!);
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
        case 3:
          message.mentions.push(reader.string());
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
      mentions: Array.isArray(object?.mentions) ? object.mentions.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: CreatedComment): unknown {
    const obj: any = {};
    message.comment !== undefined && (obj.comment = message.comment ? Comment.toJSON(message.comment) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    if (message.mentions) {
      obj.mentions = message.mentions.map((e) => e);
    } else {
      obj.mentions = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatedComment>, I>>(base?: I): CreatedComment {
    return CreatedComment.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatedComment>, I>>(object: I): CreatedComment {
    const message = createBaseCreatedComment();
    message.comment = (object.comment !== undefined && object.comment !== null)
      ? Comment.fromPartial(object.comment)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    message.mentions = object.mentions?.map((e) => e) || [];
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

  create<I extends Exact<DeepPartial<DeletedComment>, I>>(base?: I): DeletedComment {
    return DeletedComment.fromPartial(base ?? {});
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
