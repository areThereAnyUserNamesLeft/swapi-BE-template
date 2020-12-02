package structs

type Planets struct {
	Schema      string `json:"$schema" bson:"$schema"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Properties  struct {
		Name struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"name" bson:"name"`
		Diameter struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"diameter" bson:"diameter"`
		RotationPeriod struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"rotation_period" bson:"rotation_period"`
		OrbitalPeriod struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"orbital_period" bson:"orbital_period"`
		Gravity struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"gravity" bson:"gravity"`
		Population struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"population" bson:"population"`
		Climate struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"climate" bson:"climate"`
		Terrain struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"terrain" bson:"terrain"`
		SurfaceWater struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"surface_water" bson:"surface_water"`
		Films struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"films" bson:"films"`
		Residents struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"residents" bson:"residents"`
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
