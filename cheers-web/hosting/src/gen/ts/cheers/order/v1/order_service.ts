/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Ticket } from "../../ticket/v1/ticket";

export const protobufPackage = "cheers.order.v1";

export interface Order {
  id: string;
  status: string;
  amount: number;
  customerId: string;
  userId: string;
  createTime: number;
  tickets: Ticket[];
  partyId: string;
  partyName: string;
  partyHostId: string;
  email: string;
  firstName: string;
  lastName: string;
  /** The list of payment method types (e.g. card) that this PaymentIntent is allowed to use. */
  paymentMethodTypes: string[];
  /** The type of the PaymentMethod */
  paymentMethodType: string;
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

export interface ListUserOrdersRequest {
}

export interface ListUserOrdersResponse {
  orders: Order[];
}

export interface ListOrganizerOrdersRequest {
  query: string;
}

export interface ListOrganizerOrdersResponse {
  orders: Order[];
}

export interface ListPartyOrdersRequest {
  partyId: string;
}

export interface ListPartyOrdersResponse {
  orders: Order[];
}

function createBaseOrder(): Order {
  return {
    id: "",
    status: "",
    amount: 0,
    customerId: "",
    userId: "",
    createTime: 0,
    tickets: [],
    partyId: "",
    partyName: "",
    partyHostId: "",
    email: "",
    firstName: "",
    lastName: "",
    paymentMethodTypes: [],
    paymentMethodType: "",
  };
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
      writer.uint32(24).int64(message.amount);
    }
    if (message.customerId !== "") {
      writer.uint32(34).string(message.customerId);
    }
    if (message.userId !== "") {
      writer.uint32(50).string(message.userId);
    }
    if (message.createTime !== 0) {
      writer.uint32(56).int64(message.createTime);
    }
    for (const v of message.tickets) {
      Ticket.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.partyId !== "") {
      writer.uint32(74).string(message.partyId);
    }
    if (message.partyName !== "") {
      writer.uint32(82).string(message.partyName);
    }
    if (message.partyHostId !== "") {
      writer.uint32(90).string(message.partyHostId);
    }
    if (message.email !== "") {
      writer.uint32(98).string(message.email);
    }
    if (message.firstName !== "") {
      writer.uint32(106).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(114).string(message.lastName);
    }
    for (const v of message.paymentMethodTypes) {
      writer.uint32(122).string(v!);
    }
    if (message.paymentMethodType !== "") {
      writer.uint32(130).string(message.paymentMethodType);
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
          message.amount = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.customerId = reader.string();
          break;
        case 6:
          message.userId = reader.string();
          break;
        case 7:
          message.createTime = longToNumber(reader.int64() as Long);
          break;
        case 8:
          message.tickets.push(Ticket.decode(reader, reader.uint32()));
          break;
        case 9:
          message.partyId = reader.string();
          break;
        case 10:
          message.partyName = reader.string();
          break;
        case 11:
          message.partyHostId = reader.string();
          break;
        case 12:
          message.email = reader.string();
          break;
        case 13:
          message.firstName = reader.string();
          break;
        case 14:
          message.lastName = reader.string();
          break;
        case 15:
          message.paymentMethodTypes.push(reader.string());
          break;
        case 16:
          message.paymentMethodType = reader.string();
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
      createTime: isSet(object.createTime) ? Number(object.createTime) : 0,
      tickets: Array.isArray(object?.tickets) ? object.tickets.map((e: any) => Ticket.fromJSON(e)) : [],
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      partyName: isSet(object.partyName) ? String(object.partyName) : "",
      partyHostId: isSet(object.partyHostId) ? String(object.partyHostId) : "",
      email: isSet(object.email) ? String(object.email) : "",
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      paymentMethodTypes: Array.isArray(object?.paymentMethodTypes)
        ? object.paymentMethodTypes.map((e: any) => String(e))
        : [],
      paymentMethodType: isSet(object.paymentMethodType) ? String(object.paymentMethodType) : "",
    };
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.status !== undefined && (obj.status = message.status);
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.customerId !== undefined && (obj.customerId = message.customerId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.createTime !== undefined && (obj.createTime = Math.round(message.createTime));
    if (message.tickets) {
      obj.tickets = message.tickets.map((e) => e ? Ticket.toJSON(e) : undefined);
    } else {
      obj.tickets = [];
    }
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.partyName !== undefined && (obj.partyName = message.partyName);
    message.partyHostId !== undefined && (obj.partyHostId = message.partyHostId);
    message.email !== undefined && (obj.email = message.email);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    if (message.paymentMethodTypes) {
      obj.paymentMethodTypes = message.paymentMethodTypes.map((e) => e);
    } else {
      obj.paymentMethodTypes = [];
    }
    message.paymentMethodType !== undefined && (obj.paymentMethodType = message.paymentMethodType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Order>, I>>(object: I): Order {
    const message = createBaseOrder();
    message.id = object.id ?? "";
    message.status = object.status ?? "";
    message.amount = object.amount ?? 0;
    message.customerId = object.customerId ?? "";
    message.userId = object.userId ?? "";
    message.createTime = object.createTime ?? 0;
    message.tickets = object.tickets?.map((e) => Ticket.fromPartial(e)) || [];
    message.partyId = object.partyId ?? "";
    message.partyName = object.partyName ?? "";
    message.partyHostId = object.partyHostId ?? "";
    message.email = object.email ?? "";
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.paymentMethodTypes = object.paymentMethodTypes?.map((e) => e) || [];
    message.paymentMethodType = object.paymentMethodType ?? "";
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

function createBaseListUserOrdersRequest(): ListUserOrdersRequest {
  return {};
}

export const ListUserOrdersRequest = {
  encode(_: ListUserOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserOrdersRequest();
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

  fromJSON(_: any): ListUserOrdersRequest {
    return {};
  },

  toJSON(_: ListUserOrdersRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUserOrdersRequest>, I>>(_: I): ListUserOrdersRequest {
    const message = createBaseListUserOrdersRequest();
    return message;
  },
};

function createBaseListUserOrdersResponse(): ListUserOrdersResponse {
  return { orders: [] };
}

export const ListUserOrdersResponse = {
  encode(message: ListUserOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUserOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserOrdersResponse();
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

  fromJSON(object: any): ListUserOrdersResponse {
    return { orders: Array.isArray(object?.orders) ? object.orders.map((e: any) => Order.fromJSON(e)) : [] };
  },

  toJSON(message: ListUserOrdersResponse): unknown {
    const obj: any = {};
    if (message.orders) {
      obj.orders = message.orders.map((e) => e ? Order.toJSON(e) : undefined);
    } else {
      obj.orders = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUserOrdersResponse>, I>>(object: I): ListUserOrdersResponse {
    const message = createBaseListUserOrdersResponse();
    message.orders = object.orders?.map((e) => Order.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListOrganizerOrdersRequest(): ListOrganizerOrdersRequest {
  return { query: "" };
}

export const ListOrganizerOrdersRequest = {
  encode(message: ListOrganizerOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListOrganizerOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListOrganizerOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.query = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListOrganizerOrdersRequest {
    return { query: isSet(object.query) ? String(object.query) : "" };
  },

  toJSON(message: ListOrganizerOrdersRequest): unknown {
    const obj: any = {};
    message.query !== undefined && (obj.query = message.query);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListOrganizerOrdersRequest>, I>>(object: I): ListOrganizerOrdersRequest {
    const message = createBaseListOrganizerOrdersRequest();
    message.query = object.query ?? "";
    return message;
  },
};

function createBaseListOrganizerOrdersResponse(): ListOrganizerOrdersResponse {
  return { orders: [] };
}

export const ListOrganizerOrdersResponse = {
  encode(message: ListOrganizerOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListOrganizerOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListOrganizerOrdersResponse();
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

  fromJSON(object: any): ListOrganizerOrdersResponse {
    return { orders: Array.isArray(object?.orders) ? object.orders.map((e: any) => Order.fromJSON(e)) : [] };
  },

  toJSON(message: ListOrganizerOrdersResponse): unknown {
    const obj: any = {};
    if (message.orders) {
      obj.orders = message.orders.map((e) => e ? Order.toJSON(e) : undefined);
    } else {
      obj.orders = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListOrganizerOrdersResponse>, I>>(object: I): ListOrganizerOrdersResponse {
    const message = createBaseListOrganizerOrdersResponse();
    message.orders = object.orders?.map((e) => Order.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListPartyOrdersRequest(): ListPartyOrdersRequest {
  return { partyId: "" };
}

export const ListPartyOrdersRequest = {
  encode(message: ListPartyOrdersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPartyOrdersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPartyOrdersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.partyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListPartyOrdersRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: ListPartyOrdersRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListPartyOrdersRequest>, I>>(object: I): ListPartyOrdersRequest {
    const message = createBaseListPartyOrdersRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseListPartyOrdersResponse(): ListPartyOrdersResponse {
  return { orders: [] };
}

export const ListPartyOrdersResponse = {
  encode(message: ListPartyOrdersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPartyOrdersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPartyOrdersResponse();
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

  fromJSON(object: any): ListPartyOrdersResponse {
    return { orders: Array.isArray(object?.orders) ? object.orders.map((e: any) => Order.fromJSON(e)) : [] };
  },

  toJSON(message: ListPartyOrdersResponse): unknown {
    const obj: any = {};
    if (message.orders) {
      obj.orders = message.orders.map((e) => e ? Order.toJSON(e) : undefined);
    } else {
      obj.orders = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListPartyOrdersResponse>, I>>(object: I): ListPartyOrdersResponse {
    const message = createBaseListPartyOrdersResponse();
    message.orders = object.orders?.map((e) => Order.fromPartial(e)) || [];
    return message;
  },
};

export interface OrderService {
  CreateOrder(request: CreateOrderRequest): Promise<CreateOrderResponse>;
  GetOrder(request: GetOrderRequest): Promise<GetOrderResponse>;
  UpdateOrder(request: UpdateOrderRequest): Promise<UpdateOrderResponse>;
  DeleteOrder(request: DeleteOrderRequest): Promise<DeleteOrderResponse>;
  ListUserOrders(request: ListUserOrdersRequest): Promise<ListUserOrdersResponse>;
  ListOrganizerOrders(request: ListOrganizerOrdersRequest): Promise<ListOrganizerOrdersResponse>;
  ListPartyOrders(request: ListPartyOrdersRequest): Promise<ListPartyOrdersResponse>;
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
    this.ListUserOrders = this.ListUserOrders.bind(this);
    this.ListOrganizerOrders = this.ListOrganizerOrders.bind(this);
    this.ListPartyOrders = this.ListPartyOrders.bind(this);
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

  ListUserOrders(request: ListUserOrdersRequest): Promise<ListUserOrdersResponse> {
    const data = ListUserOrdersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListUserOrders", data);
    return promise.then((data) => ListUserOrdersResponse.decode(new _m0.Reader(data)));
  }

  ListOrganizerOrders(request: ListOrganizerOrdersRequest): Promise<ListOrganizerOrdersResponse> {
    const data = ListOrganizerOrdersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListOrganizerOrders", data);
    return promise.then((data) => ListOrganizerOrdersResponse.decode(new _m0.Reader(data)));
  }

  ListPartyOrders(request: ListPartyOrdersRequest): Promise<ListPartyOrdersResponse> {
    const data = ListPartyOrdersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListPartyOrders", data);
    return promise.then((data) => ListPartyOrdersResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
