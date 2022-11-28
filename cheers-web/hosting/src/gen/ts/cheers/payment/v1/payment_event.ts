/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.payment.v1";

export interface PaymentEvent {
  paymentIntentId: string;
  customerId: string;
  type: PaymentEvent_EventType;
  firstName: string;
  lastName: string;
  email: string;
  partyId: string;
  partyName: string;
}

export enum PaymentEvent_EventType {
  PAYMENT_SUCCESS = 0,
  REFUND = 1,
  UNRECOGNIZED = -1,
}

export function paymentEvent_EventTypeFromJSON(object: any): PaymentEvent_EventType {
  switch (object) {
    case 0:
    case "PAYMENT_SUCCESS":
      return PaymentEvent_EventType.PAYMENT_SUCCESS;
    case 1:
    case "REFUND":
      return PaymentEvent_EventType.REFUND;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PaymentEvent_EventType.UNRECOGNIZED;
  }
}

export function paymentEvent_EventTypeToJSON(object: PaymentEvent_EventType): string {
  switch (object) {
    case PaymentEvent_EventType.PAYMENT_SUCCESS:
      return "PAYMENT_SUCCESS";
    case PaymentEvent_EventType.REFUND:
      return "REFUND";
    case PaymentEvent_EventType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBasePaymentEvent(): PaymentEvent {
  return {
    paymentIntentId: "",
    customerId: "",
    type: 0,
    firstName: "",
    lastName: "",
    email: "",
    partyId: "",
    partyName: "",
  };
}

export const PaymentEvent = {
  encode(message: PaymentEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.paymentIntentId !== "") {
      writer.uint32(10).string(message.paymentIntentId);
    }
    if (message.customerId !== "") {
      writer.uint32(18).string(message.customerId);
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    if (message.firstName !== "") {
      writer.uint32(34).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(42).string(message.lastName);
    }
    if (message.email !== "") {
      writer.uint32(50).string(message.email);
    }
    if (message.partyId !== "") {
      writer.uint32(58).string(message.partyId);
    }
    if (message.partyName !== "") {
      writer.uint32(66).string(message.partyName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PaymentEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePaymentEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.paymentIntentId = reader.string();
          break;
        case 2:
          message.customerId = reader.string();
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        case 4:
          message.firstName = reader.string();
          break;
        case 5:
          message.lastName = reader.string();
          break;
        case 6:
          message.email = reader.string();
          break;
        case 7:
          message.partyId = reader.string();
          break;
        case 8:
          message.partyName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PaymentEvent {
    return {
      paymentIntentId: isSet(object.paymentIntentId) ? String(object.paymentIntentId) : "",
      customerId: isSet(object.customerId) ? String(object.customerId) : "",
      type: isSet(object.type) ? paymentEvent_EventTypeFromJSON(object.type) : 0,
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      email: isSet(object.email) ? String(object.email) : "",
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      partyName: isSet(object.partyName) ? String(object.partyName) : "",
    };
  },

  toJSON(message: PaymentEvent): unknown {
    const obj: any = {};
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    message.customerId !== undefined && (obj.customerId = message.customerId);
    message.type !== undefined && (obj.type = paymentEvent_EventTypeToJSON(message.type));
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.email !== undefined && (obj.email = message.email);
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.partyName !== undefined && (obj.partyName = message.partyName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PaymentEvent>, I>>(object: I): PaymentEvent {
    const message = createBasePaymentEvent();
    message.paymentIntentId = object.paymentIntentId ?? "";
    message.customerId = object.customerId ?? "";
    message.type = object.type ?? 0;
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.email = object.email ?? "";
    message.partyId = object.partyId ?? "";
    message.partyName = object.partyName ?? "";
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
