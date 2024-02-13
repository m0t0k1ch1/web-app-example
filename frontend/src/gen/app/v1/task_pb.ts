// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file app/v1/task.proto (package app.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { FieldMask, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum app.v1.TaskStatus
 */
export enum TaskStatus {
  /**
   * @generated from enum value: TASK_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: TASK_STATUS_UNCOMPLETED = 1;
   */
  UNCOMPLETED = 1,

  /**
   * @generated from enum value: TASK_STATUS_COMPLETED = 2;
   */
  COMPLETED = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(TaskStatus)
proto3.util.setEnumType(TaskStatus, "app.v1.TaskStatus", [
  { no: 0, name: "TASK_STATUS_UNSPECIFIED" },
  { no: 1, name: "TASK_STATUS_UNCOMPLETED" },
  { no: 2, name: "TASK_STATUS_COMPLETED" },
]);

/**
 * @generated from message app.v1.Task
 */
export class Task extends Message<Task> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * @generated from field: app.v1.TaskStatus status = 3;
   */
  status = TaskStatus.UNSPECIFIED;

  constructor(data?: PartialMessage<Task>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.Task";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(TaskStatus) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Task {
    return new Task().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Task {
    return new Task().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Task {
    return new Task().fromJsonString(jsonString, options);
  }

  static equals(a: Task | PlainMessage<Task> | undefined, b: Task | PlainMessage<Task> | undefined): boolean {
    return proto3.util.equals(Task, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceCreateRequest
 */
export class TaskServiceCreateRequest extends Message<TaskServiceCreateRequest> {
  /**
   * @generated from field: string title = 1;
   */
  title = "";

  constructor(data?: PartialMessage<TaskServiceCreateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceCreateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceCreateRequest {
    return new TaskServiceCreateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceCreateRequest {
    return new TaskServiceCreateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceCreateRequest {
    return new TaskServiceCreateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceCreateRequest | PlainMessage<TaskServiceCreateRequest> | undefined, b: TaskServiceCreateRequest | PlainMessage<TaskServiceCreateRequest> | undefined): boolean {
    return proto3.util.equals(TaskServiceCreateRequest, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceCreateResponse
 */
export class TaskServiceCreateResponse extends Message<TaskServiceCreateResponse> {
  /**
   * @generated from field: app.v1.Task task = 1;
   */
  task?: Task;

  constructor(data?: PartialMessage<TaskServiceCreateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceCreateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceCreateResponse {
    return new TaskServiceCreateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceCreateResponse {
    return new TaskServiceCreateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceCreateResponse {
    return new TaskServiceCreateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceCreateResponse | PlainMessage<TaskServiceCreateResponse> | undefined, b: TaskServiceCreateResponse | PlainMessage<TaskServiceCreateResponse> | undefined): boolean {
    return proto3.util.equals(TaskServiceCreateResponse, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceGetRequest
 */
export class TaskServiceGetRequest extends Message<TaskServiceGetRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<TaskServiceGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceGetRequest {
    return new TaskServiceGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceGetRequest {
    return new TaskServiceGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceGetRequest {
    return new TaskServiceGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceGetRequest | PlainMessage<TaskServiceGetRequest> | undefined, b: TaskServiceGetRequest | PlainMessage<TaskServiceGetRequest> | undefined): boolean {
    return proto3.util.equals(TaskServiceGetRequest, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceGetResponse
 */
export class TaskServiceGetResponse extends Message<TaskServiceGetResponse> {
  /**
   * @generated from field: app.v1.Task task = 1;
   */
  task?: Task;

  constructor(data?: PartialMessage<TaskServiceGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceGetResponse {
    return new TaskServiceGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceGetResponse {
    return new TaskServiceGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceGetResponse {
    return new TaskServiceGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceGetResponse | PlainMessage<TaskServiceGetResponse> | undefined, b: TaskServiceGetResponse | PlainMessage<TaskServiceGetResponse> | undefined): boolean {
    return proto3.util.equals(TaskServiceGetResponse, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceListRequest
 */
export class TaskServiceListRequest extends Message<TaskServiceListRequest> {
  constructor(data?: PartialMessage<TaskServiceListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceListRequest {
    return new TaskServiceListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceListRequest {
    return new TaskServiceListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceListRequest {
    return new TaskServiceListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceListRequest | PlainMessage<TaskServiceListRequest> | undefined, b: TaskServiceListRequest | PlainMessage<TaskServiceListRequest> | undefined): boolean {
    return proto3.util.equals(TaskServiceListRequest, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceListResponse
 */
export class TaskServiceListResponse extends Message<TaskServiceListResponse> {
  /**
   * @generated from field: repeated app.v1.Task tasks = 1;
   */
  tasks: Task[] = [];

  constructor(data?: PartialMessage<TaskServiceListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tasks", kind: "message", T: Task, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceListResponse {
    return new TaskServiceListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceListResponse {
    return new TaskServiceListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceListResponse {
    return new TaskServiceListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceListResponse | PlainMessage<TaskServiceListResponse> | undefined, b: TaskServiceListResponse | PlainMessage<TaskServiceListResponse> | undefined): boolean {
    return proto3.util.equals(TaskServiceListResponse, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceUpdateRequest
 */
export class TaskServiceUpdateRequest extends Message<TaskServiceUpdateRequest> {
  /**
   * @generated from field: app.v1.TaskServiceUpdateRequest.Fields task = 1;
   */
  task?: TaskServiceUpdateRequest_Fields;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 2;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<TaskServiceUpdateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceUpdateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: TaskServiceUpdateRequest_Fields },
    { no: 2, name: "field_mask", kind: "message", T: FieldMask },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceUpdateRequest {
    return new TaskServiceUpdateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceUpdateRequest {
    return new TaskServiceUpdateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceUpdateRequest {
    return new TaskServiceUpdateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceUpdateRequest | PlainMessage<TaskServiceUpdateRequest> | undefined, b: TaskServiceUpdateRequest | PlainMessage<TaskServiceUpdateRequest> | undefined): boolean {
    return proto3.util.equals(TaskServiceUpdateRequest, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceUpdateRequest.Fields
 */
export class TaskServiceUpdateRequest_Fields extends Message<TaskServiceUpdateRequest_Fields> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: optional string title = 2;
   */
  title?: string;

  /**
   * @generated from field: optional app.v1.TaskStatus status = 3;
   */
  status?: TaskStatus;

  constructor(data?: PartialMessage<TaskServiceUpdateRequest_Fields>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceUpdateRequest.Fields";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(TaskStatus), opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceUpdateRequest_Fields {
    return new TaskServiceUpdateRequest_Fields().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceUpdateRequest_Fields {
    return new TaskServiceUpdateRequest_Fields().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceUpdateRequest_Fields {
    return new TaskServiceUpdateRequest_Fields().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceUpdateRequest_Fields | PlainMessage<TaskServiceUpdateRequest_Fields> | undefined, b: TaskServiceUpdateRequest_Fields | PlainMessage<TaskServiceUpdateRequest_Fields> | undefined): boolean {
    return proto3.util.equals(TaskServiceUpdateRequest_Fields, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceUpdateResponse
 */
export class TaskServiceUpdateResponse extends Message<TaskServiceUpdateResponse> {
  /**
   * @generated from field: app.v1.Task task = 1;
   */
  task?: Task;

  constructor(data?: PartialMessage<TaskServiceUpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceUpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task", kind: "message", T: Task },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceUpdateResponse {
    return new TaskServiceUpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceUpdateResponse {
    return new TaskServiceUpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceUpdateResponse {
    return new TaskServiceUpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceUpdateResponse | PlainMessage<TaskServiceUpdateResponse> | undefined, b: TaskServiceUpdateResponse | PlainMessage<TaskServiceUpdateResponse> | undefined): boolean {
    return proto3.util.equals(TaskServiceUpdateResponse, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceDeleteRequest
 */
export class TaskServiceDeleteRequest extends Message<TaskServiceDeleteRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<TaskServiceDeleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceDeleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceDeleteRequest {
    return new TaskServiceDeleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceDeleteRequest {
    return new TaskServiceDeleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceDeleteRequest {
    return new TaskServiceDeleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceDeleteRequest | PlainMessage<TaskServiceDeleteRequest> | undefined, b: TaskServiceDeleteRequest | PlainMessage<TaskServiceDeleteRequest> | undefined): boolean {
    return proto3.util.equals(TaskServiceDeleteRequest, a, b);
  }
}

/**
 * @generated from message app.v1.TaskServiceDeleteResponse
 */
export class TaskServiceDeleteResponse extends Message<TaskServiceDeleteResponse> {
  constructor(data?: PartialMessage<TaskServiceDeleteResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "app.v1.TaskServiceDeleteResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskServiceDeleteResponse {
    return new TaskServiceDeleteResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskServiceDeleteResponse {
    return new TaskServiceDeleteResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskServiceDeleteResponse {
    return new TaskServiceDeleteResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TaskServiceDeleteResponse | PlainMessage<TaskServiceDeleteResponse> | undefined, b: TaskServiceDeleteResponse | PlainMessage<TaskServiceDeleteResponse> | undefined): boolean {
    return proto3.util.equals(TaskServiceDeleteResponse, a, b);
  }
}

