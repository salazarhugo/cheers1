/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.friendship.v1";

export interface AcceptFriendRequestRequest {
  userId: string;
}

export interface AcceptFriendRequestResponse {
}

export interface DeleteFriendRequestRequest {
  userId: string;
}

export interface DeleteFriendRequestResponse {
}

export interface ListFriendResponse {
  items: UserItem[];
}

export interface DeleteFriendRequest2 {
  userId: string;
  friendshipId: string;
}

export interface DeleteFriendResponse {
}

export interface ListFriendRequest {
  userId: string;
}

export interface ListFriendRequestsRequest {
  userId: string;
}

export interface ListFriendRequestsResponse {
  items: UserItem[];
}

export interface CreateFriendRequestRequest {
  userId: string;
}

export interface CreateFriendRequestResponse {
}

function createBaseAcceptFriendRequestRequest(): AcceptFriendRequestRequest {
  return { userId: "" };
}

export const AcceptFriendRequestRequest = {
  encode(message: AcceptFriendRequestRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AcceptFriendRequestRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAcceptFriendRequestRequest();
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

  fromJSON(object: any): AcceptFriendRequestRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: AcceptFriendRequestRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<AcceptFriendRequestRequest>, I>>(base?: I): AcceptFriendRequestRequest {
    return AcceptFriendRequestRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AcceptFriendRequestRequest>, I>>(object: I): AcceptFriendRequestRequest {
    const message = createBaseAcceptFriendRequestRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseAcceptFriendRequestResponse(): AcceptFriendRequestResponse {
  return {};
}

export const AcceptFriendRequestResponse = {
  encode(_: AcceptFriendRequestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AcceptFriendRequestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAcceptFriendRequestResponse();
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

  fromJSON(_: any): AcceptFriendRequestResponse {
    return {};
  },

  toJSON(_: AcceptFriendRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<AcceptFriendRequestResponse>, I>>(base?: I): AcceptFriendRequestResponse {
    return AcceptFriendRequestResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AcceptFriendRequestResponse>, I>>(_: I): AcceptFriendRequestResponse {
    const message = createBaseAcceptFriendRequestResponse();
    return message;
  },
};

function createBaseDeleteFriendRequestRequest(): DeleteFriendRequestRequest {
  return { userId: "" };
}

export const DeleteFriendRequestRequest = {
  encode(message: DeleteFriendRequestRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFriendRequestRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFriendRequestRequest();
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

  fromJSON(object: any): DeleteFriendRequestRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: DeleteFriendRequestRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFriendRequestRequest>, I>>(base?: I): DeleteFriendRequestRequest {
    return DeleteFriendRequestRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFriendRequestRequest>, I>>(object: I): DeleteFriendRequestRequest {
    const message = createBaseDeleteFriendRequestRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseDeleteFriendRequestResponse(): DeleteFriendRequestResponse {
  return {};
}

export const DeleteFriendRequestResponse = {
  encode(_: DeleteFriendRequestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFriendRequestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFriendRequestResponse();
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

  fromJSON(_: any): DeleteFriendRequestResponse {
    return {};
  },

  toJSON(_: DeleteFriendRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFriendRequestResponse>, I>>(base?: I): DeleteFriendRequestResponse {
    return DeleteFriendRequestResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFriendRequestResponse>, I>>(_: I): DeleteFriendRequestResponse {
    const message = createBaseDeleteFriendRequestResponse();
    return message;
  },
};

function createBaseListFriendResponse(): ListFriendResponse {
  return { items: [] };
}

export const ListFriendResponse = {
  encode(message: ListFriendResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListFriendResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListFriendResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendResponse>, I>>(base?: I): ListFriendResponse {
    return ListFriendResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendResponse>, I>>(object: I): ListFriendResponse {
    const message = createBaseListFriendResponse();
    message.items = object.items?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteFriendRequest2(): DeleteFriendRequest2 {
  return { userId: "", friendshipId: "" };
}

export const DeleteFriendRequest2 = {
  encode(message: DeleteFriendRequest2, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.friendshipId !== "") {
      writer.uint32(18).string(message.friendshipId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFriendRequest2 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFriendRequest2();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        case 2:
          message.friendshipId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteFriendRequest2 {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      friendshipId: isSet(object.friendshipId) ? String(object.friendshipId) : "",
    };
  },

  toJSON(message: DeleteFriendRequest2): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.friendshipId !== undefined && (obj.friendshipId = message.friendshipId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFriendRequest2>, I>>(base?: I): DeleteFriendRequest2 {
    return DeleteFriendRequest2.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFriendRequest2>, I>>(object: I): DeleteFriendRequest2 {
    const message = createBaseDeleteFriendRequest2();
    message.userId = object.userId ?? "";
    message.friendshipId = object.friendshipId ?? "";
    return message;
  },
};

function createBaseDeleteFriendResponse(): DeleteFriendResponse {
  return {};
}

export const DeleteFriendResponse = {
  encode(_: DeleteFriendResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFriendResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFriendResponse();
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

  fromJSON(_: any): DeleteFriendResponse {
    return {};
  },

  toJSON(_: DeleteFriendResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteFriendResponse>, I>>(base?: I): DeleteFriendResponse {
    return DeleteFriendResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFriendResponse>, I>>(_: I): DeleteFriendResponse {
    const message = createBaseDeleteFriendResponse();
    return message;
  },
};

function createBaseListFriendRequest(): ListFriendRequest {
  return { userId: "" };
}

export const ListFriendRequest = {
  encode(message: ListFriendRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendRequest();
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

  fromJSON(object: any): ListFriendRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: ListFriendRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendRequest>, I>>(base?: I): ListFriendRequest {
    return ListFriendRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendRequest>, I>>(object: I): ListFriendRequest {
    const message = createBaseListFriendRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListFriendRequestsRequest(): ListFriendRequestsRequest {
  return { userId: "" };
}

export const ListFriendRequestsRequest = {
  encode(message: ListFriendRequestsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendRequestsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendRequestsRequest();
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

  fromJSON(object: any): ListFriendRequestsRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: ListFriendRequestsRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendRequestsRequest>, I>>(base?: I): ListFriendRequestsRequest {
    return ListFriendRequestsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendRequestsRequest>, I>>(object: I): ListFriendRequestsRequest {
    const message = createBaseListFriendRequestsRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListFriendRequestsResponse(): ListFriendRequestsResponse {
  return { items: [] };
}

export const ListFriendRequestsResponse = {
  encode(message: ListFriendRequestsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListFriendRequestsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListFriendRequestsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListFriendRequestsResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListFriendRequestsResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListFriendRequestsResponse>, I>>(base?: I): ListFriendRequestsResponse {
    return ListFriendRequestsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListFriendRequestsResponse>, I>>(object: I): ListFriendRequestsResponse {
    const message = createBaseListFriendRequestsResponse();
    message.items = object.items?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateFriendRequestRequest(): CreateFriendRequestRequest {
  return { userId: "" };
}

export const CreateFriendRequestRequest = {
  encode(message: CreateFriendRequestRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFriendRequestRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFriendRequestRequest();
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

  fromJSON(object: any): CreateFriendRequestRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: CreateFriendRequestRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFriendRequestRequest>, I>>(base?: I): CreateFriendRequestRequest {
    return CreateFriendRequestRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFriendRequestRequest>, I>>(object: I): CreateFriendRequestRequest {
    const message = createBaseCreateFriendRequestRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseCreateFriendRequestResponse(): CreateFriendRequestResponse {
  return {};
}

export const CreateFriendRequestResponse = {
  encode(_: CreateFriendRequestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFriendRequestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFriendRequestResponse();
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

  fromJSON(_: any): CreateFriendRequestResponse {
    return {};
  },

  toJSON(_: CreateFriendRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateFriendRequestResponse>, I>>(base?: I): CreateFriendRequestResponse {
    return CreateFriendRequestResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateFriendRequestResponse>, I>>(_: I): CreateFriendRequestResponse {
    const message = createBaseCreateFriendRequestResponse();
    return message;
  },
};

export interface FriendshipService {
  /** Send a friend request */
  CreateFriendRequest(request: CreateFriendRequestRequest): Promise<CreateFriendRequestResponse>;
  /** Accept a friend request */
  AcceptFriendRequest(request: AcceptFriendRequestRequest): Promise<AcceptFriendRequestResponse>;
  /** Get friend list of a specific user */
  ListFriend(request: ListFriendRequest): Promise<ListFriendResponse>;
  /** Get friend requests list of a specific user */
  ListFriendRequests(request: ListFriendRequestsRequest): Promise<ListFriendRequestsResponse>;
  /** Refuse a friend request */
  DeleteFriendRequest(request: DeleteFriendRequestRequest): Promise<DeleteFriendRequestResponse>;
  /** Delete a friend */
  DeleteFriend(request: DeleteFriendRequest2): Promise<DeleteFriendResponse>;
}

export class FriendshipServiceClientImpl implements FriendshipService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.friendship.v1.FriendshipService";
    this.rpc = rpc;
    this.CreateFriendRequest = this.CreateFriendRequest.bind(this);
    this.AcceptFriendRequest = this.AcceptFriendRequest.bind(this);
    this.ListFriend = this.ListFriend.bind(this);
    this.ListFriendRequests = this.ListFriendRequests.bind(this);
    this.DeleteFriendRequest = this.DeleteFriendRequest.bind(this);
    this.DeleteFriend = this.DeleteFriend.bind(this);
  }
  CreateFriendRequest(request: CreateFriendRequestRequest): Promise<CreateFriendRequestResponse> {
    const data = CreateFriendRequestRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateFriendRequest", data);
    return promise.then((data) => CreateFriendRequestResponse.decode(new _m0.Reader(data)));
  }

  AcceptFriendRequest(request: AcceptFriendRequestRequest): Promise<AcceptFriendRequestResponse> {
    const data = AcceptFriendRequestRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AcceptFriendRequest", data);
    return promise.then((data) => AcceptFriendRequestResponse.decode(new _m0.Reader(data)));
  }

  ListFriend(request: ListFriendRequest): Promise<ListFriendResponse> {
    const data = ListFriendRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFriend", data);
    return promise.then((data) => ListFriendResponse.decode(new _m0.Reader(data)));
  }

  ListFriendRequests(request: ListFriendRequestsRequest): Promise<ListFriendRequestsResponse> {
    const data = ListFriendRequestsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListFriendRequests", data);
    return promise.then((data) => ListFriendRequestsResponse.decode(new _m0.Reader(data)));
  }

  DeleteFriendRequest(request: DeleteFriendRequestRequest): Promise<DeleteFriendRequestResponse> {
    const data = DeleteFriendRequestRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteFriendRequest", data);
    return promise.then((data) => DeleteFriendRequestResponse.decode(new _m0.Reader(data)));
  }

  DeleteFriend(request: DeleteFriendRequest2): Promise<DeleteFriendResponse> {
    const data = DeleteFriendRequest2.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteFriend", data);
    return promise.then((data) => DeleteFriendResponse.decode(new _m0.Reader(data)));
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
