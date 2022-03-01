// Code generated by go-swagger; DO NOT EDIT.

package cluster

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
	"github.com/go-openapi/swag"
)

// NewClusterGetClustersParams creates a new ClusterGetClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterGetClustersParams() *ClusterGetClustersParams {
	return &ClusterGetClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterGetClustersParamsWithTimeout creates a new ClusterGetClustersParams object
// with the ability to set a timeout on a request.
func NewClusterGetClustersParamsWithTimeout(timeout time.Duration) *ClusterGetClustersParams {
	return &ClusterGetClustersParams{
		timeout: timeout,
	}
}

// NewClusterGetClustersParamsWithContext creates a new ClusterGetClustersParams object
// with the ability to set a context for a request.
func NewClusterGetClustersParamsWithContext(ctx context.Context) *ClusterGetClustersParams {
	return &ClusterGetClustersParams{
		Context: ctx,
	}
}

// NewClusterGetClustersParamsWithHTTPClient creates a new ClusterGetClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterGetClustersParamsWithHTTPClient(client *http.Client) *ClusterGetClustersParams {
	return &ClusterGetClustersParams{
		HTTPClient: client,
	}
}

/* ClusterGetClustersParams contains all the parameters to send to the API endpoint
   for the cluster get clusters operation.

   Typically these are written to a http.Request.
*/
type ClusterGetClustersParams struct {

	// ID.
	ID *string

	// BlueprintRef.
	BlueprintRef *string

	// ClusterID.
	ClusterID *string

	// Count.
	//
	// Format: int64
	Count *string

	// Deleted.
	Deleted *bool

	/* DisplayName.

	   displayName only used for update queries to set displayName (READONLY).
	*/
	DisplayName *string

	// Extended.
	Extended *bool

	/* GlobalScope.

	   globalScope sets partnerID,organizationID,projectID = 0.
	*/
	GlobalScope *bool

	// Groups.
	Groups []string

	/* IgnoreScopeDefault.

	     ignoreScopeDefault ignores default values for partnerID, organizationID and
	projectID.
	*/
	IgnoreScopeDefault *bool

	// IsSSOUser.
	IsSSOUser *bool

	// Limit.
	//
	// Format: int64
	Limit *string

	/* Name.

	     name is unique ID of a resource along with (partnerID, organizationID,
	projectID).
	*/
	Name *string

	// Offset.
	//
	// Format: int64
	Offset *string

	// Order.
	Order *string

	// OrderBy.
	OrderBy *string

	// OrganizationID.
	OrganizationID *string

	// PartnerID.
	PartnerID *string

	// ProjectID.
	ProjectID string

	// PublishedVersion.
	PublishedVersion *string

	/* Selector.

	   selector is used to filter the labels of a resource.
	*/
	Selector *string

	/* URLScope.

	   urlScope is supposed to be passed in the URL as kind/HashID(value).
	*/
	URLScope *string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGetClustersParams) WithDefaults() *ClusterGetClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterGetClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster get clusters params
