package bamboohr

import (
	"errors"
	"os"
	"regexp"
	"strings"

	bamboohr_client "github.com/robertlagrant/bamboohr-client-go"
)

func makeConfig() (*bamboohr_client.Config, error) {
	apiKey, apiKeyPresent := os.LookupEnv("BAMBOOHR_API_KEY")
	tenantName, tenantNamePresent := os.LookupEnv("BAMBOOHR_TENANT")
	includeSalary := true

	if !apiKeyPresent || !tenantNamePresent || apiKey == "" || tenantName == "" {
		return nil, errors.New("Missing configuration. Check BAMBOOHR_API_KEY and BAMBOOHR_TENANT environment variables are set.")
	}

	config := bamboohr_client.Config{ApiKey: apiKey, Tenant: tenantName, IncludeSalary: includeSalary}

	return &config, nil
}

func parsePayRate(payRate string) (string, string, error) {
	r := regexp.MustCompile(`(?P<precurrency>[\D]*)(?P<salary>\d+(?:[\.,]\d\d)?)(?P<postcurrency>.*)`)

	salary := r.ReplaceAllString(payRate, "${salary}")
	preCurrency := r.ReplaceAllString(payRate, "${precurrency}")
	postCurrency := r.ReplaceAllString(payRate, "${postcurrency}")

	preAndPost := strings.TrimSpace(preCurrency) + strings.TrimSpace(postCurrency)

	return salary, preAndPost, nil
}
