/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../../google/protobuf/empty";
import { StoryState, storyStateFromJSON, storyStateToJSON, User, UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.user.v1";

export interface ListSuggestionsRequest {
  userId: string;
}

export interface ListSuggestionsResponse {
  users: UserItem[];
}

export interface CheckUsernameRequest {
  username: string;
}

export interface CheckUsernameResponse {
  valid: boolean;
}

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

export interface GetUserNodeResponse {
  user: User | undefined;
}

export interface GetUserResponse {
  user: User | undefined;
  postCount: number;
  followersCount: number;
  followingCount: number;
  hasFollowed: boolean;
  storyState: StoryState;
  friend: boolean;
  requested: boolean;
  hasRequestedViewer: boolean;
  friendsCount: number;
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

export interface GetUserNodeRequest {
  userId: string;
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

function createBaseListSuggestionsRequest(): ListSuggestionsRequest {
  return { userId: "" };
}

export const ListSuggestionsRequest = {
  encode(message: ListSuggestionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSuggestionsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSuggestionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSuggestionsRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: ListSuggestionsRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSuggestionsRequest>, I>>(base?: I): ListSuggestionsRequest {
    return ListSuggestionsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSuggestionsRequest>, I>>(object: I): ListSuggestionsRequest {
    const message = createBaseListSuggestionsRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListSuggestionsResponse(): ListSuggestionsResponse {
  return { users: [] };
}

export const ListSuggestionsResponse = {
  encode(message: ListSuggestionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSuggestionsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSuggestionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSuggestionsResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListSuggestionsResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSuggestionsResponse>, I>>(base?: I): ListSuggestionsResponse {
    return ListSuggestionsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSuggestionsResponse>, I>>(object: I): ListSuggestionsResponse {
    const message = createBaseListSuggestionsResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCheckUsernameRequest(): CheckUsernameRequest {
  return { username: "" };
}

export const CheckUsernameRequest = {
  encode(message: CheckUsernameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckUsernameRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckUsernameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.username = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CheckUsernameRequest {
    return { username: isSet(object.username) ? String(object.username) : "" };
  },

  toJSON(message: CheckUsernameRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckUsernameRequest>, I>>(base?: I): CheckUsernameRequest {
    return CheckUsernameRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CheckUsernameRequest>, I>>(object: I): CheckUsernameRequest {
    const message = createBaseCheckUsernameRequest();
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseCheckUsernameResponse(): CheckUsernameResponse {
  return { valid: false };
}

export const CheckUsernameResponse = {
  encode(message: CheckUsernameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.valid === true) {
      writer.uint32(8).bool(message.valid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckUsernameResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckUsernameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.valid = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CheckUsernameResponse {
    return { valid: isSet(object.valid) ? Boolean(object.valid) : false };
  },

  toJSON(message: CheckUsernameResponse): unknown {
    const obj: any = {};
    message.valid !== undefined && (obj.valid = message.valid);
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckUsernameResponse>, I>>(base?: I): CheckUsernameResponse {
    return CheckUsernameResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CheckUsernameResponse>, I>>(object: I): CheckUsernameResponse {
    const message = createBaseCheckUsernameResponse();
    message.valid = object.valid ?? false;
    return message;
  },
};

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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.email = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.username = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<CreateUserRequest>, I>>(base?: I): CreateUserRequest {
    return CreateUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<CreateUserResponse>, I>>(base?: I): CreateUserResponse {
    return CreateUserResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlockUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<BlockUserResponse>, I>>(base?: I): BlockUserResponse {
    return BlockUserResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnblockUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UnblockUserResponse>, I>>(base?: I): UnblockUserResponse {
    return UnblockUserResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUsersInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userIds.push(reader.string());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<GetUsersInRequest>, I>>(base?: I): GetUsersInRequest {
    return GetUsersInRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUsersInResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(User.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<GetUsersInResponse>, I>>(base?: I): GetUsersInResponse {
    return GetUsersInResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserItemsInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userIds.push(reader.string());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<GetUserItemsInRequest>, I>>(base?: I): GetUserItemsInRequest {
    return GetUserItemsInRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserItemsInResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<GetUserItemsInResponse>, I>>(base?: I): GetUserItemsInResponse {
    return GetUserItemsInResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<SearchUserResponse>, I>>(base?: I): SearchUserResponse {
    return SearchUserResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
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

  create<I extends Exact<DeepPartial<ListFollowersRequest>, I>>(base?: I): ListFollowersRequest {
    return ListFollowersRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
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

  create<I extends Exact<DeepPartial<ListFollowingRequest>, I>>(base?: I): ListFollowingRequest {
    return ListFollowingRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<ListFollowersResponse>, I>>(base?: I): ListFollowersResponse {
    return ListFollowersResponse.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFollowingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<ListFollowingResponse>, I>>(base?: I): ListFollowingResponse {
    return ListFollowingResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFollowingResponse>, I>>(object: I): ListFollowingResponse {
    const message = createBaseListFollowingResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetUserNodeResponse(): GetUserNodeResponse {
  return { user: undefined };
}

export const GetUserNodeResponse = {
  encode(message: GetUserNodeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserNodeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserNodeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserNodeResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: GetUserNodeResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserNodeResponse>, I>>(base?: I): GetUserNodeResponse {
    return GetUserNodeResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserNodeResponse>, I>>(object: I): GetUserNodeResponse {
    const message = createBaseGetUserNodeResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseGetUserResponse(): GetUserResponse {
  return {
    user: undefined,
    postCount: 0,
    followersCount: 0,
    followingCount: 0,
    hasFollowed: false,
    storyState: 0,
    friend: false,
    requested: false,
    hasRequestedViewer: false,
    friendsCount: 0,
  };
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
    if (message.friend === true) {
      writer.uint32(56).bool(message.friend);
    }
    if (message.requested === true) {
      writer.uint32(64).bool(message.requested);
    }
    if (message.hasRequestedViewer === true) {
      writer.uint32(72).bool(message.hasRequestedViewer);
    }
    if (message.friendsCount !== 0) {
      writer.uint32(80).int32(message.friendsCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.postCount = reader.int32();
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.followersCount = reader.int32();
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.followingCount = reader.int32();
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.hasFollowed = reader.bool();
          continue;
        case 6:
          if (tag != 48) {
            break;
          }

          message.storyState = reader.int32() as any;
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.friend = reader.bool();
          continue;
        case 8:
          if (tag != 64) {
            break;
          }

          message.requested = reader.bool();
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.hasRequestedViewer = reader.bool();
          continue;
        case 10:
          if (tag != 80) {
            break;
          }

          message.friendsCount = reader.int32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
      friend: isSet(object.friend) ? Boolean(object.friend) : false,
      requested: isSet(object.requested) ? Boolean(object.requested) : false,
      hasRequestedViewer: isSet(object.hasRequestedViewer) ? Boolean(object.hasRequestedViewer) : false,
      friendsCount: isSet(object.friendsCount) ? Number(object.friendsCount) : 0,
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
    message.friend !== undefined && (obj.friend = message.friend);
    message.requested !== undefined && (obj.requested = message.requested);
    message.hasRequestedViewer !== undefined && (obj.hasRequestedViewer = message.hasRequestedViewer);
    message.friendsCount !== undefined && (obj.friendsCount = Math.round(message.friendsCount));
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserResponse>, I>>(base?: I): GetUserResponse {
    return GetUserResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserResponse>, I>>(object: I): GetUserResponse {
    const message = createBaseGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.postCount = object.postCount ?? 0;
    message.followersCount = object.followersCount ?? 0;
    message.followingCount = object.followingCount ?? 0;
    message.hasFollowed = object.hasFollowed ?? false;
    message.storyState = object.storyState ?? 0;
    message.friend = object.friend ?? false;
    message.requested = object.requested ?? false;
    message.hasRequestedViewer = object.hasRequestedViewer ?? false;
    message.friendsCount = object.friendsCount ?? 0;
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlockUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<BlockUserRequest>, I>>(base?: I): BlockUserRequest {
    return BlockUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnblockUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UnblockUserRequest>, I>>(base?: I): UnblockUserRequest {
    return UnblockUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnfollowUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UnfollowUserRequest>, I>>(base?: I): UnfollowUserRequest {
    return UnfollowUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFollowUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<FollowUserRequest>, I>>(base?: I): FollowUserRequest {
    return FollowUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.query = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<SearchUserRequest>, I>>(base?: I): SearchUserRequest {
    return SearchUserRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SearchUserRequest>, I>>(object: I): SearchUserRequest {
    const message = createBaseSearchUserRequest();
    message.query = object.query ?? "";
    return message;
  },
};

function createBaseGetUserNodeRequest(): GetUserNodeRequest {
  return { userId: "" };
}

export const GetUserNodeRequest = {
  encode(message: GetUserNodeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserNodeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserNodeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserNodeRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: GetUserNodeRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserNodeRequest>, I>>(base?: I): GetUserNodeRequest {
    return GetUserNodeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserNodeRequest>, I>>(object: I): GetUserNodeRequest {
    const message = createBaseGetUserNodeRequest();
    message.userId = object.userId ?? "";
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<GetUserRequest>, I>>(base?: I): GetUserRequest {
    return GetUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.bio = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.website = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.birthday = reader.string();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.phoneNumber = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(base?: I): UpdateUserRequest {
    return UpdateUserRequest.fromPartial(base ?? {});
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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

  create<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(base?: I): DeleteUserRequest {
    return DeleteUserRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(object: I): DeleteUserRequest {
    const message = createBaseDeleteUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

export interface UserService {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  GetUserNode(request: GetUserNodeRequest): Promise<GetUserNodeResponse>;
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
  CheckUsername(request: CheckUsernameRequest): Promise<CheckUsernameResponse>;
  ListSuggestions(request: ListSuggestionsRequest): Promise<ListSuggestionsResponse>;
}

export class UserServiceClientImpl implements UserService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.user.v1.UserService";
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.GetUserNode = this.GetUserNode.bind(this);
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
    this.CheckUsername = this.CheckUsername.bind(this);
    this.ListSuggestions = this.ListSuggestions.bind(this);
  }
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse> {
    const data = CreateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateUser", data);
    return promise.then((data) => CreateUserResponse.decode(_m0.Reader.create(data)));
  }

  GetUserNode(request: GetUserNodeRequest): Promise<GetUserNodeResponse> {
    const data = GetUserNodeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserNode", data);
    return promise.then((data) => GetUserNodeResponse.decode(_m0.Reader.create(data)));
  }

  GetUser(request: GetUserRequest): Promise<GetUserResponse> {
    const data = GetUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUser", data);
    return promise.then((data) => GetUserResponse.decode(_m0.Reader.create(data)));
  }

  UpdateUser(request: UpdateUserRequest): Promise<User> {
    const data = UpdateUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateUser", data);
    return promise.then((data) => User.decode(_m0.Reader.create(data)));
  }

  DeleteUser(request: DeleteUserRequest): Promise<Empty> {
    const data = DeleteUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteUser", data);
    return promise.then((data) => Empty.decode(_m0.Reader.create(data)));
  }

  GetUsersIn(request: GetUsersInRequest): Promise<GetUsersInResponse> {
    const data = GetUsersInRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUsersIn", data);
    return promise.then((data) => GetUsersInResponse.decode(_m0.Reader.create(data)));
  }

  GetUserItemsIn(request: GetUserItemsInRequest): Promise<GetUserItemsInResponse> {
    const data = GetUserItemsInRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetUserItemsIn", data);
    return promise.then((data) => GetUserItemsInResponse.decode(_m0.Reader.create(data)));
  }

  SearchUser(request: SearchUserRequest): Promise<SearchUserResponse> {
    const data = SearchUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SearchUser", data);
    return promise.then((data) => SearchUserResponse.decode(_m0.Reader.create(data)));
  }

  FollowUser(request: FollowUserRequest): Promise<Empty> {
    const data = FollowUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FollowUser", data);
    return promise.then((data) => Empty.decode(_m0.Reader.create(data)));
  }

  UnfollowUser(request: UnfollowUserRequest): Promise<Empty> {
    const data = UnfollowUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnfollowUser", data);
    return promise.then((data) => Empty.decode(_m0.Reader.create(data)));
  }

  BlockUser(request: BlockUserRequest): Promise<BlockUserResponse> {
    const data = BlockUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "BlockUser", data);
    return promise.then((data) => BlockUserResponse.decode(_m0.Reader.create(data)));
  }

  UnblockUser(request: UnblockUserRequest): Promise<UnblockUserResponse> {
    const data = UnblockUserRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnblockUser", data);
    return promise.then((data) => UnblockUserResponse.decode(_m0.Reader.create(data)));
  }

  ListFollowers(request: ListFollowersRequest): Promise<ListFollowersResponse> {
    const data = ListFollowersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFollowers", data);
    return promise.then((data) => ListFollowersResponse.decode(_m0.Reader.create(data)));
  }

  ListFollowing(request: ListFollowingRequest): Promise<ListFollowingResponse> {
    const data = ListFollowingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFollowing", data);
    return promise.then((data) => ListFollowingResponse.decode(_m0.Reader.create(data)));
  }

  CheckUsername(request: CheckUsernameRequest): Promise<CheckUsernameResponse> {
    const data = CheckUsernameRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CheckUsername", data);
    return promise.then((data) => CheckUsernameResponse.decode(_m0.Reader.create(data)));
  }

  ListSuggestions(request: ListSuggestionsRequest): Promise<ListSuggestionsResponse> {
    const data = ListSuggestionsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListSuggestions", data);
    return promise.then((data) => ListSuggestionsResponse.decode(_m0.Reader.create(data)));
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
