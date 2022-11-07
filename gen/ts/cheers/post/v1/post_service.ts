/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../../google/protobuf/empty";
import { Post } from "../../type/post/post";
import { User } from "../../type/user/user";

export const protobufPackage = "cheers.post.v1";

export interface CreatePostRequest {
  post: Post | undefined;
}

export interface GetPostRequest {
  id: string;
}

export interface UpdatePostRequest {
  post: Post | undefined;
}

export interface DeletePostRequest {
  id: string;
}

export interface FeedPostRequest {
  parent: string;
  pageSize: number;
  page: number;
}

export interface FeedPostResponse {
  posts: PostResponse[];
  nextPageToken: string;
}

export interface LikePostRequest {
  id: string;
}

export interface LikePostResponse {
  success: boolean;
}

export interface UnlikePostRequest {
  id: string;
}

export interface UnlikePostResponse {
  success: boolean;
}

export interface SavePostRequest {
  id: string;
}

export interface SavePostResponse {
  success: boolean;
}

export interface UnsavePostRequest {
  id: string;
}

export interface UnsavePostResponse {
  success: boolean;
}

export interface PostResponse {
  post: Post | undefined;
  user: User | undefined;
  likeCount: number;
  commentCount: number;
  hasLiked: boolean;
  commentsEnabled: boolean;
  reshareEnabled: boolean;
  mediaCount: number;
}

function createBaseCreatePostRequest(): CreatePostRequest {
  return { post: undefined };
}

