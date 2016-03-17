package ngmeasurements

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kisielk/sqlstruct"
	"gopkg.in/mgo.v2"
	"log"
	"opencoredata.org/ocdJanus/utils"
)

type JanusChemCarbTestcVSW struct {
	Tables []JanusChemCarbTesttable `json:"tables"`
}

type JanusChemCarbTesttable struct {
	URL string     `json:"url"`
	Row []JanusChemCarbTestjanusRow `json:"row"`
}

type JanusChemCarbTestjanusRow struct {
	URL       string           `json:"url"`
	Rownum    int              `json:"rownum"`
	Describes []JanusChemCarbTest `json:"describes"`
}

// make name generic  How to load the body of struct
type JanusChemCarbTest struct {
    Leg                           int64               `json:"Leg" dc_description:"Number identifying the cruise."`
    Site                          int64               `json:"Site" dc_description:"Number identifying the site from which the core was retrieved. A site is the position of a beacon around which holes are drilled."`
    Hole                          string              `json:"Hole" dc_description:"Letter identifying the hole at a site from which a core was retrieved or data was collected."`
    Core                          int64               `json:"Core" dc_description:"Sequential numbers identifying the cores retrieved from a particular hole."`
    Core_type                     string              `json:"Core_type" dc_description:"Letter code identifying the drill bit/coring method used to retrieve the core."`
    Section_number                int64               `json:"Section_number" dc_description:"Sequential number of a section of a core."`
    Section_type                  string              `json:"Section_type" dc_description:"Type of a section (C=Core Catcher, S=Core)."`
    Top_cm                        string              `json:"Top_cm" dc_unit:"cm" dc_unit_descript:"centimeters (cm)" dc_description:"Top value of an interval of a section in cm." xs_type:"float"`
    Bot_cm                        string              `json:"Bot_cm" dc_unit:"cm" dc_unit_descript:"centimeters (cm)" dc_description:"Bottom value of an interval of a section in cm." xs_type:"float"`
    Depth_mbsf                    string              `json:"Depth_mbsf" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf)." xs_type:"float"`
    Depth_mbsf_c                  string              `json:"Depth_mbsf_c" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with compression applied." xs_type:"float"`
    Depth_mbsf_v                  string              `json:"Depth_mbsf_v" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with voids removed." xs_type:"float"`
    Depth_mbsf_vc                 string              `json:"Depth_mbsf_vc" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with voids removed and compression applied." xs_type:"float"`
    Depth_mbsf_e                  string              `json:"Depth_mbsf_e" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with exotics removed." xs_type:"float"`
    Depth_mbsf_ec                 string              `json:"Depth_mbsf_ec" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with exotics removed and compression applied." xs_type:"float"`
    Depth_mbsf_ve                 string              `json:"Depth_mbsf_ve" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with voids and exotics removed." xs_type:"float"`
    Depth_mbsf_vec                string              `json:"Depth_mbsf_vec" dc_unit:"mbsf" dc_unit_descript:"meters below sea floor (mbsf)" dc_description:"Depth in meters below sea floor (mbsf) with voids and exotics removed and compression applied." xs_type:"float"`
    Depth_mcd                     string              `json:"Depth_mcd" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd)." xs_type:"float"`
    Depth_mcd_c                   string              `json:"Depth_mcd_c" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with compression applied." xs_type:"float"`
    Depth_mcd_v                   string              `json:"Depth_mcd_v" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with voids removed." xs_type:"float"`
    Depth_mcd_vc                  string              `json:"Depth_mcd_vc" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with voids removed and compression applied." xs_type:"float"`
    Depth_mcd_e                   string              `json:"Depth_mcd_e" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with exotics removed." xs_type:"float"`
    Depth_mcd_ec                  string              `json:"Depth_mcd_ec" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with exotics removed and compression applied." xs_type:"float"`
    Depth_mcd_ve                  string              `json:"Depth_mcd_ve" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with voids and exotics removed." xs_type:"float"`
    Depth_mcd_vec                 string              `json:"Depth_mcd_vec" dc_unit:"mcd" dc_unit_descript:"meters composite depth" dc_description:"Depth in meters composite depth (mcd) with voids and exotics removed and compression applied." xs_type:"float"`
    Inor_c_wt_pct                 sql.NullString      `json:"Inor_C_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Inorganic Carbon in weight percent (wt%)." xs_type:"float"`
    Caco3_wt_pct                  sql.NullString      `json:"CaCO3_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Calcium Carbonate (CaCO3) in weight percent (wt%)." xs_type:"float"`
    Tot_c_wt_pct                  sql.NullString      `json:"Tot_C_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Total Carbon in weight percent (wt%)." xs_type:"float"`
    Org_c_wt_pct                  sql.NullString      `json:"Org_C_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Organic Carbon in weight percent (wt%)." xs_type:"float"`
    Nit_wt_pct                    sql.NullString      `json:"Nit_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Nitrogen (N) in weight percent (wt%)." xs_type:"float"`
    Sul_wt_pct                    sql.NullString      `json:"Sul_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Sulfur (S) in weight percent (wt%)." xs_type:"float"`
    H_wt_pct                      sql.NullString      `json:"H_wt_pct" dc_unit:"wt%" dc_unit_descript:"weight percent (wt%)" dc_description:"Measurement value of Hydrogen (H) in weight percent (wt%)." xs_type:"float"`
}

func JanusChemCarbTestModel() *JanusChemCarbTest {
	return &JanusChemCarbTest{}
}

// func JSONData(qry string, uri string, filename string) []byte {
func JanusChemCarbTestFunc(qry string, uri string, filename string, database string, collection string, conn *sql.DB, session *mgo.Session) error {

	rows, err := conn.Query(qry)
	if err != nil {
		log.Printf(`Error with "%s": %s`, qry, err)
	}

	allResults := []JanusChemCarbTestjanusRow{}
	i := 1
	for rows.Next() {
		d := []JanusChemCarbTest{}
		var t JanusChemCarbTest
		err := sqlstruct.Scan(&t, rows)
		if err != nil {
			log.Print(err)
		}
		d = append(d, t)
		rowURL := fmt.Sprintf("%s/%s#row=%v", uri, filename, i)
		aRow := JanusChemCarbTestjanusRow{rowURL, i, d}
		allResults = append(allResults, aRow)
		i = i + 1
	}

	theTable := JanusChemCarbTesttable{fmt.Sprintf("%s/%s", uri, filename), allResults}
	tableSet := []JanusChemCarbTesttable{}
	tableSet = append(tableSet, theTable)
	final := JanusChemCarbTestcVSW{tableSet}

	// Optional. Switch the session to a Strong behavior.
	session.SetMode(mgo.Strong, true)
	c := session.DB(database).C(collection)

	err = c.Insert(&final)
	if err != nil {
		log.Printf("Janus func Error %v with %s\n", err, filename)
	}

	log.Printf("File: %s  written", filename)

	jm, _ := json.MarshalIndent(final, "", " ")
	_ = utils.WriteFile(filename, jm)

	// session.Close()
	return nil
}