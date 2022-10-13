package question

import (
	"bytes"
	"encoding/json"
	"fmt"

	"context"

	"cloud.google.com/go/firestore"
	"github.com/castlemilk/qanda/server/pkg/log"
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	firebase "firebase.google.com/go"
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

	q := &questionsv1alpha1.Question{}
	dsnap, err := s.client.Collection(s.collection).Doc(id).Get(context.TODO())
	
	if status.Code(err) == codes.NotFound {
		return q, fmt.Errorf("no question question found with id: %s", id)
	}
	if err != nil {
		return  &questionsv1alpha1.Question{}, err
	}
	log.Infof("found item: %+v", dsnap.Data())
	data, err := json.Marshal(dsnap.Data())
	u := jsonpb.Unmarshaler{}
	u.Unmarshal(bytes.NewReader(data), q)

	if err != nil {
		log.Errorf("erroring unmarshalling object: %+v", err)
		return q, err
	}

	return q, nil

}

func (s firestoreStore) Create(question questionsv1alpha1.Question) (*questionsv1alpha1.Question, error) {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(question)
	json.Unmarshal(inrec, &inInterface)
	log.Infof("question: %+v", question)
	log.Infof("writing doc: %+v", inInterface)
	_, err := s.client.Collection(s.collection).Doc(question.Metadata.Id).Set(context.TODO(), inInterface)
	if err != nil {
		return  &questionsv1alpha1.Question{}, err
	}
	return &question, nil
}

func (s firestoreStore) Delete(id string) (string, error) {
	_, err := s.client.Collection(s.collection).Doc(id).Delete(context.TODO())
	if status.Code(err) == codes.NotFound {
		return id, fmt.Errorf("no question question found with id: %s", id)
	}
	if err != nil {
		return  id, err
	}
	if err != nil {
		log.Errorf("erroring unmarshalling object: %+v", err)
		return id, err
	}
	return id, nil
}

// func convertToFirebaseDoc(question questionsv1alpha1.Question) (*questionsv1alpha1.Question) {
// 	return &questionsv1alpha1.Question{
// 		Metadata: &apiv1alpha1.Metadata{

// 		},
// 	}
// }
