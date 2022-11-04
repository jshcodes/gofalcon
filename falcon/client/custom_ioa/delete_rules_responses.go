// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// DeleteRulesReader is a Reader for the DeleteRules structure.
type DeleteRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRulesOK creates a DeleteRulesOK with default headers values
func NewDeleteRulesOK() *DeleteRulesOK {
	return &DeleteRulesOK{}
}

/*
DeleteRulesOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRulesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete rules o k response has a 2xx status code
func (o *DeleteRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rules o k response has a 3xx status code
func (o *DeleteRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules o k response has a 4xx status code
func (o *DeleteRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rules o k response has a 5xx status code
func (o *DeleteRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules o k response a status code equal to that given
func (o *DeleteRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRulesOK) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesOK  %+v", 200, o.Payload)
}

func (o *DeleteRulesOK) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesOK  %+v", 200, o.Payload)
}

func (o *DeleteRulesOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesForbidden creates a DeleteRulesForbidden with default headers values
func NewDeleteRulesForbidden() *DeleteRulesForbidden {
	return &DeleteRulesForbidden{}
}

/*
DeleteRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRulesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete rules forbidden response has a 2xx status code
func (o *DeleteRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules forbidden response has a 3xx status code
func (o *DeleteRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules forbidden response has a 4xx status code
func (o *DeleteRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules forbidden response has a 5xx status code
func (o *DeleteRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules forbidden response a status code equal to that given
func (o *DeleteRulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRulesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRulesForbidden) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesNotFound creates a DeleteRulesNotFound with default headers values
func NewDeleteRulesNotFound() *DeleteRulesNotFound {
	return &DeleteRulesNotFound{}
}

/*
DeleteRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRulesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete rules not found response has a 2xx status code
func (o *DeleteRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules not found response has a 3xx status code
func (o *DeleteRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules not found response has a 4xx status code
func (o *DeleteRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules not found response has a 5xx status code
func (o *DeleteRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules not found response a status code equal to that given
func (o *DeleteRulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRulesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRulesNotFound) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRulesNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesTooManyRequests creates a DeleteRulesTooManyRequests with default headers values
func NewDeleteRulesTooManyRequests() *DeleteRulesTooManyRequests {
	return &DeleteRulesTooManyRequests{}
}

/*
DeleteRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteRulesTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete rules too many requests response has a 2xx status code
func (o *DeleteRulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules too many requests response has a 3xx status code
func (o *DeleteRulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules too many requests response has a 4xx status code
func (o *DeleteRulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules too many requests response has a 5xx status code
func (o *DeleteRulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules too many requests response a status code equal to that given
func (o *DeleteRulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRulesTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rules/v1][%d] deleteRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
