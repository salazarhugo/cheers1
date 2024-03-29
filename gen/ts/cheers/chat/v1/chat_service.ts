/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import { Timestamp } from "../../../google/protobuf/timestamp";
import { UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.chat.v1";

export enum RoomStatus {
  EMPTY = 0,
  OPENED = 1,
  SENT = 2,
  RECEIVED = 3,
  NEW = 4,
  UNRECOGNIZED = -1,
}

export function roomStatusFromJSON(object: any): RoomStatus {
  switch (object) {
    case 0:
    case "EMPTY":
      return RoomStatus.EMPTY;
    case 1:
    case "OPENED":
      return RoomStatus.OPENED;
    case 2:
    case "SENT":
      return RoomStatus.SENT;
    case 3:
    case "RECEIVED":
      return RoomStatus.RECEIVED;
    case 4:
    case "NEW":
      return RoomStatus.NEW;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RoomStatus.UNRECOGNIZED;
  }
}

export function roomStatusToJSON(object: RoomStatus): string {
  switch (object) {
    case RoomStatus.EMPTY:
      return "EMPTY";
    case RoomStatus.OPENED:
      return "OPENED";
    case RoomStatus.SENT:
      return "SENT";
    case RoomStatus.RECEIVED:
      return "RECEIVED";
    case RoomStatus.NEW:
      return "NEW";
    case RoomStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum RoomType {
  DIRECT = 0,
  GROUP = 1,
  UNRECOGNIZED = -1,
}

export function roomTypeFromJSON(object: any): RoomType {
  switch (object) {
    case 0:
    case "DIRECT":
      return RoomType.DIRECT;
    case 1:
    case "GROUP":
      return RoomType.GROUP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RoomType.UNRECOGNIZED;
  }
}

export function roomTypeToJSON(object: RoomType): string {
  switch (object) {
    case RoomType.DIRECT:
      return "DIRECT";
    case RoomType.GROUP:
      return "GROUP";
    case RoomType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum MessageType {
  TEXT = 0,
  IMAGE = 1,
  VIDEO = 2,
  UNRECOGNIZED = -1,
}

export function messageTypeFromJSON(object: any): MessageType {
  switch (object) {
    case 0:
    case "TEXT":
      return MessageType.TEXT;
    case 1:
    case "IMAGE":
      return MessageType.IMAGE;
    case 2:
    case "VIDEO":
      return MessageType.VIDEO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MessageType.UNRECOGNIZED;
  }
}

export function messageTypeToJSON(object: MessageType): string {
  switch (object) {
    case MessageType.TEXT:
      return "TEXT";
    case MessageType.IMAGE:
      return "IMAGE";
    case MessageType.VIDEO:
      return "VIDEO";
    case MessageType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface TypingEvent {
  roomId: string;
  userId: string;
  type: TypingEvent_Type;
}

export enum TypingEvent_Type {
  START = 0,
  STOP = 1,
  UNRECOGNIZED = -1,
}

export function typingEvent_TypeFromJSON(object: any): TypingEvent_Type {
  switch (object) {
    case 0:
    case "START":
      return TypingEvent_Type.START;
    case 1:
    case "STOP":
      return TypingEvent_Type.STOP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TypingEvent_Type.UNRECOGNIZED;
  }
}

export function typingEvent_TypeToJSON(object: TypingEvent_Type): string {
  switch (object) {
    case TypingEvent_Type.START:
      return "START";
    case TypingEvent_Type.STOP:
      return "STOP";
    case TypingEvent_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ListRoomRequest {
  pageSize: number;
  page: number;
}

export interface ListMembersRequest {
  roomId: string;
  pageSize: number;
  page: number;
}

export interface ListMembersResponse {
  users: UserItem[];
}

export interface UserIdReq {
  userId: string;
}

export interface Empty {
}

export interface AddTokenReq {
  token: string;
}

export interface CreateChatReq {
  userIds: string[];
  groupName: string;
}

export interface TypingReq {
  roomId: string;
  username: string;
  avatarUrl: string;
}

export interface JoinRoomRequest {
  roomId: string;
}

export interface RoomId {
  roomId: string;
}

export interface GetRoomIdReq {
  recipientId: string;
}

export interface Room {
  id: string;
  name: string;
  verified: boolean;
  typing: boolean;
  owner: string;
  type: RoomType;
  status: RoomStatus;
  admins: string[];
  members: string[];
  lastMessageText: string;
  picture: string;
  lastMessageType: MessageType;
  created: Date | undefined;
  lastMessageTime: Date | undefined;
}

export interface LikeMessageReq {
  roomId: string;
  messageId: string;
}

export interface Message {
  id: string;
  sender: string;
  room: Room | undefined;
  message: string;
  senderpicture: string;
  senderName: string;
  senderUsername: string;
  created: Date | undefined;
  photoUrl: string;
  type: MessageType;
  likeCount: number;
  status: Message_Status;
}

export enum Message_Status {
  EMPTY = 0,
  SENT = 1,
  DELIVERED = 2,
  READ = 3,
  FAILED = 4,
  UNRECOGNIZED = -1,
}

export function message_StatusFromJSON(object: any): Message_Status {
  switch (object) {
    case 0:
    case "EMPTY":
      return Message_Status.EMPTY;
    case 1:
    case "SENT":
      return Message_Status.SENT;
    case 2:
    case "DELIVERED":
      return Message_Status.DELIVERED;
    case 3:
    case "READ":
      return Message_Status.READ;
    case 4:
    case "FAILED":
      return Message_Status.FAILED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Message_Status.UNRECOGNIZED;
  }
}

export function message_StatusToJSON(object: Message_Status): string {
  switch (object) {
    case Message_Status.EMPTY:
      return "EMPTY";
    case Message_Status.SENT:
      return "SENT";
    case Message_Status.DELIVERED:
      return "DELIVERED";
    case Message_Status.READ:
      return "READ";
    case Message_Status.FAILED:
      return "FAILED";
    case Message_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface SendMessageResponse {
  status: Message_Status;
}

function createBaseTypingEvent(): TypingEvent {
  return { roomId: "", userId: "", type: 0 };
}

export const TypingEvent = {
  encode(message: TypingEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TypingEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTypingEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TypingEvent {
    return {
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      type: isSet(object.type) ? typingEvent_TypeFromJSON(object.type) : 0,
    };
  },

  toJSON(message: TypingEvent): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.type !== undefined && (obj.type = typingEvent_TypeToJSON(message.type));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TypingEvent>, I>>(object: I): TypingEvent {
    const message = createBaseTypingEvent();
    message.roomId = object.roomId ?? "";
    message.userId = object.userId ?? "";
    message.type = object.type ?? 0;
    return message;
  },
};

function createBaseListRoomRequest(): ListRoomRequest {
  return { pageSize: 0, page: 0 };
}

export const ListRoomRequest = {
  encode(message: ListRoomRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRoomRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRoomRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.pageSize = reader.int32();
          break;
        case 3:
          message.page = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListRoomRequest {
    return {
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListRoomRequest): unknown {
    const obj: any = {};
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListRoomRequest>, I>>(object: I): ListRoomRequest {
    const message = createBaseListRoomRequest();
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListMembersRequest(): ListMembersRequest {
  return { roomId: "", pageSize: 0, page: 0 };
}

export const ListMembersRequest = {
  encode(message: ListMembersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMembersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMembersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        case 2:
          message.pageSize = reader.int32();
          break;
        case 3:
          message.page = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListMembersRequest {
    return {
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      page: isSet(object.page) ? Number(object.page) : 0,
    };
  },

  toJSON(message: ListMembersRequest): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.page !== undefined && (obj.page = Math.round(message.page));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListMembersRequest>, I>>(object: I): ListMembersRequest {
    const message = createBaseListMembersRequest();
    message.roomId = object.roomId ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.page = object.page ?? 0;
    return message;
  },
};

function createBaseListMembersResponse(): ListMembersResponse {
  return { users: [] };
}

export const ListMembersResponse = {
  encode(message: ListMembersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListMembersResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMembersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(UserItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListMembersResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListMembersResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListMembersResponse>, I>>(object: I): ListMembersResponse {
    const message = createBaseListMembersResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUserIdReq(): UserIdReq {
  return { userId: "" };
}

export const UserIdReq = {
  encode(message: UserIdReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserIdReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserIdReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserIdReq {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: UserIdReq): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserIdReq>, I>>(object: I): UserIdReq {
    const message = createBaseUserIdReq();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseEmpty(): Empty {
  return {};
}

export const Empty = {
  encode(_: Empty, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Empty {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmpty();
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

  fromJSON(_: any): Empty {
    return {};
  },

  toJSON(_: Empty): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Empty>, I>>(_: I): Empty {
    const message = createBaseEmpty();
    return message;
  },
};

function createBaseAddTokenReq(): AddTokenReq {
  return { token: "" };
}

export const AddTokenReq = {
  encode(message: AddTokenReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddTokenReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddTokenReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddTokenReq {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: AddTokenReq): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AddTokenReq>, I>>(object: I): AddTokenReq {
    const message = createBaseAddTokenReq();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseCreateChatReq(): CreateChatReq {
  return { userIds: [], groupName: "" };
}

export const CreateChatReq = {
  encode(message: CreateChatReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userIds) {
      writer.uint32(10).string(v!);
    }
    if (message.groupName !== "") {
      writer.uint32(18).string(message.groupName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateChatReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateChatReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userIds.push(reader.string());
          break;
        case 2:
          message.groupName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateChatReq {
    return {
      userIds: Array.isArray(object?.userIds) ? object.userIds.map((e: any) => String(e)) : [],
      groupName: isSet(object.groupName) ? String(object.groupName) : "",
    };
  },

  toJSON(message: CreateChatReq): unknown {
    const obj: any = {};
    if (message.userIds) {
      obj.userIds = message.userIds.map((e) => e);
    } else {
      obj.userIds = [];
    }
    message.groupName !== undefined && (obj.groupName = message.groupName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateChatReq>, I>>(object: I): CreateChatReq {
    const message = createBaseCreateChatReq();
    message.userIds = object.userIds?.map((e) => e) || [];
    message.groupName = object.groupName ?? "";
    return message;
  },
};

function createBaseTypingReq(): TypingReq {
  return { roomId: "", username: "", avatarUrl: "" };
}

export const TypingReq = {
  encode(message: TypingReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(26).string(message.avatarUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TypingReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTypingReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        case 2:
          message.username = reader.string();
          break;
        case 3:
          message.avatarUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TypingReq {
    return {
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      username: isSet(object.username) ? String(object.username) : "",
      avatarUrl: isSet(object.avatarUrl) ? String(object.avatarUrl) : "",
    };
  },

  toJSON(message: TypingReq): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.username !== undefined && (obj.username = message.username);
    message.avatarUrl !== undefined && (obj.avatarUrl = message.avatarUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TypingReq>, I>>(object: I): TypingReq {
    const message = createBaseTypingReq();
    message.roomId = object.roomId ?? "";
    message.username = object.username ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    return message;
  },
};

function createBaseJoinRoomRequest(): JoinRoomRequest {
  return { roomId: "" };
}

export const JoinRoomRequest = {
  encode(message: JoinRoomRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JoinRoomRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJoinRoomRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): JoinRoomRequest {
    return { roomId: isSet(object.roomId) ? String(object.roomId) : "" };
  },

  toJSON(message: JoinRoomRequest): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<JoinRoomRequest>, I>>(object: I): JoinRoomRequest {
    const message = createBaseJoinRoomRequest();
    message.roomId = object.roomId ?? "";
    return message;
  },
};

function createBaseRoomId(): RoomId {
  return { roomId: "" };
}

export const RoomId = {
  encode(message: RoomId, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RoomId {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoomId();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoomId {
    return { roomId: isSet(object.roomId) ? String(object.roomId) : "" };
  },

  toJSON(message: RoomId): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RoomId>, I>>(object: I): RoomId {
    const message = createBaseRoomId();
    message.roomId = object.roomId ?? "";
    return message;
  },
};

function createBaseGetRoomIdReq(): GetRoomIdReq {
  return { recipientId: "" };
}

export const GetRoomIdReq = {
  encode(message: GetRoomIdReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.recipientId !== "") {
      writer.uint32(10).string(message.recipientId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRoomIdReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRoomIdReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.recipientId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetRoomIdReq {
    return { recipientId: isSet(object.recipientId) ? String(object.recipientId) : "" };
  },

  toJSON(message: GetRoomIdReq): unknown {
    const obj: any = {};
    message.recipientId !== undefined && (obj.recipientId = message.recipientId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetRoomIdReq>, I>>(object: I): GetRoomIdReq {
    const message = createBaseGetRoomIdReq();
    message.recipientId = object.recipientId ?? "";
    return message;
  },
};

function createBaseRoom(): Room {
  return {
    id: "",
    name: "",
    verified: false,
    typing: false,
    owner: "",
    type: 0,
    status: 0,
    admins: [],
    members: [],
    lastMessageText: "",
    picture: "",
    lastMessageType: 0,
    created: undefined,
    lastMessageTime: undefined,
  };
}

export const Room = {
  encode(message: Room, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.verified === true) {
      writer.uint32(32).bool(message.verified);
    }
    if (message.typing === true) {
      writer.uint32(24).bool(message.typing);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    if (message.type !== 0) {
      writer.uint32(48).int32(message.type);
    }
    if (message.status !== 0) {
      writer.uint32(64).int32(message.status);
    }
    for (const v of message.admins) {
      writer.uint32(74).string(v!);
    }
    for (const v of message.members) {
      writer.uint32(82).string(v!);
    }
    if (message.lastMessageText !== "") {
      writer.uint32(90).string(message.lastMessageText);
    }
    if (message.picture !== "") {
      writer.uint32(98).string(message.picture);
    }
    if (message.lastMessageType !== 0) {
      writer.uint32(104).int32(message.lastMessageType);
    }
    if (message.created !== undefined) {
      Timestamp.encode(toTimestamp(message.created), writer.uint32(114).fork()).ldelim();
    }
    if (message.lastMessageTime !== undefined) {
      Timestamp.encode(toTimestamp(message.lastMessageTime), writer.uint32(122).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Room {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 4:
          message.verified = reader.bool();
          break;
        case 3:
          message.typing = reader.bool();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.type = reader.int32() as any;
          break;
        case 8:
          message.status = reader.int32() as any;
          break;
        case 9:
          message.admins.push(reader.string());
          break;
        case 10:
          message.members.push(reader.string());
          break;
        case 11:
          message.lastMessageText = reader.string();
          break;
        case 12:
          message.picture = reader.string();
          break;
        case 13:
          message.lastMessageType = reader.int32() as any;
          break;
        case 14:
          message.created = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 15:
          message.lastMessageTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Room {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      verified: isSet(object.verified) ? Boolean(object.verified) : false,
      typing: isSet(object.typing) ? Boolean(object.typing) : false,
      owner: isSet(object.owner) ? String(object.owner) : "",
      type: isSet(object.type) ? roomTypeFromJSON(object.type) : 0,
      status: isSet(object.status) ? roomStatusFromJSON(object.status) : 0,
      admins: Array.isArray(object?.admins) ? object.admins.map((e: any) => String(e)) : [],
      members: Array.isArray(object?.members) ? object.members.map((e: any) => String(e)) : [],
      lastMessageText: isSet(object.lastMessageText) ? String(object.lastMessageText) : "",
      picture: isSet(object.picture) ? String(object.picture) : "",
      lastMessageType: isSet(object.lastMessageType) ? messageTypeFromJSON(object.lastMessageType) : 0,
      created: isSet(object.created) ? fromJsonTimestamp(object.created) : undefined,
      lastMessageTime: isSet(object.lastMessageTime) ? fromJsonTimestamp(object.lastMessageTime) : undefined,
    };
  },

  toJSON(message: Room): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.verified !== undefined && (obj.verified = message.verified);
    message.typing !== undefined && (obj.typing = message.typing);
    message.owner !== undefined && (obj.owner = message.owner);
    message.type !== undefined && (obj.type = roomTypeToJSON(message.type));
    message.status !== undefined && (obj.status = roomStatusToJSON(message.status));
    if (message.admins) {
      obj.admins = message.admins.map((e) => e);
    } else {
      obj.admins = [];
    }
    if (message.members) {
      obj.members = message.members.map((e) => e);
    } else {
      obj.members = [];
    }
    message.lastMessageText !== undefined && (obj.lastMessageText = message.lastMessageText);
    message.picture !== undefined && (obj.picture = message.picture);
    message.lastMessageType !== undefined && (obj.lastMessageType = messageTypeToJSON(message.lastMessageType));
    message.created !== undefined && (obj.created = message.created.toISOString());
    message.lastMessageTime !== undefined && (obj.lastMessageTime = message.lastMessageTime.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Room>, I>>(object: I): Room {
    const message = createBaseRoom();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.verified = object.verified ?? false;
    message.typing = object.typing ?? false;
    message.owner = object.owner ?? "";
    message.type = object.type ?? 0;
    message.status = object.status ?? 0;
    message.admins = object.admins?.map((e) => e) || [];
    message.members = object.members?.map((e) => e) || [];
    message.lastMessageText = object.lastMessageText ?? "";
    message.picture = object.picture ?? "";
    message.lastMessageType = object.lastMessageType ?? 0;
    message.created = object.created ?? undefined;
    message.lastMessageTime = object.lastMessageTime ?? undefined;
    return message;
  },
};

function createBaseLikeMessageReq(): LikeMessageReq {
  return { roomId: "", messageId: "" };
}

export const LikeMessageReq = {
  encode(message: LikeMessageReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.roomId !== "") {
      writer.uint32(10).string(message.roomId);
    }
    if (message.messageId !== "") {
      writer.uint32(18).string(message.messageId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikeMessageReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikeMessageReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roomId = reader.string();
          break;
        case 2:
          message.messageId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikeMessageReq {
    return {
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      messageId: isSet(object.messageId) ? String(object.messageId) : "",
    };
  },

  toJSON(message: LikeMessageReq): unknown {
    const obj: any = {};
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.messageId !== undefined && (obj.messageId = message.messageId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikeMessageReq>, I>>(object: I): LikeMessageReq {
    const message = createBaseLikeMessageReq();
    message.roomId = object.roomId ?? "";
    message.messageId = object.messageId ?? "";
    return message;
  },
};

function createBaseMessage(): Message {
  return {
    id: "",
    sender: "",
    room: undefined,
    message: "",
    senderpicture: "",
    senderName: "",
    senderUsername: "",
    created: undefined,
    photoUrl: "",
    type: 0,
    likeCount: 0,
    status: 0,
  };
}

export const Message = {
  encode(message: Message, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    if (message.room !== undefined) {
      Room.encode(message.room, writer.uint32(26).fork()).ldelim();
    }
    if (message.message !== "") {
      writer.uint32(34).string(message.message);
    }
    if (message.senderpicture !== "") {
      writer.uint32(42).string(message.senderpicture);
    }
    if (message.senderName !== "") {
      writer.uint32(50).string(message.senderName);
    }
    if (message.senderUsername !== "") {
      writer.uint32(58).string(message.senderUsername);
    }
    if (message.created !== undefined) {
      Timestamp.encode(toTimestamp(message.created), writer.uint32(66).fork()).ldelim();
    }
    if (message.photoUrl !== "") {
      writer.uint32(74).string(message.photoUrl);
    }
    if (message.type !== 0) {
      writer.uint32(80).int32(message.type);
    }
    if (message.likeCount !== 0) {
      writer.uint32(88).int32(message.likeCount);
    }
    if (message.status !== 0) {
      writer.uint32(96).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Message {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.sender = reader.string();
          break;
        case 3:
          message.room = Room.decode(reader, reader.uint32());
          break;
        case 4:
          message.message = reader.string();
          break;
        case 5:
          message.senderpicture = reader.string();
          break;
        case 6:
          message.senderName = reader.string();
          break;
        case 7:
          message.senderUsername = reader.string();
          break;
        case 8:
          message.created = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 9:
          message.photoUrl = reader.string();
          break;
        case 10:
          message.type = reader.int32() as any;
          break;
        case 11:
          message.likeCount = reader.int32();
          break;
        case 12:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Message {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
      room: isSet(object.room) ? Room.fromJSON(object.room) : undefined,
      message: isSet(object.message) ? String(object.message) : "",
      senderpicture: isSet(object.senderpicture) ? String(object.senderpicture) : "",
      senderName: isSet(object.senderName) ? String(object.senderName) : "",
      senderUsername: isSet(object.senderUsername) ? String(object.senderUsername) : "",
      created: isSet(object.created) ? fromJsonTimestamp(object.created) : undefined,
      photoUrl: isSet(object.photoUrl) ? String(object.photoUrl) : "",
      type: isSet(object.type) ? messageTypeFromJSON(object.type) : 0,
      likeCount: isSet(object.likeCount) ? Number(object.likeCount) : 0,
      status: isSet(object.status) ? message_StatusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: Message): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sender !== undefined && (obj.sender = message.sender);
    message.room !== undefined && (obj.room = message.room ? Room.toJSON(message.room) : undefined);
    message.message !== undefined && (obj.message = message.message);
    message.senderpicture !== undefined && (obj.senderpicture = message.senderpicture);
    message.senderName !== undefined && (obj.senderName = message.senderName);
    message.senderUsername !== undefined && (obj.senderUsername = message.senderUsername);
    message.created !== undefined && (obj.created = message.created.toISOString());
    message.photoUrl !== undefined && (obj.photoUrl = message.photoUrl);
    message.type !== undefined && (obj.type = messageTypeToJSON(message.type));
    message.likeCount !== undefined && (obj.likeCount = Math.round(message.likeCount));
    message.status !== undefined && (obj.status = message_StatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Message>, I>>(object: I): Message {
    const message = createBaseMessage();
    message.id = object.id ?? "";
    message.sender = object.sender ?? "";
    message.room = (object.room !== undefined && object.room !== null) ? Room.fromPartial(object.room) : undefined;
    message.message = object.message ?? "";
    message.senderpicture = object.senderpicture ?? "";
    message.senderName = object.senderName ?? "";
    message.senderUsername = object.senderUsername ?? "";
    message.created = object.created ?? undefined;
    message.photoUrl = object.photoUrl ?? "";
    message.type = object.type ?? 0;
    message.likeCount = object.likeCount ?? 0;
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseSendMessageResponse(): SendMessageResponse {
  return { status: 0 };
}

export const SendMessageResponse = {
  encode(message: SendMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SendMessageResponse {
    return { status: isSet(object.status) ? message_StatusFromJSON(object.status) : 0 };
  },

  toJSON(message: SendMessageResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = message_StatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SendMessageResponse>, I>>(object: I): SendMessageResponse {
    const message = createBaseSendMessageResponse();
    message.status = object.status ?? 0;
    return message;
  },
};

export interface ChatService {
  CreateChat(request: CreateChatReq): Promise<Room>;
  JoinRoom(request: JoinRoomRequest): Observable<Message>;
  ListRoom(request: ListRoomRequest): Observable<Room>;
  DeleteRoom(request: RoomId): Promise<Empty>;
  GetRoomId(request: GetRoomIdReq): Promise<RoomId>;
  ListMembers(request: ListMembersRequest): Promise<ListMembersResponse>;
  LeaveRoom(request: RoomId): Promise<Empty>;
  SendMessage(request: Observable<Message>): Promise<SendMessageResponse>;
  LikeMessage(request: LikeMessageReq): Promise<Empty>;
  UnlikeMessage(request: LikeMessageReq): Promise<Empty>;
  TypingChannel(request: Observable<TypingEvent>): Observable<TypingEvent>;
  TypingStart(request: TypingReq): Promise<Empty>;
  TypingEnd(request: TypingReq): Promise<Empty>;
  AddToken(request: AddTokenReq): Promise<Empty>;
  DeleteUser(request: UserIdReq): Promise<Empty>;
}

export class ChatServiceClientImpl implements ChatService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.chat.v1.ChatService";
    this.rpc = rpc;
    this.CreateChat = this.CreateChat.bind(this);
    this.JoinRoom = this.JoinRoom.bind(this);
    this.ListRoom = this.ListRoom.bind(this);
    this.DeleteRoom = this.DeleteRoom.bind(this);
    this.GetRoomId = this.GetRoomId.bind(this);
    this.ListMembers = this.ListMembers.bind(this);
    this.LeaveRoom = this.LeaveRoom.bind(this);
    this.SendMessage = this.SendMessage.bind(this);
    this.LikeMessage = this.LikeMessage.bind(this);
    this.UnlikeMessage = this.UnlikeMessage.bind(this);
    this.TypingChannel = this.TypingChannel.bind(this);
    this.TypingStart = this.TypingStart.bind(this);
    this.TypingEnd = this.TypingEnd.bind(this);
    this.AddToken = this.AddToken.bind(this);
    this.DeleteUser = this.DeleteUser.bind(this);
  }
  CreateChat(request: CreateChatReq): Promise<Room> {
    const data = CreateChatReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateChat", data);
    return promise.then((data) => Room.decode(new _m0.Reader(data)));
  }

  JoinRoom(request: JoinRoomRequest): Observable<Message> {
    const data = JoinRoomRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "JoinRoom", data);
    return result.pipe(map((data) => Message.decode(new _m0.Reader(data))));
  }

  ListRoom(request: ListRoomRequest): Observable<Room> {
    const data = ListRoomRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "ListRoom", data);
    return result.pipe(map((data) => Room.decode(new _m0.Reader(data))));
  }

  DeleteRoom(request: RoomId): Promise<Empty> {
    const data = RoomId.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteRoom", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  GetRoomId(request: GetRoomIdReq): Promise<RoomId> {
    const data = GetRoomIdReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetRoomId", data);
    return promise.then((data) => RoomId.decode(new _m0.Reader(data)));
  }

  ListMembers(request: ListMembersRequest): Promise<ListMembersResponse> {
    const data = ListMembersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListMembers", data);
    return promise.then((data) => ListMembersResponse.decode(new _m0.Reader(data)));
  }

  LeaveRoom(request: RoomId): Promise<Empty> {
    const data = RoomId.encode(request).finish();
    const promise = this.rpc.request(this.service, "LeaveRoom", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  SendMessage(request: Observable<Message>): Promise<SendMessageResponse> {
    const data = request.pipe(map((request) => Message.encode(request).finish()));
    const promise = this.rpc.clientStreamingRequest(this.service, "SendMessage", data);
    return promise.then((data) => SendMessageResponse.decode(new _m0.Reader(data)));
  }

  LikeMessage(request: LikeMessageReq): Promise<Empty> {
    const data = LikeMessageReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "LikeMessage", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  UnlikeMessage(request: LikeMessageReq): Promise<Empty> {
    const data = LikeMessageReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnlikeMessage", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  TypingChannel(request: Observable<TypingEvent>): Observable<TypingEvent> {
    const data = request.pipe(map((request) => TypingEvent.encode(request).finish()));
    const result = this.rpc.bidirectionalStreamingRequest(this.service, "TypingChannel", data);
    return result.pipe(map((data) => TypingEvent.decode(new _m0.Reader(data))));
  }

  TypingStart(request: TypingReq): Promise<Empty> {
    const data = TypingReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "TypingStart", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  TypingEnd(request: TypingReq): Promise<Empty> {
    const data = TypingReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "TypingEnd", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  AddToken(request: AddTokenReq): Promise<Empty> {
    const data = AddTokenReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "AddToken", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  DeleteUser(request: UserIdReq): Promise<Empty> {
    const data = UserIdReq.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteUser", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
