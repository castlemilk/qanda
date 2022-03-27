// package: questions.v1alpha1
// file: question/v1alpha1/question.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as api_v1alpha1_question_pb from "../../api/v1alpha1/question_pb";

export class Question extends jspb.Message { 

    hasMetadata(): boolean;
    clearMetadata(): void;
    getMetadata(): api_v1alpha1_question_pb.Metadata | undefined;
    setMetadata(value?: api_v1alpha1_question_pb.Metadata): Question;

    hasAuthor(): boolean;
    clearAuthor(): void;
    getAuthor(): api_v1alpha1_question_pb.Author | undefined;
    setAuthor(value?: api_v1alpha1_question_pb.Author): Question;

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
        metadata?: api_v1alpha1_question_pb.Metadata.AsObject,
        author?: api_v1alpha1_question_pb.Author.AsObject,
        spec?: QuestionSpec.AsObject,
    }
}

export class QuestionSpec extends jspb.Message { 
    getContent(): string;
    setContent(value: string): QuestionSpec;
    clearTagsList(): void;
    getTagsList(): Array<api_v1alpha1_question_pb.Tag>;
    setTagsList(value: Array<api_v1alpha1_question_pb.Tag>): QuestionSpec;
    addTags(value?: api_v1alpha1_question_pb.Tag, index?: number): api_v1alpha1_question_pb.Tag;

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
        content: string,
        tagsList: Array<api_v1alpha1_question_pb.Tag.AsObject>,
    }
}
