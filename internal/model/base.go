package model

type base struct {
	Name       string   `json:"task_name"`
	ScriptPath string   `json:"script_path"`
	Params     []string `json:"params"`
	//LogPath    string `json:"log_path"`
}

type rule struct {
	Rule   string   `json:"rule"`
	Params []string `json:"params,omitempty"`
}
