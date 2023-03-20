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
  partyName: string;
  partyStartTime: number;
  partyEndTime: number;
  locationName: string;
  organizer: string;
  paymentIntentId: string;
  userId: string;
  validated: boolean;
}

export interface CreateTicketRequest {
  ticket: Ticket | undefined;
}

export interface CreateTicketResponse {
  ticket: Ticket | undefined;
}

export interface GetTicketRequest {
  ticketId: string;
}

export interface GetTicketResponse {
  ticket: Ticket | undefined;
}

export interface UpdateTicketRequest {
  ticket: Ticket | undefined;
}

export interface UpdateTicketResponse {
  ticket: Ticket | undefined;
}

export interface DeleteTicketRequest {
  ticketId: string;
}

export interface DeleteTicketResponse {
}

export interface ListTicketRequest {
  partyId?: string | undefined;
  userId?: string | undefined;
}

export interface ListTicketResponse {
  tickets: Ticket[];
}

function createBaseTicket(): Ticket {
  return {
    id: "",
    price: 0,
    quantity: 0,
    name: "",
    description: "",
    partyId: "",
    partyName: "",
    partyStartTime: 0,
    partyEndTime: 0,
    locationName: "",
    organizer: "",
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
    if (message.partyName !== "") {
      writer.uint32(82).string(message.partyName);
    }
    if (message.partyStartTime !== 0) {
      writer.uint32(88).int64(message.partyStartTime);
    }
    if (message.partyEndTime !== 0) {
      writer.uint32(96).int64(message.partyEndTime);
    }
    if (message.locationName !== "") {
      writer.uint32(106).string(message.locationName);
    }
    if (message.organizer !== "") {
      writer.uint32(114).string(message.organizer);
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTicket();
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
          if (tag != 16) {
            break;
          }

          message.price = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.quantity = reader.uint32();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.name = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.description = reader.string();
          continue;
        case 6:
          if (tag != 50) {
            break;
          }

          message.partyId = reader.string();
          continue;
        case 10:
          if (tag != 82) {
            break;
          }

          message.partyName = reader.string();
          continue;
        case 11:
          if (tag != 88) {
            break;
          }

          message.partyStartTime = longToNumber(reader.int64() as Long);
          continue;
        case 12:
          if (tag != 96) {
            break;
          }

          message.partyEndTime = longToNumber(reader.int64() as Long);
          continue;
        case 13:
          if (tag != 106) {
            break;
          }

          message.locationName = reader.string();
          continue;
        case 14:
          if (tag != 114) {
            break;
          }

          message.organizer = reader.string();
          continue;
        case 7:
          if (tag != 58) {
            break;
          }

          message.paymentIntentId = reader.string();
          continue;
        case 8:
          if (tag != 66) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.validated = reader.bool();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
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
      partyName: isSet(object.partyName) ? String(object.partyName) : "",
      partyStartTime: isSet(object.partyStartTime) ? Number(object.partyStartTime) : 0,
      partyEndTime: isSet(object.partyEndTime) ? Number(object.partyEndTime) : 0,
      locationName: isSet(object.locationName) ? String(object.locationName) : "",
      organizer: isSet(object.organizer) ? String(object.organizer) : "",
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
    message.partyName !== undefined && (obj.partyName = message.partyName);
    message.partyStartTime !== undefined && (obj.partyStartTime = Math.round(message.partyStartTime));
    message.partyEndTime !== undefined && (obj.partyEndTime = Math.round(message.partyEndTime));
    message.locationName !== undefined && (obj.locationName = message.locationName);
    message.organizer !== undefined && (obj.organizer = message.organizer);
    message.paymentIntentId !== undefined && (obj.paymentIntentId = message.paymentIntentId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.validated !== undefined && (obj.validated = message.validated);
    return obj;
  },

  create<I extends Exact<DeepPartial<Ticket>, I>>(base?: I): Ticket {
    return Ticket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Ticket>, I>>(object: I): Ticket {
    const message = createBaseTicket();
    message.id = object.id ?? "";
    message.price = object.price ?? 0;
    message.quantity = object.quantity ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.partyId = object.partyId ?? "";
    message.partyName = object.partyName ?? "";
    message.partyStartTime = object.partyStartTime ?? 0;
    message.partyEndTime = object.partyEndTime ?? 0;
    message.locationName = object.locationName ?? "";
    message.organizer = object.organizer ?? "";
    message.paymentIntentId = object.paymentIntentId ?? "";
    message.userId = object.userId ?? "";
    message.validated = object.validated ?? false;
    return message;
  },
};

function createBaseCreateTicketRequest(): CreateTicketRequest {
  return { ticket: undefined };
}

export const CreateTicketRequest = {
  encode(message: CreateTicketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTicketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTicketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateTicketRequest {
    return { ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined };
  },

  toJSON(message: CreateTicketRequest): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTicketRequest>, I>>(base?: I): CreateTicketRequest {
    return CreateTicketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateTicketRequest>, I>>(object: I): CreateTicketRequest {
    const message = createBaseCreateTicketRequest();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    return message;
  },
};

function createBaseCreateTicketResponse(): CreateTicketResponse {
  return { ticket: undefined };
}

export const CreateTicketResponse = {
  encode(message: CreateTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTicketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTicketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateTicketResponse {
    return { ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined };
  },

  toJSON(message: CreateTicketResponse): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTicketResponse>, I>>(base?: I): CreateTicketResponse {
    return CreateTicketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateTicketResponse>, I>>(object: I): CreateTicketResponse {
    const message = createBaseCreateTicketResponse();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    return message;
  },
};

function createBaseGetTicketRequest(): GetTicketRequest {
  return { ticketId: "" };
}

export const GetTicketRequest = {
  encode(message: GetTicketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticketId !== "") {
      writer.uint32(10).string(message.ticketId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTicketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTicketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticketId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTicketRequest {
    return { ticketId: isSet(object.ticketId) ? String(object.ticketId) : "" };
  },

  toJSON(message: GetTicketRequest): unknown {
    const obj: any = {};
    message.ticketId !== undefined && (obj.ticketId = message.ticketId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTicketRequest>, I>>(base?: I): GetTicketRequest {
    return GetTicketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetTicketRequest>, I>>(object: I): GetTicketRequest {
    const message = createBaseGetTicketRequest();
    message.ticketId = object.ticketId ?? "";
    return message;
  },
};

function createBaseGetTicketResponse(): GetTicketResponse {
  return { ticket: undefined };
}

export const GetTicketResponse = {
  encode(message: GetTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTicketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTicketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTicketResponse {
    return { ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined };
  },

  toJSON(message: GetTicketResponse): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTicketResponse>, I>>(base?: I): GetTicketResponse {
    return GetTicketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetTicketResponse>, I>>(object: I): GetTicketResponse {
    const message = createBaseGetTicketResponse();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    return message;
  },
};

function createBaseUpdateTicketRequest(): UpdateTicketRequest {
  return { ticket: undefined };
}

export const UpdateTicketRequest = {
  encode(message: UpdateTicketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTicketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTicketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateTicketRequest {
    return { ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined };
  },

  toJSON(message: UpdateTicketRequest): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTicketRequest>, I>>(base?: I): UpdateTicketRequest {
    return UpdateTicketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTicketRequest>, I>>(object: I): UpdateTicketRequest {
    const message = createBaseUpdateTicketRequest();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    return message;
  },
};

function createBaseUpdateTicketResponse(): UpdateTicketResponse {
  return { ticket: undefined };
}

export const UpdateTicketResponse = {
  encode(message: UpdateTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTicketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTicketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateTicketResponse {
    return { ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined };
  },

  toJSON(message: UpdateTicketResponse): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTicketResponse>, I>>(base?: I): UpdateTicketResponse {
    return UpdateTicketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTicketResponse>, I>>(object: I): UpdateTicketResponse {
    const message = createBaseUpdateTicketResponse();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    return message;
  },
};

function createBaseDeleteTicketRequest(): DeleteTicketRequest {
  return { ticketId: "" };
}

export const DeleteTicketRequest = {
  encode(message: DeleteTicketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticketId !== "") {
      writer.uint32(10).string(message.ticketId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTicketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTicketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticketId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteTicketRequest {
    return { ticketId: isSet(object.ticketId) ? String(object.ticketId) : "" };
  },

  toJSON(message: DeleteTicketRequest): unknown {
    const obj: any = {};
    message.ticketId !== undefined && (obj.ticketId = message.ticketId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTicketRequest>, I>>(base?: I): DeleteTicketRequest {
    return DeleteTicketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteTicketRequest>, I>>(object: I): DeleteTicketRequest {
    const message = createBaseDeleteTicketRequest();
    message.ticketId = object.ticketId ?? "";
    return message;
  },
};

function createBaseDeleteTicketResponse(): DeleteTicketResponse {
  return {};
}

export const DeleteTicketResponse = {
  encode(_: DeleteTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTicketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTicketResponse();
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

  fromJSON(_: any): DeleteTicketResponse {
    return {};
  },

  toJSON(_: DeleteTicketResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTicketResponse>, I>>(base?: I): DeleteTicketResponse {
    return DeleteTicketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteTicketResponse>, I>>(_: I): DeleteTicketResponse {
    const message = createBaseDeleteTicketResponse();
    return message;
  },
};

function createBaseListTicketRequest(): ListTicketRequest {
  return { partyId: undefined, userId: undefined };
}

export const ListTicketRequest = {
  encode(message: ListTicketRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== undefined) {
      writer.uint32(10).string(message.partyId);
    }
    if (message.userId !== undefined) {
      writer.uint32(18).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTicketRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTicketRequest();
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

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListTicketRequest {
    return {
      partyId: isSet(object.partyId) ? String(object.partyId) : undefined,
      userId: isSet(object.userId) ? String(object.userId) : undefined,
    };
  },

  toJSON(message: ListTicketRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListTicketRequest>, I>>(base?: I): ListTicketRequest {
    return ListTicketRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListTicketRequest>, I>>(object: I): ListTicketRequest {
    const message = createBaseListTicketRequest();
    message.partyId = object.partyId ?? undefined;
    message.userId = object.userId ?? undefined;
    return message;
  },
};

function createBaseListTicketResponse(): ListTicketResponse {
  return { tickets: [] };
}

export const ListTicketResponse = {
  encode(message: ListTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tickets) {
      Ticket.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListTicketResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListTicketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.tickets.push(Ticket.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListTicketResponse {
    return { tickets: Array.isArray(object?.tickets) ? object.tickets.map((e: any) => Ticket.fromJSON(e)) : [] };
  },

  toJSON(message: ListTicketResponse): unknown {
    const obj: any = {};
    if (message.tickets) {
      obj.tickets = message.tickets.map((e) => e ? Ticket.toJSON(e) : undefined);
    } else {
      obj.tickets = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListTicketResponse>, I>>(base?: I): ListTicketResponse {
    return ListTicketResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListTicketResponse>, I>>(object: I): ListTicketResponse {
    const message = createBaseListTicketResponse();
    message.tickets = object.tickets?.map((e) => Ticket.fromPartial(e)) || [];
    return message;
  },
};

export interface TicketService {
  CreateTicket(request: CreateTicketRequest): Promise<CreateTicketResponse>;
  GetTicket(request: GetTicketRequest): Promise<GetTicketResponse>;
  UpdateTicket(request: UpdateTicketRequest): Promise<UpdateTicketResponse>;
  DeleteTicket(request: DeleteTicketRequest): Promise<DeleteTicketResponse>;
  ListTicket(request: ListTicketRequest): Promise<ListTicketResponse>;
}

export class TicketServiceClientImpl implements TicketService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.ticket.v1.TicketService";
    this.rpc = rpc;
    this.CreateTicket = this.CreateTicket.bind(this);
    this.GetTicket = this.GetTicket.bind(this);
    this.UpdateTicket = this.UpdateTicket.bind(this);
    this.DeleteTicket = this.DeleteTicket.bind(this);
    this.ListTicket = this.ListTicket.bind(this);
  }
  CreateTicket(request: CreateTicketRequest): Promise<CreateTicketResponse> {
    const data = CreateTicketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateTicket", data);
    return promise.then((data) => CreateTicketResponse.decode(_m0.Reader.create(data)));
  }

  GetTicket(request: GetTicketRequest): Promise<GetTicketResponse> {
    const data = GetTicketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetTicket", data);
    return promise.then((data) => GetTicketResponse.decode(_m0.Reader.create(data)));
  }

  UpdateTicket(request: UpdateTicketRequest): Promise<UpdateTicketResponse> {
    const data = UpdateTicketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateTicket", data);
    return promise.then((data) => UpdateTicketResponse.decode(_m0.Reader.create(data)));
  }

  DeleteTicket(request: DeleteTicketRequest): Promise<DeleteTicketResponse> {
    const data = DeleteTicketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteTicket", data);
    return promise.then((data) => DeleteTicketResponse.decode(_m0.Reader.create(data)));
  }

  ListTicket(request: ListTicketRequest): Promise<ListTicketResponse> {
    const data = ListTicketRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListTicket", data);
    return promise.then((data) => ListTicketResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
