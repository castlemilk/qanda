package question

import (
	"bytes"
	"encoding/json"
	"fmt"

	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/castlemilk/qanda/server/pkg/log"
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type firestoreStore struct {
	client     *firestore.Client
	collection string
}

func newFireStoreStorage() (Store, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Firestore(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error getting firestore client: %v", err)
	}
	return firestoreStore{client: client, collection: "qanda"}, nil
}

func (s firestoreStore) List() ([]*questionsv1alpha1.Question, error) {

	var questions []*questionsv1alpha1.Question
	log.Infof("list:firebase:")
	iter := s.client.Collection(s.collection).Documents(context.TODO())
	for {
		q := &questionsv1alpha1.Question{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return questions, err
		}
		log.Infof("found doc: %+v", doc.Data())

		if err != nil {
			log.Errorf("erroring unmarshalling object: %+v", err)
			return questions, err
		}
		b, err := json.Marshal(doc.Data())

		if err != nil {
			return questions, err
		}
		jsonpb.Unmarshal(bytes.NewBuffer(b), q)
		questions = append(questions, q)
	}

	return questions, nil

}

func (s firestoreStore) Get(id string) (*questionsv1alpha1.Question, error) {

	// we need to unmarshalva jsonpb first

	dsnap, err := s.client.Collection(s.collection).Doc(id).Get(context.TODO())

	if status.Code(err) == codes.NotFound {
		return &questionsv1alpha1.Question{}, fmt.Errorf("no question question found with id: %s", id)
	}
	if err != nil {
		return &questionsv1alpha1.Question{}, err
	}
	log.Infof("found item: %+v", dsnap.Data())
	q, err := getQuestionFromSnapshot(dsnap.Data())
	if err != nil {
		log.Errorf("erroring unmarshalling object: %+v", err)
		return q, err
	}

	return q, nil

}

func (s firestoreStore) Create(question questionsv1alpha1.Question) (*questionsv1alpha1.Question, error) {

	questionj, err := questionToMap(&question)
	_, err = s.client.Collection(s.collection).Doc(question.Metadata.Id).Set(context.TODO(), questionj)
	if err != nil {
		return &questionsv1alpha1.Question{}, err
	}
	return &question, nil
}

func (s firestoreStore) Delete(id string) (string, error) {
	_, err := s.client.Collection(s.collection).Doc(id).Delete(context.TODO())
	if status.Code(err) == codes.NotFound {
		return id, fmt.Errorf("no question question found with id: %s", id)
	}
	if err != nil {
		return id, err
	}
	if err != nil {
		log.Errorf("erroring unmarshalling object: %+v", err)
		return id, err
	}
	return id, nil
}

// getQuestionFromSnapshot converts firebase data to protobuf generated struct for a question to work around issue captured
// in https://github.com/googleapis/google-cloud-go/issues/1438
func getQuestionFromSnapshot(firebaseData map[string]interface{}) (*questionsv1alpha1.Question, error) {
	q := &questionsv1alpha1.Question{}
	data, err := json.Marshal(firebaseData)
	log.Infof("json: %+v", string(data))
	if err != nil {
		return q, err
	}
	// u := jsonpb.Unmarshaler{}
	protojson.Unmarshal(data, q)
	return q, nil
}

func questionToMap(q *questionsv1alpha1.Question) (map[string]interface{}, error) {
	var inInterface map[string]interface{}
	data, err := protojson.Marshal(q.ProtoReflect().Interface())
	if err != nil {
		log.Fatalf("Failed to JSON marhsal protobuf.\nError: %s", err.Error())
		return nil, err
	}

	json.Unmarshal(data, &inInterface)
	return inInterface, nil
}
