// package: api.v1alpha1
// file: api/v1alpha1/tag.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as api_v1alpha1_metadata_pb from "../../api/v1alpha1/metadata_pb";

export class Tag extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): api_v1alpha1_metadata_pb.Metadata | undefined;
    setMetadata(value?: api_v1alpha1_metadata_pb.Metadata): Tag;
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
        metadata?: api_v1alpha1_metadata_pb.Metadata.AsObject,
        value: string,
    }
}
