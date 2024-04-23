package consumer

import (
	"context"

	amf_context "github.com/free5gc/amf/internal/context"
	"github.com/free5gc/amf/pkg/factory"
	"github.com/free5gc/openapi/Namf_Communication"
	"github.com/free5gc/openapi/Nausf_UEAuthentication"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
	"github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/Nnssf_NSSelection"
	"github.com/free5gc/openapi/Npcf_AMPolicy"
	"github.com/free5gc/openapi/Nsmf_PDUSession"
	"github.com/free5gc/openapi/Nudm_SubscriberDataManagement"
	"github.com/free5gc/openapi/Nudm_UEContextManagement"
)

type amf interface {
	Config() *factory.Config
	Context() *amf_context.AMFContext
	CancelContext() context.Context
}

type Consumer struct {
	amf

	// consumer services
	*namfService
	*nnrfService
	*npcfService
	*nssfService
	*nsmfService
	*nudmService
	*nausfService
}

var consumer *Consumer

func GetConsumer() *Consumer {
	return consumer
}

func NewConsumer(amf amf) (*Consumer, error) {
	c := &Consumer{
		amf: amf,
	}
	consumer = c

	c.namfService = &namfService{
		consumer:   c,
		ComClients: make(map[string]*Namf_Communication.APIClient),
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*Nnrf_NFManagement.APIClient),
		nfDiscClients:   make(map[string]*Nnrf_NFDiscovery.APIClient),
	}

	c.npcfService = &npcfService{
		consumer:        c,
		AMPolicyClients: make(map[string]*Npcf_AMPolicy.APIClient),
	}

	c.nssfService = &nssfService{
		consumer:           c,
		NSSelectionClients: make(map[string]*Nnssf_NSSelection.APIClient),
	}

	c.nsmfService = &nsmfService{
		consumer:          c,
		PDUSessionClients: make(map[string]*Nsmf_PDUSession.APIClient),
	}

	c.nudmService = &nudmService{
		consumer:                 c,
		SubscriberDMngmntClients: make(map[string]*Nudm_SubscriberDataManagement.APIClient),
		UEContextMngmntClients:   make(map[string]*Nudm_UEContextManagement.APIClient),
	}

	c.nausfService = &nausfService{
		consumer:                c,
		UEAuthenticationClients: make(map[string]*Nausf_UEAuthentication.APIClient),
	}
	return c, nil
}
