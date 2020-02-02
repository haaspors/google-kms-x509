package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"

	"github.com/ericnorris/google-kms-x509/csr"
	"github.com/spf13/cobra"
)

var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

func init() {
	var key string
	var commonName string
	var country string
	var province string
	var locality string
	var organization string
	var organizationalUnit string
	var emailAddress string

	var csrCmd = &cobra.Command{
		Use:   "csr",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			subject := pkix.Name{
				CommonName: commonName,
			}

			if country != "" {
				subject.Country = []string{country}
			}

			if province != "" {
				subject.Province = []string{province}
			}

			if locality != "" {
				subject.Locality = []string{locality}
			}

			if organization != "" {
				subject.Organization = []string{organization}
			}

			if organizationalUnit != "" {
				subject.OrganizationalUnit = []string{organizationalUnit}
			}

			if emailAddress != "" {
				subject.ExtraNames = []pkix.AttributeTypeAndValue{
					{
						Type:  oidEmailAddress,
						Value: emailAddress,
					},
				}
			}

			fmt.Println(csr.Generate(key, subject))
		},
	}

	csrCmd.Flags().StringVarP(
		&key, "key", "k", "", "Google KMS key to use for signature",
	)

	csrCmd.Flags().StringVarP(
		&commonName, "common-name", "", "", "Common Name to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&country, "country", "", "", "Country to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&province, "province", "", "", "Province to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&locality, "locality", "", "", "Locality to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&organization, "organization", "", "", "Organization to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&organizationalUnit, "organizationalUnit", "", "", "Organizational Unit to use for CSR",
	)

	csrCmd.Flags().StringVarP(
		&emailAddress, "emailAddress", "", "", "Email Address to use for CSR",
	)

	csrCmd.MarkFlagRequired("key")
	csrCmd.MarkFlagRequired("common-name")

	rootCmd.AddCommand(csrCmd)
}
