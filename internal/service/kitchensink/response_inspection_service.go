package kitchensink

import (
	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

type ResponseInspectionService struct{}

func NewResponseInspectionService() *ResponseInspectionService {
	return &ResponseInspectionService{}
}

func (s *ResponseInspectionService) GetJSONResponse() *kitchensink.ResponseInspectionResponse {
	return &kitchensink.ResponseInspectionResponse{
		Message: "This is a JSON response",
		Format:  "json",
	}
}

func (s *ResponseInspectionService) GetXMLResponse() *kitchensink.ResponseInspectionResponse {
	return &kitchensink.ResponseInspectionResponse{
		Message: "This is an XML response",
		Format:  "xml",
	}
}

func (s *ResponseInspectionService) GetHTMLResponse() string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>HTML Response</title>
</head>
<body>
    <h1>This is an HTML response</h1>
</body>
</html>`
}
