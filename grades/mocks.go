package grades

func init() {
	students = Students{
		{
			ID:        1,
			FirstName: "Dovud",
			LastName:  "Inomov",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Week 1 homework",
					Type:  GradeHomework,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Iana",
			LastName:  "Chernysheva",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 79,
				},
				{
					Title: "Week 1 homework",
					Type:  GradeHomework,
					Score: 92,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 96,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Sulaimon",
			LastName:  "Inomov",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 74,
				},
				{
					Title: "Week 1 homework",
					Type:  GradeHomework,
					Score: 98,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 95,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Asror",
			LastName:  "Valiev",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 67,
				},
				{
					Title: "Week 1 homework",
					Type:  GradeHomework,
					Score: 81,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 80,
				},
			},
		},
	}

}
