/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.comment.v1";

export interface CreateLikeCommentRequest {
  commentId: string;
}

export interface CreateLikeCommentResponse {
}

export interface DeleteLikeCommentRequest {
  commentId: string;
}

export interface DeleteLikeCommentResponse {
}

export interface DeleteCommentRequest {
  postId: string;
  commentId: string;
}

export interface DeleteCommentResponse {
}

export interface ListCommentRequest {
  postId: string;
}

export interface ListRepliesRequest {
  commentId: string;
}

export interface ListRepliesResponse {
  items: CommentItem[];
}

export interface ListCommentResponse {
  items: CommentItem[];
}

export interface CommentItem {
  comment: Comment | undefined;
  userItem: UserItem | undefined;
  replyCount: number;
}

export interface Comment {
  id: string;
  text: string;
  createTime: number;
  userId: string;
  postId: string;
  replyToCommentId: string;
}

export interface CreateCommentRequest {
  postId: string;
  comment: string;
  replyToCommentId: string;
}

export interface CreateCommentResponse {
  item: CommentItem | undefined;
}

function createBaseCreateLikeCommentRequest(): CreateLikeCommentRequest {
  return { commentId: "" };
}

export const CreateLikeCommentRequest = {
  encode(message: CreateLikeCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commentId !== "") {
      writer.uint32(10).string(message.commentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateLikeCommentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateLikeCommentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateLikeCommentRequest {
    return { commentId: isSet(object.commentId) ? String(object.commentId) : "" };
  },

  toJSON(message: CreateLikeCommentRequest): unknown {
    const obj: any = {};
    message.commentId !== undefined && (obj.commentId = message.commentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateLikeCommentRequest>, I>>(base?: I): CreateLikeCommentRequest {
    return CreateLikeCommentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateLikeCommentRequest>, I>>(object: I): CreateLikeCommentRequest {
    const message = createBaseCreateLikeCommentRequest();
    message.commentId = object.commentId ?? "";
    return message;
  },
};

function createBaseCreateLikeCommentResponse(): CreateLikeCommentResponse {
  return {};
}

export const CreateLikeCommentResponse = {
  encode(_: CreateLikeCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateLikeCommentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateLikeCommentResponse();
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

  fromJSON(_: any): CreateLikeCommentResponse {
    return {};
  },

  toJSON(_: CreateLikeCommentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateLikeCommentResponse>, I>>(base?: I): CreateLikeCommentResponse {
    return CreateLikeCommentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateLikeCommentResponse>, I>>(_: I): CreateLikeCommentResponse {
    const message = createBaseCreateLikeCommentResponse();
    return message;
  },
};

function createBaseDeleteLikeCommentRequest(): DeleteLikeCommentRequest {
  return { commentId: "" };
}

export const DeleteLikeCommentRequest = {
  encode(message: DeleteLikeCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commentId !== "") {
      writer.uint32(10).string(message.commentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteLikeCommentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteLikeCommentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteLikeCommentRequest {
    return { commentId: isSet(object.commentId) ? String(object.commentId) : "" };
  },

  toJSON(message: DeleteLikeCommentRequest): unknown {
    const obj: any = {};
    message.commentId !== undefined && (obj.commentId = message.commentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteLikeCommentRequest>, I>>(base?: I): DeleteLikeCommentRequest {
    return DeleteLikeCommentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteLikeCommentRequest>, I>>(object: I): DeleteLikeCommentRequest {
    const message = createBaseDeleteLikeCommentRequest();
    message.commentId = object.commentId ?? "";
    return message;
  },
};

function createBaseDeleteLikeCommentResponse(): DeleteLikeCommentResponse {
  return {};
}

export const DeleteLikeCommentResponse = {
  encode(_: DeleteLikeCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteLikeCommentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteLikeCommentResponse();
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

  fromJSON(_: any): DeleteLikeCommentResponse {
    return {};
  },

  toJSON(_: DeleteLikeCommentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteLikeCommentResponse>, I>>(base?: I): DeleteLikeCommentResponse {
    return DeleteLikeCommentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteLikeCommentResponse>, I>>(_: I): DeleteLikeCommentResponse {
    const message = createBaseDeleteLikeCommentResponse();
    return message;
  },
};

function createBaseDeleteCommentRequest(): DeleteCommentRequest {
  return { postId: "", commentId: "" };
}

export const DeleteCommentRequest = {
  encode(message: DeleteCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    if (message.commentId !== "") {
      writer.uint32(18).string(message.commentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteCommentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteCommentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.postId = reader.string();
          break;
        case 2:
          message.commentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteCommentRequest {
    return {
      postId: isSet(object.postId) ? String(object.postId) : "",
      commentId: isSet(object.commentId) ? String(object.commentId) : "",
    };
  },

  toJSON(message: DeleteCommentRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    message.commentId !== undefined && (obj.commentId = message.commentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteCommentRequest>, I>>(base?: I): DeleteCommentRequest {
    return DeleteCommentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteCommentRequest>, I>>(object: I): DeleteCommentRequest {
    const message = createBaseDeleteCommentRequest();
    message.postId = object.postId ?? "";
    message.commentId = object.commentId ?? "";
    return message;
  },
};

function createBaseDeleteCommentResponse(): DeleteCommentResponse {
  return {};
}

export const DeleteCommentResponse = {
  encode(_: DeleteCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteCommentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteCommentResponse();
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

  fromJSON(_: any): DeleteCommentResponse {
    return {};
  },

  toJSON(_: DeleteCommentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteCommentResponse>, I>>(base?: I): DeleteCommentResponse {
    return DeleteCommentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteCommentResponse>, I>>(_: I): DeleteCommentResponse {
    const message = createBaseDeleteCommentResponse();
    return message;
  },
};

function createBaseListCommentRequest(): ListCommentRequest {
  return { postId: "" };
}

export const ListCommentRequest = {
  encode(message: ListCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListCommentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListCommentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.postId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListCommentRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: ListCommentRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListCommentRequest>, I>>(base?: I): ListCommentRequest {
    return ListCommentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListCommentRequest>, I>>(object: I): ListCommentRequest {
    const message = createBaseListCommentRequest();
    message.postId = object.postId ?? "";
    return message;
  },
};

function createBaseListRepliesRequest(): ListRepliesRequest {
  return { commentId: "" };
}

export const ListRepliesRequest = {
  encode(message: ListRepliesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commentId !== "") {
      writer.uint32(10).string(message.commentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRepliesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRepliesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListRepliesRequest {
    return { commentId: isSet(object.commentId) ? String(object.commentId) : "" };
  },

  toJSON(message: ListRepliesRequest): unknown {
    const obj: any = {};
    message.commentId !== undefined && (obj.commentId = message.commentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListRepliesRequest>, I>>(base?: I): ListRepliesRequest {
    return ListRepliesRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListRepliesRequest>, I>>(object: I): ListRepliesRequest {
    const message = createBaseListRepliesRequest();
    message.commentId = object.commentId ?? "";
    return message;
  },
};

function createBaseListRepliesResponse(): ListRepliesResponse {
  return { items: [] };
}

export const ListRepliesResponse = {
  encode(message: ListRepliesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      CommentItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRepliesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRepliesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(CommentItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListRepliesResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => CommentItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListRepliesResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? CommentItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListRepliesResponse>, I>>(base?: I): ListRepliesResponse {
    return ListRepliesResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListRepliesResponse>, I>>(object: I): ListRepliesResponse {
    const message = createBaseListRepliesResponse();
    message.items = object.items?.map((e) => CommentItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListCommentResponse(): ListCommentResponse {
  return { items: [] };
}

export const ListCommentResponse = {
  encode(message: ListCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      CommentItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListCommentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListCommentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(CommentItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListCommentResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => CommentItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListCommentResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? CommentItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListCommentResponse>, I>>(base?: I): ListCommentResponse {
    return ListCommentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListCommentResponse>, I>>(object: I): ListCommentResponse {
    const message = createBaseListCommentResponse();
    message.items = object.items?.map((e) => CommentItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCommentItem(): CommentItem {
  return { comment: undefined, userItem: undefined, replyCount: 0 };
}

export const CommentItem = {
  encode(message: CommentItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.comment !== undefined) {
      Comment.encode(message.comment, writer.uint32(10).fork()).ldelim();
    }
    if (message.userItem !== undefined) {
      UserItem.encode(message.userItem, writer.uint32(18).fork()).ldelim();
    }
    if (message.replyCount !== 0) {
      writer.uint32(24).int64(message.replyCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CommentItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCommentItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.comment = Comment.decode(reader, reader.uint32());
          break;
        case 2:
          message.userItem = UserItem.decode(reader, reader.uint32());
          break;
        case 3:
          message.replyCount = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CommentItem {
    return {
      comment: isSet(object.comment) ? Comment.fromJSON(object.comment) : undefined,
      userItem: isSet(object.userItem) ? UserItem.fromJSON(object.userItem) : undefined,
      replyCount: isSet(object.replyCount) ? Number(object.replyCount) : 0,
    };
  },

  toJSON(message: CommentItem): unknown {
    const obj: any = {};
    message.comment !== undefined && (obj.comment = message.comment ? Comment.toJSON(message.comment) : undefined);
    message.userItem !== undefined && (obj.userItem = message.userItem ? UserItem.toJSON(message.userItem) : undefined);
    message.replyCount !== undefined && (obj.replyCount = Math.round(message.replyCount));
    return obj;
  },

  create<I extends Exact<DeepPartial<CommentItem>, I>>(base?: I): CommentItem {
    return CommentItem.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CommentItem>, I>>(object: I): CommentItem {
    const message = createBaseCommentItem();
    message.comment = (object.comment !== undefined && object.comment !== null)
      ? Comment.fromPartial(object.comment)
      : undefined;
    message.userItem = (object.userItem !== undefined && object.userItem !== null)
      ? UserItem.fromPartial(object.userItem)
      : undefined;
    message.replyCount = object.replyCount ?? 0;
    return message;
  },
};

function createBaseComment(): Comment {
  return { id: "", text: "", createTime: 0, userId: "", postId: "", replyToCommentId: "" };
}

export const Comment = {
  encode(message: Comment, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    if (message.createTime !== 0) {
      writer.uint32(40).int64(message.createTime);
    }
    if (message.userId !== "") {
      writer.uint32(26).string(message.userId);
    }
    if (message.postId !== "") {
      writer.uint32(34).string(message.postId);
    }
    if (message.replyToCommentId !== "") {
      writer.uint32(50).string(message.replyToCommentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Comment {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseComment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.text = reader.string();
          break;
        case 5:
          message.createTime = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.userId = reader.string();
          break;
        case 4:
          message.postId = reader.string();
          break;
        case 6:
          message.replyToCommentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Comment {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      text: isSet(object.text) ? String(object.text) : "",
      createTime: isSet(object.createTime) ? Number(object.createTime) : 0,
      userId: isSet(object.userId) ? String(object.userId) : "",
      postId: isSet(object.postId) ? String(object.postId) : "",
      replyToCommentId: isSet(object.replyToCommentId) ? String(object.replyToCommentId) : "",
    };
  },

  toJSON(message: Comment): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.text !== undefined && (obj.text = message.text);
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    message.userId !== undefined && (obj.userId = message.userId);
    message.postId !== undefined && (obj.postId = message.postId);
    message.replyToCommentId !== undefined && (obj.replyToCommentId = message.replyToCommentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<Comment>, I>>(base?: I): Comment {
    return Comment.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Comment>, I>>(object: I): Comment {
    const message = createBaseComment();
    message.id = object.id ?? "";
    message.text = object.text ?? "";
    message.createTime = object.createTime ?? 0;
    message.userId = object.userId ?? "";
    message.postId = object.postId ?? "";
    message.replyToCommentId = object.replyToCommentId ?? "";
    return message;
  },
};

function createBaseCreateCommentRequest(): CreateCommentRequest {
  return { postId: "", comment: "", replyToCommentId: "" };
}

export const CreateCommentRequest = {
  encode(message: CreateCommentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    if (message.comment !== "") {
      writer.uint32(18).string(message.comment);
    }
    if (message.replyToCommentId !== "") {
      writer.uint32(26).string(message.replyToCommentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCommentRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCommentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.postId = reader.string();
          break;
        case 2:
          message.comment = reader.string();
          break;
        case 3:
          message.replyToCommentId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateCommentRequest {
    return {
      postId: isSet(object.postId) ? String(object.postId) : "",
      comment: isSet(object.comment) ? String(object.comment) : "",
      replyToCommentId: isSet(object.replyToCommentId) ? String(object.replyToCommentId) : "",
    };
  },

  toJSON(message: CreateCommentRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    message.comment !== undefined && (obj.comment = message.comment);
    message.replyToCommentId !== undefined && (obj.replyToCommentId = message.replyToCommentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCommentRequest>, I>>(base?: I): CreateCommentRequest {
    return CreateCommentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateCommentRequest>, I>>(object: I): CreateCommentRequest {
    const message = createBaseCreateCommentRequest();
    message.postId = object.postId ?? "";
    message.comment = object.comment ?? "";
    message.replyToCommentId = object.replyToCommentId ?? "";
    return message;
  },
};

function createBaseCreateCommentResponse(): CreateCommentResponse {
  return { item: undefined };
}

export const CreateCommentResponse = {
  encode(message: CreateCommentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.item !== undefined) {
      CommentItem.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateCommentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateCommentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.item = CommentItem.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateCommentResponse {
    return { item: isSet(object.item) ? CommentItem.fromJSON(object.item) : undefined };
  },

  toJSON(message: CreateCommentResponse): unknown {
    const obj: any = {};
    message.item !== undefined && (obj.item = message.item ? CommentItem.toJSON(message.item) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateCommentResponse>, I>>(base?: I): CreateCommentResponse {
    return CreateCommentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateCommentResponse>, I>>(object: I): CreateCommentResponse {
    const message = createBaseCreateCommentResponse();
    message.item = (object.item !== undefined && object.item !== null)
      ? CommentItem.fromPartial(object.item)
      : undefined;
    return message;
  },
};

export interface CommentService {
  CreateComment(request: CreateCommentRequest): Promise<CreateCommentResponse>;
  ListComment(request: ListCommentRequest): Promise<ListCommentResponse>;
  ListReplies(request: ListRepliesRequest): Promise<ListRepliesResponse>;
  DeleteComment(request: DeleteCommentRequest): Promise<DeleteCommentResponse>;
  CreateLikeComment(request: CreateLikeCommentRequest): Promise<CreateLikeCommentResponse>;
  DeleteLikeComment(request: DeleteLikeCommentRequest): Promise<DeleteLikeCommentResponse>;
}

export class CommentServiceClientImpl implements CommentService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.comment.v1.CommentService";
    this.rpc = rpc;
    this.CreateComment = this.CreateComment.bind(this);
    this.ListComment = this.ListComment.bind(this);
    this.ListReplies = this.ListReplies.bind(this);
    this.DeleteComment = this.DeleteComment.bind(this);
    this.CreateLikeComment = this.CreateLikeComment.bind(this);
    this.DeleteLikeComment = this.DeleteLikeComment.bind(this);
  }
  CreateComment(request: CreateCommentRequest): Promise<CreateCommentResponse> {
    const data = CreateCommentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateComment", data);
    return promise.then((data) => CreateCommentResponse.decode(new _m0.Reader(data)));
  }

  ListComment(request: ListCommentRequest): Promise<ListCommentResponse> {
    const data = ListCommentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListComment", data);
    return promise.then((data) => ListCommentResponse.decode(new _m0.Reader(data)));
  }

  ListReplies(request: ListRepliesRequest): Promise<ListRepliesResponse> {
    const data = ListRepliesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListReplies", data);
    return promise.then((data) => ListRepliesResponse.decode(new _m0.Reader(data)));
  }

  DeleteComment(request: DeleteCommentRequest): Promise<DeleteCommentResponse> {
    const data = DeleteCommentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteComment", data);
    return promise.then((data) => DeleteCommentResponse.decode(new _m0.Reader(data)));
  }

  CreateLikeComment(request: CreateLikeCommentRequest): Promise<CreateLikeCommentResponse> {
    const data = CreateLikeCommentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateLikeComment", data);
    return promise.then((data) => CreateLikeCommentResponse.decode(new _m0.Reader(data)));
  }

  DeleteLikeComment(request: DeleteLikeCommentRequest): Promise<DeleteLikeCommentResponse> {
    const data = DeleteLikeCommentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteLikeComment", data);
    return promise.then((data) => DeleteLikeCommentResponse.decode(new _m0.Reader(data)));
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