export const CreatePostRequest = {
  encode(message: CreatePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.post = Post.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreatePostRequest {
    return { post: isSet(object.post) ? Post.fromJSON(object.post) : undefined };
  },

  toJSON(message: CreatePostRequest): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatePostRequest>, I>>(object: I): CreatePostRequest {
    const message = createBaseCreatePostRequest();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    return message;
  },
};

function createBaseGetPostRequest(): GetPostRequest {
  return { id: "" };
}

export const GetPostRequest = {
  encode(message: GetPostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetPostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: GetPostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetPostRequest>, I>>(object: I): GetPostRequest {
    const message = createBaseGetPostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUpdatePostRequest(): UpdatePostRequest {
  return { post: undefined };
}

export const UpdatePostRequest = {
  encode(message: UpdatePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdatePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.post = Post.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdatePostRequest {
    return { post: isSet(object.post) ? Post.fromJSON(object.post) : undefined };
  },

  toJSON(message: UpdatePostRequest): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdatePostRequest>, I>>(object: I): UpdatePostRequest {
    const message = createBaseUpdatePostRequest();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    return message;
  },
};

function createBaseDeletePostRequest(): DeletePostRequest {
  return { id: "" };
}

export const DeletePostRequest = {
  encode(message: DeletePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletePostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: DeletePostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePostRequest>, I>>(object: I): DeletePostRequest {
    const message = createBaseDeletePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseFeedPostRequest(): FeedPostRequest {
  return { parent: "", pageSize: 0, page: 0 };
}

export const FeedPostRequest = {
  encode(message: FeedPostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.parent !== "") {
      writer.uint32(10).string(message.parent);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.parent = reader.string();
          break;
        case 2:
          message.pageSize = reader.int32();
          break;
        case 3:
          message.page = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeedPostRequest {
    return {
      parent: isSet(object.parent) ? String(object.parent) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: FeedPostRequest): unknown {
    const obj: any = {};
    message.parent !== undefined && (obj.parent = message.parent);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedPostRequest>, I>>(object: I): FeedPostRequest {
    const message = createBaseFeedPostRequest();
    message.parent = object.parent ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseFeedPostResponse(): FeedPostResponse {
  return { posts: [], nextPageToken: "" };
}

export const FeedPostResponse = {
  encode(message: FeedPostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.posts) {
      PostResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.posts.push(PostResponse.decode(reader, reader.uint32()));
          break;
        case 2:
          message.nextPageToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeedPostResponse {
    return {
      posts: Array.isArray(object?.posts) ? object.posts.map((e: any) => PostResponse.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: FeedPostResponse): unknown {
    const obj: any = {};
    if (message.posts) {
      obj.posts = message.posts.map((e) => e ? PostResponse.toJSON(e) : undefined);
    } else {
      obj.posts = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeedPostResponse>, I>>(object: I): FeedPostResponse {
    const message = createBaseFeedPostResponse();
    message.posts = object.posts?.map((e) => PostResponse.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseLikePostRequest(): LikePostRequest {
  return { id: "" };
}

export const LikePostRequest = {
  encode(message: LikePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikePostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: LikePostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikePostRequest>, I>>(object: I): LikePostRequest {
    const message = createBaseLikePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseLikePostResponse(): LikePostResponse {
  return { success: false };
}

export const LikePostResponse = {
  encode(message: LikePostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikePostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikePostResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: LikePostResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikePostResponse>, I>>(object: I): LikePostResponse {
    const message = createBaseLikePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnlikePostRequest(): UnlikePostRequest {
  return { id: "" };
}

export const UnlikePostRequest = {
  encode(message: UnlikePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnlikePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnlikePostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: UnlikePostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnlikePostRequest>, I>>(object: I): UnlikePostRequest {
    const message = createBaseUnlikePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUnlikePostResponse(): UnlikePostResponse {
  return { success: false };
}

export const UnlikePostResponse = {
  encode(message: UnlikePostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnlikePostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnlikePostResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: UnlikePostResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnlikePostResponse>, I>>(object: I): UnlikePostResponse {
    const message = createBaseUnlikePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseSavePostRequest(): SavePostRequest {
  return { id: "" };
}

export const SavePostRequest = {
  encode(message: SavePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SavePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSavePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SavePostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: SavePostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SavePostRequest>, I>>(object: I): SavePostRequest {
    const message = createBaseSavePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseSavePostResponse(): SavePostResponse {
  return { success: false };
}

export const SavePostResponse = {
  encode(message: SavePostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SavePostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSavePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SavePostResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: SavePostResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SavePostResponse>, I>>(object: I): SavePostResponse {
    const message = createBaseSavePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnsavePostRequest(): UnsavePostRequest {
  return { id: "" };
}

export const UnsavePostRequest = {
  encode(message: UnsavePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnsavePostRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsavePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnsavePostRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: UnsavePostRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnsavePostRequest>, I>>(object: I): UnsavePostRequest {
    const message = createBaseUnsavePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUnsavePostResponse(): UnsavePostResponse {
  return { success: false };
}

export const UnsavePostResponse = {
  encode(message: UnsavePostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnsavePostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsavePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UnsavePostResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: UnsavePostResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnsavePostResponse>, I>>(object: I): UnsavePostResponse {
    const message = createBaseUnsavePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBasePostResponse(): PostResponse {
  return {
    post: undefined,
    user: undefined,
    likeCount: 0,
    commentCount: 0,
    hasLiked: false,
    commentsEnabled: false,
    reshareEnabled: false,
    mediaCount: 0,
  };
}

export const PostResponse = {
  encode(message: PostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    if (message.likeCount !== 0) {
      writer.uint32(24).int64(message.likeCount);
    }
    if (message.commentCount !== 0) {
      writer.uint32(32).int64(message.commentCount);
    }
    if (message.hasLiked === true) {
      writer.uint32(40).bool(message.hasLiked);
    }
    if (message.commentsEnabled === true) {
      writer.uint32(48).bool(message.commentsEnabled);
    }
    if (message.reshareEnabled === true) {
      writer.uint32(64).bool(message.reshareEnabled);
    }
    if (message.mediaCount !== 0) {
      writer.uint32(56).int64(message.mediaCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PostResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.post = Post.decode(reader, reader.uint32());
          break;
        case 2:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 3:
          message.likeCount = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.commentCount = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.hasLiked = reader.bool();
          break;
        case 6:
          message.commentsEnabled = reader.bool();
          break;
        case 8:
          message.reshareEnabled = reader.bool();
          break;
        case 7:
          message.mediaCount = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PostResponse {
    return {
      post: isSet(object.post) ? Post.fromJSON(object.post) : undefined,
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      likeCount: isSet(object.likeCount) ? Number(object.likeCount) : 0,
      commentCount: isSet(object.commentCount) ? Number(object.commentCount) : 0,
      hasLiked: isSet(object.hasLiked) ? Boolean(object.hasLiked) : false,
      commentsEnabled: isSet(object.commentsEnabled) ? Boolean(object.commentsEnabled) : false,
      reshareEnabled: isSet(object.reshareEnabled) ? Boolean(object.reshareEnabled) : false,
      mediaCount: isSet(object.mediaCount) ? Number(object.mediaCount) : 0,
    };
  },

  toJSON(message: PostResponse): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.likeCount !== undefined && (obj.likeCount = Math.round(message.likeCount));
    message.commentCount !== undefined && (obj.commentCount = Math.round(message.commentCount));
    message.hasLiked !== undefined && (obj.hasLiked = message.hasLiked);
    message.commentsEnabled !== undefined && (obj.commentsEnabled = message.commentsEnabled);
    message.reshareEnabled !== undefined && (obj.reshareEnabled = message.reshareEnabled);
    message.mediaCount !== undefined && (obj.mediaCount = Math.round(message.mediaCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PostResponse>, I>>(object: I): PostResponse {
    const message = createBasePostResponse();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.likeCount = object.likeCount ?? 0;
    message.commentCount = object.commentCount ?? 0;
    message.hasLiked = object.hasLiked ?? false;
    message.commentsEnabled = object.commentsEnabled ?? false;
    message.reshareEnabled = object.reshareEnabled ?? false;
    message.mediaCount = object.mediaCount ?? 0;
    return message;
  },
};

export interface PostService {
  /** Create a new post */
  CreatePost(request: CreatePostRequest): Promise<PostResponse>;
  GetPost(request: GetPostRequest): Promise<PostResponse>;
  UpdatePost(request: UpdatePostRequest): Promise<PostResponse>;
  DeletePost(request: DeletePostRequest): Promise<Empty>;
  FeedPost(request: FeedPostRequest): Promise<FeedPostResponse>;
  LikePost(request: LikePostRequest): Promise<LikePostResponse>;
  UnlikePost(request: UnlikePostRequest): Promise<UnlikePostResponse>;
  SavePost(request: SavePostRequest): Promise<SavePostResponse>;
  UnsavePost(request: UnsavePostRequest): Promise<UnsavePostResponse>;
}

export class PostServiceClientImpl implements PostService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.post.v1.PostService";
    this.rpc = rpc;
    this.CreatePost = this.CreatePost.bind(this);
    this.GetPost = this.GetPost.bind(this);
    this.UpdatePost = this.UpdatePost.bind(this);
    this.DeletePost = this.DeletePost.bind(this);
    this.FeedPost = this.FeedPost.bind(this);
    this.LikePost = this.LikePost.bind(this);
    this.UnlikePost = this.UnlikePost.bind(this);
    this.SavePost = this.SavePost.bind(this);
    this.UnsavePost = this.UnsavePost.bind(this);
  }
  CreatePost(request: CreatePostRequest): Promise<PostResponse> {
    const data = CreatePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreatePost", data);
    return promise.then((data) => PostResponse.decode(new _m0.Reader(data)));
  }

  GetPost(request: GetPostRequest): Promise<PostResponse> {
    const data = GetPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPost", data);
    return promise.then((data) => PostResponse.decode(new _m0.Reader(data)));
  }

  UpdatePost(request: UpdatePostRequest): Promise<PostResponse> {
    const data = UpdatePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdatePost", data);
    return promise.then((data) => PostResponse.decode(new _m0.Reader(data)));
  }

  DeletePost(request: DeletePostRequest): Promise<Empty> {
    const data = DeletePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeletePost", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  FeedPost(request: FeedPostRequest): Promise<FeedPostResponse> {
    const data = FeedPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeedPost", data);
    return promise.then((data) => FeedPostResponse.decode(new _m0.Reader(data)));
  }

  LikePost(request: LikePostRequest): Promise<LikePostResponse> {
    const data = LikePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LikePost", data);
    return promise.then((data) => LikePostResponse.decode(new _m0.Reader(data)));
  }

  UnlikePost(request: UnlikePostRequest): Promise<UnlikePostResponse> {
    const data = UnlikePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnlikePost", data);
    return promise.then((data) => UnlikePostResponse.decode(new _m0.Reader(data)));
  }

  SavePost(request: SavePostRequest): Promise<SavePostResponse> {
    const data = SavePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SavePost", data);
    return promise.then((data) => SavePostResponse.decode(new _m0.Reader(data)));
  }

  UnsavePost(request: UnsavePostRequest): Promise<UnsavePostResponse> {
    const data = UnsavePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnsavePost", data);
    return promise.then((data) => UnsavePostResponse.decode(new _m0.Reader(data)));
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