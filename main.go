package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CodeItems struct {
	Structname string
	Structdata string
}

func main() {

	Funcnames := []string{"JanusAgeDatapoint", "JanusAgeProfile", "JanusChemCarb", "JanusCoreImage", "JanusCoreSummary",
		"JanusCryomagSection", "JanusDhtApct", "JanusGraSection", "JanusIcpSample", "JanusMadSection",
		"JanusMs2fSection", "JanusMsclSection", "JanusMslSection", "JanusNcrSection", "JanusNgrSection",
		"JanusPaleoImage", "JanusPaleoOccurrence", "JanusPaleoSample", "JanusPrimeDataImage",
		"JanusPwlSection", "JanusPws1Section", "JanusPws2Section", "JanusPws3Section",
		"JanusRscSection", "JanusSample", "JanusSedThinSectionSample", "JanusShearStrengthAvs",
		"JanusShearStrengthPen", "JanusShearStrengthTor", "JanusSmearSlide", "JanusTensorCore",
		"JanusThermalConductivity", "JanusThinSectionImage", "JanusVcdHardRockImage",
		"JanusVcdImage", "JanusVcdStructureImage", "JanusXrdImage", "JanusXrfSample"}

	// loop here....
	for _, entry := range Funcnames {

		qrytmp, err := ioutil.ReadFile("template.txt")
		check(err)

		fmt.Printf("Create from entry %s \n", entry)

		structFile := fmt.Sprintf("./inputStage2/%s.txt", entry)
		structData, err := ioutil.ReadFile(structFile)
		check(err)

		structEntry := CodeItems{Structname: entry, Structdata: string(structData)}

		var buff = bytes.NewBufferString("")
		t, err := template.New("go func template").Parse(string(qrytmp))
		if err != nil {
			log.Printf("go func template creation failed: %s", err)
		}
		err = t.Execute(buff, structEntry)
		qry := string(buff.Bytes())

		// write to a file
		filename := fmt.Sprintf("./output/%s.go", entry)
		f, err := os.Create(filename)
		check(err)
		defer f.Close()

		n2, err := f.Write([]byte(qry))
		check(err)
		fmt.Printf("Wrote %s with size %d\n", filename, n2)
		f.Sync()

	}
}
