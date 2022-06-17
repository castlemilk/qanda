// package: api.v1alpha1
// file: api/v1alpha1/content.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as validate_validate_pb from "../../validate/validate_pb";

export class Content extends jspb.Message { 

    hasApiversion(): boolean;
    clearApiversion(): void;
    getApiversion(): string | undefined;
    setApiversion(value: string): Content;
    getCreatedat(): number;
    setCreatedat(value: number): Content;
    getUpdatedat(): number;
    setUpdatedat(value: number): Content;
    getDeletedat(): number;
    setDeletedat(value: number): Content;

    hasS3(): boolean;
    clearS3(): void;
    getS3(): S3 | undefined;
    setS3(value?: S3): Content;

    hasText(): boolean;
    clearText(): void;
    getText(): Text | undefined;
    setText(value?: Text): Content;

    getContentCase(): Content.ContentCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Content.AsObject;
    static toObject(includeInstance: boolean, msg: Content): Content.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Content, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Content;
    static deserializeBinaryFromReader(message: Content, reader: jspb.BinaryReader): Content;
}

export namespace Content {
    export type AsObject = {
        apiversion?: string,
        createdat: number,
        updatedat: number,
        deletedat: number,
        s3?: S3.AsObject,
        text?: Text.AsObject,
    }

    export enum Format {
    PLAINTEXT = 0,
    MARKDOWN = 1,
    }


    export enum ContentCase {
        CONTENT_NOT_SET = 0,
        S3 = 10,
        TEXT = 11,
    }

}

export class S3 extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): S3;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): S3.AsObject;
    static toObject(includeInstance: boolean, msg: S3): S3.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: S3, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): S3;
    static deserializeBinaryFromReader(message: S3, reader: jspb.BinaryReader): S3;
}

export namespace S3 {
    export type AsObject = {
        url: string,
    }
}

export class Text extends jspb.Message { 
    getValue(): string;
    setValue(value: string): Text;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Text.AsObject;
    static toObject(includeInstance: boolean, msg: Text): Text.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Text, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Text;
    static deserializeBinaryFromReader(message: Text, reader: jspb.BinaryReader): Text;
}

export namespace Text {
    export type AsObject = {
        value: string,
    }
}
