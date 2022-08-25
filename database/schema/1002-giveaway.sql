CREATE TABLE giveaway 
(
	"index" INT NOT NULL PRIMARY KEY,
	"duration" BIGINT NOT NULL,
	"created_at" BIGINT NOT NULL,
	"name" TEXT NOT NULL,
	"completion_height" BIGINT NOT NULL,
	"winning_ticket_numbers" INT[] NOT NULL,
	"prizes" JSONB NOT NULL,
	"status" SMALLINT NOT NULL,
	"ticket_count" INT NOT NULL,
	"randomness_round" BIGINT,

	CONSTRAINT FK_giveaways_randomness FOREIGN KEY("randomness_round") REFERENCES randomness("round")
)

CREATE TABLE ticket 
(
	"index" INT NOT NULL PRIMARY KEY,
	"giveaway_id" INT,
	"participant_id" TEXT NOT NULL,
	"participant_name" TEXT NOT NULL,

	CONSTRAINT FK_tickets_giveaway FOREIGN KEY("giveaway_id") REFERENCES giveaway("index")
)