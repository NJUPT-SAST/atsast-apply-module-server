package entity

type Problem struct {
	ProblemId   *string `json:"problemId" bson:"problemId,omitempty" binding:"required"`
	ProblemName *string `json:"problemName" bson:"problemName,omitempty" binding:"required"`
	MaxScore    *int    `json:"maxScore" bson:"maxScore,omitempty" binding:"required"`
}

type Exam struct {
	ExamId      *string   `json:"examId" bson:"examId,omitempty" binding:"required"`
	ExamName    *string   `json:"examName" bson:"examName,omitempty" binding:"required"`
	ProblemList []Problem `json:"problemList" bson:"problemList,omitempty" binding:"required"`
}
