package loan

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"loanapp/business"
	loanBusiness "loanapp/business/loan"
)

type document struct {
	ID        primitive.ObjectID `bson:"_id"`
	Amount    int                `bson:"amount"`
	State     string             `bson:"state"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	Version   int                `bson:"version"`
}

func newDocument(loan loanBusiness.Loan) (*document, error) {

	idString := strconv.Itoa(loan.ID)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, err
	}

	return &document{
		ID:        id,
		Amount:    loan.Amount,
		State:     string(loan.State),
		CreatedAt: loan.CreatedAt,
		UpdatedAt: loan.UpdatedAt,
		Version:   loan.Version,
	}, nil
}

func (doc *document) ToLoan() loanBusiness.Loan {
	var l loanBusiness.Loan

	idInt, _ := strconv.Atoi(doc.ID.Hex())
	l.ID = idInt
	l.Amount = doc.Amount
	l.State = loanBusiness.State(doc.State)
	l.CreatedAt = doc.CreatedAt
	l.UpdatedAt = doc.UpdatedAt
	l.Version = doc.Version

	return l
}

type MongoDBRepository struct {
	col *mongo.Collection
}

func NewMongoDBRepository(db *mongo.Database) (*MongoDBRepository, error) {
	repo := MongoDBRepository{
		db.Collection("loan"),
	}

	return &repo, nil
}

// Insert
func (repo *MongoDBRepository) InsertLoan(loan loanBusiness.Loan) error {
	col, err := newDocument(loan)
	if err != nil {
		return err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)
	if err != nil {
		merr := err.(mongo.WriteException)
		errCode := merr.WriteErrors[0].Code
		if errCode == 11000 {
			return business.ErrDuplicate
		} else if err == mongo.ErrNoDocuments {
			return business.ErrNotFound
		}

		return err
	}

	return nil
}

// FindLoanByID
func (repo *MongoDBRepository) FindLoanByID(ID int) (*loanBusiness.Loan, error) {
	var doc document

	idString := strconv.Itoa(ID)

	objID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objID,
	}

	if err := repo.col.FindOne(context.TODO(), filter).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, business.ErrNotFound
		}

		return nil, err
	}

	loan := doc.ToLoan()

	return &loan, nil
}

// FindAll
func (repo *MongoDBRepository) FindAllLoan() ([]loanBusiness.Loan, error) {
	loans := []loanBusiness.Loan{}

	option := options.Find()
	option.SetSort(bson.D{{Key: "_id", Value: 1}})

	cursor, err := repo.col.Find(context.TODO(), bson.M{}, option)
	if err != nil {
		return loans, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var doc document

		err := cursor.Decode(&doc)
		if err != nil {
			return loans, err
		}

		data := doc.ToLoan()

		loans = append(loans, data)
	}

	return loans, nil
}

// Update
func (repo *MongoDBRepository) UpdateLoan(id int, amount int, currentVersion int) (bool, error) {

	currentTime := time.Now()

	idString := strconv.Itoa(id)

	objID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return false, err
	}

	filter := bson.M{
		"_id":     objID,
		"version": currentVersion,
	}

	updated := bson.M{
		"$set": bson.M{
			"amount":    amount,
			"updatedAt": currentTime,
			"version":   currentVersion + 1,
		},
	}

	updateResult, err := repo.col.UpdateOne(context.TODO(), filter, updated)
	if err != nil {
		return false, err
	}

	isModified := updateResult.ModifiedCount == 1

	return isModified, nil
}

// Approval
func (repo *MongoDBRepository) ApprovalLoanState(id int, state string, currentVersion int) (bool, error) {

	currentTime := time.Now()

	idString := strconv.Itoa(id)

	objID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return false, err
	}

	filter := bson.M{
		"_id":     objID,
		"version": currentVersion,
	}

	updated := bson.M{
		"$set": bson.M{
			"state":     state,
			"updatedAt": currentTime,
			"version":   currentVersion + 1,
		},
	}

	updateResult, err := repo.col.UpdateOne(context.TODO(), filter, updated)
	if err != nil {
		return false, err
	}

	isModified := updateResult.ModifiedCount == 1

	return isModified, nil
}
