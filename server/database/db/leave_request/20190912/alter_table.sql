ALTER TABLE leave_request ALTER COLUMN id SET DEFAULT NEXTVAL('leave_request_id_seq'::regclass);
