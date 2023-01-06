/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";
import { Friendship } from "./friendship_service";

export const protobufPackage = "cheers.friendship.v1";

export interface FriendshipEvent {
  created?: CreatedFriendship | undefined;
  deleted?: DeletedFriendship | undefined;
}

export interface CreatedFriendship {
  friendship: Friendship | undefined;
  user: UserItem | undefined;
}

export interface DeletedFriendship {
  friendship: Friendship | undefined;
  user: UserItem | undefined;
}

function createBaseFriendshipEvent(): FriendshipEvent {
  return { created: undefined, deleted: undefined };
}

export const FriendshipEvent = {
  encode(message: FriendshipEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.created !== undefined) {
      CreatedFriendship.encode(message.created, writer.uint32(10).fork()).ldelim();
    }
    if (message.deleted !== undefined) {
      DeletedFriendship.encode(message.deleted, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FriendshipEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFriendshipEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.created = CreatedFriendship.decode(reader, reader.uint32());
          break;
        case 2:
          message.deleted = DeletedFriendship.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FriendshipEvent {
    return {
      created: isSet(object.created) ? CreatedFriendship.fromJSON(object.created) : undefined,
      deleted: isSet(object.deleted) ? DeletedFriendship.fromJSON(object.deleted) : undefined,
    };
  },

  toJSON(message: FriendshipEvent): unknown {
    const obj: any = {};
    message.created !== undefined &&
      (obj.created = message.created ? CreatedFriendship.toJSON(message.created) : undefined);
    message.deleted !== undefined &&
      (obj.deleted = message.deleted ? DeletedFriendship.toJSON(message.deleted) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FriendshipEvent>, I>>(object: I): FriendshipEvent {
    const message = createBaseFriendshipEvent();
    message.created = (object.created !== undefined && object.created !== null)
      ? CreatedFriendship.fromPartial(object.created)
      : undefined;
    message.deleted = (object.deleted !== undefined && object.deleted !== null)
      ? DeletedFriendship.fromPartial(object.deleted)
      : undefined;
    return message;
  },
};

function createBaseCreatedFriendship(): CreatedFriendship {
  return { friendship: undefined, user: undefined };
}

export const CreatedFriendship = {
  encode(message: CreatedFriendship, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.friendship !== undefined) {
      Friendship.encode(message.friendship, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatedFriendship {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatedFriendship();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.friendship = Friendship.decode(reader, reader.uint32());
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

  fromJSON(object: any): CreatedFriendship {
    return {
      friendship: isSet(object.friendship) ? Friendship.fromJSON(object.friendship) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreatedFriendship): unknown {
    const obj: any = {};
    message.friendship !== undefined &&
      (obj.friendship = message.friendship ? Friendship.toJSON(message.friendship) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreatedFriendship>, I>>(object: I): CreatedFriendship {
    const message = createBaseCreatedFriendship();
    message.friendship = (object.friendship !== undefined && object.friendship !== null)
      ? Friendship.fromPartial(object.friendship)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseDeletedFriendship(): DeletedFriendship {
  return { friendship: undefined, user: undefined };
}

export const DeletedFriendship = {
  encode(message: DeletedFriendship, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.friendship !== undefined) {
      Friendship.encode(message.friendship, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletedFriendship {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletedFriendship();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.friendship = Friendship.decode(reader, reader.uint32());
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

  fromJSON(object: any): DeletedFriendship {
    return {
      friendship: isSet(object.friendship) ? Friendship.fromJSON(object.friendship) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: DeletedFriendship): unknown {
    const obj: any = {};
    message.friendship !== undefined &&
      (obj.friendship = message.friendship ? Friendship.toJSON(message.friendship) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletedFriendship>, I>>(object: I): DeletedFriendship {
    const message = createBaseDeletedFriendship();
    message.friendship = (object.friendship !== undefined && object.friendship !== null)
      ? Friendship.fromPartial(object.friendship)
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
