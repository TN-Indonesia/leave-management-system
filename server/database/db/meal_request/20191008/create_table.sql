CREATE TABLE IF NOT EXISTS meal_request
(
  "id" int PRIMARY KEY NOT NULL,
  
  --Request's Detail
  "requestor_id" int NOT NULL,
  "other_requestor_id" text,
  "supervisor_id" int NOT NULL,
  "amount" float NOT NULL,
  "brief_description" text NOT NULL,
  "notes" text,
  "receipt_upload_path" text NOT NULL,
  "request_date" timestamp NOT NULL,
  "status" text NOT NULL,
  
  --Approval's detail
  "reject_reason" text,
  "action_date" timestamp,
  

  "created_at" timestamp NOT NULL default CURRENT_TIMESTAMP,
  "updated_at" timestamp
);