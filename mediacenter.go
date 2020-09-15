///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/cjsmocjsmo/movgo"
	"github.com/cjsmocjsmo/tvgo"
	"net/http"
	"net/url"
	"os"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

// dBcon is exported because I want it
func dBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("MEDIACENTER_MONGODB_ADDRESS"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

func setHeaders(w http.ResponseWriter) ( http.ResponseWriter ) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	return w
}

func intActionHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var ActionMedia []map[string]string
	b1 := bson.M{"catagory": "Action"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ActionMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(ActionMedia)
}

func intCartoonsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	b1 := bson.M{"catagory": "Cartoons"}
	b2 := bson.M{"_id": 0}
	var CartoonMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&CartoonMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(CartoonMedia)
}

func intComedyHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var ComedyMedia []map[string]string
	b1 := bson.M{"catagory": "Comedy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ComedyMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(ComedyMedia)
}

func intDramaHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var DramaMedia []map[string]string
	b1 := bson.M{"catagory": "Drama"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DramaMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(DramaMedia)
}

func intGodzillaHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var GodzillaMedia []map[string]string
	b1 := bson.M{"catagory": "Godzilla"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&GodzillaMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(GodzillaMedia)
}

func intIndianaJonesHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTrc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var IndianaJonesMedia []map[string]string
	b1 := bson.M{"catagory": "IndianaJones"}
	b2 := bson.M{"_id": 0}
	err := MTrc.Find(b1).Select(b2).All(&IndianaJonesMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(IndianaJonesMedia)
}

func intJohnWayneHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var JohnWayneMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWayne"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWayneMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JohnWayneMedia)
}

func intJurassicParkHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var JurassicParkMedia []map[string]string
	b1 := bson.M{"catagory": "JurassicPark"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JurassicParkMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JurassicParkMedia)
}

func intKingsManHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var KingsmanMedia []map[string]string
	b1 := bson.M{"catagory": "Kingsman"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&KingsmanMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(KingsmanMedia)
}

func intHarryPotterHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var HarryPotterMedia []map[string]string
	b1 := bson.M{"catagory": "HarryPotter"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&HarryPotterMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(HarryPotterMedia)
}

func intMenInBlackHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var MenInBlackMedia []map[string]string
	b1 := bson.M{"catagory": "MenInBlack"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MenInBlackMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(MenInBlackMedia)
}

func intMiscHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var MiscMedia []map[string]string
	b1 := bson.M{"catagory": "Misc"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MiscMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(MiscMedia)
}

func intSciFiHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var SciFiMedia []map[string]string
	b1 := bson.M{"catagory": "SciFi"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SciFiMedia)
	if err != nil {
		log.Println("this is a database lookup error for the InitSciFiHandler function")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(SciFiMedia)
}

func intStarTrekHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var StarTrekMedia []map[string]string
	b1 := bson.M{"catagory": "StarTrek"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarTrekMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(StarTrekMedia)
}

func intStarWarsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var StarWarsMedia []map[string]string
	b1 := bson.M{"catagory": "StarWars"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarWarsMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(StarWarsMedia)
}

func intSuperHerosHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var SuperHerosMedia []map[string]string
	b1 := bson.M{"catagory": "SuperHeros"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SuperHerosMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(SuperHerosMedia)
}

func intTremorsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var TremorsMedia []map[string]string
	b1 := bson.M{"catagory": "Tremors"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TremorsMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(TremorsMedia)
}

func intJohnWickHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var JohnWickMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWick"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWickMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JohnWickMedia)
}

func intPiratesHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var PiratesMedia []map[string]string
	b1 := bson.M{"catagory": "Pirates"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&PiratesMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(PiratesMedia)
}

func intDieHardHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	var DieHardMedia []map[string]string
	b1 := bson.M{"catagory": "DieHard"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DieHardMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(DieHardMedia)
}

func playMediaHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	mf := m["movie"][0]
	ses := dBcon()
	defer ses.Close()
	var MediaInfo map[string]string
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	b1 := bson.M{"movfspath": mf}
	b2 := bson.M{"_id": 0}
	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
	if err != nil {
		fmt.Println(err)
	}
	omxAddr := os.Getenv("MEDIACENTER_OMXPLAYER_ADDRESS")
	u, _ = url.Parse(omxAddr)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("medPath", omxAddr)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	Abody := string(body)
	fmt.Printf("this is mediainfo sent to browser: %s", Abody)
	json.NewEncoder(w).Encode(&MediaInfo)
}

func playMediaReactHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	mf := m["movie"]
	omxAddr := os.Getenv("MEDIACENTER_OMXPLAYER_ADDRESS_REACT")
	u, _ = url.Parse(omxAddr)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("medPath", mf[0])
	u.RawQuery = q.Encode()

	log.Println("this is u.String()")
	log.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	Abody := string(body)
	fmt.Println(Abody)
	json.NewEncoder(w).Encode(mf)
}

//MovSetUpHandler Movupdates the db with newly added music
func MovSetUpHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	val, _ := os.LookupEnv("MEDIACENTER_MovSETUP")
	var exitstatus int
	if val == "0" {
		fmt.Println("Ok is Ok")
		log.Println("MEDIACENTER_MovSETUP environment variable is set, starting MOVIEGO")
		exitstatus = 0
	} else {
		fmt.Println("not OK")
		log.Println("MEDIACENTER_MovSETUP environment variable is not set, starting MovSETUP")
		exitstatus = movgo.MovSetUp()
	}
	json.NewEncoder(w).Encode(exitstatus)
}











// MovUpdateHandler needs exporting because I want it
func MovUpdateHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	movgo.MovUpdate()
	json.NewEncoder(w).Encode("0")
}














//MovDBCountHandler bla bla
func MovDBCountHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("MEDIACENTER").C("MEDIACENTER")
	foo, err := MTc.Count()
	if err != nil {
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(0)
		log.Println(err)
	}
	json.NewEncoder(w).Encode(foo)
}

//MovSetupVariableHandler bla bla
func MovSetupVariableHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	status := os.Getenv("MEDIACENTER_MovSETUP")
	json.NewEncoder(w).Encode(status)
}

// dBcon is exported because I want it
// func dBcon() *mgo.Session {
// 	s, err := mgo.Dial(os.Getenv("TVGOBS_MONGODB_ADDRESS"))
// 	if err != nil {
// 		fmt.Println("Session creation dial error")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Session Connection to db established")
// 	return s
// }

func intSTTVHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var STTVMedia []map[string]string
	b1 := bson.M{"catagory": "STTV", "season": s1}
	b2 := bson.M{"_id": 0}
	fmt.Printf("this is b1 %s", b1)
	errG := MTyc.Find(b1).Select(b2).All(&STTVMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&STTVMedia)
}

func intTNGHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var nextGenerationMedia []map[string]string
	b1 := bson.M{"catagory": "TNG", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&nextGenerationMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&nextGenerationMedia)
}

func intEnterpriseHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var enterpriseMedia []map[string]string
	b1 := bson.M{"catagory": "Enterprise", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&enterpriseMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&enterpriseMedia)
}

func intDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var discoveryMedia []map[string]string
	b1 := bson.M{"catagory": "Discovery", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&discoveryMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&discoveryMedia)
}

func intVoyagerHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var voyagerMedia []map[string]string
	b1 := bson.M{"catagory": "Voyager", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&voyagerMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&voyagerMedia)

}

func intLastShipHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lastshipMedia []map[string]string
	b1 := bson.M{"catagory": "LastShip", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lastshipMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&lastshipMedia)
}

func intOrvilleHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var OrvilleMedia []map[string]string
	b1 := bson.M{"catagory": "Orville", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&OrvilleMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	json.NewEncoder(w).Encode(&OrvilleMedia)
}


func intLostInSpaceHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := dBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var LostInSpaceMedia []map[string]string
	b1 := bson.M{"catagory": "Lost In Space", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&LostInSpaceMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	fmt.Println(&LostInSpaceMedia)
	json.NewEncoder(w).Encode(&LostInSpaceMedia)
}

