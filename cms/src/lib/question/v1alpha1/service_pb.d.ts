// package: questions.v1alpha1
// file: question/v1alpha1/service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../protoc-gen-openapiv2/options/annotations_pb";
import * as question_v1alpha1_question_pb from "../../question/v1alpha1/question_pb";

export class ListQuestionsRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListQuestionsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListQuestionsRequest): ListQuestionsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListQuestionsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListQuestionsRequest;
    static deserializeBinaryFromReader(message: ListQuestionsRequest, reader: jspb.BinaryReader): ListQuestionsRequest;
}

export namespace ListQuestionsRequest {
    export type AsObject = {
    }
}

export class ListQuestionsResponse extends jspb.Message { 
    clearQuestionsList(): void;
    getQuestionsList(): Array<question_v1alpha1_question_pb.Question>;
    setQuestionsList(value: Array<question_v1alpha1_question_pb.Question>): ListQuestionsResponse;
    addQuestions(value?: question_v1alpha1_question_pb.Question, index?: number): question_v1alpha1_question_pb.Question;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListQuestionsResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListQuestionsResponse): ListQuestionsResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListQuestionsResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListQuestionsResponse;
    static deserializeBinaryFromReader(message: ListQuestionsResponse, reader: jspb.BinaryReader): ListQuestionsResponse;
}

export namespace ListQuestionsResponse {
    export type AsObject = {
        questionsList: Array<question_v1alpha1_question_pb.Question.AsObject>,
    }
}

export class GetQuestionRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetQuestionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetQuestionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetQuestionRequest): GetQuestionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetQuestionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetQuestionRequest;
    static deserializeBinaryFromReader(message: GetQuestionRequest, reader: jspb.BinaryReader): GetQuestionRequest;
}

export namespace GetQuestionRequest {
    export type AsObject = {
        id: string,
    }
}

export class GetQuestionResponse extends jspb.Message { 

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): GetQuestionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetQuestionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetQuestionResponse): GetQuestionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetQuestionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetQuestionResponse;
    static deserializeBinaryFromReader(message: GetQuestionResponse, reader: jspb.BinaryReader): GetQuestionResponse;
}

export namespace GetQuestionResponse {
    export type AsObject = {
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}

export class CreateQuestionRequest extends jspb.Message { 

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): CreateQuestionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateQuestionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateQuestionRequest): CreateQuestionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateQuestionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateQuestionRequest;
    static deserializeBinaryFromReader(message: CreateQuestionRequest, reader: jspb.BinaryReader): CreateQuestionRequest;
}

export namespace CreateQuestionRequest {
    export type AsObject = {
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}

export class CreateQuestionResponse extends jspb.Message { 

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): CreateQuestionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateQuestionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateQuestionResponse): CreateQuestionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateQuestionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateQuestionResponse;
    static deserializeBinaryFromReader(message: CreateQuestionResponse, reader: jspb.BinaryReader): CreateQuestionResponse;
}

export namespace CreateQuestionResponse {
    export type AsObject = {
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}

export class UpdateQuestionRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateQuestionRequest;

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): UpdateQuestionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateQuestionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateQuestionRequest): UpdateQuestionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateQuestionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateQuestionRequest;
    static deserializeBinaryFromReader(message: UpdateQuestionRequest, reader: jspb.BinaryReader): UpdateQuestionRequest;
}

export namespace UpdateQuestionRequest {
    export type AsObject = {
        id: string,
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}

export class UpdateQuestionResponse extends jspb.Message { 

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): UpdateQuestionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateQuestionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateQuestionResponse): UpdateQuestionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateQuestionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateQuestionResponse;
    static deserializeBinaryFromReader(message: UpdateQuestionResponse, reader: jspb.BinaryReader): UpdateQuestionResponse;
}

export namespace UpdateQuestionResponse {
    export type AsObject = {
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}

export class DeleteQuestionRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): DeleteQuestionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteQuestionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteQuestionRequest): DeleteQuestionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteQuestionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteQuestionRequest;
    static deserializeBinaryFromReader(message: DeleteQuestionRequest, reader: jspb.BinaryReader): DeleteQuestionRequest;
}

export namespace DeleteQuestionRequest {
    export type AsObject = {
        id: string,
    }
}

export class DeleteQuestionResponse extends jspb.Message { 

    hasQuestion(): boolean;
    clearQuestion(): void;
    getQuestion(): question_v1alpha1_question_pb.Question | undefined;
    setQuestion(value?: question_v1alpha1_question_pb.Question): DeleteQuestionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteQuestionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteQuestionResponse): DeleteQuestionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteQuestionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteQuestionResponse;
    static deserializeBinaryFromReader(message: DeleteQuestionResponse, reader: jspb.BinaryReader): DeleteQuestionResponse;
}

export namespace DeleteQuestionResponse {
    export type AsObject = {
        question?: question_v1alpha1_question_pb.Question.AsObject,
    }
}
