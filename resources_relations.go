package corbel

import (
	"fmt"
	"net/http"
)

// RelationData is a basic structure of data relations. By default this are the simplest
// data stored in a relation, but since it's possible to add specific data to the relation
// you can create your own RelationData struct. You must include ID as minimum if you
// want to use GetFromRelation func to retrieve the resource itself.
type RelationData struct {
	Order float64                  `json:"_order,omitempty"`
	ID    string                   `json:"id,omitempty"`
	Links []map[string]interface{} `json:"links, omitempty"`
}

// AddRelation adds the required relation to the resource in the collection
// with the _related_ resource. Additionally arbitrary information can be passed
// to as relation data or nil.
func (r *ResourcesService) AddRelation(collectionName, resourceID, relationName, relatedCollectionName, relatedID string, relationInfo interface{}) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("PUT", "resources", fmt.Sprintf("/v1.0/resource/%s/%s/%s;r=%s/%s", collectionName, resourceID, relationName, relatedCollectionName, relatedID), relationInfo)
	return returnErrorHTTPSimple(r.client, req, err, 201)
}

// MoveRelation sets the required order of the related items on the relationship.
func (r *ResourcesService) MoveRelation(collectionName, resourceID, relationName, relatedCollectionName, relatedID string, order int) error {
	var (
		req *http.Request
		err error
	)

	type orderRelation struct {
		Order string `json:"_order"`
	}

	orderStruct := orderRelation{
		Order: fmt.Sprintf("$pos(%d)", order),
	}

	req, err = r.client.NewRequest("PUT", "resources", fmt.Sprintf("/v1.0/resource/%s/%s/%s;r=%s/%s", collectionName, resourceID, relationName, relatedCollectionName, relatedID), orderStruct)
	return returnErrorHTTPSimple(r.client, req, err, 204)
}

// DeleteRelation deletes the desired relation between the origin and the related
// resource
func (r *ResourcesService) DeleteRelation(collectionName, resourceID, relationName, relatedCollectionName, relatedID string) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("DELETE", "resources", fmt.Sprintf("/v1.0/resource/%s/%s/%s;r=%s/%s", collectionName, resourceID, relationName, relatedCollectionName, relatedID), nil)
	return returnErrorHTTPSimple(r.client, req, err, 204)
}

// DeleteAllRelations deletes all the relations by relationName of the desired resource
func (r *ResourcesService) DeleteAllRelations(collectionName, resourceID, relationName string) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("DELETE", "resources", fmt.Sprintf("/v1.0/resource/%s/%s/%s", collectionName, resourceID, relationName), nil)
	return returnErrorHTTPSimple(r.client, req, err, 204)
}

// SearchRelation returns an instance to the Search Builder
func (r *ResourcesService) SearchRelation(collectionName, resourceID, relationName string) *Search {
	return NewSearch(r.client, "resources", fmt.Sprintf("/v1.0/resource/%s/%s/%s", collectionName, resourceID, relationName))
}
