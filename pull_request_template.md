### THANKS FOR TAKING THE TIME TO REVIEW MY CODE <3

# GitHub Issue/Ticket Number

# Description

Please include a summary of the change and which issue is fixed. Please also include relevant motivation and context. List any dependencies that are required for this change.

Fixes # (issue)

## Type of change

Please delete options that are not relevant.

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] This change requires a documentation update

# How Has This Been Tested?

Please describe the tests that you ran to verify your changes. Provide instructions so we can reproduce. Please also list any relevant details for your test configuration

- [ ] Test A
- [ ] Test B

**Test Configuration**:
* Firmware version:
* Hardware:
* Toolchain:
* SDK:

# How To Test

- [ ] cd into intended directory
- [ ] Git clone https://github.com/MandeeRiz/quotes-starter {new file name}
- [ ] Git checkout authorization branch 
- [ ] cd into the folder
- [ ] code .
- [ ] go run gqlgen/server.go
- [ ] got to localhost:8080
- [ ] insert headers `{"x-api-key": "INSERT X-API-KEY PASSWORD"}`

for all queries and mutations copy and paste code below:

query GetQuotes{
  quoteByID(id:"84d1c189-cf52-42f7-a75c-407e9bfae601"){
    id
    quote
    author
  }
randomQuote{
    id
    quote
    author
  }
}

mutation Create{
  createAQuote(input:{quote: "s", author: "m"}){
    id
    author
    quote
  }
}
mutation Delete{
  deleteAquote(id:"ca42c66a-552c-4f7a-ba8b-5350b7bab6f3")
}

Sample ID's to test:
3736cfb7-1569-4cea-b9d7-cc438afa2ab8
980a5645-71ad-41df-939e-893560c7ac16
4589f6c5-8129-4919-b3db-de45edme255b
8090fab3-e853-466a-8e4d-daddf39790ae

Sample ID's to delete:
e6f32497-410d-46f7-b79f-1611e1de042e
50915c19-6600-410a-ba9f-7bfa07f5eef7

# Checklist:

- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published in downstream modules