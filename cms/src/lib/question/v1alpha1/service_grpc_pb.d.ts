// package: questions.v1alpha1
// file: question/v1alpha1/service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as question_v1alpha1_service_pb from "../../question/v1alpha1/service_pb";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../protoc-gen-openapiv2/options/annotations_pb";
import * as question_v1alpha1_question_pb from "../../question/v1alpha1/question_pb";

interface IQuestionServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listQuestions: IQuestionServiceService_IListQuestions;
    getQuestion: IQuestionServiceService_IGetQuestion;
    createQuestion: IQuestionServiceService_ICreateQuestion;
    deleteQuestion: IQuestionServiceService_IDeleteQuestion;
    updateQuestion: IQuestionServiceService_IUpdateQuestion;
}

interface IQuestionServiceService_IListQuestions extends grpc.MethodDefinition<question_v1alpha1_service_pb.ListQuestionsRequest, question_v1alpha1_service_pb.ListQuestionsResponse> {
    path: "/questions.v1alpha1.QuestionService/ListQuestions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<question_v1alpha1_service_pb.ListQuestionsRequest>;
    requestDeserialize: grpc.deserialize<question_v1alpha1_service_pb.ListQuestionsRequest>;
    responseSerialize: grpc.serialize<question_v1alpha1_service_pb.ListQuestionsResponse>;
    responseDeserialize: grpc.deserialize<question_v1alpha1_service_pb.ListQuestionsResponse>;
}
interface IQuestionServiceService_IGetQuestion extends grpc.MethodDefinition<question_v1alpha1_service_pb.GetQuestionRequest, question_v1alpha1_service_pb.GetQuestionResponse> {
    path: "/questions.v1alpha1.QuestionService/GetQuestion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<question_v1alpha1_service_pb.GetQuestionRequest>;
    requestDeserialize: grpc.deserialize<question_v1alpha1_service_pb.GetQuestionRequest>;
    responseSerialize: grpc.serialize<question_v1alpha1_service_pb.GetQuestionResponse>;
    responseDeserialize: grpc.deserialize<question_v1alpha1_service_pb.GetQuestionResponse>;
}
interface IQuestionServiceService_ICreateQuestion extends grpc.MethodDefinition<question_v1alpha1_service_pb.CreateQuestionRequest, question_v1alpha1_service_pb.CreateQuestionResponse> {
    path: "/questions.v1alpha1.QuestionService/CreateQuestion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<question_v1alpha1_service_pb.CreateQuestionRequest>;
    requestDeserialize: grpc.deserialize<question_v1alpha1_service_pb.CreateQuestionRequest>;
    responseSerialize: grpc.serialize<question_v1alpha1_service_pb.CreateQuestionResponse>;
    responseDeserialize: grpc.deserialize<question_v1alpha1_service_pb.CreateQuestionResponse>;
}
interface IQuestionServiceService_IDeleteQuestion extends grpc.MethodDefinition<question_v1alpha1_service_pb.DeleteQuestionRequest, question_v1alpha1_service_pb.DeleteQuestionResponse> {
    path: "/questions.v1alpha1.QuestionService/DeleteQuestion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<question_v1alpha1_service_pb.DeleteQuestionRequest>;
    requestDeserialize: grpc.deserialize<question_v1alpha1_service_pb.DeleteQuestionRequest>;
    responseSerialize: grpc.serialize<question_v1alpha1_service_pb.DeleteQuestionResponse>;
    responseDeserialize: grpc.deserialize<question_v1alpha1_service_pb.DeleteQuestionResponse>;
}
interface IQuestionServiceService_IUpdateQuestion extends grpc.MethodDefinition<question_v1alpha1_service_pb.UpdateQuestionRequest, question_v1alpha1_service_pb.UpdateQuestionResponse> {
    path: "/questions.v1alpha1.QuestionService/UpdateQuestion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<question_v1alpha1_service_pb.UpdateQuestionRequest>;
    requestDeserialize: grpc.deserialize<question_v1alpha1_service_pb.UpdateQuestionRequest>;
    responseSerialize: grpc.serialize<question_v1alpha1_service_pb.UpdateQuestionResponse>;
    responseDeserialize: grpc.deserialize<question_v1alpha1_service_pb.UpdateQuestionResponse>;
}

export const QuestionServiceService: IQuestionServiceService;

export interface IQuestionServiceServer extends grpc.UntypedServiceImplementation {
    listQuestions: grpc.handleUnaryCall<question_v1alpha1_service_pb.ListQuestionsRequest, question_v1alpha1_service_pb.ListQuestionsResponse>;
    getQuestion: grpc.handleUnaryCall<question_v1alpha1_service_pb.GetQuestionRequest, question_v1alpha1_service_pb.GetQuestionResponse>;
    createQuestion: grpc.handleUnaryCall<question_v1alpha1_service_pb.CreateQuestionRequest, question_v1alpha1_service_pb.CreateQuestionResponse>;
    deleteQuestion: grpc.handleUnaryCall<question_v1alpha1_service_pb.DeleteQuestionRequest, question_v1alpha1_service_pb.DeleteQuestionResponse>;
    updateQuestion: grpc.handleUnaryCall<question_v1alpha1_service_pb.UpdateQuestionRequest, question_v1alpha1_service_pb.UpdateQuestionResponse>;
}

export interface IQuestionServiceClient {
    listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
    updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
    updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
}

export class QuestionServiceClient extends grpc.Client implements IQuestionServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    public listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    public listQuestions(request: question_v1alpha1_service_pb.ListQuestionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.ListQuestionsResponse) => void): grpc.ClientUnaryCall;
    public getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    public getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    public getQuestion(request: question_v1alpha1_service_pb.GetQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.GetQuestionResponse) => void): grpc.ClientUnaryCall;
    public createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    public createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    public createQuestion(request: question_v1alpha1_service_pb.CreateQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.CreateQuestionResponse) => void): grpc.ClientUnaryCall;
    public deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    public deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    public deleteQuestion(request: question_v1alpha1_service_pb.DeleteQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.DeleteQuestionResponse) => void): grpc.ClientUnaryCall;
    public updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
    public updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
    public updateQuestion(request: question_v1alpha1_service_pb.UpdateQuestionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: question_v1alpha1_service_pb.UpdateQuestionResponse) => void): grpc.ClientUnaryCall;
}
