package eventmanager

import (
	"github.com/Confialink/wallet-notifications/internal/db"
	changepassword "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/change-password"
	event_dormant_profile_admin "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/dormant_profile_admin"
	email_verification "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/email-verification"
	failedloginattempts "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/failed-login-attempts"
	incomingmessage "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/incoming-message"
	incomingmessageadmins "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/incoming-message-admins"
	incomingtransaction "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/incoming-transaction"
	invoiceduedate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/invoice-due-date"
	invoicehasbeentraded "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/invoice-has-been-traded"
	invoiceoverdue1 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/invoice-overdue1"
	invoiceoverdue3 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/invoice-overdue3"
	newinvoicesuploaded "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/new-invoices-uploaded"
	upcominginvoicepayment10 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/upcoming-invoice-payment10"
	upcominginvoicepayment5 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/buyer/upcoming-invoice-payment5"
	funderinvoiceduedate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-invoice-due-date"
	funderinvoicematurityin5days "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-invoice-maturity-in-5-days"
	funderinvoiceoverdue1 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-invoice-overdue-1"
	funderinvoiceoverdue3 "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-invoice-overdue-3"
	fundernewinvoicefinancerequest "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-new-invoice-finance-request"
	fundersupplychainfinance "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/funder/funder-supply-chain-finance"
	suppliercreditmemoreceivedfrombuyer "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-credit-memo-received-from-buyer"
	supplierinvoicehasbeentraded "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-invoice-has-been-traded"
	supplierinvoiceoffersenttofunder "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-invoice-offer-sent-to-funder"
	supplierinvoicesaleofferrejected "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-invoice-sale-offer-rejected"
	suppliernewapprovedinvoiceshavebeenuploaded "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-new-approved-invoices-have-been-uploaded"
	suppliernewinvoicesaleoffer "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-new-invoice-sale-offer"
	supplieroutgoingwiretransferhasbeenexecuted "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-outgoing-wire-transfer-has-been-executed"
	supplieroutgoingwiretransferrequesthasbeenreceived "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/invoices/supplier/supplier-outgoing-wire-transfer-request-has-been-received"
	newtransferrequest "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/new-transfer-request"
	newtransferrequestbyadmin "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/new-transfer-request-by-admin"
	"github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/newmoneyrequest"
	newspublished "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/news-publiched"
	outgoingtransaction "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/outgoing-transaction"
	password_recovery "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/password-recovery"
	phone_verification "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/phone-verification"
	profileactivate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/profile-activate"
	profileactivateviaothersources "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/profile-activate-via-other-sources"
	profilecancel "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/profile-cancel"
	profile_confirmation "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/profile-confirmation"
	profilecreate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/profile-create"
	requestcancel "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/request-cancel"
	requestexecuted "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/request-executed"
	tancreate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/tan-create"
	testsmtp "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/test-smtp"
	verificationcreate "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/verification-create"
	verificationfailed "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/verification-failed"
	verificationsuccess "github.com/Confialink/wallet-notifications/internal/eventmanager/subscribers/verification-success"
	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/inconshreveable/log15"
)

// EventManager is event manager structure
type EventManager struct {
	Storage    Storage
	Dispatcher DispatcherContract
	notifier   *service.Notifier
	repo       db.RepositoryInterface
	logger     log15.Logger
}

// Attach attaches a subscriber to an event
func (e *EventManager) Attach(eventName string, sub service.Subscriber) {
	e.Storage.Attach(eventName, sub)
}

// Dispatch dispatches the event across all the subscribers
func (e *EventManager) Dispatch(eventName string, data service.CallData) {
	e.logger.Info("Event dispatching", "eventName", eventName)
	s := e.Storage.Subscribers(eventName)
	e.Dispatcher.Dispatch(eventName, data, s)
}

// Detach de attaches a subscriber from an event
func (e *EventManager) Detach(eventName string, subscriber service.Subscriber) {
	e.Storage.Detach(eventName, subscriber)
}

