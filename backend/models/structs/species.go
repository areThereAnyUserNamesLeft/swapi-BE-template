package structs

type Species struct {
	Schema      string `json:"$schema" bson:"$schema"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Properties  struct {
		Name struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"name" bson:"name"`
		Classification struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"classification" bson:"classification"`
		Designation struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"designation" bson:"designation"`
		AverageHeight struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"average_height" bson:"average_height"`
		AverageLifespan struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"average_lifespan" bson:"average_lifespan"`
		HairColors struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"hair_colors" bson:"hair_colors"`
		SkinColors struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"skin_colors" bson:"skin_colors"`
		EyeColors struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"eye_colors" bson:"eye_colors"`
		Homeworld struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"homeworld" bson:"homeworld"`
		Language struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"language" bson:"language"`
		People struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"people" bson:"people"`
		Films struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"films" bson:"films"`
		URL struct {
			Type        string `json:"type" bson:"type"`
			Format      string `json:"format" bson:"format"`
			Description string `json:"description" bson:"description"`
		} `json:"url" bson:"url"`
		Created struct {
			Type        string `json:"type" bson:"type"`
			Format      string `json:"format" bson:"format"`
			Description string `json:"description" bson:"description"`
		} `json:"created" bson:"created"`
		Edited struct {
			Type        string `json:"type" bson:"type"`
			Format      string `json:"format" bson:"format"`
			Description string `json:"description" bson:"description"`
		} `json:"edited" bson:"edited"`
	} `json:"properties" bson:"properties"`
	Required []string `json:"required" bson:"required"`
}
