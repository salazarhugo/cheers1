/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.payment.v1";

export interface CreatePaymentRequest {
  partyId: string;
  tickets: { [key: string]: number };
  firstName: string;
  lastName: string;
  email: string;
}

export interface CreatePaymentRequest_TicketsEntry {
  key: string;
  value: number;
}

export interface CreatePaymentResponse {
  clientSecret: string;
}

export interface RefundPaymentRequest {
  paymentIntentId: string;
}

export interface RefundPaymentResponse {
}

function createBaseCreatePaymentRequest(): CreatePaymentRequest {
  return { partyId: "", tickets: {}, firstName: "", lastName: "", email: "" };
}

export const CreatePaymentRequest = {
  encode(message: CreatePaymentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    Object.entries(message.tickets).forEach(([key, value]) => {
      CreatePaymentRequest_TicketsEntry.encode({ key: key as any, value }, writer.uint32(18).fork()).ldelim();
    });
    if (message.firstName !== "") {
      writer.uint32(26).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(34).string(message.lastName);
    }
    if (message.email !== "") {
      writer.uint32(42).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePaymentRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePaymentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          const entry2 = CreatePaymentRequest_TicketsEntry.decode(reader, reader.uint32());
          if (entry2.value !== undefined) {
            message.tickets[entry2.key] = entry2.value;
          }
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.firstName = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.lastName = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePaymentRequest {
    return {
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      tickets: isObject(object.tickets)
        ? Object.entries(object.tickets).reduce<{ [key: string]: number }>((acc, [key, value]) => {
          acc[key] = Number(value);
          return acc;
        }, {})
        : {},
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      email: isSet(object.email) ? String(object.email) : "",
    };
  },

  toJSON(message: CreatePaymentRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    obj.tickets = {};
    if (message.tickets) {
      Object.entries(message.tickets).forEach(([k, v]) => {
        obj.tickets[k] = Math.round(v);
      });
    }
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.email !== undefined && (obj.email = message.email);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePaymentRequest>, I>>(base?: I): CreatePaymentRequest {
    return CreatePaymentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePaymentRequest>, I>>(object: I): CreatePaymentRequest {
    const message = createBaseCreatePaymentRequest();
    message.partyId = object.partyId ?? "";
    message.tickets = Object.entries(object.tickets ?? {}).reduce<{ [key: string]: number }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Number(value);
      }
      return acc;
    }, {});
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseCreatePaymentRequest_TicketsEntry(): CreatePaymentRequest_TicketsEntry {
  return { key: "", value: 0 };
}

export const CreatePaymentRequest_TicketsEntry = {
  encode(message: CreatePaymentRequest_TicketsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).uint32(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePaymentRequest_TicketsEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePaymentRequest_TicketsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.value = reader.uint32();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePaymentRequest_TicketsEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? Number(object.value) : 0 };
  },

  toJSON(message: CreatePaymentRequest_TicketsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = Math.round(message.value));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePaymentRequest_TicketsEntry>, I>>(
    base?: I,
  ): CreatePaymentRequest_TicketsEntry {
    return CreatePaymentRequest_TicketsEntry.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePaymentRequest_TicketsEntry>, I>>(
    object: I,
  ): CreatePaymentRequest_TicketsEntry {
    const message = createBaseCreatePaymentRequest_TicketsEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseCreatePaymentResponse(): CreatePaymentResponse {
  return { clientSecret: "" };
}

export const CreatePaymentResponse = {
  encode(message: CreatePaymentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.clientSecret !== "") {
      writer.uint32(10).string(message.clientSecret);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePaymentResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePaymentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.clientSecret = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePaymentResponse {
    return { clientSecret: isSet(object.clientSecret) ? String(object.clientSecret) : "" };
  },

  toJSON(message: CreatePaymentResponse): unknown {
    const obj: any = {};
    message.clientSecret !== undefined && (obj.clientSecret = message.clientSecret);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePaymentResponse>, I>>(base?: I): CreatePaymentResponse {
    return CreatePaymentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePaymentResponse>, I>>(object: I): CreatePaymentResponse {
    const message = createBaseCreatePaymentResponse();
    message.clientSecret = object.clientSecret ?? "";
    return message;
  },
};

function createBaseRefundPaymentRequest(): RefundPaymentRequest {
  return { paymentIntentId: "" };
}

export const RefundPaymentRequest = {
  encode(message: RefundPaymentRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.paymentIntentId !== "") {
      writer.uint32(10).string(message.paymentIntentId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RefundPaymentRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRefundPaymentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.paymentIntentId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RefundPaymentRequest {
    return { paymentIntentId: isSet(object.paymentIntentId) ? String(object.paymentIntentId) : "" };
  },

  toJSON(message: RefundPaymentRequest): unknown {
    const obj: any = {};
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    return obj;
  },

  create<I extends Exact<DeepPartial<RefundPaymentRequest>, I>>(base?: I): RefundPaymentRequest {
    return RefundPaymentRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RefundPaymentRequest>, I>>(object: I): RefundPaymentRequest {
    const message = createBaseRefundPaymentRequest();
    message.paymentIntentId = object.paymentIntentId ?? "";
    return message;
  },
};

function createBaseRefundPaymentResponse(): RefundPaymentResponse {
  return {};
}

export const RefundPaymentResponse = {
  encode(_: RefundPaymentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RefundPaymentResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRefundPaymentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): RefundPaymentResponse {
    return {};
  },

  toJSON(_: RefundPaymentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<RefundPaymentResponse>, I>>(base?: I): RefundPaymentResponse {
    return RefundPaymentResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RefundPaymentResponse>, I>>(_: I): RefundPaymentResponse {
    const message = createBaseRefundPaymentResponse();
    return message;
  },
};

export interface PaymentService {
  CreatePayment(request: CreatePaymentRequest): Promise<CreatePaymentResponse>;
  RefundPayment(request: RefundPaymentRequest): Promise<RefundPaymentResponse>;
}

export class PaymentServiceClientImpl implements PaymentService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.payment.v1.PaymentService";
    this.rpc = rpc;
    this.CreatePayment = this.CreatePayment.bind(this);
    this.RefundPayment = this.RefundPayment.bind(this);
  }
  CreatePayment(request: CreatePaymentRequest): Promise<CreatePaymentResponse> {
    const data = CreatePaymentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreatePayment", data);
    return promise.then((data) => CreatePaymentResponse.decode(_m0.Reader.create(data)));
  }

  RefundPayment(request: RefundPaymentRequest): Promise<RefundPaymentResponse> {
    const data = RefundPaymentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "RefundPayment", data);
    return promise.then((data) => RefundPaymentResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
