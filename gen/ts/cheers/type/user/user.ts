/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "../../../google/protobuf/timestamp";

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
  createTime: Date | undefined;
  registrationTokens: string[];
}

export interface UserItem {
  id: string;
  name: string;
  username: string;
  verified: boolean;
  picture: string;
  hasFollowed: boolean;
  storyState: StoryState;
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
    createTime: undefined,
    registrationTokens: [],
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
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.registrationTokens) {
      writer.uint32(90).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.email = reader.string();
          break;
        case 4:
          message.verified = reader.bool();
          break;
        case 5:
          message.username = reader.string();
          break;
        case 6:
          message.picture = reader.string();
          break;
        case 7:
          message.bio = reader.string();
          break;
        case 8:
          message.website = reader.string();
          break;
        case 12:
          message.birthday = reader.string();
          break;
        case 13:
          message.gender = reader.int32() as any;
          break;
        case 9:
          message.phoneNumber = reader.string();
          break;
        case 10:
          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 11:
          message.registrationTokens.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
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
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      registrationTokens: Array.isArray(object?.registrationTokens)
        ? object.registrationTokens.map((e: any) => String(e))
        : [],
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
    message.createTime !== undefined && (obj.createTime = message.createTime.toISOString());
    if (message.registrationTokens) {
      obj.registrationTokens = message.registrationTokens.map((e) => e);
    } else {
      obj.registrationTokens = [];
    }
    return obj;
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
    message.createTime = object.createTime ?? undefined;
    message.registrationTokens = object.registrationTokens?.map((e) => e) || [];
    return message;
  },
};

function createBaseUserItem(): UserItem {
  return { id: "", name: "", username: "", verified: false, picture: "", hasFollowed: false, storyState: 0 };
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
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.username = reader.string();
          break;
        case 4:
          message.verified = reader.bool();
          break;
        case 5:
          message.picture = reader.string();
          break;
        case 6:
          message.hasFollowed = reader.bool();
          break;
        case 7:
          message.storyState = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
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
    return obj;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
