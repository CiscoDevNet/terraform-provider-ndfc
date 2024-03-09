package testing

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"terraform-provider-ndfc/internal/provider/resources/resource_vrf_bulk"
	"text/template"
	"time"
)

func GetTFConfigWithSingleResource(tt string, cfg map[string]string, vrfBulk resource_vrf_bulk.NDFCVrfBulkModel, out **string) {
	x := new(string)
	args := map[string]interface{}{
		"User":     cfg["User"],
		"Password": cfg["Password"],
		"Host":     cfg["Host"],
		"Insecure": cfg["Insecure"],
		"Vrf":      &vrfBulk,
		"RscType":  cfg["RscType"],
		"RscName":  cfg["RscName"],
	}
	root_path, _ := os.Getwd()
	tmpl, err := os.ReadFile(root_path + "/testing/config_scale.gotmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("config").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}
	output := bytes.Buffer{}
	err = t.ExecuteTemplate(&output, "HEADER", args)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(&output, "NDFC", args)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(&output, "NDFC_VRF_RESOURCE", args)
	if err != nil {
		panic(err)
	}

	log.Println(output.String())
	*x = output.String()
	if tmpDir == "" {
		tmpDir, err = os.MkdirTemp("/tmp", "tftest_*")
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.Create(fmt.Sprintf("/%s/%s.tf", tmpDir, tt))
	if err != nil {
		panic(err)
	}
	fp.Write(output.Bytes())
	fp.Close()
	*out = x
}

func GetTFConfigWithMultipleResource(tt string, cfg map[string]string, vrfBulk *[]*resource_vrf_bulk.NDFCVrfBulkModel, out **string) {
	x := new(string)
	args := map[string]interface{}{
		"User":     cfg["User"],
		"Password": cfg["Password"],
		"Host":     cfg["Host"],
		"Insecure": cfg["Insecure"],
		"Vrf":      &vrfBulk,
		"RscType":  cfg["RscType"],
		"RscName":  cfg["RscName"],
	}
	root_path, _ := os.Getwd()
	tmpl, err := os.ReadFile(root_path + "/testing/config_scale.gotmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("config").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}
	output := bytes.Buffer{}

	err = t.ExecuteTemplate(&output, "HEADER", args)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(&output, "NDFC", args)
	if err != nil {
		panic(err)
	}
	rscNames := strings.Split(cfg["RscName"], ",")

	for i := range *vrfBulk {
		args["Vrf"] = &(*vrfBulk)[i]
		args["RscName"] = rscNames[i]
		err = t.ExecuteTemplate(&output, "NDFC_VRF_RESOURCE", args)
		if err != nil {
			panic(err)
		}
	}

	log.Println(output.String())
	*x = output.String()
	if tmpDir == "" {
		ct := time.Now()
		tmpDir = fmt.Sprintf("/tmp/tftest_%s", ct.Format("2006_01_02_15-04-05"))
		err = os.MkdirAll(tmpDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.Create(fmt.Sprintf("/%s/%s.tf", tmpDir, tt))
	if err != nil {
		panic(err)
	}
	fp.Write(output.Bytes())
	fp.Close()
	*out = x
}
