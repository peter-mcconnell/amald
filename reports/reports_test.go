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
		Global: map[string]string{
			"templatesdir": "tmpl/",
		},
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
		ShowSameState: true,
	}
	Rpt.Cfg = cfg
	// some fake data to use
	yesterday := make(defs.Analysis)
	lastweek := make(defs.Analysis)
	lastmonth := make(defs.Analysis)
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
	Summaries[0] = yesterday
	Summaries[1] = lastweek
	Summaries[2] = lastmonth

	// use one of the test results from yesterday as scanResults
	Rpt.ScanResults = yesterday[0]
}

func ExampleHtmlGenerate() {
	output, _ := Rpt.GenerateHtml(Summaries)
	fmt.Print(output)
	//Output:
	// <html>
	// <head>
	// <style>
	// html, body {
	//       padding:0; margin: 0;
	//       background-color:#2c2828;
	// }
	// body, table.body, h1, h2, h3, h4, h5, h6, p, td {
	//       font-family: "Helvetica", "Arial", sans-serif;
	//       font-weight: normal;
	//       color: #222222;
	// }
	// .httpcode {
	//       font-style: bold;
	//       background: #ffffff;
	// }
	// .httpcode.hc401, .httpcode.hc403, .httpcode.hc407, .httpcode.hc550 {
	//       background: #F2F5A9;
	// }
	// .httpcode.hc200 {
	//       background: #9FF781;
	// }
	// .httpcode.hc301, .httpcode.hc307, .httpcode.hc404, .httpcode.hc410 {
	//       background: #9FF781;
	// }
	// </style>
	// </head>
	// <body style="background-color:#2c2828;">
	// <table width="100%" cellpadding="0" cellspacing="0">
	// <tr>
	//       <td align="center" style="background-color:#2c2828">
	            
	//             <table width="100%" cellpadding="0" cellspacing="0">
	//             <tr>
	//                   <td style="background-color:#757ad8" align="center">
	//                         <table width="600" cellpadding="0" cellspacing="0">
	//                         <tr>
	//                               <td style="background-color:#757ad8">
	//                                     <img src="https://raw.githubusercontent.com/pemcconnell/amald/master/reports/tmpl/email/logo.gif" width="200" height="50" style="margin:0 0 0 10px;" />
	//                               </td>
	//                         </tr>
	//                         </table>
	//                   </td>
	//             </tr>
	//             </table>


	// <table width="600" cellpadding="0" cellspacing="0">
	// 	<tr>
	// 		<td style="background-color:#2c2828">
	// 			<p>&nbsp;</p>
	// 			<h2 style="color:#efefef">Summary</h2>
				
				
	// 			<h2 style="margin:15px 0 0 0; font-style: italic; background: #757ad8; color: #efefef; padding: 5px 0 5px 10px; font-size:14px; line-height:19px; ">&raquo; since yesterday</h2>
	// 			<table width="100%" cellpadding="10" cellspacing="0">
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> removed
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://a" style="font-size:13px; color:#757ad8">http://a</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://b" style="font-size:13px; color:#757ad8">http://b</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://c" style="font-size:13px; color:#757ad8">http://c</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> created
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://d" style="font-size:13px; color:#757ad8">http://d</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://e" style="font-size:13px; color:#757ad8">http://e</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> updated
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://g" style="font-size:13px; color:#757ad8">http://g</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> same
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://l" style="font-size:13px; color:#757ad8">http://l</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 			</table>
				
	// 			<h2 style="margin:15px 0 0 0; font-style: italic; background: #757ad8; color: #efefef; padding: 5px 0 5px 10px; font-size:14px; line-height:19px; ">&raquo; since last week</h2>
	// 			<table width="100%" cellpadding="10" cellspacing="0">
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> removed
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://a" style="font-size:13px; color:#757ad8">http://a</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://b" style="font-size:13px; color:#757ad8">http://b</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://c" style="font-size:13px; color:#757ad8">http://c</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> created
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://d" style="font-size:13px; color:#757ad8">http://d</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://e" style="font-size:13px; color:#757ad8">http://e</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> updated
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://g" style="font-size:13px; color:#757ad8">http://g</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> same
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://l" style="font-size:13px; color:#757ad8">http://l</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 			</table>
				
	// 			<h2 style="margin:15px 0 0 0; font-style: italic; background: #757ad8; color: #efefef; padding: 5px 0 5px 10px; font-size:14px; line-height:19px; ">&raquo; since last month</h2>
	// 			<table width="100%" cellpadding="10" cellspacing="0">
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> removed
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://a" style="font-size:13px; color:#757ad8">http://a</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://b" style="font-size:13px; color:#757ad8">http://b</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://c" style="font-size:13px; color:#757ad8">http://c</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> created
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://d" style="font-size:13px; color:#757ad8">http://d</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://e" style="font-size:13px; color:#757ad8">http://e</a>
	// 								</td>
	// 								<td style="background:#90ea99">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 								</td>
	// 								<td class="httpcode hc401">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 				<tr>
	// 					<td style="background-color:#ECF8FF">
	// 						<h3 style="color:#757ad8; font-size:14px; font-weight:bold;">
								
	// 							<em>status:</em> updated
	// 						</h3>
	// 						<table width="100%" cellpadding="2" cellspacing="1">
	// 							<thead>
	// 								<tr>
	// 									<th align="left">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 									</th>
	// 									<th align="center" width="104">
	// 										<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 									</th>
	// 								</tr>
	// 							</thead>
	// 							<tbody>
								
	// 							<tr>
	// 								<td style="background-color:#ffffff">
	// 									<a href="http://g" style="font-size:13px; color:#757ad8">http://g</a>
	// 								</td>
	// 								<td style="background:#ff6633;">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 								</td>
	// 								<td class="httpcode hc200">
	// 									<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 								</td>
	// 							</tr>
								
	// 							</tbody>
	// 						</table>
	// 					</td>
	// 				</tr>
					
	// 			</table>
				

				
	// 			<h2 style="color:#efefef">Latest Scan Results</h2>
	// 			<h2 style="margin:15px 0 0 0; font-style: italic; background: #757ad8; color: #efefef; padding: 5px 0 5px 10px; font-size:14px; line-height:19px; ">&raquo; latest scan</h2>

	// 			<table width="600" cellpadding="0" cellspacing="0">
	// 				<tr>
	// 					<td style="background-color:#2c2828">

	// 						<table width="100%" cellpadding="10" cellspacing="0">

	// 							<tr>
	// 								<td style="background-color:#ECF8FF">

	// 									<table width="100%" cellpadding="2" cellspacing="1">
	// 										<thead>
	// 											<tr>
	// 												<th align="left">
	// 													<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">url</h4>
	// 												</th>
	// 												<th align="center" width="104">
	// 													<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">locked down</h4>
	// 												</th>
	// 												<th align="center" width="104">
	// 													<h4 style="color:#757ad8; font-size:14px; font-weight:bold; margin:0;">http code</h4>
	// 												</th>
	// 											</tr>
	// 										</thead>
	// 										<tbody>
											
	// 										<tr>
	// 											<td style="background-color:#ffffff">
	// 												<a href="http://a" style="font-size:13px; color:#757ad8">http://a</a>
	// 											</td>
	// 											<td style="background:#ff6633;">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">false</div>
	// 											</td>
	// 											<td class="httpcode hc200">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">200</div>
	// 											</td>
	// 										</tr>
											
	// 										<tr>
	// 											<td style="background-color:#ffffff">
	// 												<a href="http://b" style="font-size:13px; color:#757ad8">http://b</a>
	// 											</td>
	// 											<td style="background:#90ea99">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 											</td>
	// 											<td class="httpcode hc401">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 											</td>
	// 										</tr>
											
	// 										<tr>
	// 											<td style="background-color:#ffffff">
	// 												<a href="http://c" style="font-size:13px; color:#757ad8">http://c</a>
	// 											</td>
	// 											<td style="background:#90ea99">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">true</div>
	// 											</td>
	// 											<td class="httpcode hc401">
	// 												<div style="text-align:center; font-size:13px; padding:4px;">401</div>
	// 											</td>
	// 										</tr>
											
	// 										</tbody>
	// 									</table>

	// 								</td>
	// 							</tr>

	// 						</table>
	// 						<p style="text-align:right; font-size:10px; color:#666">powered by <a href="https://github.com/pemcconnell/amald">amald</a>.</p>
	// 						<p>&nbsp;</p>
	// 					</td>
	// 				</tr>
	// 			</table>

	// 		</td>
	// 	</tr>
	// </table>
}

func ExampleAsciiGenerate() {
	output, _ := Rpt.GenerateAscii(Summaries)
	fmt.Print(output)
	//Output:
	//[ SUMMARIES ]
	//
	// ###########  YESTERDAY  ###########
	//
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
	// ###########  LAST WEEK  ###########
	//
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
	// ###########  LAST MONTH  ###########
	//
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
	//  ~ updated [since last month]
	// +----------+------------+-------------+
	// |   URL    | LOCKEDDOWN | STATUS CODE |
	// +----------+------------+-------------+
	// | http://g | false      |         200 |
	// +----------+------------+-------------+
}
