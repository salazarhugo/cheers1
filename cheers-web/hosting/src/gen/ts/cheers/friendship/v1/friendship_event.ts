/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.friendship.v1";

export interface FriendshipEvent {
  createdFriendRequest?: CreatedFriendRequest | undefined;
  createdFriend?: CreatedFriend | undefined;
  deletedFriendRequest?: DeletedFriendRequest | undefined;
  deletedFriend?: DeletedFriend | undefined;
}

export interface CreatedFriendRequest {
  from: string;
  to: string;
}

export interface DeletedFriendRequest {
  from: string;
  to: string;
}

export interface CreatedFriend {
  from: string;
  to: string;
}

export interface DeletedFriend {
  from: string;
  to: string;
}

function createBaseFriendshipEvent(): FriendshipEvent {
  return {
    createdFriendRequest: undefined,
    createdFriend: undefined,
    deletedFriendRequest: undefined,
    deletedFriend: undefined,
  };
}

export const FriendshipEvent = {
  encode(message: FriendshipEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.createdFriendRequest !== undefined) {
      CreatedFriendRequest.encode(message.createdFriendRequest, writer.uint32(10).fork()).ldelim();
    }
    if (message.createdFriend !== undefined) {
      CreatedFriend.encode(message.createdFriend, writer.uint32(18).fork()).ldelim();
    }
    if (message.deletedFriendRequest !== undefined) {
      DeletedFriendRequest.encode(message.deletedFriendRequest, writer.uint32(26).fork()).ldelim();
    }
    if (message.deletedFriend !== undefined) {
      DeletedFriend.encode(message.deletedFriend, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FriendshipEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFriendshipEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.createdFriendRequest = CreatedFriendRequest.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.createdFriend = CreatedFriend.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.deletedFriendRequest = DeletedFriendRequest.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.deletedFriend = DeletedFriend.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FriendshipEvent {
    return {
      createdFriendRequest: isSet(object.createdFriendRequest)
        ? CreatedFriendRequest.fromJSON(object.createdFriendRequest)
        : undefined,
      createdFriend: isSet(object.createdFriend) ? CreatedFriend.fromJSON(object.createdFriend) : undefined,
      deletedFriendRequest: isSet(object.deletedFriendRequest)
        ? DeletedFriendRequest.fromJSON(object.deletedFriendRequest)
        : undefined,
      deletedFriend: isSet(object.deletedFriend) ? DeletedFriend.fromJSON(object.deletedFriend) : undefined,
    };
  },

  toJSON(message: FriendshipEvent): unknown {
    const obj: any = {};
    message.createdFriendRequest !== undefined && (obj.createdFriendRequest = message.createdFriendRequest
      ? CreatedFriendRequest.toJSON(message.createdFriendRequest)
      : undefined);
    message.createdFriend !== undefined &&
      (obj.createdFriend = message.createdFriend ? CreatedFriend.toJSON(message.createdFriend) : undefined);
    message.deletedFriendRequest !== undefined && (obj.deletedFriendRequest = message.deletedFriendRequest
      ? DeletedFriendRequest.toJSON(message.deletedFriendRequest)
      : undefined);
    message.deletedFriend !== undefined &&
      (obj.deletedFriend = message.deletedFriend ? DeletedFriend.toJSON(message.deletedFriend) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<FriendshipEvent>, I>>(base?: I): FriendshipEvent {
    return FriendshipEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FriendshipEvent>, I>>(object: I): FriendshipEvent {
    const message = createBaseFriendshipEvent();
    message.createdFriendRequest = (object.createdFriendRequest !== undefined && object.createdFriendRequest !== null)
      ? CreatedFriendRequest.fromPartial(object.createdFriendRequest)
      : undefined;
    message.createdFriend = (object.createdFriend !== undefined && object.createdFriend !== null)
      ? CreatedFriend.fromPartial(object.createdFriend)
      : undefined;
    message.deletedFriendRequest = (object.deletedFriendRequest !== undefined && object.deletedFriendRequest !== null)
      ? DeletedFriendRequest.fromPartial(object.deletedFriendRequest)
      : undefined;
    message.deletedFriend = (object.deletedFriend !== undefined && object.deletedFriend !== null)
      ? DeletedFriend.fromPartial(object.deletedFriend)
      : undefined;
    return message;
  },
};

function createBaseCreatedFriendRequest(): CreatedFriendRequest {
  return { from: "", to: "" };
}

export const CreatedFriendRequest = {
  encode(message: CreatedFriendRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.from !== "") {
      writer.uint32(10).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedFriendRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedFriendRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.from = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.to = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatedFriendRequest {
    return { from: isSet(object.from) ? String(object.from) : "", to: isSet(object.to) ? String(object.to) : "" };
  },

  toJSON(message: CreatedFriendRequest): unknown {
    const obj: any = {};
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatedFriendRequest>, I>>(base?: I): CreatedFriendRequest {
    return CreatedFriendRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatedFriendRequest>, I>>(object: I): CreatedFriendRequest {
    const message = createBaseCreatedFriendRequest();
    message.from = object.from ?? "";
    message.to = object.to ?? "";
    return message;
  },
};

function createBaseDeletedFriendRequest(): DeletedFriendRequest {
  return { from: "", to: "" };
}

export const DeletedFriendRequest = {
  encode(message: DeletedFriendRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.from !== "") {
      writer.uint32(10).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedFriendRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedFriendRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.from = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.to = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletedFriendRequest {
    return { from: isSet(object.from) ? String(object.from) : "", to: isSet(object.to) ? String(object.to) : "" };
  },

  toJSON(message: DeletedFriendRequest): unknown {
    const obj: any = {};
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletedFriendRequest>, I>>(base?: I): DeletedFriendRequest {
    return DeletedFriendRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletedFriendRequest>, I>>(object: I): DeletedFriendRequest {
    const message = createBaseDeletedFriendRequest();
    message.from = object.from ?? "";
    message.to = object.to ?? "";
    return message;
  },
};

function createBaseCreatedFriend(): CreatedFriend {
  return { from: "", to: "" };
}

export const CreatedFriend = {
  encode(message: CreatedFriend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.from !== "") {
      writer.uint32(10).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedFriend {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedFriend();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.from = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.to = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatedFriend {
    return { from: isSet(object.from) ? String(object.from) : "", to: isSet(object.to) ? String(object.to) : "" };
  },

  toJSON(message: CreatedFriend): unknown {
    const obj: any = {};
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatedFriend>, I>>(base?: I): CreatedFriend {
    return CreatedFriend.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatedFriend>, I>>(object: I): CreatedFriend {
    const message = createBaseCreatedFriend();
    message.from = object.from ?? "";
    message.to = object.to ?? "";
    return message;
  },
};

function createBaseDeletedFriend(): DeletedFriend {
  return { from: "", to: "" };
}

export const DeletedFriend = {
  encode(message: DeletedFriend, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.from !== "") {
      writer.uint32(10).string(message.from);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedFriend {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedFriend();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.from = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.to = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletedFriend {
    return { from: isSet(object.from) ? String(object.from) : "", to: isSet(object.to) ? String(object.to) : "" };
  },

  toJSON(message: DeletedFriend): unknown {
    const obj: any = {};
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletedFriend>, I>>(base?: I): DeletedFriend {
    return DeletedFriend.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletedFriend>, I>>(object: I): DeletedFriend {
    const message = createBaseDeletedFriend();
    message.from = object.from ?? "";
    message.to = object.to ?? "";
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
