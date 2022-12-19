/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../../google/protobuf/empty";
import { StoryState, storyStateFromJSON, storyStateToJSON, User, UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.user.v1";

export interface CreateUserRequest {
  email: string;
  name: string;
  picture: string;
  username: string;
}

export interface CreateUserResponse {
  user: User | undefined;
}

export interface BlockUserResponse {
}

export interface UnblockUserResponse {
}

export interface GetUsersInRequest {
  userIds: string[];
}

export interface GetUsersInResponse {
  users: User[];
}

export interface GetUserItemsInRequest {
  userIds: string[];
}

export interface GetUserItemsInResponse {
  users: UserItem[];
}

export interface SearchUserResponse {
  users: UserItem[];
}

export interface ListFollowersRequest {
  userId: string;
  pageSize: number;
  page: number;
}

export interface ListFollowingRequest {
  userId: string;
  pageSize: number;
  page: number;
}

export interface ListFollowersResponse {
  users: UserItem[];
}

export interface ListFollowingResponse {
  users: UserItem[];
}

export interface GetUserResponse {
  user: User | undefined;
  postCount: number;
  followersCount: number;
  followingCount: number;
  hasFollowed: boolean;
  storyState: StoryState;
}

export interface BlockUserRequest {
  userId: string;
}

export interface UnblockUserRequest {
  userId: string;
}

export interface UnfollowUserRequest {
  userId: string;
}

export interface FollowUserRequest {
  userId: string;
}

export interface SearchUserRequest {
  query: string;
}

export interface GetUserRequest {
  userId: string;
}

export interface UpdateUserRequest {
  name: string;
  picture: string;
  bio: string;
  website: string;
  birthday: string;
  phoneNumber: string;
}

export interface DeleteUserRequest {
  userId: string;
}

function createBaseCreateUserRequest(): CreateUserRequest {
  return { email: "", name: "", picture: "", username: "" };
}

export const CreateUserRequest = {
  encode(message: CreateUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.picture !== "") {
      writer.uint32(34).string(message.picture);
    }
    if (message.username !== "") {
      writer.uint32(42).string(message.username);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 4:
          message.picture = reader.string();
          break;
        case 5:
          message.username = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserRequest {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      name: isSet(object.name) ? String(object.name) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      username: isSet(object.username) ? String(object.username) : "",
    };
  },

  toJSON(message: CreateUserRequest): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.name !== undefined && (obj.name = message.name);
    message.picture !== undefined && (obj.picture = message.picture);
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserRequest>, I>>(object: I): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.email = object.email ?? "";
    message.name = object.name ?? "";
    message.picture = object.picture ?? "";
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseCreateUserResponse(): CreateUserResponse {
  return { user: undefined };
}

export const CreateUserResponse = {
  encode(message: CreateUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: CreateUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserResponse>, I>>(object: I): CreateUserResponse {
    const message = createBaseCreateUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseBlockUserResponse(): BlockUserResponse {
  return {};
}

export const BlockUserResponse = {
  encode(_: BlockUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BlockUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlockUserResponse();
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

  fromJSON(_: any): BlockUserResponse {
    return {};
  },

  toJSON(_: BlockUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BlockUserResponse>, I>>(_: I): BlockUserResponse {
    const message = createBaseBlockUserResponse();
    return message;
  },
};

function createBaseUnblockUserResponse(): UnblockUserResponse {
  return {};
}

export const UnblockUserResponse = {
  encode(_: UnblockUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnblockUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnblockUserResponse();
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

  fromJSON(_: any): UnblockUserResponse {
    return {};
  },

  toJSON(_: UnblockUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnblockUserResponse>, I>>(_: I): UnblockUserResponse {
    const message = createBaseUnblockUserResponse();
    return message;
  },
};

function createBaseGetUsersInRequest(): GetUsersInRequest {
  return { userIds: [] };
}

export const GetUsersInRequest = {
  encode(message: GetUsersInRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userIds) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUsersInRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUsersInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userIds.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUsersInRequest {
    return { userIds: Array.isArray(object?.userIds) ? object.userIds.map((e: any) => String(e)) : [] };
  },

  toJSON(message: GetUsersInRequest): unknown {
    const obj: any = {};
    if (message.userIds) {
      obj.userIds = message.userIds.map((e) => e);
    } else {
      obj.userIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUsersInRequest>, I>>(object: I): GetUsersInRequest {
    const message = createBaseGetUsersInRequest();
    message.userIds = object.userIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseGetUsersInResponse(): GetUsersInResponse {
  return { users: [] };
}

export const GetUsersInResponse = {
  encode(message: GetUsersInResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      User.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUsersInResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUsersInResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(User.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUsersInResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => User.fromJSON(e)) : [] };
  },

  toJSON(message: GetUsersInResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUsersInResponse>, I>>(object: I): GetUsersInResponse {
    const message = createBaseGetUsersInResponse();
    message.users = object.users?.map((e) => User.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetUserItemsInRequest(): GetUserItemsInRequest {
  return { userIds: [] };
}

export const GetUserItemsInRequest = {
  encode(message: GetUserItemsInRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userIds) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserItemsInRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserItemsInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userIds.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserItemsInRequest {
    return { userIds: Array.isArray(object?.userIds) ? object.userIds.map((e: any) => String(e)) : [] };
  },

  toJSON(message: GetUserItemsInRequest): unknown {
    const obj: any = {};
    if (message.userIds) {
      obj.userIds = message.userIds.map((e) => e);
    } else {
      obj.userIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserItemsInRequest>, I>>(object: I): GetUserItemsInRequest {
    const message = createBaseGetUserItemsInRequest();
    message.userIds = object.userIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseGetUserItemsInResponse(): GetUserItemsInResponse {
  return { users: [] };
}

export const GetUserItemsInResponse = {
  encode(message: GetUserItemsInResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserItemsInResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserItemsInResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserItemsInResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: GetUserItemsInResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserItemsInResponse>, I>>(object: I): GetUserItemsInResponse {
    const message = createBaseGetUserItemsInResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSearchUserResponse(): SearchUserResponse {
  return { users: [] };
}

export const SearchUserResponse = {
  encode(message: SearchUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchUserResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: SearchUserResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchUserResponse>, I>>(object: I): SearchUserResponse {
    const message = createBaseSearchUserResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListFollowersRequest(): ListFollowersRequest {
  return { userId: "", pageSize: 0, page: 0 };
}

export const ListFollowersRequest = {
  encode(message: ListFollowersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFollowersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
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

  fromJSON(object: any): ListFollowersRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListFollowersRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListFollowersRequest>, I>>(object: I): ListFollowersRequest {
    const message = createBaseListFollowersRequest();
    message.userId = object.userId ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListFollowingRequest(): ListFollowingRequest {
  return { userId: "", pageSize: 0, page: 0 };
}

export const ListFollowingRequest = {
  encode(message: ListFollowingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFollowingRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
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

  fromJSON(object: any): ListFollowingRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListFollowingRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListFollowingRequest>, I>>(object: I): ListFollowingRequest {
    const message = createBaseListFollowingRequest();
    message.userId = object.userId ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListFollowersResponse(): ListFollowersResponse {
  return { users: [] };
}

export const ListFollowersResponse = {
  encode(message: ListFollowersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFollowersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListFollowersResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListFollowersResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListFollowersResponse>, I>>(object: I): ListFollowersResponse {
    const message = createBaseListFollowersResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListFollowingResponse(): ListFollowingResponse {
  return { users: [] };
}

export const ListFollowingResponse = {
  encode(message: ListFollowingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFollowingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListFollowingResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListFollowingResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListFollowingResponse>, I>>(object: I): ListFollowingResponse {
    const message = createBaseListFollowingResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetUserResponse(): GetUserResponse {
  return { user: undefined, postCount: 0, followersCount: 0, followingCount: 0, hasFollowed: false, storyState: 0 };
}

export const GetUserResponse = {
  encode(message: GetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    if (message.postCount !== 0) {
      writer.uint32(16).int32(message.postCount);
    }
    if (message.followersCount !== 0) {
      writer.uint32(24).int32(message.followersCount);
    }
    if (message.followingCount !== 0) {
      writer.uint32(32).int32(message.followingCount);
    }
    if (message.hasFollowed === true) {
      writer.uint32(40).bool(message.hasFollowed);
    }
    if (message.storyState !== 0) {
      writer.uint32(48).int32(message.storyState);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.postCount = reader.int32();
          break;
        case 3:
          message.followersCount = reader.int32();
          break;
        case 4:
          message.followingCount = reader.int32();
          break;
        case 5:
          message.hasFollowed = reader.bool();
          break;
        case 6:
          message.storyState = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      postCount: isSet(object.postCount) ? Number(object.postCount) : 0,
      followersCount: isSet(object.followersCount) ? Number(object.followersCount) : 0,
      followingCount: isSet(object.followingCount) ? Number(object.followingCount) : 0,
      hasFollowed: isSet(object.hasFollowed) ? Boolean(object.hasFollowed) : false,
      storyState: isSet(object.storyState) ? storyStateFromJSON(object.storyState) : 0,
    };
  },

  toJSON(message: GetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.postCount !== undefined && (obj.postCount = Math.round(message.postCount));
    message.followersCount !== undefined && (obj.followersCount = Math.round(message.followersCount));
    message.followingCount !== undefined && (obj.followingCount = Math.round(message.followingCount));
    message.hasFollowed !== undefined && (obj.hasFollowed = message.hasFollowed);
    message.storyState !== undefined && (obj.storyState = storyStateToJSON(message.storyState));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserResponse>, I>>(object: I): GetUserResponse {
    const message = createBaseGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.postCount = object.postCount ?? 0;
    message.followersCount = object.followersCount ?? 0;
    message.followingCount = object.followingCount ?? 0;
    message.hasFollowed = object.hasFollowed ?? false;
    message.storyState = object.storyState ?? 0;
    return message;
  },
};

function createBaseBlockUserRequest(): BlockUserRequest {
  return { userId: "" };
}

export const BlockUserRequest = {
  encode(message: BlockUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BlockUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlockUserRequest();
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

  fromJSON(object: any): BlockUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: BlockUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BlockUserRequest>, I>>(object: I): BlockUserRequest {
    const message = createBaseBlockUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseUnblockUserRequest(): UnblockUserRequest {
  return { userId: "" };
}

export const UnblockUserRequest = {
  encode(message: UnblockUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnblockUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnblockUserRequest();
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

  fromJSON(object: any): UnblockUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: UnblockUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnblockUserRequest>, I>>(object: I): UnblockUserRequest {
    const message = createBaseUnblockUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseUnfollowUserRequest(): UnfollowUserRequest {
  return { userId: "" };
}

export const UnfollowUserRequest = {
  encode(message: UnfollowUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnfollowUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnfollowUserRequest();
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

  fromJSON(object: any): UnfollowUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: UnfollowUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UnfollowUserRequest>, I>>(object: I): UnfollowUserRequest {
    const message = createBaseUnfollowUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseFollowUserRequest(): FollowUserRequest {
  return { userId: "" };
}

export const FollowUserRequest = {
  encode(message: FollowUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FollowUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFollowUserRequest();
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

  fromJSON(object: any): FollowUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: FollowUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FollowUserRequest>, I>>(object: I): FollowUserRequest {
    const message = createBaseFollowUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseSearchUserRequest(): SearchUserRequest {
  return { query: "" };
}

export const SearchUserRequest = {
  encode(message: SearchUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.query = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchUserRequest {
    return { query: isSet(object.query) ? String(object.query) : "" };
  },

  toJSON(message: SearchUserRequest): unknown {
    const obj: any = {};
    message.query !== undefined && (obj.query = message.query);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchUserRequest>, I>>(object: I): SearchUserRequest {
    const message = createBaseSearchUserRequest();
    message.query = object.query ?? "";
    return message;
  },
};

function createBaseGetUserRequest(): GetUserRequest {
  return { userId: "" };
}

export const GetUserRequest = {
  encode(message: GetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
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

  fromJSON(object: any): GetUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: GetUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserRequest>, I>>(object: I): GetUserRequest {
    const message = createBaseGetUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseUpdateUserRequest(): UpdateUserRequest {
  return { name: "", picture: "", bio: "", website: "", birthday: "", phoneNumber: "" };
}

export const UpdateUserRequest = {
  encode(message: UpdateUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.picture !== "") {
      writer.uint32(18).string(message.picture);
    }
    if (message.bio !== "") {
      writer.uint32(26).string(message.bio);
    }
    if (message.website !== "") {
      writer.uint32(34).string(message.website);
    }
    if (message.birthday !== "") {
      writer.uint32(42).string(message.birthday);
    }
    if (message.phoneNumber !== "") {
      writer.uint32(50).string(message.phoneNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.picture = reader.string();
          break;
        case 3:
          message.bio = reader.string();
          break;
        case 4:
          message.website = reader.string();
          break;
        case 5:
          message.birthday = reader.string();
          break;
        case 6:
          message.phoneNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      bio: isSet(object.bio) ? String(object.bio) : "",
      website: isSet(object.website) ? String(object.website) : "",
      birthday: isSet(object.birthday) ? String(object.birthday) : "",
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
    };
  },

  toJSON(message: UpdateUserRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.picture !== undefined && (obj.picture = message.picture);
    message.bio !== undefined && (obj.bio = message.bio);
    message.website !== undefined && (obj.website = message.website);
    message.birthday !== undefined && (obj.birthday = message.birthday);
    message.phoneNumber !== undefined && (obj.phoneNumber = message.phoneNumber);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(object: I): UpdateUserRequest {
    const message = createBaseUpdateUserRequest();
    message.name = object.name ?? "";
    message.picture = object.picture ?? "";
    message.bio = object.bio ?? "";
    message.website = object.website ?? "";
    message.birthday = object.birthday ?? "";
    message.phoneNumber = object.phoneNumber ?? "";
    return message;
  },
};

function createBaseDeleteUserRequest(): DeleteUserRequest {
  return { userId: "" };
}

export const DeleteUserRequest = {
  encode(message: DeleteUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserRequest();
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

  fromJSON(object: any): DeleteUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: DeleteUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(object: I): DeleteUserRequest {
    const message = createBaseDeleteUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

export interface UserService {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  GetUser(request: GetUserRequest): Promise<GetUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<User>;
  DeleteUser(request: DeleteUserRequest): Promise<Empty>;
  GetUsersIn(request: GetUsersInRequest): Promise<GetUsersInResponse>;
  GetUserItemsIn(request: GetUserItemsInRequest): Promise<GetUserItemsInResponse>;
  SearchUser(request: SearchUserRequest): Promise<SearchUserResponse>;
  FollowUser(request: FollowUserRequest): Promise<Empty>;
  UnfollowUser(request: UnfollowUserRequest): Promise<Empty>;
  BlockUser(request: BlockUserRequest): Promise<BlockUserResponse>;
  UnblockUser(request: UnblockUserRequest): Promise<UnblockUserResponse>;
  ListFollowers(request: ListFollowersRequest): Promise<ListFollowersResponse>;
  ListFollowing(request: ListFollowingRequest): Promise<ListFollowingResponse>;
}

export class UserServiceClientImpl implements UserService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.user.v1.UserService";
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.GetUser = this.GetUser.bind(this);
    this.UpdateUser = this.UpdateUser.bind(this);
    this.DeleteUser = this.DeleteUser.bind(this);
    this.GetUsersIn = this.GetUsersIn.bind(this);
    this.GetUserItemsIn = this.GetUserItemsIn.bind(this);
    this.SearchUser = this.SearchUser.bind(this);
    this.FollowUser = this.FollowUser.bind(this);
    this.UnfollowUser = this.UnfollowUser.bind(this);
    this.BlockUser = this.BlockUser.bind(this);
    this.UnblockUser = this.UnblockUser.bind(this);
    this.ListFollowers = this.ListFollowers.bind(this);
    this.ListFollowing = this.ListFollowing.bind(this);
  }
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse> {
    const data = CreateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateUser", data);
    return promise.then((data) => CreateUserResponse.decode(new _m0.Reader(data)));
  }

  GetUser(request: GetUserRequest): Promise<GetUserResponse> {
    const data = GetUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUser", data);
    return promise.then((data) => GetUserResponse.decode(new _m0.Reader(data)));
  }

  UpdateUser(request: UpdateUserRequest): Promise<User> {
    const data = UpdateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateUser", data);
    return promise.then((data) => User.decode(new _m0.Reader(data)));
  }

  DeleteUser(request: DeleteUserRequest): Promise<Empty> {
    const data = DeleteUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteUser", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  GetUsersIn(request: GetUsersInRequest): Promise<GetUsersInResponse> {
    const data = GetUsersInRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUsersIn", data);
    return promise.then((data) => GetUsersInResponse.decode(new _m0.Reader(data)));
  }

  GetUserItemsIn(request: GetUserItemsInRequest): Promise<GetUserItemsInResponse> {
    const data = GetUserItemsInRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserItemsIn", data);
    return promise.then((data) => GetUserItemsInResponse.decode(new _m0.Reader(data)));
  }

  SearchUser(request: SearchUserRequest): Promise<SearchUserResponse> {
    const data = SearchUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SearchUser", data);
    return promise.then((data) => SearchUserResponse.decode(new _m0.Reader(data)));
  }

  FollowUser(request: FollowUserRequest): Promise<Empty> {
    const data = FollowUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FollowUser", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  UnfollowUser(request: UnfollowUserRequest): Promise<Empty> {
    const data = UnfollowUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnfollowUser", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  BlockUser(request: BlockUserRequest): Promise<BlockUserResponse> {
    const data = BlockUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "BlockUser", data);
    return promise.then((data) => BlockUserResponse.decode(new _m0.Reader(data)));
  }

  UnblockUser(request: UnblockUserRequest): Promise<UnblockUserResponse> {
    const data = UnblockUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnblockUser", data);
    return promise.then((data) => UnblockUserResponse.decode(new _m0.Reader(data)));
  }

  ListFollowers(request: ListFollowersRequest): Promise<ListFollowersResponse> {
    const data = ListFollowersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFollowers", data);
    return promise.then((data) => ListFollowersResponse.decode(new _m0.Reader(data)));
  }

  ListFollowing(request: ListFollowingRequest): Promise<ListFollowingResponse> {
    const data = ListFollowingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFollowing", data);
    return promise.then((data) => ListFollowingResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
