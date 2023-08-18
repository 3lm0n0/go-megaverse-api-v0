package domain

type Polyanet struct {
	CandidateId string
	Row int `json: "row"`
	Column int `json: "column"`
}

type MatrixConfiguration struct {
	Rows int `json: "rows"`
	Columns int `json: "columns"`
	Offset int `json: "offset"`
}

type Matrix struct {
	Data [][]int
}
