// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fx-web/internal/ent/predicate"
	"fx-web/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserName(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetPasswordDigest sets the "password_digest" field.
func (uu *UserUpdate) SetPasswordDigest(s string) *UserUpdate {
	uu.mutation.SetPasswordDigest(s)
	return uu
}

// SetNillablePasswordDigest sets the "password_digest" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswordDigest(s *string) *UserUpdate {
	if s != nil {
		uu.SetPasswordDigest(*s)
	}
	return uu
}

// ClearPasswordDigest clears the value of the "password_digest" field.
func (uu *UserUpdate) ClearPasswordDigest() *UserUpdate {
	uu.mutation.ClearPasswordDigest()
	return uu
}

// SetNickName sets the "nick_name" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// ClearNickName clears the value of the "nick_name" field.
func (uu *UserUpdate) ClearNickName() *UserUpdate {
	uu.mutation.ClearNickName()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(s string) *UserUpdate {
	uu.mutation.SetStatus(s)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(s *string) *UserUpdate {
	if s != nil {
		uu.SetStatus(*s)
	}
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetMoney sets the "money" field.
func (uu *UserUpdate) SetMoney(s string) *UserUpdate {
	uu.mutation.SetMoney(s)
	return uu
}

// SetNillableMoney sets the "money" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMoney(s *string) *UserUpdate {
	if s != nil {
		uu.SetMoney(*s)
	}
	return uu
}

// ClearMoney clears the value of the "money" field.
func (uu *UserUpdate) ClearMoney() *UserUpdate {
	uu.mutation.ClearMoney()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PasswordDigest(); ok {
		if err := user.PasswordDigestValidator(v); err != nil {
			return &ValidationError{Name: "password_digest", err: fmt.Errorf(`ent: validator failed for field "User.password_digest": %w`, err)}
		}
	}
	if v, ok := uu.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "User.nick_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Money(); ok {
		if err := user.MoneyValidator(v); err != nil {
			return &ValidationError{Name: "money", err: fmt.Errorf(`ent: validator failed for field "User.money": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.PasswordDigest(); ok {
		_spec.SetField(user.FieldPasswordDigest, field.TypeString, value)
	}
	if uu.mutation.PasswordDigestCleared() {
		_spec.ClearField(user.FieldPasswordDigest, field.TypeString)
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uu.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.Money(); ok {
		_spec.SetField(user.FieldMoney, field.TypeString, value)
	}
	if uu.mutation.MoneyCleared() {
		_spec.ClearField(user.FieldMoney, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetPasswordDigest sets the "password_digest" field.
func (uuo *UserUpdateOne) SetPasswordDigest(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordDigest(s)
	return uuo
}

// SetNillablePasswordDigest sets the "password_digest" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswordDigest(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPasswordDigest(*s)
	}
	return uuo
}

// ClearPasswordDigest clears the value of the "password_digest" field.
func (uuo *UserUpdateOne) ClearPasswordDigest() *UserUpdateOne {
	uuo.mutation.ClearPasswordDigest()
	return uuo
}

// SetNickName sets the "nick_name" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// ClearNickName clears the value of the "nick_name" field.
func (uuo *UserUpdateOne) ClearNickName() *UserUpdateOne {
	uuo.mutation.ClearNickName()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(s string) *UserUpdateOne {
	uuo.mutation.SetStatus(s)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetStatus(*s)
	}
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetMoney sets the "money" field.
func (uuo *UserUpdateOne) SetMoney(s string) *UserUpdateOne {
	uuo.mutation.SetMoney(s)
	return uuo
}

// SetNillableMoney sets the "money" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMoney(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMoney(*s)
	}
	return uuo
}

// ClearMoney clears the value of the "money" field.
func (uuo *UserUpdateOne) ClearMoney() *UserUpdateOne {
	uuo.mutation.ClearMoney()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PasswordDigest(); ok {
		if err := user.PasswordDigestValidator(v); err != nil {
			return &ValidationError{Name: "password_digest", err: fmt.Errorf(`ent: validator failed for field "User.password_digest": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.NickName(); ok {
		if err := user.NickNameValidator(v); err != nil {
			return &ValidationError{Name: "nick_name", err: fmt.Errorf(`ent: validator failed for field "User.nick_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Money(); ok {
		if err := user.MoneyValidator(v); err != nil {
			return &ValidationError{Name: "money", err: fmt.Errorf(`ent: validator failed for field "User.money": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.PasswordDigest(); ok {
		_spec.SetField(user.FieldPasswordDigest, field.TypeString, value)
	}
	if uuo.mutation.PasswordDigestCleared() {
		_spec.ClearField(user.FieldPasswordDigest, field.TypeString)
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if uuo.mutation.NickNameCleared() {
		_spec.ClearField(user.FieldNickName, field.TypeString)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.Money(); ok {
		_spec.SetField(user.FieldMoney, field.TypeString, value)
	}
	if uuo.mutation.MoneyCleared() {
		_spec.ClearField(user.FieldMoney, field.TypeString)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
