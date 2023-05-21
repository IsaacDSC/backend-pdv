package helpers

type ErrorsResponse struct {
	Error []string `json:"error"`
}

func BuildResponseListError(list_errors []string) (output ErrorsResponse) {
	for index := range list_errors {
		output.Error = append(output.Error, list_errors[index])
	}
	return
}
