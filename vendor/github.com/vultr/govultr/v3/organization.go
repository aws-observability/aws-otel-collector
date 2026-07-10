package govultr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	organizationPath = "/v2/organizations"
	invitationPath   = "/v2/invitation"
	groupPath        = "/v2/groups"
	rolePath         = "/v2/roles"
	policyPath       = "/v2/policies"
	roleTrustPath    = "/v2/role-trusts"
	roleSessionPath  = "/v2/assumed-roles"
	userPath         = "/v2/users"
)

// OrganizationService is the interface to interact with organization access endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/iam
type OrganizationService interface {
	CreateOrganization(ctx context.Context, organizationReq *OrganizationReq) (*Organization, *http.Response, error)
	GetOrganization(ctx context.Context, organizationID string) (*Organization, *http.Response, error)
	ListOrganizations(ctx context.Context, options *ListOptions) ([]Organization, *Meta, *http.Response, error)
	UpdateOrganization(ctx context.Context, organizationID string, organizationReq *OrganizationReq) (*Organization, *http.Response, error)
	DeleteOrganization(ctx context.Context, organizationID string) error
	RestoreOrganization(ctx context.Context, organizationID string) error

	DeleteUser(ctx context.Context, organizationID, userID string) error
	SuspendUser(ctx context.Context, organizationID, userID string) error
	UnsuspendUser(ctx context.Context, organizationID, userID string) error
	ListSuspendedUsers(ctx context.Context, organizationIDstring string, options *ListOptions) ([]OrganizationUser, *Meta, *http.Response, error) //nolint:lll

	CreateInvitation(ctx context.Context, invitationReq *OrganizationInvitationReq) (*OrganizationInvitation, *http.Response, error) //nolint:lll
	GetInvitation(ctx context.Context, invitationID string) (*OrganizationInvitation, *http.Response, error)
	ListInvitations(ctx context.Context, options *ListOptions) ([]OrganizationInvitation, *Meta, *http.Response, error)
	ResendInvitation(ctx context.Context, invitationID string) (*OrganizationInvitation, *http.Response, error)

	CreateGroup(ctx context.Context, groupReq *OrganizationGroupReq) (*OrganizationGroup, *http.Response, error)
	GetGroup(ctx context.Context, groupID string) (*OrganizationGroup, *http.Response, error)
	ListGroups(ctx context.Context, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error)
	UpdateGroup(ctx context.Context, groupID string, groupReq *OrganizationGroupReq) (*OrganizationGroup, *http.Response, error)
	DeleteGroup(ctx context.Context, groupID string) error
	AddGroupMember(ctx context.Context, groupID string, memberReq *OrganizationGroupMemberReq) error
	RemoveGroupMember(ctx context.Context, groupID, userID string) error
	ListGroupPolicies(ctx context.Context, groupID string) (*OrganizationGroupPolicies, *Meta, *http.Response, error)
	ListGroupRoles(ctx context.Context, groupID string) (*OrganizationGroupRoles, *Meta, *http.Response, error)

	CreateRole(ctx context.Context, roleReq *OrganizationRoleReq) (*OrganizationRole, *http.Response, error)
	GetRole(ctx context.Context, roleID string) (*OrganizationRole, *http.Response, error)
	ListRoles(ctx context.Context, options *ListOptions) ([]OrganizationRole, *Meta, *http.Response, error)
	UpdateRole(ctx context.Context, roleID string, roleReq *OrganizationRoleReq) (*OrganizationRole, *http.Response, error)
	DeleteRole(ctx context.Context, roleID string) error
	RestoreRole(ctx context.Context, roleID string) (*OrganizationRole, *http.Response, error)

	CreatePolicy(ctx context.Context, policyReq *OrganizationPolicyReq) (*OrganizationPolicy, *http.Response, error)
	GetPolicy(ctx context.Context, policyID string) (*OrganizationPolicy, *http.Response, error)
	ListPolicies(ctx context.Context, options *ListOptions) ([]OrganizationPolicy, *Meta, *http.Response, error)
	UpdatePolicy(ctx context.Context, policyID string, policyReq *OrganizationPolicyReq) (*OrganizationPolicy, *http.Response, error)
	DeletePolicy(ctx context.Context, policyID string) error
	RestorePolicy(ctx context.Context, policyID string) (*OrganizationPolicy, *http.Response, error)

	ListRolePolicies(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationPolicy, *Meta, *http.Response, error) //nolint:lll
	AttachRolePolicy(ctx context.Context, roleID, policyID string) (*OrganizationRolePolicyAttachment, *http.Response, error)       //nolint:lll
	DetachRolePolicy(ctx context.Context, roleID, policyID string) error

	ListRoleUsers(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationRoleUserAssignment, *Meta, *http.Response, error)
	AttachRoleUser(ctx context.Context, roleID, userID string) (*OrganizationRoleUserAssignment, *http.Response, error)
	DetachRoleUser(ctx context.Context, roleID, userID string) error

	ListRoleGroups(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationRoleGroupAssignment, *Meta, *http.Response, error) //nolint:lll
	AttachRoleGroup(ctx context.Context, roleID, groupID string) (*OrganizationRoleGroupAssignment, *http.Response, error)
	DetachRoleGroup(ctx context.Context, roleID, groupID string) error

	ListPolicyUsers(ctx context.Context, policyID string, options *ListOptions) ([]OrganizationUser, *Meta, *http.Response, error)
	AttachPolicyUser(ctx context.Context, policyID, userID string) error
	DetachPolicyUser(ctx context.Context, policyID, userID string) error

	ListPolicyGroups(ctx context.Context, policyID string, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error)
	AttachPolicyGroup(ctx context.Context, policyID, groupID string) error
	DetachPolicyGroup(ctx context.Context, policyID, groupID string) error

	CreateRoleTrust(ctx context.Context, roleTrustReq *OrganizationRoleTrustCreateReq) (*OrganizationRoleTrust, *http.Response, error)
	GetRoleTrust(ctx context.Context, roleTrustID string) (*OrganizationRoleTrust, *http.Response, error)
	ListRoleTrusts(ctx context.Context, options *ListOptions) ([]OrganizationRoleTrust, *Meta, *http.Response, error)
	UpdateRoleTrust(ctx context.Context, roleTrustID string, roleTrustReq *OrganizationRoleTrustUpdateReq) (*OrganizationRoleTrust, *http.Response, error) //nolint:lll
	DeleteRoleTrust(ctx context.Context, roleTrustID string) error
	RestoreRoleTrust(ctx context.Context, roleTrustID string) (*OrganizationRoleTrust, *http.Response, error)
	ListRoleTrustsByRole(ctx context.Context, roleID string) ([]OrganizationRoleTrust, *Meta, *http.Response, error)
	ListRoleTrustsByUser(ctx context.Context, userID string) ([]OrganizationRoleTrust, *Meta, *http.Response, error)

	CreateRoleSession(ctx context.Context, roleSessionReq *OrganizationRoleSessionReq) (*OrganizationRoleSession, *http.Response, error)
	GetRoleSession(ctx context.Context, token string) (*OrganizationRoleSession, *http.Response, error)
	ListRoleSessions(ctx context.Context, userID string) ([]OrganizationRoleSession, *Meta, *http.Response, error)
	RevokeRoleSession(ctx context.Context, token string) error

	ListUserGroups(ctx context.Context, userID string, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error)
	ListUserRoles(ctx context.Context, userID string, options *ListOptions) (*OrganizationRolesForUser, *Meta, *http.Response, error)
	ListUserPolicies(ctx context.Context, userID string, options *ListOptions) (*OrganizationPoliciesForUser, *Meta, *http.Response, error)

	ListCurrentUserGroups(ctx context.Context, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error)
	ListCurrentUserRoles(ctx context.Context, options *ListOptions) (*OrganizationRolesForUser, *Meta, *http.Response, error)
	ListCurrentUserPolicies(ctx context.Context, options *ListOptions) (*OrganizationPoliciesForUser, *Meta, *http.Response, error)
}

