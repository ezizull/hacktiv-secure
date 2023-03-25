package domain

type ElementRequest struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
}

type ElementResponse struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	WaterStatus string 
	WindStatus string 
}