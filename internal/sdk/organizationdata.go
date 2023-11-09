// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/operations"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/sdkerrors"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/models/shared"
	"github.com/squat/terraform-provider-saladcloud/internal/sdk/pkg/utils"
	"io"
	"net/http"
)

// OrganizationData - Auxiliary organization data and info
type OrganizationData struct {
	sdkConfiguration sdkConfiguration
}

func newOrganizationData(sdkConfig sdkConfiguration) *OrganizationData {
	return &OrganizationData{
		sdkConfiguration: sdkConfig,
	}
}

// ListGpuClasses - List the GPU Classes
// List the GPU Classes
func (s *OrganizationData) ListGpuClasses(ctx context.Context, request operations.ListGpuClassesRequest) (*operations.ListGpuClassesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/organizations/{organization_name}/gpu-classes", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListGpuClassesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.GpuClassesList
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GpuClassesList = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ProblemDetails
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ProblemDetails = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
