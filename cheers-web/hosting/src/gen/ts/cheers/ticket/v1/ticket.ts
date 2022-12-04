/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.ticket.v1";

export interface Ticket {
  id: string;
  /** Price of the ticket */
  price: number;
  /** Number of available tickets to be sold */
  quantity: number;
  name: string;
  description: string;
  partyId: string;
  paymentIntentId: string;
  userId: string;
  validated: boolean;
}

function createBaseTicket(): Ticket {
  return {
    id: "",
    price: 0,
    quantity: 0,
    name: "",
    description: "",
    partyId: "",
    paymentIntentId: "",
    userId: "",
    validated: false,
  };
}

export const Ticket = {
  encode(message: Ticket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.price !== 0) {
      writer.uint32(16).int64(message.price);
    }
    if (message.quantity !== 0) {
      writer.uint32(24).uint32(message.quantity);
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
    if (message.validated === true) {
      writer.uint32(72).bool(message.validated);
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
          message.id = reader.string();
          break;
        case 2:
          message.price = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.quantity = reader.uint32();
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
        case 9:
          message.validated = reader.bool();
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
      id: isSet(object.id) ? String(object.id) : "",
      price: isSet(object.price) ? Number(object.price) : 0,
      quantity: isSet(object.quantity) ? Number(object.quantity) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      paymentIntentId: isSet(object.paymentIntentId) ? String(object.paymentIntentId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      validated: isSet(object.validated) ? Boolean(object.validated) : false,
    };
  },

  toJSON(message: Ticket): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.price !== undefined && (obj.price = Math.round(message.price));
    message.quantity !== undefined && (obj.quantity = Math.round(message.quantity));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.validated !== undefined && (obj.validated = message.validated);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Ticket>, I>>(object: I): Ticket {
    const message = createBaseTicket();
    message.id = object.id ?? "";
    message.price = object.price ?? 0;
    message.quantity = object.quantity ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.partyId = object.partyId ?? "";
    message.paymentIntentId = object.paymentIntentId ?? "";
    message.userId = object.userId ?? "";
    message.validated = object.validated ?? false;
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
