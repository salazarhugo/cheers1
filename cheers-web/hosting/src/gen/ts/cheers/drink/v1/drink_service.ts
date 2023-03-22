/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "cheers.drink.v1";

export interface CreateDrinkRequest {
  name: string;
  icon: string;
  category: string;
}

export interface CreateDrinkResponse {
  drink: Drink | undefined;
}

export interface GetDrinkRequest {
  drinkId: string;
}

export interface GetDrinkResponse {
  drink: Drink | undefined;
}

export interface ListDrinkRequest {
}

export interface DeleteDrinkRequest {
  drinkId: string;
}

export interface DeleteDrinkResponse {
}

export interface ListDrinkResponse {
  items: Drink[];
}

export interface Drink {
  id: string;
  creatorId: string;
  name: string;
  icon: string;
  category: string;
}

export interface UpdateDrinkRequest {
  drinkId: string;
  name: string;
  icon: string;
  category: string;
}

export interface UpdateDrinkResponse {
}

function createBaseCreateDrinkRequest(): CreateDrinkRequest {
  return { name: "", icon: "", category: "" };
}

export const CreateDrinkRequest = {
  encode(message: CreateDrinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.icon !== "") {
      writer.uint32(18).string(message.icon);
    }
    if (message.category !== "") {
      writer.uint32(26).string(message.category);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateDrinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateDrinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.icon = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.category = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateDrinkRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      icon: isSet(object.icon) ? String(object.icon) : "",
      category: isSet(object.category) ? String(object.category) : "",
    };
  },

  toJSON(message: CreateDrinkRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.icon !== undefined && (obj.icon = message.icon);
    message.category !== undefined && (obj.category = message.category);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateDrinkRequest>, I>>(base?: I): CreateDrinkRequest {
    return CreateDrinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateDrinkRequest>, I>>(object: I): CreateDrinkRequest {
    const message = createBaseCreateDrinkRequest();
    message.name = object.name ?? "";
    message.icon = object.icon ?? "";
    message.category = object.category ?? "";
    return message;
  },
};

function createBaseCreateDrinkResponse(): CreateDrinkResponse {
  return { drink: undefined };
}

export const CreateDrinkResponse = {
  encode(message: CreateDrinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.drink !== undefined) {
      Drink.encode(message.drink, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateDrinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateDrinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.drink = Drink.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateDrinkResponse {
    return { drink: isSet(object.drink) ? Drink.fromJSON(object.drink) : undefined };
  },

  toJSON(message: CreateDrinkResponse): unknown {
    const obj: any = {};
    message.drink !== undefined && (obj.drink = message.drink ? Drink.toJSON(message.drink) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateDrinkResponse>, I>>(base?: I): CreateDrinkResponse {
    return CreateDrinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateDrinkResponse>, I>>(object: I): CreateDrinkResponse {
    const message = createBaseCreateDrinkResponse();
    message.drink = (object.drink !== undefined && object.drink !== null) ? Drink.fromPartial(object.drink) : undefined;
    return message;
  },
};

function createBaseGetDrinkRequest(): GetDrinkRequest {
  return { drinkId: "" };
}

export const GetDrinkRequest = {
  encode(message: GetDrinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.drinkId !== "") {
      writer.uint32(10).string(message.drinkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDrinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetDrinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.drinkId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetDrinkRequest {
    return { drinkId: isSet(object.drinkId) ? String(object.drinkId) : "" };
  },

  toJSON(message: GetDrinkRequest): unknown {
    const obj: any = {};
    message.drinkId !== undefined && (obj.drinkId = message.drinkId);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetDrinkRequest>, I>>(base?: I): GetDrinkRequest {
    return GetDrinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetDrinkRequest>, I>>(object: I): GetDrinkRequest {
    const message = createBaseGetDrinkRequest();
    message.drinkId = object.drinkId ?? "";
    return message;
  },
};

function createBaseGetDrinkResponse(): GetDrinkResponse {
  return { drink: undefined };
}

export const GetDrinkResponse = {
  encode(message: GetDrinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.drink !== undefined) {
      Drink.encode(message.drink, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDrinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetDrinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.drink = Drink.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetDrinkResponse {
    return { drink: isSet(object.drink) ? Drink.fromJSON(object.drink) : undefined };
  },

  toJSON(message: GetDrinkResponse): unknown {
    const obj: any = {};
    message.drink !== undefined && (obj.drink = message.drink ? Drink.toJSON(message.drink) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GetDrinkResponse>, I>>(base?: I): GetDrinkResponse {
    return GetDrinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetDrinkResponse>, I>>(object: I): GetDrinkResponse {
    const message = createBaseGetDrinkResponse();
    message.drink = (object.drink !== undefined && object.drink !== null) ? Drink.fromPartial(object.drink) : undefined;
    return message;
  },
};

function createBaseListDrinkRequest(): ListDrinkRequest {
  return {};
}

export const ListDrinkRequest = {
  encode(_: ListDrinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDrinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDrinkRequest();
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

  fromJSON(_: any): ListDrinkRequest {
    return {};
  },

  toJSON(_: ListDrinkRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListDrinkRequest>, I>>(base?: I): ListDrinkRequest {
    return ListDrinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListDrinkRequest>, I>>(_: I): ListDrinkRequest {
    const message = createBaseListDrinkRequest();
    return message;
  },
};

function createBaseDeleteDrinkRequest(): DeleteDrinkRequest {
  return { drinkId: "" };
}

export const DeleteDrinkRequest = {
  encode(message: DeleteDrinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.drinkId !== "") {
      writer.uint32(10).string(message.drinkId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteDrinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteDrinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.drinkId = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteDrinkRequest {
    return { drinkId: isSet(object.drinkId) ? String(object.drinkId) : "" };
  },

  toJSON(message: DeleteDrinkRequest): unknown {
    const obj: any = {};
    message.drinkId !== undefined && (obj.drinkId = message.drinkId);
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteDrinkRequest>, I>>(base?: I): DeleteDrinkRequest {
    return DeleteDrinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteDrinkRequest>, I>>(object: I): DeleteDrinkRequest {
    const message = createBaseDeleteDrinkRequest();
    message.drinkId = object.drinkId ?? "";
    return message;
  },
};

function createBaseDeleteDrinkResponse(): DeleteDrinkResponse {
  return {};
}

export const DeleteDrinkResponse = {
  encode(_: DeleteDrinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteDrinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteDrinkResponse();
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

  fromJSON(_: any): DeleteDrinkResponse {
    return {};
  },

  toJSON(_: DeleteDrinkResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteDrinkResponse>, I>>(base?: I): DeleteDrinkResponse {
    return DeleteDrinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteDrinkResponse>, I>>(_: I): DeleteDrinkResponse {
    const message = createBaseDeleteDrinkResponse();
    return message;
  },
};

function createBaseListDrinkResponse(): ListDrinkResponse {
  return { items: [] };
}

export const ListDrinkResponse = {
  encode(message: ListDrinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      Drink.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDrinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDrinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.items.push(Drink.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListDrinkResponse {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => Drink.fromJSON(e)) : [] };
  },

  toJSON(message: ListDrinkResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) => e ? Drink.toJSON(e) : undefined);
    } else {
      obj.items = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListDrinkResponse>, I>>(base?: I): ListDrinkResponse {
    return ListDrinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListDrinkResponse>, I>>(object: I): ListDrinkResponse {
    const message = createBaseListDrinkResponse();
    message.items = object.items?.map((e) => Drink.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDrink(): Drink {
  return { id: "", creatorId: "", name: "", icon: "", category: "" };
}

export const Drink = {
  encode(message: Drink, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.creatorId !== "") {
      writer.uint32(18).string(message.creatorId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.icon !== "") {
      writer.uint32(34).string(message.icon);
    }
    if (message.category !== "") {
      writer.uint32(42).string(message.category);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Drink {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDrink();
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
          if (tag != 18) {
            break;
          }

          message.creatorId = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.icon = reader.string();
          continue;
        case 5:
          if (tag != 42) {
            break;
          }

          message.category = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Drink {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      creatorId: isSet(object.creatorId) ? String(object.creatorId) : "",
      name: isSet(object.name) ? String(object.name) : "",
      icon: isSet(object.icon) ? String(object.icon) : "",
      category: isSet(object.category) ? String(object.category) : "",
    };
  },

  toJSON(message: Drink): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creatorId !== undefined && (obj.creatorId = message.creatorId);
    message.name !== undefined && (obj.name = message.name);
    message.icon !== undefined && (obj.icon = message.icon);
    message.category !== undefined && (obj.category = message.category);
    return obj;
  },

  create<I extends Exact<DeepPartial<Drink>, I>>(base?: I): Drink {
    return Drink.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Drink>, I>>(object: I): Drink {
    const message = createBaseDrink();
    message.id = object.id ?? "";
    message.creatorId = object.creatorId ?? "";
    message.name = object.name ?? "";
    message.icon = object.icon ?? "";
    message.category = object.category ?? "";
    return message;
  },
};

function createBaseUpdateDrinkRequest(): UpdateDrinkRequest {
  return { drinkId: "", name: "", icon: "", category: "" };
}

export const UpdateDrinkRequest = {
  encode(message: UpdateDrinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.drinkId !== "") {
      writer.uint32(10).string(message.drinkId);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.icon !== "") {
      writer.uint32(26).string(message.icon);
    }
    if (message.category !== "") {
      writer.uint32(34).string(message.category);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateDrinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateDrinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.drinkId = reader.string();
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.icon = reader.string();
          continue;
        case 4:
          if (tag != 34) {
            break;
          }

          message.category = reader.string();
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateDrinkRequest {
    return {
      drinkId: isSet(object.drinkId) ? String(object.drinkId) : "",
      name: isSet(object.name) ? String(object.name) : "",
      icon: isSet(object.icon) ? String(object.icon) : "",
      category: isSet(object.category) ? String(object.category) : "",
    };
  },

  toJSON(message: UpdateDrinkRequest): unknown {
    const obj: any = {};
    message.drinkId !== undefined && (obj.drinkId = message.drinkId);
    message.name !== undefined && (obj.name = message.name);
    message.icon !== undefined && (obj.icon = message.icon);
    message.category !== undefined && (obj.category = message.category);
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateDrinkRequest>, I>>(base?: I): UpdateDrinkRequest {
    return UpdateDrinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateDrinkRequest>, I>>(object: I): UpdateDrinkRequest {
    const message = createBaseUpdateDrinkRequest();
    message.drinkId = object.drinkId ?? "";
    message.name = object.name ?? "";
    message.icon = object.icon ?? "";
    message.category = object.category ?? "";
    return message;
  },
};

function createBaseUpdateDrinkResponse(): UpdateDrinkResponse {
  return {};
}

export const UpdateDrinkResponse = {
  encode(_: UpdateDrinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateDrinkResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateDrinkResponse();
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

  fromJSON(_: any): UpdateDrinkResponse {
    return {};
  },

  toJSON(_: UpdateDrinkResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateDrinkResponse>, I>>(base?: I): UpdateDrinkResponse {
    return UpdateDrinkResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateDrinkResponse>, I>>(_: I): UpdateDrinkResponse {
    const message = createBaseUpdateDrinkResponse();
    return message;
  },
};

export interface DrinkService {
  CreateDrink(request: CreateDrinkRequest): Promise<CreateDrinkResponse>;
  GetDrink(request: GetDrinkRequest): Promise<GetDrinkResponse>;
  UpdateDrink(request: UpdateDrinkRequest): Promise<UpdateDrinkResponse>;
  ListDrink(request: ListDrinkRequest): Promise<ListDrinkResponse>;
  DeleteDrink(request: DeleteDrinkRequest): Promise<DeleteDrinkResponse>;
}

export class DrinkServiceClientImpl implements DrinkService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "cheers.drink.v1.DrinkService";
    this.rpc = rpc;
    this.CreateDrink = this.CreateDrink.bind(this);
    this.GetDrink = this.GetDrink.bind(this);
    this.UpdateDrink = this.UpdateDrink.bind(this);
    this.ListDrink = this.ListDrink.bind(this);
    this.DeleteDrink = this.DeleteDrink.bind(this);
  }
  CreateDrink(request: CreateDrinkRequest): Promise<CreateDrinkResponse> {
    const data = CreateDrinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateDrink", data);
    return promise.then((data) => CreateDrinkResponse.decode(_m0.Reader.create(data)));
  }

  GetDrink(request: GetDrinkRequest): Promise<GetDrinkResponse> {
    const data = GetDrinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetDrink", data);
    return promise.then((data) => GetDrinkResponse.decode(_m0.Reader.create(data)));
  }

  UpdateDrink(request: UpdateDrinkRequest): Promise<UpdateDrinkResponse> {
    const data = UpdateDrinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateDrink", data);
    return promise.then((data) => UpdateDrinkResponse.decode(_m0.Reader.create(data)));
  }

  ListDrink(request: ListDrinkRequest): Promise<ListDrinkResponse> {
    const data = ListDrinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListDrink", data);
    return promise.then((data) => ListDrinkResponse.decode(_m0.Reader.create(data)));
  }

  DeleteDrink(request: DeleteDrinkRequest): Promise<DeleteDrinkResponse> {
    const data = DeleteDrinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteDrink", data);
    return promise.then((data) => DeleteDrinkResponse.decode(_m0.Reader.create(data)));
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