func (o *ClusterGetClustersParams) WithTimeout(timeout time.Duration) *ClusterGetClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster get clusters params
func (o *ClusterGetClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster get clusters params
func (o *ClusterGetClustersParams) WithContext(ctx context.Context) *ClusterGetClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster get clusters params
func (o *ClusterGetClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster get clusters params
func (o *ClusterGetClustersParams) WithHTTPClient(client *http.Client) *ClusterGetClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster get clusters params
func (o *ClusterGetClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cluster get clusters params
func (o *ClusterGetClustersParams) WithID(id *string) *ClusterGetClustersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cluster get clusters params
func (o *ClusterGetClustersParams) SetID(id *string) {
	o.ID = id
}

// WithBlueprintRef adds the blueprintRef to the cluster get clusters params
func (o *ClusterGetClustersParams) WithBlueprintRef(blueprintRef *string) *ClusterGetClustersParams {
	o.SetBlueprintRef(blueprintRef)
	return o
}

// SetBlueprintRef adds the blueprintRef to the cluster get clusters params
func (o *ClusterGetClustersParams) SetBlueprintRef(blueprintRef *string) {
	o.BlueprintRef = blueprintRef
}

// WithClusterID adds the clusterID to the cluster get clusters params
func (o *ClusterGetClustersParams) WithClusterID(clusterID *string) *ClusterGetClustersParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cluster get clusters params
func (o *ClusterGetClustersParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithCount adds the count to the cluster get clusters params
func (o *ClusterGetClustersParams) WithCount(count *string) *ClusterGetClustersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the cluster get clusters params
func (o *ClusterGetClustersParams) SetCount(count *string) {
	o.Count = count
}

// WithDeleted adds the deleted to the cluster get clusters params
func (o *ClusterGetClustersParams) WithDeleted(deleted *bool) *ClusterGetClustersParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the cluster get clusters params
func (o *ClusterGetClustersParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDisplayName adds the displayName to the cluster get clusters params
func (o *ClusterGetClustersParams) WithDisplayName(displayName *string) *ClusterGetClustersParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the cluster get clusters params
func (o *ClusterGetClustersParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithExtended adds the extended to the cluster get clusters params
func (o *ClusterGetClustersParams) WithExtended(extended *bool) *ClusterGetClustersParams {
	o.SetExtended(extended)
	return o
}

// SetExtended adds the extended to the cluster get clusters params
func (o *ClusterGetClustersParams) SetExtended(extended *bool) {
	o.Extended = extended
}

// WithGlobalScope adds the globalScope to the cluster get clusters params
func (o *ClusterGetClustersParams) WithGlobalScope(globalScope *bool) *ClusterGetClustersParams {
	o.SetGlobalScope(globalScope)
	return o
}

// SetGlobalScope adds the globalScope to the cluster get clusters params
func (o *ClusterGetClustersParams) SetGlobalScope(globalScope *bool) {
	o.GlobalScope = globalScope
}

// WithGroups adds the groups to the cluster get clusters params
func (o *ClusterGetClustersParams) WithGroups(groups []string) *ClusterGetClustersParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the cluster get clusters params
func (o *ClusterGetClustersParams) SetGroups(groups []string) {
	o.Groups = groups
}

// WithIgnoreScopeDefault adds the ignoreScopeDefault to the cluster get clusters params
func (o *ClusterGetClustersParams) WithIgnoreScopeDefault(ignoreScopeDefault *bool) *ClusterGetClustersParams {
	o.SetIgnoreScopeDefault(ignoreScopeDefault)
	return o
}

// SetIgnoreScopeDefault adds the ignoreScopeDefault to the cluster get clusters params
func (o *ClusterGetClustersParams) SetIgnoreScopeDefault(ignoreScopeDefault *bool) {
	o.IgnoreScopeDefault = ignoreScopeDefault
}

// WithIsSSOUser adds the isSSOUser to the cluster get clusters params
func (o *ClusterGetClustersParams) WithIsSSOUser(isSSOUser *bool) *ClusterGetClustersParams {
	o.SetIsSSOUser(isSSOUser)
	return o
}

// SetIsSSOUser adds the isSSOUser to the cluster get clusters params
func (o *ClusterGetClustersParams) SetIsSSOUser(isSSOUser *bool) {
	o.IsSSOUser = isSSOUser
}

// WithLimit adds the limit to the cluster get clusters params
func (o *ClusterGetClustersParams) WithLimit(limit *string) *ClusterGetClustersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the cluster get clusters params
func (o *ClusterGetClustersParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the cluster get clusters params
func (o *ClusterGetClustersParams) WithName(name *string) *ClusterGetClustersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the cluster get clusters params
func (o *ClusterGetClustersParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the cluster get clusters params
func (o *ClusterGetClustersParams) WithOffset(offset *string) *ClusterGetClustersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the cluster get clusters params
func (o *ClusterGetClustersParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrder adds the order to the cluster get clusters params
func (o *ClusterGetClustersParams) WithOrder(order *string) *ClusterGetClustersParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the cluster get clusters params
func (o *ClusterGetClustersParams) SetOrder(order *string) {
	o.Order = order
}

// WithOrderBy adds the orderBy to the cluster get clusters params
func (o *ClusterGetClustersParams) WithOrderBy(orderBy *string) *ClusterGetClustersParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cluster get clusters params
func (o *ClusterGetClustersParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationID adds the organizationID to the cluster get clusters params
func (o *ClusterGetClustersParams) WithOrganizationID(organizationID *string) *ClusterGetClustersParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the cluster get clusters params
func (o *ClusterGetClustersParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithPartnerID adds the partnerID to the cluster get clusters params
func (o *ClusterGetClustersParams) WithPartnerID(partnerID *string) *ClusterGetClustersParams {
	o.SetPartnerID(partnerID)
	return o
}

// SetPartnerID adds the partnerId to the cluster get clusters params
func (o *ClusterGetClustersParams) SetPartnerID(partnerID *string) {
	o.PartnerID = partnerID
}

// WithProjectID adds the projectID to the cluster get clusters params
func (o *ClusterGetClustersParams) WithProjectID(projectID string) *ClusterGetClustersParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the cluster get clusters params
func (o *ClusterGetClustersParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithPublishedVersion adds the publishedVersion to the cluster get clusters params
func (o *ClusterGetClustersParams) WithPublishedVersion(publishedVersion *string) *ClusterGetClustersParams {
	o.SetPublishedVersion(publishedVersion)
	return o
}

// SetPublishedVersion adds the publishedVersion to the cluster get clusters params
func (o *ClusterGetClustersParams) SetPublishedVersion(publishedVersion *string) {
	o.PublishedVersion = publishedVersion
}

// WithSelector adds the selector to the cluster get clusters params
func (o *ClusterGetClustersParams) WithSelector(selector *string) *ClusterGetClustersParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the cluster get clusters params
func (o *ClusterGetClustersParams) SetSelector(selector *string) {
	o.Selector = selector
}

// WithURLScope adds the uRLScope to the cluster get clusters params
func (o *ClusterGetClustersParams) WithURLScope(uRLScope *string) *ClusterGetClustersParams {
	o.SetURLScope(uRLScope)
	return o
}

// SetURLScope adds the urlScope to the cluster get clusters params
func (o *ClusterGetClustersParams) SetURLScope(uRLScope *string) {
	o.URLScope = uRLScope
}

// WithUsername adds the username to the cluster get clusters params
func (o *ClusterGetClustersParams) WithUsername(username *string) *ClusterGetClustersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the cluster get clusters params
func (o *ClusterGetClustersParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterGetClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param ID
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("ID", qID); err != nil {
				return err
			}
		}
	}

	if o.BlueprintRef != nil {

		// query param blueprintRef
		var qrBlueprintRef string

		if o.BlueprintRef != nil {
			qrBlueprintRef = *o.BlueprintRef
		}
		qBlueprintRef := qrBlueprintRef
		if qBlueprintRef != "" {

			if err := r.SetQueryParam("blueprintRef", qBlueprintRef); err != nil {
				return err
			}
		}
	}

	if o.ClusterID != nil {

		// query param clusterID
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("clusterID", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount string

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := qrCount
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string

		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {

			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}
	}

	if o.Extended != nil {

		// query param extended
		var qrExtended bool

		if o.Extended != nil {
			qrExtended = *o.Extended
		}
		qExtended := swag.FormatBool(qrExtended)
		if qExtended != "" {

			if err := r.SetQueryParam("extended", qExtended); err != nil {
				return err
			}
		}
	}

	if o.GlobalScope != nil {

		// query param globalScope
		var qrGlobalScope bool

		if o.GlobalScope != nil {
			qrGlobalScope = *o.GlobalScope
		}
		qGlobalScope := swag.FormatBool(qrGlobalScope)
		if qGlobalScope != "" {

			if err := r.SetQueryParam("globalScope", qGlobalScope); err != nil {
				return err
			}
		}
	}

	if o.Groups != nil {

		// binding items for groups
		joinedGroups := o.bindParamGroups(reg)

		// query array param groups
		if err := r.SetQueryParam("groups", joinedGroups...); err != nil {
			return err
		}
	}

	if o.IgnoreScopeDefault != nil {

		// query param ignoreScopeDefault
		var qrIgnoreScopeDefault bool

		if o.IgnoreScopeDefault != nil {
			qrIgnoreScopeDefault = *o.IgnoreScopeDefault
		}
		qIgnoreScopeDefault := swag.FormatBool(qrIgnoreScopeDefault)
		if qIgnoreScopeDefault != "" {

			if err := r.SetQueryParam("ignoreScopeDefault", qIgnoreScopeDefault); err != nil {
				return err
			}
		}
	}

	if o.IsSSOUser != nil {

		// query param isSSOUser
		var qrIsSSOUser bool

		if o.IsSSOUser != nil {
			qrIsSSOUser = *o.IsSSOUser
		}
		qIsSSOUser := swag.FormatBool(qrIsSSOUser)
		if qIsSSOUser != "" {

			if err := r.SetQueryParam("isSSOUser", qIsSSOUser); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.OrganizationID != nil {

		// query param organizationID
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationID", qOrganizationID); err != nil {
				return err
			}
		}
	}

	if o.PartnerID != nil {

		// query param partnerID
		var qrPartnerID string

		if o.PartnerID != nil {
			qrPartnerID = *o.PartnerID
		}
		qPartnerID := qrPartnerID
		if qPartnerID != "" {

			if err := r.SetQueryParam("partnerID", qPartnerID); err != nil {
				return err
			}
		}
	}

	// path param projectID
	if err := r.SetPathParam("projectID", o.ProjectID); err != nil {
		return err
	}

	if o.PublishedVersion != nil {

		// query param publishedVersion
		var qrPublishedVersion string

		if o.PublishedVersion != nil {
			qrPublishedVersion = *o.PublishedVersion
		}
		qPublishedVersion := qrPublishedVersion
		if qPublishedVersion != "" {

			if err := r.SetQueryParam("publishedVersion", qPublishedVersion); err != nil {
				return err
			}
		}
	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string

		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {

			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}
	}

	if o.URLScope != nil {

		// query param urlScope
		var qrURLScope string

		if o.URLScope != nil {
			qrURLScope = *o.URLScope
		}
		qURLScope := qrURLScope
		if qURLScope != "" {

			if err := r.SetQueryParam("urlScope", qURLScope); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamClusterGetClusters binds the parameter groups
func (o *ClusterGetClustersParams) bindParamGroups(formats strfmt.Registry) []string {
	groupsIR := o.Groups

	var groupsIC []string
	for _, groupsIIR := range groupsIR { // explode []string

		groupsIIV := groupsIIR // string as string
		groupsIC = append(groupsIC, groupsIIV)
	}

	// items.CollectionFormat: "multi"
	groupsIS := swag.JoinByFormat(groupsIC, "multi")

	return groupsIS
}