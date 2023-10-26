/* eslint-disable */
// @ts-ignore
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../../google/protobuf/empty";
import { Post } from "../../type/post/post";
import { User } from "../../type/user/user";

export const protobufPackage = "cheers.post.v1";

export interface ListMapPostRequest {
  parent: string;
  pageSize: number;
  page: number;
}

export interface ListMapPostResponse {
  posts: PostResponse[];
}

export interface CreatePostRequest {
  post: Post | undefined;
  sendNotificationToFriends: boolean;
}

export interface GetPostRequest {
  postId: string;
}

export interface GetPostResponse {
  post: Post | undefined;
}

export interface GetPostItemRequest {
  postId: string;
}

export interface UpdatePostRequest {
  post: Post | undefined;
}

export interface DeletePostRequest {
  id: string;
}

export interface ListPostRequest {
  username: string;
  pageSize: number;
  page: number;
}

export interface ListPostResponse {
  posts: PostResponse[];
  nextPageToken: string;
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
  postId: string;
}

export interface LikePostResponse {
  success: boolean;
}

export interface UnlikePostRequest {
  postId: string;
}

export interface UnlikePostResponse {
  success: boolean;
}

export interface SavePostRequest {
  postId: string;
}

export interface SavePostResponse {
  success: boolean;
}

export interface UnsavePostRequest {
  postId: string;
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
  isCreator: boolean;
}

function createBaseListMapPostRequest(): ListMapPostRequest {
  return { parent: "", pageSize: 0, page: 0 };
}

