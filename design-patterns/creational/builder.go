package creational

import "fmt"

/*
# Builder Pattern

## Concept
Consists in the constructing of complex objects step by step. The resulting
objects (product) does not need to have a common interface, they'll just share
the same building process.

## Example explanation
In this pattern we create a builder interface of setters and a getter to return
the built object. We also create a getBuilder function to instantiate concrete
builders that implements the builder interface. Finally, we need to create the
actual blueprint of the final object.
Optionally, we can create a Director to receive a builder and manage
it's creation (set all of it's properties and get the final object).
*/

// ##### Builder Interface #####
type IBuilder interface {
	setChannelType(t string)
	setFeatures(feat []string)
	setBannedContacts(bans []uint64)
	getFlow() Flow
}

func getBuilder(builderType string) (IBuilder, error) {
	if builderType == "messaging" {
		return newMessagingFlowBuilder(), nil
	}
	if builderType == "voice" {
		return newVoiceFlowBuilder(), nil
	}

	return nil, fmt.Errorf("invalid builder type passed")
}

// ##### Messaging Concrete Builder #####
type MessagingFlowBuilder struct {
	channelType    string
	features       []string
	bannedContacts []uint64
}

func newMessagingFlowBuilder() *MessagingFlowBuilder {
	return &MessagingFlowBuilder{}
}

func (b *MessagingFlowBuilder) setChannelType(t string) {
	b.channelType = t
}

func (b *MessagingFlowBuilder) setFeatures(feats []string) {
	b.features = feats
}

func (b *MessagingFlowBuilder) setBannedContacts(bans []uint64) {
	b.bannedContacts = bans
}

func (b *MessagingFlowBuilder) getFlow() Flow {
	return Flow{
		channelType:    b.channelType,
		features:       b.features,
		bannedContacts: b.bannedContacts,
	}
}

// ##### Voice Concrete Builder #####
type VoiceFlowBuilder struct {
	channelType    string
	features       []string
	bannedContacts []uint64
}

func newVoiceFlowBuilder() *VoiceFlowBuilder {
	return &VoiceFlowBuilder{}
}

func (b *VoiceFlowBuilder) setChannelType(t string) {
	b.channelType = t
}

func (b *VoiceFlowBuilder) setFeatures(feats []string) {
	b.features = feats
}

func (b *VoiceFlowBuilder) setBannedContacts(bans []uint64) {
	b.bannedContacts = bans
}

func (b *VoiceFlowBuilder) getFlow() Flow {
	return Flow{
		channelType:    b.channelType,
		features:       b.features,
		bannedContacts: b.bannedContacts,
	}
}

// ##### Product #####
type Flow struct {
	channelType    string
	features       []string
	bannedContacts []uint64
}

// ##### Director #####
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildFlow(t string, f []string, b []uint64) Flow {
	d.builder.setChannelType(t)
	d.builder.setFeatures(f)
	d.builder.setBannedContacts(b)
	return d.builder.getFlow()
}

// ##### Client Code #####
func RunBuilderExample() {
	messagingBuilder, _ := getBuilder("messaging")
	voiceBuilder, _ := getBuilder("voice")

	director := newDirector(messagingBuilder)
	messagingFlow := director.buildFlow("whatsapp", []string{"menu", "send media", "send audio"}, []uint64{123456, 654321})

	fmt.Printf("Messaging Flow channel type: %s\n", messagingFlow.channelType)
	fmt.Printf("Messaging Flow features: %s\n", messagingFlow.features)
	fmt.Printf("Messaging Flow bannedContacts: %v\n", messagingFlow.bannedContacts)

	director.setBuilder(voiceBuilder)
	voiceFlow := director.buildFlow("call", []string{"dialmenu", "queue_song"}, []uint64{123456, 654321})

	fmt.Printf("Voice Flow channel type: %s\n", voiceFlow.channelType)
	fmt.Printf("Voice Flow features: %s\n", voiceFlow.features)
	fmt.Printf("Voice Flow bannedContacts: %v\n", voiceFlow.bannedContacts)
}
