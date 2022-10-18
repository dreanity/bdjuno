CREATE TABLE "treasurer"
(
	"address" TEXT NOT NULL PRIMARY KEY,
)

CREATE TABLE "gas_bid" 
(
	"chain" TEXT NOT NULL PRIMARY KEY,
	"number" BIGINT NOT NULL,
)

CREATE TABLE "gas_price"
(
	"token_address" TEXT NOT NULL PRIMARY KEY,
	"chain" TEXT NOT NULL,
	"token_symbol" TEXT NOT NULL,
	"value" TEXT NOT NULL,

	CONSTRAINT FK_gas_price_gas_bid FOREIGN KEY("chain") REFERENCES gas_bid("chain")
)