//TVSetUpHandler Setups the db with newly added music
func TVSetUpHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	val, _ := os.LookupEnv("TVGOBS_SETUP")
	var exitstatus int
	if val == "0" {
		fmt.Println("Ok is Ok")
		fmt.Println("TVGOBS_SETUP environment variable is set, starting TVGO")
		exitstatus = 0
	} else {
		fmt.Println("not OK")
		fmt.Println("TVGOBS_SETUP environment variable is not set, starting SETUP")
		exitstatus = tvgo.TVSetUp()
	}
	json.NewEncoder(w).Encode(exitstatus)
}

//DropTVDataBaseHandler is crap
func DropTVDataBaseHandler(w http.ResponseWriter, r *http.Request) {
	sess := dBcon()
	err := sess.DB("tvgobs").DropDatabase()
	if err != nil {
		fmt.Println(err)
	}
}

//TVDBCountHandler bla bla
func TVDBCountHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	ses := dBcon()
	foo, err := ses.DB("tvgobs").C("tvgobs").Count()
	if err != nil {
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(-0)
		log.Println(err)
	}
	json.NewEncoder(w).Encode(foo)
}

//TVMovTV bla bla
func TVMovTV(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	status := os.Getenv("TVGOBS_MovSETUP")
	json.NewEncoder(w).Encode(status)
}

//TVSetupStatusHandler is exported becasued I want it
func TVSetupStatusHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	status := os.Getenv("TVGOBS_SETUP")
	json.NewEncoder(w).Encode(status)
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/intAction", intActionHandler)
	r.HandleFunc("/intCartoons", intCartoonsHandler)
	r.HandleFunc("/intComedy", intComedyHandler)
	r.HandleFunc("/intDrama", intDramaHandler)
	r.HandleFunc("/intGodzilla", intGodzillaHandler)
	r.HandleFunc("/intHarryPotter", intHarryPotterHandler)
	r.HandleFunc("/intIndianaJones", intIndianaJonesHandler)
	r.HandleFunc("/intMenInBlack", intMenInBlackHandler)
	r.HandleFunc("/intMisc", intMiscHandler)
	r.HandleFunc("/intJohnWayne", intJohnWayneHandler)
	r.HandleFunc("/intJurassicPark", intJurassicParkHandler)
	r.HandleFunc("/intKingsMan", intKingsManHandler)
	r.HandleFunc("/intSciFi", intSciFiHandler)
	r.HandleFunc("/intStarTrek", intStarTrekHandler)
	r.HandleFunc("/intStarWars", intStarWarsHandler)
	r.HandleFunc("/intSuperHeros", intSuperHerosHandler)
	r.HandleFunc("/intTremors", intTremorsHandler)
	r.HandleFunc("/intJohnWick", intJohnWickHandler)
	r.HandleFunc("/intPirates", intPiratesHandler)
	r.HandleFunc("/intDieHard", intDieHardHandler)
	r.HandleFunc("/playMedia", playMediaHandler)
	r.HandleFunc("/playMediaReact", playMediaReactHandler)
	r.HandleFunc("/MovDBCount", MovDBCountHandler)
	// r.HandleFunc("/MovSetupVariable", MovSetupVariableHandler)
	r.HandleFunc("/MovSetUp", MovSetUpHandler)
	r.HandleFunc("/MovUpdate", MovUpdateHandler)


	r.HandleFunc("/intSTTV", intSTTVHandler)
	r.HandleFunc("/intTNG", intTNGHandler)
	r.HandleFunc("/intEnterprise", intEnterpriseHandler)
	r.HandleFunc("/intDiscovery", intDiscoveryHandler)
	r.HandleFunc("/intLastShip", intLastShipHandler)
	r.HandleFunc("/intOrville", intOrvilleHandler)
	r.HandleFunc("/intLostInSpace", intLostInSpaceHandler)
	r.HandleFunc("/intVoyager", intVoyagerHandler)
	r.HandleFunc("/playMedia", playMediaHandler)
	r.HandleFunc("/playMediaReact", playMediaReactHandler)

	r.HandleFunc("/TVSetUp", TVSetUpHandler)
	// need to add UpDate
	
	r.HandleFunc("/DropTVDataBase", DropTVDataBaseHandler)
	r.HandleFunc("/TVDBCount", TVDBCountHandler)
	r.HandleFunc("/TVSetupStatus", TVSetupStatusHandler)
	
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":8888", (r))
}
