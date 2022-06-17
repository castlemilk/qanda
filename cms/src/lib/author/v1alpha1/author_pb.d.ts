// package: author.v1alpha1
// file: author/v1alpha1/author.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as api_v1alpha1_metadata_pb from "../../api/v1alpha1/metadata_pb";

export class Author extends jspb.Message { 
    getId(): string;
    setId(value: string): Author;

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
        id: string,
    }
}

export class AuthorDetails extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): api_v1alpha1_metadata_pb.Metadata | undefined;
    setMetadata(value?: api_v1alpha1_metadata_pb.Metadata): AuthorDetails;
    getName(): string;
    setName(value: string): AuthorDetails;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthorDetails.AsObject;
    static toObject(includeInstance: boolean, msg: AuthorDetails): AuthorDetails.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthorDetails, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthorDetails;
    static deserializeBinaryFromReader(message: AuthorDetails, reader: jspb.BinaryReader): AuthorDetails;
}

export namespace AuthorDetails {
    export type AsObject = {
        metadata?: api_v1alpha1_metadata_pb.Metadata.AsObject,
        name: string,
    }
}
