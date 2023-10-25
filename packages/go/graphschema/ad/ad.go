// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by Cuelang code gen. DO NOT EDIT!
// Cuelang source: github.com/specterops/bloodhound/-/tree/main/packages/cue/schemas/

package ad

import (
	"errors"
	graph "github.com/specterops/bloodhound/dawgs/graph"
)

var (
	Entity                          = graph.StringKind("Base")
	User                            = graph.StringKind("User")
	Computer                        = graph.StringKind("Computer")
	Group                           = graph.StringKind("Group")
	GPO                             = graph.StringKind("GPO")
	OU                              = graph.StringKind("OU")
	Container                       = graph.StringKind("Container")
	Domain                          = graph.StringKind("Domain")
	LocalGroup                      = graph.StringKind("ADLocalGroup")
	LocalUser                       = graph.StringKind("ADLocalUser")
	AIACA                           = graph.StringKind("AIACA")
	RootCA                          = graph.StringKind("RootCA")
	EnterpriseCA                    = graph.StringKind("EnterpriseCA")
	NTAuthStore                     = graph.StringKind("NTAuthStore")
	CertTemplate                    = graph.StringKind("CertTemplate")
	Owns                            = graph.StringKind("Owns")
	GenericAll                      = graph.StringKind("GenericAll")
	GenericWrite                    = graph.StringKind("GenericWrite")
	WriteOwner                      = graph.StringKind("WriteOwner")
	WriteDACL                       = graph.StringKind("WriteDacl")
	MemberOf                        = graph.StringKind("MemberOf")
	ForceChangePassword             = graph.StringKind("ForceChangePassword")
	AllExtendedRights               = graph.StringKind("AllExtendedRights")
	AddMember                       = graph.StringKind("AddMember")
	HasSession                      = graph.StringKind("HasSession")
	Contains                        = graph.StringKind("Contains")
	GPLink                          = graph.StringKind("GPLink")
	AllowedToDelegate               = graph.StringKind("AllowedToDelegate")
	GetChanges                      = graph.StringKind("GetChanges")
	GetChangesAll                   = graph.StringKind("GetChangesAll")
	GetChangesInFilteredSet         = graph.StringKind("GetChangesInFilteredSet")
	TrustedBy                       = graph.StringKind("TrustedBy")
	AllowedToAct                    = graph.StringKind("AllowedToAct")
	AdminTo                         = graph.StringKind("AdminTo")
	CanPSRemote                     = graph.StringKind("CanPSRemote")
	CanRDP                          = graph.StringKind("CanRDP")
	ExecuteDCOM                     = graph.StringKind("ExecuteDCOM")
	HasSIDHistory                   = graph.StringKind("HasSIDHistory")
	AddSelf                         = graph.StringKind("AddSelf")
	DCSync                          = graph.StringKind("DCSync")
	ReadLAPSPassword                = graph.StringKind("ReadLAPSPassword")
	ReadGMSAPassword                = graph.StringKind("ReadGMSAPassword")
	DumpSMSAPassword                = graph.StringKind("DumpSMSAPassword")
	SQLAdmin                        = graph.StringKind("SQLAdmin")
	AddAllowedToAct                 = graph.StringKind("AddAllowedToAct")
	WriteSPN                        = graph.StringKind("WriteSPN")
	AddKeyCredentialLink            = graph.StringKind("AddKeyCredentialLink")
	LocalToComputer                 = graph.StringKind("LocalToComputer")
	MemberOfLocalGroup              = graph.StringKind("MemberOfLocalGroup")
	RemoteInteractiveLogonPrivilege = graph.StringKind("RemoteInteractiveLogonPrivilege")
	SyncLAPSPassword                = graph.StringKind("SyncLAPSPassword")
	WriteAccountRestrictions        = graph.StringKind("WriteAccountRestrictions")
	RootCAFor                       = graph.StringKind("RootCAFor")
	PublishedTo                     = graph.StringKind("PublishedTo")
	ManageCertificates              = graph.StringKind("ManageCertificates")
	ManageCA                        = graph.StringKind("ManageCA")
	DelegatedEnrollmentAgent        = graph.StringKind("DelegatedEnrollmentAgent")
	Enroll                          = graph.StringKind("Enroll")
	HostsCAService                  = graph.StringKind("HostsCAService")
	WritePKIEnrollmentFlag          = graph.StringKind("WritePKIEnrollmentFlag")
	WritePKINameFlag                = graph.StringKind("WritePKINameFlag")
	NTAuthStoreFor                  = graph.StringKind("NTAuthStoreFor")
	TrustedForNTAuth                = graph.StringKind("TrustedForNTAuth")
	EnterpriseCAFor                 = graph.StringKind("EnterpriseCAFor")
	IssuedSignedBy                  = graph.StringKind("IssuedSignedBy")
	GoldenCert                      = graph.StringKind("GoldenCert")
	EnrollOnBehalfOf                = graph.StringKind("EnrollOnBehalfOf")
	ADCSESC1                        = graph.StringKind("ADCSESC1")
	ADCSESC2                        = graph.StringKind("ADCSESC2")
	ADCSESC3                        = graph.StringKind("ADCSESC3")
	ADCSESC4                        = graph.StringKind("ADCSESC4")
	ADCSESC5                        = graph.StringKind("ADCSESC5")
	ADCSESC6                        = graph.StringKind("ADCSESC6")
	ADCSESC7                        = graph.StringKind("ADCSESC7")
)

