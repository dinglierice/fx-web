// Code generated by ent, DO NOT EDIT.

package user

import (
	"fx-web/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// PasswordDigest applies equality check predicate on the "password_digest" field. It's identical to PasswordDigestEQ.
func PasswordDigest(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordDigest, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// Money applies equality check predicate on the "money" field. It's identical to MoneyEQ.
func Money(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMoney, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeletedAt))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordDigestEQ applies the EQ predicate on the "password_digest" field.
func PasswordDigestEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordDigest, v))
}

// PasswordDigestNEQ applies the NEQ predicate on the "password_digest" field.
func PasswordDigestNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordDigest, v))
}

// PasswordDigestIn applies the In predicate on the "password_digest" field.
func PasswordDigestIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordDigest, vs...))
}

// PasswordDigestNotIn applies the NotIn predicate on the "password_digest" field.
func PasswordDigestNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordDigest, vs...))
}

// PasswordDigestGT applies the GT predicate on the "password_digest" field.
func PasswordDigestGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordDigest, v))
}

// PasswordDigestGTE applies the GTE predicate on the "password_digest" field.
func PasswordDigestGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordDigest, v))
}

// PasswordDigestLT applies the LT predicate on the "password_digest" field.
func PasswordDigestLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordDigest, v))
}

// PasswordDigestLTE applies the LTE predicate on the "password_digest" field.
func PasswordDigestLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordDigest, v))
}

// PasswordDigestContains applies the Contains predicate on the "password_digest" field.
func PasswordDigestContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPasswordDigest, v))
}

// PasswordDigestHasPrefix applies the HasPrefix predicate on the "password_digest" field.
func PasswordDigestHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPasswordDigest, v))
}

// PasswordDigestHasSuffix applies the HasSuffix predicate on the "password_digest" field.
func PasswordDigestHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPasswordDigest, v))
}

// PasswordDigestIsNil applies the IsNil predicate on the "password_digest" field.
func PasswordDigestIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPasswordDigest))
}

// PasswordDigestNotNil applies the NotNil predicate on the "password_digest" field.
func PasswordDigestNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPasswordDigest))
}

// PasswordDigestEqualFold applies the EqualFold predicate on the "password_digest" field.
func PasswordDigestEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPasswordDigest, v))
}

// PasswordDigestContainsFold applies the ContainsFold predicate on the "password_digest" field.
func PasswordDigestContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPasswordDigest, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameIsNil applies the IsNil predicate on the "nick_name" field.
func NickNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldNickName))
}

// NickNameNotNil applies the NotNil predicate on the "nick_name" field.
func NickNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldNickName))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldStatus, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatar, v))
}

// MoneyEQ applies the EQ predicate on the "money" field.
func MoneyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMoney, v))
}

// MoneyNEQ applies the NEQ predicate on the "money" field.
func MoneyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMoney, v))
}

// MoneyIn applies the In predicate on the "money" field.
func MoneyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMoney, vs...))
}

// MoneyNotIn applies the NotIn predicate on the "money" field.
func MoneyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMoney, vs...))
}

// MoneyGT applies the GT predicate on the "money" field.
func MoneyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMoney, v))
}

// MoneyGTE applies the GTE predicate on the "money" field.
func MoneyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMoney, v))
}

// MoneyLT applies the LT predicate on the "money" field.
func MoneyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMoney, v))
}

// MoneyLTE applies the LTE predicate on the "money" field.
func MoneyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMoney, v))
}

// MoneyContains applies the Contains predicate on the "money" field.
func MoneyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMoney, v))
}

// MoneyHasPrefix applies the HasPrefix predicate on the "money" field.
func MoneyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMoney, v))
}

// MoneyHasSuffix applies the HasSuffix predicate on the "money" field.
func MoneyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMoney, v))
}

// MoneyIsNil applies the IsNil predicate on the "money" field.
func MoneyIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldMoney))
}

// MoneyNotNil applies the NotNil predicate on the "money" field.
func MoneyNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldMoney))
}

// MoneyEqualFold applies the EqualFold predicate on the "money" field.
func MoneyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMoney, v))
}

// MoneyContainsFold applies the ContainsFold predicate on the "money" field.
func MoneyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMoney, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
