/* eslint-disable */
// @ts-ignore
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.type";

export enum StoryState {
  EMPTY = 0,
  SEEN = 1,
  NOT_SEEN = 2,
  LOADING = 3,
  UNRECOGNIZED = -1,
}

export function storyStateFromJSON(object: any): StoryState {
  switch (object) {
    case 0:
    case "EMPTY":
      return StoryState.EMPTY;
    case 1:
    case "SEEN":
      return StoryState.SEEN;
    case 2:
    case "NOT_SEEN":
      return StoryState.NOT_SEEN;
    case 3:
    case "LOADING":
      return StoryState.LOADING;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StoryState.UNRECOGNIZED;
  }
}

export function storyStateToJSON(object: StoryState): string {
  switch (object) {
    case StoryState.EMPTY:
      return "EMPTY";
    case StoryState.SEEN:
      return "SEEN";
    case StoryState.NOT_SEEN:
      return "NOT_SEEN";
    case StoryState.LOADING:
      return "LOADING";
    case StoryState.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum Gender {
  MALE = 0,
  FEMALE = 1,
  UNRECOGNIZED = -1,
}

export function genderFromJSON(object: any): Gender {
  switch (object) {
    case 0:
    case "MALE":
      return Gender.MALE;
    case 1:
    case "FEMALE":
      return Gender.FEMALE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Gender.UNRECOGNIZED;
  }
}

export function genderToJSON(object: Gender): string {
  switch (object) {
    case Gender.MALE:
      return "MALE";
    case Gender.FEMALE:
      return "FEMALE";
    case Gender.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface User {
  id: string;
  name: string;
  email: string;
  verified: boolean;
  username: string;
  picture: string;
  bio: string;
  website: string;
  birthday: string;
  gender: Gender;
  phoneNumber: string;
  createTime: number;
  registrationTokens: string[];
  isBusinessAccount: boolean;
  isAdmin: boolean;
  isModerator: boolean;
  banner: string;
}

export interface UserItem {
  id: string;
  name: string;
  username: string;
  verified: boolean;
  picture: string;
  hasFollowed: boolean;
  storyState: StoryState;
  friend: boolean;
  requested: boolean;
  hasRequestedViewer: boolean;
  banner: string;
}

function createBaseUser(): User {
  return {
    id: "",
    name: "",
    email: "",
    verified: false,
    username: "",
    picture: "",
    bio: "",
    website: "",
    birthday: "",
    gender: 0,
    phoneNumber: "",
    createTime: 0,
    registrationTokens: [],
    isBusinessAccount: false,
    isAdmin: false,
    isModerator: false,
    banner: "",
  };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    if (message.verified === true) {
      writer.uint32(32).bool(message.verified);
    }
    if (message.username !== "") {
      writer.uint32(42).string(message.username);
    }
    if (message.picture !== "") {
      writer.uint32(50).string(message.picture);
    }
    if (message.bio !== "") {
      writer.uint32(58).string(message.bio);
    }
    if (message.website !== "") {
      writer.uint32(66).string(message.website);
    }
    if (message.birthday !== "") {
      writer.uint32(98).string(message.birthday);
    }
    if (message.gender !== 0) {
      writer.uint32(104).int32(message.gender);
    }
    if (message.phoneNumber !== "") {
      writer.uint32(74).string(message.phoneNumber);
    }
    if (message.createTime !== 0) {
      writer.uint32(80).int64(message.createTime);
    }
    for (const v of message.registrationTokens) {
      writer.uint32(90).string(v!);
    }
    if (message.isBusinessAccount === true) {
      writer.uint32(112).bool(message.isBusinessAccount);
    }
    if (message.isAdmin === true) {
      writer.uint32(120).bool(message.isAdmin);
    }
    if (message.isModerator === true) {
      writer.uint32(128).bool(message.isModerator);
    }
    if (message.banner !== "") {
      writer.uint32(138).string(message.banner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.email = reader.string();
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.verified = reader.bool();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.username = reader.string();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 7:
          if (tag != 58) {
            break;
          }

          message.bio = reader.string();
          continue;
        case 8:
          if (tag != 66) {
            break;
          }

          message.website = reader.string();
          continue;
        case 12:
          if (tag != 98) {
            break;
          }

          message.birthday = reader.string();
          continue;
        case 13:
          if (tag != 104) {
            break;
          }

          message.gender = reader.int32() as any;
          continue;
        case 9:
          if (tag != 74) {
            break;
          }

          message.phoneNumber = reader.string();
          continue;
        case 10:
          if (tag != 80) {
            break;
          }

          message.createTime = longToNumber(reader.int64() as Long);
          continue;
        case 11:
          if (tag != 90) {
            break;
          }

          message.registrationTokens.push(reader.string());
          continue;
        case 14:
          if (tag != 112) {
            break;
          }

          message.isBusinessAccount = reader.bool();
          continue;
        case 15:
          if (tag != 120) {
            break;
          }

          message.isAdmin = reader.bool();
          continue;
        case 16:
          if (tag != 128) {
            break;
          }

          message.isModerator = reader.bool();
          continue;
        case 17:
          if (tag != 138) {
            break;
          }

          message.banner = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      email: isSet(object.email) ? String(object.email) : "",
      verified: isSet(object.verified) ? Boolean(object.verified) : false,
      username: isSet(object.username) ? String(object.username) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      bio: isSet(object.bio) ? String(object.bio) : "",
      website: isSet(object.website) ? String(object.website) : "",
      birthday: isSet(object.birthday) ? String(object.birthday) : "",
      gender: isSet(object.gender) ? genderFromJSON(object.gender) : 0,
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
      createTime: isSet(object.createTime) ? Number(object.createTime) : 0,
      registrationTokens: Array.isArray(object?.registrationTokens)
        ? object.registrationTokens.map((e: any) => String(e))
        : [],
      isBusinessAccount: isSet(object.isBusinessAccount) ? Boolean(object.isBusinessAccount) : false,
      isAdmin: isSet(object.isAdmin) ? Boolean(object.isAdmin) : false,
      isModerator: isSet(object.isModerator) ? Boolean(object.isModerator) : false,
      banner: isSet(object.banner) ? String(object.banner) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.email !== undefined && (obj.email = message.email);
    message.verified !== undefined && (obj.verified = message.verified);
    message.username !== undefined && (obj.username = message.username);
    message.picture !== undefined && (obj.picture = message.picture);
    message.bio !== undefined && (obj.bio = message.bio);
    message.website !== undefined && (obj.website = message.website);
    message.birthday !== undefined && (obj.birthday = message.birthday);
    message.gender !== undefined && (obj.gender = genderToJSON(message.gender));
    message.phoneNumber !== undefined && (obj.phoneNumber = message.phoneNumber);
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    if (message.registrationTokens) {
      obj.registrationTokens = message.registrationTokens.map((e) => e);
    } else {
      obj.registrationTokens = [];
    }
    message.isBusinessAccount !== undefined && (obj.isBusinessAccount = message.isBusinessAccount);
    message.isAdmin !== undefined && (obj.isAdmin = message.isAdmin);
    message.isModerator !== undefined && (obj.isModerator = message.isModerator);
    message.banner !== undefined && (obj.banner = message.banner);
    return obj;
  },

  create<I extends Exact<DeepPartial<User>, I>>(base?: I): User {
    return User.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.email = object.email ?? "";
    message.verified = object.verified ?? false;
    message.username = object.username ?? "";
    message.picture = object.picture ?? "";
    message.bio = object.bio ?? "";
    message.website = object.website ?? "";
    message.birthday = object.birthday ?? "";
    message.gender = object.gender ?? 0;
    message.phoneNumber = object.phoneNumber ?? "";
    message.createTime = object.createTime ?? 0;
    message.registrationTokens = object.registrationTokens?.map((e) => e) || [];
    message.isBusinessAccount = object.isBusinessAccount ?? false;
    message.isAdmin = object.isAdmin ?? false;
    message.isModerator = object.isModerator ?? false;
    message.banner = object.banner ?? "";
    return message;
  },
};

function createBaseUserItem(): UserItem {
  return {
    id: "",
    name: "",
    username: "",
    verified: false,
    picture: "",
    hasFollowed: false,
    storyState: 0,
    friend: false,
    requested: false,
    hasRequestedViewer: false,
    banner: "",
  };
}

export const UserItem = {
  encode(message: UserItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.username !== "") {
      writer.uint32(26).string(message.username);
    }
    if (message.verified === true) {
      writer.uint32(32).bool(message.verified);
    }
    if (message.picture !== "") {
      writer.uint32(42).string(message.picture);
    }
    if (message.hasFollowed === true) {
      writer.uint32(48).bool(message.hasFollowed);
    }
    if (message.storyState !== 0) {
      writer.uint32(56).int32(message.storyState);
    }
    if (message.friend === true) {
      writer.uint32(64).bool(message.friend);
    }
    if (message.requested === true) {
      writer.uint32(72).bool(message.requested);
    }
    if (message.hasRequestedViewer === true) {
      writer.uint32(80).bool(message.hasRequestedViewer);
    }
    if (message.banner !== "") {
      writer.uint32(90).string(message.banner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserItem {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.username = reader.string();
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.verified = reader.bool();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.picture = reader.string();
          continue;
        case 6:
          if (tag != 48) {
            break;
          }

          message.hasFollowed = reader.bool();
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.storyState = reader.int32() as any;
          continue;
        case 8:
          if (tag != 64) {
            break;
          }

          message.friend = reader.bool();
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.requested = reader.bool();
          continue;
        case 10:
          if (tag != 80) {
            break;
          }

          message.hasRequestedViewer = reader.bool();
          continue;
        case 11:
          if (tag != 90) {
            break;
          }

          message.banner = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserItem {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      username: isSet(object.username) ? String(object.username) : "",
      verified: isSet(object.verified) ? Boolean(object.verified) : false,
      picture: isSet(object.picture) ? String(object.picture) : "",
      hasFollowed: isSet(object.hasFollowed) ? Boolean(object.hasFollowed) : false,
      storyState: isSet(object.storyState) ? storyStateFromJSON(object.storyState) : 0,
      friend: isSet(object.friend) ? Boolean(object.friend) : false,
      requested: isSet(object.requested) ? Boolean(object.requested) : false,
      hasRequestedViewer: isSet(object.hasRequestedViewer) ? Boolean(object.hasRequestedViewer) : false,
      banner: isSet(object.banner) ? String(object.banner) : "",
    };
  },

  toJSON(message: UserItem): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.username !== undefined && (obj.username = message.username);
    message.verified !== undefined && (obj.verified = message.verified);
    message.picture !== undefined && (obj.picture = message.picture);
    message.hasFollowed !== undefined && (obj.hasFollowed = message.hasFollowed);
    message.storyState !== undefined && (obj.storyState = storyStateToJSON(message.storyState));
    message.friend !== undefined && (obj.friend = message.friend);
    message.requested !== undefined && (obj.requested = message.requested);
    message.hasRequestedViewer !== undefined && (obj.hasRequestedViewer = message.hasRequestedViewer);
    message.banner !== undefined && (obj.banner = message.banner);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserItem>, I>>(base?: I): UserItem {
    return UserItem.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserItem>, I>>(object: I): UserItem {
    const message = createBaseUserItem();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.username = object.username ?? "";
    message.verified = object.verified ?? false;
    message.picture = object.picture ?? "";
    message.hasFollowed = object.hasFollowed ?? false;
    message.storyState = object.storyState ?? 0;
    message.friend = object.friend ?? false;
    message.requested = object.requested ?? false;
    message.hasRequestedViewer = object.hasRequestedViewer ?? false;
    message.banner = object.banner ?? "";
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
