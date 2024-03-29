/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "../../../google/protobuf/timestamp";
import { LatLng } from "../../../google/type/latlng";
import { Privacy, privacyFromJSON, privacyToJSON } from "../privacy/privacy";

export const protobufPackage = "cheers.type";

export interface Party {
  id: string;
  name: string;
  description: string;
  address: string;
  latlng: LatLng | undefined;
  privacy: Privacy;
  bannerUrl: string;
  startDate: Date | undefined;
  endDate: Date | undefined;
  hostId: string;
  locationName: string;
  /** The time when the party was created. */
  createTime: Date | undefined;
}

function createBaseParty(): Party {
  return {
    id: "",
    name: "",
    description: "",
    address: "",
    latlng: undefined,
    privacy: 0,
    bannerUrl: "",
    startDate: undefined,
    endDate: undefined,
    hostId: "",
    locationName: "",
    createTime: undefined,
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
    if (message.latlng !== undefined) {
      LatLng.encode(message.latlng, writer.uint32(42).fork()).ldelim();
    }
    if (message.privacy !== 0) {
      writer.uint32(48).int32(message.privacy);
    }
    if (message.bannerUrl !== "") {
      writer.uint32(58).string(message.bannerUrl);
    }
    if (message.startDate !== undefined) {
      Timestamp.encode(toTimestamp(message.startDate), writer.uint32(66).fork()).ldelim();
    }
    if (message.endDate !== undefined) {
      Timestamp.encode(toTimestamp(message.endDate), writer.uint32(74).fork()).ldelim();
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
        case 5:
          message.latlng = LatLng.decode(reader, reader.uint32());
          break;
        case 6:
          message.privacy = reader.int32() as any;
          break;
        case 7:
          message.bannerUrl = reader.string();
          break;
        case 8:
          message.startDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 9:
          message.endDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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
      latlng: isSet(object.latlng) ? LatLng.fromJSON(object.latlng) : undefined,
      privacy: isSet(object.privacy) ? privacyFromJSON(object.privacy) : 0,
      bannerUrl: isSet(object.bannerUrl) ? String(object.bannerUrl) : "",
      startDate: isSet(object.startDate) ? fromJsonTimestamp(object.startDate) : undefined,
      endDate: isSet(object.endDate) ? fromJsonTimestamp(object.endDate) : undefined,
      hostId: isSet(object.hostId) ? String(object.hostId) : "",
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
    };
  },

  toJSON(message: Party): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.address !== undefined && (obj.address = message.address);
    message.latlng !== undefined && (obj.latlng = message.latlng ? LatLng.toJSON(message.latlng) : undefined);
    message.privacy !== undefined && (obj.privacy = privacyToJSON(message.privacy));
    message.bannerUrl !== undefined && (obj.bannerUrl = message.bannerUrl);
    message.startDate !== undefined && (obj.startDate = message.startDate.toISOString());
    message.endDate !== undefined && (obj.endDate = message.endDate.toISOString());
    message.hostId !== undefined && (obj.hostId = message.hostId);
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.createTime !== undefined && (obj.createTime = message.createTime.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Party>, I>>(object: I): Party {
    const message = createBaseParty();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.address = object.address ?? "";
    message.latlng = (object.latlng !== undefined && object.latlng !== null)
      ? LatLng.fromPartial(object.latlng)
      : undefined;
    message.privacy = object.privacy ?? 0;
    message.bannerUrl = object.bannerUrl ?? "";
    message.startDate = object.startDate ?? undefined;
    message.endDate = object.endDate ?? undefined;
    message.hostId = object.hostId ?? "";
    message.locationName = object.locationName ?? "";
    message.createTime = object.createTime ?? undefined;
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
