INSERT INTO athletes (name, grade, personal_record, events) VALUES
    ('Marcus Johnson', 12, '16:45', '5K, 3200m'),
    ('Ethan Williams', 11, '17:12', '5K, 1600m'),
    ('David Chen', 10, '17:38', '5K'),
    ('James Wilson', 12, '16:59', '5K, 3200m'),
    ('Tyler Brooks', 9, '18:22', '5K');

INSERT INTO meets (name, date, location, description) VALUES
    ('Jones County Invitational', '2026-09-06', 'Jones County Park, Gray GA', 'Season opener hosted by Jones County High School. 5K course through the county park trails.'),
    ('Region 4-AAAA Championship', '2026-10-18', 'Heritage Park, Macon GA', 'Regional qualifying meet for the state championship. Top 10 individuals and top 3 teams advance.'),
    ('State Championship', '2026-11-01', 'Carrollton, GA', 'GHSA Class AAAA state cross country championship at University of West Georgia.');

INSERT INTO results (athlete_id, meet_id, time, place) VALUES
    (1, 1, '16:58', 2),
    (2, 1, '17:30', 5),
    (3, 1, '17:55', 8),
    (4, 1, '17:10', 3),
    (5, 1, '18:45', 15),
    (1, 2, '16:50', 1),
    (4, 2, '17:05', 3),
    (2, 2, '17:28', 6);
