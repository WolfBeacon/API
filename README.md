# API
OSH's API service


example post request to `/cms/new`:

```
{
  "ownerid": 12345,
  "name": "testhackathon1",
  "location": {
    "name": "testuni",
    "city": "cambridge",
    "state": "mass",
    "coords": [
      1.1,
      2.2
    ]
  },
  "startdate": 1,
  "enddate": 2,
  "currentstate": 0,
  "prizes": [
    {
      "name": "prize1",
      "description": "prizedesc",
      "sponsor": "osh"
    }
  ],
  "reimbursements": false,
  "busesoffered": false,
  "buslocations": [],
  "sociallinks": [],
  "hardware": [],
  "map": "http://www.google.com",
  "metadata": ""
}
```
