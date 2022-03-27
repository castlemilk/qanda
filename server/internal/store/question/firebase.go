package question

import (
	"fmt"

	"context"

	"cloud.google.com/go/firestore"
	"github.com/castlemilk/qanda/server/pkg/log"
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
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

func (s firestoreStore) Get(id string) (*questionsv1alpha1.Question, error) {
	q := &questionsv1alpha1.Question{}
	dsnap, err := s.client.Collection(s.collection).Doc(id).Get(context.TODO())
	if err != nil {
		return nil, err
	}
	if status.Code(err) == codes.NotFound {
		return q, fmt.Errorf("no question question found with id: %s", id)
	}
	log.Infof("found item: %+v", dsnap.Data())
	err = dsnap.DataTo(q)

	if err != nil {
		log.Errorf("erroring unmarshalling object: %+v", err)
		return q, err
	}

	return q, nil

}

func (s firestoreStore) Create(question questionsv1alpha1.Question) (*questionsv1alpha1.Question, error) {
	// s.questions[question.Metadata.Id] = question
	return &question, nil
}

func (s firestoreStore) Delete(id string) (*questionsv1alpha1.Question, error) {
	question, err := s.Get(id)
	if err != nil {
		return &questionsv1alpha1.Question{}, err
	}
	// delete(s.questions, id)
	return question, nil
}
