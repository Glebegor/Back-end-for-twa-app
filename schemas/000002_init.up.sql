CREATE TABLE rooms(
    id VARCHAR(1000) PRIMARY KEY,
    shortId VARCHAR(255),
	roundDuration INT,
	roomDuration INT,
	roomStartDate INT,
	roundStartDate INT
);
CREATE TABLE users_rooms(
    id_room VARCHAR(1000) REFERENCES rooms(id) ON DELETE CASCADE NOT NULL,
    id_user INT,
    id_wallet VARCHAR(1000), 
    score INT
);