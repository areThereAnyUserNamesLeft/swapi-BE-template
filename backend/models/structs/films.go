package structs

type Films struct {
	Schema      string `json:"$schema" bson:"$schema"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        string `json:"type" bson:"type"`
	Properties  struct {
		Title struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"title" bson:"title"`
		EpisodeID struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"episode_id" bson:"episode_id"`
		OpeningCrawl struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"opening_crawl" bson:"opening_crawl"`
		Director struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"director" bson:"director"`
		Producer struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"producer" bson:"producer"`
		ReleaseDate struct {
			Type        string `json:"type" bson:"type"`
			Format      string `json:"format" bson:"format"`
			Description string `json:"description" bson:"description"`
		} `json:"release_date" bson:"release_date"`
		Characters struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"characters" bson:"characters"`
		Planets struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"planets" bson:"planets"`
		Starships struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"starships" bson:"starships"`
		Vehicles struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"vehicles" bson:"vehicles"`
		Species struct {
			Type        string `json:"type" bson:"type"`
			Description string `json:"description" bson:"description"`
		} `json:"species" bson:"species"`
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
