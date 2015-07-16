package reports

import (
	"fmt"
	"github.com/pemcconnell/amald/defs"
)

var Summaries = make(defs.Summaries)
var Rpt = Report{}

func init() {
	// fake config to use
	cfg := defs.Config{
		SummaryIntervals: []defs.IntervalSettings{
			defs.IntervalSettings{
				Title:         "yesterday",
				DistanceHours: 24,
				Ansii:         "green+h:black",
			},
			defs.IntervalSettings{
				Title:         "last week",
				DistanceHours: 168,
				Ansii:         "red+h:black",
			},
			defs.IntervalSettings{
				Title:         "last month",
				DistanceHours: 720,
				Ansii:         "blue+h:black",
			},
		},
	}
	Rpt.SetCfg(cfg)
	// some fake data to use
	yesterday := make(defs.Analysis)
	lastweek := yesterday
	lastmonth := yesterday
	yesterday[0] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://a",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://b",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		defs.SiteDefinition{
			Url:            "http://c",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	yesterday[1] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://d",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://e",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	yesterday[2] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://g",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
	}
	yesterday[3] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://l",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastweek[0] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://a",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://b",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		defs.SiteDefinition{
			Url:            "http://c",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastweek[1] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://d",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://e",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastweek[2] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://g",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
	}
	lastweek[3] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://l",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastmonth[0] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://a",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://b",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
		defs.SiteDefinition{
			Url:            "http://c",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastmonth[1] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://d",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
		defs.SiteDefinition{
			Url:            "http://e",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	lastmonth[2] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://g",
			IsLockedDown:   false,
			HttpStatusCode: 200,
		},
	}
	lastmonth[3] = []defs.SiteDefinition{
		defs.SiteDefinition{
			Url:            "http://l",
			IsLockedDown:   true,
			HttpStatusCode: 401,
		},
	}
	Summaries[0] = yesterday
	Summaries[1] = lastweek
	Summaries[2] = lastmonth
}

func ExampleAsciiGenerate() {
	output, _ := Rpt.GenerateAscii(Summaries)
	fmt.Print(output)
	//Output:
	//[ SUMMARIES ]
	//
	//  yesterday
	//  ~ created [since yesterday]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://d | false      |         200 |
	// | http://e | true       |         401 |
	// +----------+------------+-------------+
	//  ~ removed [since yesterday]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://a | false      |         200 |
	// | http://b | true       |         401 |
	// | http://c | true       |         401 |
	// +----------+------------+-------------+
	//  ~ same [since yesterday]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://l | true       |         401 |
	// +----------+------------+-------------+
	//  ~ updated [since yesterday]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://g | false      |         200 |
	// +----------+------------+-------------+
	//
	//  last week
	//  ~ created [since last week]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://d | false      |         200 |
	// | http://e | true       |         401 |
	// +----------+------------+-------------+
	//  ~ removed [since last week]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://a | false      |         200 |
	// | http://b | true       |         401 |
	// | http://c | true       |         401 |
	// +----------+------------+-------------+
	//  ~ same [since last week]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://l | true       |         401 |
	// +----------+------------+-------------+
	//  ~ updated [since last week]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://g | false      |         200 |
	// +----------+------------+-------------+
	//
	//  last month
	//  ~ created [since last month]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://d | false      |         200 |
	// | http://e | true       |         401 |
	// +----------+------------+-------------+
	//  ~ removed [since last month]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://a | false      |         200 |
	// | http://b | true       |         401 |
	// | http://c | true       |         401 |
	// +----------+------------+-------------+
	//  ~ same [since last month]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://l | true       |         401 |
	// +----------+------------+-------------+
	//  ~ updated [since last month]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://g | false      |         200 |
	// +----------+------------+-------------+
}
