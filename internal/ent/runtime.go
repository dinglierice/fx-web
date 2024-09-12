// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fx-web/internal/ent/psconfig"
	"fx-web/internal/ent/psstrategy"
	"fx-web/internal/ent/schema"
	"fx-web/internal/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	psconfigMixin := schema.PsConfig{}.Mixin()
	psconfigMixinFields0 := psconfigMixin[0].Fields()
	_ = psconfigMixinFields0
	psconfigFields := schema.PsConfig{}.Fields()
	_ = psconfigFields
	// psconfigDescCreatedAt is the schema descriptor for created_at field.
	psconfigDescCreatedAt := psconfigMixinFields0[0].Descriptor()
	// psconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	psconfig.DefaultCreatedAt = psconfigDescCreatedAt.Default.(func() time.Time)
	// psconfigDescUpdatedAt is the schema descriptor for updated_at field.
	psconfigDescUpdatedAt := psconfigMixinFields0[1].Descriptor()
	// psconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	psconfig.UpdateDefaultUpdatedAt = psconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// psconfigDescPsScene is the schema descriptor for ps_scene field.
	psconfigDescPsScene := psconfigFields[1].Descriptor()
	// psconfig.PsSceneValidator is a validator for the "ps_scene" field. It is called by the builders before save.
	psconfig.PsSceneValidator = psconfigDescPsScene.Validators[0].(func(string) error)
	// psconfigDescID is the schema descriptor for id field.
	psconfigDescID := psconfigFields[0].Descriptor()
	// psconfig.IDValidator is a validator for the "id" field. It is called by the builders before save.
	psconfig.IDValidator = psconfigDescID.Validators[0].(func(uint64) error)
	psstrategyMixin := schema.PsStrategy{}.Mixin()
	psstrategyMixinFields0 := psstrategyMixin[0].Fields()
	_ = psstrategyMixinFields0
	psstrategyFields := schema.PsStrategy{}.Fields()
	_ = psstrategyFields
	// psstrategyDescCreatedAt is the schema descriptor for created_at field.
	psstrategyDescCreatedAt := psstrategyMixinFields0[0].Descriptor()
	// psstrategy.DefaultCreatedAt holds the default value on creation for the created_at field.
	psstrategy.DefaultCreatedAt = psstrategyDescCreatedAt.Default.(func() time.Time)
	// psstrategyDescUpdatedAt is the schema descriptor for updated_at field.
	psstrategyDescUpdatedAt := psstrategyMixinFields0[1].Descriptor()
	// psstrategy.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	psstrategy.UpdateDefaultUpdatedAt = psstrategyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// psstrategyDescIsDelete is the schema descriptor for is_delete field.
	psstrategyDescIsDelete := psstrategyFields[3].Descriptor()
	// psstrategy.DefaultIsDelete holds the default value on creation for the is_delete field.
	psstrategy.DefaultIsDelete = psstrategyDescIsDelete.Default.(int)
	// psstrategyDescID is the schema descriptor for id field.
	psstrategyDescID := psstrategyFields[0].Descriptor()
	// psstrategy.IDValidator is a validator for the "id" field. It is called by the builders before save.
	psstrategy.IDValidator = psstrategyDescID.Validators[0].(func(uint64) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[2].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordDigest is the schema descriptor for password_digest field.
	userDescPasswordDigest := userFields[4].Descriptor()
	// user.PasswordDigestValidator is a validator for the "password_digest" field. It is called by the builders before save.
	user.PasswordDigestValidator = userDescPasswordDigest.Validators[0].(func(string) error)
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[5].Descriptor()
	// user.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[6].Descriptor()
	// user.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	user.StatusValidator = userDescStatus.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[7].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescMoney is the schema descriptor for money field.
	userDescMoney := userFields[8].Descriptor()
	// user.MoneyValidator is a validator for the "money" field. It is called by the builders before save.
	user.MoneyValidator = userDescMoney.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint64) error)
}
