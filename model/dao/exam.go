package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/njupt-sast/atsast-apply-module-server/model/entity"
)

func ReadExamList() ([]entity.Exam, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeLimit)
	defer cancel()

	cur, err := ExamColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	examList := make([]entity.Exam, 0)
	err = cur.All(ctx, &examList)
	if err != nil {
		return nil, err
	}
	return examList, nil
}
