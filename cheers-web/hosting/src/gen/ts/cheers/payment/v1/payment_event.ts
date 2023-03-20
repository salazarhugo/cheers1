/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Order } from "../../order/v1/order_service";

export const protobufPackage = "cheers.payment.v1";

export interface PaymentEvent {
  paymentIntentId: string;
  type: PaymentEvent_EventType;
  order: Order | undefined;
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
  return { paymentIntentId: "", type: 0, order: undefined };
}

export const PaymentEvent = {
  encode(message: PaymentEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.paymentIntentId !== "") {
      writer.uint32(10).string(message.paymentIntentId);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PaymentEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePaymentEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.paymentIntentId = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.order = Order.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PaymentEvent {
    return {
      paymentIntentId: isSet(object.paymentIntentId) ? String(object.paymentIntentId) : "",
      type: isSet(object.type) ? paymentEvent_EventTypeFromJSON(object.type) : 0,
      order: isSet(object.order) ? Order.fromJSON(object.order) : undefined,
    };
  },

  toJSON(message: PaymentEvent): unknown {
    const obj: any = {};
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    message.type !== undefined && (obj.type = paymentEvent_EventTypeToJSON(message.type));
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PaymentEvent>, I>>(base?: I): PaymentEvent {
    return PaymentEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PaymentEvent>, I>>(object: I): PaymentEvent {
    const message = createBasePaymentEvent();
    message.paymentIntentId = object.paymentIntentId ?? "";
    message.type = object.type ?? 0;
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
