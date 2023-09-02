ALTER TABLE users 
ADD CONSTRAINT check_email CHECK (email <> ''),
ADD CONSTRAINT check_contact CHECK (contact <> '');
