package channel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	data := []struct {
		referrer     string
		referrerName string
		utmMedium    string
		utmCampaign  string
		utmSource    string
		clickID      string
		channel      string
	}{
		{channel: "Direct"},
		{channel: "Cross-network", utmCampaign: "foo cross-Network bar"},
		{channel: "Paid Shopping", referrer: "https://walmart.com", utmMedium: "cp"},
		{channel: "Paid Shopping", referrer: "https://walmart.com", utmMedium: "ppc"},
		{channel: "Paid Shopping", referrer: "https://walmart.com", utmMedium: "retargeting"},
		{channel: "Paid Shopping", referrer: "https://walmart.com", utmMedium: "paid"},
		{channel: "Paid Shopping", utmCampaign: "Shop Campaign", utmMedium: "paid"},
		{channel: "Paid Shopping", utmCampaign: "Shopping Campaign", utmMedium: "paid"},
		{channel: "Paid Search", referrer: "https://360.cn", utmMedium: "cp"},
		{channel: "Paid Search", referrer: "https://360.cn", utmMedium: "ppc"},
		{channel: "Paid Search", referrer: "https://360.cn", utmMedium: "retargeting"},
		{channel: "Paid Search", referrer: "https://360.cn", utmMedium: "paid"},
		{channel: "Paid Social", referrer: "https://43things.com", utmMedium: "cp"},
		{channel: "Paid Social", referrer: "https://43things.com", utmMedium: "ppc"},
		{channel: "Paid Social", referrer: "https://43things.com", utmMedium: "retargeting"},
		{channel: "Paid Social", referrer: "43things.com", utmMedium: "paid"},
		{channel: "Paid Video", referrer: "https://twitch.tv", utmMedium: "cp"},
		{channel: "Paid Video", referrer: "https://twitch.tv", utmMedium: "ppc"},
		{channel: "Paid Video", referrer: "https://twitch.tv", utmMedium: "retargeting"},
		{channel: "Paid Video", referrer: "twitch.tv", utmMedium: "paid"},
		{channel: "Display", utmMedium: "display"},
		{channel: "Display", utmMedium: "banner"},
		{channel: "Display", utmMedium: "expandable"},
		{channel: "Display", utmMedium: "interstitial"},
		{channel: "Display", utmMedium: "cpm"},
		{channel: "Paid Other", utmMedium: "foo cp bar"},
		{channel: "Paid Other", utmMedium: "ppc"},
		{channel: "Paid Other", utmMedium: "retargeting"},
		{channel: "Paid Other", utmMedium: "paid bar"},
		{channel: "Direct", utmMedium: "foo ppc"},
		{channel: "Direct", utmMedium: "bar retargeting"},
		{channel: "Organic Shopping", referrer: "https://uk.shopping.net", utmCampaign: "Shop"},
		{channel: "Organic Shopping", referrer: "https://uk.shopping.net", utmCampaign: "shopping campaign"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "social"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "social-network"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "social-media"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "sm"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "social network"},
		{channel: "Organic Social", referrer: "https://unblog.fr", utmMedium: "social media"},
		{channel: "Organic Video", referrer: "https://utreon.com", utmMedium: "video foo"},
		{channel: "Organic Video", referrer: "https://utreon.com", utmMedium: "bar video"},
		{channel: "Organic Search", referrer: "https://sogou.com", utmMedium: "organic"},
		{channel: "Referral", utmMedium: "referral"},
		{channel: "Referral", utmMedium: "app"},
		{channel: "Referral", utmMedium: "link"},
		{channel: "Email", referrer: "email"},
		{channel: "Email", referrer: "e-mail"},
		{channel: "Email", referrer: "e_mail"},
		{channel: "Email", referrer: "e mail"},
		{channel: "Email", utmSource: "email"},
		{channel: "Email", utmSource: "e-mail"},
		{channel: "Email", utmSource: "e_mail"},
		{channel: "Email", utmSource: "e mail"},
		{channel: "Email", utmMedium: "email"},
		{channel: "Email", utmMedium: "e-mail"},
		{channel: "Email", utmMedium: "e_mail"},
		{channel: "Email", utmMedium: "e mail"},
		{channel: "Affiliates", utmMedium: "affiliate"},
		{channel: "Audio", utmMedium: "audio"},
		{channel: "SMS", referrer: "sms"},
		{channel: "SMS", utmSource: "sms"},
		{channel: "SMS", utmMedium: "sms"},
		{channel: "Mobile Push Notifications", utmMedium: "this-is-a-push"},
		{channel: "Mobile Push Notifications", utmMedium: "my mobile notification"},
		{channel: "Mobile Push Notifications", utmMedium: "my app notification"},
		{channel: "Mobile Push Notifications", referrer: "firebase"},
		{channel: "Mobile Push Notifications", utmSource: "firebase"},
	}

	for _, d := range data {
		c := Get(d.referrer, d.referrerName, d.utmMedium, d.utmCampaign, d.utmSource, d.clickID)
		assert.Equal(t, d.channel, c, "%s, %s, %s, %s, %s, %s -> %s", d.referrer, d.referrerName, d.utmMedium, d.utmCampaign, d.utmSource, d.clickID, c)
	}
}