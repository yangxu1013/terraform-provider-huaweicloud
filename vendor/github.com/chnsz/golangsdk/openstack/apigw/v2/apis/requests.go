package apis

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"
)

// APIOpts is a struct which will be used to create a new API or update an existing API.
type APIOpts struct {
	// ID of the API group to which the API belongs.
	GroupId string `json:"group_id" required:"true"`
	// API name, which can contain 3 to 64 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name string `json:"name" required:"true"`
	// API type. The valid types are as following:
	//   1: public API
	//   2: private API
	Type int `json:"type" required:"true"`
	// Request protocol. The valid protocols are as following:
	//   HTTP.
	//   HTTPS (default).
	//   BOTH: The API can be accessed through both HTTP and HTTPS.
	ReqProtocol string `json:"req_protocol" required:"true"`
	// Request method. The valid values are GET,  POST,  PUT,  DELETE, HEAD, PATCH, OPTIONS and ANY.
	ReqMethod string `json:"req_method" required:"true"`
	// Request address, which can contain a maximum of 512 characters request parameters enclosed with brackets ({}).
	// For example, /getUserInfo/{userId}.
	// The request address can contain special characters, such as asterisks (), percent signs (%), hyphens (-), and
	// underscores (_) and must comply with URI specifications.
	// The address can contain environment variables, each starting with a letter and consisting of 3 to 32 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.
	ReqURI string `json:"req_uri" required:"true"`
	// Security authentication mode. The valid modes are as following:
	//   NONE
	//   APP
	//   IAM
	//   AUTHORIZER
	AuthType string `json:"auth_type" required:"true"`
	// Backend type. The valid types are as following:
	//   HTTP: web backend.
	//   FUNCTION: FunctionGraph backend.
	//   MOCK: Mock backend.
	BackendType string `json:"backend_type" required:"true"`
	// API version. The maximum length of version string is 16.
	Version *string `json:"version,omitempty"`
	// Security authentication parameter.
	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`
	// Indicates whether CORS is supported. The valid values are as following:
	//   TRUE: supported.
	//   FALSE: not supported (default).
	Cors *bool `json:"cors,omitempty"`
	// Route matching mode.  The valid modes are as following:
	//   SWA: prefix match
	//   NORMAL: exact match (default).
	MatchMode string `json:"match_mode,omitempty"`
	// Description of the API, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
	// API request body, which can be an example request body, media type, or parameters.
	// Ensure that the request body does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	BodyDescription *string `json:"body_remark,omitempty"`
	// Example response for a successful request. Ensure that the response does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	ResultNormalSample *string `json:"result_normal_sample,omitempty"`
	// Example response for a failed request. Ensure that the response does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	ResultFailureSample *string `json:"result_failure_sample,omitempty"`
	// ID of the frontend custom authorizer.
	AuthorizerId string `json:"authorizer_id,omitempty"`
	// List of tags. The length of the tags list is range from 1 to 128.
	// The value can contain only letters, digits, and underscores (_), and must start with a letter.
	Tags []string `json:"tags,omitempty"`
	// Group response ID.
	ResponseId string `json:"response_id,omitempty"`
	// Request parameters.
	ReqParams []ReqParamBase `json:"req_params,omitempty"`
	// Backend parameters.
	BackendParams []BackendParamBase `json:"backend_params,omitempty"`
	// Mock backend details.
	MockInfo *Mock `json:"mock_info,omitempty"`
	// FunctionGraph backend details.
	FuncInfo *FuncGraph `json:"func_info,omitempty"`
	// Web backend details.
	WebInfo *Web `json:"backend_api,omitempty"`
	// Mock policy backends.
	PolicyMocks []PolicyMock `json:"policy_mocks,omitempty"`
	// FunctionGraph policy backends.
	PolicyFunctions []PolicyFuncGraph `json:"policy_functions,omitempty"`
	// Web policy backends.
	PolicyWebs []PolicyWeb `json:"policy_https,omitempty"`
}

// AuthOpt is an object which will be build up an APIG application authorization.
type AuthOpt struct {
	// Indicates whether AppCode authentication is enabled. The valid types are as following:
	//   DISABLE: AppCode authentication is disabled (default).
	//   HEADER: AppCode authentication is enabled and the AppCode is located in the header.
	// This parameter is valid only if auth_type is set to App.
	AppCodeAuthType string `json:"app_code_auth_type,omitempty"`
}

// Mock is an object which will be build up a mock backend.
type Mock struct {
	// Description about the backend, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
	// Response.
	ResultContent *string `json:"result_content,omitempty"`
	// Function version Ensure that the version does not exceed 64 characters.
	Version string `json:"version,omitempty"`
	// Backend custom authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

// FuncGraph is an object which will be build up a function graph backend.
type FuncGraph struct {
	// Function URN.
	FunctionUrn string `json:"function_urn" required:"true"`
	// Invocation mode. The valid modes are as following:
	//   async: asynchronous
	//   sync: synchronous
	InvocationType string `json:"invocation_type" required:"true"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout" required:"true"`
	// Backend custom authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// Description about the backend, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
	// Function version.
	// Maximum: 64
	Version string `json:"version,omitempty"`
}

