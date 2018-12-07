// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package alexaforbusinessiface provides an interface to enable mocking the Alexa For Business service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package alexaforbusinessiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
)

// AlexaForBusinessAPI provides an interface to enable mocking the
// alexaforbusiness.AlexaForBusiness service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Alexa For Business.
//    func myFunc(svc alexaforbusinessiface.AlexaForBusinessAPI) bool {
//        // Make svc.ApproveSkill request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := alexaforbusiness.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAlexaForBusinessClient struct {
//        alexaforbusinessiface.AlexaForBusinessAPI
//    }
//    func (m *mockAlexaForBusinessClient) ApproveSkill(input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAlexaForBusinessClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AlexaForBusinessAPI interface {
	ApproveSkill(*alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error)
	ApproveSkillWithContext(aws.Context, *alexaforbusiness.ApproveSkillInput, ...request.Option) (*alexaforbusiness.ApproveSkillOutput, error)
	ApproveSkillRequest(*alexaforbusiness.ApproveSkillInput) (*request.Request, *alexaforbusiness.ApproveSkillOutput)

	AssociateContactWithAddressBook(*alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateContactWithAddressBookWithContext(aws.Context, *alexaforbusiness.AssociateContactWithAddressBookInput, ...request.Option) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateContactWithAddressBookRequest(*alexaforbusiness.AssociateContactWithAddressBookInput) (*request.Request, *alexaforbusiness.AssociateContactWithAddressBookOutput)

	AssociateDeviceWithRoom(*alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateDeviceWithRoomWithContext(aws.Context, *alexaforbusiness.AssociateDeviceWithRoomInput, ...request.Option) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateDeviceWithRoomRequest(*alexaforbusiness.AssociateDeviceWithRoomInput) (*request.Request, *alexaforbusiness.AssociateDeviceWithRoomOutput)

	AssociateSkillGroupWithRoom(*alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillGroupWithRoomWithContext(aws.Context, *alexaforbusiness.AssociateSkillGroupWithRoomInput, ...request.Option) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillGroupWithRoomRequest(*alexaforbusiness.AssociateSkillGroupWithRoomInput) (*request.Request, *alexaforbusiness.AssociateSkillGroupWithRoomOutput)

	AssociateSkillWithSkillGroup(*alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithSkillGroupWithContext(aws.Context, *alexaforbusiness.AssociateSkillWithSkillGroupInput, ...request.Option) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithSkillGroupRequest(*alexaforbusiness.AssociateSkillWithSkillGroupInput) (*request.Request, *alexaforbusiness.AssociateSkillWithSkillGroupOutput)

	CreateAddressBook(*alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateAddressBookWithContext(aws.Context, *alexaforbusiness.CreateAddressBookInput, ...request.Option) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateAddressBookRequest(*alexaforbusiness.CreateAddressBookInput) (*request.Request, *alexaforbusiness.CreateAddressBookOutput)

	CreateBusinessReportSchedule(*alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateBusinessReportScheduleWithContext(aws.Context, *alexaforbusiness.CreateBusinessReportScheduleInput, ...request.Option) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateBusinessReportScheduleRequest(*alexaforbusiness.CreateBusinessReportScheduleInput) (*request.Request, *alexaforbusiness.CreateBusinessReportScheduleOutput)

	CreateConferenceProvider(*alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateConferenceProviderWithContext(aws.Context, *alexaforbusiness.CreateConferenceProviderInput, ...request.Option) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateConferenceProviderRequest(*alexaforbusiness.CreateConferenceProviderInput) (*request.Request, *alexaforbusiness.CreateConferenceProviderOutput)

	CreateContact(*alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error)
	CreateContactWithContext(aws.Context, *alexaforbusiness.CreateContactInput, ...request.Option) (*alexaforbusiness.CreateContactOutput, error)
	CreateContactRequest(*alexaforbusiness.CreateContactInput) (*request.Request, *alexaforbusiness.CreateContactOutput)

	CreateProfile(*alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error)
	CreateProfileWithContext(aws.Context, *alexaforbusiness.CreateProfileInput, ...request.Option) (*alexaforbusiness.CreateProfileOutput, error)
	CreateProfileRequest(*alexaforbusiness.CreateProfileInput) (*request.Request, *alexaforbusiness.CreateProfileOutput)

	CreateRoom(*alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error)
	CreateRoomWithContext(aws.Context, *alexaforbusiness.CreateRoomInput, ...request.Option) (*alexaforbusiness.CreateRoomOutput, error)
	CreateRoomRequest(*alexaforbusiness.CreateRoomInput) (*request.Request, *alexaforbusiness.CreateRoomOutput)

	CreateSkillGroup(*alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateSkillGroupWithContext(aws.Context, *alexaforbusiness.CreateSkillGroupInput, ...request.Option) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateSkillGroupRequest(*alexaforbusiness.CreateSkillGroupInput) (*request.Request, *alexaforbusiness.CreateSkillGroupOutput)

	CreateUser(*alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *alexaforbusiness.CreateUserInput, ...request.Option) (*alexaforbusiness.CreateUserOutput, error)
	CreateUserRequest(*alexaforbusiness.CreateUserInput) (*request.Request, *alexaforbusiness.CreateUserOutput)

	DeleteAddressBook(*alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteAddressBookWithContext(aws.Context, *alexaforbusiness.DeleteAddressBookInput, ...request.Option) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteAddressBookRequest(*alexaforbusiness.DeleteAddressBookInput) (*request.Request, *alexaforbusiness.DeleteAddressBookOutput)

	DeleteBusinessReportSchedule(*alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteBusinessReportScheduleWithContext(aws.Context, *alexaforbusiness.DeleteBusinessReportScheduleInput, ...request.Option) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteBusinessReportScheduleRequest(*alexaforbusiness.DeleteBusinessReportScheduleInput) (*request.Request, *alexaforbusiness.DeleteBusinessReportScheduleOutput)

	DeleteConferenceProvider(*alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteConferenceProviderWithContext(aws.Context, *alexaforbusiness.DeleteConferenceProviderInput, ...request.Option) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteConferenceProviderRequest(*alexaforbusiness.DeleteConferenceProviderInput) (*request.Request, *alexaforbusiness.DeleteConferenceProviderOutput)

	DeleteContact(*alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteContactWithContext(aws.Context, *alexaforbusiness.DeleteContactInput, ...request.Option) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteContactRequest(*alexaforbusiness.DeleteContactInput) (*request.Request, *alexaforbusiness.DeleteContactOutput)

	DeleteDevice(*alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceWithContext(aws.Context, *alexaforbusiness.DeleteDeviceInput, ...request.Option) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceRequest(*alexaforbusiness.DeleteDeviceInput) (*request.Request, *alexaforbusiness.DeleteDeviceOutput)

	DeleteProfile(*alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteProfileWithContext(aws.Context, *alexaforbusiness.DeleteProfileInput, ...request.Option) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteProfileRequest(*alexaforbusiness.DeleteProfileInput) (*request.Request, *alexaforbusiness.DeleteProfileOutput)

	DeleteRoom(*alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomWithContext(aws.Context, *alexaforbusiness.DeleteRoomInput, ...request.Option) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomRequest(*alexaforbusiness.DeleteRoomInput) (*request.Request, *alexaforbusiness.DeleteRoomOutput)

	DeleteRoomSkillParameter(*alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteRoomSkillParameterWithContext(aws.Context, *alexaforbusiness.DeleteRoomSkillParameterInput, ...request.Option) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteRoomSkillParameterRequest(*alexaforbusiness.DeleteRoomSkillParameterInput) (*request.Request, *alexaforbusiness.DeleteRoomSkillParameterOutput)

	DeleteSkillAuthorization(*alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillAuthorizationWithContext(aws.Context, *alexaforbusiness.DeleteSkillAuthorizationInput, ...request.Option) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillAuthorizationRequest(*alexaforbusiness.DeleteSkillAuthorizationInput) (*request.Request, *alexaforbusiness.DeleteSkillAuthorizationOutput)

	DeleteSkillGroup(*alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteSkillGroupWithContext(aws.Context, *alexaforbusiness.DeleteSkillGroupInput, ...request.Option) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteSkillGroupRequest(*alexaforbusiness.DeleteSkillGroupInput) (*request.Request, *alexaforbusiness.DeleteSkillGroupOutput)

	DeleteUser(*alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *alexaforbusiness.DeleteUserInput, ...request.Option) (*alexaforbusiness.DeleteUserOutput, error)
	DeleteUserRequest(*alexaforbusiness.DeleteUserInput) (*request.Request, *alexaforbusiness.DeleteUserOutput)

	DisassociateContactFromAddressBook(*alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateContactFromAddressBookWithContext(aws.Context, *alexaforbusiness.DisassociateContactFromAddressBookInput, ...request.Option) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateContactFromAddressBookRequest(*alexaforbusiness.DisassociateContactFromAddressBookInput) (*request.Request, *alexaforbusiness.DisassociateContactFromAddressBookOutput)

	DisassociateDeviceFromRoom(*alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateDeviceFromRoomWithContext(aws.Context, *alexaforbusiness.DisassociateDeviceFromRoomInput, ...request.Option) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateDeviceFromRoomRequest(*alexaforbusiness.DisassociateDeviceFromRoomInput) (*request.Request, *alexaforbusiness.DisassociateDeviceFromRoomOutput)

	DisassociateSkillFromSkillGroup(*alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromSkillGroupWithContext(aws.Context, *alexaforbusiness.DisassociateSkillFromSkillGroupInput, ...request.Option) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromSkillGroupRequest(*alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*request.Request, *alexaforbusiness.DisassociateSkillFromSkillGroupOutput)

	DisassociateSkillGroupFromRoom(*alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	DisassociateSkillGroupFromRoomWithContext(aws.Context, *alexaforbusiness.DisassociateSkillGroupFromRoomInput, ...request.Option) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	DisassociateSkillGroupFromRoomRequest(*alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*request.Request, *alexaforbusiness.DisassociateSkillGroupFromRoomOutput)

	ForgetSmartHomeAppliances(*alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	ForgetSmartHomeAppliancesWithContext(aws.Context, *alexaforbusiness.ForgetSmartHomeAppliancesInput, ...request.Option) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	ForgetSmartHomeAppliancesRequest(*alexaforbusiness.ForgetSmartHomeAppliancesInput) (*request.Request, *alexaforbusiness.ForgetSmartHomeAppliancesOutput)

	GetAddressBook(*alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error)
	GetAddressBookWithContext(aws.Context, *alexaforbusiness.GetAddressBookInput, ...request.Option) (*alexaforbusiness.GetAddressBookOutput, error)
	GetAddressBookRequest(*alexaforbusiness.GetAddressBookInput) (*request.Request, *alexaforbusiness.GetAddressBookOutput)

	GetConferencePreference(*alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferencePreferenceWithContext(aws.Context, *alexaforbusiness.GetConferencePreferenceInput, ...request.Option) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferencePreferenceRequest(*alexaforbusiness.GetConferencePreferenceInput) (*request.Request, *alexaforbusiness.GetConferencePreferenceOutput)

	GetConferenceProvider(*alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetConferenceProviderWithContext(aws.Context, *alexaforbusiness.GetConferenceProviderInput, ...request.Option) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetConferenceProviderRequest(*alexaforbusiness.GetConferenceProviderInput) (*request.Request, *alexaforbusiness.GetConferenceProviderOutput)

	GetContact(*alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error)
	GetContactWithContext(aws.Context, *alexaforbusiness.GetContactInput, ...request.Option) (*alexaforbusiness.GetContactOutput, error)
	GetContactRequest(*alexaforbusiness.GetContactInput) (*request.Request, *alexaforbusiness.GetContactOutput)

	GetDevice(*alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error)
	GetDeviceWithContext(aws.Context, *alexaforbusiness.GetDeviceInput, ...request.Option) (*alexaforbusiness.GetDeviceOutput, error)
	GetDeviceRequest(*alexaforbusiness.GetDeviceInput) (*request.Request, *alexaforbusiness.GetDeviceOutput)

	GetProfile(*alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error)
	GetProfileWithContext(aws.Context, *alexaforbusiness.GetProfileInput, ...request.Option) (*alexaforbusiness.GetProfileOutput, error)
	GetProfileRequest(*alexaforbusiness.GetProfileInput) (*request.Request, *alexaforbusiness.GetProfileOutput)

	GetRoom(*alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomWithContext(aws.Context, *alexaforbusiness.GetRoomInput, ...request.Option) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomRequest(*alexaforbusiness.GetRoomInput) (*request.Request, *alexaforbusiness.GetRoomOutput)

	GetRoomSkillParameter(*alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetRoomSkillParameterWithContext(aws.Context, *alexaforbusiness.GetRoomSkillParameterInput, ...request.Option) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetRoomSkillParameterRequest(*alexaforbusiness.GetRoomSkillParameterInput) (*request.Request, *alexaforbusiness.GetRoomSkillParameterOutput)

	GetSkillGroup(*alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error)
	GetSkillGroupWithContext(aws.Context, *alexaforbusiness.GetSkillGroupInput, ...request.Option) (*alexaforbusiness.GetSkillGroupOutput, error)
	GetSkillGroupRequest(*alexaforbusiness.GetSkillGroupInput) (*request.Request, *alexaforbusiness.GetSkillGroupOutput)

	ListBusinessReportSchedules(*alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListBusinessReportSchedulesWithContext(aws.Context, *alexaforbusiness.ListBusinessReportSchedulesInput, ...request.Option) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListBusinessReportSchedulesRequest(*alexaforbusiness.ListBusinessReportSchedulesInput) (*request.Request, *alexaforbusiness.ListBusinessReportSchedulesOutput)

	ListBusinessReportSchedulesPages(*alexaforbusiness.ListBusinessReportSchedulesInput, func(*alexaforbusiness.ListBusinessReportSchedulesOutput, bool) bool) error
	ListBusinessReportSchedulesPagesWithContext(aws.Context, *alexaforbusiness.ListBusinessReportSchedulesInput, func(*alexaforbusiness.ListBusinessReportSchedulesOutput, bool) bool, ...request.Option) error

	ListConferenceProviders(*alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListConferenceProvidersWithContext(aws.Context, *alexaforbusiness.ListConferenceProvidersInput, ...request.Option) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListConferenceProvidersRequest(*alexaforbusiness.ListConferenceProvidersInput) (*request.Request, *alexaforbusiness.ListConferenceProvidersOutput)

	ListConferenceProvidersPages(*alexaforbusiness.ListConferenceProvidersInput, func(*alexaforbusiness.ListConferenceProvidersOutput, bool) bool) error
	ListConferenceProvidersPagesWithContext(aws.Context, *alexaforbusiness.ListConferenceProvidersInput, func(*alexaforbusiness.ListConferenceProvidersOutput, bool) bool, ...request.Option) error

	ListDeviceEvents(*alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListDeviceEventsWithContext(aws.Context, *alexaforbusiness.ListDeviceEventsInput, ...request.Option) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListDeviceEventsRequest(*alexaforbusiness.ListDeviceEventsInput) (*request.Request, *alexaforbusiness.ListDeviceEventsOutput)

	ListDeviceEventsPages(*alexaforbusiness.ListDeviceEventsInput, func(*alexaforbusiness.ListDeviceEventsOutput, bool) bool) error
	ListDeviceEventsPagesWithContext(aws.Context, *alexaforbusiness.ListDeviceEventsInput, func(*alexaforbusiness.ListDeviceEventsOutput, bool) bool, ...request.Option) error

	ListSkills(*alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsWithContext(aws.Context, *alexaforbusiness.ListSkillsInput, ...request.Option) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsRequest(*alexaforbusiness.ListSkillsInput) (*request.Request, *alexaforbusiness.ListSkillsOutput)

	ListSkillsPages(*alexaforbusiness.ListSkillsInput, func(*alexaforbusiness.ListSkillsOutput, bool) bool) error
	ListSkillsPagesWithContext(aws.Context, *alexaforbusiness.ListSkillsInput, func(*alexaforbusiness.ListSkillsOutput, bool) bool, ...request.Option) error

	ListSkillsStoreCategories(*alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreCategoriesWithContext(aws.Context, *alexaforbusiness.ListSkillsStoreCategoriesInput, ...request.Option) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreCategoriesRequest(*alexaforbusiness.ListSkillsStoreCategoriesInput) (*request.Request, *alexaforbusiness.ListSkillsStoreCategoriesOutput)

	ListSkillsStoreCategoriesPages(*alexaforbusiness.ListSkillsStoreCategoriesInput, func(*alexaforbusiness.ListSkillsStoreCategoriesOutput, bool) bool) error
	ListSkillsStoreCategoriesPagesWithContext(aws.Context, *alexaforbusiness.ListSkillsStoreCategoriesInput, func(*alexaforbusiness.ListSkillsStoreCategoriesOutput, bool) bool, ...request.Option) error

	ListSkillsStoreSkillsByCategory(*alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSkillsStoreSkillsByCategoryWithContext(aws.Context, *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, ...request.Option) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSkillsStoreSkillsByCategoryRequest(*alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*request.Request, *alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput)

	ListSkillsStoreSkillsByCategoryPages(*alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, func(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, bool) bool) error
	ListSkillsStoreSkillsByCategoryPagesWithContext(aws.Context, *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, func(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, bool) bool, ...request.Option) error

	ListSmartHomeAppliances(*alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListSmartHomeAppliancesWithContext(aws.Context, *alexaforbusiness.ListSmartHomeAppliancesInput, ...request.Option) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListSmartHomeAppliancesRequest(*alexaforbusiness.ListSmartHomeAppliancesInput) (*request.Request, *alexaforbusiness.ListSmartHomeAppliancesOutput)

	ListSmartHomeAppliancesPages(*alexaforbusiness.ListSmartHomeAppliancesInput, func(*alexaforbusiness.ListSmartHomeAppliancesOutput, bool) bool) error
	ListSmartHomeAppliancesPagesWithContext(aws.Context, *alexaforbusiness.ListSmartHomeAppliancesInput, func(*alexaforbusiness.ListSmartHomeAppliancesOutput, bool) bool, ...request.Option) error

	ListTags(*alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *alexaforbusiness.ListTagsInput, ...request.Option) (*alexaforbusiness.ListTagsOutput, error)
	ListTagsRequest(*alexaforbusiness.ListTagsInput) (*request.Request, *alexaforbusiness.ListTagsOutput)

	ListTagsPages(*alexaforbusiness.ListTagsInput, func(*alexaforbusiness.ListTagsOutput, bool) bool) error
	ListTagsPagesWithContext(aws.Context, *alexaforbusiness.ListTagsInput, func(*alexaforbusiness.ListTagsOutput, bool) bool, ...request.Option) error

	PutConferencePreference(*alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutConferencePreferenceWithContext(aws.Context, *alexaforbusiness.PutConferencePreferenceInput, ...request.Option) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutConferencePreferenceRequest(*alexaforbusiness.PutConferencePreferenceInput) (*request.Request, *alexaforbusiness.PutConferencePreferenceOutput)

	PutRoomSkillParameter(*alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutRoomSkillParameterWithContext(aws.Context, *alexaforbusiness.PutRoomSkillParameterInput, ...request.Option) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutRoomSkillParameterRequest(*alexaforbusiness.PutRoomSkillParameterInput) (*request.Request, *alexaforbusiness.PutRoomSkillParameterOutput)

	PutSkillAuthorization(*alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	PutSkillAuthorizationWithContext(aws.Context, *alexaforbusiness.PutSkillAuthorizationInput, ...request.Option) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	PutSkillAuthorizationRequest(*alexaforbusiness.PutSkillAuthorizationInput) (*request.Request, *alexaforbusiness.PutSkillAuthorizationOutput)

	RegisterAVSDevice(*alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RegisterAVSDeviceWithContext(aws.Context, *alexaforbusiness.RegisterAVSDeviceInput, ...request.Option) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RegisterAVSDeviceRequest(*alexaforbusiness.RegisterAVSDeviceInput) (*request.Request, *alexaforbusiness.RegisterAVSDeviceOutput)

	RejectSkill(*alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error)
	RejectSkillWithContext(aws.Context, *alexaforbusiness.RejectSkillInput, ...request.Option) (*alexaforbusiness.RejectSkillOutput, error)
	RejectSkillRequest(*alexaforbusiness.RejectSkillInput) (*request.Request, *alexaforbusiness.RejectSkillOutput)

	ResolveRoom(*alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error)
	ResolveRoomWithContext(aws.Context, *alexaforbusiness.ResolveRoomInput, ...request.Option) (*alexaforbusiness.ResolveRoomOutput, error)
	ResolveRoomRequest(*alexaforbusiness.ResolveRoomInput) (*request.Request, *alexaforbusiness.ResolveRoomOutput)

	RevokeInvitation(*alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error)
	RevokeInvitationWithContext(aws.Context, *alexaforbusiness.RevokeInvitationInput, ...request.Option) (*alexaforbusiness.RevokeInvitationOutput, error)
	RevokeInvitationRequest(*alexaforbusiness.RevokeInvitationInput) (*request.Request, *alexaforbusiness.RevokeInvitationOutput)

	SearchAddressBooks(*alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchAddressBooksWithContext(aws.Context, *alexaforbusiness.SearchAddressBooksInput, ...request.Option) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchAddressBooksRequest(*alexaforbusiness.SearchAddressBooksInput) (*request.Request, *alexaforbusiness.SearchAddressBooksOutput)

	SearchAddressBooksPages(*alexaforbusiness.SearchAddressBooksInput, func(*alexaforbusiness.SearchAddressBooksOutput, bool) bool) error
	SearchAddressBooksPagesWithContext(aws.Context, *alexaforbusiness.SearchAddressBooksInput, func(*alexaforbusiness.SearchAddressBooksOutput, bool) bool, ...request.Option) error

	SearchContacts(*alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error)
	SearchContactsWithContext(aws.Context, *alexaforbusiness.SearchContactsInput, ...request.Option) (*alexaforbusiness.SearchContactsOutput, error)
	SearchContactsRequest(*alexaforbusiness.SearchContactsInput) (*request.Request, *alexaforbusiness.SearchContactsOutput)

	SearchContactsPages(*alexaforbusiness.SearchContactsInput, func(*alexaforbusiness.SearchContactsOutput, bool) bool) error
	SearchContactsPagesWithContext(aws.Context, *alexaforbusiness.SearchContactsInput, func(*alexaforbusiness.SearchContactsOutput, bool) bool, ...request.Option) error

	SearchDevices(*alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchDevicesWithContext(aws.Context, *alexaforbusiness.SearchDevicesInput, ...request.Option) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchDevicesRequest(*alexaforbusiness.SearchDevicesInput) (*request.Request, *alexaforbusiness.SearchDevicesOutput)

	SearchDevicesPages(*alexaforbusiness.SearchDevicesInput, func(*alexaforbusiness.SearchDevicesOutput, bool) bool) error
	SearchDevicesPagesWithContext(aws.Context, *alexaforbusiness.SearchDevicesInput, func(*alexaforbusiness.SearchDevicesOutput, bool) bool, ...request.Option) error

	SearchProfiles(*alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchProfilesWithContext(aws.Context, *alexaforbusiness.SearchProfilesInput, ...request.Option) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchProfilesRequest(*alexaforbusiness.SearchProfilesInput) (*request.Request, *alexaforbusiness.SearchProfilesOutput)

	SearchProfilesPages(*alexaforbusiness.SearchProfilesInput, func(*alexaforbusiness.SearchProfilesOutput, bool) bool) error
	SearchProfilesPagesWithContext(aws.Context, *alexaforbusiness.SearchProfilesInput, func(*alexaforbusiness.SearchProfilesOutput, bool) bool, ...request.Option) error

	SearchRooms(*alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchRoomsWithContext(aws.Context, *alexaforbusiness.SearchRoomsInput, ...request.Option) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchRoomsRequest(*alexaforbusiness.SearchRoomsInput) (*request.Request, *alexaforbusiness.SearchRoomsOutput)

	SearchRoomsPages(*alexaforbusiness.SearchRoomsInput, func(*alexaforbusiness.SearchRoomsOutput, bool) bool) error
	SearchRoomsPagesWithContext(aws.Context, *alexaforbusiness.SearchRoomsInput, func(*alexaforbusiness.SearchRoomsOutput, bool) bool, ...request.Option) error

	SearchSkillGroups(*alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchSkillGroupsWithContext(aws.Context, *alexaforbusiness.SearchSkillGroupsInput, ...request.Option) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchSkillGroupsRequest(*alexaforbusiness.SearchSkillGroupsInput) (*request.Request, *alexaforbusiness.SearchSkillGroupsOutput)

	SearchSkillGroupsPages(*alexaforbusiness.SearchSkillGroupsInput, func(*alexaforbusiness.SearchSkillGroupsOutput, bool) bool) error
	SearchSkillGroupsPagesWithContext(aws.Context, *alexaforbusiness.SearchSkillGroupsInput, func(*alexaforbusiness.SearchSkillGroupsOutput, bool) bool, ...request.Option) error

	SearchUsers(*alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error)
	SearchUsersWithContext(aws.Context, *alexaforbusiness.SearchUsersInput, ...request.Option) (*alexaforbusiness.SearchUsersOutput, error)
	SearchUsersRequest(*alexaforbusiness.SearchUsersInput) (*request.Request, *alexaforbusiness.SearchUsersOutput)

	SearchUsersPages(*alexaforbusiness.SearchUsersInput, func(*alexaforbusiness.SearchUsersOutput, bool) bool) error
	SearchUsersPagesWithContext(aws.Context, *alexaforbusiness.SearchUsersInput, func(*alexaforbusiness.SearchUsersOutput, bool) bool, ...request.Option) error

	SendInvitation(*alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error)
	SendInvitationWithContext(aws.Context, *alexaforbusiness.SendInvitationInput, ...request.Option) (*alexaforbusiness.SendInvitationOutput, error)
	SendInvitationRequest(*alexaforbusiness.SendInvitationInput) (*request.Request, *alexaforbusiness.SendInvitationOutput)

	StartDeviceSync(*alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartDeviceSyncWithContext(aws.Context, *alexaforbusiness.StartDeviceSyncInput, ...request.Option) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartDeviceSyncRequest(*alexaforbusiness.StartDeviceSyncInput) (*request.Request, *alexaforbusiness.StartDeviceSyncOutput)

	StartSmartHomeApplianceDiscovery(*alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	StartSmartHomeApplianceDiscoveryWithContext(aws.Context, *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput, ...request.Option) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	StartSmartHomeApplianceDiscoveryRequest(*alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*request.Request, *alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput)

	TagResource(*alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *alexaforbusiness.TagResourceInput, ...request.Option) (*alexaforbusiness.TagResourceOutput, error)
	TagResourceRequest(*alexaforbusiness.TagResourceInput) (*request.Request, *alexaforbusiness.TagResourceOutput)

	UntagResource(*alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *alexaforbusiness.UntagResourceInput, ...request.Option) (*alexaforbusiness.UntagResourceOutput, error)
	UntagResourceRequest(*alexaforbusiness.UntagResourceInput) (*request.Request, *alexaforbusiness.UntagResourceOutput)

	UpdateAddressBook(*alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateAddressBookWithContext(aws.Context, *alexaforbusiness.UpdateAddressBookInput, ...request.Option) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateAddressBookRequest(*alexaforbusiness.UpdateAddressBookInput) (*request.Request, *alexaforbusiness.UpdateAddressBookOutput)

	UpdateBusinessReportSchedule(*alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateBusinessReportScheduleWithContext(aws.Context, *alexaforbusiness.UpdateBusinessReportScheduleInput, ...request.Option) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateBusinessReportScheduleRequest(*alexaforbusiness.UpdateBusinessReportScheduleInput) (*request.Request, *alexaforbusiness.UpdateBusinessReportScheduleOutput)

	UpdateConferenceProvider(*alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateConferenceProviderWithContext(aws.Context, *alexaforbusiness.UpdateConferenceProviderInput, ...request.Option) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateConferenceProviderRequest(*alexaforbusiness.UpdateConferenceProviderInput) (*request.Request, *alexaforbusiness.UpdateConferenceProviderOutput)

	UpdateContact(*alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateContactWithContext(aws.Context, *alexaforbusiness.UpdateContactInput, ...request.Option) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateContactRequest(*alexaforbusiness.UpdateContactInput) (*request.Request, *alexaforbusiness.UpdateContactOutput)

	UpdateDevice(*alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateDeviceWithContext(aws.Context, *alexaforbusiness.UpdateDeviceInput, ...request.Option) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateDeviceRequest(*alexaforbusiness.UpdateDeviceInput) (*request.Request, *alexaforbusiness.UpdateDeviceOutput)

	UpdateProfile(*alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateProfileWithContext(aws.Context, *alexaforbusiness.UpdateProfileInput, ...request.Option) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateProfileRequest(*alexaforbusiness.UpdateProfileInput) (*request.Request, *alexaforbusiness.UpdateProfileOutput)

	UpdateRoom(*alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateRoomWithContext(aws.Context, *alexaforbusiness.UpdateRoomInput, ...request.Option) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateRoomRequest(*alexaforbusiness.UpdateRoomInput) (*request.Request, *alexaforbusiness.UpdateRoomOutput)

	UpdateSkillGroup(*alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error)
	UpdateSkillGroupWithContext(aws.Context, *alexaforbusiness.UpdateSkillGroupInput, ...request.Option) (*alexaforbusiness.UpdateSkillGroupOutput, error)
	UpdateSkillGroupRequest(*alexaforbusiness.UpdateSkillGroupInput) (*request.Request, *alexaforbusiness.UpdateSkillGroupOutput)
}

var _ AlexaForBusinessAPI = (*alexaforbusiness.AlexaForBusiness)(nil)
