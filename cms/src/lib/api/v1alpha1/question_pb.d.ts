// package: api.v1alpha1
// file: api/v1alpha1/question.proto

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

export class Answer extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): Metadata | undefined;
    setMetadata(value?: Metadata): Answer;

    hasAuthor(): boolean;
    clearAuthor(): void;
    getAuthor(): Author | undefined;
    setAuthor(value?: Author): Answer;
    getContents(): string;
    setContents(value: string): Answer;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Answer.AsObject;
    static toObject(includeInstance: boolean, msg: Answer): Answer.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Answer, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Answer;
    static deserializeBinaryFromReader(message: Answer, reader: jspb.BinaryReader): Answer;
}

export namespace Answer {
    export type AsObject = {
        metadata?: Metadata.AsObject,
        author?: Author.AsObject,
        contents: string,
    }
}

export class Author extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): Metadata | undefined;
    setMetadata(value?: Metadata): Author;
    getName(): string;
    setName(value: string): Author;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Author.AsObject;
    static toObject(includeInstance: boolean, msg: Author): Author.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Author, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Author;
    static deserializeBinaryFromReader(message: Author, reader: jspb.BinaryReader): Author;
}

export namespace Author {
    export type AsObject = {
        metadata?: Metadata.AsObject,
        name: string,
    }
}

export class Tag extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): Metadata | undefined;
    setMetadata(value?: Metadata): Tag;
    getValue(): string;
    setValue(value: string): Tag;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Tag.AsObject;
    static toObject(includeInstance: boolean, msg: Tag): Tag.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Tag, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Tag;
    static deserializeBinaryFromReader(message: Tag, reader: jspb.BinaryReader): Tag;
}

export namespace Tag {
    export type AsObject = {
        metadata?: Metadata.AsObject,
        value: string,
    }
}