// Web is an object which will be build up a http backend.
type Web struct {
	// Request method. The valid methods are GET, POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	ReqMethod string `json:"req_method" required:"true"`
	// Request protocol. The valid protocols are HTTP and HTTPS
	ReqProtocol string `json:"req_protocol" required:"true"`
	// Request address, which can contain a maximum of 512 characters request parameters enclosed with brackets ({}).
	// For example, /getUserInfo/{userId}.
	// The request address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-), and
	// underscores (_) and must comply with URI specifications.
	// The address can contain environment variables, each starting with a letter and consisting of 3 to 32 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.
	ReqURI string `json:"req_uri" required:"true"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout" required:"true"`
	// Backend custom authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// Backend service address which consists of a domain name or IP address and a port number, with not more than 255
	// characters. It must be in the format "Host name:Port number", for example, apig.example.com:7443.
	// If the port number is not specified, the default HTTPS port 443 or the default HTTP port 80 is used.
	// The backend service address can contain environment variables, each starting with a letter and consisting of
	// 3 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed.
	DomainURL string `json:"url_domain,omitempty"`
	// Description, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
	// Web backend version, which can contain a maximum of 16 characters.
	Version *string `json:"version,omitempty"`
	// Indicates whether to enable two-way authentication.
	ClientSslEnable *bool `json:"enable_client_ssl,omitempty"`
	// VPC channel details. This parameter is required if vpc_channel_status is set to 1.
	VpcChannelInfo *VpcChannel `json:"vpc_channel_info,omitempty"`
	// Indicates whether to use a VPC channel. The valid values are as following:
	//   1: (A VPC channel is used).
	//   2: (No VPC channel is used).
	VpcChannelStatus int `json:"vpc_channel_status,omitempty"`
}

// VpcChannel is an object which will be build up a vpc channel.
type VpcChannel struct {
	// VPC channel ID.
	VpcChannelId string `json:"vpc_channel_id" required:"true"`
	// Proxy host.
	VpcChannelProxyHost string `json:"vpc_channel_proxy_host,omitempty"`
}

// ReqParamBase is an object which will be build up a front-end request parameter.
type ReqParamBase struct {
	// The parameter name, which contain of 1 to 32 characters and start with a letter.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Parameter type. The valid types are as following:
	//   STRING
	//   NUMBER
	Type string `json:"type" required:"true"`
	// Parameter location. The valid modes are as following:
	//   PATH
	//   QUERY
	//   HEADER
	Location string `json:"location" required:"true"`
	// Default value.
	DefaultValue *string `json:"default_value,omitempty"`
	// Example value.
	SampleValue string `json:"sample_value,omitempty"`
	// Indicates whether the parameter is required. The valid values are 1 (yes) and 2 (no).
	// The value of this parameter is 1 if Location is set to PATH, and 2 if Location is set to another value.
	Required int `json:"required,omitempty"`
	// Indicates whether validity check is enabled. The valid modes are as following:
	//   1: enabled.
	//   2: disabled (default).
	ValidEnable int `json:"valid_enable,omitempty"`
	// Description about the backend, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
	// Enumerated value.
	Enumerations *string `json:"enumerations,omitempty"`
	// Minimum value.
	// This parameter is valid when type is set to NUMBER.
	MinNum *int `json:"min_num,omitempty"`
	// Maximum value.
	// This parameter is valid when type is set to NUMBER.
	MaxNum *int `json:"max_num,omitempty"`
	// Minimum length.
	// This parameter is valid when type is set to STRING.
	MinSize *int `json:"min_size,omitempty"`
	// Maximum length.
	// This parameter is valid when type is set to STRING.
	MaxSize *int `json:"max_size,omitempty"`
	// Indicates whether to transparently transfer the parameter. The valid values are 1 (yes) and 2 (no).
	PassThrough string `json:"pass_through,omitempty"`
}

