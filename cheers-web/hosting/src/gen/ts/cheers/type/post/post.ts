/* eslint-disable */
// @ts-ignore
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Privacy, privacyFromJSON, privacyToJSON } from "../privacy/privacy";

export const protobufPackage = "cheers.type";

export enum PostRatio {
  RATIO_16_9 = 0,
  RATIO_1_1 = 1,
  RATIO_4_5 = 2,
  UNRECOGNIZED = -1,
}

export function postRatioFromJSON(object: any): PostRatio {
  switch (object) {
    case 0:
    case "RATIO_16_9":
      return PostRatio.RATIO_16_9;
    case 1:
    case "RATIO_1_1":
      return PostRatio.RATIO_1_1;
    case 2:
    case "RATIO_4_5":
      return PostRatio.RATIO_4_5;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PostRatio.UNRECOGNIZED;
  }
}

export function postRatioToJSON(object: PostRatio): string {
  switch (object) {
    case PostRatio.RATIO_16_9:
      return "RATIO_16_9";
    case PostRatio.RATIO_1_1:
      return "RATIO_1_1";
    case PostRatio.RATIO_4_5:
      return "RATIO_4_5";
    case PostRatio.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum PostType {
  TEXT = 0,
  IMAGE = 1,
  VIDEO = 2,
  UNRECOGNIZED = -1,
}

export function postTypeFromJSON(object: any): PostType {
  switch (object) {
    case 0:
    case "TEXT":
      return PostType.TEXT;
    case 1:
    case "IMAGE":
      return PostType.IMAGE;
    case 2:
    case "VIDEO":
      return PostType.VIDEO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PostType.UNRECOGNIZED;
  }
}

export function postTypeToJSON(object: PostType): string {
  switch (object) {
    case PostType.TEXT:
      return "TEXT";
    case PostType.IMAGE:
      return "IMAGE";
    case PostType.VIDEO:
      return "VIDEO";
    case PostType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Post {
  id: string;
  creatorId: string;
  caption: string;
  address: string;
  privacy: Privacy;
  photos: string[];
  locationName: string;
  drink: string;
  drunkenness: number;
  type: PostType;
  createTime: number;
  canComment: boolean;
  canShare: boolean;
  ratio: PostRatio;
  /** The latitude in degrees. It must be in the range [-90.0, +90.0]. */
  latitude: number;
  /** The longitude in degrees. It must be in the range [-180.0, +180.0]. */
  longitude: number;
  lastCommentText: string;
  lastCommentUsername: string;
  lastCommentCreateTime: number;
}

function createBasePost(): Post {
  return {
    id: "",
    creatorId: "",
    caption: "",
    address: "",
    privacy: 0,
    photos: [],
    locationName: "",
    drink: "",
    drunkenness: 0,
    type: 0,
    createTime: 0,
    canComment: false,
    canShare: false,
    ratio: 0,
    latitude: 0,
    longitude: 0,
    lastCommentText: "",
    lastCommentUsername: "",
    lastCommentCreateTime: 0,
  };
}

export const Post = {
  encode(message: Post, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creatorId !== "") {
      writer.uint32(18).string(message.creatorId);
    }
    if (message.caption !== "") {
      writer.uint32(26).string(message.caption);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.privacy !== 0) {
      writer.uint32(40).int32(message.privacy);
    }
    for (const v of message.photos) {
      writer.uint32(50).string(v!);
    }
    if (message.locationName !== "") {
      writer.uint32(66).string(message.locationName);
    }
    if (message.drink !== "") {
      writer.uint32(74).string(message.drink);
    }
    if (message.drunkenness !== 0) {
      writer.uint32(80).int64(message.drunkenness);
    }
    if (message.type !== 0) {
      writer.uint32(88).int32(message.type);
    }
    if (message.createTime !== 0) {
      writer.uint32(96).int64(message.createTime);
    }
    if (message.canComment === true) {
      writer.uint32(104).bool(message.canComment);
    }
    if (message.canShare === true) {
      writer.uint32(112).bool(message.canShare);
    }
    if (message.ratio !== 0) {
      writer.uint32(120).int32(message.ratio);
    }
    if (message.latitude !== 0) {
      writer.uint32(129).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(137).double(message.longitude);
    }
    if (message.lastCommentText !== "") {
      writer.uint32(146).string(message.lastCommentText);
    }
    if (message.lastCommentUsername !== "") {
      writer.uint32(154).string(message.lastCommentUsername);
    }
    if (message.lastCommentCreateTime !== 0) {
      writer.uint32(160).int64(message.lastCommentCreateTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Post {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePost();
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

          message.creatorId = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.caption = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.address = reader.string();
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.privacy = reader.int32() as any;
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.photos.push(reader.string());
          continue;
        case 8:
          if (tag != 66) {
            break;
          }

          message.locationName = reader.string();
          continue;
        case 9:
          if (tag != 74) {
            break;
          }

          message.drink = reader.string();
          continue;
        case 10:
          if (tag != 80) {
            break;
          }

          message.drunkenness = longToNumber(reader.int64() as Long);
          continue;
        case 11:
          if (tag != 88) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 12:
          if (tag != 96) {
            break;
          }

          message.createTime = longToNumber(reader.int64() as Long);
          continue;
        case 13:
          if (tag != 104) {
            break;
          }

          message.canComment = reader.bool();
          continue;
        case 14:
          if (tag != 112) {
            break;
          }

          message.canShare = reader.bool();
          continue;
        case 15:
          if (tag != 120) {
            break;
          }

          message.ratio = reader.int32() as any;
          continue;
        case 16:
          if (tag != 129) {
            break;
          }

          message.latitude = reader.double();
          continue;
        case 17:
          if (tag != 137) {
            break;
          }

          message.longitude = reader.double();
          continue;
        case 18:
          if (tag != 146) {
            break;
          }

          message.lastCommentText = reader.string();
          continue;
        case 19:
          if (tag != 154) {
            break;
          }

          message.lastCommentUsername = reader.string();
          continue;
        case 20:
          if (tag != 160) {
            break;
          }

          message.lastCommentCreateTime = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Post {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      creatorId: isSet(object.creatorId) ? String(object.creatorId) : "",
      caption: isSet(object.caption) ? String(object.caption) : "",
      address: isSet(object.address) ? String(object.address) : "",
      privacy: isSet(object.privacy) ? privacyFromJSON(object.privacy) : 0,
      photos: Array.isArray(object?.photos) ? object.photos.map((e: any) => String(e)) : [],
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      drink: isSet(object.drink) ? String(object.drink) : "",
      drunkenness: isSet(object.drunkenness) ? Number(object.drunkenness) : 0,
      type: isSet(object.type) ? postTypeFromJSON(object.type) : 0,
      createTime: isSet(object.createTime) ? Number(object.createTime) : 0,
      canComment: isSet(object.canComment) ? Boolean(object.canComment) : false,
      canShare: isSet(object.canShare) ? Boolean(object.canShare) : false,
      ratio: isSet(object.ratio) ? postRatioFromJSON(object.ratio) : 0,
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
      lastCommentText: isSet(object.lastCommentText) ? String(object.lastCommentText) : "",
      lastCommentUsername: isSet(object.lastCommentUsername) ? String(object.lastCommentUsername) : "",
      lastCommentCreateTime: isSet(object.lastCommentCreateTime) ? Number(object.lastCommentCreateTime) : 0,
    };
  },

  toJSON(message: Post): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creatorId !== undefined && (obj.creatorId = message.creatorId);
    message.caption !== undefined && (obj.caption = message.caption);
    message.address !== undefined && (obj.address = message.address);
    message.privacy !== undefined && (obj.privacy = privacyToJSON(message.privacy));
    if (message.photos) {
      obj.photos = message.photos.map((e) => e);
    } else {
      obj.photos = [];
    }
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.drink !== undefined && (obj.drink = message.drink);
    message.drunkenness !== undefined && (obj.drunkenness = Math.round(message.drunkenness));
    message.type !== undefined && (obj.type = postTypeToJSON(message.type));
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    message.canComment !== undefined && (obj.canComment = message.canComment);
    message.canShare !== undefined && (obj.canShare = message.canShare);
    message.ratio !== undefined && (obj.ratio = postRatioToJSON(message.ratio));
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    message.lastCommentText !== undefined && (obj.lastCommentText = message.lastCommentText);
    message.lastCommentUsername !== undefined && (obj.lastCommentUsername = message.lastCommentUsername);
    message.lastCommentCreateTime !== undefined &&
      (obj.lastCommentCreateTime = Math.round(message.lastCommentCreateTime));
    return obj;
  },

  create<I extends Exact<DeepPartial<Post>, I>>(base?: I): Post {
    return Post.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Post>, I>>(object: I): Post {
    const message = createBasePost();
    message.id = object.id ?? "";
    message.creatorId = object.creatorId ?? "";
    message.caption = object.caption ?? "";
    message.address = object.address ?? "";
    message.privacy = object.privacy ?? 0;
    message.photos = object.photos?.map((e) => e) || [];
    message.locationName = object.locationName ?? "";
    message.drink = object.drink ?? "";
    message.drunkenness = object.drunkenness ?? 0;
    message.type = object.type ?? 0;
    message.createTime = object.createTime ?? 0;
    message.canComment = object.canComment ?? false;
    message.canShare = object.canShare ?? false;
    message.ratio = object.ratio ?? 0;
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
    message.lastCommentText = object.lastCommentText ?? "";
    message.lastCommentUsername = object.lastCommentUsername ?? "";
    message.lastCommentCreateTime = object.lastCommentCreateTime ?? 0;
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
