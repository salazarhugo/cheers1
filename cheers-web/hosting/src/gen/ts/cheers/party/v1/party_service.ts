/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Party } from "../../type/party/party";
import { User, UserItem } from "../../type/user/user";

export const protobufPackage = "cheers.party.v1";

export enum WatchStatus {
  UNWATCHED = 0,
  GOING = 1,
  INTERESTED = 2,
  MAYBE = 3,
  UNRECOGNIZED = -1,
}

export function watchStatusFromJSON(object: any): WatchStatus {
  switch (object) {
    case 0:
    case "UNWATCHED":
      return WatchStatus.UNWATCHED;
    case 1:
    case "GOING":
      return WatchStatus.GOING;
    case 2:
    case "INTERESTED":
      return WatchStatus.INTERESTED;
    case 3:
    case "MAYBE":
      return WatchStatus.MAYBE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return WatchStatus.UNRECOGNIZED;
  }
}

export function watchStatusToJSON(object: WatchStatus): string {
  switch (object) {
    case WatchStatus.UNWATCHED:
      return "UNWATCHED";
    case WatchStatus.GOING:
      return "GOING";
    case WatchStatus.INTERESTED:
      return "INTERESTED";
    case WatchStatus.MAYBE:
      return "MAYBE";
    case WatchStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface DuplicatePartyRequest {
  partyId: string;
}

export interface DuplicatePartyResponse {
}

export interface TransferPartyRequest {
  userId: string;
  partyId: string;
}

export interface TransferPartyResponse {
}

export interface ListPartyRequest {
  page: number;
  pageSize: number;
  userId: string;
}

export interface ListPartyResponse {
  items: PartyItem[];
  nextPageToken: string;
}

export interface ListGoingRequest {
  partyId: string;
}

export interface ListGoingResponse {
  users: UserItem[];
}

export interface AnswerPartyRequest {
  partyId: string;
  watchStatus: WatchStatus;
}

export interface AnswerPartyResponse {
}

export interface CreatePartyRequest {
  party: Party | undefined;
}

export interface CreatePartyResponse {
  party: Party | undefined;
}

export interface GetPartyRequest {
  partyId: string;
}

export interface GetPartyResponse {
  party: Party | undefined;
}

export interface GetPartyItemRequest {
  partyId: string;
}

export interface GetPartyItemResponse {
  item: PartyItem | undefined;
}

export interface UpdatePartyRequest {
  party: Party | undefined;
}

export interface UpdatePartyResponse {
  party: Party | undefined;
}

export interface DeletePartyRequest {
  partyId: string;
}

export interface DeletePartyResponse {
}

export interface FeedPartyRequest {
  parent: string;
  pageSize: number;
  pageToken: string;
}

export interface FeedPartyResponse {
  items: PartyItem[];
  nextPageToken: string;
}

export interface PartyItem {
  party: Party | undefined;
  user: User | undefined;
  goingCount: number;
  interestedCount: number;
  invitedCount: number;
  isCreator: boolean;
  viewerWatchStatus: WatchStatus;
  mutualUsernames: string[];
  mutualPictures: string[];
}

function createBaseDuplicatePartyRequest(): DuplicatePartyRequest {
  return { partyId: "" };
}

export const DuplicatePartyRequest = {
  encode(message: DuplicatePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DuplicatePartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDuplicatePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DuplicatePartyRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: DuplicatePartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DuplicatePartyRequest>, I>>(base?: I): DuplicatePartyRequest {
    return DuplicatePartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DuplicatePartyRequest>, I>>(object: I): DuplicatePartyRequest {
    const message = createBaseDuplicatePartyRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseDuplicatePartyResponse(): DuplicatePartyResponse {
  return {};
}

export const DuplicatePartyResponse = {
  encode(_: DuplicatePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DuplicatePartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDuplicatePartyResponse();
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

  fromJSON(_: any): DuplicatePartyResponse {
    return {};
  },

  toJSON(_: DuplicatePartyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DuplicatePartyResponse>, I>>(base?: I): DuplicatePartyResponse {
    return DuplicatePartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DuplicatePartyResponse>, I>>(_: I): DuplicatePartyResponse {
    const message = createBaseDuplicatePartyResponse();
    return message;
  },
};

function createBaseTransferPartyRequest(): TransferPartyRequest {
  return { userId: "", partyId: "" };
}

export const TransferPartyRequest = {
  encode(message: TransferPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.partyId !== "") {
      writer.uint32(18).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransferPartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransferPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransferPartyRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
    };
  },

  toJSON(message: TransferPartyRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<TransferPartyRequest>, I>>(base?: I): TransferPartyRequest {
    return TransferPartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TransferPartyRequest>, I>>(object: I): TransferPartyRequest {
    const message = createBaseTransferPartyRequest();
    message.userId = object.userId ?? "";
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseTransferPartyResponse(): TransferPartyResponse {
  return {};
}

export const TransferPartyResponse = {
  encode(_: TransferPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransferPartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransferPartyResponse();
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

  fromJSON(_: any): TransferPartyResponse {
    return {};
  },

  toJSON(_: TransferPartyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<TransferPartyResponse>, I>>(base?: I): TransferPartyResponse {
    return TransferPartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TransferPartyResponse>, I>>(_: I): TransferPartyResponse {
    const message = createBaseTransferPartyResponse();
    return message;
  },
};

function createBaseListPartyRequest(): ListPartyRequest {
  return { page: 0, pageSize: 0, userId: "" };
}

export const ListPartyRequest = {
  encode(message: ListPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.page !== 0) {
      writer.uint32(8).int32(message.page);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.userId !== "") {
      writer.uint32(26).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 8) {
            break;
          }

          message.page = reader.int32();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag != 26) {
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

  fromJSON(object: any): ListPartyRequest {
    return {
      page: isSet(object.page) ? Number(object.page) : 0,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: ListPartyRequest): unknown {
    const obj: any = {};
    message.page !== undefined && (obj.page = Math.round(message.page));
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPartyRequest>, I>>(base?: I): ListPartyRequest {
    return ListPartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPartyRequest>, I>>(object: I): ListPartyRequest {
    const message = createBaseListPartyRequest();
    message.page = object.page ?? 0;
    message.pageSize = object.pageSize ?? 0;
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseListPartyResponse(): ListPartyResponse {
  return { items: [], nextPageToken: "" };
}

export const ListPartyResponse = {
  encode(message: ListPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      PartyItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.items.push(PartyItem.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPartyResponse {
    return {
      items: Array.isArray(object?.items) ? object.items.map((e: any) => PartyItem.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: ListPartyResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? PartyItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPartyResponse>, I>>(base?: I): ListPartyResponse {
    return ListPartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPartyResponse>, I>>(object: I): ListPartyResponse {
    const message = createBaseListPartyResponse();
    message.items = object.items?.map((e) => PartyItem.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBaseListGoingRequest(): ListGoingRequest {
  return { partyId: "" };
}

export const ListGoingRequest = {
  encode(message: ListGoingRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListGoingRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListGoingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListGoingRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: ListGoingRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<ListGoingRequest>, I>>(base?: I): ListGoingRequest {
    return ListGoingRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListGoingRequest>, I>>(object: I): ListGoingRequest {
    const message = createBaseListGoingRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseListGoingResponse(): ListGoingResponse {
  return { users: [] };
}

export const ListGoingResponse = {
  encode(message: ListGoingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.users) {
      UserItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListGoingResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListGoingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.users.push(UserItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListGoingResponse {
    return { users: Array.isArray(object?.users) ? object.users.map((e: any) => UserItem.fromJSON(e)) : [] };
  },

  toJSON(message: ListGoingResponse): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => e ? UserItem.toJSON(e) : undefined);
    } else {
      obj.users = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListGoingResponse>, I>>(base?: I): ListGoingResponse {
    return ListGoingResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListGoingResponse>, I>>(object: I): ListGoingResponse {
    const message = createBaseListGoingResponse();
    message.users = object.users?.map((e) => UserItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAnswerPartyRequest(): AnswerPartyRequest {
  return { partyId: "", watchStatus: 0 };
}

export const AnswerPartyRequest = {
  encode(message: AnswerPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    if (message.watchStatus !== 0) {
      writer.uint32(16).int32(message.watchStatus);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AnswerPartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnswerPartyRequest();
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
          if (tag != 16) {
            break;
          }

          message.watchStatus = reader.int32() as any;
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AnswerPartyRequest {
    return {
      partyId: isSet(object.partyId) ? String(object.partyId) : "",
      watchStatus: isSet(object.watchStatus) ? watchStatusFromJSON(object.watchStatus) : 0,
    };
  },

  toJSON(message: AnswerPartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    message.watchStatus !== undefined && (obj.watchStatus = watchStatusToJSON(message.watchStatus));
    return obj;
  },

  create<I extends Exact<DeepPartial<AnswerPartyRequest>, I>>(base?: I): AnswerPartyRequest {
    return AnswerPartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AnswerPartyRequest>, I>>(object: I): AnswerPartyRequest {
    const message = createBaseAnswerPartyRequest();
    message.partyId = object.partyId ?? "";
    message.watchStatus = object.watchStatus ?? 0;
    return message;
  },
};

function createBaseAnswerPartyResponse(): AnswerPartyResponse {
  return {};
}

export const AnswerPartyResponse = {
  encode(_: AnswerPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AnswerPartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnswerPartyResponse();
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

  fromJSON(_: any): AnswerPartyResponse {
    return {};
  },

  toJSON(_: AnswerPartyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<AnswerPartyResponse>, I>>(base?: I): AnswerPartyResponse {
    return AnswerPartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AnswerPartyResponse>, I>>(_: I): AnswerPartyResponse {
    const message = createBaseAnswerPartyResponse();
    return message;
  },
};

function createBaseCreatePartyRequest(): CreatePartyRequest {
  return { party: undefined };
}

export const CreatePartyRequest = {
  encode(message: CreatePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePartyRequest {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: CreatePartyRequest): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePartyRequest>, I>>(base?: I): CreatePartyRequest {
    return CreatePartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePartyRequest>, I>>(object: I): CreatePartyRequest {
    const message = createBaseCreatePartyRequest();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseCreatePartyResponse(): CreatePartyResponse {
  return { party: undefined };
}

export const CreatePartyResponse = {
  encode(message: CreatePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: CreatePartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePartyResponse>, I>>(base?: I): CreatePartyResponse {
    return CreatePartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreatePartyResponse>, I>>(object: I): CreatePartyResponse {
    const message = createBaseCreatePartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseGetPartyRequest(): GetPartyRequest {
  return { partyId: "" };
}

export const GetPartyRequest = {
  encode(message: GetPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPartyRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: GetPartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPartyRequest>, I>>(base?: I): GetPartyRequest {
    return GetPartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyRequest>, I>>(object: I): GetPartyRequest {
    const message = createBaseGetPartyRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseGetPartyResponse(): GetPartyResponse {
  return { party: undefined };
}

export const GetPartyResponse = {
  encode(message: GetPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: GetPartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPartyResponse>, I>>(base?: I): GetPartyResponse {
    return GetPartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyResponse>, I>>(object: I): GetPartyResponse {
    const message = createBaseGetPartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseGetPartyItemRequest(): GetPartyItemRequest {
  return { partyId: "" };
}

export const GetPartyItemRequest = {
  encode(message: GetPartyItemRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyItemRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyItemRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPartyItemRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: GetPartyItemRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPartyItemRequest>, I>>(base?: I): GetPartyItemRequest {
    return GetPartyItemRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyItemRequest>, I>>(object: I): GetPartyItemRequest {
    const message = createBaseGetPartyItemRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseGetPartyItemResponse(): GetPartyItemResponse {
  return { item: undefined };
}

export const GetPartyItemResponse = {
  encode(message: GetPartyItemResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.item !== undefined) {
      PartyItem.encode(message.item, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPartyItemResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPartyItemResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.item = PartyItem.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPartyItemResponse {
    return { item: isSet(object.item) ? PartyItem.fromJSON(object.item) : undefined };
  },

  toJSON(message: GetPartyItemResponse): unknown {
    const obj: any = {};
    message.item !== undefined && (obj.item = message.item ? PartyItem.toJSON(message.item) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPartyItemResponse>, I>>(base?: I): GetPartyItemResponse {
    return GetPartyItemResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPartyItemResponse>, I>>(object: I): GetPartyItemResponse {
    const message = createBaseGetPartyItemResponse();
    message.item = (object.item !== undefined && object.item !== null) ? PartyItem.fromPartial(object.item) : undefined;
    return message;
  },
};

function createBaseUpdatePartyRequest(): UpdatePartyRequest {
  return { party: undefined };
}

export const UpdatePartyRequest = {
  encode(message: UpdatePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdatePartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdatePartyRequest {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: UpdatePartyRequest): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdatePartyRequest>, I>>(base?: I): UpdatePartyRequest {
    return UpdatePartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdatePartyRequest>, I>>(object: I): UpdatePartyRequest {
    const message = createBaseUpdatePartyRequest();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseUpdatePartyResponse(): UpdatePartyResponse {
  return { party: undefined };
}

export const UpdatePartyResponse = {
  encode(message: UpdatePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdatePartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdatePartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdatePartyResponse {
    return { party: isSet(object.party) ? Party.fromJSON(object.party) : undefined };
  },

  toJSON(message: UpdatePartyResponse): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdatePartyResponse>, I>>(base?: I): UpdatePartyResponse {
    return UpdatePartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdatePartyResponse>, I>>(object: I): UpdatePartyResponse {
    const message = createBaseUpdatePartyResponse();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    return message;
  },
};

function createBaseDeletePartyRequest(): DeletePartyRequest {
  return { partyId: "" };
}

export const DeletePartyRequest = {
  encode(message: DeletePartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.partyId !== "") {
      writer.uint32(10).string(message.partyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.partyId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletePartyRequest {
    return { partyId: isSet(object.partyId) ? String(object.partyId) : "" };
  },

  toJSON(message: DeletePartyRequest): unknown {
    const obj: any = {};
    message.partyId !== undefined && (obj.partyId = message.partyId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletePartyRequest>, I>>(base?: I): DeletePartyRequest {
    return DeletePartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletePartyRequest>, I>>(object: I): DeletePartyRequest {
    const message = createBaseDeletePartyRequest();
    message.partyId = object.partyId ?? "";
    return message;
  },
};

function createBaseDeletePartyResponse(): DeletePartyResponse {
  return {};
}

export const DeletePartyResponse = {
  encode(_: DeletePartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePartyResponse();
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

  fromJSON(_: any): DeletePartyResponse {
    return {};
  },

  toJSON(_: DeletePartyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletePartyResponse>, I>>(base?: I): DeletePartyResponse {
    return DeletePartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletePartyResponse>, I>>(_: I): DeletePartyResponse {
    const message = createBaseDeletePartyResponse();
    return message;
  },
};

function createBaseFeedPartyRequest(): FeedPartyRequest {
  return { parent: "", pageSize: 0, pageToken: "" };
}

export const FeedPartyRequest = {
  encode(message: FeedPartyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.parent !== "") {
      writer.uint32(10).string(message.parent);
    }
    if (message.pageSize !== 0) {
      writer.uint32(16).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(26).string(message.pageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPartyRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPartyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.parent = reader.string();
          continue;
        case 2:
          if (tag != 16) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.pageToken = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FeedPartyRequest {
    return {
      parent: isSet(object.parent) ? String(object.parent) : "",
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? String(object.pageToken) : "",
    };
  },

  toJSON(message: FeedPartyRequest): unknown {
    const obj: any = {};
    message.parent !== undefined && (obj.parent = message.parent);
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.pageToken !== undefined && (obj.pageToken = message.pageToken);
    return obj;
  },

  create<I extends Exact<DeepPartial<FeedPartyRequest>, I>>(base?: I): FeedPartyRequest {
    return FeedPartyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FeedPartyRequest>, I>>(object: I): FeedPartyRequest {
    const message = createBaseFeedPartyRequest();
    message.parent = object.parent ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseFeedPartyResponse(): FeedPartyResponse {
  return { items: [], nextPageToken: "" };
}

export const FeedPartyResponse = {
  encode(message: FeedPartyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      PartyItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeedPartyResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeedPartyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.items.push(PartyItem.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FeedPartyResponse {
    return {
      items: Array.isArray(object?.items) ? object.items.map((e: any) => PartyItem.fromJSON(e)) : [],
      nextPageToken: isSet(object.nextPageToken) ? String(object.nextPageToken) : "",
    };
  },

  toJSON(message: FeedPartyResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? PartyItem.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    message.nextPageToken !== undefined && (obj.nextPageToken = message.nextPageToken);
    return obj;
  },

  create<I extends Exact<DeepPartial<FeedPartyResponse>, I>>(base?: I): FeedPartyResponse {
    return FeedPartyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FeedPartyResponse>, I>>(object: I): FeedPartyResponse {
    const message = createBaseFeedPartyResponse();
    message.items = object.items?.map((e) => PartyItem.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

function createBasePartyItem(): PartyItem {
  return {
    party: undefined,
    user: undefined,
    goingCount: 0,
    interestedCount: 0,
    invitedCount: 0,
    isCreator: false,
    viewerWatchStatus: 0,
    mutualUsernames: [],
    mutualPictures: [],
  };
}

export const PartyItem = {
  encode(message: PartyItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.party !== undefined) {
      Party.encode(message.party, writer.uint32(10).fork()).ldelim();
    }
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(18).fork()).ldelim();
    }
    if (message.goingCount !== 0) {
      writer.uint32(24).int64(message.goingCount);
    }
    if (message.interestedCount !== 0) {
      writer.uint32(32).int64(message.interestedCount);
    }
    if (message.invitedCount !== 0) {
      writer.uint32(40).int64(message.invitedCount);
    }
    if (message.isCreator === true) {
      writer.uint32(64).bool(message.isCreator);
    }
    if (message.viewerWatchStatus !== 0) {
      writer.uint32(72).int32(message.viewerWatchStatus);
    }
    for (const v of message.mutualUsernames) {
      writer.uint32(82).string(v!);
    }
    for (const v of message.mutualPictures) {
      writer.uint32(90).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PartyItem {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePartyItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.party = Party.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag != 24) {
            break;
          }

          message.goingCount = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag != 32) {
            break;
          }

          message.interestedCount = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag != 40) {
            break;
          }

          message.invitedCount = longToNumber(reader.int64() as Long);
          continue;
        case 8:
          if (tag != 64) {
            break;
          }

          message.isCreator = reader.bool();
          continue;
        case 9:
          if (tag != 72) {
            break;
          }

          message.viewerWatchStatus = reader.int32() as any;
          continue;
        case 10:
          if (tag != 82) {
            break;
          }

          message.mutualUsernames.push(reader.string());
          continue;
        case 11:
          if (tag != 90) {
            break;
          }

          message.mutualPictures.push(reader.string());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PartyItem {
    return {
      party: isSet(object.party) ? Party.fromJSON(object.party) : undefined,
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      goingCount: isSet(object.goingCount) ? Number(object.goingCount) : 0,
      interestedCount: isSet(object.interestedCount) ? Number(object.interestedCount) : 0,
      invitedCount: isSet(object.invitedCount) ? Number(object.invitedCount) : 0,
      isCreator: isSet(object.isCreator) ? Boolean(object.isCreator) : false,
      viewerWatchStatus: isSet(object.viewerWatchStatus) ? watchStatusFromJSON(object.viewerWatchStatus) : 0,
      mutualUsernames: Array.isArray(object?.mutualUsernames) ? object.mutualUsernames.map((e: any) => String(e)) : [],
      mutualPictures: Array.isArray(object?.mutualPictures) ? object.mutualPictures.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: PartyItem): unknown {
    const obj: any = {};
    message.party !== undefined && (obj.party = message.party ? Party.toJSON(message.party) : undefined);
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.goingCount !== undefined && (obj.goingCount = Math.round(message.goingCount));
    message.interestedCount !== undefined && (obj.interestedCount = Math.round(message.interestedCount));
    message.invitedCount !== undefined && (obj.invitedCount = Math.round(message.invitedCount));
    message.isCreator !== undefined && (obj.isCreator = message.isCreator);
    message.viewerWatchStatus !== undefined && (obj.viewerWatchStatus = watchStatusToJSON(message.viewerWatchStatus));
    if (message.mutualUsernames) {
      obj.mutualUsernames = message.mutualUsernames.map((e) => e);
    } else {
      obj.mutualUsernames = [];
    }
    if (message.mutualPictures) {
      obj.mutualPictures = message.mutualPictures.map((e) => e);
    } else {
      obj.mutualPictures = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PartyItem>, I>>(base?: I): PartyItem {
    return PartyItem.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PartyItem>, I>>(object: I): PartyItem {
    const message = createBasePartyItem();
    message.party = (object.party !== undefined && object.party !== null) ? Party.fromPartial(object.party) : undefined;
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.goingCount = object.goingCount ?? 0;
    message.interestedCount = object.interestedCount ?? 0;
    message.invitedCount = object.invitedCount ?? 0;
    message.isCreator = object.isCreator ?? false;
    message.viewerWatchStatus = object.viewerWatchStatus ?? 0;
    message.mutualUsernames = object.mutualUsernames?.map((e) => e) || [];
    message.mutualPictures = object.mutualPictures?.map((e) => e) || [];
    return message;
  },
};

export interface PartyService {
  CreateParty(request: CreatePartyRequest): Promise<CreatePartyResponse>;
  GetParty(request: GetPartyRequest): Promise<GetPartyResponse>;
  UpdateParty(request: UpdatePartyRequest): Promise<UpdatePartyResponse>;
  DeleteParty(request: DeletePartyRequest): Promise<DeletePartyResponse>;
  GetPartyItem(request: GetPartyItemRequest): Promise<GetPartyItemResponse>;
  FeedParty(request: FeedPartyRequest): Promise<FeedPartyResponse>;
  /** List parties of a specific user */
  ListParty(request: ListPartyRequest): Promise<ListPartyResponse>;
  AnswerParty(request: AnswerPartyRequest): Promise<AnswerPartyResponse>;
  ListGoing(request: ListGoingRequest): Promise<ListGoingResponse>;
  TransferParty(request: TransferPartyRequest): Promise<TransferPartyResponse>;
  DuplicateParty(request: DuplicatePartyRequest): Promise<DuplicatePartyResponse>;
}

export class PartyServiceClientImpl implements PartyService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.party.v1.PartyService";
    this.rpc = rpc;
    this.CreateParty = this.CreateParty.bind(this);
    this.GetParty = this.GetParty.bind(this);
    this.UpdateParty = this.UpdateParty.bind(this);
    this.DeleteParty = this.DeleteParty.bind(this);
    this.GetPartyItem = this.GetPartyItem.bind(this);
    this.FeedParty = this.FeedParty.bind(this);
    this.ListParty = this.ListParty.bind(this);
    this.AnswerParty = this.AnswerParty.bind(this);
    this.ListGoing = this.ListGoing.bind(this);
    this.TransferParty = this.TransferParty.bind(this);
    this.DuplicateParty = this.DuplicateParty.bind(this);
  }
  CreateParty(request: CreatePartyRequest): Promise<CreatePartyResponse> {
    const data = CreatePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateParty", data);
    return promise.then((data) => CreatePartyResponse.decode(_m0.Reader.create(data)));
  }

  GetParty(request: GetPartyRequest): Promise<GetPartyResponse> {
    const data = GetPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetParty", data);
    return promise.then((data) => GetPartyResponse.decode(_m0.Reader.create(data)));
  }

  UpdateParty(request: UpdatePartyRequest): Promise<UpdatePartyResponse> {
    const data = UpdatePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParty", data);
    return promise.then((data) => UpdatePartyResponse.decode(_m0.Reader.create(data)));
  }

  DeleteParty(request: DeletePartyRequest): Promise<DeletePartyResponse> {
    const data = DeletePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteParty", data);
    return promise.then((data) => DeletePartyResponse.decode(_m0.Reader.create(data)));
  }

  GetPartyItem(request: GetPartyItemRequest): Promise<GetPartyItemResponse> {
    const data = GetPartyItemRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetPartyItem", data);
    return promise.then((data) => GetPartyItemResponse.decode(_m0.Reader.create(data)));
  }

  FeedParty(request: FeedPartyRequest): Promise<FeedPartyResponse> {
    const data = FeedPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeedParty", data);
    return promise.then((data) => FeedPartyResponse.decode(_m0.Reader.create(data)));
  }

  ListParty(request: ListPartyRequest): Promise<ListPartyResponse> {
    const data = ListPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListParty", data);
    return promise.then((data) => ListPartyResponse.decode(_m0.Reader.create(data)));
  }

  AnswerParty(request: AnswerPartyRequest): Promise<AnswerPartyResponse> {
    const data = AnswerPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AnswerParty", data);
    return promise.then((data) => AnswerPartyResponse.decode(_m0.Reader.create(data)));
  }

  ListGoing(request: ListGoingRequest): Promise<ListGoingResponse> {
    const data = ListGoingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListGoing", data);
    return promise.then((data) => ListGoingResponse.decode(_m0.Reader.create(data)));
  }

  TransferParty(request: TransferPartyRequest): Promise<TransferPartyResponse> {
    const data = TransferPartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TransferParty", data);
    return promise.then((data) => TransferPartyResponse.decode(_m0.Reader.create(data)));
  }

  DuplicateParty(request: DuplicatePartyRequest): Promise<DuplicatePartyResponse> {
    const data = DuplicatePartyRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DuplicateParty", data);
    return promise.then((data) => DuplicatePartyResponse.decode(_m0.Reader.create(data)));
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