// PolicyMock is an object which will be build up a backend policy of the mock.
type PolicyMock struct {
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Effective mode of the backend policy. The valid modes are as following:
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Backend name, which consists of 3 to 64 characters and must start with a letter and can contain letters, digits,
	// and underscores (_).
	Name string `json:"name" required:"true"`
	// Authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// Backend parameters.
	BackendParams []BackendParamBase `json:"backend_params,omitempty"`
	// Response.
	ResultContent string `json:"result_content,omitempty"`
}

// PolicyFuncGraph is an object which will be build up a backend policy of the function graph.
type PolicyFuncGraph struct {
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Effective mode of the backend policy.
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Function URN.
	FunctionUrn string `json:"function_urn" required:"true"`
	// Invocation mode. The valid modes are as following:
	//   async: asynchronous
	//   sync: synchronous
	InvocationType string `json:"invocation_type" required:"true"`
	// The backend name consists of 3 to 64 characters, which must start with a letter and can contain letters, digits,
	// and underscores (_).
	Name string `json:"name" required:"true"`
	// Authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// Backend parameters.
	BackendParams []BackendParamBase `json:"backend_params,omitempty"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout,omitempty"`
	// Function version Ensure that the version does not exceed 64 characters.
	Version string `json:"version,omitempty"`
}

// PolicyWeb is an object which will be build up a backend policy of the http.
type PolicyWeb struct {
	// Request protocol. The value can be HTTP or HTTPS.
	ReqProtocol string `json:"req_protocol" required:"true"`
	// Request method. The valid methods are GET, POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	ReqMethod string `json:"req_method" required:"true"`
	// Request address, which can contain request parameters enclosed with brackets ({}).
	// For example, /getUserInfo/{userId}. The request address can contain special characters, such as asterisks (),
	// percent signs (%), hyphens (-), and underscores (_). It can contain a maximum of 512 characters and must comply
	// with URI specifications.
	// The request address can contain environment variables, each starting with a letter and consisting of 3 to 32
	// characters. Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.
	// The request address must comply with URI specifications.
	ReqURI string `json:"req_uri" required:"true"`
	// Effective mode of the backend policy. The valid modes are as following:
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Backend name, which contains of 3 to 64, must start with a letter and can contain letters, digits, and
	// underscores (_).
	Name string `json:"name" required:"true"`
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Backend parameters.
	BackendParams []BackendParamBase `json:"backend_params,omitempty"`
	// Authorizer ID.
	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// VPC channel details. This parameter is required if vpc_channel_status is set to 1.
	VpcChannelInfo *VpcChannel `json:"vpc_channel_info,omitempty"`
	// Indicates whether to use a VPC channel. The valid value are as following:
	//   1: A VPC channel is used.
	//   2: No VPC channel is used.
	VpcChannelStatus int `json:"vpc_channel_status,omitempty"`
	// Endpoint of the policy backend.
	// An endpoint consists of a domain name or IP address and a port number, with not more than 255 characters.
	// It must be in the format "Domain name:Port number", for example, apig.example.com:7443.
	// If the port number is not specified, the default HTTPS port 443 or the default HTTP port 80 is used.
	// The endpoint can contain environment variables, each starting with letter and consisting of 3 to 32 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	DomainURL string `json:"url_domain,omitempty"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout,omitempty"`
}

