/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "../../../google/protobuf/timestamp";

export const protobufPackage = "cheers.order.v1";

export interface Order {
  id: string;
  status: string;
  amount: number;
  customerId: string;
  userId: string;
  createTime: Date | undefined;
  tickets: { [key: string]: number };
  partyId: string;
}

export interface Order_TicketsEntry {
  key: string;
  value: number;
}

export interface CreateOrderRequest {
  order: Order | undefined;
}

export interface CreateOrderResponse {
  order: Order | undefined;
}

export interface GetOrderRequest {
  orderId: string;
}

export interface GetOrderResponse {
  order: Order | undefined;
}

export interface UpdateOrderRequest {
  order: Order | undefined;
}

export interface UpdateOrderResponse {
  order: Order | undefined;
}

export interface DeleteOrderRequest {
  orderId: string;
}

export interface DeleteOrderResponse {
}

export interface ListOrderRequest {
  partyId: string | undefined;
  userId: string | undefined;
}

export interface ListOrderResponse {
  orders: Order[];
}

function createBaseOrder(): Order {
  return { id: "", status: "", amount: 0, customerId: "", userId: "", createTime: undefined, tickets: {}, partyId: "" };
}

export const Order = {
  encode(message: Order, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.status !== "") {
      writer.uint32(18).string(message.status);
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount);
    }
    if (message.customerId !== "") {
      writer.uint32(34).string(message.customerId);
    }
    if (message.userId !== "") {
      writer.uint32(50).string(message.userId);
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(58).fork()).ldelim();
    }
    Object.entries(message.tickets).forEach(([key, value]) => {
      Order_TicketsEntry.encode({ key: key as any, value }, writer.uint32(66).fork()).ldelim();
    });
    if (message.partyId !== "") {
      writer.uint32(74).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.status = reader.string();
          break;
        case 3:
          message.amount = reader.int32();
          break;
        case 4:
          message.customerId = reader.string();
          break;
        case 6:
          message.userId = reader.string();
          break;
        case 7:
          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 8:
          const entry8 = Order_TicketsEntry.decode(reader, reader.uint32());
          if (entry8.value !== undefined) {
            message.tickets[entry8.key] = entry8.value;
          }
          break;
        case 9:
          message.partyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      status: isSet(object.status) ? String(object.status) : "",
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      customerId: isSet(object.customerId) ? String(object.customerId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
      tickets: isObject(object.tickets)
        ? Object.entries(object.tickets).reduce<{ [key: string]: number }>((acc, [key, value]) => {
          acc[key] = Number(value);
          return acc;
        }, {})
        : {},
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
    };
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.status !== undefined && (obj.status = message.status);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.customerId !== undefined && (obj.customerId = message.customerId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.createTime !== undefined && (obj.createTime = message.createTime.toISOString());
    obj.tickets = {};
    if (message.tickets) {
      Object.entries(message.tickets).forEach(([k, v]) => {
        obj.tickets[k] = Math.round(v);
      });
    }
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Order>, I>>(object: I): Order {
    const message = createBaseOrder();
    message.id = object.id ?? "";
    message.status = object.status ?? "";
    message.amount = object.amount ?? 0;
    message.customerId = object.customerId ?? "";
    message.userId = object.userId ?? "";
    message.createTime = object.createTime ?? undefined;
    message.tickets = Object.entries(object.tickets ?? {}).reduce<{ [key: string]: number }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Number(value);
      }
      return acc;
    }, {});
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseOrder_TicketsEntry(): Order_TicketsEntry {
  return { key: "", value: 0 };
}

