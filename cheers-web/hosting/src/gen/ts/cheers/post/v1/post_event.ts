/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Post } from "../../type/post/post";
import { UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.post.v1";

export interface PostEvent {
  create?: CreatePost | undefined;
  like?: LikePost | undefined;
  delete?: DeletePost | undefined;
}

export interface CreatePost {
  post: Post | undefined;
  user: UserItem | undefined;
}

export interface LikePost {
  post: Post | undefined;
  user: UserItem | undefined;
}

export interface DeletePost {
  sender: UserItem | undefined;
}

function createBasePostEvent(): PostEvent {
  return { create: undefined, like: undefined, delete: undefined };
}

export const PostEvent = {
  encode(message: PostEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.create !== undefined) {
      CreatePost.encode(message.create, writer.uint32(10).fork()).ldelim();
    }
    if (message.like !== undefined) {
      LikePost.encode(message.like, writer.uint32(18).fork()).ldelim();
    }
    if (message.delete !== undefined) {
      DeletePost.encode(message.delete, writer.uint32(26).fork()).ldelim();
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
          message.create = CreatePost.decode(reader, reader.uint32());
          break;
        case 2:
          message.like = LikePost.decode(reader, reader.uint32());
          break;
        case 3:
          message.delete = DeletePost.decode(reader, reader.uint32());
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
      create: isSet(object.create) ? CreatePost.fromJSON(object.create) : undefined,
      like: isSet(object.like) ? LikePost.fromJSON(object.like) : undefined,
      delete: isSet(object.delete) ? DeletePost.fromJSON(object.delete) : undefined,
    };
  },

  toJSON(message: PostEvent): unknown {
    const obj: any = {};
    message.create !== undefined && (obj.create = message.create ? CreatePost.toJSON(message.create) : undefined);
    message.like !== undefined && (obj.like = message.like ? LikePost.toJSON(message.like) : undefined);
    message.delete !== undefined && (obj.delete = message.delete ? DeletePost.toJSON(message.delete) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PostEvent>, I>>(object: I): PostEvent {
    const message = createBasePostEvent();
    message.create = (object.create !== undefined && object.create !== null)
      ? CreatePost.fromPartial(object.create)
      : undefined;
    message.like = (object.like !== undefined && object.like !== null) ? LikePost.fromPartial(object.like) : undefined;
    message.delete = (object.delete !== undefined && object.delete !== null)
      ? DeletePost.fromPartial(object.delete)
      : undefined;
    return message;
  },
};

function createBaseCreatePost(): CreatePost {
  return { post: undefined, user: undefined };
}

export const CreatePost = {
  encode(message: CreatePost, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePost {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePost();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.post = Post.decode(reader, reader.uint32());
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

  fromJSON(object: any): CreatePost {
    return {
      post: isSet(object.post) ? Post.fromJSON(object.post) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreatePost): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatePost>, I>>(object: I): CreatePost {
    const message = createBaseCreatePost();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseLikePost(): LikePost {
  return { post: undefined, user: undefined };
}

export const LikePost = {
  encode(message: LikePost, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikePost {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePost();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.post = Post.decode(reader, reader.uint32());
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

  fromJSON(object: any): LikePost {
    return {
      post: isSet(object.post) ? Post.fromJSON(object.post) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: LikePost): unknown {
    const obj: any = {};
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikePost>, I>>(object: I): LikePost {
    const message = createBaseLikePost();
    message.post = (object.post !== undefined && object.post !== null) ? Post.fromPartial(object.post) : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseDeletePost(): DeletePost {
  return { sender: undefined };
}

export const DeletePost = {
  encode(message: DeletePost, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sender !== undefined) {
      UserItem.encode(message.sender, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePost {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePost();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.sender = UserItem.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletePost {
    return { sender: isSet(object.sender) ? UserItem.fromJSON(object.sender) : undefined };
  },

  toJSON(message: DeletePost): unknown {
    const obj: any = {};
    message.sender !== undefined && (obj.sender = message.sender ? UserItem.toJSON(message.sender) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePost>, I>>(object: I): DeletePost {
    const message = createBaseDeletePost();
    message.sender = (object.sender !== undefined && object.sender !== null)
      ? UserItem.fromPartial(object.sender)
      : undefined;
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
