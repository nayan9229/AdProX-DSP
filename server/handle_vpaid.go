package server

import "net/http"

func vpaid(w http.ResponseWriter, r *http.Request) {
	xml := `<VAST version="3.0" xmlns:xs="http://www.w3.org/2001/XMLSchema"><Ad id="1301"><InLine><AdSystem version="4.0">InfyTV</AdSystem><AdTitle>geniun video ad</AdTitle><Pricing model="cpm" currency="USD"><![CDATA[ 0.1 ]]></Pricing><Creatives><Creative id="cr_1301" sequence="1"><Linear><Duration>00:00:30</Duration><AdParameters><![CDATA[{"publisher_id": "1229", "endpoint_id": "1301", "gen_id": "p_1800flowers", "data-ad": false, "data-c-id": "1229", "data-id": "p_1800flowers"}]]></AdParameters><MediaFiles><MediaFile id="media_1301" apiFramework="VPAID" type="application/javascript">https://infytvcode.github.io/vpaid/vpaid.js</MediaFile></MediaFiles></Linear></Creative></Creatives></InLine></Ad></VAST>`
	w.Write([]byte(xml))
}
