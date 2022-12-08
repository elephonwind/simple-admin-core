// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package coreclient

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiInfo                    = core.ApiInfo
	ApiListResp                = core.ApiListResp
	ApiPageReq                 = core.ApiPageReq
	BaseResp                   = core.BaseResp
	CallbackReq                = core.CallbackReq
	ChangePasswordReq          = core.ChangePasswordReq
	CreateOrUpdateMenuParamReq = core.CreateOrUpdateMenuParamReq
	CreateOrUpdateMenuReq      = core.CreateOrUpdateMenuReq
	CreateOrUpdateUserReq      = core.CreateOrUpdateUserReq
	CreatePolicyReq            = core.CreatePolicyReq
	DictionaryDetail           = core.DictionaryDetail
	DictionaryDetailList       = core.DictionaryDetailList
	DictionaryDetailReq        = core.DictionaryDetailReq
	DictionaryInfo             = core.DictionaryInfo
	DictionaryList             = core.DictionaryList
	DictionaryPageReq          = core.DictionaryPageReq
	Empty                      = core.Empty
	GetUserListReq             = core.GetUserListReq
	IDReq                      = core.IDReq
	IDsReq                     = core.IDsReq
	LoginReq                   = core.LoginReq
	LoginResp                  = core.LoginResp
	MenuInfo                   = core.MenuInfo
	MenuInfoList               = core.MenuInfoList
	MenuParamListResp          = core.MenuParamListResp
	MenuParamResp              = core.MenuParamResp
	MenuRoleInfo               = core.MenuRoleInfo
	MenuRoleListResp           = core.MenuRoleListResp
	Meta                       = core.Meta
	OauthLoginReq              = core.OauthLoginReq
	OauthRedirectResp          = core.OauthRedirectResp
	PageInfoReq                = core.PageInfoReq
	PolicyPartInfo             = core.PolicyPartInfo
	ProviderInfo               = core.ProviderInfo
	ProviderListResp           = core.ProviderListResp
	RoleInfo                   = core.RoleInfo
	RoleListResp               = core.RoleListResp
	RoleMenuAuthorityReq       = core.RoleMenuAuthorityReq
	RoleMenuAuthorityResp      = core.RoleMenuAuthorityResp
	StatusCodeReq              = core.StatusCodeReq
	TokenInfo                  = core.TokenInfo
	TokenListReq               = core.TokenListReq
	TokenListResp              = core.TokenListResp
	UUIDReq                    = core.UUIDReq
	UpdatePolicyReq            = core.UpdatePolicyReq
	UpdateProfileReq           = core.UpdateProfileReq
	UserInfoResp               = core.UserInfoResp
	UserListResp               = core.UserListResp

	Core interface {
		// init
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// user service
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseResp, error)
		CreateOrUpdateUser(ctx context.Context, in *CreateOrUpdateUserReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateUserStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error)
		// menu service
		CreateOrUpdateMenu(ctx context.Context, in *CreateOrUpdateMenuReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuListByRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfoList, error)
		GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error)
		CreateOrUpdateMenuParam(ctx context.Context, in *CreateOrUpdateMenuParamReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteMenuParam(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuParamListByMenuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuParamListResp, error)
		// role service
		CreateOrUpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error)
		GetRoleList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*RoleListResp, error)
		UpdateRoleStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error)
		// api management service
		CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error)
		// authorization management service
		GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error)
		CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error)
		// dictionary management service
		CreateOrUpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteDictionary(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryList(ctx context.Context, in *DictionaryPageReq, opts ...grpc.CallOption) (*DictionaryList, error)
		GetDetailByDictionaryName(ctx context.Context, in *DictionaryDetailReq, opts ...grpc.CallOption) (*DictionaryDetailList, error)
		CreateOrUpdateDictionaryDetail(ctx context.Context, in *DictionaryDetail, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteDictionaryDetail(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		// oauth management
		CreateOrUpdateProvider(ctx context.Context, in *ProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteProvider(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetProviderList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*ProviderListResp, error)
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error)
		OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*LoginResp, error)
		// Token management
		CreateOrUpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteToken(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		BatchDeleteToken(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error)
		UpdateTokenStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error)
		BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultCore struct {
		cli zrpc.Client
	}
)

func NewCore(cli zrpc.Client) Core {
	return &defaultCore{
		cli: cli,
	}
}

// init
func (m *defaultCore) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// user service
func (m *defaultCore) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultCore) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