export const ListMapPostRequest = {
  encode(message: ListMapPostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMapPostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMapPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.parent = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.page = reader.int32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMapPostRequest {
    return {
      parent: isSet(object.parent) ? String(object.parent) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListMapPostRequest): unknown {
    const obj: any = {};
    message.parent !== undefined && (obj.parent = message.parent);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  create<I extends Exact<DeepPartial<ListMapPostRequest>, I>>(base?: I): ListMapPostRequest {
    return ListMapPostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListMapPostRequest>, I>>(object: I): ListMapPostRequest {
    const message = createBaseListMapPostRequest();
    message.parent = object.parent ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListMapPostResponse(): ListMapPostResponse {
  return { posts: [] };
}

export const ListMapPostResponse = {
  encode(message: ListMapPostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.posts) {
      PostResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMapPostResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMapPostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.posts.push(PostResponse.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListMapPostResponse {
    return { posts: Array.isArray(object?.posts) ? object.posts.map((e: any) => PostResponse.fromJSON(e)) : [] };
  },

  toJSON(message: ListMapPostResponse): unknown {
    const obj: any = {};
    if (message.posts) {
      obj.posts = message.posts.map((e) => e ? PostResponse.toJSON(e) : undefined);
    } else {
      obj.posts = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListMapPostResponse>, I>>(base?: I): ListMapPostResponse {
    return ListMapPostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListMapPostResponse>, I>>(object: I): ListMapPostResponse {
    const message = createBaseListMapPostResponse();
    message.posts = object.posts?.map((e) => PostResponse.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreatePostRequest(): CreatePostRequest {
  return { post: undefined, sendNotificationToFriends: false };
}

export const CreatePostRequest = {
  encode(message: CreatePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    if (message.sendNotificationToFriends === true) {
      writer.uint32(16).bool(message.sendNotificationToFriends);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.post = Post.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.sendNotificationToFriends = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePostRequest {
    return {
      post: isSet(object.post) ? Post.fromJSON(object.post) : undefined,
      sendNotificationToFriends: isSet(object.sendNotificationToFriends)
        ? Boolean(object.sendNotificationToFriends)
        : false,
    };
  },

  toJSON(message: CreatePostRequest): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.sendNotificationToFriends !== undefined &&
      (obj.sendNotificationToFriends = message.sendNotificationToFriends);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePostRequest>, I>>(base?: I): CreatePostRequest {
    return CreatePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePostRequest>, I>>(object: I): CreatePostRequest {
    const message = createBaseCreatePostRequest();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.sendNotificationToFriends = object.sendNotificationToFriends ?? false;
    return message;
  },
};

function createBaseGetPostRequest(): GetPostRequest {
  return { postId: "" };
}

export const GetPostRequest = {
  encode(message: GetPostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPostRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: GetPostRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPostRequest>, I>>(base?: I): GetPostRequest {
    return GetPostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPostRequest>, I>>(object: I): GetPostRequest {
    const message = createBaseGetPostRequest();
    message.postId = object.postId ?? "";
    return message;
  },
};

function createBaseGetPostResponse(): GetPostResponse {
  return { post: undefined };
}

export const GetPostResponse = {
  encode(message: GetPostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPostResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.post = Post.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPostResponse {
    return { post: isSet(object.post) ? Post.fromJSON(object.post) : undefined };
  },

  toJSON(message: GetPostResponse): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPostResponse>, I>>(base?: I): GetPostResponse {
    return GetPostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPostResponse>, I>>(object: I): GetPostResponse {
    const message = createBaseGetPostResponse();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    return message;
  },
};

function createBaseGetPostItemRequest(): GetPostItemRequest {
  return { postId: "" };
}

export const GetPostItemRequest = {
  encode(message: GetPostItemRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPostItemRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPostItemRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPostItemRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: GetPostItemRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPostItemRequest>, I>>(base?: I): GetPostItemRequest {
    return GetPostItemRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPostItemRequest>, I>>(object: I): GetPostItemRequest {
    const message = createBaseGetPostItemRequest();
    message.postId = object.postId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.post = Post.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UpdatePostRequest>, I>>(base?: I): UpdatePostRequest {
    return UpdatePostRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<DeletePostRequest>, I>>(base?: I): DeletePostRequest {
    return DeletePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletePostRequest>, I>>(object: I): DeletePostRequest {
    const message = createBaseDeletePostRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseListPostRequest(): ListPostRequest {
  return { username: "", pageSize: 0, page: 0 };
}

export const ListPostRequest = {
  encode(message: ListPostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.page = reader.int32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPostRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListPostRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPostRequest>, I>>(base?: I): ListPostRequest {
    return ListPostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPostRequest>, I>>(object: I): ListPostRequest {
    const message = createBaseListPostRequest();
    message.username = object.username ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListPostResponse(): ListPostResponse {
  return { posts: [], nextPageToken: "" };
}

export const ListPostResponse = {
  encode(message: ListPostResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.posts) {
      PostResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPostResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.posts.push(PostResponse.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPostResponse {
    return {
      posts: Array.isArray(object?.posts) ? object.posts.map((e: any) => PostResponse.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: ListPostResponse): unknown {
    const obj: any = {};
    if (message.posts) {
      obj.posts = message.posts.map((e) => e ? PostResponse.toJSON(e) : undefined);
    } else {
      obj.posts = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPostResponse>, I>>(base?: I): ListPostResponse {
    return ListPostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPostResponse>, I>>(object: I): ListPostResponse {
    const message = createBaseListPostResponse();
    message.posts = object.posts?.map((e) => PostResponse.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.parent = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.page = reader.int32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<FeedPostRequest>, I>>(base?: I): FeedPostRequest {
    return FeedPostRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.posts.push(PostResponse.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<FeedPostResponse>, I>>(base?: I): FeedPostResponse {
    return FeedPostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FeedPostResponse>, I>>(object: I): FeedPostResponse {
    const message = createBaseFeedPostResponse();
    message.posts = object.posts?.map((e) => PostResponse.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseLikePostRequest(): LikePostRequest {
  return { postId: "" };
}

export const LikePostRequest = {
  encode(message: LikePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LikePostRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: LikePostRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<LikePostRequest>, I>>(base?: I): LikePostRequest {
    return LikePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LikePostRequest>, I>>(object: I): LikePostRequest {
    const message = createBaseLikePostRequest();
    message.postId = object.postId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<LikePostResponse>, I>>(base?: I): LikePostResponse {
    return LikePostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<LikePostResponse>, I>>(object: I): LikePostResponse {
    const message = createBaseLikePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnlikePostRequest(): UnlikePostRequest {
  return { postId: "" };
}

export const UnlikePostRequest = {
  encode(message: UnlikePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnlikePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UnlikePostRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: UnlikePostRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<UnlikePostRequest>, I>>(base?: I): UnlikePostRequest {
    return UnlikePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnlikePostRequest>, I>>(object: I): UnlikePostRequest {
    const message = createBaseUnlikePostRequest();
    message.postId = object.postId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnlikePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UnlikePostResponse>, I>>(base?: I): UnlikePostResponse {
    return UnlikePostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnlikePostResponse>, I>>(object: I): UnlikePostResponse {
    const message = createBaseUnlikePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseSavePostRequest(): SavePostRequest {
  return { postId: "" };
}

export const SavePostRequest = {
  encode(message: SavePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SavePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSavePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SavePostRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: SavePostRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<SavePostRequest>, I>>(base?: I): SavePostRequest {
    return SavePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SavePostRequest>, I>>(object: I): SavePostRequest {
    const message = createBaseSavePostRequest();
    message.postId = object.postId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSavePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<SavePostResponse>, I>>(base?: I): SavePostResponse {
    return SavePostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SavePostResponse>, I>>(object: I): SavePostResponse {
    const message = createBaseSavePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseUnsavePostRequest(): UnsavePostRequest {
  return { postId: "" };
}

export const UnsavePostRequest = {
  encode(message: UnsavePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postId !== "") {
      writer.uint32(10).string(message.postId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnsavePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsavePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.postId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UnsavePostRequest {
    return { postId: isSet(object.postId) ? String(object.postId) : "" };
  },

  toJSON(message: UnsavePostRequest): unknown {
    const obj: any = {};
    message.postId !== undefined && (obj.postId = message.postId);
    return obj;
  },

  create<I extends Exact<DeepPartial<UnsavePostRequest>, I>>(base?: I): UnsavePostRequest {
    return UnsavePostRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnsavePostRequest>, I>>(object: I): UnsavePostRequest {
    const message = createBaseUnsavePostRequest();
    message.postId = object.postId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnsavePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UnsavePostResponse>, I>>(base?: I): UnsavePostResponse {
    return UnsavePostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnsavePostResponse>, I>>(object: I): UnsavePostResponse {
    const message = createBaseUnsavePostResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBasePostResponse(): PostResponse {
  return { post: undefined, user: undefined, likeCount: 0, commentCount: 0, hasLiked: false, isCreator: false };
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
    if (message.isCreator === true) {
      writer.uint32(48).bool(message.isCreator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PostResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePostResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.post = Post.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.likeCount = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.commentCount = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.hasLiked = reader.bool();
          continue;
        case 6:
          if (tag != 48) {
            break;
          }

          message.isCreator = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
      isCreator: isSet(object.isCreator) ? Boolean(object.isCreator) : false,
    };
  },

  toJSON(message: PostResponse): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.likeCount !== undefined && (obj.likeCount = Math.round(message.likeCount));
    message.commentCount !== undefined && (obj.commentCount = Math.round(message.commentCount));
    message.hasLiked !== undefined && (obj.hasLiked = message.hasLiked);
    message.isCreator !== undefined && (obj.isCreator = message.isCreator);
    return obj;
  },

  create<I extends Exact<DeepPartial<PostResponse>, I>>(base?: I): PostResponse {
    return PostResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PostResponse>, I>>(object: I): PostResponse {
    const message = createBasePostResponse();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.likeCount = object.likeCount ?? 0;
    message.commentCount = object.commentCount ?? 0;
    message.hasLiked = object.hasLiked ?? false;
    message.isCreator = object.isCreator ?? false;
    return message;
  },
};

export interface PostService {
  /** Create a new post */
  CreatePost(request: CreatePostRequest): Promise<PostResponse>;
  GetPost(request: GetPostRequest): Promise<GetPostResponse>;
  GetPostItem(request: GetPostItemRequest): Promise<PostResponse>;
  UpdatePost(request: UpdatePostRequest): Promise<PostResponse>;
  DeletePost(request: DeletePostRequest): Promise<Empty>;
  /** List posts of a specific user */
  ListPost(request: ListPostRequest): Promise<ListPostResponse>;
  /** Friends post feed */
  FeedPost(request: FeedPostRequest): Promise<FeedPostResponse>;
  /** Map post feed */
  ListMapPost(request: ListMapPostRequest): Promise<ListMapPostResponse>;
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
    this.GetPostItem = this.GetPostItem.bind(this);
    this.UpdatePost = this.UpdatePost.bind(this);
    this.DeletePost = this.DeletePost.bind(this);
    this.ListPost = this.ListPost.bind(this);
    this.FeedPost = this.FeedPost.bind(this);
    this.ListMapPost = this.ListMapPost.bind(this);
    this.LikePost = this.LikePost.bind(this);
    this.UnlikePost = this.UnlikePost.bind(this);
    this.SavePost = this.SavePost.bind(this);
    this.UnsavePost = this.UnsavePost.bind(this);
  }
  CreatePost(request: CreatePostRequest): Promise<PostResponse> {
    const data = CreatePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreatePost", data);
    return promise.then((data) => PostResponse.decode(_m0.Reader.create(data)));
  }

  GetPost(request: GetPostRequest): Promise<GetPostResponse> {
    const data = GetPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPost", data);
    return promise.then((data) => GetPostResponse.decode(_m0.Reader.create(data)));
  }

  GetPostItem(request: GetPostItemRequest): Promise<PostResponse> {
    const data = GetPostItemRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPostItem", data);
    return promise.then((data) => PostResponse.decode(_m0.Reader.create(data)));
  }

  UpdatePost(request: UpdatePostRequest): Promise<PostResponse> {
    const data = UpdatePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdatePost", data);
    return promise.then((data) => PostResponse.decode(_m0.Reader.create(data)));
  }

  DeletePost(request: DeletePostRequest): Promise<Empty> {
    const data = DeletePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeletePost", data);
    return promise.then((data) => Empty.decode(_m0.Reader.create(data)));
  }

  ListPost(request: ListPostRequest): Promise<ListPostResponse> {
    const data = ListPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListPost", data);
    return promise.then((data) => ListPostResponse.decode(_m0.Reader.create(data)));
  }

  FeedPost(request: FeedPostRequest): Promise<FeedPostResponse> {
    const data = FeedPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeedPost", data);
    return promise.then((data) => FeedPostResponse.decode(_m0.Reader.create(data)));
  }

  ListMapPost(request: ListMapPostRequest): Promise<ListMapPostResponse> {
    const data = ListMapPostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListMapPost", data);
    return promise.then((data) => ListMapPostResponse.decode(_m0.Reader.create(data)));
  }

  LikePost(request: LikePostRequest): Promise<LikePostResponse> {
    const data = LikePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "LikePost", data);
    return promise.then((data) => LikePostResponse.decode(_m0.Reader.create(data)));
  }

  UnlikePost(request: UnlikePostRequest): Promise<UnlikePostResponse> {
    const data = UnlikePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnlikePost", data);
    return promise.then((data) => UnlikePostResponse.decode(_m0.Reader.create(data)));
  }

  SavePost(request: SavePostRequest): Promise<SavePostResponse> {
    const data = SavePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SavePost", data);
    return promise.then((data) => SavePostResponse.decode(_m0.Reader.create(data)));
  }

  UnsavePost(request: UnsavePostRequest): Promise<UnsavePostResponse> {
    const data = UnsavePostRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnsavePost", data);
    return promise.then((data) => UnsavePostResponse.decode(_m0.Reader.create(data)));
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
