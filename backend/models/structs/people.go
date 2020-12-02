package structs

type People struct {
	Schema      string `json:"$schema" bson:"$schema"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Properties  struct {
		Name struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"name" bson:"name"`
		Height struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"height" bson:"height"`
		Mass struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"mass" bson:"mass"`
		HairColor struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"hair_color" bson:"hair_color"`
		SkinColor struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"skin_color" bson:"skin_color"`
		EyeColor struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"eye_color" bson:"eye_color"`
		BirthYear struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"birth_year" bson:"birth_year"`
		Gender struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"gender" bson:"gender"`
		Homeworld struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"homeworld" bson:"homeworld"`
		Films struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"films" bson:"films"`
		Species struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"species" bson:"species"`
		Vehicles struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"vehicles" bson:"vehicles"`
		Starships struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"starships" bson:"starships"`
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
