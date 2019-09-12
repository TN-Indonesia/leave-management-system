ALTER TABLE leave_request 
   ADD CONSTRAINT leave_request_type_leave_id_fkey
   FOREIGN KEY (type_leave_id) 
   REFERENCES type_leave(id);

ALTER sequence leave_request_id_seq OWNED BY leave_request.id;

ALTER TABLE leave_request ALTER COLUMN id SET DEFAULT NEXTVAL('leave_request_id_seq'::regclass);