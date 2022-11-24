package models

type Tester struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ReturnTester struct {
	Name      string `json:"name"`
	DoubleAge int    `json:"double-age"`
}
