package config

type Skill struct {
	Skill_id          string
	Skill_name        string
	Skill_function    string
	Skill_describe    string
	Skill_type        int
	Skill_parameter_a string
	Skill_parameter_b string
	Skill_parameter_c string
	Skill_parameter_d string
	Skill_parameter_e string
	Skill_parameter_f string
}

var ConfSkills []Skill
