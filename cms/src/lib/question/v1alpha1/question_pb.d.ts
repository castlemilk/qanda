// package: questions.v1alpha1
// file: question/v1alpha1/question.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as api_v1alpha1_metadata_pb from "../../api/v1alpha1/metadata_pb";
import * as api_v1alpha1_tag_pb from "../../api/v1alpha1/tag_pb";
import * as api_v1alpha1_content_pb from "../../api/v1alpha1/content_pb";
import * as author_v1alpha1_author_pb from "../../author/v1alpha1/author_pb";

export class Question extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): api_v1alpha1_metadata_pb.Metadata | undefined;
    setMetadata(value?: api_v1alpha1_metadata_pb.Metadata): Question;

    hasAuthor(): boolean;
    clearAuthor(): void;
    getAuthor(): author_v1alpha1_author_pb.Author | undefined;
    setAuthor(value?: author_v1alpha1_author_pb.Author): Question;

    hasSpec(): boolean;
    clearSpec(): void;
    getSpec(): QuestionSpec | undefined;
    setSpec(value?: QuestionSpec): Question;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Question.AsObject;
    static toObject(includeInstance: boolean, msg: Question): Question.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Question, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Question;
    static deserializeBinaryFromReader(message: Question, reader: jspb.BinaryReader): Question;
}

export namespace Question {
    export type AsObject = {
        metadata?: api_v1alpha1_metadata_pb.Metadata.AsObject,
        author?: author_v1alpha1_author_pb.Author.AsObject,
        spec?: QuestionSpec.AsObject,
    }
}

export class QuestionSpec extends jspb.Message { 

    hasContent(): boolean;
    clearContent(): void;
    getContent(): api_v1alpha1_content_pb.Content | undefined;
    setContent(value?: api_v1alpha1_content_pb.Content): QuestionSpec;
    clearTagsList(): void;
    getTagsList(): Array<api_v1alpha1_tag_pb.Tag>;
    setTagsList(value: Array<api_v1alpha1_tag_pb.Tag>): QuestionSpec;
    addTags(value?: api_v1alpha1_tag_pb.Tag, index?: number): api_v1alpha1_tag_pb.Tag;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QuestionSpec.AsObject;
    static toObject(includeInstance: boolean, msg: QuestionSpec): QuestionSpec.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QuestionSpec, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QuestionSpec;
    static deserializeBinaryFromReader(message: QuestionSpec, reader: jspb.BinaryReader): QuestionSpec;
}

export namespace QuestionSpec {
    export type AsObject = {
        content?: api_v1alpha1_content_pb.Content.AsObject,
        tagsList: Array<api_v1alpha1_tag_pb.Tag.AsObject>,
    }
}
