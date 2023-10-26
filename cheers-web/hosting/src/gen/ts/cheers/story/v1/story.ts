/* eslint-disable */
// @ts-ignore
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Privacy, privacyFromJSON, privacyToJSON } from "../../type/privacy/privacy";

export const protobufPackage = "cheers.story.v1";

export interface Story {
  id: string;
  photo: string;
  address: string;
  privacy: Privacy;
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
    photo: "",
    address: "",
    privacy: 0,
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
    if (message.photo !== "") {
      writer.uint32(18).string(message.photo);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.privacy !== 0) {
      writer.uint32(32).int32(message.privacy);
    }
    if (message.isReactionEnabled === true) {
      writer.uint32(40).bool(message.isReactionEnabled);
    }
    if (message.locationName !== "") {
      writer.uint32(50).string(message.locationName);
    }
    if (message.type !== 0) {
      writer.uint32(56).int32(message.type);
    }
    if (message.createTime !== 0) {
      writer.uint32(64).int64(message.createTime);
    }
    if (message.shareEnabled === true) {
      writer.uint32(72).bool(message.shareEnabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Story {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStory();
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

          message.photo = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.address = reader.string();
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.privacy = reader.int32() as any;
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.isReactionEnabled = reader.bool();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.locationName = reader.string();
          continue;
        case 7:
          if (tag != 56) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 8:
          if (tag != 64) {
            break;
          }

          message.createTime = longToNumber(reader.int64() as Long);
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.shareEnabled = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Story {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      photo: isSet(object.photo) ? String(object.photo) : "",
      address: isSet(object.address) ? String(object.address) : "",
      privacy: isSet(object.privacy) ? privacyFromJSON(object.privacy) : 0,
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
    message.photo !== undefined && (obj.photo = message.photo);
    message.address !== undefined && (obj.address = message.address);
    message.privacy !== undefined && (obj.privacy = privacyToJSON(message.privacy));
    message.isReactionEnabled !== undefined && (obj.isReactionEnabled = message.isReactionEnabled);
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.type !== undefined && (obj.type = story_StoryTypeToJSON(message.type));
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    message.shareEnabled !== undefined && (obj.shareEnabled = message.shareEnabled);
    return obj;
  },

  create<I extends Exact<DeepPartial<Story>, I>>(base?: I): Story {
    return Story.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Story>, I>>(object: I): Story {
    const message = createBaseStory();
    message.id = object.id ?? "";
    message.photo = object.photo ?? "";
    message.address = object.address ?? "";
    message.privacy = object.privacy ?? 0;
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
