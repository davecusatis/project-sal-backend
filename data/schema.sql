CREATE TABLE ChannelConfig
(
    id SERIAL,
    customIcons BOOL,
    slotName VARCHAR
);

CREATE TABLE ChannelScores
(
    id SERIAL,
    score INT,
    recordedAt timestamptz NOT NULL DEFAULT now(),
    userId VARCHAR,
    channelId VARCHAR,
    bitsUsed INT
);

CREATE TABLE UserAlerts
(
    id SERIAL,
    opaqueId VARCHAR,
    alertsEnabled BOOL
);
