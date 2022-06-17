// package: api.v1alpha1
// file: api/v1alpha1/metadata.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as validate_validate_pb from "../../validate/validate_pb";

export class Metadata extends jspb.Message { 

    hasApiversion(): boolean;
    clearApiversion(): void;
    getApiversion(): string | undefined;
    setApiversion(value: string): Metadata;
    getId(): string;
    setId(value: string): Metadata;
    getName(): string;
    setName(value: string): Metadata;
    getCreatedat(): number;
    setCreatedat(value: number): Metadata;
    getUpdatedat(): number;
    setUpdatedat(value: number): Metadata;
    getDeletedat(): number;
    setDeletedat(value: number): Metadata;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Metadata.AsObject;
    static toObject(includeInstance: boolean, msg: Metadata): Metadata.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Metadata, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Metadata;
    static deserializeBinaryFromReader(message: Metadata, reader: jspb.BinaryReader): Metadata;
}

export namespace Metadata {
    export type AsObject = {
        apiversion?: string,
        id: string,
        name: string,
        createdat: number,
        updatedat: number,
        deletedat: number,
    }
}
