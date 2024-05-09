# Steps to run unit test with mockoon server
1. Install mockoon-cli server
```
npm install -g @mockoon/cli
```
2. Copy your mockoon environment data json file to `/terraform-provide-ndfc/mockoon_data.json`
3. Create unit test functions with prefix `TestUT_`
4. Run `./run_unit_test.sh' to run the unit test cases.
