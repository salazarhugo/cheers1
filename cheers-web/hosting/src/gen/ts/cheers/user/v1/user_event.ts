/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { User } from "../../type/user/user";

export const protobufPackage = "cheers.user.v1";

export interface UserEvent {
  create?: CreateUser | undefined;
  follow?: FollowUser | undefined;
  update?: UpdateUser | undefined;
}

export interface CreateUser {
  user: User | undefined;
}

export interface UpdateUser {
  user: User | undefined;
}

export interface FollowUser {
  user: User | undefined;
  followedUser: User | undefined;
}

function createBaseUserEvent(): UserEvent {
  return { create: undefined, follow: undefined, update: undefined };
}

export const UserEvent = {
  encode(message: UserEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.create !== undefined) {
      CreateUser.encode(message.create, writer.uint32(10).fork()).ldelim();
    }
    if (message.follow !== undefined) {
      FollowUser.encode(message.follow, writer.uint32(18).fork()).ldelim();
    }
    if (message.update !== undefined) {
      UpdateUser.encode(message.update, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.create = CreateUser.decode(reader, reader.uint32());
          break;
        case 2:
          message.follow = FollowUser.decode(reader, reader.uint32());
          break;
        case 3:
          message.update = UpdateUser.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserEvent {
    return {
      create: isSet(object.create) ? CreateUser.fromJSON(object.create) : undefined,
      follow: isSet(object.follow) ? FollowUser.fromJSON(object.follow) : undefined,
      update: isSet(object.update) ? UpdateUser.fromJSON(object.update) : undefined,
    };
  },

  toJSON(message: UserEvent): unknown {
    const obj: any = {};
    message.create !== undefined && (obj.create = message.create ? CreateUser.toJSON(message.create) : undefined);
    message.follow !== undefined && (obj.follow = message.follow ? FollowUser.toJSON(message.follow) : undefined);
    message.update !== undefined && (obj.update = message.update ? UpdateUser.toJSON(message.update) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserEvent>, I>>(base?: I): UserEvent {
    return UserEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserEvent>, I>>(object: I): UserEvent {
    const message = createBaseUserEvent();
    message.create = (object.create !== undefined && object.create !== null)
      ? CreateUser.fromPartial(object.create)
      : undefined;
    message.follow = (object.follow !== undefined && object.follow !== null)
      ? FollowUser.fromPartial(object.follow)
      : undefined;
    message.update = (object.update !== undefined && object.update !== null)
      ? UpdateUser.fromPartial(object.update)
      : undefined;
    return message;
  },
};

function createBaseCreateUser(): CreateUser {
  return { user: undefined };
}

export const CreateUser = {
  encode(message: CreateUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUser();
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

  fromJSON(object: any): CreateUser {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: CreateUser): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUser>, I>>(base?: I): CreateUser {
    return CreateUser.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateUser>, I>>(object: I): CreateUser {
    const message = createBaseCreateUser();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseUpdateUser(): UpdateUser {
  return { user: undefined };
}

export const UpdateUser = {
  encode(message: UpdateUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUser();
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

  fromJSON(object: any): UpdateUser {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: UpdateUser): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateUser>, I>>(base?: I): UpdateUser {
    return UpdateUser.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUser>, I>>(object: I): UpdateUser {
    const message = createBaseUpdateUser();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseFollowUser(): FollowUser {
  return { user: undefined, followedUser: undefined };
}

export const FollowUser = {
  encode(message: FollowUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    if (message.followedUser !== undefined) {
      User.encode(message.followedUser, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FollowUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFollowUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        case 2:
          message.followedUser = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FollowUser {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      followedUser: isSet(object.followedUser) ? User.fromJSON(object.followedUser) : undefined,
    };
  },

  toJSON(message: FollowUser): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.followedUser !== undefined &&
      (obj.followedUser = message.followedUser ? User.toJSON(message.followedUser) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<FollowUser>, I>>(base?: I): FollowUser {
    return FollowUser.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FollowUser>, I>>(object: I): FollowUser {
    const message = createBaseFollowUser();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.followedUser = (object.followedUser !== undefined && object.followedUser !== null)
      ? User.fromPartial(object.followedUser)
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
