/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Privacy, privacyFromJSON, privacyToJSON } from "../privacy/privacy";

export const protobufPackage = "cheers.type";

export interface Story {
  id: string;
  address: string;
  privacy: Privacy;
  photo: string;
  isReactionEnabled: boolean;
  locationName: string;
  type: Story_StoryType;
  createTime: number;
  shareEnabled: boolean;
}

export enum Story_StoryType {
  IMAGE = 0,
  VIDEO = 1,
  UNRECOGNIZED = -1,
}

export function story_StoryTypeFromJSON(object: any): Story_StoryType {
  switch (object) {
    case 0:
    case "IMAGE":
      return Story_StoryType.IMAGE;
    case 1:
    case "VIDEO":
      return Story_StoryType.VIDEO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Story_StoryType.UNRECOGNIZED;
  }
}

export function story_StoryTypeToJSON(object: Story_StoryType): string {
  switch (object) {
    case Story_StoryType.IMAGE:
      return "IMAGE";
    case Story_StoryType.VIDEO:
      return "VIDEO";
    case Story_StoryType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseStory(): Story {
  return {
    id: "",
    address: "",
    privacy: 0,
    photo: "",
    isReactionEnabled: false,
    locationName: "",
    type: 0,
    createTime: 0,
    shareEnabled: false,
  };
}

export const Story = {
  encode(message: Story, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.privacy !== 0) {
      writer.uint32(40).int32(message.privacy);
    }
    if (message.photo !== "") {
      writer.uint32(50).string(message.photo);
    }
    if (message.isReactionEnabled === true) {
      writer.uint32(56).bool(message.isReactionEnabled);
    }
    if (message.locationName !== "") {
      writer.uint32(66).string(message.locationName);
    }
    if (message.type !== 0) {
      writer.uint32(88).int32(message.type);
    }
    if (message.createTime !== 0) {
      writer.uint32(96).int64(message.createTime);
    }
    if (message.shareEnabled === true) {
      writer.uint32(112).bool(message.shareEnabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Story {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStory();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.privacy = reader.int32() as any;
          break;
        case 6:
          message.photo = reader.string();
          break;
        case 7:
          message.isReactionEnabled = reader.bool();
          break;
        case 8:
          message.locationName = reader.string();
          break;
        case 11:
          message.type = reader.int32() as any;
          break;
        case 12:
          message.createTime = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.shareEnabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Story {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      address: isSet(object.address) ? String(object.address) : "",
      privacy: isSet(object.privacy) ? privacyFromJSON(object.privacy) : 0,
      photo: isSet(object.photo) ? String(object.photo) : "",
      isReactionEnabled: isSet(object.isReactionEnabled) ? Boolean(object.isReactionEnabled) : false,
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      type: isSet(object.type) ? story_StoryTypeFromJSON(object.type) : 0,
      createTime: isSet(object.createTime) ? Number(object.createTime) : 0,
      shareEnabled: isSet(object.shareEnabled) ? Boolean(object.shareEnabled) : false,
    };
  },

  toJSON(message: Story): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.privacy !== undefined && (obj.privacy = privacyToJSON(message.privacy));
    message.photo !== undefined && (obj.photo = message.photo);
    message.isReactionEnabled !== undefined && (obj.isReactionEnabled = message.isReactionEnabled);
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.type !== undefined && (obj.type = story_StoryTypeToJSON(message.type));
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    message.shareEnabled !== undefined && (obj.shareEnabled = message.shareEnabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Story>, I>>(object: I): Story {
    const message = createBaseStory();
    message.id = object.id ?? "";
    message.address = object.address ?? "";
    message.privacy = object.privacy ?? 0;
    message.photo = object.photo ?? "";
    message.isReactionEnabled = object.isReactionEnabled ?? false;
    message.locationName = object.locationName ?? "";
    message.type = object.type ?? 0;
    message.createTime = object.createTime ?? 0;
    message.shareEnabled = object.shareEnabled ?? false;
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
