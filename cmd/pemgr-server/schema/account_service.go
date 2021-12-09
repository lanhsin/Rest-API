package schema

type AccountService struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The description of this Resource.
	// Used for commonality in the schema definitions. [RO]
	Description string `json:"Description"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// An indication of whether the Account Service is enabled.
	// If `true`, it is enabled.
	// If `false`, it is disabled and users cannot be created, deleted, or modified, and new sessions cannot be started.
	// However, established sessions may still continue to run.
	// Any service, such as the Session Service, that attempts to access the disabled Account Service fails.
	// However, this does not affect HTTP Basic Authentication connections. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The number of authorization failures that are allowed before the failed attempt is logged to the manager log. [RW]
	AuthFailureLoggingThreshold int `json:"AuthFailureLoggingThreshold"`

	// The minimum password length for this Account Service. [RO]
	MinPasswordLength int `json:"MinPasswordLength"`

	// The number of allowed failed login attempts before a user account is locked for a specified duration.
	// If `0`, the account is never locked. [RW]
	AccountLockoutThreshold int `json:"AccountLockoutThreshold"`

	// The period of time, in seconds, that an account is locked after the number of failed login attempts reaches the account lockout threshold,
	// within the period between the last failed login attempt and the reset of the lockout threshold counter.
	// If this value is `0`, no lockout will occur.
	// If the AccountLockoutCounterResetEnabled value is `false`, this property is ignored. [RW]
	AccountLockoutDuration int `json:"AccountLockoutDuration"`

	// The period of time, in seconds, between the last failed login attempt and the reset of the lockout threshold counter.
	// This value must be less than or equal to the AccountLockoutDuration value. A reset sets the counter to `0`. [RW]
	AccountLockoutCounterResetAfter int `json:"AccountLockoutCounterResetAfter"`

	// An indication of whether the threshold counter is reset after AccountLockoutCounterResetAfter expires.
	// If `true`, it is reset. If `false`, only a successful login resets the threshold counter and
	// if the user reaches the AccountLockoutThreshold limit,
	// the account will be locked out indefinitely and only an administrator-issued reset clears the threshold counter.
	// If this property is absent, the default is `true`. [RW]
	AccountLockoutCounterResetEnabled bool `json:"AccountLockoutCounterResetEnabled"`

	// A Collection of ManagerAccount Resource instances. [RO]
	Accounts CommonOid `json:"Accounts"`

	// The RoleCollection schema describes a collection of role instances. [RO]
	Roles CommonOid `json:"Roles"`

	// An indication of how the Service uses the accounts collection within this Account Service as part of authentication.
	// The enumerated values describe the details for each mode. [RW]
	LocalAccountAuth string `json:"LocalAccountAuth"`

	// The external account provider services that can provide accounts for this manager to use for authentication. [RW]
	LDAP AccountServiceLDAP `json:"LDAP"`

	// The external account provider services that can provide accounts for this manager to use for authentication. [RW]
	ActiveDirectory AccountServiceActiveDirectory `json:"ActiveDirectory"`

	// The collection of ExternalAccountProvider Resource instances. [RO]
	AdditionalExternalAccountProviders CommonOid `json:"AdditionalExternalAccountProviders"`

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type AccountServiceLDAP struct {
	// The type of external account provider to which this Service connects. [RO]
	// Valid values:
	// ActiveDirectoryService:	An external Active Directory service.
	// LDAPService:	            A generic external LDAP service.
	// OEM:	                    An OEM-specific external authentication or directory service.
	// RedfishService:	        An external Redfish Service.
	AccountProviderType string `json:"AccountProviderType"`

	// An indication of whether this service is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The addresses of the user account providers to which this external account provider links.
	// The format of this field depends on the type of external account provider.
	ServiceAddresses []string `json:"ServiceAddresses"`

	// The information required to authenticate to the external service. [RW]
	Authentication AccountServiceLDAPAuthentication `json:"Authentication"`

	// The settings required to parse a generic LDAP service. [RW]
	LDAPService AccountServiceLDAPLDAPService `json:"LDAPService"`

	// The mapping rules to convert the external account providers account information to the local Redfish Role. [RW]
	RemoteRoleMapping []AccountServiceLDAPRemoteRoleMapping `json:"RemoteRoleMapping"`
}

type AccountServiceLDAPAuthentication struct {
	// The type of authentication used to connect to the external account provider. [RW]
	// Valid values:
	// KerberosKeytab:	    A Kerberos keytab.
	// OEM:	                An OEM-specific authentication mechanism.
	// Token:	            An opaque authentication token.
	// UsernameAndPassword:	A user name and password combination.
	AuthenticationType string `json:"AuthenticationType"`

	// The user name for the Service. [RW]
	Username string `json:"Username"`

	// The password for this Service. A PATCH or PUT request writes the password.
	// This property is `null` in responses. [RW]
	Password string `json:"Password,omitempty"`
}

type AccountServiceLDAPLDAPService struct {
	// The settings to search a generic LDAP service. [RW]
	SearchSetting AccountServiceLDAPLDAPServiceSearchSetting `json:"SearchSetting"`
}

type AccountServiceLDAPLDAPServiceSearchSetting struct {
	// The base distinguished names to use to search an external LDAP service. [RW]
	BaseDistinguishedNames []string `json:"BaseDistinguishedNames"`

	// The attribute name that contains the LDAP user name entry. [RW]
	UsernameAttribute string `json:"UsernameAttribute"`

	// The attribute name that contains the groups for a user on the LDAP user entry. [RW]
	GroupsAttribute string `json:"GroupsAttribute"`
}

type AccountServiceLDAPRemoteRoleMapping struct {
	// The name of the remote user that maps to the local Redfish Role to which this entity links. [RW]
	RemoteUser string `json:"RemoteUser"`

	// The name of the local Redfish Role to which to map the remote user or group. [RW]
	LocalRole string `json:"LocalRole"`
}

type AccountServiceActiveDirectory struct {
	// The type of external account provider to which this Service connects. [RO]
	// Valid values:
	// ActiveDirectoryService:	An external Active Directory service.
	// LDAPService:	            A generic external LDAP service.
	// OEM:	                    An OEM-specific external authentication or directory service.
	// RedfishService:	        An external Redfish Service.
	AccountProviderType string `json:"AccountProviderType"`

	// An indication of whether this service is enabled. [RW]
	ServiceEnabled bool `json:"ServiceEnabled"`

	// The addresses of the user account providers to which this external account provider links.
	// The format of this field depends on the type of external account provider.
	ServiceAddresses []string `json:"ServiceAddresses"`

	// The information required to authenticate to the external service. [RW]
	Authentication AccountServiceActiveDirectoryAuthentication `json:"Authentication"`

	// The mapping rules to convert the external account providers account information to the local Redfish Role. [RW]
	RemoteRoleMapping []AccountServiceActiveDirectoryRemoteRoleMapping `json:"RemoteRoleMapping"`
}

type AccountServiceActiveDirectoryAuthentication struct {
	// The type of authentication used to connect to the external account provider. [RW]
	// Valid values:
	// KerberosKeytab:	    A Kerberos keytab.
	// OEM:	                An OEM-specific authentication mechanism.
	// Token:	            An opaque authentication token.
	// UsernameAndPassword:	A user name and password combination.
	AuthenticationType string `json:"AuthenticationType"`

	// The Base64-encoded version of the Kerberos keytab for this Service.
	// A PATCH or PUT operation writes the keytab. This property is `null` in responses. [RW]
	KerberosKeytab string `json:"KerberosKeytab,omitempty"`
}

type AccountServiceActiveDirectoryRemoteRoleMapping struct {
	// The name of the remote group, or the remote role in the case of a Redfish Service,
	// that maps to the local Redfish Role to which this entity links. [RW]
	RemoteGroup string `json:"RemoteGroup,omitempty"`

	// The name of the local Redfish Role to which to map the remote user or group. [RW]
	LocalRole string `json:"LocalRole,omitempty"`
}
