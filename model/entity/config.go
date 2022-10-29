package entity

type Config struct {
	Sast   SastConfig   `json:"sast" bson:"sast,omitempty" binding:"required"`
	School SchoolConfig `json:"school" bson:"school,omitempty" binding:"required"`
}

type SastConfig struct {
	DepartmentList []DepartmentConfig `json:"departmentList" bson:"departmentList,omitempty" binding:"required"`
	JobTitleList   []JobTitleConfig   `json:"jobTitleList" bson:"jobTitleList,omitempty" binding:"required"`
}

type DepartmentConfig struct {
	DepartmentId       string `json:"departmentId" bson:"departmentId,omitempty" binding:"required"`
	DepartmentName     string `json:"departmentName" bson:"departmentName,omitempty" binding:"required"`
	DepartmentCategory string `json:"departmentCategory" bson:"departmentCategory,omitempty" binding:"required"`
	CanApply           bool   `json:"canApply" bson:"canApply,omitempty" binding:"required"`
}

type JobTitleConfig struct {
	JobTitleId   *string `json:"jobTitleId" bson:"jobTitleId,omitempty" binding:"required"`
	JobTitleName *string `json:"jobTitleName" bson:"jobTitleName,omitempty" binding:"required"`
}

type SchoolConfig struct {
	CollegeList []College `json:"collegeList" bson:"collegeList,omitempty" binding:"required"`
}

type College struct {
	CollegeId   string        `json:"collegeId" bson:"collegeId,omitempty" binding:"required"`
	CollegeName string        `json:"collegeName" bson:"collegeName,omitempty" binding:"required"`
	MajorList   []MajorConfig `json:"majorList" bson:"majorList,omitempty" binding:"required"`
}

type MajorConfig struct {
	MajorId   string   `json:"majorId" bson:"majorId,omitempty" binding:"required"`
	MajorName string   `json:"majorName" bson:"majorName,omitempty" binding:"required"`
	ClassList []string `json:"classList" bson:"classList,omitempty" binding:"required"`
}
