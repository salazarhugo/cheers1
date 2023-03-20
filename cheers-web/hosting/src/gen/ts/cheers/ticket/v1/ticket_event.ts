/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";
import { Ticket } from "./ticket_service";

export const protobufPackage = "cheers.ticket.v1";

export interface TicketEvent {
  create?: CreateTicket | undefined;
  update?: UpdateTicket | undefined;
  delete?: DeleteTicket | undefined;
}

export interface CreateTicket {
  ticket: Ticket | undefined;
  user: UserItem | undefined;
}

export interface UpdateTicket {
  ticket: Ticket | undefined;
  user: UserItem | undefined;
}

export interface DeleteTicket {
  ticket: Ticket | undefined;
  user: UserItem | undefined;
}

function createBaseTicketEvent(): TicketEvent {
  return { create: undefined, update: undefined, delete: undefined };
}

export const TicketEvent = {
  encode(message: TicketEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.create !== undefined) {
      CreateTicket.encode(message.create, writer.uint32(10).fork()).ldelim();
    }
    if (message.update !== undefined) {
      UpdateTicket.encode(message.update, writer.uint32(18).fork()).ldelim();
    }
    if (message.delete !== undefined) {
      DeleteTicket.encode(message.delete, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TicketEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTicketEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.create = CreateTicket.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.update = UpdateTicket.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.delete = DeleteTicket.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TicketEvent {
    return {
      create: isSet(object.create) ? CreateTicket.fromJSON(object.create) : undefined,
      update: isSet(object.update) ? UpdateTicket.fromJSON(object.update) : undefined,
      delete: isSet(object.delete) ? DeleteTicket.fromJSON(object.delete) : undefined,
    };
  },

  toJSON(message: TicketEvent): unknown {
    const obj: any = {};
    message.create !== undefined && (obj.create = message.create ? CreateTicket.toJSON(message.create) : undefined);
    message.update !== undefined && (obj.update = message.update ? UpdateTicket.toJSON(message.update) : undefined);
    message.delete !== undefined && (obj.delete = message.delete ? DeleteTicket.toJSON(message.delete) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TicketEvent>, I>>(base?: I): TicketEvent {
    return TicketEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TicketEvent>, I>>(object: I): TicketEvent {
    const message = createBaseTicketEvent();
    message.create = (object.create !== undefined && object.create !== null)
      ? CreateTicket.fromPartial(object.create)
      : undefined;
    message.update = (object.update !== undefined && object.update !== null)
      ? UpdateTicket.fromPartial(object.update)
      : undefined;
    message.delete = (object.delete !== undefined && object.delete !== null)
      ? DeleteTicket.fromPartial(object.delete)
      : undefined;
    return message;
  },
};

function createBaseCreateTicket(): CreateTicket {
  return { ticket: undefined, user: undefined };
}

export const CreateTicket = {
  encode(message: CreateTicket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTicket {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTicket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.user = UserItem.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateTicket {
    return {
      ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreateTicket): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTicket>, I>>(base?: I): CreateTicket {
    return CreateTicket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateTicket>, I>>(object: I): CreateTicket {
    const message = createBaseCreateTicket();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseUpdateTicket(): UpdateTicket {
  return { ticket: undefined, user: undefined };
}

export const UpdateTicket = {
  encode(message: UpdateTicket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTicket {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTicket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.user = UserItem.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateTicket {
    return {
      ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: UpdateTicket): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTicket>, I>>(base?: I): UpdateTicket {
    return UpdateTicket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateTicket>, I>>(object: I): UpdateTicket {
    const message = createBaseUpdateTicket();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseDeleteTicket(): DeleteTicket {
  return { ticket: undefined, user: undefined };
}

export const DeleteTicket = {
  encode(message: DeleteTicket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ticket !== undefined) {
      Ticket.encode(message.ticket, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserItem.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTicket {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTicket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.ticket = Ticket.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.user = UserItem.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteTicket {
    return {
      ticket: isSet(object.ticket) ? Ticket.fromJSON(object.ticket) : undefined,
      user: isSet(object.user) ? UserItem.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: DeleteTicket): unknown {
    const obj: any = {};
    message.ticket !== undefined && (obj.ticket = message.ticket ? Ticket.toJSON(message.ticket) : undefined);
    message.user !== undefined && (obj.user = message.user ? UserItem.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTicket>, I>>(base?: I): DeleteTicket {
    return DeleteTicket.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteTicket>, I>>(object: I): DeleteTicket {
    const message = createBaseDeleteTicket();
    message.ticket = (object.ticket !== undefined && object.ticket !== null)
      ? Ticket.fromPartial(object.ticket)
      : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? UserItem.fromPartial(object.user) : undefined;
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