func (m *defaultCore) CreateOrUpdateUser(ctx context.Context, in *CreateOrUpdateUserReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateUser(ctx, in, opts...)
}

func (m *defaultCore) GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultCore) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultCore) DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultCore) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateProfile(ctx, in, opts...)
}

func (m *defaultCore) UpdateUserStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

// menu service
func (m *defaultCore) CreateOrUpdateMenu(ctx context.Context, in *CreateOrUpdateMenuReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateMenu(ctx, in, opts...)
}

func (m *defaultCore) DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

func (m *defaultCore) GetMenuListByRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuListByRole(ctx, in, opts...)
}

func (m *defaultCore) GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuList(ctx, in, opts...)
}

func (m *defaultCore) CreateOrUpdateMenuParam(ctx context.Context, in *CreateOrUpdateMenuParamReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateMenuParam(ctx, in, opts...)
}

func (m *defaultCore) DeleteMenuParam(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteMenuParam(ctx, in, opts...)
}

func (m *defaultCore) GetMenuParamListByMenuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuParamListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuParamListByMenuId(ctx, in, opts...)
}

// role service
func (m *defaultCore) CreateOrUpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateRole(ctx, in, opts...)
}

func (m *defaultCore) DeleteRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultCore) GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetRoleById(ctx, in, opts...)
}

func (m *defaultCore) GetRoleList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

func (m *defaultCore) UpdateRoleStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateRoleStatus(ctx, in, opts...)
}

// api management service
func (m *defaultCore) CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateApi(ctx, in, opts...)
}

func (m *defaultCore) DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

func (m *defaultCore) GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetApiList(ctx, in, opts...)
}

// authorization management service
func (m *defaultCore) GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuAuthority(ctx, in, opts...)
}

func (m *defaultCore) CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateMenuAuthority(ctx, in, opts...)
}

// dictionary management service
func (m *defaultCore) CreateOrUpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateDictionary(ctx, in, opts...)
}

func (m *defaultCore) DeleteDictionary(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteDictionary(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryList(ctx context.Context, in *DictionaryPageReq, opts ...grpc.CallOption) (*DictionaryList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryList(ctx, in, opts...)
}

func (m *defaultCore) GetDetailByDictionaryName(ctx context.Context, in *DictionaryDetailReq, opts ...grpc.CallOption) (*DictionaryDetailList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDetailByDictionaryName(ctx, in, opts...)
}

func (m *defaultCore) CreateOrUpdateDictionaryDetail(ctx context.Context, in *DictionaryDetail, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateDictionaryDetail(ctx, in, opts...)
}

func (m *defaultCore) DeleteDictionaryDetail(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteDictionaryDetail(ctx, in, opts...)
}

// oauth management
func (m *defaultCore) CreateOrUpdateProvider(ctx context.Context, in *ProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateProvider(ctx, in, opts...)
}

func (m *defaultCore) DeleteProvider(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteProvider(ctx, in, opts...)
}

func (m *defaultCore) GetProviderList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*ProviderListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetProviderList(ctx, in, opts...)
}

func (m *defaultCore) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

func (m *defaultCore) OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.OauthCallback(ctx, in, opts...)
}

// Token management
func (m *defaultCore) CreateOrUpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateToken(ctx, in, opts...)
}

func (m *defaultCore) DeleteToken(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteToken(ctx, in, opts...)
}

func (m *defaultCore) BatchDeleteToken(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.BatchDeleteToken(ctx, in, opts...)
}

func (m *defaultCore) GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetTokenList(ctx, in, opts...)
}

func (m *defaultCore) UpdateTokenStatus(ctx context.Context, in *StatusCodeReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateTokenStatus(ctx, in, opts...)
}

func (m *defaultCore) BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.BlockUserAllToken(ctx, in, opts...)
}
