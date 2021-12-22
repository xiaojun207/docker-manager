package reqDto

type ReqPublishYaml struct {
	ServerNames []string `json:"serverNames"`
	Yaml        string   `json:"yaml"`
}
