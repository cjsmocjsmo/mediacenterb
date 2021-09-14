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
	"github.com/cjsmocjsmo/movgo"
	"github.com/cjsmocjsmo/tvgo"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// DBcon is exported because I want it
func DBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("MEDIACENTER_MONGODB_ADDRESS"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

func intActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("intActionHandler start")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ActionMedia []map[string]string
	b1 := bson.M{"catagory": "Action"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ActionMedia)
	if err != nil {
		log.Println("ActionHandler db call error")
		log.Println(err)
	}
        log.Println(ActionMedia)
	json.NewEncoder(w).Encode(ActionMedia)
	count := len(ActionMedia)
	log.Printf("Sent %s files", count)
}

func intCartoonsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("intCartoonsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	b1 := bson.M{"catagory": "Cartoons"}
	b2 := bson.M{"_id": 0}
	var CartoonMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&CartoonMedia)
	if err != nil {
		log.Println("cartoonHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(CartoonMedia)
	count := len(CartoonMedia)
	log.Printf("Sent %s files", count)
}

func intComedyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("intComedyHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ComedyMedia []map[string]string
	b1 := bson.M{"catagory": "Comedy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ComedyMedia)
	if err != nil {
		log.Println("comedyHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(ComedyMedia)
	count := len(ComedyMedia)
	log.Printf("Sent %s files", count)
}

func intDramaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("intDramaHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DramaMedia []map[string]string
	b1 := bson.M{"catagory": "Drama"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DramaMedia)
	if err != nil {
		log.Println("dramaHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(DramaMedia)
	count := len(DramaMedia)
	log.Printf("Sent %s files", count)
}

func intGodzillaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("initGodzillaHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var GodzillaMedia []map[string]string
	b1 := bson.M{"catagory": "Godzilla"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&GodzillaMedia)
	if err != nil {
		log.Println("godzillaHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(GodzillaMedia)
	count := len(GodzillaMedia)
	log.Printf("Sent %s files", count)
}

func intIndianaJonesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("intIndianaJonesHandler started")
	ses := DBcon()
	defer ses.Close()
	MTrc := ses.DB("moviegobs").C("moviegobs")
	var IndianaJonesMedia []map[string]string
	b1 := bson.M{"catagory": "IndianaJones"}
	b2 := bson.M{"_id": 0}
	err := MTrc.Find(b1).Select(b2).All(&IndianaJonesMedia)
	if err != nil {
		log.Println("indianaJonesHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(IndianaJonesMedia)
	count := len(IndianaJonesMedia)
	log.Printf("Sent %s files", count)
}

func intJohnWayneHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("johnWayneHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JohnWayneMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWayne"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWayneMedia)
	if err != nil {
		log.Println("johnwayneHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JohnWayneMedia)
	count := len(JohnWayneMedia)
	log.Printf("Sent %s files", count)
}

func intJurassicParkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JurassicParkHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JurassicParkMedia []map[string]string
	b1 := bson.M{"catagory": "JurassicPark"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JurassicParkMedia)
	if err != nil {
		log.Println("JurassicParkHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JurassicParkMedia)
	count := len(JurassicParkMedia)
	log.Printf("Sent %s files", count)
}

func intKingsManHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("KingsManHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var KingsmanMedia []map[string]string
	b1 := bson.M{"catagory": "Kingsman"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&KingsmanMedia)
	if err != nil {
		log.Println("Kingsman db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(KingsmanMedia)
	count := len(KingsmanMedia)
	log.Printf("Sent %s files", count)
}

func intHarryPotterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HarryPotterHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var HarryPotterMedia []map[string]string
	b1 := bson.M{"catagory": "HarryPotter"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&HarryPotterMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(HarryPotterMedia)
	count := len(HarryPotterMedia)
	log.Printf("Sent %s files", count)
}

func intMenInBlackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MenInBlackHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MenInBlackMedia []map[string]string
	b1 := bson.M{"catagory": "MenInBlack"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MenInBlackMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(MenInBlackMedia)
	count := len(MenInBlackMedia)
	log.Printf("Sent %s files", count)
}

func intMiscHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MiscHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MiscMedia []map[string]string
	b1 := bson.M{"catagory": "Misc"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MiscMedia)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(MiscMedia)
	count := len(MiscMedia)
	log.Printf("Sent %s files", count)
}

func intSciFiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SciFiHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SciFiMedia []map[string]string
	b1 := bson.M{"catagory": "SciFi"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SciFiMedia)
	if err != nil {
		log.Println("InitSciFiHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(SciFiMedia)
	count := len(SciFiMedia)
	log.Printf("Sent %s files", count)
}

func intStarTrekHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("StarTrekHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarTrekMedia []map[string]string
	b1 := bson.M{"catagory": "StarTrek"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarTrekMedia)
	if err != nil {
		log.Println("StarTrekHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(StarTrekMedia)
	count := len(StarTrekMedia)
	log.Printf("Sent %s files", count)
}

func intStarWarsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("StarWarsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarWarsMedia []map[string]string
	b1 := bson.M{"catagory": "StarWars"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarWarsMedia)
	if err != nil {
		log.Printf("StarWarsHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(StarWarsMedia)
	count := len(StarWarsMedia)
	log.Printf("Sent %s files", count)
}

func intSuperHerosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SuperHerosHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SuperHerosMedia []map[string]string
	b1 := bson.M{"catagory": "SuperHeros"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SuperHerosMedia)
	if err != nil {
		log.Printf("SuperHerosHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(SuperHerosMedia)
	count := len(SuperHerosMedia)
	log.Printf("Sent %s files", count)
}

func intTremorsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TremorsHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TremorsMedia []map[string]string
	b1 := bson.M{"catagory": "Tremors"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TremorsMedia)
	if err != nil {
		log.Printf("TremorsHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(TremorsMedia)
	count := len(TremorsMedia)
	log.Printf("Sent %s files", count)
}

func intJohnWickHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JohnWickHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JohnWickMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWick"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWickMedia)
	if err != nil {
		log.Printf("JohnWickHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(JohnWickMedia)
	count := len(JohnWickMedia)
	log.Printf("Sent %s files", count)
}

func intPiratesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PiratesHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var PiratesMedia []map[string]string
	b1 := bson.M{"catagory": "Pirates"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&PiratesMedia)
	if err != nil {
		log.Printf("PiratesHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(PiratesMedia)
	count := len(PiratesMedia)
	log.Printf("Sent %s files", count)
}

func intBruceWillisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("BruceWillisHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DieHardMedia []map[string]string
	b1 := bson.M{"catagory": "BruceWillis"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DieHardMedia)
	if err != nil {
		log.Printf("BruceWillisHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(DieHardMedia)
	count := len(DieHardMedia)
	log.Printf("Sent %s files", count)
}

func intFantasyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("FantasyHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var FantasyMedia []map[string]string
	b1 := bson.M{"catagory": "Fantasy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&FantasyMedia)
	if err != nil {
		log.Printf("FantasyHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(FantasyMedia)
	count := len(FantasyMedia)
	log.Printf("Sent %s files", count)
}

func intRiddickHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RiddickHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var RiddickMedia []map[string]string
	b1 := bson.M{"catagory": "Riddick"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&RiddickMedia)
	if err != nil {
		log.Printf("RiddickHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(RiddickMedia)
	count := len(RiddickMedia)
	log.Printf("Sent %s files", count)
}

func intTomCruizeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TomCruizeHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TCMedia []map[string]string
	b1 := bson.M{"catagory": "TomCruize"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TCMedia)
	if err != nil {
		log.Printf("TomCruizeHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(TCMedia)
	count := len(TCMedia)
	log.Printf("Sent %s files", count)
}

func intXMenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("XMenHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var XMenMedia []map[string]string
	b1 := bson.M{"catagory": "XMen"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&XMenMedia)
	if err != nil {
		log.Printf("XMenHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(XMenMedia)
	count := len(XMenMedia)
	log.Printf("Sent %s files", count)
}

func intDocumentaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DocumentaryHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DocumentaryMedia []map[string]string
	b1 := bson.M{"catagory": "Documentary"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DocumentaryMedia)
	if err != nil {
		log.Println("DocumentaryHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(DocumentaryMedia)
	count := len(DocumentaryMedia)
	log.Printf("Sent %s files", count)
}

func intTheRockHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TheRockHandler started")
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TheRockMedia []map[string]string
	b1 := bson.M{"catagory": "TheRock"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TheRockMedia)
	if err != nil {
		log.Println("TheRockHandler db call error")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(TheRockMedia)
	count := len(TheRockMedia)
	log.Printf("Sent %s files", count)
}

// func playMediaHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("playMediaHandler started")
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		log.Println("playMediaHandler url parse error")
// 		log.Println(err)
// 	}
// 	m, _ := url.ParseQuery(u.RawQuery)
// 	mf := m["movie"][0]
// 	ses := DBcon()
// 	defer ses.Close()
// 	var MediaInfo map[string]string
// 	MTc := ses.DB("moviegobs").C("moviegobs")
// 	b1 := bson.M{"movfspath": mf}
// 	b2 := bson.M{"_id": 0}
// 	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
// 	if err != nil {
// 		log.Println("playMediaHandler db call error")
// 		log.Println(err)
// 	}


// 	omxAddr := os.Getenv("moviegobs_OMXPLAYER_ADDRESS")
// 	u, _ = url.Parse(omxAddr)
// 	q, _ := url.ParseQuery(u.RawQuery)
// 	q.Add("medPath", omxAddr)
// 	u.RawQuery = q.Encode()
// 	resp, err := http.Get(u.String())
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	Abody := string(body)
// 	fmt.Printf("this is mediainfo sent to browser: %s", Abody)
// 	json.NewEncoder(w).Encode(&MediaInfo)
// }

func playMediaReactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	Abody := string(body)
	log.Println(Abody)
	json.NewEncoder(w).Encode(mf)
}

//MovSetUpHandler Movupdates the db with newly added music
// func MovSetUpHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(" started")
// 	val, _ := os.LookupEnv("moviegobs_MovSETUP")
// 	var exitstatus int
// 	if val == "0" {
// 		log.Println("Ok is Ok")
// 		log.Println("moviegobs_MovSETUP environment variable is set, starting MOVIEGO")
// 		exitstatus = 0
// 	} else {
// 		log.Println("not OK")
// 		log.Println("moviegobs_MovSETUP environment variable is not set, starting MovSETUP")
// 		exitstatus = movgo.MovSetUp()
// 	}
// 	json.NewEncoder(w).Encode(exitstatus)
// }

// MovUpdateHandler needs exporting because I want it
func MovUpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MovUpdateHandler started")
	movgo.MovUpdate()
	json.NewEncoder(w).Encode("0")
	log.Println("MovUpdateHandler complete")
}

// //MovDBCountHandler bla bla
// func MovDBCountHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(" started")
// 	ses := DBcon()
// 	defer ses.Close()
// 	MTc := ses.DB("moviegobs").C("moviegobs")
// 	foo, err := MTc.Count()
// 	if err != nil {
// 		json.NewEncoder(w).Encode(0)
// 		log.Println(err)
// 	}
// 	json.NewEncoder(w).Encode(foo)
// }

//MovSetupVariableHandler bla bla
func MovSetupVariableHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MovSetupVariableHandler started")
	status := os.Getenv("MEDIACENTER_SETUP")
	json.NewEncoder(w).Encode(status)
	log.Println("MovSetupVariableHandler completed")
}

func intSTTVHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("STTVHandler started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println("STTVHandler url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var STTVMedia []map[string]string
	b1 := bson.M{"catagory": "STTV", "season": s1}
	b2 := bson.M{"_id": 0}
	fmt.Printf("this is b1 %s", b1)
	errG := MTyc.Find(b1).Select(b2).All(&STTVMedia)
	if errG != nil {
		log.Println("STTVHandler db call error")
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&STTVMedia)
	log.Println("STTVHandler complete")
}

func intTNGHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TNGHandler started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println("TNGHandler url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var nextGenerationMedia []map[string]string
	b1 := bson.M{"catagory": "TNG", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&nextGenerationMedia)
	if errG != nil {
		log.Println("TNGHandler db call error")
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&nextGenerationMedia)
	log.Println("TNGHandler complete")
}

func intEnterpriseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var enterpriseMedia []map[string]string
	b1 := bson.M{"catagory": "Enterprise", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&enterpriseMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&enterpriseMedia)
}

func intDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var discoveryMedia []map[string]string
	b1 := bson.M{"catagory": "Discovery", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&discoveryMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&discoveryMedia)
}

func intVoyagerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var voyagerMedia []map[string]string
	b1 := bson.M{"catagory": "Voyager", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&voyagerMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&voyagerMedia)

}

func intLastShipHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lastshipMedia []map[string]string
	b1 := bson.M{"catagory": "Last Ship", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lastshipMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&lastshipMedia)
}

func intOrvilleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var OrvilleMedia []map[string]string
	b1 := bson.M{"catagory": "Orville", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&OrvilleMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&OrvilleMedia)
}

func intLostInSpaceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var LostInSpaceMedia []map[string]string
	b1 := bson.M{"catagory": "Lost In Space", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&LostInSpaceMedia)
	if errG != nil {
		log.Println(errG)
	}
	log.Println(&LostInSpaceMedia)
	json.NewEncoder(w).Encode(&LostInSpaceMedia)
}

func intPicardHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var PicardMedia []map[string]string
	b1 := bson.M{"catagory": "Picard", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&PicardMedia)
	if errG != nil {
		log.Println(errG)
	}
	log.Println(&PicardMedia)
	json.NewEncoder(w).Encode(&PicardMedia)
}

func intMandalorianHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var MandalorianMedia []map[string]string
	b1 := bson.M{"catagory": "Mandalorian", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&MandalorianMedia)
	if errG != nil {
		log.Println(errG)
	}
	log.Println(&MandalorianMedia)
	json.NewEncoder(w).Encode(&MandalorianMedia)
}

func intAlteredCarbonHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var alteredcarbonMedia []map[string]string
	b1 := bson.M{"catagory": "AlteredCarbon", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&alteredcarbonMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&alteredcarbonMedia)
}

func intLowerDecksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lowerdecksMedia []map[string]string
	b1 := bson.M{"catagory": "LowerDecks", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lowerdecksMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&lowerdecksMedia)
}

func intRaisedByWolvesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var raisedbywolvesMedia []map[string]string
	b1 := bson.M{"catagory": "RaisedByWolves", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&raisedbywolvesMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&raisedbywolvesMedia)
}

func intForAllManKindHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting initForAllManKind")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var forallmankindMedia []map[string]string
	b1 := bson.M{"catagory": "ForAllManKind", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&forallmankindMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&forallmankindMedia)
}

func intAlienWorldsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting initAlienWorlds")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var alienworldsMedia []map[string]string
	b1 := bson.M{"catagory": "AlienWorlds", "season": `01`}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&alienworldsMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&alienworldsMedia)
}

func intWandaVisionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting initWandaVision")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var wandavisionMedia []map[string]string
	b1 := bson.M{"catagory": "WandaVision", "season": `01`}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&wandavisionMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&wandavisionMedia)
}

func intFalconWinterSoldierHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting intFalconWinterSoldier")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lokiMedia []map[string]string
	b1 := bson.M{"catagory": "FalconWinterSoldier", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lokiMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&lokiMedia)
}

func intLokiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting intLoki")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lokiMedia []map[string]string
	b1 := bson.M{"catagory": "Loki", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lokiMedia)
	if errG != nil {
		log.Println(errG)
	}
	log.Println(lokiMedia)
	json.NewEncoder(w).Encode(&lokiMedia)
}



func intWhatIfHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting intWhatIf")
	log.Println(" started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(" url parse error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	log.Println(s1)
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var whatifMedia []map[string]string
	b1 := bson.M{"catagory": "WhatIf", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&whatifMedia)
	if errG != nil {
		log.Println(errG)
	}
	log.Println(whatifMedia)
	json.NewEncoder(w).Encode(&whatifMedia)
}

func intTheBadBatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TheBadBatch started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println("parse url string error")
		log.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("usrl parsequery error")
		log.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var thebadbatchMedia []map[string]string
	b1 := bson.M{"catagory": "TheBadBatch", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&thebadbatchMedia)
	if errG != nil {
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&thebadbatchMedia)
	count := len(thebadbatchMedia)
	log.Printf("Sent %s files", count)
}

func intMastersOfTheUniverseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MastersOfTheUniverse started")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println("url parse error")
		log.Println(err)
	}
	_, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		log.Println("MastersOfTheUniverse url parse querry error")
		log.Println(eff)
	}
	// s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var mastersOfTheUniverseMedia []map[string]string
	b1 := bson.M{"catagory": "MastersOfTheUniverse", "season": `01`}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&mastersOfTheUniverseMedia)
	if errG != nil {
		log.Println("mastersOfTheUniverse db call error")
		log.Println(errG)
	}
	json.NewEncoder(w).Encode(&mastersOfTheUniverseMedia)
	count := len(mastersOfTheUniverseMedia)
	log.Printf("Sent %s files", count)
}

// MovUpdateHandler needs exporting because I want it
func TVUpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TVUpdateHandler started")
	tvgo.TVUpdate()
	json.NewEncoder(w).Encode("0")
}

//TVSetupStatusHandler is exported becasued I want it
func TVSetupStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TVSetupStatusHandler started")
	status := os.Getenv("MOVIECENTER_TVGO_SETUP")
	json.NewEncoder(w).Encode(status)
	log.Printf("TVSetupStatusHandler is %s", TVSetupStatusHandler)
}

func setupLogging() {
	logfile := os.Getenv("MEDIACENTER_LOG_PATH")
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Logging started \n")
}

func init() {
	setupLogging()
	movgo.MovSetUp()
	tvgo.TVSetUp()

}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/intAction", intActionHandler)
	r.HandleFunc("/intCartoons", intCartoonsHandler)
	r.HandleFunc("/intComedy", intComedyHandler)
	r.HandleFunc("/intDrama", intDramaHandler)
	r.HandleFunc("/intFantasy", intFantasyHandler)
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
	r.HandleFunc("/intBruceWillis", intBruceWillisHandler)
	r.HandleFunc("/intRiddick", intRiddickHandler)
	r.HandleFunc("/intTomCruize", intTomCruizeHandler)
	r.HandleFunc("/intXMen", intXMenHandler)
	r.HandleFunc("/intDocumentary", intDocumentaryHandler)
	r.HandleFunc("/intTheRock", intTheRockHandler)
	// r.HandleFunc("/MovDBCount", MovDBCountHandler)
	// r.HandleFunc("/MovSetUp", MovSetUpHandler)
	r.HandleFunc("/MovUpdate", MovUpdateHandler)

	
	//TVGOBS_SETUP
	r.HandleFunc("/intFalconWinterSoldier", intFalconWinterSoldierHandler)
	r.HandleFunc("/intAlteredCarbon", intAlteredCarbonHandler)
	r.HandleFunc("/intAlienWorlds", intAlienWorldsHandler)
	r.HandleFunc("/intDiscovery", intDiscoveryHandler)
	r.HandleFunc("/intEnterprise", intEnterpriseHandler)
	r.HandleFunc("/intForAllManKind", intForAllManKindHandler)
	r.HandleFunc("/intLastShip", intLastShipHandler)
	r.HandleFunc("/intLostInSpace", intLostInSpaceHandler)
	r.HandleFunc("/intLowerDecks", intLowerDecksHandler)
	r.HandleFunc("/intMandalorian", intMandalorianHandler)
	r.HandleFunc("/intOrville", intOrvilleHandler)
	r.HandleFunc("/intPicard", intPicardHandler)
	r.HandleFunc("/intRaisedByWolves", intRaisedByWolvesHandler)
	// r.HandleFunc("/intSpaceTime", intSpaceTimeHandler)
	// r.HandleFunc("/intSeanCarrol", intSeanCarrolHandler)
	r.HandleFunc("/intSTTV", intSTTVHandler)
	r.HandleFunc("/intTNG", intTNGHandler)
	r.HandleFunc("/intVoyager", intVoyagerHandler)
	r.HandleFunc("/intWandaVision", intWandaVisionHandler)
	r.HandleFunc("/intMastersOfTheUniverse", intMastersOfTheUniverseHandler)
	r.HandleFunc("/intLoki", intLokiHandler)
	r.HandleFunc("/intTheBadBatch", intTheBadBatchHandler)
	r.HandleFunc("/intWhatIf", intWhatIfHandler) 

	
	r.HandleFunc("/TVUpdate", TVUpdateHandler)


	// r.HandleFunc("/TVSetUp", TVSetUpHandler)
	// r.HandleFunc("/TVDBCount", TVDBCountHandler)
	// r.HandleFunc("/DropTVDataBase", DropTVDataBaseHandler)

	// r.HandleFunc("/playMedia", playMediaHandler)
	r.HandleFunc("/playMediaReact", playMediaReactHandler)
	
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), 
		handlers.AllowedOrigins([]string{"*"}))(r))
}


// //TVSetUpHandler Setups the db with newly added music
// func TVSetUpHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(" started")
// 	val, _ := os.LookupEnv("TVGOBS_SETUP")
// 	var exitstatus int
// 	if val == "0" {
// 		log.Println("Ok is Ok")
// 		log.Println("TVGOBS_SETUP environment variable is set, starting TVGO")
// 		exitstatus = 0
// 	} else {
// 		log.Println("not OK")
// 		log.Println("TVGOBS_SETUP environment variable is not set, starting SETUP")
// 		exitstatus = tvgo.TVSetUp()
// 	}
// 	json.NewEncoder(w).Encode(exitstatus)
// }
// //DropTVDataBaseHandler is crap
// func DropTVDataBaseHandler(w http.ResponseWriter, r *http.Request) {
// 	sess := DBcon()
// 	err := sess.DB("tvgobs").DropDatabase()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// //TVDBCountHandler bla bla
// func TVDBCountHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("TVDBCountHandler started")
// 	ses := DBcon()
// 	foo, err := ses.DB("tvgobs").C("tvgobs").Count()
// 	if err != nil {
// 		json.NewEncoder(w).Encode(-0)
// 		log.Println(err)
// 	}
// 	json.NewEncoder(w).Encode(foo)
// }