// OrganizationServiceHandler handles interaction with the organization methods for the Vultr API
type OrganizationServiceHandler struct {
	client *Client
}

type organizationsBase struct {
	Organizations []Organization `json:"orgs"`
	Meta          *Meta          `json:"meta"`
}

// Organization represents an organization
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	DateCreated string `json:"date_created"`
}

// OrganizationReq represents an organization modification request
type OrganizationReq struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

// OrganizationInvitation represents an organization invitation
type OrganizationInvitation struct {
	ID               string                            `json:"id"`
	InviterName      string                            `json:"inviter_name"`
	InviterEmail     string                            `json:"inviter_email"`
	EmailInvited     string                            `json:"email_invited"`
	EmailRegistered  string                            `json:"email_registered"`
	Permissions      *OrganizationInvitationPermission `json:"permissions"`
	OrganizationID   string                            `json:"org_id"`
	OrganizationName string                            `json:"org_name"`
	Status           string                            `json:"status"`
	StatusFull       string                            `json:"status_full"`
	StatusInvite     string                            `json:"invite_status"`
	DateCreated      string                            `json:"date_created"`
	DateResponded    string                            `json:"date_responded"`
	DateExpiration   string                            `json:"expiration_date"`
}

type invitationsBase struct {
	Invitations []OrganizationInvitation `json:"invites"`
	Meta        *Meta                    `json:"meta"`
}

// OrganizationInvitationReq represents an organization invitation request
type OrganizationInvitationReq struct {
	Email       string                           `json:"email_invited"`
	Permissions OrganizationInvitationPermission `json:"permissions"`
}

// OrganizationInvitationPermission represents an organization invitation permission
type OrganizationInvitationPermission struct {
	APIEnabled bool `json:"api_enabled"`
}

// OrganizationUser represents an organization user
type OrganizationUser struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Email          string `json:"email"`
	AssignmentType string `json:"assignment_type,omitempty"`
	DateCreated    string `json:"date_created"`
	DateSuspended  string `json:"date_suspended"`
	DateAssigned   string `json:"assigned_date,omitempty"`
}

type organizationUsersBase struct {
	Users []OrganizationUser `json:"users"`
	Meta  *Meta              `json:"meta"`
}

// OrganizationGroup represents an organization group
type OrganizationGroup struct {
	ID           string             `json:"id"`
	Name         string             `json:"display_name"`
	Description  string             `json:"description"`
	Members      []OrganizationUser `json:"members"`
	DateCreated  string             `json:"date_created"`
	DateAssigned string             `json:"assigned_date,omitempty"`
}

type organizationGroupBase struct {
	Group *OrganizationGroup `json:"group"`
}

// OrganizationGroupReq represents an organization group request
type OrganizationGroupReq struct {
	Name        string `json:"display_name"`
	Description string `json:"description"`
}

type organizationGroupsBase struct {
	Groups []OrganizationGroup `json:"groups"`
	Meta   *Meta               `json:"meta"`
}

// OrganizationGroupMemberReq represents an organization group member request
type OrganizationGroupMemberReq struct {
	UserID string `json:"user_id"`
	Name   string `json:"display_name"`
}

// OrganizationRole represents an organization role
type OrganizationRole struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	Type               string               `json:"role_type"`
	MaxSessionDuration int                  `json:"max_session_duration"`
	Policies           []OrganizationPolicy `json:"policies"`
	DateCreated        string               `json:"date_created"`
}

type organizationRoleBase struct {
	Role *OrganizationRole `json:"role"`
}

// OrganizationRolesForUser represents the roles for a user
type OrganizationRolesForUser struct {
	All       []OrganizationRole `json:"all"`
	Direct    []OrganizationRole `json:"direct"`
	Inherited []OrganizationRole `json:"inherited"`
}

type organizationRolesUserBase struct {
	Roles *OrganizationRolesForUser `json:"roles"`
	Meta  *Meta                     `json:"meta"`
}

type organizationRolesBase struct {
	Roles []OrganizationRole `json:"roles"`
	Meta  *Meta              `json:"meta"`
}

