/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "../../../google/protobuf/timestamp";
import { Privacy, privacyFromJSON, privacyToJSON } from "../privacy/privacy";

export const protobufPackage = "cheers.type";

export interface Party {
  id: string;
  name: string;
  description: string;
  address: string;
  privacy: Privacy;
  bannerUrl: string;
  startDate: number;
  endDate: number;
  /** The user_id of the creator */
  hostId: string;
  /** The location name */
  locationName: string;
  /** The time when the party was created. */
  createTime:
    | Date
    | undefined;
  /** The latitude in degrees. It must be in the range [-90.0, +90.0]. */
  latitude: number;
  /** The longitude in degrees. It must be in the range [-180.0, +180.0]. */
  longitude: number;
}

function createBaseParty(): Party {
  return {
    id: "",
    name: "",
    description: "",
    address: "",
    privacy: 0,
    bannerUrl: "",
    startDate: 0,
    endDate: 0,
    hostId: "",
    locationName: "",
    createTime: undefined,
    latitude: 0,
    longitude: 0,
  };
}

export const Party = {
  encode(message: Party, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.privacy !== 0) {
      writer.uint32(48).int32(message.privacy);
    }
    if (message.bannerUrl !== "") {
      writer.uint32(58).string(message.bannerUrl);
    }
    if (message.startDate !== 0) {
      writer.uint32(64).int32(message.startDate);
    }
    if (message.endDate !== 0) {
      writer.uint32(72).int32(message.endDate);
    }
    if (message.hostId !== "") {
      writer.uint32(82).string(message.hostId);
    }
    if (message.locationName !== "") {
      writer.uint32(98).string(message.locationName);
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(106).fork()).ldelim();
    }
    if (message.latitude !== 0) {
      writer.uint32(113).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(121).double(message.longitude);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Party {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParty();
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
          message.description = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 6:
          message.privacy = reader.int32() as any;
          break;
        case 7:
          message.bannerUrl = reader.string();
          break;
        case 8:
          message.startDate = reader.int32();
          break;
        case 9:
          message.endDate = reader.int32();
          break;
        case 10:
          message.hostId = reader.string();
          break;
        case 12:
          message.locationName = reader.string();
          break;
        case 13:
          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 14:
          message.latitude = reader.double();
          break;
        case 15:
          message.longitude = reader.double();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Party {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      address: isSet(object.address) ? String(object.address) : "",
      privacy: isSet(object.privacy) ? privacyFromJSON(object.privacy) : 0,
      bannerUrl: isSet(object.bannerUrl) ? String(object.bannerUrl) : "",
      startDate: isSet(object.startDate) ? Number(object.startDate) : 0,
      endDate: isSet(object.endDate) ? Number(object.endDate) : 0,
      hostId: isSet(object.hostId) ? String(object.hostId) : "",
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
    };
  },

  toJSON(message: Party): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.address !== undefined && (obj.address = message.address);
    message.privacy !== undefined && (obj.privacy = privacyToJSON(message.privacy));
    message.bannerUrl !== undefined && (obj.bannerUrl = message.bannerUrl);
    message.startDate !== undefined && (obj.startDate = Math.round(message.startDate));
    message.endDate !== undefined && (obj.endDate = Math.round(message.endDate));
    message.hostId !== undefined && (obj.hostId = message.hostId);
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.createTime !== undefined && (obj.createTime = message.createTime.toISOString());
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Party>, I>>(object: I): Party {
    const message = createBaseParty();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.address = object.address ?? "";
    message.privacy = object.privacy ?? 0;
    message.bannerUrl = object.bannerUrl ?? "";
    message.startDate = object.startDate ?? 0;
    message.endDate = object.endDate ?? 0;
    message.hostId = object.hostId ?? "";
    message.locationName = object.locationName ?? "";
    message.createTime = object.createTime ?? undefined;
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
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
