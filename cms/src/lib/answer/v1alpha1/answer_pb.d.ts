// package: answers.v1alpha1
// file: answer/v1alpha1/answer.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as api_v1alpha1_metadata_pb from "../../api/v1alpha1/metadata_pb";
import * as api_v1alpha1_tag_pb from "../../api/v1alpha1/tag_pb";
import * as api_v1alpha1_content_pb from "../../api/v1alpha1/content_pb";
import * as author_v1alpha1_author_pb from "../../author/v1alpha1/author_pb";

export class Answer extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): api_v1alpha1_metadata_pb.Metadata | undefined;
    setMetadata(value?: api_v1alpha1_metadata_pb.Metadata): Answer;

    hasAuthor(): boolean;
    clearAuthor(): void;
    getAuthor(): author_v1alpha1_author_pb.Author | undefined;
    setAuthor(value?: author_v1alpha1_author_pb.Author): Answer;

    hasSpec(): boolean;
    clearSpec(): void;
    getSpec(): AnswerSpec | undefined;
    setSpec(value?: AnswerSpec): Answer;

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
        metadata?: api_v1alpha1_metadata_pb.Metadata.AsObject,
        author?: author_v1alpha1_author_pb.Author.AsObject,
        spec?: AnswerSpec.AsObject,
    }
}

export class AnswerSpec extends jspb.Message { 

    hasContent(): boolean;
    clearContent(): void;
    getContent(): api_v1alpha1_content_pb.Content | undefined;
    setContent(value?: api_v1alpha1_content_pb.Content): AnswerSpec;
    clearTagsList(): void;
    getTagsList(): Array<api_v1alpha1_tag_pb.Tag>;
    setTagsList(value: Array<api_v1alpha1_tag_pb.Tag>): AnswerSpec;
    addTags(value?: api_v1alpha1_tag_pb.Tag, index?: number): api_v1alpha1_tag_pb.Tag;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AnswerSpec.AsObject;
    static toObject(includeInstance: boolean, msg: AnswerSpec): AnswerSpec.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AnswerSpec, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AnswerSpec;
    static deserializeBinaryFromReader(message: AnswerSpec, reader: jspb.BinaryReader): AnswerSpec;
}

export namespace AnswerSpec {
    export type AsObject = {
        content?: api_v1alpha1_content_pb.Content.AsObject,
        tagsList: Array<api_v1alpha1_tag_pb.Tag.AsObject>,
    }
}