// OrganizationRoleReq represents an organization role request
type OrganizationRoleReq struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Type               string `json:"role_type"`
	MaxSessionDuration int    `json:"max_session_duration"`
}

// OrganizationGroupRole represents an organization group role
type OrganizationGroupRole struct {
	ID          string `json:"id"`
	Name        string `json:"display_name"`
	Description string `json:"description"`
	Type        string `json:"role_type"`
	DateCreated string `json:"date_assigned"`
}

type organizationGroupRolesBase struct {
	Roles *OrganizationGroupRoles `json:"roles"`
	Meta  *Meta                   `json:"meta"`
}

// OrganizationGroupRoles represents all roles in a group
type OrganizationGroupRoles struct {
	All       []OrganizationGroupRole `json:"all"`
	Inherited []OrganizationGroupRole `json:"inherited"`
	Direct    []OrganizationGroupRole `json:"direct"`
}

// OrganizationPolicy represents an organization policy
type OrganizationPolicy struct {
	ID           string                     `json:"id"`
	Name         string                     `json:"name"`
	Description  string                     `json:"description"`
	Document     OrganizationPolicyDocument `json:"policy_document"`
	Version      string                     `json:"version"`
	SystemPolicy bool                       `json:"is_system_policy"`
	RoleID       string                     `json:"from_role,omitempty"`
	Source       string                     `json:"source,omitempty"`
	DateCreated  string                     `json:"date_created"`
}

type organizationPoliciesBase struct {
	Policies []OrganizationPolicy `json:"policies"`
	Meta     *Meta                `json:"meta"`
}

type organizationPolicyBase struct {
	Policy *OrganizationPolicy `json:"policy"`
}

// OrganizationRolesForUser represents the roles for a user
type OrganizationPoliciesForUser struct {
	All       []OrganizationPolicy `json:"all"`
	Direct    []OrganizationPolicy `json:"direct"`
	Inherited []OrganizationPolicy `json:"inherited"`
}

type organizationPolicyUserBase struct {
	Policies *OrganizationPoliciesForUser `json:"policies"`
	Meta     *Meta                        `json:"meta"`
}

// OrganizationRolePolicyReq represents an organization role policy attach/detach request
type OrganizationRolePolicyReq struct {
	PolicyID string `json:"policy_id"`
}

// OrganizationRolePolicyAttachment represents a role policy attachment
type OrganizationRolePolicyAttachment struct {
	PolicyID        string `json:"policy_id"`
	PolicyName      string `json:"policy_name"`
	RoleID          string `json:"role_id"`
	RoleDescription string `json:"role_description"`
	RoleType        string `json:"role_type"`
	DateAssigned    string `json:"date_assigned"`
	AssignedBy      string `json:"assigned_by"`
}

type organizationRolePolicyAttachmentsBase struct {
	Policies []OrganizationPolicy `json:"policies"`
	Meta     *Meta                `json:"meta"`
}

// OrganizationPolicyDocument represents a policy document
type OrganizationPolicyDocument struct {
	Version   string                        `json:"Version"`
	Statement []OrganizationPolicyStatement `json:"Statement"`
}

// OrganizationPolicyReq represents an organization policy request
type OrganizationPolicyReq struct {
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	PolicyDocument OrganizationPolicyDocument `json:"policy_document"`
}

// OrganizationPolicyStatement represents an organization policy statement
type OrganizationPolicyStatement struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource []string `json:"Resource"`
}

// UnmarshalJSON is a custom unmarshaller because the Resource and Action
// fields can, in some cases, be a string and not an array. This forces these
// values to an array.
func (s *OrganizationPolicyStatement) UnmarshalJSON(b []byte) error {
	if s == nil {
		return nil
	}

	elems := map[string]interface{}{}
	if err := json.Unmarshal(b, &elems); err != nil {
		return fmt.Errorf("unable to unmarshal organization policy statement : %s", err.Error())
	}

	switch res := elems["Resource"].(type) {
	case []interface{}:
		for i := range res {
			val, ok := res[i].(string)
			if !ok {
				return fmt.Errorf("unable to unmarshal 'Resource' interface slice string value : %v", res[i])
			}
			s.Resource = append(s.Resource, val)
		}
	case string:
		s.Resource = append(s.Resource, res)
	default:
		return fmt.Errorf("unable to unmarshal organization policy statement 'Resource' value")
	}

	switch act := elems["Action"].(type) {
	case []interface{}:
		for i := range act {
			val, ok := act[i].(string)
			if !ok {
				return fmt.Errorf("unable to unmarshal 'Action' interface slice string value : %v", act[i])
			}
			s.Action = append(s.Action, val)
		}
	case string:
		s.Action = append(s.Action, act)
	default:
		return fmt.Errorf("unable to unmarshal organization policy statement 'Action' value")
	}

	effect, ok := elems["Effect"].(string)
	if !ok {
		return fmt.Errorf("unable to unmarshal organization policy statement 'Effect' value")
	}

	s.Effect = effect

	return nil
}

// OrganizationGroupPolicies represents all organization group policies
type OrganizationGroupPolicies struct {
	Direct []OrganizationPolicy `json:"direct"`
	Role   []OrganizationPolicy `json:"from_roles"`
	All    []OrganizationPolicy `json:"all"`
}

type organizationGroupPoliciesBase struct {
	Policies *OrganizationGroupPolicies `json:"policies"`
	Meta     *Meta                      `json:"meta"`
}

// OrganizationRoleUserAssignment represents an organization role user assignment
type OrganizationRoleUserAssignment struct {
	UserID          string `json:"user_id"`
	RoleID          string `json:"role_id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	RoleType        string `json:"role_type"`
	DateCreated     string `json:"date_assigned"`
}

type organizationRoleUserAssignmentsBase struct {
	Assignments []OrganizationRoleUserAssignment `json:"assignments"`
	Meta        *Meta                            `json:"meta"`
}

// OrganizationRoleGroupAssignment represents an organization role group assignment
type OrganizationRoleGroupAssignment struct {
	GroupID         string `json:"group_id"`
	GroupName       string `json:"group_name"`
	RoleID          string `json:"role_id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	RoleType        string `json:"role_type"`
	DateCreated     string `json:"date_assigned"`
}