func (e *EventManager) Listen() {
	e.Attach(event_dormant_profile_admin.EventName, event_dormant_profile_admin.New(e.notifier, e.repo))
	e.Attach(profileactivate.EventName, profileactivate.New(e.notifier, e.repo))
	e.Attach(profilecancel.EventName, profilecancel.New(e.notifier, e.repo))
	e.Attach(profilecreate.EventName, profilecreate.New(e.notifier, e.repo))
	e.Attach(testsmtp.EventName, testsmtp.New(e.notifier, e.repo))
	e.Attach(changepassword.EventName, changepassword.New(e.notifier, e.repo))
	e.Attach(tancreate.EventName, tancreate.New(e.notifier, e.repo))
	e.Attach(incomingtransaction.EventName, incomingtransaction.New(e.notifier, e.repo))
	e.Attach(outgoingtransaction.EventName, outgoingtransaction.New(e.notifier, e.repo))
	e.Attach(newtransferrequest.EventName, newtransferrequest.New(e.notifier, e.repo))
	e.Attach(newtransferrequestbyadmin.EventName, newtransferrequestbyadmin.New(e.notifier, e.repo))
	e.Attach(incomingmessage.EventName, incomingmessage.New(e.notifier, e.repo))
	e.Attach(incomingmessageadmins.EventName, incomingmessageadmins.New(e.notifier, e.repo))
	e.Attach(newspublished.EventName, newspublished.New(e.notifier, e.repo))
	e.Attach(failedloginattempts.EventName, failedloginattempts.New(e.notifier, e.repo))
	e.Attach(verificationcreate.EventName, verificationcreate.New(e.notifier, e.repo))
	e.Attach(verificationfailed.EventName, verificationfailed.New(e.notifier, e.repo))
	e.Attach(verificationsuccess.EventName, verificationsuccess.New(e.notifier, e.repo))
	e.Attach(profileactivateviaothersources.EventName, profileactivateviaothersources.New(e.notifier, e.repo))
	e.Attach(password_recovery.EventName, password_recovery.New(e.notifier, e.repo))
	e.Attach(profile_confirmation.EventName, profile_confirmation.New(e.notifier, e.repo))
	e.Attach(requestexecuted.EventName, requestexecuted.New(e.notifier, e.logger.New("event", requestexecuted.EventName), e.repo))
	e.Attach(requestcancel.EventName, requestcancel.New(e.notifier, e.logger.New("event", requestcancel.EventName), e.repo))
	e.Attach(phone_verification.EventName, phone_verification.New(e.notifier, e.logger.New("event", phone_verification.EventName), e.repo))
	e.Attach(email_verification.EventName, email_verification.New(e.notifier, e.logger.New("event", email_verification.EventName), e.repo))
	e.Attach(newinvoicesuploaded.EventName, newinvoicesuploaded.New(e.notifier, e.logger.New("event", newinvoicesuploaded.EventName), e.repo))
	e.Attach(invoiceduedate.EventName, invoiceduedate.New(e.notifier, e.logger.New("event", invoiceduedate.EventName), e.repo))
	e.Attach(invoicehasbeentraded.EventName, invoicehasbeentraded.New(e.notifier, e.logger.New("event", invoicehasbeentraded.EventName), e.repo))
	e.Attach(invoiceoverdue1.EventName, invoiceoverdue1.New(e.notifier, e.logger.New("event", invoiceoverdue1.EventName), e.repo))
	e.Attach(invoiceoverdue3.EventName, invoiceoverdue3.New(e.notifier, e.logger.New("event", invoiceoverdue3.EventName), e.repo))
	e.Attach(upcominginvoicepayment5.EventName, upcominginvoicepayment5.New(e.notifier, e.logger.New("event", upcominginvoicepayment5.EventName), e.repo))
	e.Attach(upcominginvoicepayment10.EventName, upcominginvoicepayment10.New(e.notifier, e.logger.New("event", upcominginvoicepayment10.EventName), e.repo))
	e.Attach(funderinvoiceduedate.EventName, funderinvoiceduedate.New(e.notifier, e.logger.New("event", funderinvoiceduedate.EventName), e.repo))
	e.Attach(funderinvoicematurityin5days.EventName, funderinvoicematurityin5days.New(e.notifier, e.logger.New("event", funderinvoicematurityin5days.EventName), e.repo))
	e.Attach(funderinvoiceoverdue1.EventName, funderinvoiceoverdue1.New(e.notifier, e.logger.New("event", funderinvoiceoverdue1.EventName), e.repo))
	e.Attach(funderinvoiceoverdue3.EventName, funderinvoiceoverdue3.New(e.notifier, e.logger.New("event", funderinvoiceoverdue3.EventName), e.repo))
	e.Attach(fundernewinvoicefinancerequest.EventName, fundernewinvoicefinancerequest.New(e.notifier, e.logger.New("event", fundernewinvoicefinancerequest.EventName), e.repo))
	e.Attach(fundersupplychainfinance.EventName, fundersupplychainfinance.New(e.notifier, e.logger.New("event", fundersupplychainfinance.EventName), e.repo))
	e.Attach(suppliercreditmemoreceivedfrombuyer.EventName, suppliercreditmemoreceivedfrombuyer.New(e.notifier, e.logger.New("event", suppliercreditmemoreceivedfrombuyer.EventName), e.repo))
	e.Attach(supplierinvoicehasbeentraded.EventName, supplierinvoicehasbeentraded.New(e.notifier, e.logger.New("event", supplierinvoicehasbeentraded.EventName), e.repo))
	e.Attach(supplierinvoiceoffersenttofunder.EventName, supplierinvoiceoffersenttofunder.New(e.notifier, e.logger.New("event", supplierinvoiceoffersenttofunder.EventName), e.repo))
	e.Attach(supplierinvoicesaleofferrejected.EventName, supplierinvoicesaleofferrejected.New(e.notifier, e.logger.New("event", supplierinvoicesaleofferrejected.EventName), e.repo))
	e.Attach(suppliernewapprovedinvoiceshavebeenuploaded.EventName, suppliernewapprovedinvoiceshavebeenuploaded.New(e.notifier, e.logger.New("event", suppliernewapprovedinvoiceshavebeenuploaded.EventName), e.repo))
	e.Attach(suppliernewinvoicesaleoffer.EventName, suppliernewinvoicesaleoffer.New(e.notifier, e.logger.New("event", suppliernewinvoicesaleoffer.EventName), e.repo))
	e.Attach(supplieroutgoingwiretransferhasbeenexecuted.EventName, supplieroutgoingwiretransferhasbeenexecuted.New(e.notifier, e.logger.New("event", supplieroutgoingwiretransferhasbeenexecuted.EventName), e.repo))
	e.Attach(supplieroutgoingwiretransferrequesthasbeenreceived.EventName, supplieroutgoingwiretransferrequesthasbeenreceived.New(e.notifier, e.logger.New("event", supplieroutgoingwiretransferrequesthasbeenreceived.EventName), e.repo))
	e.Attach(newmoneyrequest.EventName, newmoneyrequest.New(e.notifier, e.logger.New("event", newmoneyrequest.EventName)))
}

// NewEventManager is factory method for the event manager
func NewEventManager(
	storage Storage,
	dispatcher DispatcherContract,
	notifier *service.Notifier,
	repo db.RepositoryInterface,
	logger log15.Logger,
) *EventManager {
	return &EventManager{
		storage,
		dispatcher,
		notifier,
		repo,
		logger,
	}
}
