package main

type inputStruct struct {
	name     string
	age      float64
	sex      float64
	cp       float64
	trestbps float64
	chol     float64
	fbs      float64
	restecg  float64
	thalach  float64
	exang    float64
	oldpeak  float64
	slope    float64
	ca       float64
	thal     float64
}

type errorStruct struct {
	Message string
	Status  int
}

type responseStruct struct {
	Label float64
}