type organizationRoleGroupAssignmentsBase struct {
	Assignments []OrganizationRoleGroupAssignment `json:"assignments"`
	Meta        *Meta                             `json:"meta"`
}

// OrganizationPolicyUserAssignment represents an organization user policy assignment
type OrganizationPolicyUserAssignment struct {
	UserID            string `json:"user_id"`
	PolicyID          string `json:"policy_id"`
	PolicyName        string `json:"policy_name"`
	PolicyDescription string `json:"policy_description"`
	DateCreated       string `json:"date_assigned"`
}

// OrganizationPolicyGroupAssignment represents an organization group policy assignment
type OrganizationPolicyGroupAssignment struct {
	GroupID           string `json:"group_id"`
	GroupName         string `json:"group_name"`
	PolicyID          string `json:"policy_id"`
	PolicyName        string `json:"policy_name"`
	PolicyDescription string `json:"policy_description"`
	DateCreated       string `json:"date_assigned"`
}

// OrganizationRoleTrust represents an organization role trust
type OrganizationRoleTrust struct {
	ID          string                         `json:"id"`
	RoleName    string                         `json:"role_name"`
	Type        string                         `json:"trust_type"`
	UserID      string                         `json:"trusted_user_id"`
	UserName    string                         `json:"user_display"`
	GroupID     string                         `json:"trusted_group_id"`
	GroupName   string                         `json:"group_display"`
	Conditions  OrganizationRoleTrustCondition `json:"conditions"`
	DateExpires string                         `json:"valid_until"`
	DateCreated string                         `json:"date_created"`
}

// OrganizationRoleTrustCondition represents a organization role trust condition
type OrganizationRoleTrustCondition struct {
	TimeOfDay OrganizationRoleTrustConditionTime `json:"time_of_day"`
	IPRanges  []string                           `json:"ip_address"`
}

// OrganizationRoleTrustConditionTime represents a organization role trust condition time
type OrganizationRoleTrustConditionTime struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type organizationRoleTrustsBase struct {
	RoleTrusts []OrganizationRoleTrust `json:"role_trusts"`
	Meta       *Meta                   `json:"meta"`
}

type organizationRoleTrustBase struct {
	RoleTrust *OrganizationRoleTrust `json:"role_trust"`
}

type organizationRoleTrustsAssumableBase struct {
	RoleTrusts []OrganizationRoleTrust `json:"assumable_roles"`
	Meta       *Meta                   `json:"meta"`
}

// OrganizationRoleTrustCreatReq represents a organization role trust create request
type OrganizationRoleTrustCreateReq struct {
	RoleID      string                             `json:"role_id"`
	UserID      string                             `json:"trusted_user_id,omitempty"`
	GroupID     string                             `json:"trusted_group_id,omitempty"`
	Type        string                             `json:"trust_type"`
	Conditions  *OrganizationRoleTrustReqCondition `json:"conditions,omitempty"`
	DateExpires *string                            `json:"valid_until,omitempty"`
}

// OrganizationRoleTrustReqCondition represents a organization role trust create request condition
type OrganizationRoleTrustReqCondition struct {
	TimeOfDay *OrganizationRoleTrustReqConditionTime `json:"time_of_day,omitempty"`
	IPRanges  []string                               `json:"ip_address,omitempty"`
}

// OrganizationRoleTrustReqConditionTime represents a organization role trust condition time
type OrganizationRoleTrustReqConditionTime struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// OrganizationRoleTrustUpdateReq represents a organization role trust update request
type OrganizationRoleTrustUpdateReq struct {
	RoleID      string                             `json:"role_id,omitempty"`
	UserID      string                             `json:"trusted_user_id,omitempty"`
	GroupID     string                             `json:"trusted_group_id,omitempty"`
	Type        string                             `json:"trust_type,omitempty"`
	Conditions  *OrganizationRoleTrustReqCondition `json:"conditions,omitempty"`
	DateExpires *string                            `json:"valid_until,omitempty"`
}

// OrganizationRoleAssumedReq represents an organization assumed role
type OrganizationRoleSessionReq struct {
	UserID      string                             `json:"user_id"`
	RoleID      string                             `json:"role_id"`
	SessionName string                             `json:"session_name"`
	Duration    int                                `json:"duration,omitempty"`
	Context     *OrganizationRoleSessionReqContext `json:"context,omitempty"`
}

// OrganizationRoleAssumedContext represents an organization assumed role context
type OrganizationRoleSessionReqContext struct {
	IPAddress *string `json:"ip_address,omitempty"`
}

// OrganizationRoleSession represents an organization role session
type OrganizationRoleSession struct {
	Token             string   `json:"session_token"`
	RoleID            string   `json:"role_id"`
	UserID            string   `json:"user_id"`
	SessionName       string   `json:"session_name"`
	AuthMethod        string   `json:"auth_method"`
	RemainingDuration int      `json:"remaining_duration"`
	ConditionsMet     []string `json:"conditions_met"`
	SourceIP          string   `json:"source_ip"`
	DateExpires       string   `json:"expires_at"`
	DateAssumed       string   `json:"assumed_at"`
}

type organizationRoleSessionsBase struct {
	Sessions []OrganizationRoleSession `json:"sessions"`
	Meta     *Meta                     `json:"meta"`
}

// Create an organization
func (o *OrganizationServiceHandler) CreateOrganization(ctx context.Context, organizationReq *OrganizationReq) (*Organization, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, organizationPath, organizationReq)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := o.client.DoWithContext(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, nil
}

// Get an organization
func (o *OrganizationServiceHandler) GetOrganization(ctx context.Context, organizationID string) (*Organization, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", organizationPath, organizationID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := o.client.DoWithContext(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, nil
}

// Update an organization
func (o *OrganizationServiceHandler) UpdateOrganization(ctx context.Context, organizationID string, organizationReq *OrganizationReq) (*Organization, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", organizationPath, organizationID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, organizationReq)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := o.client.DoWithContext(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, nil
}

