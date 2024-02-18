package forest

// these are all for the forest api
type forest_user struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Created  string `json:"created_at"`
}

type forest_repository struct {
	Id            string             `json:"id"`
	Path          string             `json:"Path"`
	BuilderPubKey string             `json:"builder_path_key"`
	Environment   forest_environment `json:"environment"`
}

type forest_environment struct {
	Id       string            `json:"id"`
	Branch   string            `json:"branch"`
	Project  forest_project    `json:"project"`
	Services []forest_services `json:"services"`
}

type forest_services struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Tools      []string `json:"tools"`
	DeployPath string   `json:"deploy_path"`
	IndexPath  string   `json:"index_path"`
}
type forest_team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type forest_project struct {
	Id   string      `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
	Team forest_team `json:"team"`
}

type forest_status struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"description"`
}

type forest_build struct {
	Id          string            `json:"id"`
	Repository  forest_repository `json:"repository"`
	User        forest_user       `json:"user"`
	Status      forest_status     `json:"build_status"`
	DeployPath  string            `json:"deploy_path"`
	IndexFile   string            `json:"index_file"`
	Before      string            `json:"before"`
	After       string            `json:"after"`
	Deployed    string            `json:"deployed_at"`
	Started     string            `json:"started_at"`
	Finished_at string            `json:"finished_at"`
	Message     string            `json:"message"`
}

type forest_builds struct {
	Builds []forest_build `json:"builds"`
}

// return object of /v1/auth/refresh_token/
type forest_refresh_token struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}
