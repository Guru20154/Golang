package main

import (
	"fmt"
	"net/url"
)

const murl = "https://www.linkedin.com/jobs/search/?alertAction=viewjobs&currentJobId=4072151400&distance=25&f_TPR=a1731011077-&geoId=105214831&keywords=data%20analyst&origin=JOB_ALERT_IN_APP_NOTIFICATION&originToLandingJobPostings=4072151400&savedSearchId=1732463674&sortBy=R"

func main() {
	fmt.Println("Url Handling")
	// fmt.Println(murl)

	result,err := url.Parse(murl)

	if err!=nil{
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("Type of qparmas %T\n",qparams)
	// fmt.Println(qparams)

	for _,val := range qparams{
		fmt.Println("Param is: ", val)
	}

	partsofURL := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawPath: "user=guru",
	}
	anoth := partsofURL.String()
	fmt.Println(anoth) 
}
