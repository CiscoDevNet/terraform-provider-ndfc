## Acceptance Tests

Acceptance tests can be run using the following command:
```shell
./run_accept_tests.sh // Run everything
./run_accept_tests.sh TestAcc<Pattern of tests> // Run Specific tests
or
make testacc // Run everything
```

* AT runs on actual NDFC environment
- The testbed settings are expected in a yaml file as seen in [`<root>/testing/testbed.yaml`](../testing/testbed.yaml)
- export `TESTBED_FILE` environment variable to indicate the testbed config to be used for AT
- The provider settings can be overrided by following environment variables
  `NDFC_URL`, `NDFC_USER`, `NDFC_PASSWORD`, `NDFC_DOMAIN`, `NDFC_TIMEOUT` and `NDFC_DOMAIN`
* The fabric names, switch serials etc in the config must match the NDFC being used