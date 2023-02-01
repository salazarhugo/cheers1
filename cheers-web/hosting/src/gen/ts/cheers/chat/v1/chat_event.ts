/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { UserItem } from "../../type/user/user";
import { Message, Room } from "./chat_service";

export const protobufPackage = "cheers.chat.v1";

export interface ChatEvent {
  create?: CreateMessage | undefined;
}

export interface CreateMessage {
  message: Message | undefined;
  sender: UserItem | undefined;
  members: UserItem[];
  room: Room | undefined;
}

function createBaseChatEvent(): ChatEvent {
  return { create: undefined };
}

export const ChatEvent = {
  encode(message: ChatEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.create !== undefined) {
      CreateMessage.encode(message.create, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChatEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChatEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.create = CreateMessage.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChatEvent {
    return { create: isSet(object.create) ? CreateMessage.fromJSON(object.create) : undefined };
  },

  toJSON(message: ChatEvent): unknown {
    const obj: any = {};
    message.create !== undefined && (obj.create = message.create ? CreateMessage.toJSON(message.create) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ChatEvent>, I>>(base?: I): ChatEvent {
    return ChatEvent.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ChatEvent>, I>>(object: I): ChatEvent {
    const message = createBaseChatEvent();
    message.create = (object.create !== undefined && object.create !== null)
      ? CreateMessage.fromPartial(object.create)
      : undefined;
    return message;
  },
};

function createBaseCreateMessage(): CreateMessage {
  return { message: undefined, sender: undefined, members: [], room: undefined };
}

export const CreateMessage = {
  encode(message: CreateMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== undefined) {
      Message.encode(message.message, writer.uint32(10).fork()).ldelim();
    }
    if (message.sender !== undefined) {
      UserItem.encode(message.sender, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.members) {
      UserItem.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.room !== undefined) {
      Room.encode(message.room, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = Message.decode(reader, reader.uint32());
          break;
        case 2:
          message.sender = UserItem.decode(reader, reader.uint32());
          break;
        case 3:
          message.members.push(UserItem.decode(reader, reader.uint32()));
          break;
        case 4:
          message.room = Room.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateMessage {
    return {
      message: isSet(object.message) ? Message.fromJSON(object.message) : undefined,
      sender: isSet(object.sender) ? UserItem.fromJSON(object.sender) : undefined,
      members: Array.isArray(object?.members) ? object.members.map((e: any) => UserItem.fromJSON(e)) : [],
      room: isSet(object.room) ? Room.fromJSON(object.room) : undefined,
    };
  },

  toJSON(message: CreateMessage): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message ? Message.toJSON(message.message) : undefined);
    message.sender !== undefined && (obj.sender = message.sender ? UserItem.toJSON(message.sender) : undefined);
    if (message.members) {
      obj.members = message.members.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.members = [];
    }
    message.room !== undefined && (obj.room = message.room ? Room.toJSON(message.room) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateMessage>, I>>(base?: I): CreateMessage {
    return CreateMessage.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateMessage>, I>>(object: I): CreateMessage {
    const message = createBaseCreateMessage();
    message.message = (object.message !== undefined && object.message !== null)
      ? Message.fromPartial(object.message)
      : undefined;
    message.sender = (object.sender !== undefined && object.sender !== null)
      ? UserItem.fromPartial(object.sender)
      : undefined;
    message.members = object.members?.map((e) => UserItem.fromPartial(e)) || [];
    message.room = (object.room !== undefined && object.room !== null) ? Room.fromPartial(object.room) : undefined;
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
