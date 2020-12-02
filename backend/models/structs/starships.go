package structs

type Starships struct {
	Schema      string `json:"$schema" bson:"$schema"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Properties  struct {
		Name struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"name" bson:"name"`
		Model struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"model" bson:"model"`
		StarshipClass struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"starship_class" bson:"starship_class"`
		Manufacturer struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"manufacturer" bson:"manufacturer"`
		CostInCredits struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"cost_in_credits" bson:"cost_in_credits"`
		Length struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"length" bson:"length"`
		Crew struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"crew" bson:"crew"`
		Passengers struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"passengers" bson:"passengers"`
		MaxAtmospheringSpeed struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"max_atmosphering_speed" bson:"max_atmosphering_speed"`
		HyperdriveRating struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"hyperdrive_rating" bson:"hyperdrive_rating"`
		MGLT struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"MGLT" bson:"MGLT"`
		CargoCapacity struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"cargo_capacity" bson:"cargo_capacity"`
		Consumables struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"consumables" bson:"consumables"`
		Films struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"films" bson:"films"`
		Pilots struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"pilots" bson:"pilots"`
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
