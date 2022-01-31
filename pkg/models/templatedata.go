package models

//This is a struct containig all the possible datatypes that can be sent to template for rendering
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}           //when we do not know datatype we use interface instead
	Flash string
	Warning string
	Error string
}