type Property string

const (
	AdminCount                                   Property = "admincount"
	CASecurityCollected                          Property = "casecuritycollected"
	CAName                                       Property = "caname"
	CertChain                                    Property = "certchain"
	CertName                                     Property = "certname"
	CertThumbprint                               Property = "certthumbprint"
	CertThumbprints                              Property = "certthumbprints"
	EnrollmentAgentRestrictionsCollected         Property = "enrollmentagentrestrictionscollected"
	IsUserSpecifiesSanEnabledCollected           Property = "isuserspecifiessanenabledcollected"
	HasBasicConstraints                          Property = "hasbasicconstraints"
	BasicConstraintPathLength                    Property = "basicconstraintpathlength"
	DNSHostname                                  Property = "dnshostname"
	CrossCertificatePair                         Property = "crosscertificatepair"
	DistinguishedName                            Property = "distinguishedname"
	DomainFQDN                                   Property = "domain"
	DomainSID                                    Property = "domainsid"
	Sensitive                                    Property = "sensitive"
	HighValue                                    Property = "highvalue"
	BlocksInheritance                            Property = "blocksinheritance"
	IsACL                                        Property = "isacl"
	IsACLProtected                               Property = "isaclprotected"
	IsDeleted                                    Property = "isdeleted"
	Enforced                                     Property = "enforced"
	Department                                   Property = "department"
	HasCrossCertificatePair                      Property = "hascrosscertificatepair"
	HasSPN                                       Property = "hasspn"
	UnconstrainedDelegation                      Property = "unconstraineddelegation"
	LastLogon                                    Property = "lastlogon"
	LastLogonTimestamp                           Property = "lastlogontimestamp"
	IsPrimaryGroup                               Property = "isprimarygroup"
	HasLAPS                                      Property = "haslaps"
	DontRequirePreAuth                           Property = "dontreqpreauth"
	LogonType                                    Property = "logontype"
	HasURA                                       Property = "hasura"
	PasswordNeverExpires                         Property = "pwdneverexpires"
	PasswordNotRequired                          Property = "passwordnotreqd"
	FunctionalLevel                              Property = "functionallevel"
	TrustType                                    Property = "trusttype"
	SidFiltering                                 Property = "sidfiltering"
	TrustedToAuth                                Property = "trustedtoauth"
	SamAccountName                               Property = "samaccountname"
	CertificateMappingMethodsCollected           Property = "certificatemappingmethodscollected"
	CertificateMappingMethodsHex                 Property = "certificatemappingmethodshex"
	CertificateMappingMethodsPretty              Property = "certificatemappingmethodspretty"
	StrongCertificateBindingEnforcementCollected Property = "strongcertificatebindingenforcementcollected"
	StrongCertificateBindingEnforcementInt       Property = "strongcertificatebindingenforcementint"
	StrongCertificateBindingEnforcementPretty    Property = "strongcertificatebindingenforcementpretty"
	EKUs                                         Property = "ekus"
	SubjectAltRequireUPN                         Property = "subjectaltrequireupn"
	AuthorizedSignatures                         Property = "authorizedsignatures"
	ApplicationPolicies                          Property = "applicationpolicies"
	SchemaVersion                                Property = "schemaversion"
	RequiresManagerApproval                      Property = "requiresmanagerapproval"
	AuthenticationEnabled                        Property = "authenticationenabled"
	EnrolleeSuppliesSubject                      Property = "enrolleesuppliessubject"
	AdminCount                             Property = "admincount"
	CASecurityCollected                    Property = "casecuritycollected"
	CAName                                 Property = "caname"
	CertChain                              Property = "certchain"
	CertName                               Property = "certname"
	CertThumbprint                         Property = "certthumbprint"
	CertThumbprints                        Property = "certthumbprints"
	EnrollmentAgentRestrictionsCollected   Property = "enrollmentagentrestrictionscollected"
	IsUserSpecifiesSanEnabledCollected     Property = "isuserspecifiessanenabledcollected"
	HasBasicConstraints                    Property = "hasbasicconstraints"
	BasicConstraintPathLength              Property = "basicconstraintpathlength"
	DNSHostname                            Property = "dnshostname"
	CrossCertificatePair                   Property = "crosscertificatepair"
	DistinguishedName                      Property = "distinguishedname"
	DomainFQDN                             Property = "domain"
	DomainSID                              Property = "domainsid"
	Sensitive                              Property = "sensitive"
	HighValue                              Property = "highvalue"
	BlocksInheritance                      Property = "blocksinheritance"
	IsACL                                  Property = "isacl"
	IsACLProtected                         Property = "isaclprotected"
	IsDeleted                              Property = "isdeleted"
	Enforced                               Property = "enforced"
	Department                             Property = "department"
	HasCrossCertificatePair                Property = "hascrosscertificatepair"
	HasSPN                                 Property = "hasspn"
	UnconstrainedDelegation                Property = "unconstraineddelegation"
	LastLogon                              Property = "lastlogon"
	LastLogonTimestamp                     Property = "lastlogontimestamp"
	IsPrimaryGroup                         Property = "isprimarygroup"
	HasLAPS                                Property = "haslaps"
	DontRequirePreAuth                     Property = "dontreqpreauth"
	LogonType                              Property = "logontype"
	HasURA                                 Property = "hasura"
	PasswordNeverExpires                   Property = "pwdneverexpires"
	PasswordNotRequired                    Property = "passwordnotreqd"
	FunctionalLevel                        Property = "functionallevel"
	TrustType                              Property = "trusttype"
	SidFiltering                           Property = "sidfiltering"
	TrustedToAuth                          Property = "trustedtoauth"
	SamAccountName                         Property = "samaccountname"
	CertificateMappingMethodsRaw           Property = "certificatemappingmethodsraw"
	CertificateMappingMethods              Property = "certificatemappingmethods"
	StrongCertificateBindingEnforcementRaw Property = "strongcertificatebindingenforcementraw"
	StrongCertificateBindingEnforcement    Property = "strongcertificatebindingenforcement"
	EKUs                                   Property = "ekus"
	SubjectAltRequireUPN                   Property = "subjectaltrequireupn"
	AuthorizedSignatures                   Property = "authorizedsignatures"
	ApplicationPolicies                    Property = "applicationpolicies"
	SchemaVersion                          Property = "schemaversion"
)

