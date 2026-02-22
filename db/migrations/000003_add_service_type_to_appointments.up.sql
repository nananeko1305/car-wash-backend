CREATE TYPE service_type AS ENUM ('washing' , 'vacuuming', 'washing_and_vacuming', 'deep_cleaning');

ALTER TABLE appointments ADD COLUMN service_type service_type;