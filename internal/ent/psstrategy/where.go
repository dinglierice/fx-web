// Code generated by ent, DO NOT EDIT.

package psstrategy

import (
	"fx-web/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldUpdatedAt, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldOwner, v))
}

// ScriptContent applies equality check predicate on the "script_content" field. It's identical to ScriptContentEQ.
func ScriptContent(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldScriptContent, v))
}

// IsDelete applies equality check predicate on the "is_delete" field. It's identical to IsDeleteEQ.
func IsDelete(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldIsDelete, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldUpdatedAt, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v uint64) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldOwner, v))
}

// ScriptContentEQ applies the EQ predicate on the "script_content" field.
func ScriptContentEQ(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldScriptContent, v))
}

// ScriptContentNEQ applies the NEQ predicate on the "script_content" field.
func ScriptContentNEQ(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldScriptContent, v))
}

// ScriptContentIn applies the In predicate on the "script_content" field.
func ScriptContentIn(vs ...string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldScriptContent, vs...))
}

// ScriptContentNotIn applies the NotIn predicate on the "script_content" field.
func ScriptContentNotIn(vs ...string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldScriptContent, vs...))
}

// ScriptContentGT applies the GT predicate on the "script_content" field.
func ScriptContentGT(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldScriptContent, v))
}

// ScriptContentGTE applies the GTE predicate on the "script_content" field.
func ScriptContentGTE(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldScriptContent, v))
}

// ScriptContentLT applies the LT predicate on the "script_content" field.
func ScriptContentLT(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldScriptContent, v))
}

// ScriptContentLTE applies the LTE predicate on the "script_content" field.
func ScriptContentLTE(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldScriptContent, v))
}

// ScriptContentContains applies the Contains predicate on the "script_content" field.
func ScriptContentContains(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldContains(FieldScriptContent, v))
}

// ScriptContentHasPrefix applies the HasPrefix predicate on the "script_content" field.
func ScriptContentHasPrefix(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldHasPrefix(FieldScriptContent, v))
}

// ScriptContentHasSuffix applies the HasSuffix predicate on the "script_content" field.
func ScriptContentHasSuffix(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldHasSuffix(FieldScriptContent, v))
}

// ScriptContentEqualFold applies the EqualFold predicate on the "script_content" field.
func ScriptContentEqualFold(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEqualFold(FieldScriptContent, v))
}

// ScriptContentContainsFold applies the ContainsFold predicate on the "script_content" field.
func ScriptContentContainsFold(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldContainsFold(FieldScriptContent, v))
}

// IsDeleteEQ applies the EQ predicate on the "is_delete" field.
func IsDeleteEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldIsDelete, v))
}

// IsDeleteNEQ applies the NEQ predicate on the "is_delete" field.
func IsDeleteNEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldIsDelete, v))
}

// IsDeleteIn applies the In predicate on the "is_delete" field.
func IsDeleteIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldIsDelete, vs...))
}

// IsDeleteNotIn applies the NotIn predicate on the "is_delete" field.
func IsDeleteNotIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldIsDelete, vs...))
}

// IsDeleteGT applies the GT predicate on the "is_delete" field.
func IsDeleteGT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldIsDelete, v))
}

// IsDeleteGTE applies the GTE predicate on the "is_delete" field.
func IsDeleteGTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldIsDelete, v))
}

// IsDeleteLT applies the LT predicate on the "is_delete" field.
func IsDeleteLT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldIsDelete, v))
}

// IsDeleteLTE applies the LTE predicate on the "is_delete" field.
func IsDeleteLTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldIsDelete, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.NotPredicates(p))
}
