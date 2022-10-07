package entity

type Config struct {
	Sast   SastConfig   `json:"sast" bson:"sast" bind:"required"`
	School SchoolConfig `json:"school" bson:"school" bind:"required"`
}

type SastConfig struct {
	DepartmentList []DepartmentConfig `json:"departmentList" bson:"departmentList" bind:"required"`
}

type DepartmentConfig struct {
	DepartmentId       string `json:"departmentId" bson:"departmentId" bind:"required"`
	DepartmentName     string `json:"departmentName" bson:"departmentName" bind:"required"`
	DepartmentCategory string `json:"departmentCategory" bson:"departmentCategory" bind:"required"`
	CanApply           bool   `json:"canApply" bson:"canApply" bind:"required"`
}

type SchoolConfig struct {
	CollegeList []College `json:"collegeList" bson:"collegeList" bind:"required"`
}

type College struct {
	CollegeId   string        `json:"collegeId" bson:"collegeId" bind:"required"`
	CollegeName string        `json:"collegeName" bson:"collegeName" bind:"required"`
	MajorList   []MajorConfig `json:"majorList" bson:"majorList" bind:"required"`
}

type MajorConfig struct {
	MajorId   string   `json:"majorId" bson:"majorId" bind:"required"`
	MajorName string   `json:"majorName" bson:"majorName" bind:"required"`
	ClassList []string `json:"classList" bson:"classList" bind:"required"`
}
