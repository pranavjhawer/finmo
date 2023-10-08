1. Top Up Account API
    -Endpoint: POST /accounts/topup
    -Request body: { "currency": "USD", "amount": 100 }
    -Description: This API allows users to top up their account with a specified amount in a given currency.
2. FX Rate API
    -Endpoint: GET /fx-rates
    -Description: This API fetches live FX conversion rates using alpha vantage API.
    -The system generates a quoteId and sends it in the response to the client along with expiry time(now + 20s).
    -Response: { "quoteId": "12345", expiry_at: "12345"}
3. FX Conversion API
    -Endpoint: POST /fx-conversion
    -Request body: { "quoteId": "12345", "fromCurrency": "USD", "toCurrency": "EUR", "amount": 100 }
    -Description: This API performs an FX conversion using the provided quoteId and converts the specified amount from one currency to another.
    -Response: { "convertedAmount": 90.53, "currency": "EUR"}
4. Balance API
    -Endpoint: GET /accounts/balance
    -Description: This API retrieves the balances in all currencies for the user's account.
    -Response: { "balances": { "USD": 1000, "EUR": 500, "GBP": 300 } }


Updated
1. Top Up Account API
    -Endpoint: POST /topup
2. FX Rate API
    -Request body: { "fromCurrency": "USD", "toCurrency": "EUR" }
4. Balance API
    -Endpoint: GET /balance

## Running the server
1. Download/unzip the repo
2. In the server.go directory run : go run server.go


## Sample APIs

1. Top Up Account API
curl --location --request POST '127.0.0.1:1323/v1/account/finmo/topup' \
--header 'Content-Type: application/json' \
--data-raw '{ "currency": "USD", "amount": 50 }'

2. FX Rate API
curl --location --request GET '127.0.0.1:1323/v1/account/finmo/fx-rates' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fromCurrency": "USD",
    "toCurrency": "EUR"
}'

3. FX Conversion API
curl --location --request POST '127.0.0.1:1323/v1/account/finmo/fx-conversion' \
--header 'Content-Type: application/json' \
--data-raw '{
    "quoteId": quote_id,
    "fromCurrency": "USD",
    "toCurrency": "EUR",
    "amount": 10
}'

4. Balance API
curl --location --request GET '127.0.0.1:1323/v1/account/finmo/balance' \
--data-raw ''


Scope of Improvement
- Add authentication for accounts
- Add mutex lock for read/write to DB variables
- Add validation of request
- Move expiration timeout and alphavantage details to config 