// Delete an organization
func (o *OrganizationServiceHandler) DeleteOrganization(ctx context.Context, organizationID string) error {
	uri := fmt.Sprintf("%s/%s", organizationPath, organizationID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// Restore an organization
func (o *OrganizationServiceHandler) RestoreOrganization(ctx context.Context, organizationID string) error {
	uri := fmt.Sprintf("%s/%s/restore", organizationPath, organizationID)

	req, err := o.client.NewRequest(ctx, http.MethodPatch, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// List all organizations
func (o *OrganizationServiceHandler) ListOrganizations(ctx context.Context, options *ListOptions) ([]Organization, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, organizationPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	orgs := new(organizationsBase)
	resp, err := o.client.DoWithContext(ctx, req, orgs)
	if err != nil {
		return nil, nil, resp, err
	}

	return orgs.Organizations, orgs.Meta, resp, nil
}

// DeleteUser deletes an organization user
func (o *OrganizationServiceHandler) DeleteUser(ctx context.Context, organizationID, userID string) error {
	uri := fmt.Sprintf("%s/%s/user/%s", organizationPath, organizationID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// SuspendUser suspends an organization user
func (o *OrganizationServiceHandler) SuspendUser(ctx context.Context, organizationID, userID string) error {
	uri := fmt.Sprintf("%s/%s/user/%s/suspend", organizationPath, organizationID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// UnsuspendUser un-suspends an organization user
func (o *OrganizationServiceHandler) UnsuspendUser(ctx context.Context, organizationID, userID string) error {
	uri := fmt.Sprintf("%s/%s/user/%s/unsuspend", organizationPath, organizationID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	return err
}

// ListSuspendedUsers lists suspended organization users
func (o *OrganizationServiceHandler) ListSuspendedUsers(ctx context.Context, organizationID string, options *ListOptions) ([]OrganizationUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/suspended-users", organizationPath, organizationID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	users := new(organizationUsersBase)
	resp, err := o.client.DoWithContext(ctx, req, users)
	if err != nil {
		return nil, nil, resp, err
	}

	return users.Users, users.Meta, resp, err
}

// CreateInvitation creates an organization invitation
func (o *OrganizationServiceHandler) CreateInvitation(ctx context.Context, invitationReq *OrganizationInvitationReq) (*OrganizationInvitation, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, invitationPath, invitationReq)
	if err != nil {
		return nil, nil, err
	}

	invite := new(OrganizationInvitation)
	resp, err := o.client.DoWithContext(ctx, req, invite)
	if err != nil {
		return nil, resp, err
	}

	return invite, resp, nil
}

// GetInvitation will get an invitation
func (o *OrganizationServiceHandler) GetInvitation(ctx context.Context, invitationID string) (*OrganizationInvitation, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", invitationPath, invitationID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	invite := new(OrganizationInvitation)
	resp, err := o.client.DoWithContext(ctx, req, invite)
	if err != nil {
		return nil, resp, err
	}

	return invite, resp, nil
}

// ListInvitations will list all invitations
func (o *OrganizationServiceHandler) ListInvitations(ctx context.Context, options *ListOptions) ([]OrganizationInvitation, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, invitationPath, options)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	invitations := new(invitationsBase)
	resp, err := o.client.DoWithContext(ctx, req, invitations)
	if err != nil {
		return nil, nil, resp, err
	}

	return invitations.Invitations, invitations.Meta, resp, nil
}

// ResendInvitation will resend an invitation
func (o *OrganizationServiceHandler) ResendInvitation(ctx context.Context, invitationID string) (*OrganizationInvitation, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/resend", invitationPath, invitationID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	invite := new(OrganizationInvitation)
	resp, err := o.client.DoWithContext(ctx, req, invite)
	if err != nil {
		return nil, resp, err
	}

	return invite, resp, nil
}

// CreateGroup creates an organization group
func (o *OrganizationServiceHandler) CreateGroup(ctx context.Context, groupReq *OrganizationGroupReq) (*OrganizationGroup, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, groupPath, groupReq)
	if err != nil {
		return nil, nil, err
	}

	group := new(OrganizationGroup)
	resp, err := o.client.DoWithContext(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// GetGroup retrieves an organization group
func (o *OrganizationServiceHandler) GetGroup(ctx context.Context, groupID string) (*OrganizationGroup, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(organizationGroupBase)
	resp, err := o.client.DoWithContext(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group.Group, resp, nil
}

// ListGroups retrieves all organization groups
func (o *OrganizationServiceHandler) ListGroups(ctx context.Context, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, groupPath, options)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	groups := new(organizationGroupsBase)
	resp, err := o.client.DoWithContext(ctx, req, groups)
	if err != nil {
		return nil, nil, resp, err
	}

	return groups.Groups, groups.Meta, resp, nil
}

// UpdateGroup updates an organization group
func (o *OrganizationServiceHandler) UpdateGroup(ctx context.Context, groupID string, groupReq *OrganizationGroupReq) (*OrganizationGroup, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, groupReq)
	if err != nil {
		return nil, nil, err
	}

	group := new(OrganizationGroup)
	resp, err := o.client.DoWithContext(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// DeleteGroup deletes an organization group
func (o *OrganizationServiceHandler) DeleteGroup(ctx context.Context, groupID string) error {
	uri := fmt.Sprintf("%s/%s", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// AddGroupMember adds an organization group member
func (o *OrganizationServiceHandler) AddGroupMember(ctx context.Context, groupID string, memberReq *OrganizationGroupMemberReq) error {
	uri := fmt.Sprintf("%s/%s/members", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, memberReq)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupMember deletes an organization group member
func (o *OrganizationServiceHandler) RemoveGroupMember(ctx context.Context, groupID, userID string) error {
	uri := fmt.Sprintf("%s/%s/members/%s", groupPath, groupID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListGroupPolicies retrieves the organization group policies
func (o *OrganizationServiceHandler) ListGroupPolicies(ctx context.Context, groupID string) (*OrganizationGroupPolicies, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/policies", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	pols := new(organizationGroupPoliciesBase)
	resp, err := o.client.DoWithContext(ctx, req, pols)
	if err != nil {
		return nil, nil, resp, err
	}

	return pols.Policies, pols.Meta, resp, nil
}

// ListGroupRoles retrieves the organization group roles
func (o *OrganizationServiceHandler) ListGroupRoles(ctx context.Context, groupID string) (*OrganizationGroupRoles, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/roles", groupPath, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	roles := new(organizationGroupRolesBase)
	resp, err := o.client.DoWithContext(ctx, req, roles)
	if err != nil {
		return nil, nil, resp, err
	}

	return roles.Roles, roles.Meta, resp, nil
}

// CreateRole creates a role
func (o *OrganizationServiceHandler) CreateRole(ctx context.Context, roleReq *OrganizationRoleReq) (*OrganizationRole, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, rolePath, roleReq)
	if err != nil {
		return nil, nil, err
	}

	role := new(OrganizationRole)
	resp, err := o.client.DoWithContext(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// GetRole retrieves a role
func (o *OrganizationServiceHandler) GetRole(ctx context.Context, roleID string) (*OrganizationRole, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	role := new(organizationRoleBase)
	resp, err := o.client.DoWithContext(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role.Role, resp, nil
}

// ListRoles retrieves a list of all roles
func (o *OrganizationServiceHandler) ListRoles(ctx context.Context, options *ListOptions) ([]OrganizationRole, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, rolePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	roles := new(organizationRolesBase)
	resp, err := o.client.DoWithContext(ctx, req, roles)
	if err != nil {
		return nil, nil, resp, err
	}

	return roles.Roles, roles.Meta, resp, nil
}

// UpdateRole updates a role
func (o *OrganizationServiceHandler) UpdateRole(ctx context.Context, roleID string, roleReq *OrganizationRoleReq) (*OrganizationRole, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, roleReq)
	if err != nil {
		return nil, nil, err
	}

	role := new(OrganizationRole)
	resp, err := o.client.DoWithContext(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// DeleteRole deletes a role
func (o *OrganizationServiceHandler) DeleteRole(ctx context.Context, roleID string) error {
	uri := fmt.Sprintf("%s/%s", rolePath, roleID)
	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// RestoreRole restores a deleted role
func (o *OrganizationServiceHandler) RestoreRole(ctx context.Context, roleID string) (*OrganizationRole, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/restore", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodPatch, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	role := new(OrganizationRole)
	resp, err := o.client.DoWithContext(ctx, req, role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// CreatePolicy creates a policy
func (o *OrganizationServiceHandler) CreatePolicy(ctx context.Context, policyReq *OrganizationPolicyReq) (*OrganizationPolicy, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, policyPath, policyReq)
	if err != nil {
		return nil, nil, err
	}

	policy := new(OrganizationPolicy)
	resp, err := o.client.DoWithContext(ctx, req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// GetPolicy retrieves a policy
func (o *OrganizationServiceHandler) GetPolicy(ctx context.Context, policyID string) (*OrganizationPolicy, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", policyPath, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(organizationPolicyBase)
	resp, err := o.client.DoWithContext(ctx, req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy.Policy, resp, nil
}

// ListPolicies retrieves a list of all policys
func (o *OrganizationServiceHandler) ListPolicies(ctx context.Context, options *ListOptions) ([]OrganizationPolicy, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, policyPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	policies := new(organizationPoliciesBase)
	resp, err := o.client.DoWithContext(ctx, req, policies)
	if err != nil {
		return nil, nil, resp, err
	}

	return policies.Policies, policies.Meta, resp, nil
}

// UpdatePolicy updates a policy
func (o *OrganizationServiceHandler) UpdatePolicy(ctx context.Context, policyID string, policyReq *OrganizationPolicyReq) (*OrganizationPolicy, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", policyPath, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, policyReq)
	if err != nil {
		return nil, nil, err
	}

	policy := new(OrganizationPolicy)
	resp, err := o.client.DoWithContext(ctx, req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeletePolicy deletes a policy
func (o *OrganizationServiceHandler) DeletePolicy(ctx context.Context, policyID string) error {
	uri := fmt.Sprintf("%s/%s", policyPath, policyID)
	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// RestorePolicy restores a deleted policy
func (o *OrganizationServiceHandler) RestorePolicy(ctx context.Context, policyID string) (*OrganizationPolicy, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/restore", policyPath, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodPatch, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(OrganizationPolicy)
	resp, err := o.client.DoWithContext(ctx, req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// ListRolePolicies retrieves a list of role policies
func (o *OrganizationServiceHandler) ListRolePolicies(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationPolicy, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/policies", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	policies := new(organizationRolePolicyAttachmentsBase)
	resp, err := o.client.DoWithContext(ctx, req, policies)
	if err != nil {
		return nil, nil, resp, err
	}

	return policies.Policies, policies.Meta, resp, nil
}

// AttachRolePolicy attaches a role policy
func (o *OrganizationServiceHandler) AttachRolePolicy(ctx context.Context, roleID, policyID string) (*OrganizationRolePolicyAttachment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/policies/%s", rolePath, roleID, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	att := new(OrganizationRolePolicyAttachment)
	resp, err := o.client.DoWithContext(ctx, req, att)
	if err != nil {
		return nil, resp, err
	}

	return att, resp, nil
}

// DetachRolePolicy detaches a role policy
func (o *OrganizationServiceHandler) DetachRolePolicy(ctx context.Context, roleID, policyID string) error {
	uri := fmt.Sprintf("%s/%s/policies/%s", rolePath, roleID, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err := o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// ListRoleUsers lists a role's users
func (o *OrganizationServiceHandler) ListRoleUsers(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationRoleUserAssignment, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/users", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	users := new(organizationRoleUserAssignmentsBase)
	resp, err := o.client.DoWithContext(ctx, req, users)
	if err != nil {
		return nil, nil, resp, err
	}

	return users.Assignments, users.Meta, resp, nil
}

// AttachRoleUser attaches a user to a role
func (o *OrganizationServiceHandler) AttachRoleUser(ctx context.Context, roleID, userID string) (*OrganizationRoleUserAssignment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/users/%s", rolePath, roleID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(OrganizationRoleUserAssignment)
	resp, err := o.client.DoWithContext(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// DetachRoleUser detaches a user from a role
func (o *OrganizationServiceHandler) DetachRoleUser(ctx context.Context, roleID, userID string) error {
	uri := fmt.Sprintf("%s/%s/users/%s", rolePath, roleID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err = o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// ListRoleGroups lists a role's groups
func (o *OrganizationServiceHandler) ListRoleGroups(ctx context.Context, roleID string, options *ListOptions) ([]OrganizationRoleGroupAssignment, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/groups", rolePath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	groups := new(organizationRoleGroupAssignmentsBase)
	resp, err := o.client.DoWithContext(ctx, req, groups)
	if err != nil {
		return nil, nil, resp, err
	}

	return groups.Assignments, groups.Meta, resp, nil
}

// AttachRoleGroup attaches a group to a role
func (o *OrganizationServiceHandler) AttachRoleGroup(ctx context.Context, roleID, groupID string) (*OrganizationRoleGroupAssignment, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/groups/%s", rolePath, roleID, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(OrganizationRoleGroupAssignment)
	resp, err := o.client.DoWithContext(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// DetachRoleGroup detaches a group from a role
func (o *OrganizationServiceHandler) DetachRoleGroup(ctx context.Context, roleID, groupID string) error {
	uri := fmt.Sprintf("%s/%s/groups/%s", rolePath, roleID, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err = o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// ListPolicyUsers lists the policy's users
func (o *OrganizationServiceHandler) ListPolicyUsers(ctx context.Context, policyID string, options *ListOptions) ([]OrganizationUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/users", policyPath, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	users := new(organizationUsersBase)
	resp, err := o.client.DoWithContext(ctx, req, users)
	if err != nil {
		return nil, nil, resp, err
	}

	return users.Users, users.Meta, resp, nil
}

// AttachPolicyUser attaches a policy to a user
func (o *OrganizationServiceHandler) AttachPolicyUser(ctx context.Context, policyID, userID string) error {
	uri := fmt.Sprintf("%s/%s/users/%s", policyPath, policyID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DetachPolicyUser detaches a user from a policy
func (o *OrganizationServiceHandler) DetachPolicyUser(ctx context.Context, policyID, userID string) error {
	uri := fmt.Sprintf("%s/%s/users/%s", policyPath, policyID, userID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListPolicyGroups lists the policy's groups
func (o *OrganizationServiceHandler) ListPolicyGroups(ctx context.Context, policyID string, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/groups", policyPath, policyID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	groups := new(organizationGroupsBase)
	resp, err := o.client.DoWithContext(ctx, req, groups)
	if err != nil {
		return nil, nil, resp, err
	}

	return groups.Groups, groups.Meta, resp, nil
}

// AttachPolicyGroup attaches a policy to a group
func (o *OrganizationServiceHandler) AttachPolicyGroup(ctx context.Context, policyID, groupID string) error {
	uri := fmt.Sprintf("%s/%s/groups/%s", policyPath, policyID, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DetachPolicyGroup detaches a group from a policy
func (o *OrganizationServiceHandler) DetachPolicyGroup(ctx context.Context, policyID, groupID string) error {
	uri := fmt.Sprintf("%s/%s/groups/%s", policyPath, policyID, groupID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateRoleTrust creates a role trust
func (o *OrganizationServiceHandler) CreateRoleTrust(ctx context.Context, roleTrustReq *OrganizationRoleTrustCreateReq) (*OrganizationRoleTrust, *http.Response, error) { //nolint:lll
	req, err := o.client.NewRequest(ctx, http.MethodPost, roleTrustPath, roleTrustReq)
	if err != nil {
		return nil, nil, err
	}

	roleTrust := new(OrganizationRoleTrust)
	resp, err := o.client.DoWithContext(ctx, req, roleTrust)
	if err != nil {
		return nil, resp, err
	}

	return roleTrust, resp, nil
}

// GetRoleTrust retrieves a role trust
func (o *OrganizationServiceHandler) GetRoleTrust(ctx context.Context, roleTrustID string) (*OrganizationRoleTrust, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", roleTrustPath, roleTrustID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	roleTrust := new(organizationRoleTrustBase)
	resp, err := o.client.DoWithContext(ctx, req, roleTrust)
	if err != nil {
		return nil, resp, err
	}

	return roleTrust.RoleTrust, resp, nil
}

// ListRoleTrusts retrieves a list of roll trusts
func (o *OrganizationServiceHandler) ListRoleTrusts(ctx context.Context, options *ListOptions) ([]OrganizationRoleTrust, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := o.client.NewRequest(ctx, http.MethodGet, roleTrustPath, options)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	roleTrusts := new(organizationRoleTrustsBase)
	resp, err := o.client.DoWithContext(ctx, req, roleTrusts)
	if err != nil {
		return nil, nil, resp, err
	}

	return roleTrusts.RoleTrusts, roleTrusts.Meta, resp, nil
}

// UpdateRoleTrust updates a role trust
func (o *OrganizationServiceHandler) UpdateRoleTrust(ctx context.Context, roleTrustID string, roleTrustReq *OrganizationRoleTrustUpdateReq) (*OrganizationRoleTrust, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s", roleTrustPath, roleTrustID)

	req, err := o.client.NewRequest(ctx, http.MethodPut, uri, roleTrustReq)
	if err != nil {
		return nil, nil, err
	}

	roleTrust := new(OrganizationRoleTrust)
	resp, err := o.client.DoWithContext(ctx, req, roleTrust)
	if err != nil {
		return nil, resp, err
	}

	return roleTrust, resp, nil
}

// DeleteRoleTrust deletes a role trust
func (o *OrganizationServiceHandler) DeleteRoleTrust(ctx context.Context, roleTrustID string) error {
	uri := fmt.Sprintf("%s/%s", roleTrustPath, roleTrustID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// RestoreRoleTrust restores a deleted role trust
func (o *OrganizationServiceHandler) RestoreRoleTrust(ctx context.Context, roleTrustID string) (*OrganizationRoleTrust, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/restore", roleTrustPath, roleTrustID)

	req, err := o.client.NewRequest(ctx, http.MethodPatch, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	roleTrust := new(organizationRoleTrustBase)
	resp, err := o.client.DoWithContext(ctx, req, roleTrust)
	if err != nil {
		return nil, resp, err
	}

	return roleTrust.RoleTrust, resp, nil
}

// ListRoleTrustsByRole retrieves a list of all role trusts for the specified role
func (o *OrganizationServiceHandler) ListRoleTrustsByRole(ctx context.Context, roleID string) ([]OrganizationRoleTrust, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/by-role/%s", roleTrustPath, roleID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	roleTrusts := new(organizationRoleTrustsBase)
	resp, err := o.client.DoWithContext(ctx, req, roleTrusts)
	if err != nil {
		return nil, nil, resp, err
	}

	return roleTrusts.RoleTrusts, roleTrusts.Meta, resp, nil
}

// ListRoleTrustsByUser retrieves a list of all role trusts for the specified user
func (o *OrganizationServiceHandler) ListRoleTrustsByUser(ctx context.Context, userID string) ([]OrganizationRoleTrust, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/assumable/%s", roleTrustPath, userID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	roleTrusts := new(organizationRoleTrustsAssumableBase)
	resp, err := o.client.DoWithContext(ctx, req, roleTrusts)
	if err != nil {
		return nil, nil, resp, err
	}

	return roleTrusts.RoleTrusts, roleTrusts.Meta, resp, nil
}

// CreateRoleSession creates a temporary role session
func (o *OrganizationServiceHandler) CreateRoleSession(ctx context.Context, roleSessionReq *OrganizationRoleSessionReq) (*OrganizationRoleSession, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/assume", roleSessionPath)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, roleSessionReq)
	if err != nil {
		return nil, nil, err
	}

	session := new(OrganizationRoleSession)
	resp, err := o.client.DoWithContext(ctx, req, session)
	if err != nil {
		return nil, resp, err
	}

	return session, resp, nil
}

// GetRoleSession retrieves details on a temporary role session
func (o *OrganizationServiceHandler) GetRoleSession(ctx context.Context, token string) (*OrganizationRoleSession, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", roleSessionPath, token)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	roleSession := new(OrganizationRoleSession)
	resp, err := o.client.DoWithContext(ctx, req, roleSession)
	if err != nil {
		return nil, resp, err
	}

	return roleSession, resp, nil
}

// ListRoleSessions retrieves a list of a user's temporary sessions
func (o *OrganizationServiceHandler) ListRoleSessions(ctx context.Context, userID string) ([]OrganizationRoleSession, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/users/%s/sessions", roleSessionPath, userID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	sessions := new(organizationRoleSessionsBase)
	resp, err := o.client.DoWithContext(ctx, req, sessions)
	if err != nil {
		return nil, nil, resp, err
	}

	return sessions.Sessions, sessions.Meta, resp, nil
}

// RevokeRoleSession revokes a temporary role session
func (o *OrganizationServiceHandler) RevokeRoleSession(ctx context.Context, token string) error {
	uri := fmt.Sprintf("%s/%s", roleSessionPath, token)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = o.client.DoWithContext(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListUserGroups retrieves a list of groups for a user
func (o *OrganizationServiceHandler) ListUserGroups(ctx context.Context, userID string, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/groups", userPath, userID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	groups := new(organizationGroupsBase)
	resp, err := o.client.DoWithContext(ctx, req, groups)
	if err != nil {
		return nil, nil, resp, err
	}

	return groups.Groups, groups.Meta, resp, nil
}

// ListUserRoles retrieves a list of roles for a user
func (o *OrganizationServiceHandler) ListUserRoles(ctx context.Context, userID string, options *ListOptions) (*OrganizationRolesForUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/roles", userPath, userID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	roles := new(organizationRolesUserBase)
	resp, err := o.client.DoWithContext(ctx, req, roles)
	if err != nil {
		return nil, nil, resp, err
	}

	return roles.Roles, roles.Meta, resp, nil
}

// ListUserPolicies retrieves a list of policies for a user
func (o *OrganizationServiceHandler) ListUserPolicies(ctx context.Context, userID string, options *ListOptions) (*OrganizationPoliciesForUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/policies", userPath, userID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	policies := new(organizationPolicyUserBase)
	resp, err := o.client.DoWithContext(ctx, req, policies)
	if err != nil {
		return nil, nil, resp, err
	}

	return policies.Policies, policies.Meta, resp, nil
}

// ListCurrentUserGroups retrieves a list of groups for the current user
func (o *OrganizationServiceHandler) ListCurrentUserGroups(ctx context.Context, options *ListOptions) ([]OrganizationGroup, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/me/groups", userPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	groups := new(organizationGroupsBase)
	resp, err := o.client.DoWithContext(ctx, req, groups)
	if err != nil {
		return nil, nil, resp, err
	}

	return groups.Groups, groups.Meta, resp, nil
}

// ListCurrentUserRoles retrieves a list of roles for the current user
func (o *OrganizationServiceHandler) ListCurrentUserRoles(ctx context.Context, options *ListOptions) (*OrganizationRolesForUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/me/roles", userPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	roles := new(organizationRolesUserBase)
	resp, err := o.client.DoWithContext(ctx, req, roles)
	if err != nil {
		return nil, nil, resp, err
	}

	return roles.Roles, roles.Meta, resp, nil
}

// ListCurrentUserPolicies retrieves a list of policies for the current user
func (o *OrganizationServiceHandler) ListCurrentUserPolicies(ctx context.Context, options *ListOptions) (*OrganizationPoliciesForUser, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/me/policies", userPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	policies := new(organizationPolicyUserBase)
	resp, err := o.client.DoWithContext(ctx, req, policies)
	if err != nil {
		return nil, nil, resp, err
	}

	return policies.Policies, policies.Meta, resp, nil
}
