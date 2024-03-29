/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.type";

export interface Ticket {
  validated: boolean;
  price: number;
  id: string;
  name: string;
  description: string;
  partyId: string;
  paymentIntentId: string;
  userId: string;
}

function createBaseTicket(): Ticket {
  return {
    validated: false,
    price: 0,
    id: "",
    name: "",
    description: "",
    partyId: "",
    paymentIntentId: "",
    userId: "",
  };
}

export const Ticket = {
  encode(message: Ticket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validated === true) {
      writer.uint32(8).bool(message.validated);
    }
    if (message.price !== 0) {
      writer.uint32(16).int64(message.price);
    }
    if (message.id !== "") {
      writer.uint32(26).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.partyId !== "") {
      writer.uint32(50).string(message.partyId);
    }
    if (message.paymentIntentId !== "") {
      writer.uint32(58).string(message.paymentIntentId);
    }
    if (message.userId !== "") {
      writer.uint32(66).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Ticket {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTicket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validated = reader.bool();
          break;
        case 2:
          message.price = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.id = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.partyId = reader.string();
          break;
        case 7:
          message.paymentIntentId = reader.string();
          break;
        case 8:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ticket {
    return {
      validated: isSet(object.validated) ? Boolean(object.validated) : false,
      price: isSet(object.price) ? Number(object.price) : 0,
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      paymentIntentId: isSet(object.paymentIntentId) ? String(object.paymentIntentId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: Ticket): unknown {
    const obj: any = {};
    message.validated !== undefined && (obj.validated = message.validated);
    message.price !== undefined && (obj.price = Math.round(message.price));
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Ticket>, I>>(object: I): Ticket {
    const message = createBaseTicket();
    message.validated = object.validated ?? false;
    message.price = object.price ?? 0;
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.partyId = object.partyId ?? "";
    message.paymentIntentId = object.paymentIntentId ?? "";
    message.userId = object.userId ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
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
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
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
