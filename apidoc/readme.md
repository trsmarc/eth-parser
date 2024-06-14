# eth-parser

Available format:
- OpenAPI
- Swagger

## End-point: Subscribe
### Method: POST
>```
>/subscribe
>```
### Body (**raw**)

```json
{
    "address": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"
}
```

### Response: 200
```json
{
    "message": "Subscribed successfully"
}
```

### Response: 400
```json
address already subscribed

```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Transactions
### Method: GET
>```
>/transactions?address=0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5
>```
### Query Params

|Param|value|
|---|---|
|address|0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5|


### Response: 200
```json
{
    "total": 9,
    "transactions": [
        {
            "Hash": "0x57834a4cff528eb28047c7c6ea563d1ca4ab48e33ccb202314bda06871b4d999",
            "Block": 20083815,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x673D2EBe4B6BAA946345C7b1F8d3Cc2FfB3429Bf",
            "Value": 757592515680231
        },
        {
            "Hash": "0xd91cf16418a67f39de81a89037fe983121388f455b5b123fe4280c9d6bb76960",
            "Block": 20083815,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0xE94f1fa4F27D9d288FFeA234bB62E1fBC086CA0c",
            "Value": 520195376512174153
        },
        {
            "Hash": "0x22595bfa0e1646334e912921fdb14fb66430184cc63c1305d7b11ed06cd4fa8b",
            "Block": 20083816,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0xeE4F22594EbEac5b2D94e03d7b43b266c945C2C6",
            "Value": 47715160639726325
        },
        {
            "Hash": "0x0825a505825d6fca162214bbe2a7b58b5664b228d67fc645fc24bd11d81b924b",
            "Block": 20083818,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x673D2EBe4B6BAA946345C7b1F8d3Cc2FfB3429Bf",
            "Value": 367370170322553
        },
        {
            "Hash": "0x9fa33c9d6272abc5864215e01e84191d0bf941ff9886328d55f4f35d6f61f6f0",
            "Block": 20083818,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263",
            "Value": 560484777143059122
        },
        {
            "Hash": "0x54f286e44a5fd9d8f9244fcd4ae10870f82ad3d821478acbd1a6576656e5c1e0",
            "Block": 20083820,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x388C818CA8B9251b393131C08a736A67ccB19297",
            "Value": 80822635783348288
        },
        {
            "Hash": "0x741444ae81808b64da65a0ec7f2817105ca677af6015477849d2d485e6a66a91",
            "Block": 20083821,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x15D59433Aea693cDE0E82793Edd3b6F3d5E24E22",
            "Value": 47755544193859562
        },
        {
            "Hash": "0x8f15db3ca062705759c6ad9e7a13d2bf5de3caa913df406db9f72423b4048ed5",
            "Block": 20083822,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x00669Df8c991b8eFf3285671927e7320a33F7291",
            "Value": 6591897028779115
        },
        {
            "Hash": "0xff024d28914be16fbe9aa8087312838945a3568a297ffd67fa3b0c0dec111202",
            "Block": 20083822,
            "From": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
            "To": "0x388C818CA8B9251b393131C08a736A67ccB19297",
            "Value": 70871357721586999
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Latest Block
### Method: GET
>```
>/block
>```
### Response: 200
```json
{
    "blockNumber": 20083828
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Subscriber
### Method: GET
>```
>/subscriber
>```
### Response: 200
```json
{
    "addresses": [
        "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"
    ],
    "total": 1
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
