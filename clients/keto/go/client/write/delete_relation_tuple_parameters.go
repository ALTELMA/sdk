// Code generated by go-swagger; DO NOT EDIT.

package write

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRelationTupleParams creates a new DeleteRelationTupleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRelationTupleParams() *DeleteRelationTupleParams {
	return &DeleteRelationTupleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRelationTupleParamsWithTimeout creates a new DeleteRelationTupleParams object
// with the ability to set a timeout on a request.
func NewDeleteRelationTupleParamsWithTimeout(timeout time.Duration) *DeleteRelationTupleParams {
	return &DeleteRelationTupleParams{
		timeout: timeout,
	}
}

// NewDeleteRelationTupleParamsWithContext creates a new DeleteRelationTupleParams object
// with the ability to set a context for a request.
func NewDeleteRelationTupleParamsWithContext(ctx context.Context) *DeleteRelationTupleParams {
	return &DeleteRelationTupleParams{
		Context: ctx,
	}
}

// NewDeleteRelationTupleParamsWithHTTPClient creates a new DeleteRelationTupleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRelationTupleParamsWithHTTPClient(client *http.Client) *DeleteRelationTupleParams {
	return &DeleteRelationTupleParams{
		HTTPClient: client,
	}
}

/* DeleteRelationTupleParams contains all the parameters to send to the API endpoint
   for the delete relation tuple operation.

   Typically these are written to a http.Request.
*/
type DeleteRelationTupleParams struct {

	/* Namespace.

	   Namespace of the Relation Tuple
	*/
	Namespace string

	/* Object.

	   Object of the Relation Tuple
	*/
	Object string

	/* Relation.

	   Relation of the Relation Tuple
	*/
	Relation string

	/* SubjectID.

	   SubjectID of the Relation Tuple
	*/
	SubjectID *string

	/* SubjectSetNamespace.

	   Namespace of the Subject Set
	*/
	SubjectSetNamespace *string

	/* SubjectSetObject.

	   Object of the Subject Set
	*/
	SubjectSetObject *string

	/* SubjectSetRelation.

	   Relation of the Subject Set
	*/
	SubjectSetRelation *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete relation tuple params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRelationTupleParams) WithDefaults() *DeleteRelationTupleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete relation tuple params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRelationTupleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithTimeout(timeout time.Duration) *DeleteRelationTupleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithContext(ctx context.Context) *DeleteRelationTupleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithHTTPClient(client *http.Client) *DeleteRelationTupleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithNamespace(namespace string) *DeleteRelationTupleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithObject adds the object to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithObject(object string) *DeleteRelationTupleParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetObject(object string) {
	o.Object = object
}

// WithRelation adds the relation to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithRelation(relation string) *DeleteRelationTupleParams {
	o.SetRelation(relation)
	return o
}

// SetRelation adds the relation to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetRelation(relation string) {
	o.Relation = relation
}

// WithSubjectID adds the subjectID to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithSubjectID(subjectID *string) *DeleteRelationTupleParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetSubjectID(subjectID *string) {
	o.SubjectID = subjectID
}

// WithSubjectSetNamespace adds the subjectSetNamespace to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithSubjectSetNamespace(subjectSetNamespace *string) *DeleteRelationTupleParams {
	o.SetSubjectSetNamespace(subjectSetNamespace)
	return o
}

// SetSubjectSetNamespace adds the subjectSetNamespace to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetSubjectSetNamespace(subjectSetNamespace *string) {
	o.SubjectSetNamespace = subjectSetNamespace
}

// WithSubjectSetObject adds the subjectSetObject to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithSubjectSetObject(subjectSetObject *string) *DeleteRelationTupleParams {
	o.SetSubjectSetObject(subjectSetObject)
	return o
}

// SetSubjectSetObject adds the subjectSetObject to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetSubjectSetObject(subjectSetObject *string) {
	o.SubjectSetObject = subjectSetObject
}

// WithSubjectSetRelation adds the subjectSetRelation to the delete relation tuple params
func (o *DeleteRelationTupleParams) WithSubjectSetRelation(subjectSetRelation *string) *DeleteRelationTupleParams {
	o.SetSubjectSetRelation(subjectSetRelation)
	return o
}

// SetSubjectSetRelation adds the subjectSetRelation to the delete relation tuple params
func (o *DeleteRelationTupleParams) SetSubjectSetRelation(subjectSetRelation *string) {
	o.SubjectSetRelation = subjectSetRelation
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRelationTupleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {

		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	// query param object
	qrObject := o.Object
	qObject := qrObject
	if qObject != "" {

		if err := r.SetQueryParam("object", qObject); err != nil {
			return err
		}
	}

	// query param relation
	qrRelation := o.Relation
	qRelation := qrRelation
	if qRelation != "" {

		if err := r.SetQueryParam("relation", qRelation); err != nil {
			return err
		}
	}

	if o.SubjectID != nil {

		// query param subject_id
		var qrSubjectID string

		if o.SubjectID != nil {
			qrSubjectID = *o.SubjectID
		}
		qSubjectID := qrSubjectID
		if qSubjectID != "" {

			if err := r.SetQueryParam("subject_id", qSubjectID); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetNamespace != nil {

		// query param subject_set.namespace
		var qrSubjectSetNamespace string

		if o.SubjectSetNamespace != nil {
			qrSubjectSetNamespace = *o.SubjectSetNamespace
		}
		qSubjectSetNamespace := qrSubjectSetNamespace
		if qSubjectSetNamespace != "" {

			if err := r.SetQueryParam("subject_set.namespace", qSubjectSetNamespace); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetObject != nil {

		// query param subject_set.object
		var qrSubjectSetObject string

		if o.SubjectSetObject != nil {
			qrSubjectSetObject = *o.SubjectSetObject
		}
		qSubjectSetObject := qrSubjectSetObject
		if qSubjectSetObject != "" {

			if err := r.SetQueryParam("subject_set.object", qSubjectSetObject); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetRelation != nil {

		// query param subject_set.relation
		var qrSubjectSetRelation string

		if o.SubjectSetRelation != nil {
			qrSubjectSetRelation = *o.SubjectSetRelation
		}
		qSubjectSetRelation := qrSubjectSetRelation
		if qSubjectSetRelation != "" {

			if err := r.SetQueryParam("subject_set.relation", qSubjectSetRelation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
