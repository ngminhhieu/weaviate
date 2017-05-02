/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateACLEntriesListOKCode is the HTTP code returned for type WeaviateACLEntriesListOK
const WeaviateACLEntriesListOKCode int = 200

/*WeaviateACLEntriesListOK Successful response.

swagger:response weaviateAclEntriesListOK
*/
type WeaviateACLEntriesListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ACLEntriesListResponse `json:"body,omitempty"`
}

// NewWeaviateACLEntriesListOK creates WeaviateACLEntriesListOK with default headers values
func NewWeaviateACLEntriesListOK() *WeaviateACLEntriesListOK {
	return &WeaviateACLEntriesListOK{}
}

// WithPayload adds the payload to the weaviate Acl entries list o k response
func (o *WeaviateACLEntriesListOK) WithPayload(payload *models.ACLEntriesListResponse) *WeaviateACLEntriesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate Acl entries list o k response
func (o *WeaviateACLEntriesListOK) SetPayload(payload *models.ACLEntriesListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateACLEntriesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateACLEntriesListNotFoundCode is the HTTP code returned for type WeaviateACLEntriesListNotFound
const WeaviateACLEntriesListNotFoundCode int = 404

/*WeaviateACLEntriesListNotFound Successful query result but no resource was found

swagger:response weaviateAclEntriesListNotFound
*/
type WeaviateACLEntriesListNotFound struct {
}

// NewWeaviateACLEntriesListNotFound creates WeaviateACLEntriesListNotFound with default headers values
func NewWeaviateACLEntriesListNotFound() *WeaviateACLEntriesListNotFound {
	return &WeaviateACLEntriesListNotFound{}
}

// WriteResponse to the client
func (o *WeaviateACLEntriesListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateACLEntriesListNotImplementedCode is the HTTP code returned for type WeaviateACLEntriesListNotImplemented
const WeaviateACLEntriesListNotImplementedCode int = 501

/*WeaviateACLEntriesListNotImplemented Not (yet) implemented.

swagger:response weaviateAclEntriesListNotImplemented
*/
type WeaviateACLEntriesListNotImplemented struct {
}

// NewWeaviateACLEntriesListNotImplemented creates WeaviateACLEntriesListNotImplemented with default headers values
func NewWeaviateACLEntriesListNotImplemented() *WeaviateACLEntriesListNotImplemented {
	return &WeaviateACLEntriesListNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateACLEntriesListNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