// BackendParamBase is an object which will be build up a back-end parameter.
type BackendParamBase struct {
	// Parameter type. The valid types are as following:
	//   REQUEST: Backend parameter.
	//   CONSTANT: Constant parameter.
	//   SYSTEM: System parameter.
	Origin string `json:"origin" required:"true"`
	// Parameter name, which can contains 1 to 32 characters, must start with a letter and can only contain letters,
	// digits, hyphens (-), underscores (_) and periods (.).
	Name string `json:"name" required:"true"`
	// Parameter location. The valid values are PATH, QUERY and HEADER.
	Location string `json:"location" required:"true"`
	// Parameter value, which can contain a maximum of 255 characters. If the origin type is REQUEST, the value of this parameter is the parameter name in req_params.
	// If the origin type is CONSTANT, the value is a constant.
	// If the origin type is SYSTEM, the value is a system parameter name. System parameters include gateway parameters, frontend authentication parameters, and backend authentication parameters. You can set the frontend or backend authentication parameters after enabling custom frontend or backend authentication.
	// The gateway parameters are as follows:
	//   $context.sourceIp: source IP address of the API caller.
	//   $context.stage: deployment environment in which the API is called.
	//   $context.apiId: API ID.
	//   $context.appId: ID of the app used by the API caller.
	//   $context.requestId: request ID generated when the API is called.
	//   $context.serverAddr: address of the gateway server.
	//   $context.serverName: name of the gateway server.
	//   $context.handleTime: time when the API request is processed.
	//   $context.providerAppId: ID of the app used by the API owner. This parameter is currently not supported.
	// Frontend authentication parameter: prefixed with "$context.authorizer.frontend.". For example, to return "aaa" upon successful custom authentication, set this parameter to "$context.authorizer.frontend.aaa".
	// Backend authentication parameter: prefixed with "$context.authorizer.backend.". For example, to return "aaa" upon successful custom authentication, set this parameter to "$context.authorizer.backend.aaa".
	Value string `json:"value" required:"true"`
	// Description, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description *string `json:"remark,omitempty"`
}

// APIConditionBase is an object which will be build up a policy condition.
type APIConditionBase struct {
	// Policy type. The valid types are as following:
	//   param: input parameter
	//   source: source IP address
	ConditionOrigin string `json:"condition_origin" required:"true"`
	// Condition value.
	ConditionValue string `json:"condition_value" required:"true"`
	// Input parameter name. This parameter is required if the policy type is param.
	ReqParamName string `json:"req_param_name,omitempty"`
	// Policy condition. The valid values are as following:
	//   exact: exact match
	//   enum: enumeration
	//   pattern: regular expression
	// This parameter is required if the policy type is param.
	ConditionType string `json:"condition_type,omitempty"`
}

// APIOptsBuilder is an interface which to support request body build of the API creation and updation.
type APIOptsBuilder interface {
	ToAPIOptsMap() (map[string]interface{}, error)
}

// ToAPIOptsMap is a method which to build a request body by the APIOpts.
func (opts APIOpts) ToAPIOptsMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method by which to create function that create a new custom API.
func Create(client *golangsdk.ServiceClient, instanceId string, opts APIOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToAPIOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client, instanceId), reqBody, &r.Body, nil)
	return
}

// Update is a method to update an existing custom API.
func Update(client *golangsdk.ServiceClient, instanceId, appId string, opts APIOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToAPIOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, instanceId, appId), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get is a method to obtain the specified API according to the instanceId and API ID.
func Get(client *golangsdk.ServiceClient, instanceId, apiId string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, instanceId, apiId), &r.Body, nil)
	return
}

// ListOpts allows to filter list data using given parameters.
type ListOpts struct {
	// API ID.
	ID string `q:"id"`
	// API name.
	Name string `q:"name"`
	// API group ID.
	GroupId string `q:"group_id"`
	// Request protocol.
	ReqProtocol string `q:"req_protocol"`
	// Request method.
	ReqMethod string `q:"req_method"`
	// Request path.
	ReqURI string `q:"req_uri"`
	// Security authentication mode.
	AuthType string `q:"auth_type"`
	// ID of the environment in which the API has been published.
	EnvId string `q:"env_id"`
	// API type.
	Type int `q:"type"`
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page. The range of number is form 1 to 500, default to 20.
	Limit int `q:"limit"`
	// Parameter name (name or req_uri) for exact matching.
	PreciseSearch string `q:"precise_search"`
}

// ListOptsBuilder is an interface which to support request query build of the API search.
type ListOptsBuilder interface {
	ToListOptsQuery() (string, error)
}

// ToListOptsQuery is a method which to build a request query by the ListOpts.
func (opts ListOpts) ToListOptsQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// List is a method to obtain an array of one or more APIs according to the query parameters.
func List(client *golangsdk.ServiceClient, instanceId string, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client, instanceId)
	if opts != nil {
		query, err := opts.ToListOptsQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return APIPage{pagination.SinglePageBase(r)}
	})
}

// Delete is a method to delete an existing custom API.
func Delete(client *golangsdk.ServiceClient, instanceId, apiId string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, instanceId, apiId), nil)
	return
}
