// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type HTTPHeaders struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *HTTPHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}
