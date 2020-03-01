package scanners

import (
	"github.com/sundowndev/dorkgen"
	"github.com/sundowndev/phoneinfoga/pkg/utils"
)

// GoogleSearchDork is the common format for dork requests
type GoogleSearchDork struct {
	Number string `json:"number"`
	Dork   string `json:"dork"`
	URL    string
}

func getDisposableProvidersDorks(number *Number) []*dorkgen.GoogleSearch {
	return []*dorkgen.GoogleSearch{
		(&dorkgen.GoogleSearch{}).
			Site("hs3x.com").
			Intext(number.International),
		(&dorkgen.GoogleSearch{}).
			Site("receive-sms-now.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("smslisten.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("smsnumbersonline.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("freesmscode.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("catchsms.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("smstibo.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("smsreceiving.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("getfreesmsnumber.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("sellaite.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receive-sms-online.info").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receivesmsonline.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receive-a-sms.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("sms-receive.net").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receivefreesms.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receive-sms.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receivetxt.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("freephonenum.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("freesmsverification.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("receive-sms-online.com").
			Intext(number.International).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("smslive.co").
			Intext(number.International).
			Or().
			Intext(number.Local),
	}
}

func getIndividualsDorks(number *Number) []*dorkgen.GoogleSearch {
	return []*dorkgen.GoogleSearch{
		(&dorkgen.GoogleSearch{}).
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("numinfo.net").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("sync.me").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("whocallsyou.de").
			Intext("0" + number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("pastebin.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("whycall.me").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("locatefamily.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("spytox.com").
			Intext(number.Local),
	}
}

func getSocialMediaDorks(number *Number) []*dorkgen.GoogleSearch {
	return []*dorkgen.GoogleSearch{
		(&dorkgen.GoogleSearch{}).
			Site("facebook.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("twitter.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("linkedin.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("instagram.com").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
	}
}

func getReputationDorks(number *Number) []*dorkgen.GoogleSearch {
	return []*dorkgen.GoogleSearch{
		(&dorkgen.GoogleSearch{}).
			Site("whosenumber.info").
			Intext(number.E164).
			Intitle("who called"),
		(&dorkgen.GoogleSearch{}).
			Intitle("Phone Fraud").
			Intext(number.International).
			Or().
			Intext(number.E164).
			Or().
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("findwhocallsme.com").
			Intext(number.E164).
			Or().
			Intext(number.International),
		(&dorkgen.GoogleSearch{}).
			Site("yellowpages.ca").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("phonenumbers.ie").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("who-calledme.com").
			Intext(number.E164),
		(&dorkgen.GoogleSearch{}).
			Site("usphonesearch.net").
			Intext(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("whocalled.us").
			Inurl(number.Local),
		(&dorkgen.GoogleSearch{}).
			Site("quinumero.info").
			Intext(number.Local).
			Or().
			Intext(number.International),
		(&dorkgen.GoogleSearch{}).
			Site("uk.popularphotolook.com").
			Inurl(number.Local),
	}
}

// GoogleSearchScan creates several Google requests to search footprints of
// the number online through Google search.
func GoogleSearchScan(number *Number, formats ...string) []*GoogleSearchDork {
	utils.LoggerService.Infoln("Generating Google search dork requests...")

	dorks := []*dorkgen.GoogleSearch{}
	dorks = append(dorks, getIndividualsDorks(number)...)
	dorks = append(dorks, getSocialMediaDorks(number)...)
	dorks = append(dorks, getReputationDorks(number)...)
	dorks = append(dorks, getDisposableProvidersDorks(number)...)

	results := []*GoogleSearchDork{}

	for _, dork := range dorks {
		for _, f := range formats {
			dork.Or().Intext(f)
		}

		results = append(results, &GoogleSearchDork{
			Number: number.E164,
			Dork:   dork.ToString(),
			URL:    dork.ToURL(),
		})
	}

	return results
}