func AllProperties() []Property {
	return []Property{AdminCount, CASecurityCollected, CAName, CertChain, CertName, CertThumbprint, CertThumbprints, EnrollmentAgentRestrictionsCollected, IsUserSpecifiesSanEnabledCollected, HasBasicConstraints, BasicConstraintPathLength, DNSHostname, CrossCertificatePair, DistinguishedName, DomainFQDN, DomainSID, Sensitive, HighValue, BlocksInheritance, IsACL, IsACLProtected, IsDeleted, Enforced, Department, HasCrossCertificatePair, HasSPN, UnconstrainedDelegation, LastLogon, LastLogonTimestamp, IsPrimaryGroup, HasLAPS, DontRequirePreAuth, LogonType, HasURA, PasswordNeverExpires, PasswordNotRequired, FunctionalLevel, TrustType, SidFiltering, TrustedToAuth, SamAccountName, CertificateMappingMethodsRaw, CertificateMappingMethods, StrongCertificateBindingEnforcementRaw, StrongCertificateBindingEnforcement, EKUs, SubjectAltRequireUPN, AuthorizedSignatures, ApplicationPolicies, SchemaVersion}
}
func ParseProperty(source string) (Property, error) {
	switch source {
	case "admincount":
		return AdminCount, nil
	case "casecuritycollected":
		return CASecurityCollected, nil
	case "caname":
		return CAName, nil
	case "certchain":
		return CertChain, nil
	case "certname":
		return CertName, nil
	case "certthumbprint":
		return CertThumbprint, nil
	case "certthumbprints":
		return CertThumbprints, nil
	case "enrollmentagentrestrictionscollected":
		return EnrollmentAgentRestrictionsCollected, nil
	case "isuserspecifiessanenabledcollected":
		return IsUserSpecifiesSanEnabledCollected, nil
	case "hasbasicconstraints":
		return HasBasicConstraints, nil
	case "basicconstraintpathlength":
		return BasicConstraintPathLength, nil
	case "dnshostname":
		return DNSHostname, nil
	case "crosscertificatepair":
		return CrossCertificatePair, nil
	case "distinguishedname":
		return DistinguishedName, nil
	case "domain":
		return DomainFQDN, nil
	case "domainsid":
		return DomainSID, nil
	case "sensitive":
		return Sensitive, nil
	case "highvalue":
		return HighValue, nil
	case "blocksinheritance":
		return BlocksInheritance, nil
	case "isacl":
		return IsACL, nil
	case "isaclprotected":
		return IsACLProtected, nil
	case "isdeleted":
		return IsDeleted, nil
	case "enforced":
		return Enforced, nil
	case "department":
		return Department, nil
	case "hascrosscertificatepair":
		return HasCrossCertificatePair, nil
	case "hasspn":
		return HasSPN, nil
	case "unconstraineddelegation":
		return UnconstrainedDelegation, nil
	case "lastlogon":
		return LastLogon, nil
	case "lastlogontimestamp":
		return LastLogonTimestamp, nil
	case "isprimarygroup":
		return IsPrimaryGroup, nil
	case "haslaps":
		return HasLAPS, nil
	case "dontreqpreauth":
		return DontRequirePreAuth, nil
	case "logontype":
		return LogonType, nil
	case "hasura":
		return HasURA, nil
	case "pwdneverexpires":
		return PasswordNeverExpires, nil
	case "passwordnotreqd":
		return PasswordNotRequired, nil
	case "functionallevel":
		return FunctionalLevel, nil
	case "trusttype":
		return TrustType, nil
	case "sidfiltering":
		return SidFiltering, nil
	case "trustedtoauth":
		return TrustedToAuth, nil
	case "samaccountname":
		return SamAccountName, nil
	case "certificatemappingmethodsraw":
		return CertificateMappingMethodsRaw, nil
	case "certificatemappingmethods":
		return CertificateMappingMethods, nil
	case "strongcertificatebindingenforcementraw":
		return StrongCertificateBindingEnforcementRaw, nil
	case "strongcertificatebindingenforcement":
		return StrongCertificateBindingEnforcement, nil
	case "ekus":
		return EKUs, nil
	case "subjectaltrequireupn":
		return SubjectAltRequireUPN, nil
	case "authorizedsignatures":
		return AuthorizedSignatures, nil
	case "applicationpolicies":
		return ApplicationPolicies, nil
	case "schemaversion":
		return SchemaVersion, nil
	case "requiresmanagerapproval":
		return RequiresManagerApproval, nil
	case "authenticationenabled":
		return AuthenticationEnabled, nil
	case "enrolleesuppliessubject":
		return EnrolleeSuppliesSubject, nil
	default:
		return "", errors.New("Invalid enumeration value: " + source)
	}
}
func (s Property) String() string {
	switch s {
	case AdminCount:
		return string(AdminCount)
	case CASecurityCollected:
		return string(CASecurityCollected)
	case CAName:
		return string(CAName)
	case CertChain:
		return string(CertChain)
	case CertName:
		return string(CertName)
	case CertThumbprint:
		return string(CertThumbprint)
	case CertThumbprints:
		return string(CertThumbprints)
	case EnrollmentAgentRestrictionsCollected:
		return string(EnrollmentAgentRestrictionsCollected)
	case IsUserSpecifiesSanEnabledCollected:
		return string(IsUserSpecifiesSanEnabledCollected)
	case HasBasicConstraints:
		return string(HasBasicConstraints)
	case BasicConstraintPathLength:
		return string(BasicConstraintPathLength)
	case DNSHostname:
		return string(DNSHostname)
	case CrossCertificatePair:
		return string(CrossCertificatePair)
	case DistinguishedName:
		return string(DistinguishedName)
	case DomainFQDN:
		return string(DomainFQDN)
	case DomainSID:
		return string(DomainSID)
	case Sensitive:
		return string(Sensitive)
	case HighValue:
		return string(HighValue)
	case BlocksInheritance:
		return string(BlocksInheritance)
	case IsACL:
		return string(IsACL)
	case IsACLProtected:
		return string(IsACLProtected)
	case IsDeleted:
		return string(IsDeleted)
	case Enforced:
		return string(Enforced)
	case Department:
		return string(Department)
	case HasCrossCertificatePair:
		return string(HasCrossCertificatePair)
	case HasSPN:
		return string(HasSPN)
	case UnconstrainedDelegation:
		return string(UnconstrainedDelegation)
	case LastLogon:
		return string(LastLogon)
	case LastLogonTimestamp:
		return string(LastLogonTimestamp)
	case IsPrimaryGroup:
		return string(IsPrimaryGroup)
	case HasLAPS:
		return string(HasLAPS)
	case DontRequirePreAuth:
		return string(DontRequirePreAuth)
	case LogonType:
		return string(LogonType)
	case HasURA:
		return string(HasURA)
	case PasswordNeverExpires:
		return string(PasswordNeverExpires)
	case PasswordNotRequired:
		return string(PasswordNotRequired)
	case FunctionalLevel:
		return string(FunctionalLevel)
	case TrustType:
		return string(TrustType)
	case SidFiltering:
		return string(SidFiltering)
	case TrustedToAuth:
		return string(TrustedToAuth)
	case SamAccountName:
		return string(SamAccountName)
	case CertificateMappingMethodsRaw:
		return string(CertificateMappingMethodsRaw)
	case CertificateMappingMethods:
		return string(CertificateMappingMethods)
	case StrongCertificateBindingEnforcementRaw:
		return string(StrongCertificateBindingEnforcementRaw)
	case StrongCertificateBindingEnforcement:
		return string(StrongCertificateBindingEnforcement)
	case EKUs:
		return string(EKUs)
	case SubjectAltRequireUPN:
		return string(SubjectAltRequireUPN)
	case AuthorizedSignatures:
		return string(AuthorizedSignatures)
	case ApplicationPolicies:
		return string(ApplicationPolicies)
	case SchemaVersion:
		return string(SchemaVersion)
	case RequiresManagerApproval:
		return string(RequiresManagerApproval)
	case AuthenticationEnabled:
		return string(AuthenticationEnabled)
	case EnrolleeSuppliesSubject:
		return string(EnrolleeSuppliesSubject)
	default:
		panic("Invalid enumeration case: " + string(s))
	}
}
func (s Property) Name() string {
	switch s {
	case AdminCount:
		return "Admin Count"
	case CASecurityCollected:
		return "CA Security Collected"
	case CAName:
		return "CA Name"
	case CertChain:
		return "Certificate Chain"
	case CertName:
		return "Certificate Name"
	case CertThumbprint:
		return "Certificate Thumbprint"
	case CertThumbprints:
		return "Certificate Thumbprints"
	case EnrollmentAgentRestrictionsCollected:
		return "Enrollment Agent Restrictions Collected"
	case IsUserSpecifiesSanEnabledCollected:
		return "Is User Specifies San Enabled Collected"
	case HasBasicConstraints:
		return "Has Basic Constraints"
	case BasicConstraintPathLength:
		return "Basic Constraint Path Length"
	case DNSHostname:
		return "DNS Hostname"
	case CrossCertificatePair:
		return "Cross Certificate Pair"
	case DistinguishedName:
		return "Distinguished Name"
	case DomainFQDN:
		return "Domain FQDN"
	case DomainSID:
		return "Domain SID"
	case Sensitive:
		return "Marked sensitive"
	case HighValue:
		return "High Value"
	case BlocksInheritance:
		return "Blocks Inheritance"
	case IsACL:
		return "Is ACL"
	case IsACLProtected:
		return "ACL Inheritance Denied"
	case IsDeleted:
		return "Is Deleted"
	case Enforced:
		return "Enforced"
	case Department:
		return "Department"
	case HasCrossCertificatePair:
		return "Has Cross Certificate Pair"
	case HasSPN:
		return "Has SPN"
	case UnconstrainedDelegation:
		return "Allows Unconstrained Delegation"
	case LastLogon:
		return "Last Logon"
	case LastLogonTimestamp:
		return "Last Logon (Replicated)"
	case IsPrimaryGroup:
		return "Is Primary Group"
	case HasLAPS:
		return "LAPS Enabled"
	case DontRequirePreAuth:
		return "Do Not Require Pre-Authentication"
	case LogonType:
		return "Logon Type"
	case HasURA:
		return "Has User Rights Assignment Collection"
	case PasswordNeverExpires:
		return "Password Never Expires"
	case PasswordNotRequired:
		return "Password Not Required"
	case FunctionalLevel:
		return "Functional Level"
	case TrustType:
		return "Trust Type"
	case SidFiltering:
		return "SID Filtering Enabled"
	case TrustedToAuth:
		return "Trusted For Constrained Delegation"
	case SamAccountName:
		return "SAM Account Name"
	case CertificateMappingMethodsRaw:
		return "Certificate Mapping Methods (Raw)"
	case CertificateMappingMethods:
		return "Certificate Mapping Methods"
	case StrongCertificateBindingEnforcementRaw:
		return "Strong Certificate Binding Enforcement (Raw)"
	case StrongCertificateBindingEnforcement:
		return "Strong Certificate Binding Enforcement"
	case EKUs:
		return "EKUs"
	case SubjectAltRequireUPN:
		return "Subject Alt Require UPN"
	case AuthorizedSignatures:
		return "Authorized Signatures"
	case ApplicationPolicies:
		return "Application Policies"
	case SchemaVersion:
		return "Schema Version"
	case RequiresManagerApproval:
		return "Requires Manager Approval"
	case AuthenticationEnabled:
		return "Authentication Enabled"
	case EnrolleeSuppliesSubject:
		return "Enrollee Suppliess Subject"
	default:
		panic("Invalid enumeration case: " + string(s))
	}
}
func (s Property) Is(others ...graph.Kind) bool {
	for _, other := range others {
		if value, err := ParseProperty(other.String()); err == nil && value == s {
			return true
		}
	}
	return false
}
func Nodes() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser, AIACA, RootCA, EnterpriseCA, NTAuthStore, CertTemplate}
}
func Relationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, GetChanges, GetChangesAll, GetChangesInFilteredSet, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, LocalToComputer, MemberOfLocalGroup, RemoteInteractiveLogonPrivilege, SyncLAPSPassword, WriteAccountRestrictions, RootCAFor, PublishedTo, ManageCertificates, ManageCA, DelegatedEnrollmentAgent, Enroll, HostsCAService, WritePKIEnrollmentFlag, WritePKINameFlag, NTAuthStoreFor, TrustedForNTAuth, EnterpriseCAFor, IssuedSignedBy, GoldenCert, EnrollOnBehalfOf, ADCSESC1, ADCSESC2, ADCSESC3, ADCSESC4, ADCSESC5, ADCSESC6, ADCSESC7}
}
func ACLRelationships() []graph.Kind {
	return []graph.Kind{AllExtendedRights, ForceChangePassword, AddMember, AddAllowedToAct, GenericAll, WriteDACL, WriteOwner, GenericWrite, ReadLAPSPassword, ReadGMSAPassword, Owns, AddSelf, WriteSPN, AddKeyCredentialLink, GetChanges, GetChangesAll, GetChangesInFilteredSet, WriteAccountRestrictions, SyncLAPSPassword, DCSync, ManageCertificates, ManageCA, Enroll, WritePKIEnrollmentFlag, WritePKINameFlag}
}
func PathfindingRelationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, SyncLAPSPassword, WriteAccountRestrictions, GoldenCert, ADCSESC1, ADCSESC2, ADCSESC3, ADCSESC4, ADCSESC5, ADCSESC6, ADCSESC7}
}
func IsACLKind(s graph.Kind) bool {
	for _, acl := range ACLRelationships() {
		if s == acl {
			return true
		}
	}
	return false
}
func NodeKinds() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser, AIACA, RootCA, EnterpriseCA, NTAuthStore, CertTemplate}
}