export const Order_TicketsEntry = {
  encode(message: Order_TicketsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== 0) {
      writer.uint32(16).int32(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Order_TicketsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrder_TicketsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order_TicketsEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? Number(object.value) : 0 };
  },

  toJSON(message: Order_TicketsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = Math.round(message.value));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Order_TicketsEntry>, I>>(object: I): Order_TicketsEntry {
    const message = createBaseOrder_TicketsEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseCreateOrderRequest(): CreateOrderRequest {
  return { order: undefined };
}

export const CreateOrderRequest = {
  encode(message: CreateOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateOrderRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateOrderRequest {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: CreateOrderRequest): unknown {
    const obj: any = {};
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateOrderRequest>, I>>(object: I): CreateOrderRequest {
    const message = createBaseCreateOrderRequest();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseCreateOrderResponse(): CreateOrderResponse {
  return { order: undefined };
}

export const CreateOrderResponse = {
  encode(message: CreateOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateOrderResponse {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: CreateOrderResponse): unknown {
    const obj: any = {};
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateOrderResponse>, I>>(object: I): CreateOrderResponse {
    const message = createBaseCreateOrderResponse();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseGetOrderRequest(): GetOrderRequest {
  return { orderId: "" };
}

export const GetOrderRequest = {
  encode(message: GetOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.orderId !== "") {
      writer.uint32(10).string(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrderRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetOrderRequest {
    return { orderId: isSet(object.orderId) ? String(object.orderId) : "" };
  },

  toJSON(message: GetOrderRequest): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetOrderRequest>, I>>(object: I): GetOrderRequest {
    const message = createBaseGetOrderRequest();
    message.orderId = object.orderId ?? "";
    return message;
  },
};

function createBaseGetOrderResponse(): GetOrderResponse {
  return { order: undefined };
}

export const GetOrderResponse = {
  encode(message: GetOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetOrderResponse {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: GetOrderResponse): unknown {
    const obj: any = {};
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetOrderResponse>, I>>(object: I): GetOrderResponse {
    const message = createBaseGetOrderResponse();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseUpdateOrderRequest(): UpdateOrderRequest {
  return { order: undefined };
}

export const UpdateOrderRequest = {
  encode(message: UpdateOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateOrderRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateOrderRequest {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: UpdateOrderRequest): unknown {
    const obj: any = {};
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateOrderRequest>, I>>(object: I): UpdateOrderRequest {
    const message = createBaseUpdateOrderRequest();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseUpdateOrderResponse(): UpdateOrderResponse {
  return { order: undefined };
}

export const UpdateOrderResponse = {
  encode(message: UpdateOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateOrderResponse {
    return { order: isSet(object.order) ? Order.fromJSON(object.order) : undefined };
  },

  toJSON(message: UpdateOrderResponse): unknown {
    const obj: any = {};
    message.order !== undefined && (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateOrderResponse>, I>>(object: I): UpdateOrderResponse {
    const message = createBaseUpdateOrderResponse();
    message.order = (object.order !== undefined && object.order !== null) ? Order.fromPartial(object.order) : undefined;
    return message;
  },
};

function createBaseDeleteOrderRequest(): DeleteOrderRequest {
  return { orderId: "" };
}

export const DeleteOrderRequest = {
  encode(message: DeleteOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.orderId !== "") {
      writer.uint32(10).string(message.orderId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteOrderRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orderId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteOrderRequest {
    return { orderId: isSet(object.orderId) ? String(object.orderId) : "" };
  },

  toJSON(message: DeleteOrderRequest): unknown {
    const obj: any = {};
    message.orderId !== undefined && (obj.orderId = message.orderId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteOrderRequest>, I>>(object: I): DeleteOrderRequest {
    const message = createBaseDeleteOrderRequest();
    message.orderId = object.orderId ?? "";
    return message;
  },
};

function createBaseDeleteOrderResponse(): DeleteOrderResponse {
  return {};
}

export const DeleteOrderResponse = {
  encode(_: DeleteOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): DeleteOrderResponse {
    return {};
  },

  toJSON(_: DeleteOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteOrderResponse>, I>>(_: I): DeleteOrderResponse {
    const message = createBaseDeleteOrderResponse();
    return message;
  },
};

function createBaseListOrderRequest(): ListOrderRequest {
  return { partyId: undefined, userId: undefined };
}

export const ListOrderRequest = {
  encode(message: ListOrderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== undefined) {
      writer.uint32(10).string(message.partyId);
    }
    if (message.userId !== undefined) {
      writer.uint32(18).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListOrderRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListOrderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListOrderRequest {
    return {
      partyId: isSet(object.partyId) ? String(object.partyId) : undefined,
      userId: isSet(object.userId) ? String(object.userId) : undefined,
    };
  },

  toJSON(message: ListOrderRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListOrderRequest>, I>>(object: I): ListOrderRequest {
    const message = createBaseListOrderRequest();
    message.partyId = object.partyId ?? undefined;
    message.userId = object.userId ?? undefined;
    return message;
  },
};

function createBaseListOrderResponse(): ListOrderResponse {
  return { orders: [] };
}

export const ListOrderResponse = {
  encode(message: ListOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListOrderResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orders.push(Order.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListOrderResponse {
    return { orders: Array.isArray(object?.orders) ? object.orders.map((e: any) => Order.fromJSON(e)) : [] };
  },

  toJSON(message: ListOrderResponse): unknown {
    const obj: any = {};
    if (message.orders) {
      obj.orders = message.orders.map((e) => e ? Order.toJSON(e) : undefined);
    } else {
      obj.orders = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListOrderResponse>, I>>(object: I): ListOrderResponse {
    const message = createBaseListOrderResponse();
    message.orders = object.orders?.map((e) => Order.fromPartial(e)) || [];
    return message;
  },
};

export interface OrderService {
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse>;
  GetOrder(request: GetOrderRequest): Promise<GetOrderResponse>;
  UpdateOrder(request: UpdateOrderRequest): Promise<UpdateOrderResponse>;
  DeleteOrder(request: DeleteOrderRequest): Promise<DeleteOrderResponse>;
  ListOrder(request: ListOrderRequest): Promise<ListOrderResponse>;
}

export class OrderServiceClientImpl implements OrderService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.order.v1.OrderService";
    this.rpc = rpc;
    this.CreateOrder = this.CreateOrder.bind(this);
    this.GetOrder = this.GetOrder.bind(this);
    this.UpdateOrder = this.UpdateOrder.bind(this);
    this.DeleteOrder = this.DeleteOrder.bind(this);
    this.ListOrder = this.ListOrder.bind(this);
  }
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse> {
    const data = CreateOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateOrder", data);
    return promise.then((data) => CreateOrderResponse.decode(new _m0.Reader(data)));
  }

  GetOrder(request: GetOrderRequest): Promise<GetOrderResponse> {
    const data = GetOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetOrder", data);
    return promise.then((data) => GetOrderResponse.decode(new _m0.Reader(data)));
  }

  UpdateOrder(request: UpdateOrderRequest): Promise<UpdateOrderResponse> {
    const data = UpdateOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateOrder", data);
    return promise.then((data) => UpdateOrderResponse.decode(new _m0.Reader(data)));
  }

  DeleteOrder(request: DeleteOrderRequest): Promise<DeleteOrderResponse> {
    const data = DeleteOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteOrder", data);
    return promise.then((data) => DeleteOrderResponse.decode(new _m0.Reader(data)));
  }

  ListOrder(request: ListOrderRequest): Promise<ListOrderResponse> {
    const data = ListOrderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListOrder", data);
    return promise.then((data) => ListOrderResponse.decode(new _m0.Reader(data)));
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
