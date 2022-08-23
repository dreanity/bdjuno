CREATE TABLE giveaway 
(
	"index" BIGINT NOT NULL PRIMARY KEY,
	"duration" BIGINT NOT NULL,
	"created_at" BIGINT NOT NULL,
	"name" TEXT NOT NULL,
	"completion_height" BIGINT NOT NULL,
	"winning_ticket_numbers" INT[] NOT NULL,
	"prizes" JSONB NOT NULL,
	"status" SMALLINT NOT NULL,
	"ticket_count" INT NOT NULL,
	"randomnessRound" BIGINT 
)